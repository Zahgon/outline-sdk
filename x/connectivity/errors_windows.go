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

//go:build windows

package connectivity

import (
	"syscall"
)

func systemErrnoName(errno syscall.Errno) string {
	_ = "STUB: not implemented"
	// Windows socket API errors
	// Official list at https://learn.microsoft.com/en-us/windows/win32/winsock/windows-sockets-error-codes-2.
	// Easy to parse list at https://cs.opensource.google/go/x/sys/+/master:windows/zerrors_windows.go,
	// then restricted to those starting in "WSAE"
	return ""
}
