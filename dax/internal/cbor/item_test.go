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
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/aws/aws-dax-go/dax/internal/lru"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func TestItemKey(t *testing.T) {
	cases := []struct {
		keydef []types.AttributeDefinition
		item   map[string]types.AttributeValue
		enc    []byte
	}{
		{
			keydef: []types.AttributeDefinition{{AttributeName: aws.String("hks"), AttributeType: types.ScalarAttributeTypeS}},
			item: map[string]types.AttributeValue{
				"hks": &types.AttributeValueMemberS{Value: "hkv"},
			},
			enc: fromHex("0x43686b76"),
		},
		{
			keydef: []types.AttributeDefinition{{AttributeName: aws.String("hkn"), AttributeType: types.ScalarAttributeTypeN}},
			item: map[string]types.AttributeValue{
				"hkn": &types.AttributeValueMemberN{Value: "5"},
			},
			enc: fromHex("0x4105"),
		},
		{
			keydef: []types.AttributeDefinition{{AttributeName: aws.String("hkb"), AttributeType: types.ScalarAttributeTypeB}},
			item: map[string]types.AttributeValue{
				"hkb": &types.AttributeValueMemberB{Value: fromHex("0x010203")},
			},
			enc: fromHex("0x43010203"),
		},
		{
			keydef: []types.AttributeDefinition{
				{AttributeName: aws.String("hks"), AttributeType: types.ScalarAttributeTypeS},
				{AttributeName: aws.String("rks"), AttributeType: types.ScalarAttributeTypeS},
			},
			item: map[string]types.AttributeValue{
				"hks": &types.AttributeValueMemberS{Value: "hkv"},
				"rks": &types.AttributeValueMemberS{Value: "rkv"},
			},
			enc: fromHex("0x4763686b76726b76"),
		},
		{
			keydef: []types.AttributeDefinition{
				{AttributeName: aws.String("hks"), AttributeType: types.ScalarAttributeTypeS},
				{AttributeName: aws.String("rkn"), AttributeType: types.ScalarAttributeTypeN},
			},
			item: map[string]types.AttributeValue{
				"hks": &types.AttributeValueMemberS{Value: "hkv"},
				"rkn": &types.AttributeValueMemberN{Value: "5"},
			},
			//enc:fromHex("0x4563686b76724105"), TODO lex decimal
		},
		{
			keydef: []types.AttributeDefinition{
				{AttributeName: aws.String("hks"), AttributeType: types.ScalarAttributeTypeS},
				{AttributeName: aws.String("rkb"), AttributeType: types.ScalarAttributeTypeB},
			},
			item: map[string]types.AttributeValue{
				"hks": &types.AttributeValueMemberS{Value: "hkv"},
				"rkb": &types.AttributeValueMemberB{Value: fromHex("0x010203")},
			},
			enc: fromHex("0x4763686b76010203"),
		},
		{
			keydef: []types.AttributeDefinition{
				{AttributeName: aws.String("hkn"), AttributeType: types.ScalarAttributeTypeN},
				{AttributeName: aws.String("rks"), AttributeType: types.ScalarAttributeTypeS},
			},
			item: map[string]types.AttributeValue{
				"hkn": &types.AttributeValueMemberN{Value: "5"},
				"rks": &types.AttributeValueMemberS{Value: "rkv"},
			},
			enc: fromHex("0x4405726b76"),
		},
		{
			keydef: []types.AttributeDefinition{
				{AttributeName: aws.String("hkn"), AttributeType: types.ScalarAttributeTypeN},
				{AttributeName: aws.String("rkn"), AttributeType: types.ScalarAttributeTypeN},
			},
			item: map[string]types.AttributeValue{
				"hkn": &types.AttributeValueMemberN{Value: "5"},
				"rkn": &types.AttributeValueMemberN{Value: "1"},
			},
			//enc:fromHex("0x4105726b76"), TODO lex decimal
		},
		{
			keydef: []types.AttributeDefinition{
				{AttributeName: aws.String("hkn"), AttributeType: types.ScalarAttributeTypeN},
				{AttributeName: aws.String("rkb"), AttributeType: types.ScalarAttributeTypeB},
			},
			item: map[string]types.AttributeValue{
				"hkn": &types.AttributeValueMemberN{Value: "5"},
				"rkb": &types.AttributeValueMemberB{Value: fromHex("0x010203")},
			},
			enc: fromHex("0x4405010203"),
		},
		{
			keydef: []types.AttributeDefinition{
				{AttributeName: aws.String("hkb"), AttributeType: types.ScalarAttributeTypeB},
				{AttributeName: aws.String("rks"), AttributeType: types.ScalarAttributeTypeS},
			},
			item: map[string]types.AttributeValue{
				"hkb": &types.AttributeValueMemberB{Value: fromHex("0x040506")},
				"rks": &types.AttributeValueMemberS{Value: "rkv"},
			},
			enc: fromHex("0x4743040506726b76"),
		},
		{
			keydef: []types.AttributeDefinition{
				{AttributeName: aws.String("hkb"), AttributeType: types.ScalarAttributeTypeB},
				{AttributeName: aws.String("rkn"), AttributeType: types.ScalarAttributeTypeN},
			},
			item: map[string]types.AttributeValue{
				"hkb": &types.AttributeValueMemberB{Value: fromHex("0x040506")},
				"rkn": &types.AttributeValueMemberN{Value: "123"},
			},
			//enc:fromHex("0x4743040506726b76"), TODO lex decimal
		},
		{
			keydef: []types.AttributeDefinition{
				{AttributeName: aws.String("hkb"), AttributeType: types.ScalarAttributeTypeB},
				{AttributeName: aws.String("rkb"), AttributeType: types.ScalarAttributeTypeB},
			},
			item: map[string]types.AttributeValue{
				"hkb": &types.AttributeValueMemberB{Value: fromHex("0x040506")},
				"rkb": &types.AttributeValueMemberB{Value: fromHex("0x010203")},
			},
			enc: fromHex("0x4743040506010203"),
		},
	}

	for _, c := range cases {
		var buf bytes.Buffer
		lval := c.item
		w := NewWriter(&buf)

		if err := EncodeItemKey(c.item, c.keydef, w); err != nil {
			t.Errorf("unexpected error %v", err)
			continue
		}
		if err := w.Flush(); err != nil {
			t.Errorf("unexpected error %v", err)
			continue
		}

		if c.enc != nil {
			b := buf.Bytes()
			if !reflect.DeepEqual(c.enc, b) {
				t.Errorf("expected %v, actual %v for for %v", c.enc, b, c.item)
			}
		}

		r := NewReader(&buf)
		rval, err := DecodeItemKey(r, c.keydef)
		if err != nil {
			t.Errorf("unexpected error %v", err)
			continue
		}

		if !reflect.DeepEqual(lval, rval) {
			t.Errorf("expected: %v, actual: %v", lval, rval)
		}

		w.Close()
	}
}

