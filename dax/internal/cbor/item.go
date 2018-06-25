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
	"fmt"
	"github.com/aws/aws-dax-go/dax/internal/lru"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"sort"
)

var ErrMissingKey = awserr.New(request.ParamRequiredErrCode, "One of the required keys was not given a value", nil)

func EncodeItemKey(item map[string]*dynamodb.AttributeValue, keydef []dynamodb.AttributeDefinition, writer *Writer) error {
	if item == nil {
		return awserr.New(request.InvalidParameterErrCode, "item cannot be nil", nil)
	}

	hk := keydef[0]
	hkval, ok := item[*hk.AttributeName]
	if !ok {
		return ErrMissingKey
	}

	var buf bytes.Buffer
	w := NewWriter(&buf)
	defer w.Close()

	if len(keydef) == 1 {
		switch *hk.AttributeType {
		case dynamodb.ScalarAttributeTypeS:
			sp := hkval.S
			if sp == nil {
				return ErrMissingKey
			}
			if err := w.Write([]byte(*sp)); err != nil {
				return err
			}
		case dynamodb.ScalarAttributeTypeN:
			if hkval.N == nil {
				return ErrMissingKey
			}
			if err := EncodeAttributeValue(hkval, w); err != nil {
				return err
			}
		case dynamodb.ScalarAttributeTypeB:
			b := hkval.B
			if b == nil {
				return ErrMissingKey
			}
			if err := w.Write(b); err != nil {
				return err
			}
		default:
			return awserr.New(request.InvalidParameterErrCode, fmt.Sprintf("Unsupported KeyType encountered in Hash Attribute: "+*hk.AttributeType), nil)
		}
	} else {
		switch *hk.AttributeType {
		case dynamodb.ScalarAttributeTypeS:
			sp := hkval.S
			if sp == nil {
				return ErrMissingKey
			}
			if err := w.WriteString(*sp); err != nil {
				return err
			}
		case dynamodb.ScalarAttributeTypeN:
			if hkval.N == nil {
				return ErrMissingKey
			}
			if err := EncodeAttributeValue(hkval, w); err != nil {
				return err
			}
		case dynamodb.ScalarAttributeTypeB:
			b := hkval.B
			if b == nil {
				return ErrMissingKey
			}
			if err := w.WriteBytes(b); err != nil {
				return err
			}
		default:
			return awserr.New(request.InvalidParameterErrCode, fmt.Sprintf("Unsupported KeyType encountered in Hash Attribute: "+*hk.AttributeType), nil)
		}

		rk := keydef[1]
		rkval, ok := item[*rk.AttributeName]
		if !ok {
			return ErrMissingKey
		}
		switch *rk.AttributeType {
		case dynamodb.ScalarAttributeTypeS:
			sp := rkval.S
			if sp == nil {
				return ErrMissingKey
			}
			if err := w.Write([]byte(*sp)); err != nil {
				return err
			}
		case dynamodb.ScalarAttributeTypeN:
			n := rkval.N
			if n == nil {
				return ErrMissingKey
			}
			d := new(Decimal)
			d, ok := d.SetString(*n)
			if !ok {
				return awserr.New(request.InvalidParameterErrCode, fmt.Sprintf("invalid number "+*n), nil)
			}
			if _, err := EncodeLexDecimal(d, w.bw); err != nil {
				return err
			}
		case dynamodb.ScalarAttributeTypeB:
			b := rkval.B
			if b == nil {
				return ErrMissingKey
			}
			if err := w.Write(b); err != nil {
				return err
			}
		default:
			return awserr.New(request.InvalidParameterErrCode, fmt.Sprintf("Unsupported KeyType encountered in Range Attribute: "+*rk.AttributeType), nil)
		}
	}

	if err := w.Flush(); err != nil {
		return err
	}
	return writer.WriteBytes(buf.Bytes())
}

