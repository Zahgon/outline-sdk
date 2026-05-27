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
	"github.com/vishvananda/netlink"
)

var ipRule *netlink.Rule = nil

func startRouting(proxyIP string, config *RoutingConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func stopRouting(routingTable int) { _ = "STUB: not implemented"; return }

func setupRoutingTable(routingTable int, tunName, gwSubnet string, tunIP string) error {
	_ = "STUB: not implemented"
	return nil
}

func cleanUpRoutingTable(routingTable int) error { _ = "STUB: not implemented"; return nil }

func setupIpRule(svrIp string, routingTable, routingPriority int) error {
	_ = "STUB: not implemented"
	return nil
}

// todo: exclude server IP will cause issues when accessing services on the same server,
//       use fwmask to protect the shadowsocks socket instead

// assuming duplicate from previous run, just to make sure it does not stays stale forever in the routing table

func cleanUpRule() error { _ = "STUB: not implemented"; return nil }
