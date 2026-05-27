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
	"io"
	"net"
	"time"

	"golang.getoutline.org/sdk/transport"
)

var _ transport.StreamConn = (*pipeConn)(nil)

type pipeConn struct {
	reader     readCloseDeadliner
	writer     writeCloseDeadliner
	remoteAddr net.Addr
}

type readCloseDeadliner interface {
	io.ReadCloser
	SetDeadline(deadline time.Time) error
}

type writeCloseDeadliner interface {
	io.WriteCloser
	SetDeadline(deadline time.Time) error
}

func newPipeConn(writer writeCloseDeadliner, reader readCloseDeadliner, remoteAddr net.Addr) *pipeConn {
	_ = "STUB: not implemented"
	return nil
}

func (p *pipeConn) Read(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (p *pipeConn) Write(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (p *pipeConn) CloseRead() error { _ = "STUB: not implemented"; return nil }

func (p *pipeConn) CloseWrite() error { _ = "STUB: not implemented"; return nil }

func (p *pipeConn) Close() error { _ = "STUB: not implemented"; return nil }

func (p *pipeConn) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (p *pipeConn) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (p *pipeConn) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (p *pipeConn) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (p *pipeConn) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }
