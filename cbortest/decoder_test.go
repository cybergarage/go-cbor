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
	"reflect"
	"testing"

	"github.com/cybergarage/go-cbor/cbor"
)

func TestDecoder(t *testing.T) {
	t.Run("RFC-8949", func(t *testing.T) {
		t.Run("AppendixA", func(t *testing.T) {
			t.Run("uint8", func(t *testing.T) {
				tests := []struct {
					encoded  string
					expected any
				}{
					{encoded: "00", expected: uint8(0)},
					{encoded: "01", expected: uint8(1)},
					{encoded: "0a", expected: uint8(10)},
					{encoded: "17", expected: uint8(23)},
					{encoded: "1818", expected: uint8(24)},
					{encoded: "1819", expected: uint8(25)},
					{encoded: "1864", expected: uint8(100)},
				}
				for _, test := range tests {
					testBytes, err := hex.DecodeString(test.encoded)
					if err != nil {
						t.Errorf("%s => (%s)", test.encoded, err.Error())
						continue
					}
					decoder := cbor.NewDecoder(bytes.NewReader(testBytes))
					v, err := decoder.Decode()
					if err != nil {
						t.Errorf("%s => %v", test.encoded, test.expected)
					}
					if !reflect.DeepEqual(v, test.expected) {
						t.Errorf("%v != %v", v, test.expected)
					}
				}
			})
		})
	})
}
