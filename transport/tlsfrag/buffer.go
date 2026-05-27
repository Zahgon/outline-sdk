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

package tlsfrag

import (
	"bytes"
	"errors"
	"io"
)

var (
	// errTLSClientHelloFullyReceived is returned when a full TLS Client Hello has been received and no
	// more data can be pushed to the buffer.
	errTLSClientHelloFullyReceived = errors.New("already received a complete TLS Client Hello packet")
)

// clientHelloBuffer is a byte buffer used to receive and buffer a TLS Client Hello packet.
type clientHelloBuffer struct {
	// The buffer that hosts both header and content, cap: 5 -> 5+len(content)+padding
	data []byte
	// Indicates whether the content in data is a valid TLS Client Hello record
	validationErr error
	// A reader used to read from the slice passed to Write
	bufrd *bytes.Reader
}

var _ io.Writer = (*clientHelloBuffer)(nil)
var _ io.ReaderFrom = (*clientHelloBuffer)(nil)

// newClientHelloBuffer creates and initializes a new buffer to receive a TLS Client Hello packet.
func newClientHelloBuffer() *clientHelloBuffer {
	_ = "STUB: not implemented"
	// Allocate the 5 bytes header first, and then reallocate it to contain the entire packet later
	return nil
}

// It will be Reset in Write

// Bytes returns the full Client Hello packet including both the 5 bytes header and the content.
func (b *clientHelloBuffer) Bytes() []byte {
	_ = "STUB: not implemented"

	// Write appends p to the buffer and returns the number of bytes actually used.
	// If this data completes a valid TLS Client Hello, it returns errTLSClientHelloFullyReceived.
	// If an invalid TLS Client Hello message is detected, it returns the error errInvalidTLSClientHello.
	// If all bytes in p have been used and the buffer still requires more data to build a complete TLS Client Hello
	// message, it returns (len(p), nil).
	return nil
}

func (b *clientHelloBuffer) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// ReadFrom reads all the data from r and appends it to this buffer until a complete Client Hello packet has been
// received, or r returns EOF or error. It returns the number of bytes read. Any error except EOF encountered during
// the read is also returned.
//
// If this buffer completes a valid TLS Client Hello, it returns errTLSClientHelloFullyReceived.
// If an invalid TLS Client Hello message is detected, it returns the error errInvalidTLSClientHello.
// If this buffer still requires more data to build a complete TLS Client Hello message, it returns nil error.
//
// You can call ReadFrom multiple times if r doesn't provide enough data to build a complete Client Hello packet.
//
// ReadFrom will hang indefinitely if r provides fewer than 5 bytes and doesn't return the io.EOF error (e.g., "PING").
func (b *clientHelloBuffer) ReadFrom(r io.Reader) (n int64, err error) {
	_ = "STUB: not implemented"
	// Waiting to finish the header of 5 bytes
	return 0, nil
}

// If the buffer is already invalid

// Waiting to finish the payload of cap(b.data)-5 bytes
