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
	"reflect"
	"time"

	"github.com/cybergarage/go-safecast/safecast"
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
			return dec.unmarshalMapToStruct(from, reflect.ValueOf(toObj))
		case reflect.Map:
			return dec.unmarshalMapToMap(from, toObj)
		case reflect.Pointer:
			elem := reflect.ValueOf(toObj).Elem()
			if elem.Type().Kind() != reflect.Struct {
				return newErrorUnmarshalDataTypes(fromObj, toObj)
			}
			return dec.unmarshalMapToStruct(from, elem)
		default:
			return newErrorUnmarshalDataTypes(fromObj, toObj)
		}
	case []any:
		switch reflect.ValueOf(toObj).Type().Kind() {
		case reflect.Array, reflect.Slice, reflect.Pointer:
			return dec.unmarshalArrayTo(reflect.ValueOf(fromObj), reflect.ValueOf(toObj))
		}
		return newErrorUnmarshalDataTypes(fromObj, toObj)
	case time.Time:
		return dec.unmarshalToSpecialStruct(fromObj, toObj)
	}

	return dec.unmarshalToBasicType(fromObj, toObj)
}

// nolint: exhaustive
func (dec *Decoder) unmarshalArrayTo(fromArrayVal reflect.Value, toArrayVal reflect.Value) error {
	// NOTE: The Laws of Reflection - The Go Programming Language
	// https://go.dev/blog/laws-of-reflection

	fromArrayType := fromArrayVal.Type()
	fromArrayLen := fromArrayVal.Len()
	toArrayType := toArrayVal.Type()
	switch toArrayType.Kind() {
	case reflect.Array:
		if toArrayVal.Len() < fromArrayLen {
			return newErrorUnmarshalArraySize(fromArrayVal, toArrayVal)
		}
	case reflect.Slice:
		if toArrayVal.Len() < fromArrayLen {
			if !toArrayVal.CanSet() {
				return newErrorUnmarshalArraySize(fromArrayVal, toArrayVal)
			}
			toArrayVal.Set(reflect.MakeSlice(fromArrayType, fromArrayLen, fromArrayLen))
		}
	case reflect.Pointer:
		elem := toArrayVal.Elem()
		switch elem.Type().Kind() {
		case reflect.Array:
			if elem.Len() < fromArrayLen {
				return newErrorUnmarshalArraySize(fromArrayVal, toArrayVal)
			}
		case reflect.Slice:
			if elem.Len() < fromArrayLen {
				if !elem.CanSet() {
					return newErrorUnmarshalArraySize(fromArrayVal, toArrayVal)
				}
				toArrayType = elem.Type()
				appendLen := fromArrayLen - elem.Len()
				elem.Set(reflect.AppendSlice(elem, reflect.MakeSlice(toArrayType, appendLen, appendLen)))
				toArrayVal = elem
			}
		default:
			return newErrorUnmarshalDataTypes(fromArrayVal.Interface(), toArrayVal.Interface())
		}
	default:
		return newErrorUnmarshalDataTypes(fromArrayVal.Interface(), toArrayVal.Interface())
	}

	for n := 0; n < fromArrayLen; n++ {
		fromVal := fromArrayVal.Index(n)
		fromObj := fromVal.Interface()
		fromObjType := reflect.TypeOf(fromObj)
		fromObjKind := fromObjType.Kind()
		toObj := toArrayVal.Index(n).Interface()
		toObjType := reflect.TypeOf(toObj)
		toObjKind := toObjType.Kind()
		if fromObjKind == toObjKind {
			toArrayIndex := toArrayVal.Index(n)
			toArrayIndex.Set(reflect.ValueOf(fromObj))
			continue
		}
		if fromVal.CanConvert(toObjType) {
			toArrayIndex := toArrayVal.Index(n)
			toArrayIndex.Set(fromVal.Convert(toObjType))
			continue
		}
		return newErrorUnmarshalDataTypes(fromObj, toObj)
	}
	return nil
}

