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

package dns

import (
	"context"
	"errors"
	"io"

	"golang.getoutline.org/sdk/transport"
	"golang.org/x/net/dns/dnsmessage"
)

var (
	ErrBadRequest  = errors.New("request input is invalid")
	ErrDial        = errors.New("dial DNS resolver failed")
	ErrSend        = errors.New("send DNS message failed")
	ErrReceive     = errors.New("receive DNS message failed")
	ErrBadResponse = errors.New("response message is invalid")
)

// nestedError allows us to use errors.Is and still preserve the error cause.
// This is unlike fmt.Errorf, which creates a new error and preserves the cause,
// but you can't specify the type of the resulting top-level error.
type nestedError struct {
	is      error
	wrapped error
}

func (e *nestedError) Is(target error) bool { _ = "STUB: not implemented"; return false }

func (e *nestedError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e *nestedError) Error() string { _ = "STUB: not implemented"; return "" }

// Resolver can query the DNS with a question, and obtain a DNS message as response.
// This abstraction helps hide the underlying transport protocol.
type Resolver interface {
	Query(ctx context.Context, q dnsmessage.Question) (*dnsmessage.Message, error)
}

// FuncResolver is a [Resolver] that uses the given function to query DNS.
type FuncResolver func(ctx context.Context, q dnsmessage.Question) (*dnsmessage.Message, error)

// Query implements the [Resolver] interface.
func (f FuncResolver) Query(ctx context.Context, q dnsmessage.Question) (*dnsmessage.Message, error) {
	_ = "STUB: not implemented"

	// NewQuestion is a convenience function to create a [dnsmessage.Question].
	// The input domain is interpreted as fully-qualified. If the end "." is missing, it's added.
	return nil, nil
}

func NewQuestion(domain string, qtype dnsmessage.Type) (*dnsmessage.Question, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Maximum UDP message size that we support.
// The value is taken from https://dnsflagday.net/2020/, which says:
// "An EDNS buffer size of 1232 bytes will avoid fragmentation on nearly all current networks.
// This is based on an MTU of 1280, which is required by the IPv6 specification, minus 48 bytes
// for the IPv6 and UDP headers".
const maxUDPMessageSize = 1232

// appendRequest appends the bytes a DNS request using the id and question to buf.
func appendRequest(id uint16, q dnsmessage.Question, buf []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Set the maximum payload size we support, as per https://datatracker.ietf.org/doc/html/rfc6891#section-4.3

// Fold case as clarified in https://datatracker.ietf.org/doc/html/rfc4343#section-3.
func foldCase(char byte) byte { _ = "STUB: not implemented"; return 0 }

// equalASCIIName compares DNS name as specified in https://datatracker.ietf.org/doc/html/rfc1035#section-3.1 and
// https://datatracker.ietf.org/doc/html/rfc4343#section-3.
func equalASCIIName(x, y dnsmessage.Name) bool { _ = "STUB: not implemented"; return false }

func checkResponse(reqID uint16, reqQues dnsmessage.Question, respHdr dnsmessage.Header, respQs []dnsmessage.Question) error {
	_ = "STUB: not implemented"
	return nil
}

// https://datatracker.ietf.org/doc/html/rfc5452#section-4.3

// https://datatracker.ietf.org/doc/html/rfc5452#section-4.2

// queryDatagram implements a DNS query over a datagram protocol.
func queryDatagram(conn io.ReadWriter, q dnsmessage.Question) (*dnsmessage.Message, error) {
	_ = "STUB: not implemented"
	// Reference: https://cs.opensource.google/go/go/+/master:src/net/dnsclient_unix.go?q=func:dnsPacketRoundTrip&ss=go%2Fgo
	return nil, nil
}

// Handle bad io.Reader.

// Ignore invalid packets that fail to parse. It could be injected.

// queryStream implements a DNS query over a stream protocol. It frames the messages by prepending them with a 2-byte length prefix.
func queryStream(conn io.ReadWriter, q dnsmessage.Question) (*dnsmessage.Message, error) {
	_ = "STUB: not implemented"
	// Reference: https://cs.opensource.google/go/go/+/master:src/net/dnsclient_unix.go?q=func:dnsStreamRoundTrip&ss=go%2Fgo
	return nil, nil
}

// Buffer length must fit in a uint16.

// TODO: Consider writer.ReadFrom(net.Buffers) in case the writer is a TCPConn.

func ensurePort(address string, defaultPort string) string { _ = "STUB: not implemented"; return "" }

// Failed to parse as host:port. Assume address is a host.

// NewUDPResolver creates a [Resolver] that implements the DNS-over-UDP protocol, using a [transport.PacketDialer] for transport.
// It uses a different port for every request.
//
// [DNS-over-UDP]: https://datatracker.ietf.org/doc/html/rfc1035#section-4.2.1
func NewUDPResolver(pd transport.PacketDialer, resolverAddr string) Resolver {
	_ = "STUB: not implemented"
	return *new(Resolver)
}

type streamResolver struct {
	NewConn func(context.Context) (transport.StreamConn, error)
}

func (r *streamResolver) Query(ctx context.Context, q dnsmessage.Question) (*dnsmessage.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: reuse connection, as per https://datatracker.ietf.org/doc/html/rfc7766#section-6.2.1.

// NewTCPResolver creates a [Resolver] that implements the [DNS-over-TCP] protocol, using a [transport.StreamDialer] for transport.
// It creates a new connection to the resolver for every request.
//
// [DNS-over-TCP]: https://datatracker.ietf.org/doc/html/rfc1035#section-4.2.2
func NewTCPResolver(sd transport.StreamDialer, resolverAddr string) Resolver {
	_ = "STUB: not implemented"
	// TODO: Consider handling Authenticated Data.
	return *new(Resolver)
}

// NewTLSResolver creates a [Resolver] that implements the [DNS-over-TLS] protocol, using a [transport.StreamDialer]
// to connect to the resolverAddr, and the resolverName as the TLS server name.
// It creates a new connection to the resolver for every request.
//
// [DNS-over-TLS]: https://datatracker.ietf.org/doc/html/rfc7858
func NewTLSResolver(sd transport.StreamDialer, resolverAddr string, resolverName string) Resolver {
	_ = "STUB: not implemented"
	return *new(Resolver)
}

// NewHTTPSResolver creates a [Resolver] that implements the [DNS-over-HTTPS] protocol, using a [transport.StreamDialer]
// to connect to the resolverAddr, and the url as the DoH template URI.
// It uses an internal HTTP client that reuses connections when possible.
//
// [DNS-over-HTTPS]: https://datatracker.ietf.org/doc/html/rfc8484
func NewHTTPSResolver(sd transport.StreamDialer, resolverAddr string, url string) Resolver {
	_ = "STUB: not implemented"
	return *new(Resolver)
}

// TODO: Support UDP for QUIC.

// TODO: add mechanism to close idle connections.
// Copied from Intra: https://github.com/Jigsaw-Code/Intra/blob/d3554846a1146ae695e28a8ed6dd07f0cd310c5a/Android/tun2socks/intra/doh/doh.go#L213-L219

// Same value as Android DNS-over-TLS

// Prepare request.

// Send request and get response.

// Process response.
