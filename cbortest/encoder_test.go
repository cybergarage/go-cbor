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
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/cybergarage/go-cbor/cbor"
)

func TestEncoder(t *testing.T) {
	t.Run("RFC-8949", func(t *testing.T) {
		t.Run("AppendixA", func(t *testing.T) {
			tests := []struct {
				value    any
				expected string
			}{
				{value: uint8(0), expected: "00"},
				{value: uint8(1), expected: "01"},
				{value: uint8(10), expected: "0a"},
				{value: uint8(23), expected: "17"},
				{value: uint8(24), expected: "1818"},
				{value: uint8(25), expected: "1819"},
				{value: uint8(100), expected: "1864"},
				{value: uint16(1000), expected: "1903e8"},
				{value: uint32(1000000), expected: "1a000f4240"},
				{value: uint64(1000000000000), expected: "1b000000e8d4a51000"},
				{value: uint64(18446744073709551615), expected: "1bffffffffffffffff"},
				{value: uint(1000000000000), expected: "1b000000e8d4a51000"},
				{value: uint(18446744073709551615), expected: "1bffffffffffffffff"},
				// {value: uint(18446744073709551616), expected: "c249010000000000000000"},
				// {value: int(-18446744073709551616), expected: "3bffffffffffffffff"},
				// {value: int(-18446744073709551617), expected: "c349010000000000000000"},
				{value: int8(-1), expected: "20"},
				{value: int8(-10), expected: "29"},
				{value: int8(-100), expected: "3863"},
				{value: int16(-1000), expected: "3903e7"},
				// {value: float16(0.0), expected: "f90000"},
				// {value: float16(0.0), expected: "f98000"},
				// {value: float16(1.0), expected: "f93c00"},
				{value: float64(1.1), expected: "fb3ff199999999999a"},
				// {value: float16(1.5), expected: "f93e00"},
				// {value: float16(65504.0), expected: "f97bff"},
				{value: float32(100000.0), expected: "fa47c35000"},
				// {value: float64(3.4028234663852886e+38), expected: "fa7f7fffff"},
				{value: float64(1.0e+300), expected: "fb7e37e43c8800759c"},
				// {value: float16(5.960464477539063e-8), expected: "f90001"},
				// {value: float16(0.00006103515625), expected: "f90400"},
				// {value: float16(-4.0), expected: "f9c400"},
				{value: float64(-4.1), expected: "fbc010666666666666"},
				// {value: float64(math.Inf), expected: "f97c00"},
				// {value: float64(math.NaN), expected: "f97e00"},
				// {value: float64(-math.Inf), expected: "f9fc00"},
				// {value: float64(math.Inf), expected: "fa7f800000"},
				// {value: float64(math.NaN), expected: "fa7fc00000"},
				// {value: float64(-math.Inf), expected: "faff800000"},
				// {value: float64(math.Inf), expected: "fb7ff0000000000000"},
				// {value: float64(math.NaN), expected: "fb7ff8000000000000"},
				// {value: float64(-math.Inf), expected: "fbfff0000000000000"},
			}
			for _, test := range tests {
				t.Run(fmt.Sprintf("%T/%v=>%s", test.value, test.value, test.expected), func(t *testing.T) {
					var writer bytes.Buffer
					encoder := cbor.NewEncoder(&writer)
					err := encoder.Encode(test.value)
					if err != nil {
						t.Errorf("%v (%s)", test.value, err.Error())
						return
					}
					encoded := hex.EncodeToString(writer.Bytes())
					if encoded != test.expected {
						t.Errorf("%v (%T) != %v (%T)", encoded, encoded, test.expected, test.expected)
					}
				})
			}
		})
	})
}
