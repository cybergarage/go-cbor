// Copyright (C) 2022 The go-cbor Authors All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cbor

// Config represents a configuration for CBOR encoder and decoder.
type Config struct {
	MapSortEnabled       bool
	TypeRetentionEnabled bool
}

// NewConfig returns a new config instance.
func NewConfig() *Config {
	return &Config{
		MapSortEnabled:       false,
		TypeRetentionEnabled: true,
	}
}

// SetMapSortEnabled sets a flag to sort map keys.
func (config *Config) SetMapSortEnabled(flag bool) {
	config.MapSortEnabled = flag
}

// IsMapSortEnabled returns true whether the map keys are sorted.
func (config *Config) IsMapSortEnabled() bool {
	return config.MapSortEnabled
}

// SetTypeRetentionEnabled sets a flag to retain the type information.
func (config *Config) SetTypeRetentionEnabled(flag bool) {
	config.TypeRetentionEnabled = flag
}

// IsTypeRetentionEnabled returns true whether the type information is retained.
func (config *Config) IsTypeRetentionEnabled() bool {
	return config.TypeRetentionEnabled
}
