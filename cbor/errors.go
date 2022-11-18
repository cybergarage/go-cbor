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
	errorUnkonwnNativeType     = "%T (%v) is %w"
	errorUnkonwnMajorType      = "major type (%d) is %w"
	errorUnkonwnAdditionalInfo = "major type (%d:%d) is %w"
	errorUnmarshalDataTypes    = " %w : cound not convert from %v(%T) to %T"
	errorUnmarshalShortArray   = " %w : short array size (%T[%d] < %T[%d])"
	errorCastDataTypes         = " %w : cound not convert from %v(%T) to %T"
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

func newErrorUnmarshalDataTypes(fromItem any, toItem any) error {
	return fmt.Errorf(errorUnmarshalDataTypes, ErrUnmarshal, fromItem, fromItem, toItem)
}

func newErrorUnmarshalArraySize(fromArray []any, toObj any, toArrayVal reflect.Value) error {
	return fmt.Errorf(errorUnmarshalShortArray, ErrUnmarshal, fromArray, len(fromArray), toObj, toArrayVal.Cap())
}

func newErrorUnmarshalCastTypes(fromItem any, toItem any) error {
	return fmt.Errorf(errorCastDataTypes, ErrUnmarshal, fromItem, fromItem, toItem)
}
