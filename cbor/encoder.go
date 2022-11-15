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

// An Encoder writes CBOR values to an output stream.
type Encoder struct {
	writer io.Writer
}

// NewEncoder returns a new encoder that writes to the specified writer.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		writer: w,
	}
}

// Encode writes the specified object to the specified writer.
func (enc *Encoder) Encode(item any) error {
	// Special data types that cannot be determined by reflect package
	switch item.(type) {
	case []byte: // Recognize as a byte array instead of a uint8 array。
		return enc.encodePrimitiveTypes(item)
	case time.Time:
		return enc.encodeStdStruct(item)
	case nil:
		return enc.encodePrimitiveTypes(item)
	}

	switch reflect.TypeOf(item).Kind() {
	// Major type 5: A map of pairs of data items.
	case reflect.Map:
		return enc.encodeMap(item)
	// Major type 4: An array of data items.
	case reflect.Array, reflect.Slice:
		return enc.encodeArray(item)
	case reflect.Struct:
		return enc.encodeStruct(item)
	// 3. Specification of the CBOR Encoding.
	case reflect.Bool,
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Float32,
		reflect.Float64,
		reflect.String,
		reflect.Pointer:
		return enc.encodePrimitiveTypes(item)
	case reflect.Complex64,
		reflect.Complex128:
	case reflect.Invalid,
		reflect.Chan,
		reflect.Func,
		reflect.Interface,
		reflect.Uintptr,
		reflect.UnsafePointer:
		return newErrorNotSupportedNativeType(item)
	}

	return newErrorNotSupportedNativeType(item)
}

func (enc *Encoder) encodeNumberOfBytes(mt majorType, n int) error {
	header := byte(mt)
	switch {
	case n < int(aiOneByte):
		header |= uint8(n)
	case n < math.MaxUint8:
		header |= byte(aiOneByte)
	case n < math.MaxUint16:
		header |= byte(aiTwoByte)
	case n < math.MaxUint32:
		header |= byte(aiFourByte)
	default:
		header |= byte(aiEightByte)
	}
	if err := writeByte(enc.writer, header); err != nil {
		return err
	}

	switch {
	case n < int(aiOneByte):
		return nil
	case n < math.MaxUint8:
		return writeUint8Bytes(enc.writer, uint8(n))
	case n < math.MaxUint16:
		return writeUint16Bytes(enc.writer, uint16(n))
	case n < math.MaxUint32:
		return writeUint32Bytes(enc.writer, uint32(n))
	default:
		return writeUint64Bytes(enc.writer, uint64(n))
	}
}

func (enc *Encoder) encodeTextString(v string) error {
	n := len(v)
	if err := enc.encodeNumberOfBytes(mtText, n); err != nil {
		return err
	}
	return writeString(enc.writer, v)
}

func (enc *Encoder) encodeByteString(v []byte) error {
	n := len(v)
	if err := enc.encodeNumberOfBytes(mtBytes, n); err != nil {
		return err
	}
	return writeBytes(enc.writer, v)
}

