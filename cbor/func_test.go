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

import (
	"bytes"
	"math"
	"testing"
)

func TestEncodeDecodeFunc(t *testing.T) {
	t.Run("int8", func(t *testing.T) {
		testValues := []int8{
			math.MinInt8,
			math.MinInt8 / 2,
			0,
			math.MaxInt8 / 2,
			math.MaxInt8,
		}
		for _, testVal := range testValues {
			var w bytes.Buffer
			err := writeInt8Bytes(&w, testVal)
			if err != nil {
				t.Error(err)
				continue
			}
			reader := bytes.NewReader(w.Bytes())
			val, err := readInt8Bytes(reader)
			if err != nil {
				t.Error(err)
				continue
			}
			if val != testVal {
				t.Errorf("%d != %d", val, testVal)
			}
		}
	})
	t.Run("uint8", func(t *testing.T) {
		testValues := []uint8{
			0,
			math.MaxUint8 / 2,
			math.MaxUint8,
		}
		for _, testVal := range testValues {
			testBytes := appendUint8Bytes([]byte{}, testVal)
			val, _, err := readUint8Bytes(testBytes)
			if err != nil {
				t.Error(err)
				continue
			}
			if val != testVal {
				t.Errorf("%d != %d", val, testVal)
			}
		}
	})
	t.Run("int16", func(t *testing.T) {
		testValues := []int16{
			math.MinInt16,
			math.MinInt16 / 2,
			0,
			math.MaxInt16 / 2,
			math.MaxInt16,
		}
		for _, testVal := range testValues {
			testBytes := appendInt16Bytes([]byte{}, testVal)
			val, _, err := readInt16Bytes(testBytes)
			if err != nil {
				t.Error(err)
				continue
			}
			if val != testVal {
				t.Errorf("%d != %d", val, testVal)
			}
		}
	})
	t.Run("uint16", func(t *testing.T) {
		testValues := []uint16{
			0,
			1,
			math.MaxUint16 / 2,
			math.MaxUint16,
		}
		for _, testVal := range testValues {
			testBytes := appendUint16Bytes([]byte{}, testVal)
			val, _, err := readUint16Bytes(testBytes)
			if err != nil {
				t.Error(err)
				continue
			}
			if val != testVal {
				t.Errorf("%d != %d", val, testVal)
			}
		}
	})
	t.Run("int32", func(t *testing.T) {
		testValues := []int32{
			math.MinInt32,
			math.MinInt32 / 2,
			0,
			math.MaxInt32 / 2,
			math.MaxInt32,
		}
		for _, testVal := range testValues {
			testBytes := appendInt32Bytes([]byte{}, testVal)
			val, _, err := readInt32Bytes(testBytes)
			if err != nil {
				t.Error(err)
				continue
			}
			if val != testVal {
				t.Errorf("%d != %d", val, testVal)
			}
		}
	})
	t.Run("uint32", func(t *testing.T) {
		testValues := []uint32{
			0,
			1,
			math.MaxUint32 / 2,
			math.MaxUint32,
		}
		for _, testVal := range testValues {
			testBytes := appendUint32Bytes([]byte{}, testVal)
			val, _, err := readUint32Bytes(testBytes)
			if err != nil {
				t.Error(err)
				continue
			}
			if val != testVal {
				t.Errorf("%d != %d", val, testVal)
			}
		}
	})
	t.Run("int64", func(t *testing.T) {
	})
	t.Run("uint64", func(t *testing.T) {
		testValues := []uint64{
			0,
			1,
			math.MaxUint64 / 2,
			math.MaxUint64,
		}
		for _, testVal := range testValues {
			testBytes := appendUint64Bytes([]byte{}, testVal)
			val, _, err := readUint64Bytes(testBytes)
			if err != nil {
				t.Error(err)
				continue
			}
			if val != testVal {
				t.Errorf("%d != %d", val, testVal)
			}
		}
	})
}
