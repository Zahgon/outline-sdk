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
	"context"

	"golang.getoutline.org/sdk/transport"
	"golang.getoutline.org/sdk/transport/socks5"
)

func registerSOCKS5StreamDialer(r TypeRegistry[transport.StreamDialer], typeID string, newSD BuildFunc[transport.StreamDialer]) {
	_ = "STUB: not implemented"
	return
}

func registerSOCKS5PacketDialer(r TypeRegistry[transport.PacketDialer], typeID string, newSD BuildFunc[transport.StreamDialer], newPD BuildFunc[transport.PacketDialer]) {
	_ = "STUB: not implemented"
	return
}

func registerSOCKS5PacketListener(r TypeRegistry[transport.PacketListener], typeID string, newSD BuildFunc[transport.StreamDialer], newPD BuildFunc[transport.PacketDialer]) {
	_ = "STUB: not implemented"
	return
}

func newSOCKS5Client(ctx context.Context, config Config, newSD BuildFunc[transport.StreamDialer]) (*socks5.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
