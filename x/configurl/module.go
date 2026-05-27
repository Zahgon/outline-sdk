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
	"net/url"

	"golang.getoutline.org/sdk/transport"
)

// ProviderContainer contains providers for the creation of network objects based on a config. The config is
// extensible by registering providers for different config subtypes.
type ProviderContainer struct {
	StreamDialers   ExtensibleProvider[transport.StreamDialer]
	PacketDialers   ExtensibleProvider[transport.PacketDialer]
	PacketListeners ExtensibleProvider[transport.PacketListener]
}

// NewProviderContainer creates a [ProviderContainer] with the base instances properly initialized.
func NewProviderContainer() *ProviderContainer { _ = "STUB: not implemented"; return nil }

// RegisterDefaultProviders registers a set of default providers with the providers in [ProviderContainer].
func RegisterDefaultProviders(c *ProviderContainer) *ProviderContainer {
	_ = "STUB: not implemented"
	// Please keep the list in alphabetical order.
	return nil
}

// NewDefaultProviders creates a [ProviderContainer] with a set of default providers already registered.
func NewDefaultProviders() *ProviderContainer { _ = "STUB: not implemented"; return nil }

// NewStreamDialer creates a [transport.StreamDialer] according to the config text.
func (p *ProviderContainer) NewStreamDialer(ctx context.Context, configText string) (transport.StreamDialer, error) {
	_ = "STUB: not implemented"
	return *new(transport.StreamDialer), nil
}

// NewPacketDialer creates a [transport.PacketDialer] according to the config text.
func (p *ProviderContainer) NewPacketDialer(ctx context.Context, configText string) (transport.PacketDialer, error) {
	_ = "STUB: not implemented"
	return *new(transport.PacketDialer), nil
}

// NewPacketListner creates a [transport.PacketListener] according to the config text.
func (p *ProviderContainer) NewPacketListener(ctx context.Context, configText string) (transport.PacketListener, error) {
	_ = "STUB: not implemented"
	return *new(transport.PacketListener), nil
}

// SanitizeConfig removes sensitive information from the given config so it can be safely be used in logging and debugging.
func SanitizeConfig(configStr string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Do nothing if the config is empty

// No sanitization needed

func sanitizeSOCKS5URL(u *url.URL) (string, error) { _ = "STUB: not implemented"; return "", nil }
