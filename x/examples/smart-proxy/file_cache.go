// Copyright 2025 The Outline Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"sync"
)

// JSONFileCache implements a key-value cache using a JSON file.
type JSONFileCache struct {
	path  string
	mu    sync.RWMutex
	cache map[string]string
}

// NewJSONFileCache creates a new JSONFileCache.
func NewJSONFileCache(path string) (*JSONFileCache, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// flushNoLock writes the current cache data to the JSON file without holding the lock.
func (c *JSONFileCache) flushNoLock() error { _ = "STUB: not implemented"; return nil }

// Get retrieves a strategy result string associated with the given key.
func (c *JSONFileCache) Get(key string) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Put adds the strategy result string to the cache with the given key.
func (c *JSONFileCache) Put(key string, val []byte) { _ = "STUB: not implemented"; return }
