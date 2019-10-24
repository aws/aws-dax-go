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
	"net"
	"os"
	"sync"
	"time"
)

const network = "tcp"

// Acts as the gate to create new tubes
// and keeps track of tubes which are currently not in use.
type tubePool struct {
	address              string
	gate                 gate
	errCh                chan error
	timeout              time.Duration
	connectFn            func(string, string) (net.Conn, error)
	closeTubeImmediately bool

	mutex      sync.Mutex
	closed     bool  // protected by mutex
	top        *tube // protected by mutex
	lastActive *tube // protected by mutex
	waiters    chan *tube
}

type tubePoolOptions struct {
	maxConcurrentConnAttempts int
	timeout                   time.Duration
}

var defaultTubePoolOptions = tubePoolOptions{10, time.Second * 5}

// Creates a new pool using defaultTubePoolOptions and associated with given address.
func newTubePool(address string) *tubePool {
	return newTubePoolWithOptions(address, defaultTubePoolOptions)
}

// Creates a new pool with provided options associated with the given address.
func newTubePoolWithOptions(address string, options tubePoolOptions) *tubePool {
	if options.maxConcurrentConnAttempts <= 0 {
		options.maxConcurrentConnAttempts = defaultTubePoolOptions.maxConcurrentConnAttempts
	}
	return &tubePool{
		address:   address,
		gate:      make(gate, options.maxConcurrentConnAttempts),
		errCh:     make(chan error),
		waiters:   make(chan *tube),
		timeout:   options.timeout,
		connectFn: connect,
	}
}

// Gets a new or reuses existing tube with timeout context set to tubePool#timeout
func (p *tubePool) get() (*tube, error) {
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
func (p *tubePool) getWithContext(ctx context.Context, highPriority bool) (*tube, error) {
	for {
		p.mutex.Lock()
		if p.closed {
			p.mutex.Unlock()
			return nil, os.ErrClosed
		}

		// look for idle tubes in stack
		if p.top != nil {
			tube := p.top
			p.top = tube.next
			if p.lastActive == tube {
				p.lastActive = p.top
			}
			tube.next = nil
			p.mutex.Unlock()
			return tube, nil
		}

		// no tubes in stack, create wait channel
		if p.waiters == nil {
			p.waiters = make(chan *tube)
		}
		waitCh := p.waiters
		p.mutex.Unlock()

		var done chan *tube
		if p.gate.tryEnter() {
			go p.allocAndReleaseGate(done, true)
		} else if highPriority {
			done = make(chan *tube)
			go p.allocAndReleaseGate(done, false)
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
func (p *tubePool) allocAndReleaseGate(done chan *tube, releaseGate bool) {
	tube, err := p.alloc()
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
// Tube will be closed if the pool is closed
// or it will be handed over to a waiter if any
// or it will be added to the idle tubes stack.
func (p *tubePool) put(tube *tube) {
	if tube == nil {
		return
	}

	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.closed {
		tube.Close()
		// Waiters channel was already closed in Close
		return
	}

	if p.waiters != nil {
		select {
		case p.waiters <- tube:
			return
		default:
			close(p.waiters) // unblock future waiters
			p.waiters = nil
		}
	}

	tube.next = p.top
	p.top = tube
	return
}

// Discards the given tube indicating that it must no longer be used and has to be closed.
func (p *tubePool) discard(tube *tube) {
	if tube == nil {
		return
	}
	if p.closeTubeImmediately {
		tube.Close()
	} else {
		go func() {
			tube.Close()
		}()
	}

	p.mutex.Lock()
	defer p.mutex.Unlock()

	// Waiters enter the waiting queue when there's no existing tube
	// or when they failed to acquire a permit to create a new tube.
	// There's also a chance the newly created tube was stolen and
	// the thief must return it back into the pool or discard it.
	if p.waiters != nil {
		select {
		case p.waiters <- nil: // wake up a single waiter, if any
			return
		default:
			close(p.waiters) // or unblock all future waiters who are yet to enter the waiters queue
			p.waiters = nil
		}
	}
}

// Sets the deadline on the underlying net.Conn object
func (p *tubePool) setDeadline(ctx context.Context, tube *tube) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	var deadline time.Time
	if d, ok := ctx.Deadline(); ok {
		deadline = d
	}
	return tube.setDeadline(deadline)
}

// Closes the pool and tubes in the pool.
func (p *tubePool) Close() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.closed {
		return nil
	}

	p.closed = true

	t := p.top
	for t != nil {
		t.Close()
		t = t.next
	}
	p.top = nil
	p.lastActive = nil
	if p.waiters != nil {
		close(p.waiters)
		p.waiters = nil
	}
	close(p.errCh)
	// cannot close(p.gate) as send on closed channel will panic. new connections will be closed immediately.
	return nil
}

// Closes tubes which weren't used since the last time this method was called.
func (p *tubePool) reapIdleConnections() {
	p.mutex.Lock()

	if p.closed {
		p.mutex.Unlock()
		return
	}

	var reapHead *tube
	if p.lastActive != nil {
		reapHead = p.lastActive.next
		p.lastActive.next = nil
	}
	p.lastActive = p.top
	p.mutex.Unlock()

	// closing tubes synchronously as this method is expected to be called from a background goroutine
	p.closeAll(reapHead)
}

// Allocates a new tube by establishing a new connection and performing initialization.
func (p *tubePool) alloc() (*tube, error) {
	conn, err := p.connectFn(network, p.address)
	if err != nil {
		return nil, err
	}
	tube, err := newTube(conn)
	if err != nil {
		tube.Close()
		return nil, err
	}
	if err = tube.init(); err != nil {
		tube.Close()
		return nil, err
	}
	return tube, nil
}

// Traverses the passed stack and closes all tubes in it.
func (p *tubePool) closeAll(head *tube) {
	var next *tube
	for head != nil {
		next = head.next
		head.next = nil
		head.Close()
		head = next
	}
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

func connect(network, address string) (net.Conn, error) {
	return net.Dial(network, address)
}

type connectionReaper interface {
	reapIdleConnections()
}
