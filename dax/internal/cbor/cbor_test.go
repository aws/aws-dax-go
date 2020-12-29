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
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"math/big"
	"reflect"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
)

type IntBoundary struct {
	name  string
	value *big.Int
	cbor  []byte
}

var (
	// CBOR BigNum; uses a tagged value rather than the CBOR negative int major type
	MinCborNegativeIntMinusOne = IntBoundary{
		"min cbor negative int - 1",
		new(big.Int).Sub(MinCborNegativeInt.value, big.NewInt(1)),
		fromHex("0xc349010000000000000000"),
	}
	MinCborNegativeInt = IntBoundary{
		"min cbor negative int",
		new(big.Int).Neg(new(big.Int).Exp(big.NewInt(2), big.NewInt(64), nil)),
		fromHex("0x3bffffffffffffffff"),
	}
	MinCborNegativeIntPlusOne = IntBoundary{
		"min cbor negative int",
		new(big.Int).Add(new(big.Int).Neg(new(big.Int).Exp(big.NewInt(2), big.NewInt(64), nil)), big.NewInt(1)),
		fromHex("0x3bfffffffffffffffe"),
	}
	MinInt64MinusOne = IntBoundary{
		"min int64 - 1",
		new(big.Int).Sub(big.NewInt(math.MinInt64), big.NewInt(1)),
		fromHex("0x3b8000000000000000"),
	}
	MinInt64 = IntBoundary{
		"min int64",
		big.NewInt(math.MinInt64),
		fromHex("0x3b7fffffffffffffff"),
	}
	MinusOne = IntBoundary{
		"minus 1",
		new(big.Int).Neg(big.NewInt(1)),
		fromHex("0x20"),
	}
	Zero = IntBoundary{
		"zero",
		big.NewInt(0),
		fromHex("0x00"),
	}
	MaxInt64 = IntBoundary{
		"max int64",
		big.NewInt(math.MaxInt64),
		fromHex("0x1b7fffffffffffffff"),
	}
	MaxInt64PlusOne = IntBoundary{
		"max int64 + 1",
		new(big.Int).Add(MaxInt64.value, big.NewInt(1)),
		fromHex("0x1b8000000000000000"),
	}
	MaxCborPositiveInt = IntBoundary{
		"max cbor positive int (also max uint64)",
		new(big.Int).SetUint64(math.MaxUint64),
		fromHex("0x1bffffffffffffffff"),
	}
	MaxUint64 = MaxCborPositiveInt
	// CBOR BigNum; uses a tagged value rather than the CBOR positive int major type
	MaxUint64PlusOne = IntBoundary{
		"max uint64 + 1 (also max cbor positive int + 1)",
		new(big.Int).Add(MaxUint64.value, big.NewInt(1)),
		fromHex("0xc249010000000000000000"),
	}
	MaxCborPositiveIntPlusOne = MaxUint64PlusOne
)

// Occasionally useful for debugging tests
func TestPrintBoundaries(t *testing.T) {
	t.Skip()
	for _, b := range []IntBoundary{
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
		t.Logf("%s %d %x", b.name, b.value, b.cbor)
	}
}

func TestCborString(t *testing.T) {
	ss := []string{
		"",
		" ",
		"\n",
		"hello",
		"helloworld",
		"1234567890abcdefgh",
		"1234567890\n abcdefgh",

		"Ж",
		"ЖЖ",
		"брэд-ЛГТМ",
		"日本語日本語日本語日",
		"☺☻☹",
		"日a本b語ç日ð本Ê語þ日¥本¼語i日©",
		"日a本b語ç日ð本Ê語þ日¥本¼語i日©日a本b語ç日ð本Ê語þ日¥本¼語i日©日a本b語ç日ð本Ê語þ日¥本¼語i日©",
		"\x80\x80\x80\x80",
	}

	for _, s := range ss {
		var buf bytes.Buffer
		w := NewWriter(&buf)
		w.WriteInt(337)
		w.WriteString(s)
		w.WriteInt(1)
		w.Flush()

		r := NewReader(&buf)
		r.ReadInt()
		if v, err := r.ReadString(); err != nil {
			t.Errorf("ReadString(%v) error = %v, want = nil", s, err)
		} else if !reflect.DeepEqual(s, v) {
			t.Errorf("ReadString(%v) got = %v", s, v)
		}
	}
}

