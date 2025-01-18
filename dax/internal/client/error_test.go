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
	"io"
	"net"
	"reflect"
	"testing"

	"github.com/aws/aws-dax-go/dax/internal/cbor"
	"github.com/aws/aws-dax-go/dax/internal/lru"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/smithy-go"
)

func TestDecodeError(t *testing.T) {
	var b bytes.Buffer
	errCodes := []int{4, 37, 38, 39, 40}
	requestID := "request-1"
	statusCode := 400
	exception := types.ProvisionedThroughputExceededException{
		Message: aws.String("ProvisionedThroughputExceededException Message"),
	}

	w := cbor.NewWriter(&b)
	_ = w.WriteArrayHeader(len(errCodes))
	for _, c := range errCodes {
		_ = w.WriteInt(c)
	}
	_ = w.WriteString(exception.ErrorMessage())

	_ = w.WriteArrayHeader(3)
	_ = w.WriteString(requestID)
	_ = w.WriteString(exception.ErrorCode())
	_ = w.WriteInt(statusCode)
	_ = w.Flush()

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
		GenericAPIError: &smithy.GenericAPIError{
			Code:    exception.ErrorCode(),
			Message: exception.ErrorMessage(),
			Fault:   smithy.FaultServer,
		},
		codes:      errCodes,
		requestID:  requestID,
		statusCode: statusCode,
	}

	if !reflect.DeepEqual(expected, d) {
		t.Errorf("expected %v, got %v", expected, d)
	}
}

