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
	"context"
	"crypto/tls"
	"net"
	"os"
	"sync"
	"time"

	"github.com/aws/aws-dax-go/dax/internal/proxy"
)

const network = "tcp"

type dialContext func(ctx context.Context, network string, address string) (net.Conn, error)

// Acts as the gate to create new tubes
// and keeps track of tubes which are currently not in use.
type tubePool struct {
	address              string
	gate                 gate
	errCh                chan error
	timeout              time.Duration
	dialContext          dialContext
	closeTubeImmediately bool

	mutex      sync.Mutex
	closed     bool    // protected by mutex
	top        tube    // protected by mutex
	lastActive tube    // protected by mutex
	session    session // protected by mutex
	waiters    chan tube

	connConfig connConfig
}

type tubePoolOptions struct {
	maxConcurrentConnAttempts int
	timeout                   time.Duration
	dialContext               dialContext
}

var defaultDialer = &net.Dialer{}

var defaultTubePoolOptions = tubePoolOptions{maxConcurrentConnAttempts: 10, timeout: time.Second * 5}

// Creates a new pool using defaultTubePoolOptions and associated with given address.
func newTubePool(address string, connConfigData connConfig) *tubePool {
	return newTubePoolWithOptions(address, defaultTubePoolOptions, connConfigData)
}

// Creates a new pool with provided options associated with the given address.
func newTubePoolWithOptions(address string, options tubePoolOptions, connConfigData connConfig) *tubePool {
	if options.maxConcurrentConnAttempts <= 0 {
		options.maxConcurrentConnAttempts = defaultTubePoolOptions.maxConcurrentConnAttempts
	}

	if options.dialContext == nil {
		if connConfigData.isEncrypted {
			dialer := &proxy.Dialer{}
			var cfg tls.Config
			if connConfigData.skipHostnameVerification {
				cfg = tls.Config{InsecureSkipVerify: true}
			} else {
				cfg = tls.Config{ServerName: connConfigData.hostname}
			}
			dialer.Config = &cfg
			options.dialContext = dialer.DialContext
		} else {
			dialer := &net.Dialer{}
			options.dialContext = dialer.DialContext
		}
	}

	return &tubePool{
		address:     address,
		gate:        make(gate, options.maxConcurrentConnAttempts),
		errCh:       make(chan error),
		waiters:     make(chan tube),
		timeout:     options.timeout,
		dialContext: options.dialContext,

		connConfig: connConfigData,
	}
}

// Gets a new or reuses existing tube with timeout context set to tubePool#timeout
func (p *tubePool) get() (tube, error) {
	ctx := context.Background()
	if p.timeout > 0 {
		var cancelFn func()
		ctx, cancelFn = context.WithTimeout(ctx, p.timeout)
		defer cancelFn()
	}
	return p.getWithContext(ctx, false)
}