func TestCborObjTooBig(t *testing.T) {
	for _, typ := range []int{Utf, Bytes} {
		var buf bytes.Buffer
		w := NewWriter(&buf)
		w.writeType(uint64(typ), 2<<30)
		w.Flush()

		r := NewReader(&buf)
		var err error
		switch typ {
		case Utf:
			_, err = r.ReadString()
		case Bytes:
			_, err = r.ReadBytes()
		}

		if !reflect.DeepEqual(err, ErrObjTooBig) {
			t.Errorf("expected error %v, got %v", ErrObjTooBig, err)
		}
	}
}

func TestCborType(t *testing.T) {
	for _, wt := range []int{PosInt, Utf, Bytes, Map, Array} {
		for _, rt := range []int{PosInt, Utf, Bytes, Map, Array} {
			if wt != rt {
				var buf bytes.Buffer
				w := NewWriter(&buf)
				w.writeType(uint64(wt), 0)
				w.Flush()

				r := NewReader(&buf)
				var err error
				switch rt {
				case PosInt:
					_, err = r.ReadInt()
				case Utf:
					_, err = r.ReadString()
				case Bytes:
					_, err = r.ReadBytes()
				case Map:
					_, err = r.ReadMapLength()
				case Array:
					_, err = r.ReadArrayLength()
				}

				var exp error
				if rt == PosInt || rt == NegInt {
					exp = ErrNaN
				} else {
					exp = awserr.New(request.ErrCodeSerialization, fmt.Sprintf("cbor: expected major type %d, got %d", rt, wt&MajorTypeMask), nil)
				}
				if !reflect.DeepEqual(exp, err) {
					t.Errorf("expected %v, got %v", exp, err)
				}
			}
		}
	}
}

func TestCborReadRawTypeHeader(t *testing.T) {
	for _, wt := range []int{PosInt, Utf, Bytes, Map, Array} {
		var buf bytes.Buffer
		w := NewWriter(&buf)
		switch wt {
		case PosInt:
			w.WriteInt(0)
		case Utf:
			w.WriteString("")
		case Bytes:
			w.WriteBytes([]byte{})
		case Map:
			w.WriteMapHeader(0)
		case Array:
			w.WriteArrayHeader(0)
		}
		w.Flush()
		r := NewReader(&buf)

		o := &bytes.Buffer{}
		hrd, value, err := r.readRawTypeHeader(o)
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
		nr := NewReader(o)
		nhrd, nvalue, err := nr.readTypeHeader()
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
		if hrd != nhrd || value != nvalue {
			t.Errorf("results from readRawTypeHeader and readTypeHeader are different: hdr(%d != %d), value(%d != %d", hrd, nhrd, value, nvalue)
		}
		ert := wt
		rt := hrd & MajorTypeMask
		if ert != rt {
			t.Errorf("cbor: expected major type %d, got %d", ert, rt)
		}
	}
}

func TestCborReadRawBytes(t *testing.T) {
	var buf bytes.Buffer
	w := NewWriter(&buf)
	exp := []byte{byte(1)}
	w.WriteBytes(exp)
	w.Flush()

	r := NewReader(&buf)
	var obuf bytes.Buffer
	if err := r.ReadRawBytes(&obuf); err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	nr := NewReader(&obuf)
	act, err := nr.ReadBytes()
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if !reflect.DeepEqual(exp, act) {
		t.Errorf("expected %v, got %v", exp, err)
	}
}

func fromHex(s string) []byte {
	if strings.HasPrefix(s, "0x") {
		s = s[2:]
	}
	src := []byte(s)
	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}
	return dst[:n]
}

