// Copyright 2025 The Outline Authors
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

package soax

import (
	"net/url"
	"time"

	"golang.getoutline.org/sdk/transport"
	"golang.getoutline.org/sdk/transport/socks5"
	"golang.getoutline.org/sdk/x/httpconnect"
)

// The official address of the SOAX proxy.
const proxyAddress = "proxy.soax.com:5000"

// ProxyAuthConfig represents the authentication credentials for the SOAX proxy.
type ProxyAuthConfig struct {
	PackageID  int
	PackageKey string
}

// ProxyNodeConfig defines the geographic and network restrictions for the proxy node.
type ProxyNodeConfig struct {
	CountryCode string
	RegionID    string
	CityID      string
	ISPID       string
}

// SessionConfig defines the session management parameters for the proxy.
type SessionConfig struct {
	// Unique ID for this session
	ID string
	// How long this session should last
	Duration time.Duration
	// Rotate the node if the session goes idle for longer than this.
	IdleTTL time.Duration
}

// Indicates that the session should not persist. No session ID will be set
// and each request will get a different node.
var SessionNotPersistent = SessionConfig{Duration: -1}

// ProxySessionConfig defines how to connect to the SOAX proxy.
type ProxySessionConfig struct {
	// Authentication credentials
	Auth ProxyAuthConfig

	// Node restrictions
	Node ProxyNodeConfig

	// Session management
	// Set to SessionNotPersistent to disable session persistence and get different nodes per request.
	Session SessionConfig

	// Endpoint address. If empty, defaults to proxy.soax.com:5000.
	// Useful for testing or reverse proxies.
	Endpoint string
}

// ProxySession represents a session with unique SessionID, created from a SessionConfig.
type ProxySession struct {
	config ProxySessionConfig
}

func (c *ProxySessionConfig) newUserPassword() *url.Userinfo { _ = "STUB: not implemented"; return nil }

func (c *ProxySessionConfig) NewSession() *ProxySession { _ = "STUB: not implemented"; return nil }

// 1 hour is the maximum session length as per
// https://helpcenter.soax.com/en/articles/9925415-building-sticky-sessions-connection

// Ensure IPs won't be released if they become idle during the session as per
// https://helpcenter.soax.com/en/articles/9939557-understanding-session-parameters#h_ffc53ee8d5

// NewSOCKS5Client creates a [socks5.Client] that connects through the SOAX proxy.
func (c *ProxySession) NewSOCKS5Client() (*socks5.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewWebProxyStreamDialer creates a [transport.StreamDialer] that connects through the SOAX proxy.
// It uses HTTP CONNECT, so it only supports TCP.
func (c *ProxySession) NewWebProxyStreamDialer(opts ...httpconnect.TransportOption) (transport.StreamDialer, error) {
	_ = "STUB: not implemented"
	return *new(transport.StreamDialer), nil
}

// TODO: Add support for Respond-With, as explained in
// https://helpcenter.soax.com/en/articles/9956908-getting-node-information-with-the-respond-with-header
