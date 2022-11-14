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
	"bytes"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/cybergarage/go-cbor/cbor"
)

func ExampleDecoder_Decode() {
	cborObjs := []string{
		"0a",
		"1903e8",
		"3903e7",
		"fb3ff199999999999a",
		"f90001",
		"f4",
		"f5",
		"f6",
		"c074323031332d30332d32315432303a30343a30305a",
		"4449455446",
		"6449455446",
		"83010203",
		"a201020304",
	}
	for _, cborObj := range cborObjs {
		cborBytes, _ := hex.DecodeString(cborObj)
		decoder := cbor.NewDecoder(bytes.NewReader(cborBytes))
		goObj, _ := decoder.Decode()
		fmt.Printf("%s => %v\n", cborObj, goObj)
	}
}

func ExampleUnmarshal() {
	cborObjs := []string{
		"0a",
		"1903e8",
		"3903e7",
		"fb3ff199999999999a",
		"f90001",
		"f4",
		"f5",
		"f6",
		"c074323031332d30332d32315432303a30343a30305a",
		"4449455446",
		"6449455446",
		"83010203",
		"a201020304",
	}
	for _, cborObj := range cborObjs {
		cborBytes, _ := hex.DecodeString(cborObj)
		goObj, _ := cbor.Unmarshal(cborBytes)
		fmt.Printf("%s => %v\n", cborObj, goObj)
	}
}

func ExampleEncoder_Encode() {
	goTimeObj, _ := time.Parse(time.RFC3339, "2013-03-21T20:04:00Z")
	goObjs := []any{
		uint(1000),
		int(-1000),
		float32(100000.0),
		float64(-4.1),
		false,
		true,
		nil,
		goTimeObj,
		[]byte("IETF"),
		"IETF",
		[]int{1, 2, 3},
		map[any]any{"a": "A", "b": "B", "c": "C"},
	}
	for _, goObj := range goObjs {
		var w bytes.Buffer
		encoder := cbor.NewEncoder(&w)
		encoder.Encode(goObj)
		cborBytes := w.Bytes()
		fmt.Printf("%v => %s\n", goObj, hex.EncodeToString(cborBytes))
	}
}

func ExampleMarshal() {
	goTimeObj, _ := time.Parse(time.RFC3339, "2013-03-21T20:04:00Z")
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
		goTimeObj,
		[]int{1, 2, 3},
		map[any]any{"a": "A", "b": "B", "c": "C"},
	}
	for _, goObj := range goObjs {
		cborBytes, _ := cbor.Marshal(goObj)
		fmt.Printf("%v => %s\n", goObj, hex.EncodeToString(cborBytes))
	}
}