func TestItemNonKeyAttributes(t *testing.T) {
	keydef := []types.AttributeDefinition{
		{AttributeName: aws.String("hks"), AttributeType: types.ScalarAttributeTypeS},
		{AttributeName: aws.String("rkn"), AttributeType: types.ScalarAttributeTypeN},
	}
	item := map[string]types.AttributeValue{
		"hks": &types.AttributeValueMemberS{Value: "hkv"},
		"rkn": &types.AttributeValueMemberN{Value: "123"},
		"av1": &types.AttributeValueMemberS{Value: "avs"},
		"av2": &types.AttributeValueMemberN{Value: "456"},
		"av3": &types.AttributeValueMemberB{Value: fromHex("0x010203")},
	}
	attrNames := []string{"av1", "av2", "av3"}
	var attrListId int64 = 1
	km := func(key lru.Key) lru.Key {
		return fmt.Sprintf("%q", key)
	}
	attrNamesListToId := &lru.Lru{
		LoadFunc: func(ctx context.Context, key lru.Key) (interface{}, error) {
			an := key.([]string)
			if !reflect.DeepEqual(an, attrNames) {
				return nil, errors.New(fmt.Sprintf("unknown attribute list %v %v", an, strings.Join(an, ",")))
			}
			return attrListId, nil
		},
		KeyMarshaller: km,
	}
	attrListIdToNames := &lru.Lru{
		LoadFunc: func(ctx context.Context, key lru.Key) (interface{}, error) {
			id := key.(int64)
			if id != attrListId {
				return nil, errors.New(fmt.Sprintf("unknown attribute list id %v", id))
			}
			return attrNames, nil
		},
	}

	var buf bytes.Buffer
	w := NewWriter(&buf)
	if err := EncodeItemNonKeyAttributes(nil, item, keydef, attrNamesListToId, w); err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if err := w.Flush(); err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	r := NewReader(&buf)
	actual, err := DecodeItemNonKeyAttributes(nil, r, attrListIdToNames)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	expected := make(map[string]types.AttributeValue)
	for k, v := range item {
		if k != *keydef[0].AttributeName && k != *keydef[1].AttributeName {
			expected[k] = v
		}
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected: %v, actual: %v", expected, actual)
	}
}
