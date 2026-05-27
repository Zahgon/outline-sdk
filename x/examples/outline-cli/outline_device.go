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

import (
	"net"

	"golang.getoutline.org/sdk/network"
	"golang.getoutline.org/sdk/transport"
	"golang.getoutline.org/sdk/x/configurl"
)

const (
	connectivityTestDomain   = "www.google.com"
	connectivityTestResolver = "1.1.1.1:53"
)

type OutlineDevice struct {
	network.IPDevice
	sd    transport.StreamDialer
	pp    *outlinePacketProxy
	svrIP net.IP
}

var configModule = configurl.NewDefaultProviders()

func NewOutlineDevice(transportConfig string) (od *OutlineDevice, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *OutlineDevice) Close() error { _ = "STUB: not implemented"; return nil }

func (d *OutlineDevice) Refresh() error { _ = "STUB: not implemented"; return nil }

func (d *OutlineDevice) GetServerIP() net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

func resolveShadowsocksServerIPFromConfig(transportConfig string) (net.IP, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil
}

// todo: we only tested IPv4 routing table, need to test IPv6 in the future
