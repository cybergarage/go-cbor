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

package cbor

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func TestUnmarshalTo(t *testing.T) {
	fromObjs := []any{
		&struct {
			Key   string
			Value string
		}{
			Key: "hello", Value: "world",
		},
		[]string{"one", "two"},
		[]string{"one", "two"},
		map[string]int{"one": 1, "two": 2},
		map[any]any{"one": 1, "two": 2},
		map[string]int{"one": 1, "two": 2},
		map[int]string{1: "one", 2: "two"},
		map[any]any{1: "one", 2: "two"},
		map[int]string{1: "one", 2: "two"},
	}

	toObjs := []any{
		&struct {
			Key   string
			Value string
		}{},
		&[]string{},
		make([]string, 2),
		map[string]int{},
		map[string]int{},
		map[any]any{},
		map[int]string{},
		map[int]string{},
		map[any]any{},
	}

	for n, fromObj := range fromObjs {
		toObj := toObjs[n]
		t.Run(fmt.Sprintf("%T=>%T", fromObj, toObj), func(t *testing.T) {
			encBytes, err := Marshal(fromObj)
			if err != nil {
				t.Error(err)
				return
			}
			err = UnmarshalTo(encBytes, toObj)
			if err != nil {
				t.Error(err)
				return
			}
			if !reflect.DeepEqual(fromObj, toObj) {
				re := regexp.MustCompile(fmt.Sprintf("[&]?%v", fromObj))
				if !re.MatchString(fmt.Sprintf("%v", toObj)) {
					t.Errorf("%v != %v", fromObj, toObj)
				}
			}
		})
	}
}
