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
	"errors"
	"io"
	"log"
	"net"
	"strings"
	"sync"
	"testing"
	"time"
)

const localConnTimeoutMillis = 10

func TestTubePoolConnectionCache(t *testing.T) {
	endpoint := ":8181"
	var actualConnections, expectedConnections int
	startConnNotifier := make(chan net.Conn, 25)
	endConnNotifier := make(chan net.Conn, 25)
	listener, err := startServer(endpoint, startConnNotifier, endConnNotifier, drainAndCloseConn)
	if err != nil {
		t.Fatalf("cannot start server")
	}
	defer listener.Close()

	pool := newTubePoolWithOptions(endpoint, tubePoolOptions{10, time.Second * 1})

	// verify tube is re-used
	expectedConnections = 1
	attempts := 3
	for i := 0; i < attempts; i++ {
		tube, err := pool.get()
		if err != nil {
			t.Errorf("unexpected error %v", err)
		}
		select {
		case <-startConnNotifier:
			actualConnections++
		case <-endConnNotifier:
			t.Errorf("unexpected connection term")
		case <-time.After(time.Millisecond * localConnTimeoutMillis):
		}

		if actualConnections != expectedConnections {
			t.Errorf("expected %v, actual %v", expectedConnections, actualConnections)
		}

		pool.put(tube)
	}

	// verify new connections are established if cached tube not available
	expectedConnections = 0
	attempts = 3
	tubes := make([]*tube, attempts)
	for i := 0; i < attempts; i++ {
		tube, err := pool.get()
		tubes[i] = tube
		if err != nil {
			t.Errorf("unexpected error %v", err)
		}
		expectedConnections++
		select {
		case <-startConnNotifier:
			actualConnections++
		case <-endConnNotifier:
			t.Errorf("unexpected connection term")
		case <-time.After(time.Millisecond * localConnTimeoutMillis):
		}

		if actualConnections != expectedConnections {
			t.Errorf("expected %v, actual %v", expectedConnections, actualConnections)
		}
	}

	// verify connections are kept alive when returned
	for i := 0; i < attempts; i++ {
		pool.put(tubes[i])
		select {
		case <-startConnNotifier:
			t.Errorf("unexpected connection init")
		case <-endConnNotifier:
			t.Errorf("unexpected connection term")
		case <-time.After(time.Millisecond * localConnTimeoutMillis):
		}
	}

	// verify tubes cache is lifo
	for i := 0; i < len(tubes); i++ {
		tube, err := pool.get()
		if err != nil {
			t.Errorf("unexpected error %v", err)
		}
		select {
		case <-startConnNotifier:
			t.Errorf("unexpected connection init")
		case <-endConnNotifier:
			t.Errorf("unexpected connection term")
		case <-time.After(time.Millisecond * localConnTimeoutMillis):
		}

		if tube != tubes[len(tubes)-i-1] {
			t.Errorf("expected most recent tube")
		}
	}
}

func TestTubePool_reapIdleTubes(t *testing.T) {
	endpoint := ":8182"
	startConnNotifier := make(chan net.Conn, 25)
	endConnNotifier := make(chan net.Conn, 25)
	listener, err := startServer(endpoint, startConnNotifier, endConnNotifier, drainAndCloseConn)
	if err != nil {
		t.Fatalf("cannot start server")
	}
	defer listener.Close()

	pool := newTubePool(endpoint)

	tubeCount := 10
	tubes := make([]*tube, tubeCount)
	for i := 0; i < tubeCount; i++ {
		tubes[i], err = pool.get()
		if err != nil {
			t.Errorf("unexpected error %v", err)
		}
	}

	for i := 0; i < tubeCount; i++ {
		pool.put(tubes[i])
	}

	pool.reapIdleConnections()
	if countTubes(pool) != tubeCount {
		t.Errorf("expected cached tube count %v, actual %v", tubeCount, countTubes(pool))
	}

	active := make([]*tube, 0, tubeCount)
	activeCount := 5
	for i := 0; i < activeCount; i++ {
		tb, err := pool.get()
		if err != nil {
			t.Errorf("unexpected error %v", err)
		}
		active = append([]*tube{tb}, active...)
	}

	pool.reapIdleConnections()
	if countTubes(pool) != 1 {
		t.Errorf("expected cached tube count %v, actual %v", 1, countTubes(pool))
	}

	putCount := (activeCount + 1) / 2
	for i := 0; i < putCount; i++ {
		activeCount--
		a := active[activeCount]
		active = active[0:activeCount]
		pool.put(a)
	}

	pool.reapIdleConnections()
	if countTubes(pool) != putCount+1 {
		t.Errorf("expected cached tube count %v, actual %v", putCount+1, countTubes(pool))
	}

	count := len(active)
	for _, a := range active {
		pool.put(a)
	}

	pool.reapIdleConnections()
	if countTubes(pool) != count+1 {
		t.Errorf("expected cached tube count %v, actual %v", count+1, countTubes(pool))
	}
}

