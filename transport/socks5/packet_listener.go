// Copyright 2024 The Outline Authors
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

package socks5

import (
	"context"
	"io"
	"net"
	"time"

	"golang.getoutline.org/sdk/internal/slicepool"
)

// clientUDPBufferSize is the maximum supported UDP packet size in bytes.
const clientUDPBufferSize = 16 * 1024

// udpPool stores the byte slices used for storing packets.
var udpPool = slicepool.MakePool(clientUDPBufferSize)

type packetConn struct {
	pc net.Conn
	sc io.Closer
}

var _ net.PacketConn = (*packetConn)(nil)

func (p *packetConn) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (p *packetConn) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (p *packetConn) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (c *packetConn) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// ReadFrom reads the packet from the SOCKS5 server and extract the payload
// The packet format is specified in https://datatracker.ietf.org/doc/html/rfc1928#section-7
func (p *packetConn) ReadFrom(b []byte) (int, net.Addr, error) {
	_ = "STUB: not implemented"
	return 0, *new(net.Addr), nil
}

// Minimum packet size

// Using bytes.Buffer to handle data

// Read and check reserved bytes

// Read fragment byte

// Read address using socks.ReadAddr which must now accept a bytes.Buffer directly

// Convert the address to a net.Addr

// Payload handling: remaining bytes in the buffer are the payload

// WriteTo encapsulates the payload in a SOCKS5 UDP packet as specified in
// https://datatracker.ietf.org/doc/html/rfc1928#section-7
// and write it to the SOCKS5 server via the underlying connection.
func (p *packetConn) WriteTo(b []byte, addr net.Addr) (int, error) {
	_ = "STUB: not implemented"

	// The minimum preallocated header size (10 bytes)
	return 0, nil
}

// Reserved
// Fragment number
// To be appended below:
// ATYP, IPv4, IPv6, Domain Name, Port

// Combine the header and the payload

// Close closes both the underlying stream and packet connections.
func (p *packetConn) Close() error { _ = "STUB: not implemented"; return nil }

// ListenPacket creates a [net.PacketConn] for UDP communication via the SOCKS5 server.
func (c *Client) ListenPacket(ctx context.Context) (net.PacketConn, error) {
	_ = "STUB: not implemented"
	// Connect to the SOCKS5 server and perform UDP association
	// Since local address is not known in advance, we use unspecified address
	// which means the server is going to accept incoming packets from any address
	// on the bind port on the server. The bind address is determined and returned by
	// the server.
	// https://datatracker.ietf.org/doc/html/rfc1928#section-6
	// Whoile binding address to specific client address has its advantages, it also creates some
	// challenges such as NAT traveral if client is behind NAT.
	return *new(net.PacketConn), nil
}

// If the returned bind IP address is unspecified (i.e. "0.0.0.0" or "::"),
// then use the IP address of the SOCKS5 server
