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
	"fmt"
	"github.com/aws/aws-dax-go/dax/internal/cbor"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"net"
)

const (
	ErrCodeNotImplemented      = "NotImplemented"
	ErrCodeValidationException = "ValidationException"
	ErrCodeServiceUnavailable  = "ServiceUnavailable"
	ErrCodeUnknown             = "Unknown"
)

type daxRequestFailure struct {
	awserr.RequestFailure
	codes []int
}

func newDaxRequestFailure(codes []int, errorCode, message, requestId string, statusCode int) *daxRequestFailure {
	return &daxRequestFailure{
		RequestFailure: awserr.NewRequestFailure(awserr.New(errorCode, message, nil), statusCode, requestId),
		codes:          codes,
	}
}

func (f *daxRequestFailure) retryable() bool {
	return len(f.codes) > 0 && (f.codes[0] == 1 || f.codes[0] == 2)
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
	len, err := reader.ReadArrayLength()
	if err != nil {
		return nil, err
	}
	if len == 0 {
		return nil, nil
	}

	codes := make([]int, len)
	for i := 0; i < len; i++ {
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
	hdr, err := reader.PeekHeader()
	if err != nil {
		return nil, err
	}
	if hdr == cbor.Nil {
		if err := reader.ReadNil(); err != nil {
			return nil, err
		}
	} else {
		len, err = reader.ReadArrayLength()
		if err != nil {
			return nil, err
		}
		if len != 3 {
			return nil, awserr.New(request.ErrCodeSerialization, fmt.Sprintf("expected 3 elements for error info, got %d", len), nil)
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
	}

	if errorCode == "" {
		errorCode = inferErrorCode(codes)
	}
	if statusCode == 0 {
		statusCode = inferStatusCode(codes)
	}

	return newDaxRequestFailure(codes, errorCode, msg, requestId, statusCode), nil
}

func inferErrorCode(codes []int) string {
	if len(codes) < 2 {
		return ""
	}
	switch codes[1] {
	case 23:
		if len(codes) > 2 {
			switch codes[2] {
			case 24:
				return dynamodb.ErrCodeResourceNotFoundException
			case 35:
				return dynamodb.ErrCodeResourceInUseException
			}
		}
	case 37:
		if len(codes) > 3 {
			switch codes[3] {
			case 39:
				if len(codes) > 4 {
					switch codes[4] {
					case 40:
						return dynamodb.ErrCodeProvisionedThroughputExceededException
					case 41:
						return dynamodb.ErrCodeResourceNotFoundException
					case 43:
						return dynamodb.ErrCodeConditionalCheckFailedException
					case 45:
						return dynamodb.ErrCodeResourceInUseException
					case 46:
						return ErrCodeValidationException
					case 47:
						return dynamodb.ErrCodeInternalServerError
					case 48:
						return dynamodb.ErrCodeItemCollectionSizeLimitExceededException
					case 49:
						return dynamodb.ErrCodeLimitExceededException
					}
				}
			case 44:
				return ErrCodeNotImplemented
			}
		}
	}
	return ErrCodeUnknown
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
