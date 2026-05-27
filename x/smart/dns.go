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
	"net"
	"time"

	"golang.org/x/net/dns/dnsmessage"
)

// makeFullyQualified makes the domain fully-qualified, ending on a dot (".").
// This is useful in domain resolution to avoid ambiguity with local domains
// and domain search.
func makeFullyQualified(domain string) string { _ = "STUB: not implemented"; return "" }

// mixCase randomizes the case of the domain letters.
func mixCase(domain string) string { _ = "STUB: not implemented"; return "" }

func evaluateNetResolver(ctx context.Context, resolver *net.Resolver, testDomain string) ([]net.IP, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: consider validating the IPs: fingerprint, TCP connection, hardcoded ground truth, trusted response, TLS connection.

func getIPs(answers []dnsmessage.Resource) []net.IP { _ = "STUB: not implemented"; return nil }

func evaluateAddressResponse(response dnsmessage.Message, requestDomain string) ([]net.IP, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// All popular recursive resolvers we tested maintain the domain case of the request.
// Note that this is not the case of authoritative resolvers. Some of them will return
// a fully normalized domain name, or normalize part of it.

func evaluateCNAMEResponse(response dnsmessage.Message, requestDomain string) error {
	_ = "STUB: not implemented"
	return nil
}

func testDNSResolver(baseCtx context.Context, oneTestTimeout time.Duration, resolver *smartResolver, testDomain string) ([]net.IP, error) {
	_ = "STUB: not implemented"
	// We special case the system resolver, since we can't get a dns.RoundTripper.
	return nil, nil
}

// For secure DNS, we just need to check if we can communicate with it.
// No need to analyze content, since it is protected by TLS.

// TODO(fortuna): Consider testing whether we can establish a TCP connection to ip:443.

// Run CNAME test, which helps in case the resolver returns a public IP, as is the
// case in China.
