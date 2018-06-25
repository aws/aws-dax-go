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
	"errors"
	"github.com/aws/aws-dax-go/dax/internal/cbor"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"net"
	"reflect"
	"testing"
)

func TestDecodeError(t *testing.T) {
	var b bytes.Buffer
	errcode := []int{4, 37, 38, 39, 40}
	awserr := awserr.NewRequestFailure(awserr.New(dynamodb.ErrCodeProvisionedThroughputExceededException, "ProvisionedThroughputExceededException Message", nil), 400, "request-1")

	w := cbor.NewWriter(&b)
	w.WriteArrayHeader(len(errcode))
	for _, c := range errcode {
		w.WriteInt(c)
	}
	w.WriteString(awserr.Message())

	w.WriteArrayHeader(3)
	w.WriteString(awserr.RequestID())
	w.WriteString(awserr.Code())
	w.WriteInt(awserr.StatusCode())
	w.Flush()

	r := cbor.NewReader(&b)
	e, err := decodeError(r)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	d, ok := e.(*daxRequestFailure)
	if !ok {
		t.Errorf("expected daxRequestFailure type")
	}

	expected := &daxRequestFailure{
		RequestFailure: awserr,
		codes:          errcode,
	}

	if !reflect.DeepEqual(expected, d) {
		t.Errorf("expected %v, got %v", expected, d)
	}
}

func TestDecodeErrorInfer(t *testing.T) {
	var b bytes.Buffer
	errcode := []int{4, 37, 38, 39, 43}
	awserr := awserr.NewRequestFailure(awserr.New(dynamodb.ErrCodeConditionalCheckFailedException, "ConditionalCheckFailedException Message", nil), 400, "")

	w := cbor.NewWriter(&b)
	w.WriteArrayHeader(len(errcode))
	for _, c := range errcode {
		w.WriteInt(c)
	}
	w.WriteString(awserr.Message())
	w.WriteNull()
	w.Flush()

	r := cbor.NewReader(&b)
	e, err := decodeError(r)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	d, ok := e.(*daxRequestFailure)
	if !ok {
		t.Errorf("expected daxRequestFailure type")
	}

	expected := &daxRequestFailure{
		RequestFailure: awserr,
		codes:          errcode,
	}

	if !reflect.DeepEqual(expected, d) {
		t.Errorf("expected %v, got %v", expected, d)
	}
}

func TestDecodeNilErrorDetail(t *testing.T) {
	var b bytes.Buffer
	errcode := []int{4, 37, 38, 39, 43}
	awserr := awserr.NewRequestFailure(awserr.New(dynamodb.ErrCodeConditionalCheckFailedException, "ConditionalCheckFailedException Message", nil), 400, "")

	w := cbor.NewWriter(&b)
	w.WriteArrayHeader(len(errcode))
	for _, c := range errcode {
		w.WriteInt(c)
	}
	w.WriteString(awserr.Message())

	w.WriteArrayHeader(3)
	w.WriteNull()
	w.WriteString(awserr.Code())
	w.WriteNull() // status code will be inferred from error code
	w.Flush()

	r := cbor.NewReader(&b)
	e, err := decodeError(r)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	d, ok := e.(*daxRequestFailure)
	if !ok {
		t.Errorf("expected daxRequestFailure type")
	}

	expected := &daxRequestFailure{
		RequestFailure: awserr,
		codes:          errcode,
	}

	if !reflect.DeepEqual(expected, d) {
		t.Errorf("expected %v, got %v", expected, d)
	}

}

func TestTranslateError(t *testing.T) {
	cases := []struct {
		input  error
		output error
	}{
		{
			input:  newDaxRequestFailure([]int{1, 2, 3}, "ec", "msg", "rid", 500),
			output: newDaxRequestFailure([]int{1, 2, 3}, "ec", "msg", "rid", 500),
		},
		{
			input:  awserr.NewRequestFailure(awserr.New("ec", "msg", nil), 500, "rid"),
			output: awserr.NewRequestFailure(awserr.New("ec", "msg", nil), 500, "rid"),
		},
		{
			input:  awserr.New("ec", "msg", nil),
			output: awserr.New("ec", "msg", nil),
		},
		{
			input:  new(net.UnknownNetworkError),
			output: awserr.New(dynamodb.ErrCodeInternalServerError, "network error", new(net.UnknownNetworkError)),
		},
		{
			input:  errors.New("ex"),
			output: awserr.New("UnknownError", "unknown error", errors.New("ex")),
		},
	}

	for _, c := range cases {
		actual := translateError(c.input)
		if !reflect.DeepEqual(c.output, actual) {
			t.Errorf("expected %v, got %v", c.output, actual)
		}
	}
}
