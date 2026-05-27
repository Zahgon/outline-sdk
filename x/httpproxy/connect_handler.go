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

package httpproxy

import (
	"context"
	"io"
	"net/http"

	"golang.getoutline.org/sdk/transport"
)

type sanitizeErrorDialer struct {
	transport.StreamDialer
}

func isCancelledError(err error) bool { _ = "STUB: not implemented"; return false }

// Works around the fact that DNS doesn't return typed errors.

func (d *sanitizeErrorDialer) DialStream(ctx context.Context, addr string) (transport.StreamConn, error) {
	_ = "STUB: not implemented"
	return *new(transport.StreamConn), nil
}

// StreamDialerParser creates a [transport.StreamDialer] from a config string.
// It is used by [NewConnectHandler] to support the Transport request header.
type StreamDialerParser func(ctx context.Context, config string) (transport.StreamDialer, error)

// HandlerOption configures a connect handler.
type HandlerOption func(*connectHandler)

// WithStreamDialerParser sets a factory that creates a dialer from the Transport request header value.
// When set, clients can override the transport per-request by sending a Transport header.
func WithStreamDialerParser(f StreamDialerParser) HandlerOption {
	_ = "STUB: not implemented"
	return *new(HandlerOption)
}

type connectHandler struct {
	dialer        *sanitizeErrorDialer
	dialerFactory StreamDialerParser
}

var _ http.Handler = (*connectHandler)(nil)

func (h *connectHandler) ServeHTTP(proxyResp http.ResponseWriter, proxyReq *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Validate the target address.

// As per https://httpwg.org/specs/rfc9110.html#CONNECT.

// Dial the target, optionally using a per-request transport from the Transport header.

// Because we sanitize the base dialer error, it's safe to return error details here.

// Set up protocol-specific client I/O. H1 hijacks the raw connection; H2/H3 stream
// through the ResponseWriter with explicit flushing after each write.

// H1: hijack the raw connection and relay using the underlying bufio.ReadWriter.

// TODO(fortuna): Use context.AfterFunc after we migrate to Go 1.21.

// We close the hijacked connection when the context is done. This way
// we allow the HTTP server to control the request lifetime.
// The request context will be cancelled right after ServeHTTP returns,
// but it can be cancelled before, if the server uses a custom BaseContext.

// clientRW (bufio.ReadWriter) implements io.ReaderFrom via its embedded bufio.Writer.

// afterCopy flushes the bufio buffer to push any remaining bytes to the client.

// H2/H3: hijacking is not available on multiplexed connections.

// flushingWriter flushes after every write, so no afterCopy flush is needed.

// Relay data between client and target in both directions.

// io.Copy prefers WriteTo, which clientRW implements. However,
// bufio.ReadWriter.WriteTo issues an empty Write() call, which flushes
// the Shadowsocks IV and connect request, breaking the coalescing with
// the initial data. By preferring ReaderFrom, the coalescing of IV,
// request and initial data is preserved.

// We can't use io.Copy here because it doesn't call Flush on writes, so the first
// write is never sent and the entire relay gets stuck. bufio.Writer.ReadFrom (H1)
// and flushingWriter.ReadFrom (H2/H3) take care of that.

// flushingWriter wraps an http.ResponseWriter and flushes after every write,
// ensuring bytes are sent to the client immediately over H2/H3 streams.
type flushingWriter struct {
	w http.ResponseWriter
	f http.Flusher
}

func (fw *flushingWriter) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// ReadFrom shadows http.ResponseWriter's own ReadFrom (present in net/http's *response),
// which does not flush. This implementation flushes after every write so bytes reach
// the client immediately, and prefers r.WriteTo to avoid an intermediate buffer.
func (fw *flushingWriter) ReadFrom(r io.Reader) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// NewConnectHandler creates a [http.Handler] that handles CONNECT requests and forwards
// the requests using the given [transport.StreamDialer].
//
// Use [WithStreamDialerParser] to support the Transport request header, which allows clients
// to specify a per-request transport config.
//
// The resulting handler is currently vulnerable to probing attacks. It's ok as a localhost proxy
// but it may be vulnerable if used as a public proxy.
func NewConnectHandler(dialer transport.StreamDialer, opts ...HandlerOption) http.Handler {
	_ = "STUB: not implemented"
	// We sanitize the errors from the input Dialer because we don't want to leak sensitive details
	// of the base dialer (e.g. access key credentials) to the user.
	return *new(http.Handler)
}
