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
	"sort"

	"github.com/aws/aws-dax-go/dax/internal/lru"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var ErrMissingKey = errors.New("one of the required keys was not given a value")

func EncodeItemKey(item map[string]types.AttributeValue, keydef []types.AttributeDefinition, writer *Writer) error {
	keyBytes, err := GetEncodedItemKey(item, keydef)
	if err != nil {
		return err
	}
	return writer.WriteBytes(keyBytes)
}

func GetEncodedItemKey(item map[string]types.AttributeValue, keydef []types.AttributeDefinition) ([]byte, error) {
	if item == nil {
		return nil, errors.New("item cannot be nil")
	}

	hk := keydef[0]
	hkval, found := item[*hk.AttributeName]
	if !found {
		return nil, ErrMissingKey
	}

	var buf bytes.Buffer
	w := NewWriter(&buf)
	defer w.Close()

	if len(keydef) == 1 {
		switch hk.AttributeType {
		case types.ScalarAttributeTypeS:
			sp, ok := hkval.(*types.AttributeValueMemberS)
			if !ok {
				return nil, ErrMissingKey
			}
			if err := w.Write([]byte(sp.Value)); err != nil {
				return nil, err
			}
		case types.ScalarAttributeTypeN:
			_, ok := hkval.(*types.AttributeValueMemberN)
			if !ok {
				return nil, ErrMissingKey
			}
			if err := EncodeAttributeValue(hkval, w); err != nil {
				return nil, err
			}
		case types.ScalarAttributeTypeB:
			b, ok := hkval.(*types.AttributeValueMemberB)
			if !ok {
				return nil, ErrMissingKey
			}
			if err := w.Write(b.Value); err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("unsupported KeyType encountered in Hash Attribute: %s", hk.AttributeType)
		}
	} else {
		switch hk.AttributeType {
		case types.ScalarAttributeTypeS:
			sp, ok := hkval.(*types.AttributeValueMemberS)
			if !ok {
				return nil, ErrMissingKey
			}
			if err := w.WriteString(sp.Value); err != nil {
				return nil, err
			}
		case types.ScalarAttributeTypeN:
			_, ok := hkval.(*types.AttributeValueMemberN)
			if !ok {
				return nil, ErrMissingKey
			}
			if err := EncodeAttributeValue(hkval, w); err != nil {
				return nil, err
			}
		case types.ScalarAttributeTypeB:
			b, ok := hkval.(*types.AttributeValueMemberB)
			if !ok {
				return nil, ErrMissingKey
			}
			if err := w.WriteBytes(b.Value); err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("unsupported KeyType encountered in Hash Attribute: %s", hk.AttributeType)
		}

		rk := keydef[1]
		rkval, found := item[*rk.AttributeName]
		if !found {
			return nil, ErrMissingKey
		}
		switch rk.AttributeType {
		case types.ScalarAttributeTypeS:
			sp, ok := rkval.(*types.AttributeValueMemberS)
			if !ok {
				return nil, ErrMissingKey
			}
			if err := w.Write([]byte(sp.Value)); err != nil {
				return nil, err
			}
		case types.ScalarAttributeTypeN:
			n, ok := rkval.(*types.AttributeValueMemberN)
			if !ok {
				return nil, ErrMissingKey
			}
			d := new(Decimal)
			d, ok = d.SetString(n.Value)
			if !ok {
				return nil, errors.New("invalid number " + n.Value)
			}
			if _, err := EncodeLexDecimal(d, w.bw); err != nil {
				return nil, err
			}
		case types.ScalarAttributeTypeB:
			b, ok := rkval.(*types.AttributeValueMemberB)
			if !ok {
				return nil, ErrMissingKey
			}
			if err := w.Write(b.Value); err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("unsupported KeyType encountered in Range Attribute: %s", rk.AttributeType)
		}
	}

	if err := w.Flush(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func DecodeItemKey(reader *Reader, keydef []types.AttributeDefinition) (map[string]types.AttributeValue, error) {
	hk := keydef[0]
	keys := make(map[string]types.AttributeValue)

	if len(keydef) == 1 {
		switch hk.AttributeType {
		case types.ScalarAttributeTypeS:
			kb, err := reader.ReadBytes()
			if err != nil {
				return nil, err
			}
			s := string(kb)
			keys[*hk.AttributeName] = &types.AttributeValueMemberS{Value: s}
		case types.ScalarAttributeTypeN:
			r, err := reader.BytesReader()
			if err != nil {
				return nil, err
			}
			defer r.Close()
			av, err := DecodeAttributeValue(r)
			if err != nil {
				return nil, err
			}
			_, ok := av.(*types.AttributeValueMemberN)
			if !ok {
				return nil, ErrMissingKey
			}
			keys[*hk.AttributeName] = av
		case types.ScalarAttributeTypeB:
			kb, err := reader.ReadBytes()
			if err != nil {
				return nil, err
			}
			keys[*hk.AttributeName] = &types.AttributeValueMemberB{Value: kb}
		default:
			return nil, fmt.Errorf("unsupported KeyType encountered in Hash Attribute: %s", hk.AttributeType)
		}
	} else {
		r, err := reader.BytesReader()
		if err != nil {
			return nil, err
		}
		defer r.Close()
		switch hk.AttributeType {
		case types.ScalarAttributeTypeS:
			s, err := r.ReadString()
			if err != nil {
				return nil, err
			}
			keys[*hk.AttributeName] = &types.AttributeValueMemberS{Value: s}
		case types.ScalarAttributeTypeN:
			av, err := DecodeAttributeValue(r)
			if err != nil {
				return nil, err
			}
			_, ok := av.(*types.AttributeValueMemberN)
			if !ok {
				return nil, ErrMissingKey
			}
			keys[*hk.AttributeName] = av
		case types.ScalarAttributeTypeB:
			b, err := r.ReadBytes()
			if err != nil {
				return nil, err
			}
			keys[*hk.AttributeName] = &types.AttributeValueMemberB{Value: b}
		default:
			return nil, fmt.Errorf("unsupported KeyType encountered in Hash Attribute: %s", hk.AttributeType)
		}

		rk := keydef[1]
		switch rk.AttributeType {
		case types.ScalarAttributeTypeS:
			var buf bytes.Buffer
			if _, err := r.br.WriteTo(&buf); err != nil {
				return nil, err
			}
			s := string(buf.Bytes())
			keys[*rk.AttributeName] = &types.AttributeValueMemberS{Value: s}
		case types.ScalarAttributeTypeN:
			d, err := DecodeLexDecimal(r.br)
			if err != nil {
				return nil, err
			}
			s := d.String()
			keys[*rk.AttributeName] = &types.AttributeValueMemberN{Value: s}
		case types.ScalarAttributeTypeB:
			var buf bytes.Buffer
			if _, err := r.br.WriteTo(&buf); err != nil {
				return nil, err
			}
			keys[*rk.AttributeName] = &types.AttributeValueMemberB{Value: buf.Bytes()}
		default:
			return nil, fmt.Errorf("unsupported KeyType encountered in Range Attribute: %s", rk.AttributeType)
		}
	}

	return keys, nil
}

func EncodeItemNonKeyAttributes(ctx context.Context, item map[string]types.AttributeValue, keydef []types.AttributeDefinition,
	attrNamesListToId *lru.Lru, writer *Writer) error {

	keydeflen := len(keydef)
	nonKeyAttrNames := make([]string, 0, len(item)-keydeflen)
	for k := range item {
		if k != *keydef[0].AttributeName && (keydeflen == 1 || k != *keydef[1].AttributeName) {
			nonKeyAttrNames = append(nonKeyAttrNames, k)
		}
	}
	sort.Strings(nonKeyAttrNames)

	nonKeyAttrValues := make([]types.AttributeValue, len(nonKeyAttrNames))
	for i, k := range nonKeyAttrNames {
		nonKeyAttrValues[i] = item[k]
	}

	id, err := attrNamesListToId.GetWithContext(ctx, nonKeyAttrNames)
	if err != nil {
		return err
	}

	if err = writer.WriteInt64(id.(int64)); err != nil {
		return err
	}
	for _, v := range nonKeyAttrValues {
		if err := EncodeAttributeValue(v, writer); err != nil {
			return err
		}
	}

	return nil
}

func DecodeItemNonKeyAttributes(ctx context.Context, reader *Reader, attrListIdToNames *lru.Lru) (map[string]types.AttributeValue, error) {
	id, err := reader.ReadInt64()
	if err != nil {
		return nil, err
	}
	attrNames, err := attrListIdToNames.GetWithContext(ctx, id)
	if err != nil {
		return nil, err
	}

	attrs := make(map[string]types.AttributeValue)
	for _, n := range attrNames.([]string) {
		av, err := DecodeAttributeValue(reader)
		if err != nil {
			return nil, err
		}
		attrs[n] = av
	}
	return attrs, nil
}
