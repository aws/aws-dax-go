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

package lru

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestLruGet(t *testing.T) {
	c := &Lru{
		LoadFunc: func(ctx aws.Context, key Key) (interface{}, error) {
			return key, nil
		},
	}

	for i := 0; i < 123; i++ {
		v, err := c.GetWithContext(nil, i)
		if err != nil {
			t.Fatalf("Lru.Get(%v) got error %v", i, err)
		}
		if !reflect.DeepEqual(i, v) {
			t.Fatalf("Lru.Get(%v) got %v want %v", i, v, i)
		}
	}
}

func TestLruKeyMarshaller(t *testing.T) {
	loadCount := 0
	c := &Lru{
		LoadFunc: func(ctx aws.Context, key Key) (interface{}, error) {
			loadCount++
			return key, nil
		},
		KeyMarshaller: func(key Key) Key {
			return fmt.Sprintf("%q", key)
		},
	}

	k := []string{"a", "b", "c"}
	for i := 0; i < 3; i++ {
		if v, err := c.GetWithContext(nil, k); err != nil {
			t.Errorf("unexpected error %v", err)
		} else if !reflect.DeepEqual(v, k) {
			t.Errorf("expected %v, got %v", k, v)
		}
		if loadCount != 1 {
			t.Errorf("expected %d, got %d", 1, loadCount)
		}
	}
}

func TestLruEvict(t *testing.T) {
	loads := 0
	loadFn := func(ctx aws.Context, key Key) (interface{}, error) {
		loads++
		return key, nil
	}

	c := &Lru{
		MaxEntries: 100,
		LoadFunc:   loadFn,
	}

	for i := 0; i < 123; i++ {
		v, err := c.GetWithContext(nil, i)
		if err != nil {
			t.Fatalf("Lru.Get(%v) got error %v", i, err)
		}
		if !reflect.DeepEqual(i, v) {
			t.Fatalf("Lru.Get(%v) got %v want %v", i, v, i)
		}
		if loads != i+1 {
			t.Fatalf("Lru.Get(%v) load calls got %v want %v", i, loads, i+1)
		}
	}

	for i := 0; i < 23; i++ {
		if c.contains(i) {
			t.Fatalf("Lru.contains(%v) want false", i)
		}
	}

	curr := loads
	for i := 23; i < 123; i++ {
		if !c.contains(i) {
			t.Fatalf("Lru.contains(%v) want true", i)
		}
		v, err := c.GetWithContext(nil, i)
		if err != nil {
			t.Fatalf("Lru.Get(%v) got error %v", i, err)
		}
		if !reflect.DeepEqual(i, v) {
			t.Fatalf("Lru.Get(%v) got %v want %v", i, v, i)
		}

		// Cached values should not reload.
		if loads != curr {
			t.Fatalf("Lru.Get(%v) load calls got %v want %v", i, loads, curr)
		}
	}
}

func TestLruTimeout(t *testing.T) {
	loadFn := func(ctx aws.Context, key Key) (interface{}, error) {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		}
		return key, nil
	}

	c := &Lru{
		MaxEntries: 100,
		LoadFunc:   loadFn,
	}

	ctx, cfn := context.WithTimeout(aws.BackgroundContext(), 1*time.Millisecond)
	defer cfn()
	key := "key1"
	v, err := c.GetWithContext(ctx, key)
	if err != ctx.Err() {
		t.Errorf("Lru.Get(%v) expected error %v, error %v", key, ctx.Err(), err)
	}
	if v != nil {
		t.Errorf("Lru.Get(%v) expected nil, got %v", key, v)
	}
}

func TestLruConcurrentLoad(t *testing.T) {
	var loads int32
	loadTime := 10 * time.Millisecond
	loadFn := func(ctx aws.Context, key Key) (interface{}, error) {
		<-time.After(loadTime)
		atomic.AddInt32(&loads, 1)
		return key, nil
	}

	c := &Lru{
		MaxEntries: 1000,
		LoadFunc:   loadFn,
	}

	keys := 100
	gets := 100
	var wg sync.WaitGroup
	wg.Add(keys * gets)
	st := time.Now()
	for k := 0; k < keys; k++ {
		for g := 0; g < gets; g++ {
			var key Key = k
			go func(key Key) {
				v, err := c.GetWithContext(nil, key)
				if err != nil {
					t.Errorf("Lru.Get(%v) got error %v", key, err)
				}
				if !reflect.DeepEqual(key, v) {
					t.Errorf("Lru.Get(%v) got %v want %v", key, v, key)
				}
				wg.Done()
			}(key)
		}
	}
	wg.Wait()
	elapsed := time.Since(st)

	if atomic.LoadInt32(&loads) != int32(keys) {
		t.Fatalf("Lru expected %d loads, had %d", keys, atomic.LoadInt32(&loads))
	}
	maxLoadTime := time.Duration(keys) * loadTime
	exp := maxLoadTime / 2
	if elapsed > exp {
		t.Fatalf("Lru loads (%v) were slower than expected (%v) for %d keys", elapsed, exp, keys)
	}
}

func TestLruSingleLoader(t *testing.T) {
	valueCh := make(chan interface{})
	loadFn := func(ctx aws.Context, key Key) (interface{}, error) {
		return <-valueCh, nil
	}

	c := &Lru{
		MaxEntries: 100,
		LoadFunc:   loadFn,
	}

	key := "key1"
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			v, err := c.GetWithContext(nil, key)
			if err != nil {
				t.Errorf("Lru.Get(%v) got error %v", key, err)
			}
			if !reflect.DeepEqual(key, v) {
				t.Errorf("Lru.Get(%v) got %v want %v", key, v, key)
			}
			wg.Done()
		}()
	}
	valueCh <- key
	wg.Wait()
}

func TestLoadGroup(t *testing.T) {
	loadCh := make(chan interface{})
	loadFn := func() (interface{}, error) {
		return <-loadCh, nil
	}

	key := "key1"
	l := &loadGroup{}
	done := make(chan struct{})
	go func() {
		v, err := l.do(key, loadFn)
		if err != nil {
			t.Errorf("unexpected error %v", err)
		}
		if !reflect.DeepEqual(key, v) {
			t.Errorf("expected %v, got %v", key, v)
		}
		done <- struct{}{}
	}()
	loadCh <- key
	<-done

	if len(l.m) > 0 {
		t.Errorf("expected 0, got %v", len(l.m))
	}
}

func BenchmarkLruGet(b *testing.B) {
	c := &Lru{
		LoadFunc: func(ctx aws.Context, key Key) (interface{}, error) {
			return key, nil
		},
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c.GetWithContext(nil, 123)
	}
}
