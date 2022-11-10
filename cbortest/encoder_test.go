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
			t.Run("uint8", func(t *testing.T) {
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
				}
				for _, test := range tests {
					t.Run(fmt.Sprintf("%v:%s", test.value, test.expected), func(t *testing.T) {
						var writer bytes.Buffer
						encoder := cbor.NewEncoder(&writer)
						err := encoder.Encode(test.value)
						if err != nil {
							t.Errorf("%v (%s)", test.value, err.Error())
							return
						}
						encoded := hex.EncodeToString(writer.Bytes())
						if encoded != test.expected {
							t.Errorf("%s != %s", encoded, test.expected)
						}
					})
				}
			})
		})
	})
}