// nolint: gocyclo, maintidx
func (enc *Encoder) encodePrimitiveTypes(item any) error {
	encodeNull := func() error {
		return writeByte(enc.writer, byte(mtFloat)|byte(simpNull))
	}

	encodeBool := func(v bool) error {
		header := byte(mtFloat)
		if v {
			header |= byte(simpTrue)
		} else {
			header |= byte(simpFalse)
		}
		return writeByte(enc.writer, header)
	}

	encodeUint8 := func(v uint8) error {
		header := byte(mtUint)
		if v < 24 {
			header |= v
			return writeByte(enc.writer, header)
		}
		header |= byte(aiOneByte)
		if err := writeByte(enc.writer, header); err != nil {
			return err
		}
		return writeUint8Bytes(enc.writer, v)
	}

	encodeUint16 := func(v uint16) error {
		if err := writeHeader(enc.writer, mtUint, aiTwoByte); err != nil {
			return err
		}
		return writeUint16Bytes(enc.writer, v)
	}

	encodeUint32 := func(v uint32) error {
		if err := writeHeader(enc.writer, mtUint, aiFourByte); err != nil {
			return err
		}
		return writeUint32Bytes(enc.writer, v)
	}

	encodeUint64 := func(v uint64) error {
		if err := writeHeader(enc.writer, mtUint, aiEightByte); err != nil {
			return err
		}
		return writeUint64Bytes(enc.writer, v)
	}

	// 3. Specification of the CBOR Encoding.

	switch v := item.(type) {
	case uint8:
		return encodeUint8(v)
	case uint16:
		return encodeUint16(v)
	case uint32:
		return encodeUint32(v)
	case uint64:
		return encodeUint64(v)
	case uint:
		return encodeUint64(uint64(v))
	case int8:
		if 0 <= v {
			return encodeUint8(uint8(v))
		}
		header := byte(mtNInt)
		if (-v) < 24 {
			header |= uint8(-v) - 1
			return writeByte(enc.writer, header)
		}
		header |= byte(aiOneByte)
		if err := writeByte(enc.writer, header); err != nil {
			return err
		}
		return writeNint8Bytes(enc.writer, v)
	case int16:
		if 0 <= v {
			return encodeUint16(uint16(v))
		}
		if err := writeHeader(enc.writer, mtNInt, aiTwoByte); err != nil {
			return err
		}
		return writeNint16Bytes(enc.writer, v)
	case int32:
		if 0 <= v {
			return encodeUint32(uint32(v))
		}
		if err := writeHeader(enc.writer, mtNInt, aiFourByte); err != nil {
			return err
		}
		return writeNint32Bytes(enc.writer, v)
	case int64:
		if 0 <= v {
			return encodeUint64(uint64(v))
		}
		if err := writeHeader(enc.writer, mtNInt, aiEightByte); err != nil {
			return err
		}
		return writeNint64Bytes(enc.writer, v)
	case int:
		if 0 <= v {
			return encodeUint64(uint64(v))
		}
		if err := writeHeader(enc.writer, mtNInt, aiEightByte); err != nil {
			return err
		}
		return writeNint64Bytes(enc.writer, int64(v))
	case float32:
		if err := writeHeader(enc.writer, mtFloat, fpnFloat32); err != nil {
			return err
		}
		return writeFloat32Bytes(enc.writer, v)
	case float64:
		if err := writeHeader(enc.writer, mtFloat, fpnFloat64); err != nil {
			return err
		}
		return writeFloat64Bytes(enc.writer, v)
	case bool:
		return encodeBool(v)
	case nil:
		return encodeNull()
	case []byte:
		return enc.encodeByteString(v)
	case string:
		return enc.encodeTextString(v)
	}

	return newErrorNotSupportedNativeType(item)
}

func (enc *Encoder) encodeArray(item any) error {
	writeAnyArray := func(v []any) error {
		cnt := len(v)
		if err := enc.encodeNumberOfBytes(mtArray, cnt); err != nil {
			return err
		}
		for n := 0; n < cnt; n++ {
			if err := enc.Encode(v[n]); err != nil {
				return err
			}
		}
		return nil
	}

	// Major type 4: An array of data items.

	switch v := item.(type) {
	case []int8:
		return writeAnyArray(toAnyArray(v))
	case []int16:
		return writeAnyArray(toAnyArray(v))
	case []int32:
		return writeAnyArray(toAnyArray(v))
	case []int64:
		return writeAnyArray(toAnyArray(v))
	case []int:
		return writeAnyArray(toAnyArray(v))
	case []uint8:
		return writeAnyArray(toAnyArray(v))
	case []uint16:
		return writeAnyArray(toAnyArray(v))
	case []uint32:
		return writeAnyArray(toAnyArray(v))
	case []uint64:
		return writeAnyArray(toAnyArray(v))
	case []uint:
		return writeAnyArray(toAnyArray(v))
	case []float32:
		return writeAnyArray(toAnyArray(v))
	case []float64:
		return writeAnyArray(toAnyArray(v))
	case []bool:
		return writeAnyArray(toAnyArray(v))
	case []string:
		return writeAnyArray(toAnyArray(v))
	case []any: // NOTE: Any array is not match.
		return writeAnyArray(v)
	}

	return newErrorNotSupportedNativeType(item)
}

func (enc *Encoder) encodeMap(item any) error {
	writeAnyMap := func(m map[any]any) error {
		if err := enc.encodeNumberOfBytes(mtMap, len(m)); err != nil {
			return err
		}
		for k, v := range m {
			if err := enc.Encode(k); err != nil {
				return err
			}
			if err := enc.Encode(v); err != nil {
				return err
			}
		}
		return nil
	}

	// Major type 5: A map of pairs of data items.

	v, ok := item.(map[any]any)
	if ok {
		return writeAnyMap(v)
	}

	return newErrorNotSupportedNativeType(item)
}

func (enc *Encoder) encodeStruct(item any) error {
	structMap := map[any]any{}
	itemStruct := reflect.ValueOf(item)
	for n := 0; n < itemStruct.NumField(); n++ {
		typeField := itemStruct.Type().Field(n)
		structMap[typeField.Name] = itemStruct.Field(n).Interface()
	}
	return enc.encodeMap(structMap)
}

func (enc *Encoder) encodeStdStruct(item any) error {
	switch v := item.(type) {
	case time.Time:
		if err := writeHeader(enc.writer, mtTag, tagStdDateTime); err != nil {
			return err
		}
		return enc.encodeTextString(v.Format(time.RFC3339))
	default:
		return newErrorNotSupportedNativeType(item)
	}
}
