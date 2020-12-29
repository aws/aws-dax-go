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
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"math/big"
	"strconv"
	"strings"
)

const (
	tagStringSet = 3321 + iota
	tagNumberSet
	tagBinarySet
	tagDocumentPathOrdinal
)

func EncodeAttributeValue(value *dynamodb.AttributeValue, writer *Writer) error {
	if value == nil {
		return awserr.New(request.InvalidParameterErrCode, "invalid attribute value: nil", nil)
	}

	var err error
	switch {
	case value.S != nil:
		err = writer.WriteString(*value.S)
	case value.N != nil:
		err = writeStringNumber(*value.N, writer)
	case value.B != nil:
		err = writer.WriteBytes(value.B)
	case value.SS != nil:
		if err = writer.writeType(Tag, tagStringSet); err != nil {
			return err
		}
		if err = writer.WriteArrayHeader(len(value.SS)); err != nil {
			return err
		}
		for _, sp := range value.SS {
			if err := writer.WriteString(*sp); err != nil {
				return err
			}
		}
	case value.NS != nil:
		if err = writer.writeType(Tag, tagNumberSet); err != nil {
			return err
		}
		if err = writer.WriteArrayHeader(len(value.NS)); err != nil {
			return err
		}
		for _, sp := range value.NS {
			if err := writeStringNumber(*sp, writer); err != nil {
				return err
			}
		}
	case value.BS != nil:
		if err = writer.writeType(Tag, tagBinarySet); err != nil {
			return err
		}
		if err = writer.WriteArrayHeader(len(value.BS)); err != nil {
			return err
		}
		for _, bp := range value.BS {
			if err := writer.WriteBytes(bp); err != nil {
				return err
			}
		}
	case value.L != nil:
		if err = writer.WriteArrayHeader(len(value.L)); err != nil {
			return err
		}
		for _, v := range value.L {
			if err := EncodeAttributeValue(v, writer); err != nil {
				return err
			}
		}
	case value.M != nil:
		if err = writer.WriteMapHeader(len(value.M)); err != nil {
			return err
		}
		for k, v := range value.M {
			if err := writer.WriteString(k); err != nil {
				return err
			}
			if err = EncodeAttributeValue(v, writer); err != nil {
				return err
			}
		}
	case value.BOOL != nil:
		err = writer.WriteBoolean(*value.BOOL)
	case value.NULL != nil:
		if !(*value.NULL) {
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

func DecodeAttributeValue(reader *Reader) (*dynamodb.AttributeValue, error) {
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
		return &dynamodb.AttributeValue{S: &s}, nil
	case Bytes:
		b, err := reader.ReadBytes()
		if err != nil {
			return nil, err
		}
		return &dynamodb.AttributeValue{B: b}, nil
	case Array:
		len, err := reader.ReadArrayLength()
		if err != nil {
			return nil, err
		}
		as := make([]*dynamodb.AttributeValue, len)
		for i := 0; i < len; i++ {
			a, err := DecodeAttributeValue(reader)
			if err != nil {
				return nil, err
			}
			as[i] = a
		}
		return &dynamodb.AttributeValue{L: as}, nil
	case Map:
		len, err := reader.ReadMapLength()
		if err != nil {
			return nil, err
		}
		m := make(map[string]*dynamodb.AttributeValue, len)
		for i := 0; i < len; i++ {
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
		return &dynamodb.AttributeValue{M: m}, nil
	case PosInt, NegInt:
		s, err := reader.ReadCborIntegerToString()
		if err != nil {
			return nil, err
		}
		return &dynamodb.AttributeValue{N: &s}, nil
	case Simple:
		if _, _, err := reader.readTypeHeader(); err != nil {
			return nil, err
		}
		switch hdr {
		case False:
			return &dynamodb.AttributeValue{BOOL: aws.Bool(false)}, nil
		case True:
			return &dynamodb.AttributeValue{BOOL: aws.Bool(true)}, nil
		case Nil:
			return &dynamodb.AttributeValue{NULL: aws.Bool(true)}, nil
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
			return &dynamodb.AttributeValue{N: aws.String(i.String())}, nil
		case TagDecimal:
			d, err := reader.ReadDecimal()
			if err != nil {
				return nil, err
			}
			return &dynamodb.AttributeValue{N: aws.String(d.String())}, nil
		default:
			_, tag, err := reader.readTypeHeader()
			if err != nil {
				return nil, err
			}
			switch tag {
			case tagStringSet:
				len, err := reader.ReadArrayLength()
				if err != nil {
					return nil, err
				}
				ss := make([]*string, len)
				for i := 0; i < len; i++ {
					s, err := reader.ReadString()
					if err != nil {
						return nil, err
					}
					ss[i] = &s
				}
				return &dynamodb.AttributeValue{SS: ss}, nil
			case tagNumberSet:
				len, err := reader.ReadArrayLength()
				if err != nil {
					return nil, err
				}
				ss := make([]*string, len)
				for i := 0; i < len; i++ {
					av, err := DecodeAttributeValue(reader)
					if err != nil {
						return nil, err
					}
					ss[i] = av.N
				}
				return &dynamodb.AttributeValue{NS: ss}, nil
			case tagBinarySet:
				len, err := reader.ReadArrayLength()
				if err != nil {
					return nil, err
				}
				bs := make([][]byte, len)
				for i := 0; i < len; i++ {
					b, err := reader.ReadBytes()
					if err != nil {
						return nil, err
					}
					bs[i] = b
				}
				return &dynamodb.AttributeValue{BS: bs}, nil
			default:
				return nil, awserr.New(request.ErrCodeSerialization, fmt.Sprintf("unknown minor type %d or tag %d", minor, tag), nil)
			}
		}
	default:
		return nil, awserr.New(request.ErrCodeSerialization, fmt.Sprintf("unknown major type %d", major), nil)
	}
}
