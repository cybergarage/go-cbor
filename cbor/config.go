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

// DecodeMode specifies which encodings Decoder accepts.
type DecodeMode int

const (
	// DecodeModeLenient accepts any well-formed encoding, including non-shortest arguments,
	// non-shortest floating-point values, indefinite-length items, and unsorted map keys.
	// This is the default mode.
	DecodeModeLenient DecodeMode = iota
	// DecodeModeCoreDeterministic accepts only the core deterministic encoding defined in
	// RFC 8949 Section 4.2.1, and returns ErrDecode for non-shortest arguments, non-shortest
	// floating-point values, indefinite-length items, map keys that are not in the bytewise
	// lexicographic order of their encodings (including duplicate keys), and bignums that are
	// not in the preferred serialization (Section 3.4.3).
	DecodeModeCoreDeterministic
)

// String returns the name of the decode mode.
func (mode DecodeMode) String() string {
	switch mode {
	case DecodeModeLenient:
		return "Lenient"
	case DecodeModeCoreDeterministic:
		return "CoreDeterministic"
	}
	return "Unknown"
}

// DuplicateMapKeyMode specifies how Decoder handles duplicate map keys (RFC 8949 Section 5.6).
type DuplicateMapKeyMode int

const (
	// DuplicateMapKeyAllowed keeps the value of the last duplicate key. This is the default mode.
	DuplicateMapKeyAllowed DuplicateMapKeyMode = iota
	// DuplicateMapKeyRejected returns ErrDecode when a map has duplicate keys.
	// Keys are compared as decoded Go values, so keys encoded with different argument
	// sizes, such as 0x01 (int8) and 0x190001 (int16), are not detected as duplicates.
	// Use DecodeModeCoreDeterministic to reject them as well.
	DuplicateMapKeyRejected
)

// DefaultMaxNestedLevels is the default maximum nesting depth of arrays, maps, and tags accepted by Decoder.
const DefaultMaxNestedLevels = 128

// Config represents a configuration for CBOR encoder and decoder.
type Config struct {
	// MapSortEnabled sorts map keys in the bytewise lexicographic order of their encodings
	// (RFC 8949 Section 4.2.1) in EncodeModePreferred and EncodeModeTypePreserving.
	// The deterministic encode modes always sort map keys.
	MapSortEnabled bool

	encodeMode             EncodeMode
	decodeMode             DecodeMode
	duplicateMapKeyMode    DuplicateMapKeyMode
	maxNestedLevels        int
	utf8ValidationDisabled bool
}

// NewConfig returns a new config instance.
func NewConfig() *Config {
	return &Config{
		MapSortEnabled:         false,
		encodeMode:             EncodeModePreferred,
		decodeMode:             DecodeModeLenient,
		duplicateMapKeyMode:    DuplicateMapKeyAllowed,
		maxNestedLevels:        DefaultMaxNestedLevels,
		utf8ValidationDisabled: false,
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

// SetDecodeMode sets the decode mode.
func (config *Config) SetDecodeMode(mode DecodeMode) {
	config.decodeMode = mode
}

// DecodeMode returns the decode mode.
func (config *Config) DecodeMode() DecodeMode {
	return config.decodeMode
}

// SetDuplicateMapKeyMode sets how duplicate map keys are handled when decoding.
func (config *Config) SetDuplicateMapKeyMode(mode DuplicateMapKeyMode) {
	config.duplicateMapKeyMode = mode
}

// DuplicateMapKeyMode returns how duplicate map keys are handled when decoding.
func (config *Config) DuplicateMapKeyMode() DuplicateMapKeyMode {
	return config.duplicateMapKeyMode
}

// SetMaxNestedLevels sets the maximum nesting depth of arrays, maps, and tags accepted when decoding
// (RFC 8949 Section 10). A value less than 1 resets it to DefaultMaxNestedLevels.
func (config *Config) SetMaxNestedLevels(n int) {
	config.maxNestedLevels = n
}

// MaxNestedLevels returns the maximum nesting depth of arrays, maps, and tags accepted when decoding.
func (config *Config) MaxNestedLevels() int {
	if config.maxNestedLevels < 1 {
		return DefaultMaxNestedLevels
	}
	return config.maxNestedLevels
}

// SetUTF8ValidationEnabled sets whether text strings are validated as UTF-8 when encoding and decoding
// (RFC 8949 Section 5.3.1). It is enabled by default, and invalid text strings return an error
// wrapping ErrInvalidUTF8 with ErrEncode or ErrDecode.
func (config *Config) SetUTF8ValidationEnabled(flag bool) {
	config.utf8ValidationDisabled = !flag
}

// IsUTF8ValidationEnabled returns true if text strings are validated as UTF-8 when encoding and decoding.
func (config *Config) IsUTF8ValidationEnabled() bool {
	return !config.utf8ValidationDisabled
}
