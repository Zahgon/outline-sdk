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

package disorder

import (
	"io"

	"golang.getoutline.org/sdk/x/sockopt"
)

type disorderWriter struct {
	conn             io.Writer
	tcpOptions       sockopt.TCPOptions
	writesToDisorder int
}

var _ io.Writer = (*disorderWriter)(nil)

func NewWriter(conn io.Writer, tcpOptions sockopt.TCPOptions, runAtPacketN int) io.Writer {
	_ = "STUB: not implemented"
	// TODO: Support ReadFrom.
	return *new(io.Writer)
}

func (w *disorderWriter) Write(data []byte) (written int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Setting number of hops to 1 will lead to data to get lost on host.

// The packet with low hop limit was sent.
// Make next calls send data normally.
//
// The packet with the low hop limit will get resent by the kernel later.
// The network filters will receive data out of order.

// The packet will get lost at the first send, since the hop limit is too low.

// TODO: Wait for queued data to be sent by the kernel to the socket.
