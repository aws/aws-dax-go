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
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func TestAttrVal(t *testing.T) {
	cases := []struct {
		val types.AttributeValue
		enc []byte
	}{
		{val: &types.AttributeValueMemberS{Value: "abc"}},
		{val: &types.AttributeValueMemberS{Value: "abcdefghijklmnopqrstuvwxyz0123456789"}},
		{val: &types.AttributeValueMemberN{Value: "123"}},
		{val: &types.AttributeValueMemberN{Value: "-123"}},
		{val: &types.AttributeValueMemberN{Value: "123456789012345678901234567890"}},
		{val: &types.AttributeValueMemberN{Value: "-123456789012345678901234567890"}},
		{val: &types.AttributeValueMemberN{Value: "314E-2"}},
		{val: &types.AttributeValueMemberN{Value: "-314E-2"}},
		//{val: types.AttributeValue{N: stringptr("3.14")}},	// Decimal.String() return 314E-2
		{val: &types.AttributeValueMemberB{Value: fromHex("0x010203")}},
		{val: &types.AttributeValueMemberSS{Value: []string{"abc", "def", "xyz"}}},
		{val: &types.AttributeValueMemberNS{Value: []string{"123", "456", "789"}}},
		{val: &types.AttributeValueMemberBS{Value: [][]byte{fromHex("0x010203"), fromHex("0x040506")}}},
		{val: &types.AttributeValueMemberL{Value: []types.AttributeValue{&types.AttributeValueMemberS{Value: "abc"}, &types.AttributeValueMemberN{Value: "123"}}}},
		{val: &types.AttributeValueMemberM{Value: map[string]types.AttributeValue{"s": &types.AttributeValueMemberS{Value: "abc"}, "n": &types.AttributeValueMemberN{Value: "123"}}}},
		{val: &types.AttributeValueMemberBOOL{Value: true}},
		{val: &types.AttributeValueMemberBOOL{Value: false}},
		{val: &types.AttributeValueMemberNULL{Value: true}},
	}

	for _, c := range cases {
		lval := c.val
		var buf bytes.Buffer
		w := NewWriter(&buf)
		if err := EncodeAttributeValue(lval, w); err != nil {
			t.Errorf("unexpected error %v for %v", err, lval)
			continue
		}
		if err := w.Flush(); err != nil {
			t.Errorf("unexpected error %v for %v", err, lval)
			continue
		}

		bufBytes := buf.Bytes()
		if c.enc != nil && !reflect.DeepEqual(c.enc, bufBytes) {
			t.Errorf("incorrect encoding for %v", c.val)
		}

		r := NewReader(&buf)
		rval, err := DecodeAttributeValue(r)
		if err != nil {
			t.Errorf("unexpected error %v for %v", err, lval)
			continue
		}

		if !reflect.DeepEqual(lval, rval) {
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
		if eAttr := types.AttributeValue(&types.AttributeValueMemberN{Value: e.value.String()}); !reflect.DeepEqual(eAttr, a) {
			t.Errorf("test %s expected: %v, actual: %v", e.name, eAttr, a)
		}
	}
}
