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
		// []string{"one", "two"},
		[]string{"one", "two"},
		map[any]any{1: "one", 2: "two"},
	}

	toObjs := []any{
		&struct {
			Key   string
			Value string
		}{},
		// &[]string{},
		make([]string, 2),
		map[int]string{},
	}

	for n, fromObj := range fromObjs {
		t.Run(fmt.Sprintf("%v", fromObj), func(t *testing.T) {
			encBytes, err := Marshal(fromObj)
			if err != nil {
				t.Error(err)
				return
			}
			err = UnmarshalTo(encBytes, toObjs[n])
			if err != nil {
				t.Error(err)
				return
			}
			if !reflect.DeepEqual(fromObj, toObjs[n]) {
				if fmt.Sprintf("%v", fromObj) != fmt.Sprintf("%v", toObjs[n]) {
					t.Errorf("%v != %v", fromObj, toObjs[n])
				}
			}
		})
	}
}
