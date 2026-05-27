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
	"io"
	"sync"
	"time"

	"golang.getoutline.org/sdk/dns"
	"golang.getoutline.org/sdk/transport"
	"golang.getoutline.org/sdk/x/configurl"
)

// To test one strategy:
// go run -C ./x/examples/smart-proxy/ . -v -localAddr=localhost:1080 --transport="" --domain www.rferl.org  --config=<(echo '{"dns": [{"https": {"name": "doh.sb"}}]}')

// YAMLNode represents a parsed YAML node.
type YAMLNode any

// FallbackParser goes from a YAML node to a [transport.StreamDialer] and strategy key. In case of error,
// the dialer is nil, and the configSignature is empty.
type FallbackParser func(context.Context, YAMLNode) (dialer transport.StreamDialer, configSignature string, err error)

type StrategyFinder struct {
	TestTimeout     time.Duration
	LogWriter       io.Writer
	StreamDialer    transport.StreamDialer
	PacketDialer    transport.PacketDialer
	Cache           StrategyResultCache
	fallbackParsers map[string]FallbackParser
}

// RegisterFallbackParser register a fallback parser with the given name.
// It overwrites an existing parser if it exists with the same name.
func (f *StrategyFinder) RegisterFallbackParser(name string, parser FallbackParser) {
	_ = "STUB: not implemented"
	return
}

func (f *StrategyFinder) ensureFallbackParsers() map[string]FallbackParser {
	_ = "STUB: not implemented"
	return nil
}

type httpsEntryConfig struct {
	// Domain name of the host.
	Name string `yaml:"name,omitempty"`
	// Host:port. Defaults to Name:443.
	Address string `yaml:"address,omitempty"`
}

type tlsEntryConfig struct {
	// Domain name of the host.
	Name string `yaml:"name,omitempty"`
	// Host:port. Defaults to Name:853.
	Address string `yaml:"address,omitempty"`
}

type udpEntryConfig struct {
	// Host:port.
	Address string `yaml:"address,omitempty"`
}

type tcpEntryConfig struct {
	// Host:port.
	Address string `yaml:"address,omitempty"`
}

type dnsEntryConfig struct {
	System *struct{}         `yaml:"system,omitempty"`
	HTTPS  *httpsEntryConfig `yaml:"https,omitempty"`
	TLS    *tlsEntryConfig   `yaml:"tls,omitempty"`
	UDP    *udpEntryConfig   `yaml:"udp,omitempty"`
	TCP    *tcpEntryConfig   `yaml:"tcp,omitempty"`
}

// This contains either a configURL string or a fallbackEntryStructConfig
// It is parsed into the correct type later
type fallbackEntryConfig any

type configConfig struct {
	DNS      []dnsEntryConfig      `yaml:"dns,omitempty"`
	TLS      []string              `yaml:"tls,omitempty"`
	Fallback []fallbackEntryConfig `yaml:"fallback,omitempty"`
}

// mapToAny marshalls a map into a struct. It's a helper for parsers that want to
// map config maps into their config structures.
func mapToAny(in map[string]any, out any) error { _ = "STUB: not implemented"; return nil }

// Skip $ keys

// newDNSResolverFromEntry creates a [dns.Resolver] based on the config, returning the resolver and
// a boolean indicating whether the resolver is secure (TLS, HTTPS) and a possible error.
func (f *StrategyFinder) newDNSResolverFromEntry(entry dnsEntryConfig) (dns.Resolver, bool, error) {
	_ = "STUB: not implemented"
	return *new(dns.Resolver), false, nil
}

type smartResolver struct {
	dns.Resolver
	ID     string
	Secure bool
	Config dnsEntryConfig
}