func TestTubePool_Close(t *testing.T) {
	endpoint := ":8183"
	startConnNotifier := make(chan net.Conn, 25)
	endConnNotifier := make(chan net.Conn, 25)
	listener, err := startServer(endpoint, startConnNotifier, endConnNotifier, drainAndCloseConn)
	if err != nil {
		t.Fatalf("could not start server")
	}
	defer listener.Close()

	pool := newTubePoolWithOptions(endpoint, tubePoolOptions{1, time.Second * 1})
	tubes := make([]*tube, 2)
	for i := 0; i < 2; i++ {
		tubes[i], err = pool.get()
		if err != nil {
			t.Fatalf("unexpected error %v", err)
		}
		select {
		case <-startConnNotifier:
		case <-time.After(time.Second * 1):
			t.Fatalf("could not establish connection")
		}
	}

	pool.put(tubes[0])
	pool.Close()

	select {
	case <-endConnNotifier:
	case <-time.After(time.Millisecond * localConnTimeoutMillis):
		t.Fatalf("cached connection was not terminated")
	}

	pool.put(tubes[1])
	select {
	case <-endConnNotifier:
	case <-time.After(time.Millisecond * localConnTimeoutMillis):
		t.Fatalf("tube returned to a closed pool was not terminated")
	}
}

func TestTubePoolError(t *testing.T) {
	endpoint := ":8184"
	pool := newTubePoolWithOptions(endpoint, tubePoolOptions{10, time.Second * 1})
	_, err := pool.get()
	if err == nil || !strings.Contains(err.Error(), "connection refused") {
		t.Errorf("expected 'dial tcp :8184: connection refused', actual '%v'\n", err)
	}
}

func TestConnectionPriority(t *testing.T) {
	endpoint := ":8185"
	listener, err := startServer(endpoint, nil, nil, drainAndCloseConn)
	if err != nil {
		t.Fatalf("cannot start server")
	}
	defer listener.Close()

	maxAttempts := 5
	var delay sync.WaitGroup
	delay.Add(maxAttempts + 1)
	connectFn := func(network, address string) (net.Conn, error) {
		delay.Done()
		delay.Wait()
		return net.Dial(network, address)
	}

	pool := newTubePoolWithOptions(endpoint, tubePoolOptions{maxAttempts, 1 * time.Second})
	pool.connectFn = connectFn
	defer pool.Close()

	var wg sync.WaitGroup
	wg.Add(maxAttempts)
	for i := 0; i < maxAttempts; i++ {
		go func() {
			defer wg.Done()
			tb, err := pool.getWithContext(context.Background(), false)
			if err != nil {
				t.Errorf("unexpected error %v", err)
			} else {
				tb.Close()
			}
		}()
	}

	ctx, cfn := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cfn()
	_, err = pool.getWithContext(ctx, false)
	if err == nil {
		t.Errorf("expected error, got none")
	}
	if err != ctx.Err() {
		t.Errorf("expected %v, got %v", ctx.Err(), err)
	}

	tb, err := pool.getWithContext(context.Background(), true)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	} else {
		tb.Close()
	}

	wg.Wait()
}

func TestGate(t *testing.T) {
	size := 3
	g := make(gate, size)
	// fill all slots
	for i := 0; i < size; i++ {
		if !g.tryEnter() {
			t.Errorf("expected gate access allowed")
		}
	}

	// verify further access is not allowed
	for i := 0; i < (size * 3); i++ {
		if g.tryEnter() {
			t.Errorf("expected gate access denied")
		}
	}

	// free up one slot, verify access is allowed
	g.exit()
	if !g.tryEnter() {
		t.Errorf("expected gate access allowed")
	}

	// free up all slots and more
	for i := 0; i < (size * 3); i++ {
		g.exit()
	}

	// fill all slots
	for i := 0; i < size; i++ {
		if !g.tryEnter() {
			t.Errorf("expected gate access allowed")
		}
	}

	// verify further access is not allowed
	for i := 0; i < (size * 3); i++ {
		if g.tryEnter() {
			t.Errorf("expected gate access denied")
		}
	}
}

func startServer(endpoint string, startConnNotifier chan net.Conn, endConnNotifier chan net.Conn,
	connectionHandler func(conn net.Conn, endConnNotifier chan net.Conn)) (net.Listener, error) {
	listener, err := net.Listen(network, endpoint)
	if err != nil {
		return nil, errors.New("cannot create server")
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				//log.Printf("cannot accept connection %v\n", err)
				return
			}
			startConnNotifier <- conn
			go connectionHandler(conn, endConnNotifier)
		}
	}()

	return listener, nil
}

func drainAndCloseConn(conn net.Conn, endConnNotifier chan net.Conn) {
	b := make([]byte, 1024)
	for {
		_, err := conn.Read(b)
		if err == io.EOF {
			endConnNotifier <- conn
			if err = conn.Close(); err != nil {
				log.Printf("unexpected error %v", err)
			}
			return
		} else if err != nil {
			log.Printf("unexpected error %v", err)
			return
		}
	}
}

func countTubes(pool *tubePool) int {
	head := pool.top
	count := 0
	for head != nil {
		count++
		head = head.next
	}
	return count
}
