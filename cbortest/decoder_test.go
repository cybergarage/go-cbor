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
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/cybergarage/go-cbor/cbor"
)

func TestDecoder(t *testing.T) {
	dencoderTest := func(t *testing.T, encoded string, expected any) {
		t.Helper()
		testBytes, err := hex.DecodeString(encoded)
		if err != nil {
			return
		}
		decoder := cbor.NewDecoder(bytes.NewReader(testBytes))
		v, err := decoder.Decode()
		if err != nil {
			if errors.Is(err, cbor.ErrNotSupported) {
				t.Skipf("%v (%s)", encoded, err.Error())
			} else {
				t.Errorf("%v (%s)", encoded, err.Error())
			}
			return
		}
		err = deepEqual(v, expected)
		if err != nil {
			t.Errorf("%v (%T) != %v (%T)", v, v, expected, expected)
			return
		}
	}

	t.Run("RFC-8949", func(t *testing.T) {
		t.Run("AppendixA", func(t *testing.T) {
			t20120321, err := time.Parse(time.RFC3339, "2013-03-21T20:04:00Z")
			if err != nil {
				t.Error(err)
				return
			}
			tests := []struct {
				encoded  string
				expected any
			}{
				{encoded: "00", expected: int8(0)},
				{encoded: "01", expected: int8(1)},
				{encoded: "0a", expected: int8(10)},
				{encoded: "17", expected: int8(23)},
				{encoded: "1818", expected: int8(24)},
				{encoded: "1819", expected: int8(25)},
				{encoded: "1864", expected: int8(100)},
				{encoded: "1903e8", expected: int16(1000)},
				{encoded: "1a000f4240", expected: int32(1000000)},
				{encoded: "1b000000e8d4a51000", expected: int64(1000000000000)},
				{encoded: "1bffffffffffffffff", expected: uint64(18446744073709551615)},
				// {encoded: "c249010000000000000000", expected: uint64(18446744073709551616)},
				// {encoded: "3bffffffffffffffff", expected: int64(-18446744073709551616)},
				// {encoded: "c349010000000000000000", expected: int64(-18446744073709551617)},
				{encoded: "20", expected: int8(-1)},
				{encoded: "29", expected: int8(-10)},
				{encoded: "3863", expected: int8(-100)},
				{encoded: "3903e7", expected: int16(-1000)},
				{encoded: "f90000", expected: float64(0.0)},
				{encoded: "f98000", expected: float64(0.0)},
				{encoded: "f93c00", expected: float64(1.0)},
				{encoded: "fb3ff199999999999a", expected: float64(1.1)},
				{encoded: "f93e00", expected: float64(1.5)},
				{encoded: "f97bff", expected: float64(65504.0)},
				{encoded: "fa47c35000", expected: float32(100000.0)},
				{encoded: "fa7f7fffff", expected: float32(3.4028234663852886e+38)},
				{encoded: "fb7e37e43c8800759c", expected: float64(1.0e+300)},
				{encoded: "f90001", expected: float64(5.960464477539063e-8)},
				{encoded: "f90400", expected: float64(0.00006103515625)},
				{encoded: "f9c400", expected: float64(-4.0)},
				{encoded: "fbc010666666666666", expected: float64(-4.1)},
				// {encoded: "f97c00", expected: math.Inf},
				// {encoded: "f97e00", expected: math.NaN},
				// {encoded: "f9fc00", expected: -math.Inf},
				// {encoded: "fa7f800000", expected: math.Inf},
				// {encoded: "fa7fc00000", expected: math.NaN},
				// {encoded: "faff800000", expected: -math.Inf},
				// {encoded: "fb7ff0000000000000", expected: math.Inf},
				// {encoded: "fb7ff8000000000000", expected: math.NaN},
				// {encoded: "fbfff0000000000000", expected: -math.Inf},
				{encoded: "f4", expected: false},
				{encoded: "f5", expected: true},
				{encoded: "f6", expected: nil},
				{encoded: "c074323031332d30332d32315432303a30343a30305a", expected: t20120321},
				{encoded: "60", expected: ""},
				{encoded: "6161", expected: "a"},
				{encoded: "6449455446", expected: "IETF"},
				{encoded: "62225c", expected: "\"\\"},
				{encoded: "62c3bc", expected: "\u00fc"},
				{encoded: "63e6b0b4", expected: "\u6c34"},
				// {encoded: "64f0908591", expected: "\ud800\udd51"},
				{encoded: "80", expected: []int8{}},
				{encoded: "83010203", expected: []int8{1, 2, 3}},
				{encoded: "98190102030405060708090a0b0c0d0e0f101112131415161718181819", expected: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}},
				{encoded: "a0", expected: map[any]any{}},
				{encoded: "a201020304", expected: map[any]any{1: 2, 3: 4}},
				{encoded: "a56161614161626142616361436164614461656145", expected: map[any]any{"a": "A", "b": "B", "c": "C", "d": "D", "e": "E"}},
			}
			for _, test := range tests {
				t.Run(fmt.Sprintf("%T/%s=>%v", test.expected, test.encoded, test.expected), func(t *testing.T) {
					dencoderTest(t, test.encoded, test.expected)
				})
			}
		})
		t.Run("Extra", func(t *testing.T) {
			tests := []struct {
				encoded  string
				expected any
			}{
				{encoded: "40", expected: []byte("")},
				{encoded: "4161", expected: []byte("a")},
				{encoded: "4449455446", expected: []byte("IETF")},
				{encoded: "42225c", expected: []byte("\"\\")},
			}
			for _, test := range tests {
				t.Run(fmt.Sprintf("%T/%s=>%v", test.expected, test.encoded, test.expected), func(t *testing.T) {
					dencoderTest(t, test.encoded, test.expected)
				})
			}
		})
	})
}