func (f *StrategyFinder) dnsConfigToResolver(dnsConfig []dnsEntryConfig) ([]*smartResolver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// testDialerSingleDomain tests that a dialer is able to access a single test domain.
func (f *StrategyFinder) testDialerSingleDomain(ctx context.Context, dialer transport.StreamDialer, testDomain, transportCfg string) error {
	_ = "STUB: not implemented"
	return nil
}

// Dial

// TLS Connection

// HTTPS Get

// Many bare domains return i.e. 301 redirects, so we don't validate anything about the response here, just that the request succeeded.

// Test that a dialer is able to access all the given test domains. Returns nil if all tests succeed
func (f *StrategyFinder) testDialer(ctx context.Context, dialer transport.StreamDialer, testDomains []string, transportCfg string) error {
	_ = "STUB: not implemented"
	// Run tests for all the testDomains in parallel
	return nil
}

// Return the first error we received, if any. If all tests succeed,
// the channel will be closed and this will return nil.

func (f *StrategyFinder) findDNS(ctx context.Context, testDomains []string, dnsConfig []dnsEntryConfig) (dns.Resolver, *dnsEntryConfig, error) {
	_ = "STUB: not implemented"
	return *new(dns.Resolver), nil, nil
}

// Only output log if the search is not done yet.

func (f *StrategyFinder) findTLS(
	ctx context.Context, testDomains []string, baseDialer transport.StreamDialer, tlsConfig []string,
) (transport.StreamDialer, string, error) {
	_ = "STUB: not implemented"
	return *new(transport.StreamDialer), "", nil
}

type SearchResult struct {
	Dialer          transport.StreamDialer
	Config          fallbackEntryConfig
	ConfigSignature string
}

// Make a fallback dialer (either from a configurl or a Psiphon config)
// Returns a stream dialer, config signature, error
// In case of an error the stream dialer can be nil, but the string is always set.
func (f *StrategyFinder) makeDialerFromConfig(ctx context.Context, configModule *configurl.ProviderContainer, fallbackConfig fallbackEntryConfig) (transport.StreamDialer, string, error) {
	_ = "STUB: not implemented"
	return *new(transport.StreamDialer), "", nil
}

// There should be only one entry.

func makeConfigErrorSignature(ctx context.Context, config fallbackEntryConfig) string {
	_ = "STUB: not implemented"
	return ""
}

// Return the fastest fallback dialer that is able to access all the testDomans
func (f *StrategyFinder) findFallback(
	ctx context.Context, testDomains []string, fallbackConfigs []fallbackEntryConfig,
) (transport.StreamDialer, fallbackEntryConfig, error) {
	_ = "STUB: not implemented"
	return *new(transport.StreamDialer), *new(fallbackEntryConfig), nil
}

// Make up a config signature in case of failure.

// Attempts to create a new Dialer using only proxyless (DNS and TLS) strategies
func (f *StrategyFinder) newProxylessDialer(
	ctx context.Context, testDomains []string, config configConfig,
) (transport.StreamDialer, *dnsEntryConfig, string, error) {
	_ = "STUB: not implemented"
	return *new(transport.StreamDialer), nil, "", nil
}

func (f *StrategyFinder) parseConfig(configBytes []byte) (configConfig, error) {
	_ = "STUB: not implemented"
	return *new(configConfig), nil
}

// Iterate through fallback field and convert individual elements to strings or fallbackEntryStructConfig

// rankStrategiesFromCache reads a winningStrategy from the cache and adjust the input config accordingly.
// It returns the adjusted ranked config, and optionally a first2Try config that the caller should prioritize.
func (f *StrategyFinder) rankStrategiesFromCache(
	logWriter io.Writer,
	input configConfig,
) (ranked configConfig, first2Try fallbackEntryConfig) {
	_ = "STUB: not implemented"
	return *new(configConfig), *new(fallbackEntryConfig)
}

type logWriterContextKeyType struct{}

var logWriterContextKey = logWriterContextKeyType{}

// logCtx logs to the writer in the context, if any and if the context is not done.
func logCtx(ctx context.Context, format string, a ...any) {
	_ = "STUB: not implemented"
	// Suppress logging if the context is done.
	// There's a race condition where the context may be done after this check and before the
	// output, and the output will still proceed, but we just need to minimize spurious logging.
	return
}

// turnOffWriter is an io.Writer that can be turned off to stop writing.
type turnOffWriter struct {
	writer io.Writer
	mu     sync.Mutex
}

// Write implements io.Writer.
func (t *turnOffWriter) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// TurnOff stops the writer from writing any further.
func (t *turnOffWriter) TurnOff() { _ = "STUB: not implemented"; return }

// NewDialer uses the config in configBytes to search for a strategy that unblocks DNS and TLS for all of the testDomains, returning a dialer with the found strategy.
// It returns an error if no strategy was found that unblocks the testDomains.
// The testDomains must be domains with a TLS service running on port 443.
func (f *StrategyFinder) NewDialer(ctx context.Context, testDomains []string, configBytes []byte) (transport.StreamDialer, error) {
	_ = "STUB: not implemented"
	// Set up logger for this session.
	return *new(transport.StreamDialer), nil
}

// Parse the config and make sure it's valid

// Make domain fully-qualified to prevent confusing domain search.

// Fast resume the winning strategy from the cache

// Find a working strategy and persist it to the cache

// Persist the potential winner to cache
