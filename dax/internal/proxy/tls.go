package proxy

// This package is copied from crypto/tls package of Go 1.15.
// This package is used to provide functionality
// of crypto/tls package of Go 1.15 to lower versions.

import (
	"context"
	"crypto/tls"
	"net"
	"strings"
	"time"
)

type Dialer struct {
	NetDialer *net.Dialer
	Config    *tls.Config
}

type timeoutError struct {
}

func (timeoutError) Error() string   { return "tls: DialWithDialer timed out" }
func (timeoutError) Timeout() bool   { return true }
func (timeoutError) Temporary() bool { return true }

func (d *Dialer) DialContext(ctx context.Context, network, addr string) (net.Conn, error) {
	c, err := dial(ctx, d.netDialer(), network, addr, d.Config)
	if err != nil {
		// Don't return c (a typed nil) in an interface.
		return nil, err
	}
	return c, nil
}

// WARNING: this can leak a goroutine for as long as the underlying Dialer implementation takes to timeout
func dial(ctx context.Context, netDialer *net.Dialer, network, addr string, config *tls.Config) (*tls.Conn, error) {
	// We want the Timeout and Deadline values from dialer to cover the
	// whole process: TCP connection and TLS handshake. This means that we
	// also need to start our own timers now.
	timeout := netDialer.Timeout

	if !netDialer.Deadline.IsZero() {
		deadlineTimeout := time.Until(netDialer.Deadline)
		if timeout == 0 || deadlineTimeout < timeout {
			timeout = deadlineTimeout
		}
	}

	// hsErrCh is non-nil if we might not wait for Handshake to complete.
	var hsErrCh chan error
	if timeout != 0 || ctx.Done() != nil {
		hsErrCh = make(chan error, 2)
	}
	if timeout != 0 {
		timer := time.AfterFunc(timeout, func() {
			hsErrCh <- timeoutError{}
		})
		defer timer.Stop()
	}

	rawConn, err := netDialer.DialContext(ctx, network, addr)
	if err != nil {
		return nil, err
	}

	colonPos := strings.LastIndex(addr, ":")
	if colonPos == -1 {
		colonPos = len(addr)
	}
	hostname := addr[:colonPos]

	// If no ServerName is set, infer the ServerName
	// from the hostname we're connecting to.
	if config.ServerName == "" {
		// Make a copy to avoid polluting argument or default.
		c := config.Clone()
		c.ServerName = hostname
		config = c
	}

	conn := tls.Client(rawConn, config)

	if hsErrCh == nil {
		err = conn.Handshake()
	} else {
		go func() {
			hsErrCh <- conn.Handshake()
		}()

		select {
		case <-ctx.Done():
			err = ctx.Err()
		case err = <-hsErrCh:
			if err != nil {
				// If the error was due to the context
				// closing, prefer the context's error, rather
				// than some random network teardown error.
				if e := ctx.Err(); e != nil {
					err = e
				}
			}
		}
	}

	if err != nil {
		rawConn.Close()
		return nil, err
	}

	return conn, nil
}

func (d *Dialer) netDialer() *net.Dialer {
	if d.NetDialer != nil {
		return d.NetDialer
	}
	return new(net.Dialer)
}
