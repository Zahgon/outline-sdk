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

package network

import (
	"net"
	"net/netip"
	"sync"
	"time"

	"golang.getoutline.org/sdk/internal/slicepool"
	"golang.getoutline.org/sdk/transport"
)

// this was the buffer size used before, we may consider update it in the future
const packetMaxSize = 2048

// packetBufferPool is used to create buffers to read UDP response packets
var packetBufferPool = slicepool.MakePool(packetMaxSize)

// Compilation guard against interface implementation
var _ PacketProxy = (*PacketListenerProxy)(nil)
var _ PacketRequestSender = (*packetListenerRequestSender)(nil)

type PacketListenerProxy struct {
	listener         transport.PacketListener
	writeIdleTimeout time.Duration
}

type packetListenerRequestSender struct {
	mu     sync.Mutex // Protects closed and timer function calls
	closed bool

	proxyConn        net.PacketConn
	writeIdleTimeout time.Duration
	writeIdleTimer   *time.Timer
}

// NewPacketProxyFromPacketListener creates a new [PacketProxy] that uses the existing [transport.PacketListener] to
// create connections to a proxy. You can also specify additional options.
// This function is useful if you already have an implementation of [transport.PacketListener] and you want to use it
// with one of the network stacks (for example, network/lwip2transport) as a UDP traffic handler.
func NewPacketProxyFromPacketListener(pl transport.PacketListener, options ...func(*PacketListenerProxy) error) (*PacketListenerProxy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WithPacketListenerWriteIdleTimeout sets the write idle timeout of the [PacketListenerProxy].
// This means that if there are no WriteTo operations on the UDP session created by NewSession for the specified amount
// of time, the proxy will end this session.
//
// This should be used together with the [NewPacketProxyFromPacketListenerWithOptions] function.
func WithPacketListenerWriteIdleTimeout(timeout time.Duration) func(*PacketListenerProxy) error {
	_ = "STUB: not implemented"
	return nil
}

// NewSession implements [PacketProxy].NewSession function. It uses [transport.PacketListener].ListenPacket to create
// a [net.PacketConn], and constructs a new [PacketRequestSender] that is based on this [net.PacketConn].
func (proxy *PacketListenerProxy) NewSession(respWriter PacketResponseReceiver) (PacketRequestSender, error) {
	_ = "STUB: not implemented"
	return *new(PacketRequestSender), nil
}

// Terminate the session after timeout with no outgoing writes (deadline is refreshed by WriteTo)

// Relay incoming UDP responses from the proxy asynchronously until EOF, session expiration or error

// Allocate buffer from slicepool, because `go build -gcflags="-m"` shows a local array will escape to heap

// Ignore some specific recoverable errors

// WriteTo implements [PacketRequestSender].WriteTo function. It simply forwards the packet to the underlying
// [net.PacketConn].WriteTo function.
func (s *packetListenerRequestSender) WriteTo(p []byte, destination netip.AddrPort) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Close implements [PacketRequestSender].Close function. It closes the underlying [net.PacketConn]. This will also
// terminate the goroutine created in NewSession because s.conn.ReadFrom will return [io.EOF].
func (s *packetListenerRequestSender) Close() error { _ = "STUB: not implemented"; return nil }

// resetWriteIdleTimer extends the writeIdleTimer's timeout to now() + writeIdleTimeout. If `s` is closed, it will
// return ErrClosed.
func (s *packetListenerRequestSender) resetWriteIdleTimer() error {
	_ = "STUB: not implemented"
	return nil
}
