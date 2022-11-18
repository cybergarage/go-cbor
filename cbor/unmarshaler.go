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
	"bytes"
	"math"
	"reflect"
)

// Unmarshal decodes the specified CBOR-encoded bytes and returns the data representation of Go. Unmarshal is a sugar function of Decoder::Decode().
func Unmarshal(cborBytes []byte) (any, error) {
	decoder := NewDecoder(bytes.NewReader(cborBytes))
	return decoder.Decode()
}

// UnmarshalTo decodes the specified CBOR-encoded bytes and stores the decoded item to the specified data type if appropriate. UnmarshalTo is a sugar function of Decoder::Unmarshal().
func UnmarshalTo(cborBytes []byte, s any) error {
	decoder := NewDecoder(bytes.NewReader(cborBytes))
	return decoder.Unmarshal(s)
}

// nolint: exhaustive
// Unmarshal decodes a next encoded item from the specified reader and stores the decoded item to the specified data type if appropriate.
func (dec *Decoder) Unmarshal(toObj any) error {
	fromObj, err := dec.Decode()
	if err != nil {
		return err
	}

	switch from := fromObj.(type) {
	case map[any]any:
		switch reflect.ValueOf(toObj).Type().Kind() {
		case reflect.Struct:
			return dec.unmarshalMapToStrct(from, reflect.ValueOf(toObj))
		case reflect.Map:
			return dec.unmarshalMapToMap(from, toObj)
		case reflect.Pointer:
			elem := reflect.ValueOf(toObj).Elem()
			if elem.Type().Kind() != reflect.Struct {
				return newErrorUnmarshalDataTypes(fromObj, toObj)
			}
			return dec.unmarshalMapToStrct(from, elem)
		default:
			return newErrorUnmarshalDataTypes(fromObj, toObj)
		}
	case []any:
		switch reflect.ValueOf(toObj).Type().Kind() {
		case reflect.Array, reflect.Slice, reflect.Pointer:
			return dec.unmarshalArrayTo(from, toObj)
		}
		return newErrorUnmarshalDataTypes(fromObj, toObj)
	case int:
		switch to := toObj.(type) {
		case *int:
			*to = from
		case *int8:
			if math.MaxInt8 < from {
				return newErrorUnmarshalCastOverflow(from, toObj)
			}
			*to = int8(from)
		case *int16:
			if math.MaxInt16 < from {
				return newErrorUnmarshalCastOverflow(from, toObj)
			}
			*to = int16(from)
		case *int32:
			if math.MaxInt32 < from {
				return newErrorUnmarshalCastOverflow(from, toObj)
			}
			*to = int32(from)
		case *int64:
			*to = int64(from)
		}
	case int64:
		switch to := toObj.(type) {
		case *int:
			if math.MaxInt < from {
				return newErrorUnmarshalCastOverflow(from, toObj)
			}
			*to = int(from)
		case *int8:
			if math.MaxInt8 < from {
				return newErrorUnmarshalCastOverflow(from, toObj)
			}
			*to = int8(from)
		case *int16:
			if math.MaxInt16 < from {
				return newErrorUnmarshalCastOverflow(from, toObj)
			}
			*to = int16(from)
		case *int32:
			if math.MaxInt32 < from {
				return newErrorUnmarshalCastOverflow(from, toObj)
			}
			*to = int32(from)
		case *int64:
			*to = from
		}
	}

	return newErrorUnmarshalDataTypes(fromObj, toObj)
}

