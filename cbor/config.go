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

// EncodeMode specifies how Encoder serializes data items.
type EncodeMode int

const (
	// EncodeModePreferred encodes data items with the preferred serialization defined in RFC 8949 Section 4.1.
	// Integers, lengths, and tag numbers use the shortest form, and floating-point values use the shortest
	// of float16, float32, and float64 that preserves the value. This is the default mode.
	// Map keys are written in the Go map iteration order unless MapSortEnabled is set.
	EncodeModePreferred EncodeMode = iota
	// EncodeModeTypePreserving encodes integers and floating-point values with the width of their Go types,
	// for example int as an 8-byte integer and float64 as a double-precision value, so that Decoder
	// returns values of the same width (an unsigned value that fits in the signed type of the same
	// width is returned as that signed type). This was the behavior of go-cbor v1.3.3 and earlier.
	// It does not satisfy the preferred serialization of RFC 8949.
	EncodeModeTypePreserving
	// EncodeModeCoreDeterministic encodes data items with the core deterministic encoding defined in
	// RFC 8949 Section 4.2.1: the preferred serialization with map keys sorted in the bytewise
	// lexicographic order of their deterministic encodings.
	EncodeModeCoreDeterministic
	// EncodeModeLengthFirstDeterministic encodes data items with the preferred serialization and the
	// length-first map key ordering defined in RFC 8949 Section 4.2.3 (the canonical CBOR of RFC 7049):
	// shorter encoded keys sort earlier, and keys of the same length sort in bytewise lexicographic order.
	EncodeModeLengthFirstDeterministic
)

// String returns the name of the encode mode.
func (mode EncodeMode) String() string {
	switch mode {
	case EncodeModePreferred:
		return "Preferred"
	case EncodeModeTypePreserving:
		return "TypePreserving"
	case EncodeModeCoreDeterministic:
		return "CoreDeterministic"
	case EncodeModeLengthFirstDeterministic:
		return "LengthFirstDeterministic"
	}
	return "Unknown"
}

// Config represents a configuration for CBOR encoder and decoder.
type Config struct {
	// MapSortEnabled sorts map keys in the bytewise lexicographic order of their encodings
	// (RFC 8949 Section 4.2.1) in EncodeModePreferred and EncodeModeTypePreserving.
	// The deterministic encode modes always sort map keys.
	MapSortEnabled bool

	encodeMode EncodeMode
}

// NewConfig returns a new config instance.
func NewConfig() *Config {
	return &Config{
		MapSortEnabled: false,
		encodeMode:     EncodeModePreferred,
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

// SetEncodeMode sets the encode mode.
func (config *Config) SetEncodeMode(mode EncodeMode) {
	config.encodeMode = mode
}

// EncodeMode returns the encode mode.
func (config *Config) EncodeMode() EncodeMode {
	return config.encodeMode
}
