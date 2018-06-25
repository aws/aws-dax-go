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
)

func (r *Reader) readBigInt(flip int) (*big.Int, error) {
	bs, err := r.ReadBytes()
	if err != nil {
		return nil, err
	}

	v := new(big.Int)
	v.SetBytes(bs)
	if flip < 0 {
		v.Not(v)
	}
	return v, nil
}

func (r *Reader) ReadBigInt() (*big.Int, error) {
	hdr, value, err := r.readTypeHeader()
	if err != nil {
		return nil, err
	}
	switch hdr & MajorTypeMask {
	case PosInt:
		v := new(big.Int)
		v.SetUint64(value)
		return v, nil

	case NegInt:
		v := new(big.Int)
		v.SetUint64(value)
		v.Not(v)
		return v, nil

	case Tag:
		// TODO skip other tags.
		switch int(value) {
		case TagPosBigInt:
			return r.readBigInt(0)
		case TagNegBigInt:
			return r.readBigInt(-1)
		}

	case Simple:
		// FIXME: convertible floats.
	}
	return nil, ErrNaN
}

func (w *Writer) WriteBigInt(value *big.Int) error {
	sign := value.Sign() >> 1
	bitlen := value.BitLen()
	switch {
	case bitlen <= 63:
		return w.WriteInt64(value.Int64())
	case bitlen == 64:
		v := value.Uint64()
		if sign < 0 {
			v = v - 1
		}
		return w.writeType64(uint64(PosInt64|(sign&NegInt)), v)
	}

	var cmp big.Int
	if sign < 0 {
		cmp.Not(value)
	} else {
		cmp.Set(value)
	}
	if cmp.BitLen() <= 64 {
		return w.writeType64(uint64(PosInt64|(sign&NegInt)), cmp.Uint64())
	}

	err := w.write(byte(Tag | TagPosBigInt | (sign & 1)))
	if err != nil {
		return err
	}
	bs := cmp.Bytes()
	err = w.writeType(Bytes, uint64(len(bs)))
	if err == nil {
		_, err = w.bw.Write(bs)
	}
	return err
}
