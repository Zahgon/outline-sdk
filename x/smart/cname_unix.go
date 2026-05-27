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

//go:build unix

package smart

/*
#include <stdlib.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netdb.h>
*/
import "C"

import (
	"context"
)

// lookupCNAME provides functionality equivalent to net.DefaultResolver.LookupCNAME. However,
// net.DefaultResolver.LookupCNAME uses libresolv on unix, and, on Android and iOS, it tries
// to connect to [::1]:53 (probably from /etc/resolv.conf) and the connection is refused.
// Instead, we use getaddrinfo to get the canonical name.
func lookupCNAME(ctx context.Context, domain string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func lookupCNAMEBlocking(host string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Extract canonical name
