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

package configurl

import (
	"net/url"

	"golang.getoutline.org/sdk/dns"
	"golang.getoutline.org/sdk/transport"
)

func registerDO53StreamDialer(r TypeRegistry[transport.StreamDialer], typeID string, newSD BuildFunc[transport.StreamDialer], newPD BuildFunc[transport.PacketDialer]) {
	_ = "STUB: not implemented"
	return
}

func registerDOHStreamDialer(r TypeRegistry[transport.StreamDialer], typeID string, newSD BuildFunc[transport.StreamDialer]) {
	_ = "STUB: not implemented"
	return
}

func newDO53Resolver(config url.URL, sd transport.StreamDialer, pd transport.PacketDialer) (dns.Resolver, error) {
	_ = "STUB: not implemented"
	return *new(dns.Resolver), nil
}

// If the message is truncated, retry over TCP.
// See https://datatracker.ietf.org/doc/html/rfc1123#page-75.

func newDOHResolver(config url.URL, sd transport.StreamDialer) (dns.Resolver, error) {
	_ = "STUB: not implemented"
	return *new(dns.Resolver), nil
}