func DecodeItemKey(reader *Reader, keydef []dynamodb.AttributeDefinition) (map[string]*dynamodb.AttributeValue, error) {
	hk := keydef[0]
	keys := make(map[string]*dynamodb.AttributeValue)

	if len(keydef) == 1 {
		switch *hk.AttributeType {
		case dynamodb.ScalarAttributeTypeS:
			kb, err := reader.ReadBytes()
			if err != nil {
				return nil, err
			}
			s := string(kb)
			keys[*hk.AttributeName] = &dynamodb.AttributeValue{S: &s}
		case dynamodb.ScalarAttributeTypeN:
			r, err := reader.BytesReader()
			if err != nil {
				return nil, err
			}
			defer r.Close()
			av, err := DecodeAttributeValue(r)
			if err != nil {
				return nil, err
			}
			if av.N == nil {
				return nil, ErrMissingKey
			}
			keys[*hk.AttributeName] = av
		case dynamodb.ScalarAttributeTypeB:
			kb, err := reader.ReadBytes()
			if err != nil {
				return nil, err
			}
			keys[*hk.AttributeName] = &dynamodb.AttributeValue{B: kb}
		default:
			return nil, awserr.New(request.InvalidParameterErrCode, fmt.Sprintf("Unsupported KeyType encountered in Hash Attribute: "+*hk.AttributeType), nil)
		}
	} else {
		r, err := reader.BytesReader()
		if err != nil {
			return nil, err
		}
		defer r.Close()
		switch *hk.AttributeType {
		case dynamodb.ScalarAttributeTypeS:
			s, err := r.ReadString()
			if err != nil {
				return nil, err
			}
			keys[*hk.AttributeName] = &dynamodb.AttributeValue{S: &s}
		case dynamodb.ScalarAttributeTypeN:
			av, err := DecodeAttributeValue(r)
			if err != nil {
				return nil, err
			}
			if av.N == nil {
				return nil, ErrMissingKey
			}
			keys[*hk.AttributeName] = av
		case dynamodb.ScalarAttributeTypeB:
			b, err := r.ReadBytes()
			if err != nil {
				return nil, err
			}
			keys[*hk.AttributeName] = &dynamodb.AttributeValue{B: b}
		default:
			return nil, awserr.New(request.InvalidParameterErrCode, fmt.Sprintf("Unsupported KeyType encountered in Hash Attribute: "+*hk.AttributeType), nil)
		}

		rk := keydef[1]
		switch *rk.AttributeType {
		case dynamodb.ScalarAttributeTypeS:
			var buf bytes.Buffer
			if _, err := r.br.WriteTo(&buf); err != nil {
				return nil, err
			}
			s := string(buf.Bytes())
			keys[*rk.AttributeName] = &dynamodb.AttributeValue{S: &s}
		case dynamodb.ScalarAttributeTypeN:
			d, err := DecodeLexDecimal(r.br)
			if err != nil {
				return nil, err
			}
			s := d.String()
			keys[*rk.AttributeName] = &dynamodb.AttributeValue{N: &s}
		case dynamodb.ScalarAttributeTypeB:
			var buf bytes.Buffer
			if _, err := r.br.WriteTo(&buf); err != nil {
				return nil, err
			}
			keys[*rk.AttributeName] = &dynamodb.AttributeValue{B: buf.Bytes()}
		default:
			return nil, awserr.New(request.InvalidParameterErrCode, fmt.Sprintf("Unsupported KeyType encountered in Range Attribute: "+*rk.AttributeType), nil)
		}
	}

	return keys, nil
}

func EncodeItemNonKeyAttributes(ctx aws.Context, item map[string]*dynamodb.AttributeValue, keydef []dynamodb.AttributeDefinition,
	attrNamesListToId *lru.Lru, writer *Writer) error {

	keydeflen := len(keydef)
	nonKeyAttrNames := make([]string, 0, len(item)-keydeflen)
	for k, _ := range item {
		if k != *keydef[0].AttributeName && (keydeflen == 1 || k != *keydef[1].AttributeName) {
			nonKeyAttrNames = append(nonKeyAttrNames, k)
		}
	}
	sort.Strings(nonKeyAttrNames)

	nonKeyAttrValues := make([]*dynamodb.AttributeValue, len(nonKeyAttrNames))
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

func DecodeItemNonKeyAttributes(ctx aws.Context, reader *Reader, attrListIdToNames *lru.Lru) (map[string]*dynamodb.AttributeValue, error) {
	id, err := reader.ReadInt64()
	if err != nil {
		return nil, err
	}
	attrNames, err := attrListIdToNames.GetWithContext(ctx, id)
	if err != nil {
		return nil, err
	}

	attrs := make(map[string]*dynamodb.AttributeValue)
	for _, n := range attrNames.([]string) {
		av, err := DecodeAttributeValue(reader)
		if err != nil {
			return nil, err
		}
		attrs[n] = av
	}
	return attrs, nil
}
