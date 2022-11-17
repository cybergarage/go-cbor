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

func DeepEqual(fromObj any, toObj any) error {
	if reflect.DeepEqual(fromObj, toObj) {
		return nil
	}

	// String comparisons
	fromObjStr := fmt.Sprintf("%+v", fromObj)
	toObjStr := fmt.Sprintf("%+v", toObj)
	if toObjStr == fromObjStr {
		return nil
	}

	// Array string comparisons
	// fromObj: "[one two]"
	// toObj  : "&[one two]"
	// NOTE: (}`): error parsing regexp: missing closing )
	// re := regexp.MustCompile(fmt.Sprintf("[&]?%s", fromObjStr))
	// if re.MatchString(toObjStr) {
	// 	return nil
	// }
	toObjStr = strings.ReplaceAll(toObjStr, "&", "")
	if toObjStr == fromObjStr {
		return nil
	}

	// Map string comparisons
	// fromObj:    "{Key:0 Val:-3.4028234663852886e+38}"
	// toObj  : "map[Key:0 Val:-3.4028234663852886e+38]"
	toObjStr = strings.ReplaceAll(toObjStr, "[", "{")
	toObjStr = strings.ReplaceAll(toObjStr, "]", "}")
	toObjStr = strings.ReplaceAll(toObjStr, "map", "")
	if toObjStr == fromObjStr {
		return nil
	}
	// "{Key:32767 Val:&} != "{Key:32767 Val:}""
	// {Key:32767 Val:} != {Key:32767 Val:&}
	fromObjStr = strings.ReplaceAll(fromObjStr, "&", "")
	toObjStr = strings.ReplaceAll(toObjStr, "&", "")
	if toObjStr == fromObjStr {
		return nil
	}

	return fmt.Errorf("%v != %v", toObjStr, fromObjStr)
}
