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
	"time"

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

	t.Run("basic_types", func(t *testing.T) {
		t20120321, err := time.Parse(time.RFC3339, "2013-03-21T20:04:00Z")
		if err != nil {
			t.Error(err)
			return
		}

		var vi int
		var vi8 int8
		var vi16 int16
		var vi32 int32
		var vi64 int64
		var uvi uint
		var uvi8 uint8
		var uvi16 uint16
		var uvi32 uint32
		var uvi64 uint64
		var vf32 float32
		var vf64 float64
		var vs string
		var vt time.Time

		tests := []struct {
			from any
			to   any
		}{
			{from: t20120321, to: &vt},
		}

		froms := []any{int(1), int8(1), int16(1), int32(1), int64(1), uint(1), uint8(1), uint16(1), uint32(1), uint64(1), float32(1), float64(1)}
		toTypes := []any{&vi, &vi8, &vi16, &vi32, &vi64, &uvi, &uvi8, &uvi16, &uvi32, &uvi64, &vf32, &vf64, &vs}
		for _, from := range froms {
			for _, toType := range toTypes {
				test := struct {
					from any
					to   any
				}{from: from, to: toType}
				tests = append(tests, test)
			}
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
					Name    string
					Address struct {
						Code uint
					}
				}{
					Name: "hello", Address: struct{ Code uint }{Code: 1},
				},
				to: &struct {
					Name    string
					Address struct {
						Code uint
					}
				}{},
			},
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
