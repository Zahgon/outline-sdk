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
	"io"
)

// clientHelloFragWriter intercepts the initial TLS Client Hello record and splits it into two TLS records based on the
// return value of frag function. These fragmented records are then written to the base [io.Writer]. Subsequent packets
// are not modified and are directly transmitted through the base [io.Writer].
type clientHelloFragWriter struct {
	base io.Writer
	// Indicates all splitted rcds have been already written to base
	done bool
	frag FragFunc

	// The buffer containing and parsing a TLS Client Hello record
	helloBuf *clientHelloBuffer
	// The buffer containing splitted records what will be written to base
	record *bytes.Buffer
}

// clientHelloFragReaderFrom serves as an optimized version of clientHelloFragWriter when the base [io.Writer] also
// implements the [io.ReaderFrom] interface.
type clientHelloFragReaderFrom struct {
	*clientHelloFragWriter
	baseRF io.ReaderFrom
}

// Compilation guard against interface implementation
var _ io.Writer = (*clientHelloFragWriter)(nil)
var _ io.Writer = (*clientHelloFragReaderFrom)(nil)
var _ io.ReaderFrom = (*clientHelloFragReaderFrom)(nil)

// newClientHelloFragWriter creates a [io.Writer] that splits the first TLS Client Hello record into two records
// based on the provided [FragFunc] callback.
// It then writes these records and all subsequent messages to the base [io.Writer].
// If the first message isn't a Client Hello, no splitting occurs and all messages are written directly to base.
//
// The returned [io.Writer] will implement the [io.ReaderFrom] interface for optimized performance if the base
// [io.Writer] implements [io.ReaderFrom].
//
// If you just want to split the record at a fixed position (e.g., always at the 5th byte or 2nd from the last
// byte), use [NewRecordLenFuncWriter]. It consumes less resources and is more efficient.
func newClientHelloFragWriter(base io.Writer, frag FragFunc) (io.Writer, error) {
	_ = "STUB: not implemented"
	return *new(io.Writer), nil
}

// Write implements io.Writer.Write. It attempts to split the data received in the first one or more Write call(s)
// into two TLS records if the data corresponds to a TLS Client Hello record.
func (w *clientHelloFragWriter) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"

	// not yet splitted, append to the buffer
	return 0, nil
}

// all written, but Client Hello is not fully received yet

// already splitted (but previous Writes might fail), try to flush all remaining w.record to w.base

// ReadFrom implements io.ReaderFrom.ReadFrom. It attempts to split the first packet into two TLS records if the data
// corresponds to a TLS Client Hello record. And then copies the remaining data from r to the base io.Writer until EOF
// or error.
//
// If the first packet is not a valid TLS Client Hello, everything from r gets copied to the base io.Writer as is.
//
// It returns the number of bytes read. Any error except EOF encountered during the read is also returned.
//
// ReadFrom will hang indefinitely if r provides fewer than 5 bytes and doesn't return the io.EOF error (e.g., "PING").
func (w *clientHelloFragReaderFrom) ReadFrom(r io.Reader) (n int64, err error) {
	_ = "STUB: not implemented"

	// not yet splitted, append to the buffer
	return 0, nil
}

// EOF, but Client Hello is not fully received yet

// already splitted (but previous Writes might fail), try to flush all remaining w.record to w.base

// copyHelloBufToRecord copies w.helloBuf into w.record without allocations.
func (w *clientHelloFragWriter) copyHelloBufToRecord() { _ = "STUB: not implemented"; return }

// allows the GC to recycle the memory

// splitHelloBufToRecord splits w.helloBuf into two records and put them into w.record without allocations.
func (w *clientHelloFragWriter) splitHelloBufToRecord() { _ = "STUB: not implemented"; return }

//           |  header   |         payload         |  cap==len+5
// original: | <= (5) => | <= head => | <= tail => | <= (5) => |
//                       |            |\            \
//                       |            | \-------\    \-------\
//                       |            |          \            \
// splitted: | <= (5) => | <= head => | <= (5) => | <= tail => |
//           |  header1  |  payload1  |  header2  |  payload2  |

// Shift tail fragment to make space for record header.

// Insert header for second fragment.

// allows the GC to recycle the memory

// flushRecord writes all bytes from w.record to base.
func (w *clientHelloFragWriter) flushRecord() (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// allows the GC to recycle the memory
