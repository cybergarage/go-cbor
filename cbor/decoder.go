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
	"io"
	"math"
	"reflect"
	"time"
)

// An Decoder reads CBOR values from an output stream.
type Decoder struct {
	reader io.Reader
	header []byte
}

// NewDecoder returns a new decoder that reads from the specified writer.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		reader: r,
		header: make([]byte, 1),
	}
}

// nolint: gocyclo, maintidx, exhaustive
// Decode returns the next item if available, otherwise returns EOF or error.
func (dec *Decoder) Decode() (any, error) {
	returnDecordedUint8 := func(v uint8) any {
		if math.MaxInt8 < v {
			return v
		}
		return int8(v)
	}

	returnDecordedUint16 := func(v uint16) any {
		if math.MaxInt16 < v {
			return v
		}
		return int16(v)
	}

	returnDecordedUint32 := func(v uint32) any {
		if math.MaxInt32 < v {
			return v
		}
		return int32(v)
	}

	returnDecordedUint64 := func(v uint64) any {
		if math.MaxInt64 < v {
			return v
		}
		return int64(v)
	}

	readNumberOfItems := func(mt majorType, ai majorInfo) (int, error) {
		if ai < aiOneByte {
			return int(ai), nil
		}
		switch ai {
		case aiOneByte:
			v, err := readUint8Bytes(dec.reader)
			if err != nil {
				return 0, err
			}
			return int(v), nil
		case aiTwoByte:
			v, err := readUint16Bytes(dec.reader)
			if err != nil {
				return 0, err
			}
			return int(v), nil
		case aiFourByte:
			v, err := readUint32Bytes(dec.reader)
			if err != nil {
				return 0, err
			}
			return int(v), nil
		case aiEightByte:
			v, err := readUint64Bytes(dec.reader)
			if err != nil {
				return 0, err
			}
			return int(v), nil
		}
		return 0, newErrorNotSupportedAddInfo(mt, ai)
	}

	readByteString := func(m majorType, i majorInfo) ([]byte, error) {
		n, err := readNumberOfItems(m, i)
		if err != nil {
			return nil, err
		}
		return readBytes(dec.reader, n)
	}

	readTextString := func(m majorType, i majorInfo) (string, error) {
		bytes, err := readByteString(m, i)
		if err != nil {
			return "", err
		}
		return string(bytes), nil
	}

	// 3. Specification of the CBOR Encoding.

	if _, err := io.ReadFull(dec.reader, dec.header); err != nil {
		return nil, err
	}

	majorType := majorType(dec.header[0] & majorTypeMask)
	majorInfo := majorInfo(dec.header[0] & majorInfoMask)

	switch majorType {
	case mtUint:
		if majorInfo < aiOneByte {
			return returnDecordedUint8(uint8(majorInfo)), nil
		}
		switch majorInfo {
		case aiOneByte:
			v, err := readUint8Bytes(dec.reader)
			if err != nil {
				return 0, err
			}
			return returnDecordedUint8(v), nil
		case aiTwoByte:
			v, err := readUint16Bytes(dec.reader)
			if err != nil {
				return 0, err
			}
			return returnDecordedUint16(v), nil
		case aiFourByte:
			v, err := readUint32Bytes(dec.reader)
			if err != nil {
				return 0, err
			}
			return returnDecordedUint32(v), nil
		case aiEightByte:
			v, err := readUint64Bytes(dec.reader)
			if err != nil {
				return 0, err
			}
			return returnDecordedUint64(v), nil
		}
		return nil, newErrorNotSupportedAddInfo(mtUint, majorInfo)
	case mtNInt:
		if majorInfo < aiOneByte {
			return -int8(majorInfo + 1), nil
		}
		switch majorInfo {
		case aiOneByte:
			return readNint8Bytes(dec.reader)
		case aiTwoByte:
			return readNint16Bytes(dec.reader)
		case aiFourByte:
			return readNint32Bytes(dec.reader)
		case aiEightByte:
			return readNint64Bytes(dec.reader)
		}
		return nil, newErrorNotSupportedAddInfo(mtNInt, majorInfo)
	case mtBytes:
		return readByteString(mtBytes, majorInfo)
	case mtText:
		return readTextString(mtText, majorInfo)
	case mtArray:
		itemCount, err := readNumberOfItems(mtArray, majorInfo)
		if err != nil {
			return nil, err
		}
		itemArray := make([]any, 0)
		for n := 0; n < itemCount; n++ {
			item, err := dec.Decode()
			if err != nil {
				return nil, err
			}
			itemArray = append(itemArray, item)
		}
		return itemArray, nil
	case mtMap:
		itemArray, err := readNumberOfItems(mtArray, majorInfo)
		if err != nil {
			return nil, err
		}
		itemMap := map[any]any{}
		for n := 0; n < itemArray; n++ {
			key, err := dec.Decode()
			if err != nil {
				return nil, err
			}
			val, err := dec.Decode()
			if err != nil {
				return nil, err
			}
			itemMap[key] = val
		}
		return itemMap, nil
	case mtTag:
		switch majorInfo {
		case tagStdDateTime:
			dateTime, err := dec.Decode()
			if err != nil {
				return nil, err
			}
			dateTimeStr, ok := dateTime.(string)
			if !ok {
				return nil, newErrorNotSupportedAddInfo(mtTag, majorInfo)
			}
			return time.Parse(time.RFC3339, dateTimeStr)
		case tagEpochDateTime:
		}
		return nil, newErrorNotSupportedMajorType(majorType)
	case mtFloat:
		switch majorInfo {
		case simpFalse:
			return false, nil
		case simpTrue:
			return true, nil
		case simpNull:
			return nil, nil
		case fpnFloat16:
			return nil, newErrorNotSupportedAddInfo(mtFloat, majorInfo)
		case fpnFloat32:
			return readFloat32Bytes(dec.reader)
		case fpnFloat64:
			return readFloat64Bytes(dec.reader)
		}
		return nil, newErrorNotSupportedAddInfo(mtFloat, majorInfo)
	}

	return nil, newErrorNotSupportedMajorType(majorType)
}

