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

package cbortest

import (
	"bytes"
	"math"
	"testing"

	"github.com/cybergarage/go-cbor/cbor"
)

func TestNegativeIntegers(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
	}{
		{"NegativeInt8", int8(-1)},
		{"NegativeInt8Min", int8(math.MinInt8)},
		{"NegativeInt16", int16(-1000)},
		{"NegativeInt16Min", int16(math.MinInt16)},
		{"NegativeInt32", int32(-100000)},
		{"NegativeInt32Min", int32(math.MinInt32)},
		{"NegativeInt64", int64(-10000000000)},
		{"NegativeInt64Min", int64(math.MinInt64)},
		{"NegativeInt", int(-42)},
		{"NegativeIntMin", int(math.MinInt)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := cbor.Marshal(tt.value)
			if err != nil {
				t.Fatalf("Marshal failed: %v", err)
			}

			result, err := cbor.Unmarshal(data)
			if err != nil {
				t.Fatalf("Unmarshal failed: %v", err)
			}

			if err := deepEqual(tt.value, result); err != nil {
				t.Errorf("Values not equal: %v", err)
			}
		})
	}
}

func TestEdgeCaseNumbers(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
	}{
		{"MaxFloat32", float32(math.MaxFloat32)},
		{"SmallestFloat32", float32(math.SmallestNonzeroFloat32)},
		{"MaxFloat64", float64(math.MaxFloat64)},
		{"SmallestFloat64", float64(math.SmallestNonzeroFloat64)},
		{"PositiveInf32", float32(math.Inf(1))},
		{"NegativeInf32", float32(math.Inf(-1))},
		{"PositiveInf64", math.Inf(1)},
		{"NegativeInf64", math.Inf(-1)},
		{"NaN32", float32(math.NaN())},
		{"NaN64", math.NaN()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := cbor.Marshal(tt.value)
			if err != nil {
				t.Fatalf("Marshal failed: %v", err)
			}

			result, err := cbor.Unmarshal(data)
			if err != nil {
				t.Fatalf("Unmarshal failed: %v", err)
			}

			// Special handling for NaN values
			switch v := tt.value.(type) {
			case float32:
				if math.IsNaN(float64(v)) {
					if !math.IsNaN(float64(result.(float32))) {
						t.Errorf("Expected NaN, got %v", result)
					}
				} else {
					if err := deepEqual(tt.value, result); err != nil {
						t.Errorf("Values not equal: %v", err)
					}
				}
			case float64:
				if math.IsNaN(v) {
					if !math.IsNaN(result.(float64)) {
						t.Errorf("Expected NaN, got %v", result)
					}
				} else {
					if err := deepEqual(tt.value, result); err != nil {
						t.Errorf("Values not equal: %v", err)
					}
				}
			default:
				if err := deepEqual(tt.value, result); err != nil {
					t.Errorf("Values not equal: %v", err)
				}
			}
		})
	}
}

func TestSortedMapEncoding(t *testing.T) {
	// Test with string keys (should work with sorting)
	stringMap := map[string]int{
		"zebra":  1,
		"apple":  2,
		"banana": 3,
	}

	// Create encoder with sorting enabled
	var buf bytes.Buffer
	encoder := cbor.NewEncoder(&buf)
	encoder.SetMapSortEnabled(true)

	err := encoder.Encode(stringMap)
	if err != nil {
		t.Fatalf("Failed to encode string map with sort: %v", err)
	}

	result, err := cbor.Unmarshal(buf.Bytes())
	if err != nil {
		t.Fatalf("Failed to unmarshal sorted map: %v", err)
	}

	if err := deepEqual(stringMap, result); err != nil {
		t.Errorf("Sorted map values not equal: %v", err)
	}

	// Test with complex keys (should fail when sorting is enabled)
	complexMap := map[complex64]int{
		complex(1, 2): 1,
		complex(3, 4): 2,
	}

	buf.Reset()
	err = encoder.Encode(complexMap)
	if err == nil {
		t.Error("Expected error when encoding map with complex keys for sorting")
	}
}

func TestUnmarshalToSpecificTypes(t *testing.T) {
	// Test unmarshaling basic types to specific variables
	stringData, _ := cbor.Marshal("hello")
	var str string
	err := cbor.UnmarshalTo(stringData, &str)
	if err != nil {
		t.Fatalf("Failed to unmarshal to string: %v", err)
	}
	if str != "hello" {
		t.Errorf("Expected 'hello', got '%s'", str)
	}

	// Test unmarshaling number to different integer types
	intData, _ := cbor.Marshal(42)
	var i64 int64
	err = cbor.UnmarshalTo(intData, &i64)
	if err != nil {
		t.Fatalf("Failed to unmarshal to int64: %v", err)
	}
	if i64 != 42 {
		t.Errorf("Expected 42, got %d", i64)
	}

	// Test array to slice
	arrayData, _ := cbor.Marshal([]int{1, 2, 3})
	var slice []int
	err = cbor.UnmarshalTo(arrayData, &slice)
	if err != nil {
		t.Fatalf("Failed to unmarshal to slice: %v", err)
	}
	if len(slice) != 3 || slice[0] != 1 || slice[1] != 2 || slice[2] != 3 {
		t.Errorf("Expected [1, 2, 3], got %v", slice)
	}
}
