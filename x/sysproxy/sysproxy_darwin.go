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

//go:build darwin && !ios

package sysproxy

type ProxyType string

const (
	proxyTypeHTTP  ProxyType = "web"
	proxyTypeHTTPS ProxyType = "secureweb"
	proxyTypeSOCKS ProxyType = "socks"
)

type proxySettings struct {
	host    string
	port    string
	enabled bool
}

func SetWebProxy(host string, port string) error {
	_ = "STUB: not implemented"
	// Get the active network interface
	return nil
}

// Set the web proxy and secure web proxy

// revert previous changes

func DisableWebProxy() error {
	_ = "STUB: not implemented"
	// Get the active network interface
	return nil
}

// disable the web proxy and secure web proxy

func SetSOCKSProxy(host string, port string) error {
	_ = "STUB: not implemented"
	// Get the active network interface
	return nil
}

// Set the SOCKS proxy

func DisableSOCKSProxy() error {
	_ = "STUB: not implemented"
	// Get the active network interface
	return nil
}

// disable the SOCKS proxy

// getActiveNetworkInterface finds the active network interface using shell commands.
// https://keith.github.io/xcode-man-pages/networksetup.8.html#listnetworkserviceorder
func getActiveNetworkInterface() (string, error) {
	_ = "STUB: not implemented"
	// cmd := "networksetup -listnetworkserviceorder | grep `route -n get 0.0.0.0 | grep 'interface' | cut -d ':' -f2` -B 1 | head -n 1 | cut -d ' ' -f2"
	return "", nil
}

// getDefaultRouteInterface gets the default route interface using os command.
// Example output of `route get default` on macOS:
//
//	route to: default
//	destination: default
//	mask: default
//	gateway: 192.168.1.1
//	interface: en0
//	flags: <UP,GATEWAY,DONE,STATIC,PRCLONING,GLOBAL>
//	recvpipe  sendpipe  ssthresh  rtt,msec    rttvar  hopcount      mtu     expire
//	0         0         0         0         0         0      1500         0
func getDefaultRouteInterface() (string, error) {
	_ = "STUB: not implemented"
	// Execute a command to get the default route
	return "", nil
}

// Extract the interface name from the command output

// getNetworkServiceName parses the output of networksetup -listnetworkserviceorder to find
// the network service name for a given hardware port (e.g. Wi-Fi for en0)
func getNetworkServiceName(output, hardwarePort string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// example line: (Hardware Port: Wi-Fi, Device: en0)

// setProxySettings sets the specified type of proxy on the given network interface.
// https://keith.github.io/xcode-man-pages/networksetup.8.html#getsecurewebproxy
func setProxySettings(p ProxyType, interfaceName string, host string, port string) error {
	_ = "STUB: not implemented"
	return nil
}

// disableProxy turns off the specified type of proxy on the given network interface.
// https://keith.github.io/xcode-man-pages/networksetup.8.html#setwebproxystate
func disableProxy(p ProxyType, interfaceName string) error { _ = "STUB: not implemented"; return nil }

func getProxySettings(p ProxyType, interfaceName string) (*proxySettings, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseProxySettings(commandOutput string) (*proxySettings, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getWebProxy() (host string, port string, enabled bool, err error) {
	_ = "STUB: not implemented"
	return "", "", false, nil
}

func getSOCKSProxy() (host string, port string, enabled bool, err error) {
	_ = "STUB: not implemented"
	return "", "", false, nil
}
