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
	"reflect"
	"strings"
)

func deepEqual(fromObj any, toObj any) error {
	tostring := func(v any) string {
		if reflect.TypeOf(v).Kind() == reflect.Pointer {
			return fmt.Sprintf("%+v", reflect.ValueOf(v).Elem())
		}
		return fmt.Sprintf("%+v", v)
	}

	replaceUnmatchedStrs := func(v string) string {
		// Map string comparisons
		// fromObj:    "{Key:0 Val:-3.4028234663852886e+38}"
		// toObj  : "map[Key:0 Val:-3.4028234663852886e+38]"
		// Struct string comparisons
		// fromObj: {Elem1:{120 121 122}}
		// toObj  : {Elem1:[120 121 122]}
		v = strings.ReplaceAll(v, "[", "{")
		v = strings.ReplaceAll(v, "]", "}")
		v = strings.ReplaceAll(v, "map", "")

		// Array string comparisons
		// fromObj: "[one two]"
		// toObj  : "&[one two]"
		// fromObj: "{Key:32767 Val:&} != "{Key:32767 Val:}""
		// toObj  :{Key:32767 Val:} != {Key:32767 Val:&}
		v = strings.ReplaceAll(v, "&", "")

		return v
	}

	// reflect comparisons

	if reflect.DeepEqual(fromObj, toObj) {
		return nil
	}

	// String comparisons

	fromObjStr := tostring(fromObj)
	toObjStr := tostring(toObj)
	if toObjStr == fromObjStr {
		return nil
	}

	fromObjStr = replaceUnmatchedStrs(fromObjStr)
	toObjStr = replaceUnmatchedStrs(toObjStr)
	if toObjStr == fromObjStr {
		return nil
	}

	return fmt.Errorf("%v != %v", toObjStr, fromObjStr)
}
