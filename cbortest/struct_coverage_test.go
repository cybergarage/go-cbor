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
	"testing"
	"time"

	"github.com/cybergarage/go-cbor/cbor"
)

type TestStruct struct {
	Name   string
	Age    int
	Active bool
}

type NestedStruct struct {
	ID      int
	Details TestStruct
	Tags    []string
}

type EmbeddedStruct struct {
	TestStruct

	Extra string
}

func TestStructEncoding(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
	}{
		{
			"SimpleStruct",
			TestStruct{
				Name:   "John",
				Age:    30,
				Active: true,
			},
		},
		{
			"NestedStruct",
			NestedStruct{
				ID: 123,
				Details: TestStruct{
					Name:   "Jane",
					Age:    25,
					Active: false,
				},
				Tags: []string{"tag1", "tag2"},
			},
		},
		{
			"EmbeddedStruct",
			EmbeddedStruct{
				TestStruct: TestStruct{
					Name:   "Bob",
					Age:    40,
					Active: true,
				},
				Extra: "additional",
			},
		},
		{
			"EmptyStruct",
			struct{}{},
		},
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

			// Check that the result is a map with the expected fields
			resultMap, ok := result.(map[interface{}]interface{})
			if !ok {
				t.Errorf("Expected map result for struct, got %T", result)
				return
			}

			// Basic validation that the struct was encoded correctly
			// Empty structs are valid and should encode to empty maps
			if tt.name != "EmptyStruct" && len(resultMap) == 0 {
				t.Errorf("Non-empty struct encoded to empty map")
			}
		})
	}
}

func TestTimeEncoding(t *testing.T) {
	now := time.Now().Truncate(time.Second) // Remove nanoseconds for comparison
	data, err := cbor.Marshal(now)
	if err != nil {
		t.Fatalf("Marshal time failed: %v", err)
	}

	result, err := cbor.Unmarshal(data)
	if err != nil {
		t.Fatalf("Unmarshal time failed: %v", err)
	}

	// Check that a time value was successfully encoded and decoded
	if result == nil {
		t.Error("Time value became nil after marshal/unmarshal")
	}
}

func TestUnmarshalToStruct(t *testing.T) {
	original := TestStruct{
		Name:   "Alice",
		Age:    28,
		Active: true,
	}

	data, err := cbor.Marshal(original)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var result TestStruct
	err = cbor.UnmarshalTo(data, &result)
	if err != nil {
		t.Fatalf("UnmarshalTo failed: %v", err)
	}

	if result.Name != original.Name || result.Age != original.Age || result.Active != original.Active {
		t.Errorf("Struct fields not equal: got %+v, want %+v", result, original)
	}
}

func TestArrayToStructErrors(t *testing.T) {
	// Test unmarshaling array to struct (should fail)
	arrayData, _ := cbor.Marshal([3]int{1, 2, 3})
	var s TestStruct
	err := cbor.UnmarshalTo(arrayData, &s)
	if err == nil {
		t.Error("Expected error when unmarshaling array to struct")
	}

	// Test unmarshaling struct to array (should fail)
	structData, _ := cbor.Marshal(TestStruct{Name: "test", Age: 25, Active: true})
	var arr [3]int
	err = cbor.UnmarshalTo(structData, &arr)
	if err == nil {
		t.Error("Expected error when unmarshaling struct to array")
	}
}

func TestSliceOperations(t *testing.T) {
	// Test different slice types
	intSlice := []int{1, 2, 3, 4, 5}
	data, err := cbor.Marshal(intSlice)
	if err != nil {
		t.Fatalf("Marshal slice failed: %v", err)
	}

	result, err := cbor.Unmarshal(data)
	if err != nil {
		t.Fatalf("Unmarshal slice failed: %v", err)
	}

	if err := deepEqual(intSlice, result); err != nil {
		t.Errorf("Slice values not equal: %v", err)
	}

	// Test unmarshal to specific slice type
	var resultSlice []int
	err = cbor.UnmarshalTo(data, &resultSlice)
	if err != nil {
		t.Fatalf("UnmarshalTo slice failed: %v", err)
	}

	if len(resultSlice) != len(intSlice) {
		t.Errorf("Slice lengths not equal: got %d, want %d", len(resultSlice), len(intSlice))
	}

	for i, v := range intSlice {
		if resultSlice[i] != v {
			t.Errorf("Slice element %d not equal: got %d, want %d", i, resultSlice[i], v)
		}
	}
}
