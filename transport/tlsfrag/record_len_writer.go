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

package tlsfrag

import (
	"io"
)

// RecordLenFragFunc takes the length of the first [handshake record]'s content (without the 5-byte header),
// and returns an integer that determines where the record should be fragmented.
//
// The returned splitLen should be in range 1 to recordLen-1.
// The record content will then be fragmented into two parts: record[:splitLen] and record[splitLen:].
// If splitLen is either ≤ 0 or ≥ recordLen, no fragmentation will occur.
//
// [handshake record]: https://datatracker.ietf.org/doc/html/rfc8446#section-5.1
type RecordLenFragFunc func(recordLen int) (splitLen int)

// recordLenFragWriter splits the initial TLS Client Hello record into two TLS records based on a fixed length
// returned by a [RecordLenFragFunc] callback.
// These fragmented records are then written to the base [io.Writer]. Subsequent packets are not modified and are
// directly transmitted through the base [io.Writer].
type recordLenFragWriter struct {
	base   io.Writer
	frag   RecordLenFragFunc
	done   bool                     // the first fragmented header and payload are flushed (or don't split)
	hdr    []byte                   // the raw 5 bytes header, use tlsHdr to update PayloadLen
	tlsHdr tlsHandshakeRecordHeader // non-nil if hdr is a valid TLS Handshake record header

	// the records' sizes and written bytes (including 5 bytes header)
	r1Size, r1Written, r2Size int
}

var _ io.Writer = (*recordLenFragWriter)(nil)

// NewRecordLenFuncWriter creates a [io.Writer] that splits the first TLS Client Hello record into two records
// based on the provided [RecordLenFragFunc] callback.
// It then writes these records and all subsequent messages to the base [io.Writer].
// If the first message isn't a Client Hello, no splitting occurs and all messages are written directly to base.
//
// The returned [io.Writer] will implement the [io.ReaderFrom] interface for optimized performance if the base
// [io.Writer] implements [io.ReaderFrom].
func NewRecordLenFuncWriter(base io.Writer, frag RecordLenFragFunc) (io.Writer, error) {
	_ = "STUB: not implemented"
	return *new(io.Writer), nil
}

// Write implements io.Writer.Write. It attempts to split the data received in the first one or more Write call(s)
// into two TLS records if the data corresponds to a TLS Client Hello record without using any additional buffers.
func (w *recordLenFragWriter) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// try to fill r.w.hdr of totally 5 bytes

// construct the structured TLS Handshake header object

// invalid TLS header, or invalid split lens, stop splitting

// update header to be the first record's header

// update w.tlsHdr (aliases w.hdr) to the second record's header

// Internally writeBoth might copy p to a temporary buffer, if p is too big
// This is wasting CPU and memory, so we limit the maximum buffer size to be 16K
// which would be way more enough than a single TLS Client Hello record.

// fixedLenReaderFrom optimizes for fixedLenWriter when the base [io.Writer] implements [io.ReaderFrom].
type fixedLenReaderFrom struct {
	*recordLenFragWriter
	baseRF io.ReaderFrom
}

var _ io.ReaderFrom = (*fixedLenReaderFrom)(nil)

// fixedLenFirstRecordReader reads 5 bytes from r into w.hdr (aliased by w.tlsHdr), and calculate the split length.
// It then copies up to (w.r1Size - w.r1Written) bytes from (w.tlsHdr & r) into Read's buffer.
// It will update w.r1Written and w.done accordingly.
// After the first record is fully copied, it will set w.tlsHdr (aliases w.hdr) to be the second record's header.
type fixedLenFirstRecordReader struct {
	w        *recordLenFragWriter
	r        io.Reader
	rReadLen int64
}

// fixedLenRemainingRecordReader flushes the content of w.hdr and all remaining r into Read's buffer.
type fixedLenRemainingRecordReader struct {
	w        *recordLenFragWriter
	r        io.Reader
	rReadLen int64
}

var _ io.Reader = (*fixedLenFirstRecordReader)(nil)
var _ io.Reader = (*fixedLenRemainingRecordReader)(nil)

func (r *fixedLenFirstRecordReader) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// try to fill r.w.hdr of totally 5 bytes

// construct the structured TLS Handshake header object

// invalid TLS header, or invalid split lens, stop splitting

// update header to be the first record's header

// update r.w.tlsHdr (aliases r.w.hdr) to the second record's header

func (r *fixedLenRemainingRecordReader) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ReadFrom implements io.ReaderFrom.ReadFrom. It attempts to split the first packet into two TLS records if the data
// corresponds to a TLS Client Hello record without using any additional buffers.
// And then copies the remaining data from r to the base io.Writer until EOF or error.
//
// If the first packet is not a valid TLS Client Hello, everything from r gets copied to the base io.Writer as is.
//
// It returns the number of bytes read. Any error except EOF encountered during the read is also returned.
//
// ReadFrom will hang indefinitely if r provides fewer than 5 bytes and doesn't return the io.EOF error (e.g., "PING").
func (w *fixedLenReaderFrom) ReadFrom(r io.Reader) (n int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// We should return the actual bytes read from r, not the bytes passed to base or baseRF

// updateSplitLen determines the split length by calling w.frag with the input of w.tlsHdr.PayloadLen().
// It returns nil error if w.frag returns a valid split length, otherwise it returns non-nil error.
//
// The corresponding record lengths (including the 5 bytes header) will be set to w.r1Size and w.r2Size.
func (w *recordLenFragWriter) updateSplitRecordLens() error { _ = "STUB: not implemented"; return nil }

// writeN writes at most limit bytes from p to dst.
func writeN(dst io.Writer, p []byte, limit int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// writeBoth writes both p1 and p2 to dst in a single Write or writev call.
// It returns the number of bytes that are written from p1 and p2, respectively.
//
// Issuing a single Write or writev call to dst is required because otherwise dst
// will receive two TCP packets, which introduces unwanted TCP split.
//
// Performance note, internally we might allocate a temporary buffer and copy the
// data from p1 and p2 to that buffer, please be careful about the data size.
func writeBoth(dst io.Writer, p1 []byte, p2 []byte) (int, int, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// If the underlying writer implements writev system call
// UDPConn and IPConn also implement writev, but TLS is TCP so we only care about TCP

// We must allocate temporary buffer to hold both content and issue a single Write.
// This will add some pressure to GC because the temporary buffer will escape to heap.
// Go's proposal of memory arena can be a remedy, but the proposal is on hold indefinitely.

// writeBothN writes at most limit bytes from p1 and p2 to dst in a single Write orwritev call.
// It returns the number of bytes that are written from p1 and p2, respectively.
func writeBothN(dst io.Writer, p1 []byte, p2 []byte, limit int) (int, int, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}
