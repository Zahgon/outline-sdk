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

package smart

import (
	"context"
	"time"
)

// Returns a read channel that is already closed.
func newClosedChanel() <-chan struct{} { _ = "STUB: not implemented"; return nil }

// raceTests will call the test function on each entry until it finds an entry for which the test returns nil error.
// That entry is returned. A test is only started after the previous test finished or maxWait is done, whichever
// happens first. That way you bound the wait for a test, and they may overlap.
// The test function should make use of the context to stop doing work when the race is done and it is no longer needed.
func raceTests[E any, R any](ctx context.Context, maxWait time.Duration, entries []E, test func(index int, entry E) (R, error)) (R, error) {
	_ = "STUB: not implemented"
	return *new(R), nil
}

// Communicates the result of each test.

// Search cancelled, quit.

// Ready to start testing another resolver.

// Done with entries. No longer trigger on waitCh.

// Got a test result.
