// Copyright 2018 The Outline Authors
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
	"bytes"
	"crypto/cipher"
	"fmt"
	"io"
	"sync"

	"golang.getoutline.org/sdk/internal/slicepool"
)

// payloadSizeMask is the maximum size of payload in bytes, as per https://shadowsocks.org/guide/aead.html#tcp.
const payloadSizeMask = 0x3FFF // 16*1024 - 1

// Buffer pool used for decrypting Shadowsocks streams.
// The largest buffer we could need is for decrypting a max-length payload.
var readBufPool = slicepool.MakePool(payloadSizeMask + maxTagSize)

// Writer is an [io.Writer] that also implements [io.ReaderFrom] to
// allow for piping the data without extra allocations and copies.
// The LazyWrite and Flush methods allow a header to be
// added but delayed until the first write, for concatenation.
// All methods except Flush must be called from a single thread.
type Writer struct {
	// This type is single-threaded except when needFlush is true.
	// mu protects needFlush, and also protects everything
	// else while needFlush could be true.
	mu sync.Mutex
	// Indicates that a concurrent flush is currently allowed.
	needFlush     bool
	writer        io.Writer
	key           *EncryptionKey
	saltGenerator SaltGenerator
	// Wrapper for input that arrives as a slice.
	byteWrapper bytes.Reader
	// Number of plaintext bytes that are currently buffered.
	pending int
	// These are populated by init():
	buf  []byte
	aead cipher.AEAD
	// Index of the next encrypted chunk to write.
	counter []byte
}

var (
	_ io.Writer     = (*Writer)(nil)
	_ io.ReaderFrom = (*Writer)(nil)
)

// NewWriter creates a [Writer] that encrypts the given [io.Writer] using
// the shadowsocks protocol with the given encryption key.
func NewWriter(writer io.Writer, key *EncryptionKey) *Writer { _ = "STUB: not implemented"; return nil }

// SetSaltGenerator sets the salt generator to be used. Must be called before the first write.
func (sw *Writer) SetSaltGenerator(saltGenerator SaltGenerator) { _ = "STUB: not implemented"; return }

// init generates a random salt, sets up the AEAD object and writes
// the salt to the inner Writer.
func (sw *Writer) init() (err error) {
	if sw.aead == nil {
		salt := make([]byte, sw.key.SaltSize())
		if err := sw.saltGenerator.GetSalt(salt); err != nil {
			return fmt.Errorf("failed to generate salt: %w", err)
		}
		sw.aead, err = sw.key.NewAEAD(salt)
		if err != nil {
			return fmt.Errorf("failed to create AEAD: %w", err)
		}
		sw.saltGenerator = nil // No longer needed, so release reference.
		sw.counter = make([]byte, sw.aead.NonceSize())
		// The maximum length message is the salt (first message only), length, length tag,
		// payload, and payload tag.
		sizeBufSize := 2 + sw.aead.Overhead()
		maxPayloadBufSize := payloadSizeMask + sw.aead.Overhead()
		sw.buf = make([]byte, len(salt)+sizeBufSize+maxPayloadBufSize)
		// Store the salt at the start of sw.buf.
		copy(sw.buf, salt)
	}
	return nil
}

// encryptBlock encrypts `plaintext` in-place.  The slice must have enough capacity
// for the tag. Returns the total ciphertext length.
func (sw *Writer) encryptBlock(plaintext []byte) int { _ = "STUB: not implemented"; return 0 }

