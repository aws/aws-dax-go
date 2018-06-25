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

type tubePool struct {
	address   string
	gate      gate
	errCh     chan error
	timeout   time.Duration
	connectFn func(string, string) (net.Conn, error)

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

func newTubePool(address string) *tubePool {
	return newTubePoolWithOptions(address, defaultTubePoolOptions)
}

func newTubePoolWithOptions(a string, options tubePoolOptions) *tubePool {
	if options.maxConcurrentConnAttempts <= 0 {
		options.maxConcurrentConnAttempts = defaultTubePoolOptions.maxConcurrentConnAttempts
	}
	pool := tubePool{
		address:   a,
		gate:      make(gate, options.maxConcurrentConnAttempts),
		errCh:     make(chan error),
		waiters:   make(chan *tube),
		timeout:   options.timeout,
		connectFn: connect,
	}
	return &pool
}

func (p *tubePool) get() (*tube, error) {
	ctx := context.Background()
	if p.timeout > 0 {
		var cancelFn func()
		ctx, cancelFn = context.WithTimeout(ctx, p.timeout)
		defer cancelFn()
	}
	return p.getWithContext(ctx, false)
}

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
		if highPriority {
			done = make(chan *tube)
		}
		if p.gate.tryEnter() {
			go p.allocAndReleaseGate(done, true)
		} else if highPriority {
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
			return nil, err
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}

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

func (p *tubePool) put(tube *tube) {
	if tube == nil {
		return
	}

	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.closed {
		tube.Close()
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

func (p *tubePool) discard(tube *tube) {
	if tube == nil {
		return
	}
	go func() {
		tube.Close()
	}()
}

func (p *tubePool) setDeadline(tube *tube, ctx context.Context) error {
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

func (p *tubePool) Close() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
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

func (p *tubePool) alloc() (*tube, error) {
	conn, err := p.connectFn(network, p.address)
	if err != nil {
		return nil, err
	}
	tube, err := newTube(conn)
	if err != nil {
		return nil, err
	}
	if err = tube.init(); err != nil {
		tube.Close()
		return nil, err
	}
	return tube, nil
}

func (p *tubePool) closeAll(head *tube) {
	var next *tube
	for head != nil {
		next = head.next
		head.next = nil
		head.Close()
		head = next
	}
}

type gate chan struct{}

func (g gate) tryEnter() bool {
	select {
	case g <- struct{}{}:
		return true
	default:
		return false
	}
}

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
