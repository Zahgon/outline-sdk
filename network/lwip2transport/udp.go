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

package lwip2transport

import (
	"net"
	"sync"
	"sync/atomic"

	lwip "github.com/eycorsican/go-tun2socks/core"
	"golang.getoutline.org/sdk/network"
)

// Compilation guard against interface implementation
var _ lwip.UDPConnHandler = (*udpHandler)(nil)
var _ network.PacketResponseReceiver = (*udpConnResponseWriter)(nil)

type udpHandler struct {
	mu      sync.Mutex                             // Protects the senders field
	proxy   network.PacketProxy                    // A network stack neutral implementation of UDP PacketProxy
	senders map[string]network.PacketRequestSender // Maps local lwIP UDP socket to PacketRequestSender
}

// newUDPHandler returns a lwIP UDP connection handler.
//
// `pktProxy` is a PacketProxy that handles UDP packets.
func newUDPHandler(pktProxy network.PacketProxy) *udpHandler { _ = "STUB: not implemented"; return nil }

// Connect does nothing. New UDP sessions will be created in ReceiveTo.
func (h *udpHandler) Connect(tunConn lwip.UDPConn, _ *net.UDPAddr) error {
	_ = "STUB: not implemented"

	// ReceiveTo relays packets from the lwIP TUN device to the proxy. It's called by lwIP. ReceiveTo will also create a
	// new UDP session if `data` is the first packet from the `tunConn`.
	return nil
}

func (h *udpHandler) ReceiveTo(tunConn lwip.UDPConn, data []byte, destAddr *net.UDPAddr) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// TODO: address synchronization issues with h.senders[laddr] and the following `if` in the future

// newSession creates a new PacketRequestSender related to conn. The caller needs to put the new PacketRequestSender
// to the h.senders map.
func (h *udpHandler) newSession(conn lwip.UDPConn) (network.PacketRequestSender, error) {
	_ = "STUB: not implemented"
	return *new(network.PacketRequestSender), nil
}

// closeSession cleans up resources related to conn.
func (h *udpHandler) closeSession(conn lwip.UDPConn) error { _ = "STUB: not implemented"; return nil }

// The PacketResponseWriter that will write responses to the lwip network stack.
type udpConnResponseWriter struct {
	closed atomic.Bool
	conn   lwip.UDPConn
	h      *udpHandler
}

// Write relays packets from the proxy to the lwIP TUN device.
func (r *udpConnResponseWriter) WriteFrom(p []byte, source net.Addr) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// net.Addr -> *net.UDPAddr, because r.conn.WriteFrom requires *net.UDPAddr
// and this is more reliable than type assertion
// also the source address host will be an IP address, no actual resolution will be done

// Close informs the udpHandler to close the UDPConn and clean up the UDP session.
func (r *udpConnResponseWriter) Close() error { _ = "STUB: not implemented"; return nil }
