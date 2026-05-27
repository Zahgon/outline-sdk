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

package split

import (
	"io"
)

type splitWriter struct {
	writer io.Writer
	// Bytes until the next split. This must always be > 0, unless splits are done.
	nextSplitBytes    int64
	nextSegmentLength func() int64
}

var _ io.Writer = (*splitWriter)(nil)

type splitWriterReaderFrom struct {
	*splitWriter
	rf io.ReaderFrom
}

var _ io.ReaderFrom = (*splitWriterReaderFrom)(nil)

// SplitIterator is a function that returns how many bytes until the next split point, or zero if there are no more splits to do.
type SplitIterator func() int64

// NewFixedSplitIterator is a helper function that returns a [SplitIterator] that returns the input number once, followed by zero.
// This is helpful for when you want to split the stream once in a fixed position.
func NewFixedSplitIterator(n int64) SplitIterator {
	_ = "STUB: not implemented"
	return *new(SplitIterator)
}

// RepeatedSplit represents a split sequence of count segments with bytes length.
type RepeatedSplit struct {
	Count int
	Bytes int64
}

// NewRepeatedSplitIterator is a helper function that returns a [SplitIterator] that returns split points according to splits.
// The splits input represents pairs of (count, bytes), meaning a sequence of count splits with bytes length.
// This is helpful for when you want to split the stream repeatedly at different positions and lengths.
func NewRepeatedSplitIterator(splits ...RepeatedSplit) SplitIterator {
	_ = "STUB: not implemented"
	// Make sure we don't edit the original slice.
	return *new(SplitIterator)
}

// Remove no-op splits.

// NewWriter creates a split Writer that calls the nextSegmentLength [SplitIterator] to determine the number bytes until the next split
// point until it returns zero.
func NewWriter(writer io.Writer, nextSegmentLength SplitIterator) io.Writer {
	_ = "STUB: not implemented"
	return *new(io.Writer)
}

// ReadFrom implements io.ReaderFrom.
func (w *splitWriterReaderFrom) ReadFrom(source io.Reader) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Source is done before the split happened. Return.

func (w *splitWriter) advance(n int64) { _ = "STUB: not implemented"; return }

// Done with splits: return.

// Split done, set up the next split.

// Write implements io.Writer.
func (w *splitWriter) Write(data []byte) (written int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}
