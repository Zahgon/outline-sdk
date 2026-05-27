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

//go:build windows

package sysproxy

import (
	"golang.org/x/sys/windows"
)

type proxySettings struct {
	proxyServer   string
	proxyOverride string
}

var (
	modwininet            = windows.NewLazySystemDLL("wininet.dll")
	procInternetSetOption = modwininet.NewProc("InternetSetOptionW")
)

// https://learn.microsoft.com/en-us/windows/win32/wininet/option-flags
// INTERNET_OPTION_SETTINGS_CHANGED: 39
// Notifies the system that the registry settings have been changed so that it verifies the settings on the next call to InternetConnect.
// INTERNET_OPTION_REFRESH: 37
// Causes the proxy data to be reread from the registry for a handle. No buffer is required.
// This option can be used on the HINTERNET handle returned by InternetOpen.
// This is used by InternetSetOption.
const (
	INTERNET_OPTION_SETTINGS_CHANGED = 39
	INTERNET_OPTION_REFRESH          = 37
)

func SetWebProxy(host string, port string) error { _ = "STUB: not implemented"; return nil }

func DisableWebProxy() error {
	_ = "STUB: not implemented"
	// disable proxy settings
	return nil
}

// SetProxy does nothing on windows platforms.
func SetSOCKSProxy(host string, port string) error { _ = "STUB: not implemented"; return nil }

// SetProxy does nothing on windows platforms.
func DisableSOCKSProxy() error { _ = "STUB: not implemented"; return nil }

func setProxySettings(settings *proxySettings) error { _ = "STUB: not implemented"; return nil }

// Finally, enable the proxy

// Refresh the settings

func disableProxy() error { _ = "STUB: not implemented"; return nil }

// Set ProxyEnable to 0

// Refresh the settings

// https://learn.microsoft.com/en-us/windows/win32/api/wininet/nf-wininet-internetsetoptionw
// internetSetOption sets an Internet option.
func internetSetOption(hInternet uintptr, dwOption int, lpBuffer uintptr, dwBufferLength uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func notifyWinInetProxySettingsChanged() error { _ = "STUB: not implemented"; return nil }

func getWebProxy() (host string, port string, enabled bool, err error) {
	_ = "STUB: not implemented"
	return "", "", false, nil
}

// Read back the value of ProxyEnable

func getSOCKSProxy() (host string, port string, enabled bool, err error) {
	_ = "STUB: not implemented"
	return "", "", false, nil
}

// Read back the value of ProxyEnable
