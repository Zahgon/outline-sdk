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

package configurl

import (
	"net/url"

	"golang.getoutline.org/sdk/transport"
	"golang.getoutline.org/sdk/transport/shadowsocks"
)

func registerShadowsocksStreamDialer(r TypeRegistry[transport.StreamDialer], typeID string, newSD BuildFunc[transport.StreamDialer]) {
	_ = "STUB: not implemented"
	return
}

func registerShadowsocksPacketDialer(r TypeRegistry[transport.PacketDialer], typeID string, newPD BuildFunc[transport.PacketDialer]) {
	_ = "STUB: not implemented"
	return
}

// TODO: support UDP prefix.

func registerShadowsocksPacketListener(r TypeRegistry[transport.PacketListener], typeID string, newPD BuildFunc[transport.PacketDialer]) {
	_ = "STUB: not implemented"
	return
}

type shadowsocksConfig struct {
	serverAddress string
	cryptoKey     *shadowsocks.EncryptionKey
	prefix        []byte
}

func parseShadowsocksURL(url url.URL) (*shadowsocksConfig, error) {
	_ = "STUB: not implemented"
	// attempt to decode as SIP002 URI format and
	// fall back to legacy base64 format if decoding fails
	return nil, nil
}

// parseShadowsocksLegacyBase64URL parses URL based on legacy base64 format:
// https://shadowsocks.org/doc/configs.html#uri-and-qr-code
func parseShadowsocksLegacyBase64URL(url url.URL) (*shadowsocksConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If decoding fails, return the original url with error

// if parsing fails, return the original url with error

// extend this check to see if decoded string contains contains other valid fields

// parseShadowsocksSIP002URL parses URL based on SIP002 format:
// https://shadowsocks.org/doc/sip002.html
func parseShadowsocksSIP002URL(url url.URL) (*shadowsocksConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cipher info can be optionally encoded with Base64URL.

// Try base64 decoding in legacy mode

func parseStringPrefix(utf8Str string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func sanitizeShadowsocksURL(u url.URL) (string, error) { _ = "STUB: not implemented"; return "", nil }
