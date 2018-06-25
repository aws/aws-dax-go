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
	"bufio"
	"encoding/binary"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"io"
	"math"
	"sync"
)

const (
	defaultBufSize = 8192
	maxObjLenBytes = 1024 * 1024 * 1024
)

var ErrNaN = awserr.New(request.InvalidParameterErrCode, "cbor: not a number", nil)
var ErrObjTooBig = awserr.New(request.ErrCodeSerialization, "cbor: object too big", nil)
var ErrNegLength = awserr.New(request.ErrCodeSerialization, "cbor: negative length", nil)

// A Writer writes cbor-encoded data.
type Writer struct {
	w       io.Writer
	bw      *bufio.Writer
	buf     []byte
	scratch [9]byte
	recycle bool
}

var bufferedWriterPool = sync.Pool{
	New: func() interface{} {
		return bufio.NewWriterSize(nil, defaultBufSize)
	},
}

func NewWriter(w io.Writer) *Writer {
	// Check if writer is already buffered.
	bw, ok := w.(*bufio.Writer)
	if !ok {
		bw = bufferedWriterPool.Get().(*bufio.Writer)
		bw.Reset(w)
	}

	cw := Writer{
		w:       w,
		bw:      bw,
		recycle: !ok,
	}
	cw.buf = cw.scratch[:]
	return &cw
}

func (w *Writer) Flush() error {
	return w.bw.Flush()
}

func (w *Writer) WriteFloat(v float32) error {
	bits := math.Float32bits(v)
	// Abuse append.
	_ = append(w.buf[:0],
		byte(Float32),
		byte(bits>>24),
		byte(bits>>16),
		byte(bits>>8),
		byte(bits))

	_, err := w.bw.Write(w.buf[:5])
	return err
}

func (w *Writer) WriteFloat64(v float64) error {
	bits := math.Float64bits(v)
	if isNaN(v) {
		bits &= 0xfffffffffffffff0
	}
	return w.writeType64(Float64, bits)
}

func (w *Writer) WriteBoolean(b bool) error {
	v := False
	if b {
		v = True
	}
	return w.bw.WriteByte(byte(v))
}

func (w *Writer) WriteBytes(b []byte) error {
	if len(b) == 0 {
		b = w.scratch[0:0]
	}
	err := w.writeType(Bytes, uint64(len(b)))
	if err == nil {
		_, err = w.bw.Write(b)
	}
	return err
}

func (w *Writer) WriteString(s string) error {
	err := w.writeType(Utf, uint64(len(s)))
	if err == nil {
		_, err = w.bw.WriteString(s)
	}
	return err
}

func (w *Writer) WriteTag(tag uint64) error {
	return w.writeType(Tag, tag)
}

func (w *Writer) WriteMapHeader(pairs int) error {
	return w.writeType(Map, uint64(pairs))
}

func (w *Writer) WriteArrayHeader(elems int) error {
	return w.writeType(Array, uint64(elems))
}

func (w *Writer) WriteMapStreamHeader() error {
	return w.write(byte(MapStream))
}

func (w *Writer) WriteArrayStreamHeader() error {
	return w.write(byte(ArrayStream))
}

func (w *Writer) WriteStreamBreak() error {
	return w.write(byte(Break))
}

func (w *Writer) WriteNull() error {
	return w.write(byte(Nil))
}

func (w *Writer) write(b byte) error {
	err := w.bw.WriteByte(b)
	return err
}

func (w *Writer) Write(b []byte) error {
	_, err := w.bw.Write(b)
	return err
}

func (w *Writer) WriteInt(value int) error {
	return w.WriteInt64(int64(value))
}

func (w *Writer) WriteInt64(value int64) error {
	ui := uint64(value >> 63)
	return w.writeType(ui&0x20, ui^uint64(value))
}

func (w *Writer) writeType(mt uint64, ui uint64) (err error) {
	switch {
	case ui < Size8:
		err = w.bw.WriteByte(byte(mt + ui))
	case ui < 1<<8:
		err = w.bw.WriteByte(byte(mt + Size8))
		if err == nil {
			err = w.bw.WriteByte(byte(ui))
		}
	// Abuse append.
	case ui < 1<<16:
		_ = append(w.buf[:0],
			byte(mt+Size16),
			byte(ui>>8),
			byte(ui))
		_, err = w.bw.Write(w.buf[:3])
	case ui < 1<<32:
		_ = append(w.buf[:0],
			byte(mt+Size32),
			byte(ui>>24),
			byte(ui>>16),
			byte(ui>>8),
			byte(ui))
		_, err = w.bw.Write(w.buf[:5])
	default:
		_ = append(w.buf[:0],
			byte(mt+Size64),
			byte(ui>>56),
			byte(ui>>48),
			byte(ui>>40),
			byte(ui>>32),
			byte(ui>>24),
			byte(ui>>16),
			byte(ui>>8),
			byte(ui))
		_, err = w.bw.Write(w.buf)
	}
	return
}

func (w *Writer) writeType64(typ uint64, ui uint64) error {
	// Abuse append.
	_ = append(w.buf[:0],
		byte(typ),
		byte(ui>>56),
		byte(ui>>48),
		byte(ui>>40),
		byte(ui>>32),
		byte(ui>>24),
		byte(ui>>16),
		byte(ui>>8),
		byte(ui))

	_, err := w.bw.Write(w.buf)
	return err
}

func (w *Writer) Close() error {
	if w.recycle {
		bufferedWriterPool.Put(w.bw)
	}
	return nil
}

var bufferedReaderPool = sync.Pool{
	New: func() interface{} {
		return bufio.NewReaderSize(nil, defaultBufSize)
	},
}

type Reader struct {
	r  io.Reader
	br *bufio.Reader

	buf     []byte
	scratch [8]byte
	recycle bool
}