func TestCborDecimal(t *testing.T) {
	values := []struct {
		value string
		mant  int64
		scale int
		cbor  []byte
	}{
		{value: "0", mant: 0, scale: 0},
		{value: "0.00", mant: 0, scale: 2},
		{value: "123", mant: 123, scale: 0},
		{value: "-123", mant: -123, scale: 0},
		{value: "1.23E3", mant: 123, scale: -1},
		{value: "1.23E+3", mant: 123, scale: -1},
		{value: "12.3E+7", mant: 123, scale: -6},
		{value: "12.0", mant: 120, scale: 1},
		{value: "12.3", mant: 123, scale: 1},
		{value: "0.00123", mant: 123, scale: 5},
		{value: "-1.23E-12", mant: -123, scale: 14},
		{value: "1234.5E-4", mant: 12345, scale: 5},
		{value: "0E+7", mant: 0, scale: -7},
		{value: "-0", mant: 0, scale: 0},

		{"273.15", 27315, 2, fromHex("0xc48221196ab3")},
	}

	for _, tt := range values {
		d := new(Decimal)
		if _, ok := d.SetString(tt.value); !ok {
			t.Errorf("Decimal.SetString(%v) not ok", tt.value)
			continue
		}
		if d.Scale() != tt.scale || d.Unscaled().Int64() != tt.mant {
			t.Errorf("Decimal.SetString(%v) got %v", tt.value, d)
			continue
		}

		// Should be able to round-trip using Decimal.String().
		rts := new(Decimal)
		if _, ok := rts.SetString(d.String()); !ok {
			t.Errorf("Decimal.SetString(%v) string %v not ok", tt.value, d)
		} else if !reflect.DeepEqual(d, rts) {
			t.Errorf("Decimal.SetString(%v) string %v got %v", tt.value, d, rts)
		}

		var buf bytes.Buffer
		w := NewWriter(&buf)
		if err := w.WriteDecimal(d); err != nil {
			t.Errorf("WriteDecimal(%v) got error %v", tt.value, err)
		}
		w.Flush()

		if tt.cbor != nil {
			// Test expected bytes if provided.
			bs := buf.Bytes()
			if !reflect.DeepEqual(bs, tt.cbor) {
				t.Errorf("WriteDecimal(%v)\ngot  %v\nwant %v", tt.value, hex.Dump(bs), hex.Dump(tt.cbor))
				continue
			}
		}

		r := NewReader(&buf)
		rt, err := r.ReadDecimal()
		if err != nil {
			t.Errorf("ReadDecimal(%v) got error %v", tt.value, err)
		} else if !reflect.DeepEqual(d, rt) {
			t.Errorf("ReadDecimal(%v) got %v", tt.value, rt)
		}
	}
}

func TestCborBigInt(t *testing.T) {
	values := []struct {
		value string
		cbor  []byte
	}{
		{"18446744073709551614", fromHex("0x1bfffffffffffffffe")},
		{"18446744073709551615", fromHex("0x1bffffffffffffffff")},
		{"18446744073709551616", fromHex("0xc249010000000000000000")},
		{"18446744073709551617", fromHex("0xc249010000000000000001")},

		{"-446744073709551614", fromHex("0x3b0633275e3af7fffd")},
		{"-8446744073709551614", fromHex("0x3b7538dcfb7617fffd")},
		{"-13835058055282163712", fromHex("0x3bbfffffffffffffff")},

		{"-18446744073709551614", fromHex("0x3bfffffffffffffffd")},
		{"-18446744073709551614", fromHex("0x3bfffffffffffffffd")},
		{"-18446744073709551615", fromHex("0x3bfffffffffffffffe")},
		{"-18446744073709551616", fromHex("0x3bffffffffffffffff")},
		{"-18446744073709551617", fromHex("0xc349010000000000000000")},
		{"-18446744073709551618", fromHex("0xc349010000000000000001")},

		{"9223372036854775807", fromHex("0x1b7fffffffffffffff")},
		{"9223372036854775808", fromHex("0x1b8000000000000000")},
		{"-9223372036854775808", fromHex("0x3b7fffffffffffffff")},
		{"-9223372036854775809", fromHex("0x3b8000000000000000")},

		{"0", fromHex("0x00")},
		{"1", fromHex("0x01")},
		{"-1", fromHex("0x20")},
	}

	for _, tt := range values {
		v := new(big.Int)
		v.SetString(tt.value, 10)

		var buf bytes.Buffer
		w := NewWriter(&buf)
		if err := w.WriteBigInt(v); err != nil {
			t.Errorf("WriteBigInt(%v) got error %v", v, err)
		}
		w.Flush()

		bs := buf.Bytes()
		if !reflect.DeepEqual(bs, tt.cbor) {
			t.Errorf("WriteBigInt(%v)\ngot  %v\nwant %v", v, hex.Dump(bs), hex.Dump(tt.cbor))
			continue
		}

		r := NewReader(&buf)
		rt, err := r.ReadBigInt()
		if err != nil {
			t.Errorf("ReadBigInt(%v) got error %v", v, err)
		} else if !reflect.DeepEqual(v, rt) {
			t.Errorf("ReadBigInt(%v) got %v", v, rt)
		}
	}
}

