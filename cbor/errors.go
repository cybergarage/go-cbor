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
	"errors"
	"fmt"
	"reflect"
)

var ErrNotSupported = errors.New("not supported")
var ErrUnmarshal = errors.New("unmarshal error")
var ErrDecode = errors.New("decode error")
var ErrEncode = errors.New("encode error")

const (
	errorUnkonwnNativeType      = "%T (%v) is %w"
	errorUnkonwnMajorType       = "major type (%d) is %w"
	errorUnkonwnAdditionalInfo  = "major type (%d:%d) is %w"
	errorUnmarshalDataTypes     = "%w : cound not convert from %v (%T) to %T"
	errorUnmarshalShortArray    = "%w : short array size ([%d]%T < [%d]%T)"
	errorUnmarshalCastTypes     = "%w : cound not cast from %v (%T) to %T"
	errorSortedMapEncode        = "%w : map key (%v:%T) could not be sorted"
	errorUnmarshalReflectValues = "%w : cound not convert from %v to %T"
)

func newErrorNotSupportedMajorType(m majorType) error {
	return fmt.Errorf(errorUnkonwnMajorType, (m >> 5), ErrNotSupported)
}

func newErrorNotSupportedAddInfo(m majorType, a majorInfo) error {
	return fmt.Errorf(errorUnkonwnAdditionalInfo, (m >> 5), a, ErrNotSupported)
}

func newErrorNotSupportedNativeType(item any) error {
	return fmt.Errorf(errorUnkonwnNativeType, item, item, ErrNotSupported)
}

func newErrorUnmarshalDataTypes(from any, to any) error {
	return fmt.Errorf(errorUnmarshalDataTypes, ErrUnmarshal, from, from, to)
}

func newErrorUnmarshalArraySize(fromArray []any, toObj any, toArrayVal reflect.Value) error {
	return fmt.Errorf(errorUnmarshalShortArray, ErrUnmarshal, fromArray, len(fromArray), toObj, toArrayVal.Cap())
}

func newErrorUnmarshalCastTypes(from any, to any) error {
	return fmt.Errorf(errorUnmarshalCastTypes, ErrUnmarshal, from, from, to)
}

func newErrorSortedMapEncode(key any) error {
	return fmt.Errorf(errorSortedMapEncode, ErrEncode, key, key)
}

func newErrorUnmarshalReflectValues(from reflect.Value, to reflect.Value) error {
	return fmt.Errorf(errorUnmarshalReflectValues, ErrUnmarshal, from.Kind().String(), to.Kind().String())
}
