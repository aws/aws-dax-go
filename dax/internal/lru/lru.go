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
	"sync"
)

// Lru is a cache which is safe for concurrent access.
type Lru struct {
	// MaxEntries is the maximum number of cache entries
	// before an item is evicted. Zero means no limit.
	MaxEntries int

	// LoadFunc specifies the function that loads a value
	// for a specific key when not found in the cache.
	LoadFunc  func(ctx context.Context, key Key) (interface{}, error)
	loadGroup loadGroup

	// Optional KeyMarshaller. Caller should provide one when using
	// Key type which is not comparable. eg. slice
	KeyMarshaller func(key Key) Key

	mu         sync.RWMutex
	cache      map[Key]*entry
	head, tail *entry
}

type Key interface{}

type entry struct {
	key        Key
	value      interface{}
	prev, next *entry
}

func (c *Lru) contains(key Key) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, ok := c.cache[key]
	return ok
}

func (c *Lru) lookup(key Key) (*entry, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.cache[key]
	return v, ok
}

func (c *Lru) GetWithContext(ctx context.Context, okey Key) (interface{}, error) {
	ikey := okey
	if c.KeyMarshaller != nil {
		ikey = c.KeyMarshaller(okey)
	}

	if en, ok := c.lookup(ikey); ok {
		return en.value, nil
	}

	v, err := c.loadGroup.do(ikey, func() (interface{}, error) {
		if en, ok := c.lookup(ikey); ok {
			return en.value, nil
		}

		val, err := c.LoadFunc(ctx, okey)
		if err != nil {
			return nil, err
		}

		c.mu.Lock()
		defer c.mu.Unlock()
		en := &entry{key: ikey, value: val}
		if c.tail == nil {
			c.head = en
			c.tail = en
		} else {
			en.prev = c.tail
			c.tail.next = en
			c.tail = en
		}

		if c.cache == nil {
			c.cache = make(map[Key]*entry)
		}
		c.cache[ikey] = en

		// Evict oldest entry if over the max.
		if c.MaxEntries > 0 && len(c.cache) > c.MaxEntries {
			evict := c.head
			if evict != nil {
				delete(c.cache, evict.key)
				c.head = evict.next
				if c.head != nil {
					c.head.prev = nil
				}
				evict.next = nil
			}
		}
		return val, nil
	})
	return v, err
}

type loader struct {
	wg    sync.WaitGroup
	value interface{}
	err   error
}

type loadGroup struct {
	mu sync.Mutex
	m  map[Key]*loader
}

func (g *loadGroup) do(key Key, loadFn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[Key]*loader)
	}
	if l, ok := g.m[key]; ok {
		g.mu.Unlock()
		l.wg.Wait()
		return l.value, l.err
	}
	v := &loader{}
	v.wg.Add(1)
	g.m[key] = v
	g.mu.Unlock()

	v.value, v.err = loadFn()
	v.wg.Done()

	// cleanup loader
	g.mu.Lock()
	defer g.mu.Unlock()
	delete(g.m, key)
	if len(g.m) <= 0 {
		g.m = nil
	}
	return v.value, v.err
}
