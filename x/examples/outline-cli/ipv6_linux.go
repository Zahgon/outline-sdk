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

package main

const disableIPv6ProcFile = "/proc/sys/net/ipv6/conf/all/disable_ipv6"

// enableIPv6 enables or disables the IPv6 support for the Linux system.
// It returns the previous setting value so the caller can restore it.
// Non-nil error means we cannot find the IPv6 setting.
func enableIPv6(enabled bool) (bool, error) { _ = "STUB: not implemented"; return false, nil }
