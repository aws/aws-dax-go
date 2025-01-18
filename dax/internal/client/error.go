/*
  Copyright 2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.

  Licensed under the Apache License, Version 2.0 (the "License").
  You may not use this file except in compliance with the License.
  A copy of the License is located at

      http://www.apache.org/licenses/LICENSE-2.0

  or in the "license" file accompanying this file. This file is distributed
  on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
  express or implied. See the License for the specific language governing
  permissions and limitations under the License.
*/

package client

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/aws/aws-dax-go/dax/internal/cbor"
	"github.com/aws/aws-dax-go/dax/internal/lru"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/smithy-go"
)

const (
	ErrCodeNotImplemented      = "NotImplemented"
	ErrCodeValidationException = "ValidationException"
	ErrCodeServiceUnavailable  = "ServiceUnavailable"
	ErrCodeUnknown             = "Unknown"
	ErrCodeThrottlingException = "ThrottlingException"
)

type daxError interface {
	smithy.APIError
	CodeSequence() []int
	RequestID() string
	StatusCode() int
}

type daxRequestFailure struct {
	*smithy.GenericAPIError
	codes      []int
	requestID  string
	statusCode int
}

var _ daxError = (*daxRequestFailure)(nil)

type daxTransactionCanceledFailure struct {
	*daxRequestFailure
	cancellationReasonCodes []*string
	cancellationReasonMsgs  []*string
	cancellationReasonItems []byte
	cancellationReasons     []types.CancellationReason
}

func newDaxRequestFailure(codes []int, errorCode, message, requestId string, statusCode int) *daxRequestFailure {
	return &daxRequestFailure{
		GenericAPIError: &smithy.GenericAPIError{
			Code:    errorCode,
			Message: message,
			Fault:   smithy.FaultServer,
		},
		codes:      codes,
		requestID:  requestId,
		statusCode: statusCode,
	}
}

func newDaxTransactionCanceledFailure(codes []int, errorCode, message, requestId string, statusCode int,
	cancellationReasonCodes, cancellationReasonMsgs []*string, cancellationReasonItems []byte) *daxTransactionCanceledFailure {
	return &daxTransactionCanceledFailure{
		daxRequestFailure:       newDaxRequestFailure(codes, errorCode, message, requestId, statusCode),
		cancellationReasonCodes: cancellationReasonCodes,
		cancellationReasonMsgs:  cancellationReasonMsgs,
		cancellationReasonItems: cancellationReasonItems,
	}
}

func (f *daxRequestFailure) CodeSequence() []int {
	return f.codes
}

func (f *daxRequestFailure) RequestID() string {
	return f.requestID
}

func (f *daxRequestFailure) StatusCode() int {
	return f.statusCode
}

func (f *daxRequestFailure) recoverable() bool {
	return len(f.codes) > 0 && f.codes[0] == 2
}

func (f *daxRequestFailure) authError() bool {
	return len(f.codes) > 3 && (f.codes[1] == 23 && f.codes[2] == 31 &&
		(f.codes[3] == 32 || f.codes[3] == 33 || f.codes[3] == 34))
}

