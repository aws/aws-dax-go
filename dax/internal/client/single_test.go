package client

import (
	"errors"
	"fmt"
	"github.com/aws/aws-dax-go/dax/internal/cbor"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"net"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestExecuteErrorHandling(t *testing.T) {
	cases := []struct {
		conn *mockConn
		enc  func(writer *cbor.Writer) error
		dec  func(reader *cbor.Reader) error
		ee   error
		ec   map[string]int
	}{
		{ // write error, discard tube
			&mockConn{we: errors.New("io")},
			nil,
			nil,
			errors.New("io"),
			map[string]int{"Write": 1, "Close": 1},
		},
		{ // encoding error, discard tube
			&mockConn{},
			func(writer *cbor.Writer) error { return errors.New("ser") },
			nil,
			errors.New("ser"),
			map[string]int{"Write": 2, "SetDeadline": 1, "Close": 1},
		},
		{ // read error, discard tube
			&mockConn{re: errors.New("IO")},
			func(writer *cbor.Writer) error { return nil },
			nil,
			errors.New("IO"),
			map[string]int{"Write": 2, "Read": 1, "SetDeadline": 1, "Close": 1},
		},
		{ // serialization error, discard tube
			&mockConn{rd: []byte{cbor.NegInt}},
			func(writer *cbor.Writer) error { return nil },
			nil,
			awserr.New(request.ErrCodeSerialization, fmt.Sprintf("cbor: expected major type %d, got %d", cbor.Array, cbor.NegInt), nil),
			map[string]int{"Write": 2, "Read": 1, "SetDeadline": 1, "Close": 1},
		},
		{ // decode error, discard tube
			&mockConn{rd: []byte{cbor.Array + 0}},
			func(writer *cbor.Writer) error { return nil },
			func(reader *cbor.Reader) error { return errors.New("IO") },
			errors.New("IO"),
			map[string]int{"Write": 2, "Read": 1, "SetDeadline": 1, "Close": 1},
		},
		{ // dax error, do not discard tube
			&mockConn{rd: []byte{cbor.Array + 3, cbor.PosInt + 4, cbor.PosInt + 0, cbor.PosInt + 0, cbor.Utf, cbor.Nil}},
			func(writer *cbor.Writer) error { return nil },
			nil,
			newDaxRequestFailure([]int{4, 0, 0}, ErrCodeUnknown, "", "", 400),
			map[string]int{"Write": 2, "Read": 1, "SetDeadline": 1},
		},
		{ // no error, do not discard tube
			&mockConn{rd: []byte{cbor.Array + 0}},
			func(writer *cbor.Writer) error { return nil },
			func(reader *cbor.Reader) error { return nil },
			nil,
			map[string]int{"Write": 2, "Read": 1, "SetDeadline": 1},
		},
	}

	for _, c := range cases {
		cli, err := newSingleClientWithOptions(":9121", "us-west-2", credentials.NewStaticCredentials("id", "secret", "tok"), 1)
		if err != nil {
			t.Fatalf("unexpected error %v", err)
		}
		cli.pool.connectFn = func(a, n string) (net.Conn, error) {
			return c.conn, nil
		}
		cli.pool.closeTubeImmediately = true

		err = cli.executeWithContext(aws.BackgroundContext(), OpGetItem, c.enc, c.dec)
		if !reflect.DeepEqual(c.ee, err) {
			t.Errorf("expected error %v, got error %v", c.ee, err)
		}
		if !reflect.DeepEqual(c.ec, c.conn.cc) {
			t.Errorf("expected %v calls, got %v", c.ec, c.conn.cc)
		}
		cli.Close()
	}
}

type mockConn struct {
	net.Conn
	we, re error
	wd, rd []byte
	cc     map[string]int
}

func (m *mockConn) Read(b []byte) (n int, err error) {
	m.register()
	if m.re != nil {
		return 0, m.re
	}
	if len(m.rd) > 0 {
		l := copy(b, m.rd)
		m.rd = m.rd[l:]
		return l, nil
	}
	return 0, nil
}

func (m *mockConn) Write(b []byte) (n int, err error) {
	m.register()
	if m.we != nil {
		return 0, m.we
	}
	if len(m.wd) > 0 {
		l := copy(m.wd, b)
		m.wd = m.wd[l:]
		return l, nil
	}
	return len(b), nil
}

func (m *mockConn) Close() error {
	m.register()
	return nil
}

func (m *mockConn) SetDeadline(t time.Time) error {
	m.register()
	return nil
}

func (m *mockConn) register() {
	pc, _, _, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	s := strings.Split(fn.Name(), ".")
	n := s[len(s)-1]
	if m.cc == nil {
		m.cc = make(map[string]int)
	}
	m.cc[n]++
}
