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
	"encoding/binary"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"io"
	"math"
	"math/big"
)

/* Encoding of header:
0x00:       null low (unused)
0x01:       negative signum; four bytes follow for positive exponent
0x02..0x3f: negative signum; positive exponent; 3e range, 61..0
0x40..0x7d: negative signum; negative exponent; 3e range, -1..-62
0x7e:       negative signum; four bytes follow for negative exponent
0x7f:       negative zero (unused)
0x80:       zero
0x81:       positive signum; four bytes follow for negative exponent
0x82..0xbf: positive signum; negative exponent; 3e range, -62..-1
0xc0..0xfd: positive signum; positive exponent; 3e range, 0..61
0xfe:       positive signum; four bytes follow for positive exponent
0xff:       null high

Significand must be decimal encoded to maintain proper sort order. Base 1000 is more
efficient than base 10 and still maintains proper sort order. A minimum of two bytes must
be generated, however.
*/

const nullLow = 0x00
const nullHigh = 0xff
const log1000M = math.Ln2 / math.Ln10

var billion = big.NewInt(1000000000)
var thousand = big.NewInt(1000)
var hundred = big.NewInt(100)
var ten = big.NewInt(10)

type BytesWriter interface {
	io.Writer
	io.ByteWriter
}

func EncodeLexDecimal(decimal *Decimal, writer BytesWriter) (int, error) {
	return encode(decimal, writer, 0)
}

func encode(decimal *Decimal, writer BytesWriter, xormask int) (int, error) {
	if decimal.Unscaled().Sign() == 0 {
		if err := writer.WriteByte(byte(0x80 ^ xormask)); err != nil {
			return 0, err
		}
		return 1, nil
	}

	len := 0
	precision := precision(decimal)
	exponent := precision - int(decimal.scale)
	val := decimal.Unscaled()

	if val.Sign() < 0 {
		if exponent >= -0x3e && exponent < 0x3e {
			if err := writer.WriteByte(byte((0x3f - exponent) ^ xormask)); err != nil {
				return 0, err
			}
			len++
		} else {
			if exponent < 0 {
				if err := writer.WriteByte(byte(0x7e ^ xormask)); err != nil {
					return 0, err
				}
			} else {
				if err := writer.WriteByte(byte(1 ^ xormask)); err != nil {
					return 0, err
				}
			}
			if err := encodeInt32BE(exponent^xormask^0x7fffffff, writer); err != nil {
				return 0, err
			}
			len += 5
		}
	} else {
		if exponent >= -0x3e && exponent < 0x3e {
			if err := writer.WriteByte(byte((exponent + 0xc0) ^ xormask)); err != nil {
				return 0, err
			}
			len++
		} else {
			if exponent < 0 {
				if err := writer.WriteByte(byte(0x81 ^ xormask)); err != nil {
					return 0, err
				}
			} else {
				if err := writer.WriteByte(byte(0xfe ^ xormask)); err != nil {
					return 0, err
				}
			}
			if err := encodeInt32BE(exponent^xormask^0x80000000, writer); err != nil {
				return 0, err
			}
			len += 5
		}
	}

	var terminator int
	switch precision % 3 {
	case 0:
		terminator = 2
	case 1:
		terminator = 0
		val = val.Mul(val, hundred)
	case 2:
		terminator = 1
		val = val.Mul(val, ten)
	default:
		terminator = 2
	}

	var digitAdjust int
	if val.Sign() >= 0 {
		digitAdjust = 12
	} else {
		digitAdjust = 999 + 12
		terminator = 1023 - terminator
	}

	pos := ((val.BitLen() + 9) / 10) + 1
	digits := make([]int, pos)
	pos--
	digits[pos] = terminator

	var rem big.Int
	for val.Sign() != 0 {
		val.QuoRem(val, thousand, &rem)

		pos--
		v := int(rem.Int64()) + digitAdjust
		if pos < 0 {
			digits = append([]int{v}, digits...)
		} else {
			digits[pos] = v
		}
	}

	accum := 0
	var bits uint = 0
	for _, v := range digits {
		accum = accum<<10 | v
		bits += 10
		for {
			bits -= 8
			if err := writer.WriteByte(byte((accum >> bits) ^ xormask)); err != nil {
				return 0, err
			}
			len++
			if bits < 8 {
				break
			}
		}
	}

	if bits != 0 {
		if err := writer.WriteByte(byte((accum << uint(8-bits)) ^ xormask)); err != nil {
			return 0, err
		}
		len++
	}
	return len, nil
}

var numDigits = []int64{0, 9, 99, 999, 9999, 99999, 999999, 9999999, 99999999, 999999999, 1000000000}

