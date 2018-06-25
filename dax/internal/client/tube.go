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
	"github.com/aws/aws-dax-go/dax/internal/cbor"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

const magic = "J7yne5G"
const agent = "DaxGoClient-1.0"

var optional = map[string]string{"UserAgent": agent}

type tube struct {
	session    string
	conn       net.Conn
	cborReader *cbor.Reader
	cborWriter *cbor.Writer
	next       *tube

	authExpiryUnix int64
	lock           sync.Mutex
	authId         string
}

func newTube(c net.Conn) (*tube, error) {
	tube := tube{
		conn:       c,
		cborReader: cbor.NewReader(bufio.NewReader(c)),
		cborWriter: cbor.NewWriter(bufio.NewWriter(c)),
		session:    "sess-1", // TODO session mgmt
	}
	return &tube, nil
}

func (t *tube) init() error {
	w := t.cborWriter
	if err := w.WriteString(magic); err != nil {
		return err
	}
	if err := w.WriteInt(0); err != nil { // layering
		return err
	}
	if err := w.WriteString(t.session); err != nil {
		return err
	}
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
	if err := w.WriteInt(0); err != nil { // client mode
		return err
	}

	return w.Flush()
}

func (t *tube) getAuthExpiryUnix() int64 {
	return atomic.LoadInt64(&t.authExpiryUnix)
}

func (t *tube) setAuthExpiryUnix(expiry int64) {
	atomic.StoreInt64(&t.authExpiryUnix, expiry)
}

func (t *tube) compareAndSwapAuthId(id string) bool {
	t.lock.Lock()
	defer t.lock.Unlock()
	if t.authId != id {
		t.authId = id
		return true
	}
	return false
}

func (t *tube) setDeadline(time time.Time) error {
	return t.conn.SetDeadline(time)
}

func (t *tube) Close() error {
	t.cborWriter.Close()
	t.cborReader.Close()
	return t.conn.Close()
}
