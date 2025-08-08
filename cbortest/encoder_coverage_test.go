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
	"fmt"
	"testing"

	"github.com/cybergarage/go-cbor/cbor"
)

func TestByteStringEncoding(t *testing.T) {
	tests := []struct {
		name string
		data []byte
	}{
		{"EmptyBytes", []byte{}},
		{"SingleByte", []byte{0x42}},
		{"SmallBytes", []byte{0x01, 0x02, 0x03}},
		{"LargeBytes", make([]byte, 1000)},
		{"NilBytes", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := cbor.Marshal(tt.data)
			if err != nil {
				t.Fatalf("Marshal failed: %v", err)
			}

			result, err := cbor.Unmarshal(data)
			if err != nil {
				t.Fatalf("Unmarshal failed: %v", err)
			}

			if err := deepEqual(tt.data, result); err != nil {
				t.Errorf("Values not equal: %v", err)
			}
		})
	}
}

func TestTextStringEncoding(t *testing.T) {
	tests := []struct {
		name string
		text string
	}{
		{"EmptyString", ""},
		{"SimpleString", "hello"},
		{"UnicodeString", "こんにちは"},
		{"LongString", string(make([]byte, 1000))},
		{"SpecialChars", "\n\t\r"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := cbor.Marshal(tt.text)
			if err != nil {
				t.Fatalf("Marshal failed: %v", err)
			}

			result, err := cbor.Unmarshal(data)
			if err != nil {
				t.Fatalf("Unmarshal failed: %v", err)
			}

			if err := deepEqual(tt.text, result); err != nil {
				t.Errorf("Values not equal: %v", err)
			}
		})
	}
}

func TestPrimitiveTypesEncoding(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
	}{
		{"BoolTrue", true},
		{"BoolFalse", false},
		{"Nil", nil},
		{"UintZero", uint(0)},
		{"UintMax", ^uint(0)},
		{"IntZero", int(0)},
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

func TestInvalidCBORData(t *testing.T) {
	invalidData := [][]byte{
		{0xFF}, // Invalid major type
		{0x1F}, // Invalid additional info
		{0x18}, // Missing byte after uint8 indicator
		{0x19}, // Missing bytes after uint16 indicator
		{0x1A}, // Missing bytes after uint32 indicator
		{0x1B}, // Missing bytes after uint64 indicator
	}

	for i, data := range invalidData {
		t.Run(fmt.Sprintf("InvalidData%d", i), func(t *testing.T) {
			_, err := cbor.Unmarshal(data)
			if err == nil {
				t.Errorf("Expected error for invalid data %v", data)
			}
		})
	}
}

func TestEncoderBufferWriting(t *testing.T) {
	var buf bytes.Buffer
	encoder := cbor.NewEncoder(&buf)

	// Test multiple writes to the same buffer
	values := []interface{}{
		42,
		"hello",
		[]int{1, 2, 3},
		map[string]int{"key": 123},
	}

	totalBytes := 0
	for _, v := range values {
		initialSize := buf.Len()
		err := encoder.Encode(v)
		if err != nil {
			t.Fatalf("Encode failed: %v", err)
		}
		bytesWritten := buf.Len() - initialSize
		if bytesWritten <= 0 {
			t.Errorf("No bytes written for value %v", v)
		}
		totalBytes += bytesWritten
	}

	if buf.Len() != totalBytes {
		t.Errorf("Buffer size mismatch: expected %d, got %d", totalBytes, buf.Len())
	}
}
