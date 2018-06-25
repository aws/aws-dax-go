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

// Type encoding sizes.
const (
	Size8      = 0x18
	Size16     = 0x19
	Size32     = 0x1a
	Size64     = 0x1b
	SizeStream = 0x1f
)

// Basic types.
const (
	PosInt   = 0x00 // 0 << 5
	PosInt8  = PosInt + Size8
	PosInt16 = PosInt + Size16
	PosInt32 = PosInt + Size32
	PosInt64 = PosInt + Size64

	NegInt   = 0x20 // 1 << 5
	NegInt8  = NegInt + Size8
	NegInt16 = NegInt + Size16
	NegInt32 = NegInt + Size32
	NegInt64 = NegInt + Size64

	Bytes       = 0x40 // 2 << 5
	Bytes8      = Bytes + Size8
	Bytes16     = Bytes + Size16
	Bytes32     = Bytes + Size32
	Bytes64     = Bytes + Size64
	BytesStream = Bytes + SizeStream

	Utf       = 0x60 // 3 << 5
	Utf8      = Utf + Size8
	Utf16     = Utf + Size16
	Utf32     = Utf + Size32
	Utf64     = Utf + Size64
	UtfStream = Utf + SizeStream

	Array       = 0x80 // 4 << 5
	Array8      = Array + Size8
	Array16     = Array + Size16
	Array32     = Array + Size32
	Array64     = Array + Size64
	ArrayStream = Array + SizeStream

	Map       = 0xa0 // 5 << 5
	Map8      = Map + Size8
	Map16     = Map + Size16
	Map32     = Map + Size32
	Map64     = Map + Size64
	MapStream = Map + SizeStream

	Tag   = 0xc0 // 6 << 5
	Tag8  = Tag + Size8
	Tag16 = Tag + Size16
	Tag32 = Tag + Size32
	Tag64 = Tag + Size64

	Simple = 0xe0 // 7 << 5
)

// Simple and special types.
const (
	False = Simple + 0x14 + iota
	True
	Nil
	Undefined
	Simple8
	Float16
	Float32
	Float64
	Break = Simple + SizeStream
)

// Standard tags.
const (
	TagDatetime  = iota // string
	TagTimestamp        // seconds from epoch
	TagPosBigInt
	TagNegBigInt
	TagDecimal
	TagBigFloat
)

const (
	MajorTypeMask = 0xe0 // Upper 3 bits of type header defines the major type.
	MinorTypeMask = 0x1f // Lower 5 bits of type header defines the minor type.
)
