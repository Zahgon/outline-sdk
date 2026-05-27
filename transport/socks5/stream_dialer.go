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

package socks5

import (
	"context"
	"io"

	"golang.getoutline.org/sdk/transport"
)

// https://datatracker.ietf.org/doc/html/rfc1929
// Credentials can be nil, and that means no authentication.
type credentials struct {
	username []byte
	password []byte
}

// NewClient creates a SOCKS5 client that routes connections to a SOCKS5
// proxy listening at the given [transport.StreamEndpoint].
func NewClient(streamEndpoint transport.StreamEndpoint) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Client struct {
	se   transport.StreamEndpoint
	pd   transport.PacketDialer
	cred *credentials
}

var _ transport.StreamDialer = (*Client)(nil)
var _ transport.PacketListener = (*Client)(nil)

func (c *Client) SetCredentials(username, password []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// EnablePacket enables the use of the [Client] as a [transport.PacketListener]. It takes the [transport.PacketDialer] used to connect to the SOCKS5 packet endpoint.
func (c *Client) EnablePacket(packetDialer transport.PacketDialer) {
	_ = "STUB: not implemented"
	return

	// request sends a SOCKS5 request to the server to perform a command (e.g., connect, udp associate),
	// performs authentication (if provided), returns the bound address.
}

func (c *Client) request(conn io.ReadWriter, cmd byte, dstAddr string) (*address, error) {
	_ = "STUB: not implemented"
	// For protocol details, see https://datatracker.ietf.org/doc/html/rfc1928#section-3
	// Creating a single buffer for method selection, authentication, and connection request
	// Buffer large enough for method, auth, and connect requests with a domain name address.
	// The maximum buffer size is:
	// 3 (1 socks version + 1 method selection + 1 methods)
	// + 1 (auth version) + 1 (username length) + 255 (username) + 1 (password length) + 255 (password)
	// + 256 (max domain name length)
	return nil, nil
}

// Method selection part: VER = 5, NMETHODS = 1, METHODS = 0 (no auth)
// +----+----------+----------+
// |VER | NMETHODS | METHODS  |
// +----+----------+----------+
// | 1  |    1     | 1 to 255 |
// +----+----------+----------+

// https://datatracker.ietf.org/doc/html/rfc1929
// Method selection part: VER = 5, NMETHODS = 1, METHODS = 2 (username/password)

// Authentication part: VER = 1, ULEN = 1, UNAME = 1~255, PLEN = 1, PASSWD = 1~255
// +----+------+----------+------+----------+
// |VER | ULEN |  UNAME   | PLEN |  PASSWD  |
// +----+------+----------+------+----------+
// | 1  |  1   | 1 to 255 |  1   | 1 to 255 |
// +----+------+----------+------+----------+

// CMD Request:
// VER = 5, CMD = cmd, RSV = 0, DST.ADDR, DST.PORT
// +----+-----+-------+------+----------+----------+
// |VER | CMD |  RSV  | ATYP | DST.ADDR | DST.PORT |
// +----+-----+-------+------+----------+----------+
// | 1  |  1  | X'00' |  1   | Variable |    2     |
// +----+-----+-------+------+----------+----------+

// TODO: Probably more memory efficient if remoteAddr is added to the buffer directly.

// We merge the method and CMD requests and only perform one write
// because we send a single authentication method, so there's no point
// in waiting for the response. This eliminates a roundtrip.

// Reading the response:
// 1. Read method response (VER, METHOD).
// +----+--------+
// |VER | METHOD |
// +----+--------+
// | 1  |   1    |
// +----+--------+
// buffer[0]: VER, buffer[1]: METHOD
// Reuse buffer for better performance.

// No authentication required.

// 2. Read authentication version and status
// VER = 1, STATUS = 0
// +----+--------+
// |VER | STATUS |
// +----+--------+
// | 1  |   1    |
// +----+--------+
// VER = 1 means the server should be expecting username/password authentication.
// buffer[2]: VER, buffer[3]: STATUS

// 3. Read connect response (VER, REP, RSV, ATYP, BND.ADDR, BND.PORT).
// See https://datatracker.ietf.org/doc/html/rfc1928#section-6.
// +----+-----+-------+------+----------+----------+
// |VER | REP |  RSV  | ATYP | BND.ADDR | BND.PORT |
// +----+-----+-------+------+----------+----------+
// | 1  |  1  | X'00' |  1   | Variable |    2     |
// +----+-----+-------+------+----------+----------+
// buffer[0]: VER
// buffer[1]: REP
// buffer[2]: RSV
// buffer[3]: ATYP

// if REP is not 0, it means the server returned an error.

// 4. Read BND.ADDR.

// connectAndRequest manages the connection lifecycle and delegates the SOCKS5 communication to the request function.
func (c *Client) connectAndRequest(ctx context.Context, cmd byte, dstAddr string) (transport.StreamConn, *address, error) {
	_ = "STUB: not implemented"
	return *new(transport.StreamConn), nil, nil
}

// DialStream implements [transport.StreamDialer].DialStream using SOCKS5.
// It will send the auth method, auth credentials (if auth is chosen), and
// the connect requests in one packet, to avoid an additional roundtrip.
// The returned [error] will be of type [ReplyCode] if the server sends a SOCKS error reply code, which
// you can check against the error constants in this package using [errors.Is].
func (c *Client) DialStream(ctx context.Context, dstAddr string) (transport.StreamConn, error) {
	_ = "STUB: not implemented"
	return *new(transport.StreamConn), nil
}
