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
	"net"
	"reflect"
	"testing"

	"github.com/aws/aws-dax-go/dax/internal/cbor"
	"github.com/aws/aws-dax-go/dax/internal/lru"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
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

func TestDecodeTransactionCanceledException(t *testing.T) {
	errcode := []int{4, 37, 38, 39, 58}
	awserr := awserr.NewRequestFailure(awserr.New(dynamodb.ErrCodeTransactionCanceledException, "TransactionCanceledException Message", nil), 400, "request-1")
	reasonLen := 2
	reasonCodes := []*string{aws.String("reasonCode1"), aws.String("reasonCode2")}
	reasonMsgs := []*string{aws.String("reasonMsg1"), aws.String("reasonMsg2")}
	items := []byte{}
	var expItems []byte

	var b bytes.Buffer
	w := cbor.NewWriter(&b)
	w.WriteArrayHeader(len(errcode))
	for _, c := range errcode {
		w.WriteInt(c)
	}
	w.WriteString(awserr.Message())

	w.WriteArrayHeader(4)
	w.WriteString(awserr.RequestID())
	w.WriteString(awserr.Code())
	w.WriteInt(awserr.StatusCode())
	w.WriteArrayHeader(3 * reasonLen)
	for i := 0; i < reasonLen; i++ {
		w.WriteString(*reasonCodes[i])
		w.WriteString(*reasonMsgs[i])
		w.WriteBytes(items)

		buf := bytes.Buffer{}
		nw := cbor.NewWriter(&buf)
		nw.WriteBytes(items)
		nw.Flush()

		r := cbor.NewReader(&buf)
		obuf := bytes.Buffer{}
		r.ReadRawBytes(&obuf)

		expItems = append(expItems, obuf.Bytes()...)
	}
	w.Flush()

	r := cbor.NewReader(&b)
	e, err := decodeError(r)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	d, ok := e.(*daxTransactionCanceledFailure)
	if !ok {
		t.Errorf("expected daxTransactionCanceledFailure type")
	}

	expected := &daxTransactionCanceledFailure{
		daxRequestFailure: &daxRequestFailure{
			RequestFailure: awserr,
			codes:          errcode,
		},
		cancellationReasonCodes: reasonCodes,
		cancellationReasonMsgs:  reasonMsgs,
		cancellationReasonItems: expItems,
	}

	if !reflect.DeepEqual(expected, d) {
		t.Errorf("expected %v, got %v", expected, d)
	}
}

// TestDecodeTransactionCancellationReasons tests decoding transaction cancellations reasons in daxTransactionCanceledFailure.
//
// Specifically, the decoding of items in cancellation reasons are being testing here. It covers three situations:
//    1. transact item didn't fail conditional check
//    2. transact item failed conditional check and was configured to return ALL_OLD item
//    3. transact item failed conditional check and was configured to return NONE item
func TestDecodeTransactionCancellationReasons(t *testing.T) {
	expCodes := []int{1, 2, 3, 4}
	expErrCode := dynamodb.ErrCodeTransactionCanceledException
	expMsg := "Transaction was cancelled."
	expReqID := "134213414395861"
	expStatusCode := 400
	expCanceledCodes := []*string{
		aws.String("NONE"),
		aws.String(dynamodb.ErrCodeConditionalCheckFailedException),
		aws.String(dynamodb.ErrCodeTransactionInProgressException),
	}
	expCanceledReasons := []*string{
		nil,
		aws.String("first reason"),
		aws.String("second reason"),
	}
	keyDef := []types.AttributeDefinition{
		{AttributeName: aws.String("hk")},
	}
	keys := []map[string]types.AttributeValue{
		{"hk": &types.AttributeValueMemberN{Value: "0"}},
		{"hk": &types.AttributeValueMemberN{Value: "0"}},
		{"hk": &types.AttributeValueMemberN{Value: "0"}},
	}
	canceledItems := []map[string]types.AttributeValue{
		nil,
		{"attr": &types.AttributeValueMemberN{Value: "0"}},
		nil,
	}
	attrs := []string{"attr"}
	attrsToID := &lru.Lru{
		LoadFunc: func(ctx aws.Context, key lru.Key) (interface{}, error) {
			return int64(12345), nil
		},
		KeyMarshaller: func(key lru.Key) lru.Key {
			var buf bytes.Buffer
			w := cbor.NewWriter(&buf)
			defer w.Close()
			for _, v := range key.([]string) {
				w.WriteString(v)
			}
			w.Flush()
			return string(buf.Bytes())
		},
	}
	idToAttrs := &lru.Lru{
		LoadFunc: func(ctx aws.Context, key lru.Key) (interface{}, error) {
			return attrs, nil
		},
	}

	// nbuf mocks CBOR output from server
	buf := bytes.Buffer{}
	w := cbor.NewWriter(&buf)
	cbor.EncodeItemNonKeyAttributes(nil, canceledItems[1], keyDef, attrsToID, w)
	w.Flush()

	nbuf := bytes.Buffer{}
	nw := cbor.NewWriter(&nbuf)
	nw.WriteNull()
	nw.WriteBytes(buf.Bytes())
	nw.WriteNull()
	nw.Flush()

	for k, v := range keys[1] {
		canceledItems[1][k] = v
	}

	expCancellationReason := []types.CancellationReason{
		{
			Code: expCanceledCodes[0],
		},
		{
			Code:    expCanceledCodes[1],
			Message: expCanceledReasons[1],
			Item:    canceledItems[1],
		},
		{
			Code:    expCanceledCodes[2],
			Message: expCanceledReasons[2],
		},
	}

	expTcErr := newDaxTransactionCanceledFailure(expCodes, expErrCode, expMsg, expReqID, expStatusCode, expCanceledCodes, expCanceledReasons, nbuf.Bytes())
	expTcErr.cancellationReasons = expCancellationReason

	// tcErr mocks partial decoded output from error.decodeError(*cbor.Reader)
	tcErr := newDaxTransactionCanceledFailure(expCodes, expErrCode, expMsg, expReqID, expStatusCode, expCanceledCodes, expCanceledReasons, nbuf.Bytes())

	// Method under test
	cancellationReason, err := decodeTransactionCancellationReasons(nil, tcErr, keys, idToAttrs)
	tcErr.cancellationReasons = cancellationReason

	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	tcErr.cancellationReasons = cancellationReason

	if !reflect.DeepEqual(expTcErr, tcErr) {
		t.Errorf("expected %v, got %v", expTcErr, tcErr)
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
