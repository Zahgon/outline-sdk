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

package shadowsocks

import (
	"context"
	"net"

	"golang.getoutline.org/sdk/internal/slicepool"
	"golang.getoutline.org/sdk/transport"
)

// clientUDPBufferSize is the maximum supported UDP packet size in bytes.
const clientUDPBufferSize = 16 * 1024

// udpPool stores the byte slices used for storing encrypted packets.
var udpPool = slicepool.MakePool(clientUDPBufferSize)

type packetListener struct {
	endpoint      transport.PacketEndpoint
	key           *EncryptionKey
	saltGenerator SaltGenerator
}

var _ transport.PacketListener = (*packetListener)(nil)

type PacketListener = *packetListener

// NewPacketListener creates a new Shadowsocks PacketListener that connects to the proxy on the given endpoint
// and uses the given key for encryption.
func NewPacketListener(endpoint transport.PacketEndpoint, key *EncryptionKey) (PacketListener, error) {
	_ = "STUB: not implemented"
	return *new(PacketListener), nil
}

// SetSaltGenerator sets the SaltGenerator to use for encryption. If not set, it used the [RandomSaltGenerator] by default.
func (pl *packetListener) SetSaltGenerator(sg SaltGenerator) { _ = "STUB: not implemented"; return }

// ListenPacket creates a net.PackeConn to send packets from the remote endpoint.
func (pl *packetListener) ListenPacket(ctx context.Context) (net.PacketConn, error) {
	_ = "STUB: not implemented"
	return *new(net.PacketConn), nil
}

type packetConn struct {
	net.Conn
	key           *EncryptionKey
	saltGenerator SaltGenerator
}

var _ net.PacketConn = (*packetConn)(nil)

// NewPacketConn wraps a [net.Conn] and returns a [net.PacketConn] that encrypts/decrypts
// packets before writing/reading them to/from the underlying connection using the provided
// encryption key.
//
// Closing the returned [net.PacketConn] will also close the underlying [net.Conn].
func NewPacketConn(conn net.Conn, key *EncryptionKey) net.PacketConn {
	_ = "STUB: not implemented"
	return *new(net.PacketConn)
}

// WriteTo encrypts `b` and writes to `addr` through the proxy.
func (c *packetConn) WriteTo(b []byte, addr net.Addr) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Copy the SOCKS target address and payload, reserving space for the generated salt to avoid
// partially overlapping the plaintext and cipher slices since `Pack` skips the salt when calling
// `AEAD.Seal` (see https://golang.org/pkg/crypto/cipher/#AEAD).

// ReadFrom reads from the embedded PacketConn and decrypts into `b`.
func (c *packetConn) ReadFrom(b []byte) (int, net.Addr, error) {
	_ = "STUB: not implemented"
	return 0, *new(net.Addr), nil
}

// Decrypt in-place.

// Strip the SOCKS source address