func TestDecodeTransactionCanceledException(t *testing.T) {
	errCodes := []int{4, 37, 38, 39, 58}
	requestID := "request-1"
	statusCode := 400
	exception := types.TransactionCanceledException{
		Message: aws.String("TransactionCanceledException Message"),
		CancellationReasons: []types.CancellationReason{
			{Code: aws.String("reasonCode1"), Item: map[string]types.AttributeValue{}, Message: aws.String("reasonMsg1")},
			{Code: aws.String("reasonCode2"), Item: map[string]types.AttributeValue{}, Message: aws.String("reasonMsg2")},
		},
	}
	items := []byte{}
	var expItems []byte

	var b bytes.Buffer
	w := cbor.NewWriter(&b)
	_ = w.WriteArrayHeader(len(errCodes))
	for _, c := range errCodes {
		_ = w.WriteInt(c)
	}
	_ = w.WriteString(exception.ErrorMessage())

	_ = w.WriteArrayHeader(4)
	_ = w.WriteString(requestID)
	_ = w.WriteString(exception.ErrorCode())
	_ = w.WriteInt(statusCode)
	_ = w.WriteArrayHeader(3 * len(exception.CancellationReasons))
	for i := 0; i < len(exception.CancellationReasons); i++ {
		_ = w.WriteString(*exception.CancellationReasons[i].Code)
		_ = w.WriteString(*exception.CancellationReasons[i].Message)
		_ = w.WriteBytes(items)

		buf := bytes.Buffer{}
		nw := cbor.NewWriter(&buf)
		_ = nw.WriteBytes(items)
		_ = nw.Flush()

		r := cbor.NewReader(&buf)
		obuf := bytes.Buffer{}
		_ = r.ReadRawBytes(&obuf)

		expItems = append(expItems, obuf.Bytes()...)
	}
	_ = w.Flush()

	r := cbor.NewReader(&b)
	e, err := decodeError(r)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	d, ok := e.(*daxTransactionCanceledFailure)
	if !ok {
		t.Errorf("expected daxTransactionCanceledFailure type")
	}

	reasonCodes := make([]*string, 0, len(exception.CancellationReasons))
	reasonMsgs := make([]*string, 0, len(exception.CancellationReasons))
	for _, r := range exception.CancellationReasons {
		reasonCodes = append(reasonCodes, r.Code)
		reasonMsgs = append(reasonMsgs, r.Message)
	}

	expected := &daxTransactionCanceledFailure{
		daxRequestFailure: &daxRequestFailure{
			GenericAPIError: &smithy.GenericAPIError{
				Code:    exception.ErrorCode(),
				Message: exception.ErrorMessage(),
				Fault:   smithy.FaultServer,
			},
			codes:      errCodes,
			requestID:  requestID,
			statusCode: statusCode,
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
//  1. transact item didn't fail conditional check
//  2. transact item failed conditional check and was configured to return ALL_OLD item
//  3. transact item failed conditional check and was configured to return NONE item
func TestDecodeTransactionCancellationReasons(t *testing.T) {
	expCodes := []int{1, 2, 3, 4}

	expErrCode := (&types.TransactionCanceledException{}).ErrorCode()
	expMsg := "Transaction was cancelled."
	expReqID := "134213414395861"
	expStatusCode := 400
	expCanceledCodes := []*string{
		aws.String("NONE"),
		aws.String((&types.ConditionalCheckFailedException{}).ErrorCode()),
		aws.String((&types.TransactionInProgressException{}).ErrorCode()),
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
		LoadFunc: func(ctx context.Context, key lru.Key) (interface{}, error) {
			return int64(12345), nil
		},
		KeyMarshaller: func(key lru.Key) lru.Key {
			var buf bytes.Buffer
			w := cbor.NewWriter(&buf)
			defer w.Close()
			for _, v := range key.([]string) {
				_ = w.WriteString(v)
			}
			_ = w.Flush()
			return string(buf.Bytes())
		},
	}
	idToAttrs := &lru.Lru{
		LoadFunc: func(ctx context.Context, key lru.Key) (interface{}, error) {
			return attrs, nil
		},
	}

	// nbuf mocks CBOR output from server
	buf := bytes.Buffer{}
	w := cbor.NewWriter(&buf)
	cbor.EncodeItemNonKeyAttributes(nil, canceledItems[1], keyDef, attrsToID, w)
	_ = w.Flush()

	nbuf := bytes.Buffer{}
	nw := cbor.NewWriter(&nbuf)
	_ = nw.WriteNull()
	_ = nw.WriteBytes(buf.Bytes())
	_ = nw.WriteNull()
	_ = nw.Flush()

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
	errCodes := []int{4, 37, 38, 39, 43}
	exception := types.ConditionalCheckFailedException{
		Message: aws.String("ConditionalCheckFailedException Message"),
	}
	//awserr := awserr.NewRequestFailure(awserr.New(dynamodb.ErrCodeConditionalCheckFailedException, "ConditionalCheckFailedException Message", nil), 400, "")

	w := cbor.NewWriter(&b)
	_ = w.WriteArrayHeader(len(errCodes))
	for _, c := range errCodes {
		_ = w.WriteInt(c)
	}
	_ = w.WriteString(exception.ErrorMessage())

	_ = w.WriteArrayHeader(3)
	_ = w.WriteNull()
	_ = w.WriteString(exception.ErrorCode())
	_ = w.WriteNull() // status code will be inferred from error code
	_ = w.Flush()

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
		GenericAPIError: &smithy.GenericAPIError{
			Code:    exception.ErrorCode(),
			Message: exception.ErrorMessage(),
			Fault:   smithy.FaultServer,
		},
		codes:      errCodes,
		requestID:  "",
		statusCode: 400,
	}

	if !reflect.DeepEqual(expected, d) {
		t.Errorf("expected %v, got %v", expected, d)
	}

}

func TestIsIoError(t *testing.T) {
	if isIOError(context.DeadlineExceeded) != true {
		t.Errorf("isIOError Failed to identify timeout error")
	}

	if isIOError(context.Canceled) != true {
		t.Errorf("isIOError Failed to identify timeout error")
	}

	if isIOError(io.EOF) != true {
		t.Errorf("isIOError Failed to identify EOF error")
	}

	if isIOError(&smithy.OperationError{Err: errors.New("network error: caused by: dial tcp 172.31.0.205:8111: connect: connection refused")}) != true {
		t.Errorf("isIOError Failed to identify network error")
	}

	if isIOError(&smithy.OperationError{Err: nil}) == true {
		t.Errorf("isIOError Failed to identify non-timeout error")
	}

	if isIOError(net.ErrClosed) != true {
		t.Errorf("isIOError Failed to identify timeout error")
	}

	if isIOError(nil) == true {
		t.Errorf("isIOError Failed to identify non-timeout error")
	}
}