func decodeError(reader *cbor.Reader) (error, error) {
	length, err := reader.ReadArrayLength()
	if err != nil {
		return nil, err
	}
	if length == 0 {
		return nil, nil
	}

	codes := make([]int, length)
	for i := 0; i < length; i++ {
		codes[i], err = reader.ReadInt()
		if err != nil {
			return nil, err
		}
	}

	msg, err := reader.ReadString()
	if err != nil {
		return nil, err
	}

	var requestId, errorCode string
	var statusCode int
	var cancellationReasonCodes, cancellationReasonMsgs []*string
	var cancellationReasonItems []byte
	hdr, err := reader.PeekHeader()
	if err != nil {
		return nil, err
	}
	if hdr == cbor.Nil {
		if err := reader.ReadNil(); err != nil {
			return nil, err
		}
	} else {
		length, err = reader.ReadArrayLength()
		if err != nil {
			return nil, err
		}
		if (length < 3) || (length > 4) {
			return nil, &smithy.DeserializationError{Err: fmt.Errorf("expected 3 or 4 elements for error info, got %d", length)}
		}
		if hdr, err = reader.PeekHeader(); err != nil {
			return nil, err
		} else if hdr == cbor.Nil {
			if err := reader.ReadNil(); err != nil {
				return nil, err
			}
		} else if requestId, err = reader.ReadString(); err != nil {
			return nil, err
		}

		if hdr, err = reader.PeekHeader(); err != nil {
			return nil, err
		} else if hdr == cbor.Nil {
			if err := reader.ReadNil(); err != nil {
				return nil, err
			}
		} else if errorCode, err = reader.ReadString(); err != nil {
			return nil, err
		}

		if hdr, err = reader.PeekHeader(); err != nil {
			return nil, err
		} else if hdr == cbor.Nil {
			if err := reader.ReadNil(); err != nil {
				return nil, err
			}
		} else if statusCode, err = reader.ReadInt(); err != nil {
			return nil, err
		}

		if length == 4 {
			arrLen, err := reader.ReadArrayLength()
			if err != nil {
				return nil, err
			}
			if arrLen%3 != 0 {
				return nil, &smithy.DeserializationError{Err: fmt.Errorf("error found when parsing CancellationReasons")}
			}
			cancellationReasonsLen := arrLen / 3
			cancellationReasonCodes = make([]*string, cancellationReasonsLen)
			cancellationReasonMsgs = make([]*string, cancellationReasonsLen)
			itemsBuf := bytes.Buffer{}
			for i := 0; i < cancellationReasonsLen; i++ {
				if consumed, err := consumeNil(reader); err != nil {
					return nil, err
				} else if !consumed {
					s, err := reader.ReadString()
					cancellationReasonCodes[i] = aws.String(s)
					if err != nil {
						return nil, err
					}
				}
				if consumed, err := consumeNil(reader); err != nil {
					return nil, err
				} else if !consumed {
					s, err := reader.ReadString()
					cancellationReasonMsgs[i] = aws.String(s)
					if err != nil {
						return nil, err
					}
				}
				if consumed, err := consumeNil(reader); err != nil {
					return nil, err
				} else if !consumed {
					if err := reader.ReadRawBytes(&itemsBuf); err != nil {
						return nil, err
					}
				} else {
					itemsBuf.WriteByte(byte(cbor.Nil))
				}
			}
			cancellationReasonItems = itemsBuf.Bytes()
		}
	}

	if statusCode == 0 {
		statusCode = inferStatusCode(codes)
	}

	// user or server error
	if cancellationReasonCodes != nil && len(cancellationReasonCodes) > 0 {
		return newDaxTransactionCanceledFailure(codes, errorCode, msg, requestId, statusCode,
			cancellationReasonCodes, cancellationReasonMsgs, cancellationReasonItems), nil
	}
	return newDaxRequestFailure(codes, errorCode, msg, requestId, statusCode), nil
}

