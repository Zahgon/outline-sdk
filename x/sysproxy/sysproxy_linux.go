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

//go:build linux && !android

// TODO(fortuna): remove once Linux tests are re-enabled for Ubuntu 24.
//lint:file-ignore U1000 getter functions are only called from tests, which are temporarily disabled on Linux

package sysproxy

type ProxyType string

const (
	proxyTypeHTTP  ProxyType = "http"
	proxyTypeHTTPS ProxyType = "https"
	proxyTypeSOCKS ProxyType = "socks"
)

func SetWebProxy(host string, port string) error {
	_ = "STUB: not implemented"
	// Set HTTP and HTTPS proxy settings
	return nil
}

func DisableWebProxy() error { _ = "STUB: not implemented"; return nil }

func SetSOCKSProxy(host string, port string) error {
	_ = "STUB: not implemented"
	// Set SOCKS proxy settings
	return nil
}

func DisableSOCKSProxy() error { _ = "STUB: not implemented"; return nil }

func setManualMode() error { _ = "STUB: not implemented"; return nil }

func setProxySettings(p ProxyType, host string, port string) error {
	_ = "STUB: not implemented"
	return nil
}

func gnomeSettingsSetString(settings, key, value string) error {
	_ = "STUB: not implemented"
	return nil
}

func getWebProxy() (host string, port string, enabled bool, err error) {
	_ = "STUB: not implemented"
	return "", "", false, nil
}

func getSOCKSProxy() (host string, port string, enabled bool, err error) {
	_ = "STUB: not implemented"
	return "", "", false, nil
}

func gnomeSettingsGetString(settings, key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
