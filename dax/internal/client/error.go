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
	"fmt"
	"net"

	"github.com/aws/aws-dax-go/dax/internal/cbor"
	"github.com/aws/aws-dax-go/dax/internal/lru"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	ErrCodeNotImplemented      = "NotImplemented"
	ErrCodeValidationException = "ValidationException"
	ErrCodeServiceUnavailable  = "ServiceUnavailable"
	ErrCodeUnknown             = "Unknown"
	ErrCodeThrottlingException = "ThrottlingException"
)

type daxError interface {
	awserr.RequestFailure
	CodeSequence() []int
}

type daxRequestFailure struct {
	awserr.RequestFailure
	codes []int
}

type daxTransactionCanceledFailure struct {
	*daxRequestFailure
	cancellationReasonCodes []*string
	cancellationReasonMsgs  []*string
	cancellationReasonItems []byte
	cancellationReasons     []*dynamodb.CancellationReason
}

func newDaxRequestFailure(codes []int, errorCode, message, requestId string, statusCode int) *daxRequestFailure {
	return &daxRequestFailure{
		RequestFailure: awserr.NewRequestFailure(awserr.New(errorCode, message, nil), statusCode, requestId),
		codes:          codes,
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

func (f *daxRequestFailure) recoverable() bool {
	return len(f.codes) > 0 && f.codes[0] == 2
}

func (f *daxRequestFailure) authError() bool {
	return len(f.codes) > 3 && (f.codes[1] == 23 && f.codes[2] == 31 &&
		(f.codes[3] == 32 || f.codes[3] == 33 || f.codes[3] == 34))
}

func translateError(err error) awserr.Error {
	if err == nil {
		return nil
	}
	switch err.(type) {
	case awserr.Error:
		e := err.(awserr.Error)
		return e
	case net.Error:
		e := err.(net.Error)
		code := dynamodb.ErrCodeInternalServerError
		if e.Timeout() {
			code = request.ErrCodeResponseTimeout
		}
		return awserr.New(code, "network error", e)
	default:
		return awserr.New("UnknownError", "unknown error", err)
	}
}

func decodeError(reader *cbor.Reader) (awserr.Error, error) {
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
			return nil, awserr.New(request.ErrCodeSerialization, fmt.Sprintf("expected 3 or 4 elements for error info, got %d", length), nil)
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
				return nil, awserr.New(request.ErrCodeSerialization, "error found when parsing CancellationReasons", nil)
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

	if cancellationReasonCodes != nil && len(cancellationReasonCodes) > 0 {
		return newDaxTransactionCanceledFailure(codes, errorCode, msg, requestId, statusCode,
			cancellationReasonCodes, cancellationReasonMsgs, cancellationReasonItems), nil
	}
	return newDaxRequestFailure(codes, errorCode, msg, requestId, statusCode), nil
}

// convertDAXError converts DAX error to specific error type based on error code sequense returned from server.
func convertDaxError(e daxError) error {
	codes := e.CodeSequence()
	if len(codes) < 2 {
		return e
	}
	md := protocol.ResponseMetadata{
		StatusCode: e.StatusCode(),
		RequestID:  e.RequestID(),
	}
	switch codes[1] {
	case 23:
		if len(codes) > 2 {
			switch codes[2] {
			case 24:
				return &dynamodb.ResourceNotFoundException{
					RespMetadata: md,
					Message_:     aws.String(e.Message()),
				}
			case 35:
				return &dynamodb.ResourceInUseException{
					RespMetadata: md,
					Message_:     aws.String(e.Message()),
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
						return &dynamodb.ProvisionedThroughputExceededException{
							RespMetadata: md,
							Message_:     aws.String(e.Message()),
						}
					case 41:
						return &dynamodb.ResourceNotFoundException{
							RespMetadata: md,
							Message_:     aws.String(e.Message()),
						}
					case 43:
						return &dynamodb.ConditionalCheckFailedException{
							RespMetadata: md,
							Message_:     aws.String(e.Message()),
						}
					case 45:
						return &dynamodb.ResourceInUseException{
							RespMetadata: md,
							Message_:     aws.String(e.Message())}
					case 46:
						// there's no dynamodb.ValidationException type
						return awserr.NewRequestFailure(awserr.New(ErrCodeValidationException, e.Message(), nil), e.StatusCode(), e.RequestID())
					case 47:
						return &dynamodb.InternalServerError{
							RespMetadata: md,
							Message_:     aws.String(e.Message()),
						}
					case 48:
						return &dynamodb.ItemCollectionSizeLimitExceededException{
							RespMetadata: md,
							Message_:     aws.String(e.Message()),
						}
					case 49:
						return &dynamodb.LimitExceededException{
							RespMetadata: md,
							Message_:     aws.String(e.Message()),
						}
					case 50:
						// there's no dynamodb.ThrottlingException type
						return awserr.NewRequestFailure(awserr.New(ErrCodeThrottlingException, e.Message(), nil), e.StatusCode(), e.RequestID())
					case 57:
						return &dynamodb.TransactionConflictException{
							RespMetadata: md,
							Message_:     aws.String(e.Message()),
						}
					case 58:
						tcFailure, ok := e.(*daxTransactionCanceledFailure)
						if ok {
							return &dynamodb.TransactionCanceledException{
								RespMetadata:        md,
								Message_:            aws.String(e.Message()),
								CancellationReasons: tcFailure.cancellationReasons,
							}
						}
					case 59:
						return &dynamodb.TransactionInProgressException{
							RespMetadata: md,
							Message_:     aws.String(e.Message()),
						}
					case 60:
						return &dynamodb.IdempotentParameterMismatchException{
							RespMetadata: md,
							Message_:     aws.String(e.Message()),
						}
					}
				}
			case 44:
				return awserr.NewRequestFailure(awserr.New(ErrCodeNotImplemented, e.Message(), nil), e.StatusCode(), e.RequestID())
			}
		}
	}
	return awserr.NewRequestFailure(awserr.New(ErrCodeUnknown, e.Message(), nil), e.StatusCode(), e.RequestID())
}

func decodeTransactionCancellationReasons(ctx aws.Context, failure *daxTransactionCanceledFailure,
	keys []map[string]*dynamodb.AttributeValue, attrListIdToNames *lru.Lru) ([]*dynamodb.CancellationReason, error) {
	inputL := len(keys)
	outputL := len(failure.cancellationReasonCodes)
	if inputL != outputL {
		return nil, awserr.New(request.ErrCodeSerialization, "Cancellation reasons must be the same length as transact items in the request", nil)
	}
	reasons := make([]*dynamodb.CancellationReason, outputL)
	r := cbor.NewReader(bytes.NewReader(failure.cancellationReasonItems))
	for i := 0; i < outputL; i++ {
		reason := new(dynamodb.CancellationReason)
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
