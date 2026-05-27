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

	"golang.getoutline.org/sdk/transport"
)

func registerOverrideStreamDialer(r TypeRegistry[transport.StreamDialer], typeID string, newSD BuildFunc[transport.StreamDialer]) {
	_ = "STUB: not implemented"
	return
}

func registerOverridePacketDialer(r TypeRegistry[transport.PacketDialer], typeID string, newPD BuildFunc[transport.PacketDialer]) {
	_ = "STUB: not implemented"
	return
}

func newOverrideFromURL(configURL url.URL) (func(string) (string, error), error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Optimization when we fully override the address.