func TestCborFloat64(t *testing.T) {
	values := []struct {
		value float64
		cbor  []byte
	}{
		// Examples from rfc7409: https://tools.ietf.org/html/rfc7049#appendix-A
		// TODO: float16/32
		//{0.0, fromHex("0xf90000")},
		//{-0.0, fromHex("0xf98000")},
		//{1.0, fromHex("0xf93c00")},
		//{1.5, fromHex("0xf93e00")},
		//{65504.0, fromHex("0xf97bff")},
		//{5.960464477539063e-8, fromHex("0xf90001")},
		//{0.00006103515625, fromHex("0xf90400")},
		//{-4.0, fromHex("0xf9c400")},
		//{math.Inf(0), fromHex("0xf97c00")},
		//{math.NaN(), fromHex("0xf97e00")},
		//{math.Inf(-1), fromHex("0xf9fc00")},
		//{100000.0, fromHex("0xfa47c35000")},
		//{3.4028234663852886e+38, fromHex("0xfa7f7fffff")},
		//{math.Inf(0), fromHex("0xfa7f800000")},
		//{math.NaN(), fromHex("0xfa7fc00000")},
		//{math.Inf(-1), fromHex("0xfaff800000")},

		{0.0, fromHex("0xfb0000000000000000")},
		// Golang does not support fp constant of -0.0
		// Use math.Copysign for explicit IEEE754 negative zero.
		// See https://github.com/golang/go/issues/2196
		{math.Copysign(0, -1), fromHex("0xfb8000000000000000")},

		{1.1, fromHex("0xfb3ff199999999999a")},
		{1.0e+300, fromHex("0xfb7e37e43c8800759c")},
		{-4.1, fromHex("0xfbc010666666666666")},
		{math.Inf(0), fromHex("0xfb7ff0000000000000")},
		{math.NaN(), fromHex("0xfb7ff8000000000000")},
		{math.Inf(-1), fromHex("0xfbfff0000000000000")},
	}

	for _, tt := range values {
		v := tt.value
		cbor := tt.cbor

		var buf bytes.Buffer
		w := NewWriter(&buf)
		if err := w.WriteFloat64(v); err != nil {
			t.Errorf("WriteFloat64(%v) error = %v, want = nil", v, err)
		} else {
			if err := w.Flush(); err != nil {
				t.Errorf("Flush(%v) error = %v, want = nil", v, err)
			}

			if bs := buf.Bytes(); !reflect.DeepEqual(bs, cbor) {
				t.Errorf("WriteFloat64(%v) got %v want %v", v, hex.Dump(bs), hex.Dump(cbor))
			}

			r := NewReader(&buf)
			rt, err := r.ReadFloat64()
			if err != nil {
				t.Errorf("ReadFloat64(%v) error = %v, want = nil", v, err)
			} else if math.IsNaN(v) {
				if !math.IsNaN(rt) || !isNaN(rt) || !isNaN(v) {
					t.Errorf("ReadFloat64(%v) want NaN got %v", v, rt)
				}
			} else if rt != v || isNaN(rt) {
				t.Errorf("ReadFloat64(%v) got %v", v, rt)
			}
		}
	}
}

func TestCborMap(t *testing.T) {
	var buf bytes.Buffer
	w := NewWriter(&buf)

	m := map[int]int{
		0: 2,
		1: 4,
		2: 8,
	}
	w.WriteMapHeader(len(m))
	for k, v := range m {
		w.WriteInt(k)
		w.WriteInt(v)
	}
	w.Flush()

	r := NewReader(&buf)
	size, _ := r.ReadMapLength()
	rm := make(map[int]int, size)
	for i := 0; i < size; i++ {
		k, _ := r.ReadInt()
		v, _ := r.ReadInt()
		rm[k] = v
	}

	if !reflect.DeepEqual(m, rm) {
		t.Errorf("ReadMap(%v) got %v", m, rm)
	}
}

func TestCborInt64(t *testing.T) {
	values := []int64{0, 256, 16 << 10, 32 << 10, 64 << 10,
		math.MaxInt32, math.MinInt32, math.MaxInt64, math.MinInt64}

	tt := func(v int64) {
		var buf bytes.Buffer
		w := NewWriter(&buf)
		if err := w.WriteInt64(v); err != nil {
			t.Errorf("WriteInt(%v) error = %v, want = nil", v, err)
		} else {
			w.Flush()
			r := NewReader(&buf)
			rt, err := r.ReadInt64()
			if err != nil {
				t.Errorf("ReadInt(%v) error = %v, want = nil", v, err)
			} else if rt != v {
				t.Errorf("ReadInt(%v) got %v", v, rt)
			}
		}
	}

	for _, v := range values {
		tt(v)
		tt(v - 1)
		tt(v + 1)
	}
}