// convertDAXError converts DAX error to specific error type based on error code sequence returned from server.
func convertDaxError(e daxError) error {
	codes := e.CodeSequence()
	if len(codes) < 2 {
		return e
	}
	switch codes[1] {
	case 23:
		if len(codes) > 2 {
			switch codes[2] {
			case 24:
				return &types.ResourceNotFoundException{
					Message: aws.String(e.Error()),
				}
			case 35:
				return &types.ResourceInUseException{
					Message: aws.String(e.Error()),
				}
			}
		}
	case 37:
		if len(codes) > 3 {
			switch codes[3] {
			case 39:
				if len(codes) > 4 {
					switch codes[4] {
					case 40:
						return &types.ProvisionedThroughputExceededException{
							Message: aws.String(e.Error()),
						}
					case 41:
						return &types.ResourceNotFoundException{
							Message: aws.String(e.Error()),
						}
					case 43:
						return &types.ConditionalCheckFailedException{
							Message: aws.String(e.Error()),
						}
					case 45:
						return &types.ResourceInUseException{
							Message: aws.String(e.Error())}
					case 46:
						// there's no dynamodb.ValidationException type
						return &smithy.GenericAPIError{
							Code:    ErrCodeValidationException,
							Message: e.Error(),
							Fault:   smithy.FaultServer,
						}
					case 47:
						return &types.InternalServerError{
							Message: aws.String(e.Error()),
						}
					case 48:
						return &types.ItemCollectionSizeLimitExceededException{
							Message: aws.String(e.Error()),
						}
					case 49:
						return &types.LimitExceededException{
							Message: aws.String(e.Error()),
						}
					case 50:
						// there's no dynamodb.ThrottlingException type
						return &smithy.GenericAPIError{
							Code:    ErrCodeThrottlingException,
							Message: e.Error(),
							Fault:   smithy.FaultServer,
						}
					case 57:
						return &types.TransactionConflictException{
							Message: aws.String(e.Error()),
						}
					case 58:
						tcFailure, ok := e.(*daxTransactionCanceledFailure)
						if ok {
							return &types.TransactionCanceledException{
								Message:             aws.String(e.Error()),
								CancellationReasons: tcFailure.cancellationReasons,
							}
						} else {
							return &types.TransactionCanceledException{
								Message: aws.String(e.Error()),
							}
						}
					case 59:
						return &types.TransactionInProgressException{
							Message: aws.String(e.Error()),
						}
					case 60:
						return &types.IdempotentParameterMismatchException{
							Message: aws.String(e.Error()),
						}
					}
				}
			case 44:
				return &smithy.GenericAPIError{
					Code:    ErrCodeNotImplemented,
					Message: e.Error(),
					Fault:   smithy.FaultServer,
				}
			}
		}
	}
	return &smithy.GenericAPIError{
		Code:    ErrCodeUnknown,
		Message: e.Error(),
		Fault:   smithy.FaultServer,
	}
}

func decodeTransactionCancellationReasons(ctx context.Context, failure *daxTransactionCanceledFailure,
	keys []map[string]types.AttributeValue, attrListIdToNames *lru.Lru) ([]types.CancellationReason, error) {
	inputL := len(keys)
	outputL := len(failure.cancellationReasonCodes)
	if inputL != outputL {
		return nil, &smithy.DeserializationError{Err: errors.New("cancellation reasons must be the same length as transact items in the request")}
	}
	reasons := make([]types.CancellationReason, outputL)
	r := cbor.NewReader(bytes.NewReader(failure.cancellationReasonItems))
	for i := 0; i < outputL; i++ {
		reason := types.CancellationReason{}
		reason.Code = failure.cancellationReasonCodes[i]
		reason.Message = failure.cancellationReasonMsgs[i]
		if consumed, err := consumeNil(r); err != nil {
			return nil, err
		} else if !consumed {
			item, err := decodeNonKeyAttributes(ctx, r, attrListIdToNames, nil)
			if err != nil {
				return nil, err
			}
			if item != nil {
				for k, v := range keys[i] {
					item[k] = v
				}
			}
			reason.Item = item
		}
		reasons[i] = reason
	}
	return reasons, nil
}

func inferStatusCode(codes []int) int {
	if len(codes) == 0 {
		return 0
	}
	/*
		1. Retryable server error.
		2. Recoverable failures in cluster. Retry after recovery.
		3. Unretryable server error.
		4. Client error.
	*/
	if codes[0] == 4 {
		return 400
	}
	return 500
}

func isIOError(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, context.DeadlineExceeded) ||
		errors.Is(err, context.Canceled) ||
		errors.Is(err, io.EOF) ||
		strings.Contains(strings.ToLower(err.Error()), "network error") {
		return true
	}

	var netErr net.Error
	if errors.As(err, &netErr) {
		return true
	}
	return false
}
