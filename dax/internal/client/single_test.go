package client

import (
	"context"
	"errors"
	"fmt"
	"net"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-dax-go/dax/internal/cbor"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var unEncryptedConnConfig = connConfig{isEncrypted: false}

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
			newDaxRequestFailure([]int{4, 0, 0}, "", "", "", 400),
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

	for i, c := range cases {
		cli, err := newSingleClientWithOptions(":9121", unEncryptedConnConfig, "us-west-2", credentials.NewStaticCredentials("id", "secret", "tok"), tubePoolOptions{}, func(ctx context.Context, a, n string) (net.Conn, error) {
			return c.conn, nil
		}, nil)
		if err != nil {
			t.Fatalf("unexpected error %v", err)
		}
		cli.pool.closeTubeImmediately = true

		err = cli.executeWithContext(aws.BackgroundContext(), OpGetItem, c.enc, c.dec, RequestOptions{})
		if !reflect.DeepEqual(c.ee, err) {
			t.Errorf("case[%d] expected error %v, got error %v", i, c.ee, err)
		}
		if !reflect.DeepEqual(c.ec, c.conn.cc) {
			t.Errorf("case[%d] expected %v calls, got %v", i, c.ec, c.conn.cc)
		}
		cli.Close()
	}
}

func TestRetryPropogatesContextError(t *testing.T) {
	client, clientErr := newSingleClientWithOptions(":9121", unEncryptedConnConfig, "us-west-2", credentials.NewStaticCredentials("id", "secret", "tok"), tubePoolOptions{}, func(ctx context.Context, a, n string) (net.Conn, error) {
		return &mockConn{rd: []byte{cbor.Array + 0}}, nil
	}, nil)
	defer client.Close()
	if clientErr != nil {
		t.Fatalf("unexpected error %v", clientErr)
	}

	client.pool.closeTubeImmediately = true

	ctx, cancel := context.WithCancel(aws.BackgroundContext())
	requestOptions := RequestOptions{
		MaxRetries: 2,
		Context:    ctx,
	}

	writer := func(writer *cbor.Writer) error { return nil }
	reader := func(reader *cbor.Reader) error { return nil }

	// Cancel context to fail the execution
	cancel()
	err := client.executeWithRetries(OpGetItem, requestOptions, writer, reader)

	// Context related error should be returned
	awsError, ok := err.(awserr.Error)
	if !ok {
		t.Fatal("Error type is not awserr.Error")
	}

	if awsError.Code() != request.CanceledErrorCode || awsError.OrigErr() != context.Canceled {
		t.Errorf("aws error doesn't match expected. %v", awsError)
	}
}

func TestRetryPropogatesOtherErrors(t *testing.T) {
	client, clientErr := newSingleClientWithOptions(":9121", unEncryptedConnConfig, "us-west-2", credentials.NewStaticCredentials("id", "secret", "tok"), tubePoolOptions{}, func(ctx context.Context, a, n string) (net.Conn, error) {
		return &mockConn{rd: []byte{cbor.Array + 0}}, nil
	}, nil)
	defer client.Close()
	if clientErr != nil {
		t.Fatalf("unexpected error %v", clientErr)
	}

	client.pool.closeTubeImmediately = true

	requestOptions := RequestOptions{
		MaxRetries: 1,
	}
	expectedError := errors.New("IO")

	writer := func(writer *cbor.Writer) error { return nil }
	reader := func(reader *cbor.Reader) error { return errors.New("IO") }

	err := client.executeWithRetries(OpGetItem, requestOptions, writer, reader)

	// IO error should be returned
	awsError, ok := err.(awserr.Error)
	if !ok {
		t.Fatal("Error type is not awserr.Error")
	}

	if awsError.OrigErr() == nil {
		t.Fatal("Original error is empty")
	}

	if awsError.Code() != "UnknownError" || awsError.OrigErr().Error() != expectedError.Error() {
		t.Errorf("aws error doesn't match expected. %v", awsError)
	}
}

