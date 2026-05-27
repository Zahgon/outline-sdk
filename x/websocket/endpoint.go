// Copyright 2025 The Outline Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package websocket provides the Websocket transport.
package websocket

import (
	"context"
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"golang.getoutline.org/sdk/transport"
)

// NewStreamEndpoint creates a new Websocket Stream Endpoint. Streams are sent over
// Websockets, with each Write becoming a separate message. Half-close is supported:
// CloseRead will not close the Websocket connection, while CloseWrite sends a Websocket
// close but continues reading until a close is received from the server.
func NewStreamEndpoint(urlStr string, se transport.StreamEndpoint, opts ...Option) (func(context.Context) (transport.StreamConn, error), error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewPacketEndpoint creates a new Websocket Packet Endpoint. Each packet is exchanged as a Websocket message.
func NewPacketEndpoint(urlStr string, se transport.StreamEndpoint, opts ...Option) (func(context.Context) (net.Conn, error), error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type options struct {
	tlsConfig *tls.Config
	headers   http.Header
}

// Option for building the Websocket endpoint.
type Option func(c *options)

// WithTLSConfig specifies the TLS configuration to use.
// TODO(fortuna): Use Outline TLS instead.
func WithTLSConfig(tlsConfig *tls.Config) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithHTTPHeaders specifies the HTTP headers to use.
func WithHTTPHeaders(headers http.Header) Option { _ = "STUB: not implemented"; return *new(Option) }

func newEndpoint[ConnType net.Conn](urlStr string, se transport.StreamEndpoint, wsToConn func(*gorillaConn) ConnType, opts ...Option) (func(context.Context) (ConnType, error), error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// By default, we use this User-Agent.

func newGorillaConn(wsConn *websocket.Conn) *gorillaConn { _ = "STUB: not implemented"; return nil }

type gorillaConn struct {
	wsConn *websocket.Conn

	// websocket.Conn is not safe for concurrent use
	// https://github.com/OutlineFoundation/outline-apps/issues/2573
	readMu, writeMu sync.Mutex

	writeErr      error
	readErr       error
	pendingReader io.Reader
}

var _ transport.StreamConn = (*gorillaConn)(nil)

func (c *gorillaConn) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (c *gorillaConn) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (c *gorillaConn) SetDeadline(deadline time.Time) error { _ = "STUB: not implemented"; return nil }

func (c *gorillaConn) SetReadDeadline(deadline time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *gorillaConn) SetWriteDeadline(deadline time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *gorillaConn) Read(buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *gorillaConn) Write(buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *gorillaConn) CloseRead() error { _ = "STUB: not implemented"; return nil }

func (c *gorillaConn) CloseWrite() error {
	_ = "STUB: not implemented"
	// Send close message.
	return nil
}

func (c *gorillaConn) Close() error { _ = "STUB: not implemented"; return nil }

// Upgrade upgrades an HTTP connection to a WebSocket connection. It returns a
// [transport.StreamConn] representing the WebSocket connection, or an error if
// the upgrade fails.
func Upgrade(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (transport.StreamConn, error) {
	_ = "STUB: not implemented"
	return *new(transport.StreamConn), nil
}
