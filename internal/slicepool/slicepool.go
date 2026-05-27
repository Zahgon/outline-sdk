// Copyright 2020 The Outline Authors
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

package slicepool

import (
	"sync"
)

// Pool wraps a sync.Pool of *[]byte.  To encourage correct usage,
// all public methods are on slicepool.LazySlice.
//
// All copies of a Pool refer to the same underlying pool.
//
// "*[]byte" is used to avoid a heap allocation when passing a
// []byte to sync.Pool.Put, which leaks its argument to the heap.
type Pool struct {
	pool *sync.Pool
	len  int
}

// MakePool returns a Pool of slices with the specified length.
func MakePool(sliceLen int) Pool { _ = "STUB: not implemented"; return *new(Pool) }

// Return a *[]byte instead of []byte ensures that
// the []byte is not copied, which would cause a heap
// allocation on every call to sync.pool.Put

func (p *Pool) get() *[]byte { _ = "STUB: not implemented"; return nil }

func (p *Pool) put(b *[]byte) { _ = "STUB: not implemented"; return }

// LazySlice returns an empty LazySlice tied to this Pool.
func (p *Pool) LazySlice() LazySlice {
	_ = "STUB: not implemented"
	return *

	// LazySlice holds 0 or 1 slices from a particular Pool.
	new(LazySlice)
}

type LazySlice struct {
	slice *[]byte
	pool  *Pool
}

// Acquire this slice from the pool and return it.
// This slice must not already be acquired.
func (b *LazySlice) Acquire() []byte { _ = "STUB: not implemented"; return nil }

// Release the buffer back to the pool, unless the box is empty.
// The caller must discard any references to the buffer.
func (b *LazySlice) Release() { _ = "STUB: not implemented"; return }
