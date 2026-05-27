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

package smart

// StrategyResultCache is a cache of strategy results that can be used by [StrategyFinder]
// to resume a strategy efficiently.
// Implementations are expected to be called concurrently from different goroutines.
type StrategyResultCache interface {
	// Get retrieves a strategy result value associated with the given key.
	// It returns the value text encoded in UTF-8 and true if found.
	Get(key string) (value []byte, ok bool)

	// Put adds the strategy result value encoded in UTF-8 to the cache with the given key.
	// If called with nil value, it should remove the cache entry.
	Put(key string, value []byte)
}

// winningStrategyCacheKey is the key for storing the winning strategy in the
// [StrategyResultCache] each time [StrategyFinder].NewDialer is invoked.
const winningStrategyCacheKey = "winning_strategy"

// winningConfig holds the configuration of a successful strategy.
// It contains either one entry of proxyless or one entry of fallback.
type winningConfig configConfig

func newProxylessWinningConfig(dns *dnsEntryConfig, tls string) winningConfig {
	_ = "STUB: not implemented"
	return *new(winningConfig)
}

func newFallbackWinningConfig(fallback fallbackEntryConfig) winningConfig {
	_ = "STUB: not implemented"
	return *new(winningConfig)
}

// getFallbackIfExclusive checks if the winningConfig is a fallback strategy.
// It returns the fallback entry and true if there is only one exclusive fallback entry
// in the config; otherwise it returns nil and false.
func (w winningConfig) getFallbackIfExclusive(cfg *configConfig) (fallbackEntryConfig, bool) {
	_ = "STUB: not implemented"
	return *new(fallbackEntryConfig), false
}

// promoteProxylessToFront reorders the DNS and TLS configs within the provided configConfig
// to move the entries matching the winning strategy to the front.
func (w winningConfig) promoteProxylessToFront(cfg *configConfig) {
	_ = "STUB: not implemented"
	return

	// If not found, IndexFunc will return -1, and moveToFront will ignore
}

// If not found, Index will return -1, and moveToFront will ignore

func (w winningConfig) toYAML() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
