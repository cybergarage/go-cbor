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
	"strings"
	"testing"

	"github.com/cybergarage/go-cbor/cbor"
)

func TestErrors(t *testing.T) {
	// Test error creation for unsupported types
	tests := []struct {
		name     string
		testFunc func() error
		contains string
	}{
		{
			name: "NotSupportedMajorType",
			testFunc: func() error {
				// Trigger by encoding an unsupported type
				data := []byte{0xFF} // Invalid major type
				_, err := cbor.Unmarshal(data)
				return err
			},
			contains: "not supported",
		},
		{
			name: "NotSupportedNativeType",
			testFunc: func() error {
				// Trigger by encoding a cmplex type
				cmplex := complex(1, 2)
				_, err := cbor.Marshal(cmplex)
				return err
			},
			contains: "not supported",
		},
		{
			name: "UnmarshalDataTypes",
			testFunc: func() error {
				// Trigger by unmarshaling to incompatible type
				data, _ := cbor.Marshal([]int{1, 2, 3})
				var intVal int
				err := cbor.UnmarshalTo(data, &intVal)
				return err
			},
			contains: "unmarshal error",
		},
		{
			name: "UnmarshalArraySize",
			testFunc: func() error {
				// Trigger by unmarshaling array to smaller array
				data, _ := cbor.Marshal([3]int{1, 2, 3})
				var smallArray [2]int
				err := cbor.UnmarshalTo(data, &smallArray)
				return err
			},
			contains: "unmarshal error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.testFunc()
			if err == nil {
				t.Errorf("Expected error for %s, got nil", tt.name)
				return
			}
			if !strings.Contains(err.Error(), tt.contains) {
				t.Errorf("Expected error to contain '%s', got: %s", tt.contains, err.Error())
			}
		})
	}
}

func TestReflectTypeErrors(t *testing.T) {
	// Test reflect.Value related errors through complex unmarshaling scenarios
	data, _ := cbor.Marshal(map[string]interface{}{"key": 42})
	var stringVar string
	err := cbor.UnmarshalTo(data, &stringVar)
	if err == nil {
		t.Error("Expected error when unmarshaling map to string")
	}
}
