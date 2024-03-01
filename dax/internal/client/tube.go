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

package client

import (
	"bufio"
	"net"
	"strconv"
	"time"

	"github.com/aws/aws-dax-go/dax/internal/cbor"
)

const magic = "J7yne5G"
const agent = "DaxGoClient-1.2.14"

var optional = map[string]string{"UserAgent": agent}

type session = int64

// Interface to represent a data stream connection to the Dax server
type tube interface {
	AuthExpiryUnix() int64
	SetAuthExpiryUnix(int64)
	CompareAndSwapAuthID(string) bool
	SetDeadline(time.Time) error
	Session() session
	Next() tube
	SetNext(tube)
	CborReader() *cbor.Reader
	CborWriter() *cbor.Writer

	Close() error
}

// A concrete tube implementation based on net.Conn
type netConnTube struct {
	sess       session
	conn       net.Conn
	cborReader *cbor.Reader
	cborWriter *cbor.Writer
	next       tube

	authExpiryUnix int64
	authID         string
}

// Creates and initializes a new tube belonging to the given session
// and using the provided connection.
func newTube(c net.Conn, s session) (tube, error) {
	w := cbor.NewWriter(bufio.NewWriter(c))
	closeResources := func() {
		w.Close()
		c.Close()
	}
	if err := writeMagic(w); err != nil {
		closeResources()
		return nil, err
	}
	if err := writeLayering(w); err != nil {
		closeResources()
		return nil, err
	}
	if err := writeSession(w, s); err != nil {
		closeResources()
		return nil, err
	}
	if err := writeHeader(w); err != nil {
		closeResources()
		return nil, err
	}
	if err := writeClientMode(w); err != nil {
		closeResources()
		return nil, err
	}
	if err := w.Flush(); err != nil {
		closeResources()
		return nil, err
	}

	// pack pointer inside the struct to prevent excessive copying
	return &netConnTube{
		sess:       s,
		conn:       c,
		cborReader: cbor.NewReader(bufio.NewReader(c)),
		cborWriter: w,
	}, nil

}

func (t *netConnTube) AuthExpiryUnix() int64 {
	return t.authExpiryUnix
}

func (t *netConnTube) SetAuthExpiryUnix(expiry int64) {
	t.authExpiryUnix = expiry
}

// Swaps auth id if it differs from the current one.
// Returns true if we swapped the auth id, false otherwise.
func (t *netConnTube) CompareAndSwapAuthID(id string) bool {
	if t.authID != id {
		t.authID = id
		return true
	}
	return false
}

// Sets the deadline on the underlying net.Conn object
func (t *netConnTube) SetDeadline(time time.Time) error {
	return t.conn.SetDeadline(time)
}

func (t *netConnTube) Session() session {
	return t.sess
}

func (t *netConnTube) Next() tube {
	return t.next
}

func (t *netConnTube) SetNext(n tube) {
	t.next = n
}

func (t *netConnTube) CborReader() *cbor.Reader {
	return t.cborReader
}

func (t *netConnTube) CborWriter() *cbor.Writer {
	return t.cborWriter
}

func (t *netConnTube) Close() error {
	t.cborWriter.Close()
	t.cborReader.Close()
	return t.conn.Close()
}

func writeMagic(w *cbor.Writer) error {
	return w.WriteString(magic)
}

func writeSession(w *cbor.Writer, s session) error {
	// server expects a string
	sess := strconv.FormatInt(s, 10)
	return w.WriteString(sess)
}

func writeLayering(w *cbor.Writer) error {
	return w.WriteInt(0)
}

func writeHeader(w *cbor.Writer) error {
	if err := w.WriteMapHeader(len(optional)); err != nil {
		return err
	}
	for k, v := range optional {
		if err := w.WriteString(k); err != nil {
			return err
		}
		if err := w.WriteString(v); err != nil {
			return err
		}
	}
	return nil
}

func writeClientMode(w *cbor.Writer) error {
	// client mode
	return w.WriteInt(0)
}
