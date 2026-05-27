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

package configurl

import (
	"net/url"

	"golang.getoutline.org/sdk/transport"
	"golang.getoutline.org/sdk/x/httpconnect"
)

// connectOptions holds parsed transport and client options for an HTTP CONNECT proxy URL.
type connectOptions struct {
	transport []httpconnect.TransportOption
	client    []httpconnect.ClientOption
}

// parseConnectOptions parses query parameters and userinfo from a hierarchical config URL into
// transport and client options.
//
// Supported query parameters:
//   - sni: TLS server name for SNI.
//   - certname: name to validate against the server certificate.
//   - plain: if "true", use cleartext (no TLS). Only meaningful for h2connect (h2c).
//   - auth: raw Proxy-Authorization header value (e.g. "Bearer mytoken").
//
// URL userinfo (user:password@host) is translated to a Proxy-Authorization: Basic header.
// Use ?auth= for other schemes such as Bearer.
func parseConnectOptions(configURL url.URL) (connectOptions, error) {
	_ = "STUB: not implemented"
	return *

	// Userinfo → Proxy-Authorization: Basic
	new(connectOptions), nil
}

// sanitizeConnectURL redacts credentials from an HTTP CONNECT proxy URL:
// userinfo (user:pass) is replaced with REDACTED, and the ?auth= parameter value is redacted.
func sanitizeConnectURL(u url.URL) (string, error) { _ = "STUB: not implemented"; return "", nil }

// registerHTTPConnectStreamDialer registers an HTTP CONNECT proxy transport (H1.1, or H2 via ALPN).
//
// Config format: httpconnect://[user:pass@]host:port[?sni=SNI][&certname=CERTNAME][&auth=TOKEN]
//
// The base dialer (from the previous element in the pipe chain) is used to establish
// the TCP connection to the proxy. TLS is negotiated by the transport itself.
// When H2 is negotiated via ALPN, CONNECT streams are multiplexed over the single TCP connection.
func registerHTTPConnectStreamDialer(r TypeRegistry[transport.StreamDialer], typeID string, newSD BuildFunc[transport.StreamDialer]) {
	_ = "STUB: not implemented"
	return
}

// registerH2ConnectStreamDialer registers a pure HTTP/2 CONNECT proxy transport.
//
// Config format: h2connect://[user:pass@]host:port[?sni=SNI][&certname=CERTNAME][&auth=TOKEN][&plain=true]
//
// Unlike httpconnect, all CONNECT streams are multiplexed over a single TCP connection
// to the proxy. The base dialer is used to establish that connection.
func registerH2ConnectStreamDialer(r TypeRegistry[transport.StreamDialer], typeID string, newSD BuildFunc[transport.StreamDialer]) {
	_ = "STUB: not implemented"
	return
}

// registerH3ConnectStreamDialer registers an HTTP/3 CONNECT proxy transport over QUIC.
//
// Config format: h3connect://[user:pass@]host:port[?sni=SNI][&certname=CERTNAME][&auth=TOKEN]
//
// A UDP socket is created internally and shared across all CONNECT streams (QUIC multiplexing).
// The base stream dialer is not used; QUIC always runs over a fresh UDP connection.
func registerH3ConnectStreamDialer(r TypeRegistry[transport.StreamDialer], typeID string) {
	_ = "STUB: not implemented"
	return
}
