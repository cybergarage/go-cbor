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

package cbor_test

import (
	"encoding/hex"
	"fmt"

	"github.com/cybergarage/go-cbor/cbor"
)

// nolint: nosnakecase
func ExampleMarshal() {
	goObjs := []any{
		uint(1000),
		int(-1000),
		float32(100000.0),
		float64(-4.1),
		false,
		true,
		nil,
		[]byte("IETF"),
		"IETF",
		[]int{1, 2, 3},
		map[any]any{"a": "A", "b": "B", "c": "C"},
	}
	for _, goObj := range goObjs {
		cborBytes, _ := cbor.Marshal(goObj)
		fmt.Printf("%v => %s", goObj, hex.EncodeToString(cborBytes))
	}
}
