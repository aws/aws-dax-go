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
	"math/big"
	"strconv"
	"strings"
)

// Decimal represents an arbitrary-precision signed decimal number. It consists
// of an arbitrary precision integer unscaled value and a scale. If zero or
// positive, the scale is the number of digits to the right of the decimal
// point. If negative, the unscaled value of the number is multiplied by ten to
// the power of the negation of the scale: unscaled * (10 ** -scale).
type Decimal struct {
	value big.Int
	scale int
}

func NewDecimal(unscaled *big.Int, scale int) *Decimal {
	return new(Decimal).SetIntScale(unscaled, scale)
}

func (d *Decimal) SetString(s string) (*Decimal, bool) {
	if len(s) == 0 {
		return nil, false
	}

	var exp int64
	var err error
	if sep := strings.IndexAny(s, "Ee"); sep >= 0 {
		exp, err = strconv.ParseInt(s[sep+1:], 10, 32)
		if err != nil {
			return nil, false
		}
		s = s[:sep]
	}

	var v big.Int
	dot := strings.LastIndexByte(s, '.')
	if dot < 0 {
		_, ok := v.SetString(s, 10)
		if !ok {
			return nil, false
		}
		d.value.Set(&v)
		d.scale = int(-exp)
		return d, true
	}

	if signCh := strings.IndexAny(s, "+-"); signCh > 0 {
		// Optional sign must be the first char.
		return nil, false
	}
	exp -= int64(len(s) - dot - 1)

	_, ok := v.SetString(s[:dot]+s[dot+1:], 10)
	if !ok {
		return nil, false
	}

	d.value.Set(&v)
	d.scale = int(-exp)
	return d, true
}

func (d *Decimal) String() string {
	if d.scale == 0 {
		return d.value.String()
	}
	var buf []byte
	buf = d.value.Append(buf, 10)
	buf = append(buf, 'E')
	buf = strconv.AppendInt(buf, -int64(d.scale), 10)
	return string(buf)
}

func (d *Decimal) Scale() int {
	return d.scale
}

func (d *Decimal) Unscaled() *big.Int {
	return &d.value
}

func (d *Decimal) SetInt(v *big.Int) *Decimal {
	d.value.Set(v)
	return d
}

func (d *Decimal) SetScale(scale int) *Decimal {
	d.scale = scale
	return d
}

func (d *Decimal) SetIntScale(unscaled *big.Int, scale int) *Decimal {
	d.value.Set(unscaled)
	d.scale = scale
	return d
}

func (r *Reader) ReadDecimal() (*Decimal, error) {
	hdr, value, err := r.readTypeHeader()
	if err != nil {
		return nil, err
	}
	switch hdr & MajorTypeMask {
	case PosInt:
		v := new(big.Int).SetUint64(value)
		return NewDecimal(v, 0), nil

	case NegInt:
		v := new(big.Int)
		v.SetUint64(value)
		v.Not(v)
		return NewDecimal(v, 0), nil

	case Tag:
		// TODO skip other tags.
		switch int(value) {
		case TagPosBigInt:
			v, err := r.readBigInt(0)
			if err == nil {
				return NewDecimal(v, 0), nil
			}
			return nil, err

		case TagNegBigInt:
			v, err := r.readBigInt(-1)
			if err == nil {
				return NewDecimal(v, 0), nil
			}
			return nil, err

		case TagDecimal:
			size, err := r.ReadArrayLength()
			if err == nil && size == 2 {
				var scale int
				var v *big.Int
				scale, err = r.ReadInt()
				if err == nil {
					v, err = r.ReadBigInt()
				}
				if err == nil {
					return NewDecimal(v, -scale), nil
				}
			}
			if err != nil {
				return nil, err
			}
		}

	case Simple:
		// FIXME: convertible floats.
	}
	return nil, ErrNaN
}

func (w *Writer) WriteDecimal(value *Decimal) (err error) {
	err = w.write(byte(Tag | TagDecimal))
	if err != nil {
		return err
	}

	err = w.WriteArrayHeader(2)
	if err == nil {
		err = w.WriteInt(-value.Scale())
	}
	if err == nil {
		err = w.WriteBigInt(value.Unscaled())
	}
	return
}