func precision(decimal *Decimal) int {
	big := new(big.Int).Set(decimal.Unscaled())
	if big.Sign() == 0 {
		return 1
	}
	big = big.Abs(big)

	digits := 0
	for big.Cmp(billion) > 0 {
		big = big.Quo(big, billion)
		digits += 9
	}

	small := big.Int64()
	for i, d := range numDigits {
		if small <= d {
			return digits + i
		}
	}

	// unreachable code
	return digits
}

type BytesReader interface {
	io.Reader
	io.ByteReader
}

func DecodeLexDecimal(reader BytesReader) (*Decimal, error) {
	return decode(reader, 0)
}

func decode(reader BytesReader, xormask int) (*Decimal, error) {
	b, err := reader.ReadByte()
	if err != nil {
		return nil, err
	}
	header := (int(b) ^ xormask) & 0xff

	var digitAdjust int
	var exponent int

	switch header {
	case nullHigh & 0xff, nullLow & 0xff:
		return nil, nil

	case 0x7f, 0x80:
		d0 := NewDecimal(big.NewInt(0), 0)
		return d0, nil

	case 1, 0x7e:
		digitAdjust = 999 + 12
		v, err := decodeInt32BE(reader)
		if err != nil {
			return nil, err
		}
		exponent = v ^ xormask ^ 0x7fffffff

	case 0x81, 0xfe:
		digitAdjust = 12
		v, err := decodeInt32BE(reader)
		if err != nil {
			return nil, err
		}
		exponent = v ^ xormask ^ 0x80000000

	default:
		exponent = (int(b) ^ xormask) & 0xff
		if exponent >= 0x82 {
			digitAdjust = 12
			exponent -= 0xc0
		} else {
			digitAdjust = 999 + 12
			exponent = 0x3f - exponent
		}
	}

	precision := 0
	accum := 0
	var bits uint = 0
	var lastDigit *big.Int
	var unscaledValue *big.Int

	done := false
	for !done {
		b, err := reader.ReadByte()
		if err != nil {
			return nil, err
		}
		accum = (accum << 8) | ((int(b) ^ xormask) & 0xff)
		bits += 8
		if bits >= 10 {
			digit := (accum >> (bits - 10)) & 0x3ff

			switch digit {
			case 0, 1023:
				lastDigit = lastDigit.Quo(lastDigit, hundred)
				if unscaledValue == nil {
					unscaledValue = lastDigit
				} else {
					unscaledValue = unscaledValue.Mul(unscaledValue, ten)
					unscaledValue = unscaledValue.Add(unscaledValue, lastDigit)
				}
				precision += 1
				done = true
			case 1, 1022:
				lastDigit = lastDigit.Quo(lastDigit, ten)
				if unscaledValue == nil {
					unscaledValue = lastDigit
				} else {
					unscaledValue = unscaledValue.Mul(unscaledValue, hundred)
					unscaledValue = unscaledValue.Add(unscaledValue, lastDigit)
				}
				precision += 2
				done = true
			case 2, 1021:
				if unscaledValue == nil {
					unscaledValue = lastDigit
				} else {
					unscaledValue = unscaledValue.Mul(unscaledValue, thousand)
					unscaledValue = unscaledValue.Add(unscaledValue, lastDigit)
				}
				precision += 3
				done = true
			default:
				if unscaledValue == nil {
					unscaledValue = lastDigit
					if unscaledValue != nil {
						precision += 3
					}
				} else {
					unscaledValue = unscaledValue.Mul(unscaledValue, thousand)
					unscaledValue = unscaledValue.Add(unscaledValue, lastDigit)
					precision += 3
				}
				bits -= 10
				lastDigit = big.NewInt(int64(digit - digitAdjust))
			}
		}
	}

	scale := precision - exponent
	dec := new(Decimal)
	dec.SetIntScale(unscaledValue, scale)
	return dec, nil
}

func encodeInt32BE(v int, writer BytesWriter) error {
	var s uint = 24
	for s <= 24 {
		err := writer.WriteByte(byte(v >> s))
		if err != nil {
			return err
		}
		s -= 8
	}
	return nil
}

func decodeInt32BE(reader BytesReader) (int, error) {
	var bytes [4]byte
	len, err := reader.Read(bytes[:])
	if err != nil {
		return len, nil
	}
	if len != 4 {
		return len, awserr.New(request.ErrCodeSerialization, "incomplete lexdecimal", nil)
	}
	v := int32(binary.BigEndian.Uint32(bytes[:]))
	return int(v), nil
}
