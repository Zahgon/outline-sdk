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

package shared_backend

import (
	"time"

	"golang.getoutline.org/sdk/transport/shadowsocks"

	_ "golang.org/x/mobile/bind"
)

type ConnectivityTestProtocolConfig struct {
	TCP bool `json:"tcp"`
	UDP bool `json:"udp"`
}

type ConnectivityTestResult struct {
	// Inputs
	Proxy    string `json:"proxy"`
	Resolver string `json:"resolver"`
	Proto    string `json:"proto"`
	Prefix   string `json:"prefix"`
	// Observations
	Time       time.Time              `json:"time"`
	DurationMs int64                  `json:"durationMs"`
	Error      *ConnectivityTestError `json:"error"`
}

type ConnectivityTestError struct {
	// TODO: add Shadowsocks/Transport error
	Op string `json:"operation"`
	// Posix error, when available
	PosixError string `json:"posixError"`
	// TODO: remove IP addresses
	Msg string `json:"message"`
}

type ConnectivityTestRequest struct {
	AccessKey string                         `json:"accessKey"`
	Domain    string                         `json:"domain"`
	Resolvers []string                       `json:"resolvers"`
	Protocols ConnectivityTestProtocolConfig `json:"protocols"`
}

type sessionConfig struct {
	Hostname  string
	Port      int
	CryptoKey *shadowsocks.EncryptionKey
	Prefix    Prefix
}

type Prefix []byte

func ConnectivityTest(request ConnectivityTestRequest) ([]ConnectivityTestResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: limit number of IPs. Or force an input IP?

type PlatformMetadata struct {
	OS string `json:"operatingSystem"`
}

func Platform() PlatformMetadata { _ = "STUB: not implemented"; return *new(PlatformMetadata) }

func makeErrorRecord(err error) *ConnectivityTestError { _ = "STUB: not implemented"; return nil }

func unwrapAll(err error) error { _ = "STUB: not implemented"; return nil }

func (p Prefix) String() string { _ = "STUB: not implemented"; return "" }

func parseAccessKey(accessKey string) (*sessionConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Host is a <host>:<port> string

func ParseStringPrefix(utf8Str string) (Prefix, error) {
	_ = "STUB: not implemented"
	return *new(Prefix), nil
}
