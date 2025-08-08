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
	"testing"

	"github.com/cybergarage/go-cbor/cbor"
)

func TestSpecificIntegerSizes(t *testing.T) {
	// Test specific byte size encodings
	tests := []struct {
		name  string
		value interface{}
	}{
		// Test values that trigger specific byte sizes
		{"Int8Max", int8(127)},
		{"Int8Min", int8(-128)},
		{"Uint8Max", uint8(255)},
		{"Int16Max", int16(32767)},
		{"Int16Min", int16(-32768)},
		{"Uint16Max", uint16(65535)},
		{"Int32Max", int32(2147483647)},
		{"Int32Min", int32(-2147483648)},
		{"Uint32Max", uint32(4294967295)},
		{"Int64Max", int64(9223372036854775807)},
		{"Int64Min", int64(-9223372036854775808)},
		{"Uint64Max", uint64(18446744073709551615)},

		// Values that require specific byte encodings
		{"Value23", int(23)},       // Should use single byte
		{"Value24", int(24)},       // Should use uint8
		{"Value255", int(255)},     // Should use uint8
		{"Value256", int(256)},     // Should use uint16
		{"Value65535", int(65535)}, // Should use uint16
		{"Value65536", int(65536)}, // Should use uint32
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

func TestDirectDecoderUsage(t *testing.T) {
	// Test direct decoder usage to hit more code paths
	data, _ := cbor.Marshal(map[string]interface{}{
		"number": 42,
		"string": "hello",
		"array":  []int{1, 2, 3},
		"nested": map[string]int{"inner": 99},
	})

	decoder := cbor.NewDecoder(bytes.NewReader(data))
	result, err := decoder.Decode()
	if err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	// Verify the structure
	if resultMap, ok := result.(map[interface{}]interface{}); ok {
		if resultMap["number"] != int64(42) {
			t.Errorf("Expected number 42, got %v", resultMap["number"])
		}
		if resultMap["string"] != "hello" {
			t.Errorf("Expected string 'hello', got %v", resultMap["string"])
		}
	} else {
		t.Errorf("Expected map result, got %T", result)
	}
}

func TestLargeDataStructures(t *testing.T) {
	// Test large arrays to trigger different encoding paths
	largeArray := make([]int, 1000)
	for i := range largeArray {
		largeArray[i] = i
	}

	data, err := cbor.Marshal(largeArray)
	if err != nil {
		t.Fatalf("Marshal large array failed: %v", err)
	}

	result, err := cbor.Unmarshal(data)
	if err != nil {
		t.Fatalf("Unmarshal large array failed: %v", err)
	}

	if err := deepEqual(largeArray, result); err != nil {
		t.Errorf("Large array values not equal: %v", err)
	}

	// Test large map
	largeMap := make(map[string]int)
	for i := range 100 {
		largeMap[string(rune('a'+i%26))+string(rune('0'+i/26))] = i
	}

	data, err = cbor.Marshal(largeMap)
	if err != nil {
		t.Fatalf("Marshal large map failed: %v", err)
	}

	result, err = cbor.Unmarshal(data)
	if err != nil {
		t.Fatalf("Unmarshal large map failed: %v", err)
	}

	if err := deepEqual(largeMap, result); err != nil {
		t.Errorf("Large map values not equal: %v", err)
	}
}

func TestBoundaryValues(t *testing.T) {
	// Test boundary values for different integer types
	tests := []struct {
		name  string
		value interface{}
	}{
		{"Zero", 0},
		{"One", 1},
		{"TwentyThree", 23},   // Boundary for direct encoding
		{"TwentyFour", 24},    // Requires uint8
		{"FiftyNine", 59},     // Maximum for direct encoding
		{"Sixty", 60},         // Requires uint8
		{"TwoFiftyFive", 255}, // Maximum uint8
		{"TwoFiftySix", 256},  // Requires uint16
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

			// Test different byte lengths
			if len(data) == 0 {
				t.Errorf("Encoded data is empty for value %v", tt.value)
			}
		})
	}
}

func TestArrayConversionFunctions(t *testing.T) {
	// Test array to interface{} array conversion
	intArray := [3]int{1, 2, 3}
	data, err := cbor.Marshal(intArray)
	if err != nil {
		t.Fatalf("Marshal array failed: %v", err)
	}

	result, err := cbor.Unmarshal(data)
	if err != nil {
		t.Fatalf("Unmarshal array failed: %v", err)
	}

	// Verify conversion worked
	if resultArray, ok := result.([]interface{}); ok {
		if len(resultArray) != 3 {
			t.Errorf("Expected array length 3, got %d", len(resultArray))
		}
		for i, v := range resultArray {
			if v != int64(i+1) {
				t.Errorf("Array element %d: expected %d, got %v", i, i+1, v)
			}
		}
	} else {
		t.Errorf("Expected []interface{}, got %T", result)
	}
}

func TestMapConversionFunctions(t *testing.T) {
	// Test map conversion with various key types
	stringKeyMap := map[string]int{"a": 1, "b": 2}
	data, err := cbor.Marshal(stringKeyMap)
	if err != nil {
		t.Fatalf("Marshal map failed: %v", err)
	}

	result, err := cbor.Unmarshal(data)
	if err != nil {
		t.Fatalf("Unmarshal map failed: %v", err)
	}

	// Verify conversion worked
	if resultMap, ok := result.(map[interface{}]interface{}); ok {
		if len(resultMap) != 2 {
			t.Errorf("Expected map length 2, got %d", len(resultMap))
		}
		if resultMap["a"] != int64(1) || resultMap["b"] != int64(2) {
			t.Errorf("Map values incorrect: %v", resultMap)
		}
	} else {
		t.Errorf("Expected map[interface{}]interface{}, got %T", result)
	}
}
