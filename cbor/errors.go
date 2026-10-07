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
	"math"
	"reflect"
)

var ErrNotSupported = errors.New("not supported")
var ErrUnmarshal = errors.New("unmarshal error")
var ErrDecode = errors.New("decode error")
var ErrEncode = errors.New("encode error")

// ErrJSON is returned when data cannot be converted between CBOR and JSON (RFC 8949 Section 6).
var ErrJSON = errors.New("JSON conversion error")

// ErrInvalidUTF8 is wrapped with ErrEncode or ErrDecode when a text string is not valid UTF-8 (RFC 8949 Section 5.3.1).
var ErrInvalidUTF8 = errors.New("not valid UTF-8")

const (
	errorUnkonwnNativeType      = "%T (%v) is %w"
	errorUnkonwnMajorType       = "major type (%d) is %w"
	errorUnkonwnAdditionalInfo  = "major type (%d:%d) is %w"
	errorUnmarshalDataTypes     = "%w : cound not convert from %v (%T) to %T"
	errorUnmarshalShortArray    = "%w : short array size (%T(%d) < %T(%d))"
	errorUnmarshalCastTypes     = "%w : cound not cast from %v (%T) to %T"
	errorSortedMapEncode        = "%w : map key (%v:%T) could not be sorted"
	errorUnmarshalReflectValues = "%w : cound not convert from %v to %T"
	errorDecodeLengthTooLarge   = "%w : length (%d) exceeds the maximum supported length (%d)"
	errorDecodeUnhashableKey    = "%w : map key (%T) cannot be used as a Go map key"
	errorDecodeUnexpectedBreak  = "%w : unexpected break stop code"
	errorDecodeInvalidChunk     = "%w : invalid chunk (0x%02X) in indefinite-length string of major type (%d)"
	errorDecodeInvalidSimple    = "%w : simple value (%d) must not be encoded in the following byte"
	errorDecodeInvalidTag       = "%w : invalid content (%v:%T) for tag (%d)"
	errorEncodeReservedSimple   = "%w : simple value (%d) is reserved"
	errorEncodeDuplicateMapKey  = "%w : duplicate map key (0x%X)"
	errorEncodeIndefinite       = "%w : %s"
	errorDecodeNotDeterministic = "%w : %s is not allowed in the core deterministic encoding"
	errorInvalidUTF8            = "%w : text string is %w"
	errorDecodeTooDeep          = "%w : nesting depth exceeds the maximum (%d)"
	errorDecodeDuplicateMapKey  = "%w : duplicate map key (%v)"
	errorDecodeExtraneousData   = "%w : %d bytes of extraneous data after the data item"
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

func newErrorUnmarshalArraySize(fromArrayVal reflect.Value, toArrayVal reflect.Value) error {
	return fmt.Errorf(errorUnmarshalShortArray, ErrUnmarshal, fromArrayVal, fromArrayVal.Len(), toArrayVal, toArrayVal.Cap())
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

func newErrorDecodeLengthTooLarge(n uint64) error {
	return fmt.Errorf(errorDecodeLengthTooLarge, ErrDecode, n, math.MaxInt)
}

func newErrorDecodeUnhashableKey(key any) error {
	return fmt.Errorf(errorDecodeUnhashableKey, ErrDecode, key)
}

func newErrorDecodeUnexpectedBreak() error {
	return fmt.Errorf(errorDecodeUnexpectedBreak, ErrDecode)
}

func newErrorDecodeInvalidChunk(m majorType, header byte) error {
	return fmt.Errorf(errorDecodeInvalidChunk, ErrDecode, header, (m >> 5))
}

func newErrorDecodeInvalidSimpleValue(v uint8) error {
	return fmt.Errorf(errorDecodeInvalidSimple, ErrDecode, v)
}

func newErrorDecodeInvalidTagContent(number uint64, content any) error {
	return fmt.Errorf(errorDecodeInvalidTag, ErrDecode, content, content, number)
}

func newErrorEncodeReservedSimpleValue(v SimpleValue) error {
	return fmt.Errorf(errorEncodeReservedSimple, ErrEncode, uint8(v))
}

func newErrorEncodeDuplicateMapKey(key []byte) error {
	return fmt.Errorf(errorEncodeDuplicateMapKey, ErrEncode, key)
}

func newErrorDecodeNotDeterministic(what string) error {
	return fmt.Errorf(errorDecodeNotDeterministic, ErrDecode, what)
}

func newErrorDecodeInvalidUTF8() error {
	return fmt.Errorf(errorInvalidUTF8, ErrDecode, ErrInvalidUTF8)
}

func newErrorEncodeInvalidUTF8() error {
	return fmt.Errorf(errorInvalidUTF8, ErrEncode, ErrInvalidUTF8)
}

func newErrorDecodeTooDeep(maxLevels int) error {
	return fmt.Errorf(errorDecodeTooDeep, ErrDecode, maxLevels)
}

func newErrorDecodeDuplicateMapKey(key any) error {
	return fmt.Errorf(errorDecodeDuplicateMapKey, ErrDecode, key)
}

func newErrorDecodeExtraneousData(n int) error {
	return fmt.Errorf(errorDecodeExtraneousData, ErrDecode, n)
}

func newErrorEncodeIndefinite(reason string) error {
	return fmt.Errorf(errorEncodeIndefinite, ErrEncode, reason)
}

func newErrorJSON(err error) error {
	return fmt.Errorf("%w : %w", ErrJSON, err)
}

func newErrorJSONKeyCollision(key string) error {
	return fmt.Errorf("%w : duplicate object member name (%q)", ErrJSON, key)
}

func newErrorJSONUnsupportedKey(key string) error {
	return fmt.Errorf("%w : map key (%s) cannot be converted to a JSON object member name", ErrJSON, key)
}
