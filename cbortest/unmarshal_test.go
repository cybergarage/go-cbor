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
	"fmt"
	"testing"

	"github.com/cybergarage/go-cbor/cbor"
)

func TestUnmarshalTo(t *testing.T) {
	unmarshalToTest := func(t *testing.T, from any, to any) {
		t.Helper()
		encBytes, err := cbor.Marshal(from)
		if err != nil {
			t.Error(err)
			return
		}
		err = cbor.UnmarshalTo(encBytes, to)
		if err != nil {
			t.Error(err)
			return
		}

		err = deepEqual(from, to)
		if err != nil {
			t.Error(err)
			return
		}
	}

	t.Run("basicTypes", func(t *testing.T) {
		var i8 int8
		tests := []struct {
			from any
			to   any
		}{
			{from: 1, to: &i8},
		}

		for _, test := range tests {
			t.Run(fmt.Sprintf("%T=>%T", test.from, test.to), func(t *testing.T) {
				unmarshalToTest(t, test.from, test.to)
			})
		}
	})

	t.Run("map/array/struct", func(t *testing.T) {
		tests := []struct {
			from any
			to   any
		}{
			{
				from: &struct {
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
			{
				from: []string{"one", "two"},
				to:   &[]string{},
			},
			{
				from: []string{"one", "two"},
				to:   make([]string, 2),
			},
			{
				from: map[string]int{"one": 1, "two": 2},
				to:   map[string]int{},
			},
			{
				from: map[any]any{"one": 1, "two": 2},
				to:   map[string]int{},
			},
			{
				from: map[string]int{"one": 1, "two": 2},
				to:   map[any]any{},
			},
			{
				from: map[int]string{1: "one", 2: "two"},
				to:   map[int]string{},
			},
			{
				from: map[any]any{1: "one", 2: "two"},
				to:   map[int]string{},
			},
			{
				from: map[int]string{1: "one", 2: "two"},
				to:   map[any]any{},
			},
		}

		for _, test := range tests {
			t.Run(fmt.Sprintf("%T=>%T", test.from, test.to), func(t *testing.T) {
				unmarshalToTest(t, test.from, test.to)
			})
		}
	})
}
