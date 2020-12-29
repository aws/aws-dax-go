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

package cbor

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"reflect"
	"testing"
)

func TestAttrVal(t *testing.T) {
	cases := []struct {
		val dynamodb.AttributeValue
		enc []byte
	}{
		{val: dynamodb.AttributeValue{S: aws.String("abc")}},
		{val: dynamodb.AttributeValue{S: aws.String("abcdefghijklmnopqrstuvwxyz0123456789")}},
		{val: dynamodb.AttributeValue{N: aws.String("123")}},
		{val: dynamodb.AttributeValue{N: aws.String("-123")}},
		{val: dynamodb.AttributeValue{N: aws.String("123456789012345678901234567890")}},
		{val: dynamodb.AttributeValue{N: aws.String("-123456789012345678901234567890")}},
		{val: dynamodb.AttributeValue{N: aws.String("314E-2")}},
		{val: dynamodb.AttributeValue{N: aws.String("-314E-2")}},
		//{val: dynamodb.AttributeValue{N: stringptr("3.14")}},	// Decimal.String() return 314E-2
		{val: dynamodb.AttributeValue{B: fromHex("0x010203")}},
		{val: dynamodb.AttributeValue{SS: []*string{aws.String("abc"), aws.String("def"), aws.String("xyz")}}},
		{val: dynamodb.AttributeValue{NS: []*string{aws.String("123"), aws.String("456"), aws.String("789")}}},
		{val: dynamodb.AttributeValue{BS: [][]byte{fromHex("0x010203"), fromHex("0x040506")}}},
		{val: dynamodb.AttributeValue{L: []*dynamodb.AttributeValue{&dynamodb.AttributeValue{S: aws.String("abc")}, &dynamodb.AttributeValue{N: aws.String("123")}}}},
		{val: dynamodb.AttributeValue{M: map[string]*dynamodb.AttributeValue{"s": &dynamodb.AttributeValue{S: aws.String("abc")}, "n": &dynamodb.AttributeValue{N: aws.String("123")}}}},
		{val: dynamodb.AttributeValue{BOOL: aws.Bool(true)}},
		{val: dynamodb.AttributeValue{BOOL: aws.Bool(false)}},
		{val: dynamodb.AttributeValue{NULL: aws.Bool(true)}},
	}

	for _, c := range cases {
		lval := c.val
		var buf bytes.Buffer
		w := NewWriter(&buf)
		if err := EncodeAttributeValue(&lval, w); err != nil {
			t.Errorf("unexpected error %v for %v", err, lval)
			continue
		}
		if err := w.Flush(); err != nil {
			t.Errorf("unexpected error %v for %v", err, lval)
			continue
		}

		bytes := buf.Bytes()
		if c.enc != nil && !reflect.DeepEqual(c.enc, bytes) {
			t.Errorf("incorrect encoding for %v", c.val)
		}

		r := NewReader(&buf)
		rval, err := DecodeAttributeValue(r)
		if err != nil {
			t.Errorf("unexpected error %v for %v", err, lval)
			continue
		}

		if !reflect.DeepEqual(lval, *rval) {
			t.Errorf("expected: %v, actual: %v", lval, rval)
		}
	}
}

func TestDecodeIntBoundariesFromCbor(t *testing.T) {
	for _, e := range []IntBoundary{
		MinCborNegativeIntMinusOne,
		MinCborNegativeInt,
		MinCborNegativeIntPlusOne,
		MinInt64MinusOne,
		MinInt64,
		MinusOne,
		Zero,
		MaxInt64,
		MaxInt64PlusOne,
		MaxCborPositiveInt,
		MaxUint64,
		MaxUint64PlusOne,
		MaxCborPositiveIntPlusOne,
	} {
		var buf bytes.Buffer
		buf.Write(e.cbor)
		a, err := DecodeAttributeValue(NewReader(&buf))
		if err != nil {
			t.Errorf("unexpected error %v for %s", err, e.name)
		}
		if eAttr := (dynamodb.AttributeValue{N: aws.String(e.value.String())}); !reflect.DeepEqual(eAttr, *a) {
			t.Errorf("test %s expected: %v, actual: %v", e.name, eAttr, a)
		}
	}
}