func (sw *Writer) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// LazyWrite queues p to be written, but doesn't send it until Flush() is
// called, a non-lazy write is made, or the buffer is filled.
func (sw *Writer) LazyWrite(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Locking is needed due to potential concurrency with the Flush()
// for a previous call to LazyWrite().

// p didn't fit in the buffer.  Flush the buffer and try
// again.

// Flush sends the pending data, if any.  This method is thread-safe.
func (sw *Writer) Flush() error { _ = "STUB: not implemented"; return nil }

func isZero(b []byte) bool { _ = "STUB: not implemented"; return false }

// Returns the slices of sw.buf in which to place plaintext for encryption.
func (sw *Writer) buffers() (sizeBuf, payloadBuf []byte) {
	_ = "STUB: not implemented"
	// sw.buf starts with the salt.
	return nil, nil
}

// Each Shadowsocks-TCP message consists of a fixed-length size block,
// followed by a variable-length payload block.

// ReadFrom implements the [io.ReaderFrom] interface.
func (sw *Writer) ReadFrom(r io.Reader) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// Special case: one thread-safe read, if necessary

// The first pending+overhead bytes of payloadBuf are potentially
// in use, and may be modified on the flush thread.  Data after
// that is safe to use on this thread.

// Main transfer loop

// ignore EOF as per io.ReaderFrom contract

// Adds as much of `plaintext` into the buffer as will fit, and increases
// sw.pending accordingly.  Returns the number of bytes consumed.
func (sw *Writer) enqueue(plaintext []byte) int { _ = "STUB: not implemented"; return 0 }

// Encrypts all pending data and writes it to the output.
func (sw *Writer) flush() error { _ = "STUB: not implemented"; return nil }

// sw.buf starts with the salt.

// Normally we ignore the salt at the beginning of sw.buf.

// For the first message, include the salt.  Compared to writing the salt
// separately, this saves one packet during TCP slow-start and potentially
// avoids having a distinctive size for the first packet.

// genericChunkReader is similar to io.Reader, except that it controls its own
// buffer granularity.
type genericChunkReader interface {
	// ReadChunk reads the next chunk and returns its payload.  The caller must
	// complete its use of the returned buffer before the next call.
	// The buffer is nil iff there is an error.  io.EOF indicates a close.
	ReadChunk() ([]byte, error)
}

type chunkReader struct {
	reader io.Reader
	key    *EncryptionKey
	// These are lazily initialized:
	aead cipher.AEAD
	// Index of the next encrypted chunk to read.
	counter []byte
	// Buffer for the uint16 size and its AEAD tag.  Made in init().
	payloadSizeBuf []byte
	// Holds a buffer for the payload and its AEAD tag, when needed.
	payload slicepool.LazySlice
}

// Reader is an [io.Reader] that also implements [io.WriterTo] to
// allow for piping the data without extra allocations and copies.
type Reader interface {
	io.Reader
	io.WriterTo
}

// NewReader creates a [Reader] that decrypts the given [io.Reader] using
// the shadowsocks protocol with the given encryption key.
func NewReader(reader io.Reader, key *EncryptionKey) Reader {
	_ = "STUB: not implemented"
	return *new(Reader)
}

// init reads the salt from the inner Reader and sets up the AEAD object
func (cr *chunkReader) init() (err error) {
	if cr.aead == nil {
		// For chacha20-poly1305, SaltSize is 32, NonceSize is 12 and Overhead is 16.
		salt := make([]byte, cr.key.SaltSize())
		if _, err := io.ReadFull(cr.reader, salt); err != nil {
			if err != io.EOF && err != io.ErrUnexpectedEOF {
				err = fmt.Errorf("failed to read salt: %w", err)
			}
			return err
		}
		cr.aead, err = cr.key.NewAEAD(salt)
		if err != nil {
			return fmt.Errorf("failed to create AEAD: %w", err)
		}
		cr.counter = make([]byte, cr.aead.NonceSize())
		cr.payloadSizeBuf = make([]byte, 2+cr.aead.Overhead())
	}
	return nil
}

// readMessage reads, decrypts, and verifies a single AEAD ciphertext.
// The ciphertext and tag (i.e. "overhead") must exactly fill `buf`,
// and the decrypted message will be placed in buf[:len(buf)-overhead].
// Returns an error only if the block could not be read.
func (cr *chunkReader) readMessage(buf []byte) error { _ = "STUB: not implemented"; return nil }

// ReadChunk returns the next chunk from the stream.  Callers must fully
// consume and discard the previous chunk before calling ReadChunk again.
func (cr *chunkReader) ReadChunk() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Release the previous payload buffer.

// In Shadowsocks-AEAD, each chunk consists of two
// encrypted messages.  The first message contains the payload length,
// and the second message is the payload.  Idle read threads will
// block here until the next chunk.

// This code is unreachable if the constants are set correctly.

// EOF is not expected mid-chunk.

// readConverter adapts from ChunkReader, with source-controlled
// chunk sizes, to Go-style IO.
type readConverter struct {
	cr       genericChunkReader
	leftover []byte
}

func (c *readConverter) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *readConverter) WriteTo(w io.Writer) (written int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Ensures that c.leftover is nonempty.  If leftover is empty, this method
// waits for incoming data and decrypts it.
// Returns an error only if c.leftover could not be populated.
func (c *readConverter) ensureLeftover() error { _ = "STUB: not implemented"; return nil }

// increment little-endian encoded unsigned integer b. Wrap around on overflow.
func increment(b []byte) { _ = "STUB: not implemented"; return }
