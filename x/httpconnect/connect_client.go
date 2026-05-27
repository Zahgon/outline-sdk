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
	"context"
	"net/http"

	"golang.getoutline.org/sdk/transport"
)

// ConnectClient is a [transport.StreamDialer] that establishes an HTTP CONNECT tunnel over an abstract HTTP transport.
//
// The package also includes transport builders:
// - NewHTTPProxyTransport
// - NewH2ProxyTransport
// - NewH3ProxyTransport
//
// Options:
// - WithHeaders appends the provided headers to every CONNECT request.
type ConnectClient struct {
	proxyRT ProxyRoundTripper
	headers http.Header
}

var _ transport.StreamDialer = (*ConnectClient)(nil)

// ProxyRoundTripper is the minimal interface required by ConnectClient to send HTTP CONNECT requests.
// The Scheme method is used to construct the request URL, and the RoundTrip method is used to send the request.
type ProxyRoundTripper interface {
	http.RoundTripper
	Scheme() string
}

// ClientOption is an option for configuring the ConnectClient.
type ClientOption func(c *clientConfig)

// NewConnectClient creates a new ConnectClient that uses the provided ProxyRoundTripper to send HTTP CONNECT requests.
// The returned client implements the [transport.StreamDialer] interface.
func NewConnectClient(proxyRT ProxyRoundTripper, opts ...ClientOption) (*ConnectClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WithHeaders appends the given headers to the CONNECT request.
func WithHeaders(headers http.Header) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

type clientConfig struct {
	headers http.Header
}

// DialStream implements the [transport.StreamDialer] interface by sending an HTTP CONNECT request to the proxy and returning a connection that tunnels to the target address.
func (cc *ConnectClient) DialStream(ctx context.Context, remoteAddr string) (transport.StreamConn, error) {
	_ = "STUB: not implemented"
	return *new(transport.StreamConn), nil
}

// -1 means length unknown

// to provide the SetReadDeadline function of the returned connection, at the expense of an extra copy

func mergeHeaders(dst http.Header, src http.Header) { _ = "STUB: not implemented"; return }
