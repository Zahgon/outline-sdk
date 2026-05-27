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
	"context"
	"net/url"
)

// Config is a pre-parsed generic config created from pipe-separated URLs.
type Config struct {
	URL        url.URL
	BaseConfig *Config
}

// BuildFunc is a function that creates an instance of ObjectType given a [Config].
type BuildFunc[ObjectType any] func(ctx context.Context, config *Config) (ObjectType, error)

// TypeRegistry registers config types.
type TypeRegistry[ObjectType any] interface {
	RegisterType(subtype string, newInstance BuildFunc[ObjectType])
}

// ExtensibleProvider creates instances of ObjectType in a way that can be extended via its [TypeRegistry] interface.
type ExtensibleProvider[ObjectType comparable] struct {
	// Instance to return when config is nil.
	BaseInstance ObjectType
	builders     map[string]BuildFunc[ObjectType]
}

var (
	_ BuildFunc[any]    = (*ExtensibleProvider[any])(nil).NewInstance
	_ TypeRegistry[any] = (*ExtensibleProvider[any])(nil)
)

// NewExtensibleProvider creates an [ExtensibleProvider] with the given base instance.
func NewExtensibleProvider[ObjectType comparable](baseInstance ObjectType) ExtensibleProvider[ObjectType] {
	_ = "STUB: not implemented"
	return nil
}

func (p *ExtensibleProvider[ObjectType]) ensureBuildersMap() map[string]BuildFunc[ObjectType] {
	_ = "STUB: not implemented"
	return nil
}

// RegisterType will register a factory for the given subtype.
func (p *ExtensibleProvider[ObjectType]) RegisterType(subtype string, newInstance BuildFunc[ObjectType]) {
	_ = "STUB: not implemented"
	return
}

// NewInstance creates a new instance of ObjectType according to the config.
func (p *ExtensibleProvider[ObjectType]) NewInstance(ctx context.Context, config *Config) (ObjectType, error) {
	_ = "STUB: not implemented"
	return *new(ObjectType), nil
}

// ParseConfig will parse a config given as a string and return the structured [Config].
func ParseConfig(configText string) (*Config, error) { _ = "STUB: not implemented"; return nil, nil }

// Make it "<scheme>:" if it's only "<scheme>" to parse as a URL.