// nolint: exhaustive
func (dec *Decoder) unmarshalArrayTo(fromArray []any, toObj any) error {
	// NOTE: The Laws of Reflection - The Go Programming Language
	// https://go.dev/blog/laws-of-reflection

	fromArrayVal := reflect.ValueOf(fromArray)
	fromArrayType := fromArrayVal.Type()
	toArrayVal := reflect.ValueOf(toObj)
	fromArrayLen := len(fromArray)
	toArrayType := toArrayVal.Type()
	switch toArrayType.Kind() {
	case reflect.Array:
		if toArrayVal.Len() < fromArrayLen {
			return newErrorUnmarshalArraySize(fromArray, toObj, toArrayVal)
		}
	case reflect.Slice:
		if toArrayVal.Len() < fromArrayLen {
			if !toArrayVal.CanSet() {
				return newErrorUnmarshalArraySize(fromArray, toObj, toArrayVal)
			}
			toArrayVal.Set(reflect.MakeSlice(fromArrayType, fromArrayLen, fromArrayLen))
		}
	case reflect.Pointer:
		elem := toArrayVal.Elem()
		switch elem.Type().Kind() {
		case reflect.Array:
			if elem.Len() < fromArrayLen {
				return newErrorUnmarshalArraySize(fromArray, toObj, toArrayVal)
			}
		case reflect.Slice:
			if elem.Len() < fromArrayLen {
				if !elem.CanSet() {
					return newErrorUnmarshalArraySize(fromArray, toObj, toArrayVal)
				}
				toArrayType = elem.Type()
				appendLen := fromArrayLen - elem.Len()
				elem.Set(reflect.AppendSlice(elem, reflect.MakeSlice(toArrayType, appendLen, appendLen)))
				toArrayVal = elem
			}
		default:
			return newErrorUnmarshalDataTypes(fromArray, toObj)
		}
	default:
		return newErrorUnmarshalDataTypes(fromArray, toObj)
	}

	toObjType := toArrayType.Elem().Kind()
	for n, fromObj := range fromArray {
		fromObjType := reflect.TypeOf(fromObj).Kind()
		if fromObjType != toObjType {
			return newErrorUnmarshalDataTypes(fromObj, toObj)
		}
		toArrayIndex := toArrayVal.Index(n)
		toArrayIndex.Set(reflect.ValueOf(fromObj))
	}

	return nil
}

func (dec *Decoder) unmarshalMapToStrct(fromMap map[any]any, toStructVal reflect.Value) error {
	if toStructVal.Type().Kind() != reflect.Struct {
		return newErrorUnmarshalDataTypes(fromMap, toStructVal)
	}
	for fromMapKey, fromMapElem := range fromMap {
		key, ok := fromMapKey.(string)
		if !ok {
			return newErrorUnmarshalDataTypes(fromMap, toStructVal)
		}
		toStructField := toStructVal.FieldByName(key)
		if !ok {
			return newErrorUnmarshalDataTypes(fromMap, toStructVal)
		}
		fromMapElemVal := reflect.ValueOf(fromMapElem)
		if fromMapElemVal.Type().Kind() != toStructField.Type().Kind() {
			return newErrorUnmarshalDataTypes(fromMap, toStructVal)
		}
		toStructField.Set(fromMapElemVal)
	}
	return nil
}

func (dec *Decoder) unmarshalMapToMap(fromMap map[any]any, toMap any) error {
	toMapVal := reflect.ValueOf(toMap)
	toMapType := toMapVal.Type()
	if toMapType.Kind() != reflect.Map {
		return newErrorUnmarshalDataTypes(fromMap, toMap)
	}
	toMapKeyType := toMapType.Key()
	toMapElemType := toMapType.Elem()
	for fromMapKey, fromMapValue := range fromMap {
		fromMapKeyVal := reflect.ValueOf(fromMapKey)
		if !fromMapKeyVal.CanConvert(toMapKeyType) {
			return newErrorUnmarshalDataTypes(fromMapKey, toMapVal)
		}
		fromMapElemVal := reflect.ValueOf(fromMapValue)
		if !fromMapElemVal.CanConvert(toMapElemType) {
			return newErrorUnmarshalDataTypes(fromMapKey, toMapVal)
		}
		toMapVal.SetMapIndex(fromMapKeyVal.Convert(toMapKeyType), fromMapElemVal.Convert(toMapElemType))
	}
	return nil
}
