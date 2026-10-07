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
	"reflect"
	"testing"

	"github.com/cybergarage/go-cbor/cbor"
)

func TestUnmarshalNestedArrays(t *testing.T) {
	tests := []struct {
		from any
		to   any
	}{
		{[][]byte{[]byte("x"), []byte("yz")}, &[][]byte{}},
		{[][]int{{1, 2}, {3}}, &[][]int{}},
		{[]any{[]any{"a"}, []any{"b", "c"}}, &[][]string{}},
		{struct{ List [][]int }{List: [][]int{{1}, {2, 3}}}, &struct{ List [][]int }{}},
	}
	for _, test := range tests {
		b, err := cbor.Marshal(test.from)
		if err != nil {
			t.Fatal(err)
		}
		if err := cbor.UnmarshalTo(b, test.to); err != nil {
			t.Errorf("%v: %v", test.from, err)
			continue
		}
		got := reflect.ValueOf(test.to).Elem().Interface()
		if err := deepEqual(test.from, got); err != nil {
			t.Error(err)
		}
	}
}