// Gets a new or reuses existing tube with provided context.
// Create a new tube even if pool reached maxConcurrentConnAttempts if highPriority is true.
func (p *tubePool) getWithContext(ctx context.Context, highPriority bool) (tube, error) {
	for {
		p.mutex.Lock()
		if p.closed {
			p.mutex.Unlock()
			return nil, os.ErrClosed
		}

		// look for idle tubes in stack
		if p.top != nil {
			t := p.top
			p.top = t.Next()
			if p.lastActive == t {
				p.lastActive = p.top
			}
			t.SetNext(nil)
			p.mutex.Unlock()
			return t, nil
		}

		// no tubes in stack, create wait channel
		if p.waiters == nil {
			p.waiters = make(chan tube)
		}
		waitCh := p.waiters
		session := p.session
		p.mutex.Unlock()

		var done chan tube
		if p.gate.tryEnter() {
			go p.allocAndReleaseGate(session, done, true)
		} else if highPriority {
			done = make(chan tube)
			go p.allocAndReleaseGate(session, done, false)
		}

		select {
		case tube := <-waitCh:
			if tube != nil {
				return tube, nil
			}
			// if channel is closed, continue to look for idle tubes in stack
		case tube := <-done:
			if tube != nil {
				return tube, nil
			}
		case err := <-p.errCh:
			// if channel was closed, the error will be nil
			if err != nil {
				return nil, err
			}
			return nil, os.ErrClosed
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}

// Allocates a new tube and optionally releases the gate.
// If done channel isn't nil the new tube will be send there as opposed to idle tubes stack.
func (p *tubePool) allocAndReleaseGate(session int64, done chan tube, releaseGate bool) {
	tube, err := p.alloc(session)
	if releaseGate {
		p.gate.exit()
	}
	if err == nil {
		select {
		case done <- tube:
		default:
			p.put(tube)
		}
	} else {
		p.mutex.Lock()
		cls := p.closed
		p.mutex.Unlock()
		if !cls {
			select {
			case p.errCh <- err:
			default:
			}
		}
	}
	if done != nil {
		close(done)
	}
}

// Returns a previously allocated tube back into the pool.
// Tube will be closed if the pool is closed or its coming from a different session
// Otherwise it will be handed over to a waiter, if any
// or it will be added on top of the idle tubes stack.
func (p *tubePool) put(t tube) {
	if t == nil {
		return
	}

	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.closed || t.Session() != p.session {
		t.Close()
		// Waiters channel was already closed in Close
		return
	}

	if p.waiters != nil {
		select {
		case p.waiters <- t:
			return
		default:
			close(p.waiters) // unblock future waiters
			p.waiters = nil
		}
	}

	t.SetNext(p.top)
	p.top = t
}

// Closes the specified tube, and if the tube is using the same version as the current session,
// then also closes all other idle tubes and performs a version bump.
func (p *tubePool) discard(t tube) {
	if t == nil {
		return
	}
	if p.closeTubeImmediately {
		t.Close()
	} else {
		go func() {
			t.Close()
		}()
	}

	p.mutex.Lock()

	var head tube
	if t.Session() == p.session {
		p.sessionBump()
		head = p.clearIdleConnections()
	}

	// Waiters enter the waiting queue when there's no existing tube
	// or when they failed to acquire a permit to create a new tube.
	// There's also a chance the newly created tube was stolen and
	// the thief must return it back into the pool or discard it.
	if p.waiters != nil {
		select {
		case p.waiters <- nil: // wake up a single waiter, if any
			break
		default:
			close(p.waiters) // or unblock all future waiters who are yet to enter the waiters queue
			p.waiters = nil
		}
	}
	p.mutex.Unlock()
	p.closeAll(head)
}

// Sets the deadline on the underlying net.Conn object
func (p *tubePool) setDeadline(ctx context.Context, tube tube) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	var deadline time.Time
	if d, ok := ctx.Deadline(); ok {
		deadline = d
	}
	return tube.SetDeadline(deadline)
}

// Closes the pool and all idle tubes in it.
func (p *tubePool) Close() error {
	p.mutex.Lock()

	var head tube
	if !p.closed {
		p.closed = true
		p.sessionBump()
		head = p.clearIdleConnections()
		if p.waiters != nil {
			close(p.waiters)
			p.waiters = nil
		}
		close(p.errCh)
		// cannot close(p.gate) as send on closed channel will panic. new connections will be closed immediately.
	}
	p.mutex.Unlock()
	p.closeAll(head)
	return nil
}

// Resets the idle tube stack by detaching existing tubes from it.
// p.mutex must be held when calling this method
func (p *tubePool) clearIdleConnections() tube {
	head := p.top
	p.top = nil
	p.lastActive = nil
	return head
}

// Closes tubes which weren't used since the last time this method was called.
func (p *tubePool) reapIdleConnections() {
	p.mutex.Lock()

	var reapHead tube
	if !p.closed {
		if p.lastActive != nil {
			reapHead = p.lastActive.Next()
			p.lastActive.SetNext(nil)
		}
		p.lastActive = p.top
	}
	p.mutex.Unlock()

	// closing tubes synchronously as this method is expected to be called from a background goroutine
	p.closeAll(reapHead)
}

// Allocates a new tube by establishing a new connection and performing initialization.
func (p *tubePool) alloc(session int64) (tube, error) {
	conn, err := p.dialContext(context.TODO(), network, p.address)
	if err != nil {
		return nil, err
	}
	t, err := newTube(conn, session)
	if err != nil {
		return nil, err
	}
	return t, nil
}

// Traverses the passed stack and closes all tubes in it.
func (p *tubePool) closeAll(head tube) {
	var next tube
	for head != nil {
		next = head.Next()
		head.SetNext(nil)
		head.Close()
		head = next
	}
}

// Increases the session version.
// Recycled or newly created tubes with the old session will be immediately closed
// p.mutex must be held when calling this method
func (p *tubePool) sessionBump() {
	p.session++
}

// Represents a semaphore limiting the total number of in-flight connection attempts.
// Being a channel it must be initialized with the desired limit as the buffer size.
type gate chan struct{}

// Returns true if we successfully acquired a permit, false otherwise
// gate#exit() must be called once the permit is no longer needed
func (g gate) tryEnter() bool {
	select {
	case g <- struct{}{}:
		return true
	default:
		return false
	}
}

// Exits the gate effectively returning a permit back into the pool
func (g gate) exit() {
	select { // do not block
	case <-g:
	default:
	}
}

type connectionReaper interface {
	reapIdleConnections()
}