func TestReadCborIntegerToString(t *testing.T) {
	for _, tt := range []IntBoundary{
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
	} {
		r := NewReader(bytes.NewBuffer(tt.cbor))
		a, err := r.ReadCborIntegerToString()
		if err != nil {
			t.Errorf("reading %s, expected: %d, got error %v", tt.name, tt.value, err)
		} else if tt.value.String() != a {
			t.Errorf("reading %s, expected %s, actual: %s", tt.name, tt.value.String(), a)
		}
	}
}

func BenchmarkEncodeCborIntSmall(b *testing.B) {
	benchmarkEncodeInt(1, b)
}

func BenchmarkEncodeCborIntBig(b *testing.B) {
	benchmarkEncodeInt(math.MaxInt64, b)
}

func benchmarkEncodeInt(value int, b *testing.B) {
	var buf bytes.Buffer
	w := NewWriter(&buf)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if err := w.WriteInt(value); err != nil {
			b.Errorf("WriteInt(%v) error = %v, want = nil", value, err)
		}
		w.Flush()
		buf.Reset()
	}
}

func BenchmarkEncodeBigInt(b *testing.B) {
	value := new(big.Int)
	value.SetString("-18446744073709551617", 10)

	var buf bytes.Buffer
	w := NewWriter(&buf)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if err := w.WriteBigInt(value); err != nil {
			b.Errorf("WriteBigInt(%v) error = %v, want = nil", value, err)
		}
		w.Flush()
		buf.Reset()
	}
}

func BenchmarkPeekHeader(b *testing.B) {
	value := byte(1)
	buf := bytes.NewBuffer([]byte{value})
	r := NewReader(buf)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if v, err := r.PeekHeader(); err != nil {
			b.Errorf("PeekHeader(%v) error = %v, want = nil", value, err)
		} else if v != value {
			b.Errorf("PeekHeader(%v) got %v", value, v)
		}
	}
}

func BenchmarkCborReadRawHeaderValue(b *testing.B) {
	var buf bytes.Buffer
	w := NewWriter(&buf)
	if err := w.WriteInt(0); err != nil {
		b.Errorf("WriteInt(%v) error = %v, want = nil", 0, err)
	}
	w.Flush()
	bs := buf.Bytes()
	buf.Reset()
	br := bytes.NewReader(bs)
	rdr := NewReader(br)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if _, _, err := rdr.readRawTypeHeader(&buf); err != nil {
			b.Errorf("readRawTypeHeader error expect nil, got %v", err)
		}
		buf.Reset()
		br.Seek(0, 0)
	}
}

func BenchmarkCborReadHeaderValue(b *testing.B) {
	var buf bytes.Buffer
	w := NewWriter(&buf)
	if err := w.WriteInt(0); err != nil {
		b.Errorf("WriteInt(%v) error = %v, want = nil", 0, err)
	}
	w.Flush()
	bs := buf.Bytes()
	br := bytes.NewReader(bs)
	rdr := NewReader(br)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if _, _, err := rdr.readTypeHeader(); err != nil {
			b.Errorf("readRawTypeHeader error expect nil, got %v", err)
		}
		br.Seek(0, 0)
	}
}

func BenchmarkDecodeCborIntSmall(b *testing.B) {
	benchmarkDecodeInt(1, b)
}

func BenchmarkDecodeCborIntBig(b *testing.B) {
	benchmarkDecodeInt(math.MaxInt64, b)
}

func benchmarkDecodeInt(value int, b *testing.B) {
	var buf bytes.Buffer
	w := NewWriter(&buf)
	if err := w.WriteInt(value); err != nil {
		b.Errorf("WriteInt(%v) error = %v, want = nil", value, err)
	}
	w.Flush()
	bs := buf.Bytes()
	br := bytes.NewReader(bs)
	rdr := NewReader(br)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		v, err := rdr.ReadInt()
		if err != nil {
			b.Errorf("ReadInt(%v) error = %v, want = nil", value, err)
		} else if v != value {
			b.Errorf("ReadInt(%v) got %v", value, v)
		}
		br.Seek(0, 0)
	}
}
