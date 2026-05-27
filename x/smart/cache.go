// Copyright 2023 The Outline Authors
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

package smart

import (
	"context"
	"sync"
	"time"

	"golang.getoutline.org/sdk/dns"
	"golang.org/x/net/dns/dnsmessage"
)

// canonicalName returns the domain name in canonical form. A name in canonical
// form is lowercase and fully qualified. Only US-ASCII letters are affected. See
// Section 6.2 in RFC 4034.
func canonicalName(s string) string { _ = "STUB: not implemented"; return "" }

type cacheEntry struct {
	key    string
	msg    *dnsmessage.Message
	expire time.Time
}

// simpleLRUCacheResolver is a very simple caching [dns.Resolver].
// It doesn't use the response TTL.
// It also doesn't dedup duplicate in-flight requests.
type simpleLRUCacheResolver struct {
	resolver dns.Resolver
	cache    []cacheEntry
	mux      sync.Mutex
}

var _ dns.Resolver = (*simpleLRUCacheResolver)(nil)

func newSimpleLRUCacheResolver(resolver dns.Resolver, numEntries int) dns.Resolver {
	_ = "STUB: not implemented"
	return *new(dns.Resolver)
}

func (r *simpleLRUCacheResolver) RemoveExpired() { _ = "STUB: not implemented"; return }

func (r *simpleLRUCacheResolver) moveToFront(index int) { _ = "STUB: not implemented"; return }

func makeCacheKey(q dnsmessage.Question) string { _ = "STUB: not implemented"; return "" }

func (r *simpleLRUCacheResolver) SearchCache(key string) *dnsmessage.Message {
	_ = "STUB: not implemented"
	return nil
}

// TODO: update TTLs
// TODO: make names match

func (r *simpleLRUCacheResolver) AddToCache(key string, msg *dnsmessage.Message) {
	_ = "STUB: not implemented"
	return
}

// TODO: copy and normalize names

// Query implements [dns.Resolver].
func (r *simpleLRUCacheResolver) Query(ctx context.Context, q dnsmessage.Question) (*dnsmessage.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: cache server failures. See https://datatracker.ietf.org/doc/html/rfc2308.
