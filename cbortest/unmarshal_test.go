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

	t.Run("basicTypes", func(t *testing.T) {
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
		var vt time.Time

		tests := []struct {
			from any
			to   any
		}{
			{from: t20120321, to: &vt},
			{from: int(1), to: &vi},
			{from: int(1), to: &vi8},
			{from: int(1), to: &vi16},
			{from: int(1), to: &vi32},
			{from: int(1), to: &vi64},
			{from: int(1), to: &uvi},
			{from: int(1), to: &uvi8},
			{from: int(1), to: &uvi16},
			{from: int(1), to: &uvi32},
			{from: int(1), to: &uvi64},
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
