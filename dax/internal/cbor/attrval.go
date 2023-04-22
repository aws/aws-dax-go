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
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
)

const (
	tagStringSet = 3321 + iota
	tagNumberSet
	tagBinarySet
	tagDocumentPathOrdinal
)

func EncodeAttributeValue(value types.AttributeValue, writer *Writer) error {
	if value == nil {
		return awserr.New(request.InvalidParameterErrCode, "invalid attribute value: nil", nil)
	}

	var err error
	switch v := value.(type) {
	case *types.AttributeValueMemberS:
		err = writer.WriteString(v.Value)
	case *types.AttributeValueMemberN:
		err = writeStringNumber(v.Value, writer)
	case *types.AttributeValueMemberB:
		err = writer.WriteBytes(v.Value)
	case *types.AttributeValueMemberSS:
		if err = writer.writeType(Tag, tagStringSet); err != nil {
			return err
		}
		if err = writer.WriteArrayHeader(len(v.Value)); err != nil {
			return err
		}
		for _, sp := range v.Value {
			if err := writer.WriteString(sp); err != nil {
				return err
			}
		}
	case *types.AttributeValueMemberNS:
		if err = writer.writeType(Tag, tagNumberSet); err != nil {
			return err
		}
		if err = writer.WriteArrayHeader(len(v.Value)); err != nil {
			return err
		}
		for _, sp := range v.Value {
			if err := writeStringNumber(sp, writer); err != nil {
				return err
			}
		}
	case *types.AttributeValueMemberBS:
		if err = writer.writeType(Tag, tagBinarySet); err != nil {
			return err
		}
		if err = writer.WriteArrayHeader(len(v.Value)); err != nil {
			return err
		}
		for _, bp := range v.Value {
			if err := writer.WriteBytes(bp); err != nil {
				return err
			}
		}
	case *types.AttributeValueMemberL:
		if err = writer.WriteArrayHeader(len(v.Value)); err != nil {
			return err
		}
		for _, v := range v.Value {
			if err := EncodeAttributeValue(v, writer); err != nil {
				return err
			}
		}
	case *types.AttributeValueMemberM:
		if err = writer.WriteMapHeader(len(v.Value)); err != nil {
			return err
		}
		for k, v := range v.Value {
			if err := writer.WriteString(k); err != nil {
				return err
			}
			if err = EncodeAttributeValue(v, writer); err != nil {
				return err
			}
		}
	case *types.AttributeValueMemberBOOL:
		err = writer.WriteBoolean(v.Value)
	case *types.AttributeValueMemberNULL:
		if !v.Value {
			return awserr.New(request.InvalidParameterErrCode, "invalid null attribute value", nil) // DaxJavaClient suppress this error
		}
		err = writer.WriteNull()
	}
	return err
}

func writeStringNumber(val string, writer *Writer) error {
	if strings.IndexAny(val, ".eE") >= 0 {
		dec := new(Decimal)
		if _, ok := dec.SetString(val); !ok {
			return awserr.New(request.InvalidParameterErrCode, fmt.Sprintf("invalid number %v", val), nil)
		}
		err := writer.WriteDecimal(dec)
		return err
	}
	if len(val) > 18 {
		bint := new(big.Int)
		bint.SetString(val, 10)
		err := writer.WriteBigInt(bint)
		return err
	}
	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return awserr.New(request.InvalidParameterErrCode, fmt.Sprintf("invalid number %v", val), err)
	}
	err = writer.WriteInt64(i)
	return err
}