// nolint: exhaustive
// Unmarshal returns the next item to the specified data type if available, otherwise returns EOF or error.
func (dec *Decoder) Unmarshal(toObj any) error {
	fromObj, err := dec.Decode()
	if err != nil {
		return err
	}

	switch v := fromObj.(type) {
	case map[any]any:
		switch reflect.ValueOf(toObj).Type().Kind() {
		case reflect.Struct, reflect.Pointer:
			return dec.unmarshalMapTo(v, toObj)
		default:
			return newErrorNotSupportedNativeType(toObj)
		}
	case []any:
		return dec.unmarshalArrayTo(v, toObj)
	}

	return newErrorNotSupportedNativeType(toObj)
}

// nolint: exhaustive
func (dec *Decoder) unmarshalArrayTo(fromArray []any, toArray any) error {
	fromArrayType := reflect.TypeOf(fromArray).Kind()
	toArrayType := reflect.TypeOf(toArray).Kind()
	if fromArrayType != toArrayType {
		return newErrorNotSupportedUnmarshalingDataTypes(fromArray, toArray)
	}

	toArrayVal := reflect.ValueOf(toArray)
	toObjType := reflect.TypeOf(toArray).Elem().Kind()
	for _, fromObj := range fromArray {
		fromObjType := reflect.TypeOf(fromObj).Kind()
		if fromObjType != toObjType {
			return newErrorNotSupportedUnmarshalingDataTypes(fromObj, toArray)
		}
		toArrayVal = reflect.Append(toArrayVal, reflect.ValueOf(fromObj))
	}

	return nil
}

// nolint: exhaustive
func (dec *Decoder) unmarshalMapTo(fromMap map[any]any, toObj any) error {
	switch reflect.TypeOf(toObj).Kind() {
	case reflect.Struct:
		return dec.unmarshalMapToStrct(fromMap, reflect.ValueOf(toObj))
	case reflect.Map:
		return dec.unmarshalMapToMap(fromMap, reflect.ValueOf(toObj))
	case reflect.Pointer:
		elem := reflect.ValueOf(toObj).Elem()
		if elem.Type().Kind() != reflect.Struct {
			return newErrorNotSupportedNativeType(toObj)
		}
		return dec.unmarshalMapToStrct(fromMap, elem)
	}

	return newErrorNotSupportedUnmarshalingDataTypes(fromMap, toObj)
}

// nolint: exhaustive
func (dec *Decoder) unmarshalMapToStrct(fromMap map[any]any, toStruct reflect.Value) error {
	for fromMapKey, fromMapValue := range fromMap {
		key, ok := fromMapKey.(string)
		if !ok {
			return newErrorNotSupportedUnmarshalingDataTypes(fromMap, toStruct)
		}
		toStructField := toStruct.FieldByName(key)
		if !ok {
			return newErrorNotSupportedUnmarshalingDataTypes(fromMap, toStruct)
		}
		fromMapValueStruct := reflect.ValueOf(fromMapValue)
		if fromMapValueStruct.Type().Kind() != toStructField.Type().Kind() {
			return newErrorNotSupportedUnmarshalingDataTypes(fromMap, toStruct)
		}
		toStructField.Set(fromMapValueStruct)
	}
	return nil
}

// nolint: exhaustive
func (dec *Decoder) unmarshalMapToMap(fromObj map[any]any, toMap reflect.Value) error {
	return newErrorNotSupportedUnmarshalingDataTypes(fromObj, toMap)
}