func TestRetryPropogatesOtherErrorsWithDelay(t *testing.T) {
	client, clientErr := newSingleClientWithOptions(":9121", unEncryptedConnConfig, "us-west-2", credentials.NewStaticCredentials("id", "secret", "tok"), tubePoolOptions{}, func(ctx context.Context, a, n string) (net.Conn, error) {
		return &mockConn{rd: []byte{cbor.Array + 0}}, nil
	}, nil)
	defer client.Close()
	if clientErr != nil {
		t.Fatalf("unexpected error %v", clientErr)
	}

	client.pool.closeTubeImmediately = true

	requestOptions := RequestOptions{
		MaxRetries: 1,
		RetryDelay: 1,
	}
	expectedError := errors.New("IO")

	writer := func(writer *cbor.Writer) error { return nil }
	reader := func(reader *cbor.Reader) error { return expectedError }

	err := client.executeWithRetries(OpGetItem, requestOptions, writer, reader)

	// IO error should be returned
	awsError, ok := err.(awserr.Error)
	if !ok {
		t.Fatal("Error type is not awserr.Error")
	}

	if awsError.OrigErr() == nil {
		t.Fatal("Original error is empty")
	}

	if awsError.Code() != "UnknownError" || awsError.OrigErr().Error() != expectedError.Error() {
		t.Errorf("aws error doesn't match expected. %v", awsError)
	}
}

func TestRetrySleepCycleCount(t *testing.T) {
	client, clientErr := newSingleClientWithOptions(":9121", unEncryptedConnConfig, "us-west-2", credentials.NewStaticCredentials("id", "secret", "tok"), tubePoolOptions{}, func(ctx context.Context, a, n string) (net.Conn, error) {
		return &mockConn{rd: []byte{cbor.Array + 0}}, nil
	}, nil)
	defer client.Close()
	if clientErr != nil {
		t.Fatalf("unexpected error %v", clientErr)
	}

	client.pool.closeTubeImmediately = true

	sleepCallCount := 0
	requestOptions := RequestOptions{
		MaxRetries:   0,
		RetryDelay:   0,
		SleepDelayFn: func(d time.Duration) { sleepCallCount++ },
	}

	writer := func(writer *cbor.Writer) error { return nil }
	reader := func(reader *cbor.Reader) error { return errors.New("IO") }
	client.executeWithRetries(OpGetItem, requestOptions, writer, reader)

	if sleepCallCount != 0 {
		t.Fatalf("Sleep was called %d times, but expected none", sleepCallCount)
	}

	requestOptions.MaxRetries = 3
	requestOptions.RetryDelay = 1
	client.executeWithRetries(OpGetItem, requestOptions, writer, reader)

	if sleepCallCount != requestOptions.MaxRetries {
		t.Fatalf("Sleep was called %d times, but expected %d", sleepCallCount, requestOptions.MaxRetries)
	}
}

func TestRetryLastError(t *testing.T) {
	client, clientErr := newSingleClientWithOptions(":9121", unEncryptedConnConfig, "us-west-2", credentials.NewStaticCredentials("id", "secret", "tok"), tubePoolOptions{}, func(ctx context.Context, a, n string) (net.Conn, error) {
		return &mockConn{rd: []byte{cbor.Array + 0}}, nil
	}, nil)
	defer client.Close()
	if clientErr != nil {
		t.Fatalf("unexpected error %v", clientErr)
	}

	client.pool.closeTubeImmediately = true

	var sleepCallCount uint
	requestOptions := RequestOptions{
		MaxRetries:   2,
		RetryDelay:   1,
		SleepDelayFn: func(d time.Duration) { sleepCallCount++ },
	}

	writer := func(writer *cbor.Writer) error { return nil }
	reader := func(reader *cbor.Reader) error {
		if sleepCallCount == 1 {
			return errors.New("IO")
		} else {
			return errors.New("LastError")
		}
	}
	err := client.executeWithRetries(OpGetItem, requestOptions, writer, reader)
	awsError, ok := err.(awserr.Error)
	if !ok {
		t.Fatal("Error type is not awserr.Error")
	}

	if awsError.OrigErr() == nil {
		t.Fatal("Original error is empty")
	}

	if awsError.Code() != "UnknownError" || awsError.OrigErr().Error() != "LastError" {
		t.Fatalf("aws error doesn't match expected. %v", awsError)
	}
}

func TestSingleClient_customDialer(t *testing.T) {
	conn := &mockConn{}
	var dialContextFn dialContext = func(ctx context.Context, address string, network string) (net.Conn, error) {
		return conn, nil
	}
	client, err := newSingleClientWithOptions(":9121", unEncryptedConnConfig, "us-west-2", credentials.NewStaticCredentials("id", "secret", "tok"), tubePoolOptions{}, dialContextFn, nil)
	require.NoError(t, err)
	defer client.Close()

	c, _ := client.pool.dialContext(context.TODO(), "address", "network")
	assert.Equal(t, conn, c)
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

func (m *mockConn) LocalAddr() net.Addr {
	return nil
}

func (m *mockConn) RemoteAddr() net.Addr {
	return nil
}

func (m *mockConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (m *mockConn) SetWriteDeadline(t time.Time) error {
	return nil
}
