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
		"a161616141",
	}
	for _, cborObj := range cborObjs {
		cborBytes, _ := hex.DecodeString(cborObj)
		decoder := cbor.NewDecoder(bytes.NewReader(cborBytes))
		goObj, _ := decoder.Decode()
		fmt.Printf("%v\n", goObj)
	}

	// Output:
	// 10
	// 1000
	// -1000
	// 1.1
	// <nil>
	// false
	// true
	// <nil>
	// 2013-03-21 20:04:00 +0000 UTC
	// [73 69 84 70]
	// IETF
	// [1 2 3]
	// map[a:A]
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
		"a161616141",
	}
	for _, cborObj := range cborObjs {
		cborBytes, _ := hex.DecodeString(cborObj)
		goObj, _ := cbor.Unmarshal(cborBytes)
		fmt.Printf("%v\n", goObj)
	}

	// Output:
	// 10
	// 1000
	// -1000
	// 1.1
	// <nil>
	// false
	// true
	// <nil>
	// 2013-03-21 20:04:00 +0000 UTC
	// [73 69 84 70]
	// IETF
	// [1 2 3]
	// map[a:A]
}

func ExampleDecoder_Unmarshal() {
	fromObjs := []any{
		[]string{"one", "two"},
		map[string]int{"one": 1, "two": 2},
		struct {
			Key   string
			Value string
		}{
			Key: "hello", Value: "world",
		},
	}

	toObjs := []any{
		&[]string{},
		map[string]int{},
		&struct {
			Key   string
			Value string
		}{},
	}

	for n, fromObj := range fromObjs {
		var w bytes.Buffer
		encoder := cbor.NewEncoder(&w)
		encoder.Encode(fromObj)
		cborBytes := w.Bytes()

		toObj := toObjs[n]
		decoder := cbor.NewDecoder(bytes.NewReader(cborBytes))
		decoder.Unmarshal(toObj)
		fmt.Printf("%v\n", toObj)
	}

	// Output:
	// &[one two]
	// map[one:1 two:2]
	// &{hello world}
}

func ExampleUnmarshalTo() {
	examples := []struct {
		from any
		to   any
	}{
		{
			from: []string{"one", "two"},
			to:   &[]string{},
		},
		{
			from: map[string]int{"one": 1, "two": 2},
			to:   map[string]int{},
		},
		{
			from: struct {
				Key   string
				Value string
			}{
				Key: "hello", Value: "world",
			},
			to: &struct {
				Key   string
				Value string
			}{},
		},
	}

	for _, e := range examples {
		encBytes, _ := cbor.Marshal(e.from)
		cbor.UnmarshalTo(encBytes, e.to)
		fmt.Printf("%v\n", e.to)
	}

	// Output:
	// &[one two]
	// map[one:1 two:2]
	// &{hello world}
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
		map[any]any{"a": "A"},
	}
	for _, goObj := range goObjs {
		var w bytes.Buffer
		encoder := cbor.NewEncoder(&w)
		encoder.Encode(goObj)
		cborBytes := w.Bytes()
		fmt.Printf("%s\n", hex.EncodeToString(cborBytes))
	}

	// Output:
	// 1b00000000000003e8
	// 3b00000000000003e7
	// fa47c35000
	// fbc010666666666666
	// f4
	// f5
	// f6
	// c074323031332d30332d32315432303a30343a30305a
	// 4449455446
	// 6449455446
	// 831b00000000000000011b00000000000000021b0000000000000003
	// a161616141
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
		map[any]any{"a": "A"},
	}
	for _, goObj := range goObjs {
		cborBytes, _ := cbor.Marshal(goObj)
		fmt.Printf("%s\n", hex.EncodeToString(cborBytes))
	}

	// Output:
	// 1b00000000000003e8
	// 3b00000000000003e7
	// fa47c35000
	// fbc010666666666666
	// f4
	// f5
	// f6
	// 4449455446
	// 6449455446
	// c074323031332d30332d32315432303a30343a30305a
	// 831b00000000000000011b00000000000000021b0000000000000003
	// a161616141
}