func NewReader(r io.Reader) *Reader {
	br, ok := r.(*bufio.Reader)
	if !ok {
		br = bufferedReaderPool.Get().(*bufio.Reader)
		br.Reset(r)
	}
	rdr := Reader{
		r:       r,
		br:      br,
		recycle: !ok,
	}
	rdr.buf = rdr.scratch[:]
	return &rdr
}

func (r *Reader) ReadString() (string, error) {
	// TODO skip tags, indef length strings
	hdr, value, err := r.readTypeHeader()
	if err != nil {
		return "", err
	}
	if err = r.verifyMajorType(hdr, Utf); err != nil {
		return "", err
	}
	if value > maxObjLenBytes {
		return "", ErrObjTooBig
	} else if value < 0 {
		return "", ErrNegLength
	} else if value == 0 {
		return "", nil
	}
	b := make([]byte, value)
	_, err = io.ReadFull(r.br, b)
	if err != nil {
		return "", err
	}
	return string(b), err
}

func (r *Reader) ReadBytes() ([]byte, error) {
	// TODO skip tags, indef length bytes
	hdr, value, err := r.readTypeHeader()
	if err != nil {
		return nil, err
	}
	if err = r.verifyMajorType(hdr, Bytes); err != nil {
		return nil, err
	}
	if value > maxObjLenBytes {
		return nil, ErrObjTooBig
	} else if value < 0 {
		return nil, ErrNegLength
	} else if value == 0 {
		return []byte{}, nil
	}
	b := make([]byte, value)
	_, err = io.ReadFull(r.br, b)
	if err != nil {
		return nil, err
	}
	return b, err
}

func (r *Reader) BytesReader() (*Reader, error) {
	// TODO skip tags.
	hdr, value, err := r.readTypeHeader()
	if err != nil {
		return nil, err
	}
	if err = r.verifyMajorType(hdr, Bytes); err != nil {
		return nil, err
	}
	// TODO avoid double buffering
	lr := io.LimitReader(r.br, int64(value))
	return NewReader(lr), nil
}

func (r *Reader) ReadMapLength() (int, error) {
	hdr, value, err := r.readTypeHeader()
	if err != nil {
		return 0, err
	}
	if err = r.verifyMajorType(hdr, Map); err != nil {
		return 0, err
	}
	return int(value), err
}

func (r *Reader) ReadArrayLength() (int, error) {
	hdr, value, err := r.readTypeHeader()
	if err != nil {
		return 0, err
	}
	if err = r.verifyMajorType(hdr, Array); err != nil {
		return 0, err
	}
	return int(value), err
}

func (r *Reader) ReadFloat64() (float64, error) {
	hdr, value, err := r.readTypeHeader()
	if err != nil {
		return 0, err
	}
	switch hdr & MajorTypeMask {
	case PosInt:
		//return math.Float64frombits(value), nil
		return float64(value), nil
	case NegInt:
		return float64(^value), nil
	case Simple:
	default:
		return 0, ErrNaN
	}
	switch hdr & MinorTypeMask {
	case Float16 & MinorTypeMask:
		// TODO b16 to b32
		panic("unimpl")
	case Float32 & MinorTypeMask:
		return float64(math.Float32frombits(uint32(value))), nil
	case Float64 & MinorTypeMask:
		v := math.Float64frombits(value)
		if isNaN(v) {
			v = math.NaN()
		}
		return v, nil
	}
	return 0, ErrNaN
}

func (r *Reader) PeekHeader() (hdr byte, err error) {
	b, err := r.br.Peek(1)
	if err != nil {
		return 0, err
	}
	return b[0], nil
}

func (r *Reader) ReadNil() (err error) {
	_, _, err = r.readTypeHeader()
	return err
}

func (r *Reader) ReadBreak() (err error) {
	_, _, err = r.readTypeHeader()
	return err
}

func (r *Reader) readTypeHeader() (hdr int, value uint64, err error) {
	b, err := r.br.ReadByte()
	if err != nil {
		return 0, 0, err
	}
	hdr = int(b)
	switch hdr & MinorTypeMask {
	default:
		value = uint64(hdr) & MinorTypeMask
	case Size8:
		var nb byte
		nb, err = r.br.ReadByte()
		if err == nil {
			value = uint64(nb)
		}
	case Size16:
		if _, err = io.ReadFull(r.br, r.buf[:2]); err == nil {
			value = uint64(binary.BigEndian.Uint16(r.buf))
		}
	case Size32:
		if _, err = io.ReadFull(r.br, r.buf[:4]); err == nil {
			value = uint64(binary.BigEndian.Uint32(r.buf))
		}
	case Size64:
		if _, err = io.ReadFull(r.br, r.buf); err == nil {
			value = binary.BigEndian.Uint64(r.buf)
		}
	}
	return
}

func (r *Reader) verifyMajorType(hdr, exp int) error {
	if (hdr & MajorTypeMask) != exp {
		return awserr.New(request.ErrCodeSerialization, fmt.Sprintf("cbor: expected major type %d, got %d", exp, hdr&MajorTypeMask), nil)
	}
	return nil
}

func (r *Reader) ReadInt() (int, error) {
	v, err := r.ReadInt64()
	return int(v), err
}

func (r *Reader) ReadInt64() (int64, error) {
	hdr, value, err := r.readTypeHeader()
	if err != nil {
		return 0, err
	}
	//TODO skip tags.
	switch hdr & MajorTypeMask {
	case NegInt:
		return ^int64(value), nil
	case PosInt:
		return int64(value), nil
	default:
		return 0, ErrNaN
	}
}

func (r *Reader) Close() error {
	if r.recycle {
		bufferedReaderPool.Put(r.br)
	}
	return nil
}