func DecodeAttributeValue(reader *Reader) (types.AttributeValue, error) {
	hdr, err := reader.PeekHeader()
	if err != nil {
		return nil, err
	}
	major := hdr & MajorTypeMask
	minor := hdr & MinorTypeMask

	switch major {
	case Utf:
		s, err := reader.ReadString()
		if err != nil {
			return nil, err
		}
		return &types.AttributeValueMemberS{Value: s}, nil
	case Bytes:
		b, err := reader.ReadBytes()
		if err != nil {
			return nil, err
		}
		return &types.AttributeValueMemberB{Value: b}, nil
	case Array:
		length, err := reader.ReadArrayLength()
		if err != nil {
			return nil, err
		}
		as := make([]types.AttributeValue, length)
		for i := 0; i < length; i++ {
			a, err := DecodeAttributeValue(reader)
			if err != nil {
				return nil, err
			}
			as[i] = a
		}
		return &types.AttributeValueMemberL{Value: as}, nil
	case Map:
		length, err := reader.ReadMapLength()
		if err != nil {
			return nil, err
		}
		m := make(map[string]types.AttributeValue, length)
		for i := 0; i < length; i++ {
			k, err := reader.ReadString()
			if err != nil {
				return nil, err
			}
			v, err := DecodeAttributeValue(reader)
			if err != nil {
				return nil, err
			}
			m[k] = v
		}
		return &types.AttributeValueMemberM{Value: m}, nil
	case PosInt, NegInt:
		s, err := reader.ReadCborIntegerToString()
		if err != nil {
			return nil, err
		}
		return &types.AttributeValueMemberN{Value: s}, nil
	case Simple:
		if _, _, err := reader.readTypeHeader(); err != nil {
			return nil, err
		}
		switch hdr {
		case False:
			return &types.AttributeValueMemberBOOL{Value: false}, nil
		case True:
			return &types.AttributeValueMemberBOOL{Value: true}, nil
		case Nil:
			return &types.AttributeValueMemberNULL{Value: true}, nil
		default:
			return nil, awserr.New(request.ErrCodeSerialization, fmt.Sprintf("unknown minor type %d for simple major type", minor), nil)
		}
	case Tag:
		switch minor {
		case TagPosBigInt, TagNegBigInt:
			i, err := reader.ReadBigInt()
			if err != nil {
				return nil, err
			}
			return &types.AttributeValueMemberN{Value: i.String()}, nil
		case TagDecimal:
			d, err := reader.ReadDecimal()
			if err != nil {
				return nil, err
			}
			return &types.AttributeValueMemberN{Value: d.String()}, nil
		default:
			_, tag, err := reader.readTypeHeader()
			if err != nil {
				return nil, err
			}
			switch tag {
			case tagStringSet:
				length, err := reader.ReadArrayLength()
				if err != nil {
					return nil, err
				}
				ss := make([]string, length)
				for i := 0; i < length; i++ {
					s, err := reader.ReadString()
					if err != nil {
						return nil, err
					}
					ss[i] = s
				}
				return &types.AttributeValueMemberSS{Value: ss}, nil
			case tagNumberSet:
				length, err := reader.ReadArrayLength()
				if err != nil {
					return nil, err
				}
				ss := make([]string, length)
				for i := 0; i < length; i++ {
					av, err := DecodeAttributeValue(reader)
					if err != nil {
						return nil, err
					}
					n, ok := av.(*types.AttributeValueMemberN)
					if !ok {
						return nil, awserr.New(request.ErrCodeSerialization, fmt.Sprintf("attribute type is not number. type: %T", av), nil)
					}
					ss[i] = n.Value
				}
				return &types.AttributeValueMemberNS{Value: ss}, nil
			case tagBinarySet:
				length, err := reader.ReadArrayLength()
				if err != nil {
					return nil, err
				}
				bs := make([][]byte, length)
				for i := 0; i < length; i++ {
					b, err := reader.ReadBytes()
					if err != nil {
						return nil, err
					}
					bs[i] = b
				}
				return &types.AttributeValueMemberBS{Value: bs}, nil
			default:
				return nil, awserr.New(request.ErrCodeSerialization, fmt.Sprintf("unknown minor type %d or tag %d", minor, tag), nil)
			}
		}
	default:
		return nil, awserr.New(request.ErrCodeSerialization, fmt.Sprintf("unknown major type %d", major), nil)
	}
}
