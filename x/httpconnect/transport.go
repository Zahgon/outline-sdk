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

package httpconnect

import (
	stdTLS "crypto/tls"
	"net"
	"net/http"

	"golang.getoutline.org/sdk/transport"
	"golang.getoutline.org/sdk/transport/tls"
)

type TransportOption func(c *transportConfig)

// WithTLSOptions configures the transport to use the given TLS options.
// The default behavior is to use TLS.
func WithTLSOptions(opts ...tls.ClientOption) TransportOption {
	_ = "STUB: not implemented"
	return *new(TransportOption)
}

// WithPlainHTTP configures the transport to use HTTP instead of HTTPS.
func WithPlainHTTP() TransportOption { _ = "STUB: not implemented"; return *new(TransportOption) }

// NewHTTPProxyTransport creates a net/http Transport that establishes a connection to the proxy using the given [transport.StreamDialer].
// The proxy address must be in the form "host:port".
//
// For HTTP/1 (plain and over TLS) and HTTP/2 (over TLS) over a stream connection.
// When using TLS, pass WithTLSOptions(tls.WithALPN()) to enable or enforce HTTP/2.
func NewHTTPProxyTransport(dialer transport.StreamDialer, proxyAddr string, opts ...TransportOption) (ProxyRoundTripper, error) {
	_ = "STUB: not implemented"
	return *new(ProxyRoundTripper), nil
}

// TLS config must be applied AFTER http2.ConfigureTransport, as it appends h2 to the list of supported protocols.

// NewH2ProxyTransport creates a pure HTTP/2 transport that establishes a connection to the proxy
// using the given [transport.StreamDialer].
// The proxy address must be in the form "host:port".
//
// Unlike [NewHTTPProxyTransport], this uses [golang.org/x/net/http2.Transport] directly, enabling:
//   - h2c (cleartext HTTP/2 via prior knowledge) with [WithPlainHTTP] — no TLS required
//   - Pure H2 from byte 1: multiple concurrent CONNECT tunnels share one TCP connection
func NewH2ProxyTransport(dialer transport.StreamDialer, proxyAddr string, opts ...TransportOption) (ProxyRoundTripper, error) {
	_ = "STUB: not implemented"
	return *new(ProxyRoundTripper), nil
}

// DialTLSContext is used even for plaintext when AllowHTTP is true.

// Ensure "h2" is in ALPN NextProtos so the server negotiates HTTP/2.

// http2.Transport type-asserts to *tls.Conn to read NegotiatedProtocol,
// so we must return *stdTLS.Conn directly — not an sdk tls.WrapConn wrapper.

// NewH3ProxyTransport creates an HTTP/3 transport that establishes a QUIC connection to the proxy using the given [net.PacketConn].
// The proxy address must be in the form "host:port".
//
// For HTTP/3 over QUIC over a datagram connection.
// [tls.WithALPN] has no effect on this transport.
func NewH3ProxyTransport(conn net.PacketConn, proxyAddr string, opts ...TransportOption) (ProxyRoundTripper, error) {
	_ = "STUB: not implemented"
	return *new(ProxyRoundTripper), nil
}

// HTTP/3 is always over TLS

type transportConfig struct {
	tlsOptions []tls.ClientOption
	plainHTTP  bool
}

func (c *transportConfig) applyOptions(opts ...TransportOption) { _ = "STUB: not implemented"; return }

type scheme string

const (
	schemeHTTP  scheme = "http"
	schemeHTTPS scheme = "https"
)

type proxyRT struct {
	http.RoundTripper
	scheme scheme
}

func (rt proxyRT) Scheme() string { _ = "STUB: not implemented"; return "" }

// TODO: Replace with tls.ToGoTLSConfig call once outline-sdk dependency version for this module is bumped.
// It is basically a copy of the implementation ToGoTLSConfig
func toStdConfig(cfg tls.ClientConfig) *stdTLS.Config { _ = "STUB: not implemented"; return nil }

// Set InsecureSkipVerify to skip the default validation we are
// replacing. This will not disable VerifyConnection.
