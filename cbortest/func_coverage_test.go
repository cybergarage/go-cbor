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
	"time"

	"github.com/cybergarage/go-cbor/cbor"
)

func TestSpecificIntegerSizes(t *testing.T) {
	// Test specific byte size encodings
	tests := []struct {
		name  string
		value any
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
	data, _ := cbor.Marshal(map[string]any{
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
	if resultMap, ok := result.(map[any]any); ok {
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
		value any
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
	if resultArray, ok := result.([]any); ok {
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
	if resultMap, ok := result.(map[any]any); ok {
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

func TestNegativeIntegerEncoding(t *testing.T) {
	// Test negative integers that trigger nint32 and nint64 byte encodations
	tests := []struct {
		name  string
		value any
	}{
		{"NegInt32_1", int32(-256)},
		{"NegInt32_2", int32(-65536)},
		{"NegInt32_3", int32(-2147483648)},
		{"NegInt64_1", int64(-4294967296)},
		{"NegInt64_2", int64(-9223372036854775808)},
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

func TestByteStringConversions(t *testing.T) {
	// Test []byte to string and string to []byte conversions
	t.Run("ByteArrayToString", func(t *testing.T) {
		original := []byte("hello world")
		data, err := cbor.Marshal(original)
		if err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}

		var resultStr string
		err = cbor.UnmarshalTo(data, &resultStr)
		if err != nil {
			t.Fatalf("UnmarshalTo string failed: %v", err)
		}

		if resultStr != string(original) {
			t.Errorf("Expected '%s', got '%s'", string(original), resultStr)
		}
	})

	t.Run("ByteArrayToByteArray", func(t *testing.T) {
		original := []byte{0x01, 0x02, 0x03, 0x04}
		data, err := cbor.Marshal(original)
		if err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}

		var resultBytes []byte
		err = cbor.UnmarshalTo(data, &resultBytes)
		if err != nil {
			t.Fatalf("UnmarshalTo []byte failed: %v", err)
		}

		if !bytes.Equal(original, resultBytes) {
			t.Errorf("Expected %v, got %v", original, resultBytes)
		}
	})

	t.Run("StringToString", func(t *testing.T) {
		original := "test string"
		data, err := cbor.Marshal(original)
		if err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}

		var resultStr string
		err = cbor.UnmarshalTo(data, &resultStr)
		if err != nil {
			t.Fatalf("UnmarshalTo string failed: %v", err)
		}

		if resultStr != original {
			t.Errorf("Expected '%s', got '%s'", original, resultStr)
		}
	})
}

func TestTimeConversions(t *testing.T) {
	// Test time.Time conversions
	t.Run("TimeToString", func(t *testing.T) {
		original, err := time.Parse(time.RFC3339, "2024-03-21T20:04:00Z")
		if err != nil {
			t.Fatalf("Parse time failed: %v", err)
		}

		data, err := cbor.Marshal(original)
		if err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}

		var resultStr string
		err = cbor.UnmarshalTo(data, &resultStr)
		if err != nil {
			t.Fatalf("UnmarshalTo string failed: %v", err)
		}

		expectedStr := original.Format(time.RFC3339)
		if resultStr != expectedStr {
			t.Errorf("Expected '%s', got '%s'", expectedStr, resultStr)
		}
	})

	t.Run("TimeToTime", func(t *testing.T) {
		original, err := time.Parse(time.RFC3339, "2024-03-21T20:04:00Z")
		if err != nil {
			t.Fatalf("Parse time failed: %v", err)
		}

		data, err := cbor.Marshal(original)
		if err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}

		var resultTime time.Time
		err = cbor.UnmarshalTo(data, &resultTime)
		if err != nil {
			t.Fatalf("UnmarshalTo time.Time failed: %v", err)
		}

		if !original.Equal(resultTime) {
			t.Errorf("Expected '%v', got '%v'", original, resultTime)
		}
	})
}

func TestArrayToArrayConversions(t *testing.T) {
	// Test array to array conversions with different target types
	t.Run("SliceToSlice", func(t *testing.T) {
		original := []int{1, 2, 3, 4, 5}
		data, err := cbor.Marshal(original)
		if err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}

		var result []int
		err = cbor.UnmarshalTo(data, &result)
		if err != nil {
			t.Fatalf("UnmarshalTo slice failed: %v", err)
		}

		if len(result) != len(original) {
			t.Errorf("Length mismatch: expected %d, got %d", len(original), len(result))
		}

		for i, v := range original {
			if int64(result[i]) != int64(v) {
				t.Errorf("Index %d: expected %d, got %d", i, v, result[i])
			}
		}
	})

	t.Run("SliceOfStrings", func(t *testing.T) {
		original := []string{"hello", "world", "test"}
		data, err := cbor.Marshal(original)
		if err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}

		var result []string
		err = cbor.UnmarshalTo(data, &result)
		if err != nil {
			t.Fatalf("UnmarshalTo string slice failed: %v", err)
		}

		if len(result) != len(original) {
			t.Errorf("Length mismatch: expected %d, got %d", len(original), len(result))
		}

		for i, v := range original {
			if result[i] != v {
				t.Errorf("Index %d: expected %s, got %s", i, v, result[i])
			}
		}
	})

	t.Run("EmptySlice", func(t *testing.T) {
		original := []int{}
		data, err := cbor.Marshal(original)
		if err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}

		var result []int
		err = cbor.UnmarshalTo(data, &result)
		if err != nil {
			t.Fatalf("UnmarshalTo empty slice failed: %v", err)
		}

		if len(result) != 0 {
			t.Errorf("Expected empty slice, got length %d", len(result))
		}
	})
}

func TestMapToMapConversions(t *testing.T) {
	// Test map to map conversions to improve coverage
	t.Run("MapWithIntKeys", func(t *testing.T) {
		original := map[int]string{1: "one", 2: "two", 3: "three"}
		data, err := cbor.Marshal(original)
		if err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}

		result := make(map[int]string)
		err = cbor.UnmarshalTo(data, result)
		if err != nil {
			t.Fatalf("UnmarshalTo map failed: %v", err)
		}

		if len(result) != len(original) {
			t.Errorf("Length mismatch: expected %d, got %d", len(original), len(result))
		}

		for k, v := range original {
			if result[k] != v {
				t.Errorf("Key %d: expected %s, got %s", k, v, result[k])
			}
		}
	})

	t.Run("MapWithMixedTypes", func(t *testing.T) {
		original := map[string]any{"number": 42, "text": "hello", "flag": true}
		data, err := cbor.Marshal(original)
		if err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}

		result := make(map[string]any)
		err = cbor.UnmarshalTo(data, result)
		if err != nil {
			t.Fatalf("UnmarshalTo map failed: %v", err)
		}

		if len(result) != len(original) {
			t.Errorf("Length mismatch: expected %d, got %d", len(original), len(result))
		}
	})
}