func (dec *Decoder) unmarshalMapElemToStructField(fromVal reflect.Value, toVal reflect.Value) error {
	from := fromVal.Interface()
	fromType := fromVal.Type()
	fromKind := fromType.Kind()
	toType := toVal.Type()
	toKind := toType.Kind()
	if fromKind == toKind {
		toVal.Set(fromVal)
		return nil
	}
	if fromVal.CanConvert(toType) {
		toVal.Set(fromVal.Convert(toType))
		return nil
	}
	switch toKind { // nolint: exhaustive
	case reflect.Int:
		var v int
		if err := safecast.ToInt(from, &v); err == nil {
			toVal.Set(reflect.ValueOf(v))
			return nil
		}
	case reflect.Int8:
		var v int8
		if err := safecast.ToInt8(from, &v); err == nil {
			toVal.Set(reflect.ValueOf(v))
			return nil
		}
	case reflect.Int16:
		var v int16
		if err := safecast.ToInt16(from, &v); err == nil {
			toVal.Set(reflect.ValueOf(v))
			return nil
		}
	case reflect.Int32:
		var v int32
		if err := safecast.ToInt32(from, &v); err == nil {
			toVal.Set(reflect.ValueOf(v))
			return nil
		}
	case reflect.Int64:
		var v int64
		if err := safecast.ToInt64(from, &v); err == nil {
			toVal.Set(reflect.ValueOf(v))
			return nil
		}
	case reflect.Uint:
		var v uint
		if err := safecast.ToUint(from, &v); err == nil {
			toVal.Set(reflect.ValueOf(v))
			return nil
		}
	case reflect.Uint8:
		var v uint8
		if err := safecast.ToUint8(from, &v); err == nil {
			toVal.Set(reflect.ValueOf(v))
			return nil
		}
	case reflect.Uint16:
		var v uint16
		if err := safecast.ToUint16(from, &v); err == nil {
			toVal.Set(reflect.ValueOf(v))
			return nil
		}
	case reflect.Uint32:
		var v uint32
		if err := safecast.ToUint32(from, &v); err == nil {
			toVal.Set(reflect.ValueOf(v))
			return nil
		}
	case reflect.Uint64:
		var v uint64
		if err := safecast.ToUint64(from, &v); err == nil {
			toVal.Set(reflect.ValueOf(v))
			return nil
		}
	case reflect.Float32:
		var v float32
		if err := safecast.ToFloat32(from, &v); err == nil {
			toVal.Set(reflect.ValueOf(v))
			return nil
		}
	case reflect.Float64:
		var v float64
		if err := safecast.ToFloat64(from, &v); err == nil {
			toVal.Set(reflect.ValueOf(v))
			return nil
		}
	case reflect.Bool:
		var v bool
		if err := safecast.ToBool(from, &v); err == nil {
			toVal.Set(reflect.ValueOf(v))
			return nil
		}
	case reflect.String:
		var v string
		if err := safecast.ToString(from, &v); err == nil {
			toVal.Set(reflect.ValueOf(v))
			return nil
		}
	case reflect.Array, reflect.Slice:
		return dec.unmarshalArrayTo(fromVal, toVal)
	}
	return newErrorUnmarshalReflectValues(fromVal, toVal)
}

func (dec *Decoder) unmarshalMapToStruct(fromMap map[any]any, toStructVal reflect.Value) error {
	if toStructVal.Type().Kind() != reflect.Struct {
		return newErrorUnmarshalDataTypes(fromMap, toStructVal)
	}
	for fromMapKey, fromMapElem := range fromMap {
		key, ok := fromMapKey.(string)
		if !ok {
			return newErrorUnmarshalDataTypes(fromMap, toStructVal)
		}
		toStructField := toStructVal.FieldByName(key)
		if !toStructField.IsValid() {
			return newErrorUnmarshalDataTypes(fromMap, toStructVal)
		}
		fromMapElemVal := reflect.ValueOf(fromMapElem)
		fromMapElemKind := fromMapElemVal.Type().Kind()
		toStructFieldKind := toStructField.Type().Kind()
		if fromMapElemKind == toStructFieldKind {
			toStructField.Set(fromMapElemVal)
			continue
		}
		switch toStructFieldKind { //nolint:exhaustive
		case reflect.Struct:
			fromMapElemMap, ok := fromMapElem.(map[any]any)
			if !ok {
				return newErrorUnmarshalDataTypes(fromMap, toStructVal)
			}
			if err := dec.unmarshalMapToStruct(fromMapElemMap, toStructField); err != nil {
				return err
			}
		default:
			if err := dec.unmarshalMapElemToStructField(fromMapElemVal, toStructField); err != nil {
				return err
			}
		}
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

func (dec *Decoder) unmarshalToBasicType(fromObj any, toObj any) error {
	switch from := fromObj.(type) {
	case int:
		return safecast.FromInt(from, toObj)
	case int8:
		return safecast.FromInt8(from, toObj)
	case int16:
		return safecast.FromInt16(from, toObj)
	case int32:
		return safecast.FromInt32(from, toObj)
	case int64:
		return safecast.FromInt64(from, toObj)
	case uint:
		return safecast.FromUint(from, toObj)
	case uint8:
		return safecast.FromUint8(from, toObj)
	case uint16:
		return safecast.FromUint16(from, toObj)
	case uint32:
		return safecast.FromUint32(from, toObj)
	case uint64:
		return safecast.FromUint64(from, toObj)
	case float32:
		return safecast.FromFloat32(from, toObj)
	case float64:
		return safecast.FromFloat64(from, toObj)
	case bool:
		return safecast.FromBool(from, toObj)
	case []byte:
		switch to := toObj.(type) {
		case *string:
			*to = string(from)
		case *[]byte:
			*to = from
		}
		return nil
	case string:
		return safecast.FromString(from, toObj)
	default:
	}
	return newErrorUnmarshalDataTypes(fromObj, toObj)
}

func (dec *Decoder) unmarshalToSpecialStruct(fromObj any, toObj any) error {
	switch from := fromObj.(type) {
	case time.Time:
		switch to := toObj.(type) {
		case *string:
			*to = from.Format(time.RFC3339)
		case *time.Time:
			*to = from
		}
	default:
		return newErrorUnmarshalDataTypes(fromObj, toObj)
	}
	return nil
}