func TestAdditionalStructEncoding(t *testing.T) {
	// Test struct encoding to improve encodeStruct coverage
	type Person struct {
		Name   string
		Age    int
		Active bool
	}

	t.Run("SimpleStruct", func(t *testing.T) {
		original := Person{Name: "John", Age: 30, Active: true}
		data, err := cbor.Marshal(original)
		if err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}

		var result Person
		err = cbor.UnmarshalTo(data, &result)
		if err != nil {
			t.Fatalf("UnmarshalTo failed: %v", err)
		}

		if result.Name != original.Name || result.Age != original.Age || result.Active != original.Active {
			t.Errorf("Struct mismatch: expected %+v, got %+v", original, result)
		}
	})

	t.Run("NestedStruct", func(t *testing.T) {
		type Address struct {
			City    string
			ZipCode int
		}
		type Employee struct {
			Name    string
			Address Address
		}

		original := Employee{
			Name:    "Jane",
			Address: Address{City: "Tokyo", ZipCode: 12345},
		}
		data, err := cbor.Marshal(original)
		if err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}

		var result Employee
		err = cbor.UnmarshalTo(data, &result)
		if err != nil {
			t.Fatalf("UnmarshalTo failed: %v", err)
		}

		if result.Name != original.Name || result.Address.City != original.Address.City {
			t.Errorf("Nested struct mismatch: expected %+v, got %+v", original, result)
		}
	})
}

func TestComplexDataStructures(t *testing.T) {
	// Test more complex structures to improve overall coverage
	t.Run("SliceOfMaps", func(t *testing.T) {
		original := []map[string]int{
			{"a": 1, "b": 2},
			{"c": 3, "d": 4},
		}
		data, err := cbor.Marshal(original)
		if err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}

		result, err := cbor.Unmarshal(data)
		if err != nil {
			t.Fatalf("Unmarshal failed: %v", err)
		}

		if result == nil {
			t.Errorf("Expected non-nil result")
		}
	})

	t.Run("MapOfSlices", func(t *testing.T) {
		original := map[string][]int{
			"evens": {2, 4, 6},
			"odds":  {1, 3, 5},
		}
		data, err := cbor.Marshal(original)
		if err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}

		result, err := cbor.Unmarshal(data)
		if err != nil {
			t.Fatalf("Unmarshal failed: %v", err)
		}

		if result == nil {
			t.Errorf("Expected non-nil result")
		}
	})

	t.Run("BoolValues", func(t *testing.T) {
		tests := []bool{true, false}
		for _, original := range tests {
			data, err := cbor.Marshal(original)
			if err != nil {
				t.Fatalf("Marshal %v failed: %v", original, err)
			}

			result, err := cbor.Unmarshal(data)
			if err != nil {
				t.Fatalf("Unmarshal %v failed: %v", original, err)
			}

			if result != original {
				t.Errorf("Expected %v, got %v", original, result)
			}
		}
	})

	t.Run("NilValue", func(t *testing.T) {
		data, err := cbor.Marshal(nil)
		if err != nil {
			t.Fatalf("Marshal nil failed: %v", err)
		}

		result, err := cbor.Unmarshal(data)
		if err != nil {
			t.Fatalf("Unmarshal nil failed: %v", err)
		}

		if result != nil {
			t.Errorf("Expected nil, got %v", result)
		}
	})
}
