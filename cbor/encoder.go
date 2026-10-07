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
	"io"
	"math"
	"math/big"
	"reflect"
	"sort"
	"time"
)

// An Encoder writes CBOR values to an output stream.
type Encoder struct {
	*Config

	writer io.Writer
}

// NewEncoder returns a new encoder that writes to the specified writer.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		Config: NewConfig(),
		writer: w,
	}
}

// Encode writes the specified object to the specified writer.
func (enc *Encoder) Encode(item any) error {
	// Special data types that cannot be determined by reflect package
	switch v := item.(type) {
	case []byte: // Recognize as a byte array instead of a uint8 array。
		return enc.encodePrimitiveTypes(item)
	case time.Time:
		return enc.encodeStdStruct(item)
	case SimpleValue:
		return enc.encodeSimpleValue(v)
	case big.Int:
		return enc.encodeBigInt(&v)
	case *big.Int:
		if v == nil {
			return enc.encodePrimitiveTypes(nil)
		}
		return enc.encodeBigInt(v)
	case Tag:
		return enc.encodeTag(v.Number, v.Content)
	case *Tag:
		if v == nil {
			return enc.encodePrimitiveTypes(nil)
		}
		return enc.encodeTag(v.Number, v.Content)
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
	case reflect.Struct, reflect.Pointer:
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
		reflect.String:
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

func (enc *Encoder) encodeTextString(v string) error {
	n := len(v)
	if err := enc.encodeArgument(mtText, uint64(n)); err != nil {
		return err
	}
	return writeString(enc.writer, v)
}

func (enc *Encoder) encodeByteString(v []byte) error {
	n := len(v)
	if err := enc.encodeArgument(mtBytes, uint64(n)); err != nil {
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

	if enc.encodeMode != EncodeModeTypePreserving {
		// 4.1. Preferred Serialization
		switch v := item.(type) {
		case uint8:
			return enc.encodeArgument(mtUint, uint64(v))
		case uint16:
			return enc.encodeArgument(mtUint, uint64(v))
		case uint32:
			return enc.encodeArgument(mtUint, uint64(v))
		case uint64:
			return enc.encodeArgument(mtUint, v)
		case uint:
			return enc.encodeArgument(mtUint, uint64(v))
		case int8:
			return enc.encodeInt(int64(v))
		case int16:
			return enc.encodeInt(int64(v))
		case int32:
			return enc.encodeInt(int64(v))
		case int64:
			return enc.encodeInt(v)
		case int:
			return enc.encodeInt(int64(v))
		case float32:
			return enc.encodePreferredFloat(float64(v))
		case float64:
			return enc.encodePreferredFloat(v)
		}
	}

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
		iv := -(v + 1)
		if iv < 24 {
			header |= uint8(iv)
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
		if math.IsNaN(float64(v)) {
			if err := writeHeader(enc.writer, mtFloat, fpnFloat16); err != nil {
				return err
			}
			return writeUint16Bytes(enc.writer, 0x7e00) // canonical float16 NaN
		}
		if math.IsInf(float64(v), 1) {
			if err := writeHeader(enc.writer, mtFloat, fpnFloat16); err != nil {
				return err
			}
			return writeUint16Bytes(enc.writer, 0x7c00) // float16 +Inf
		}
		if math.IsInf(float64(v), -1) {
			if err := writeHeader(enc.writer, mtFloat, fpnFloat16); err != nil {
				return err
			}
			return writeUint16Bytes(enc.writer, 0xfc00) // float16 -Inf
		}
		if err := writeHeader(enc.writer, mtFloat, fpnFloat32); err != nil {
			return err
		}
		return writeFloat32Bytes(enc.writer, v)
	case float64:
		if math.IsNaN(v) {
			if err := writeHeader(enc.writer, mtFloat, fpnFloat16); err != nil {
				return err
			}
			return writeUint16Bytes(enc.writer, 0x7e00) // canonical float16 NaN
		}
		if math.IsInf(v, 1) {
			if err := writeHeader(enc.writer, mtFloat, fpnFloat16); err != nil {
				return err
			}
			return writeUint16Bytes(enc.writer, 0x7c00) // float16 +Inf
		}
		if math.IsInf(v, -1) {
			if err := writeHeader(enc.writer, mtFloat, fpnFloat16); err != nil {
				return err
			}
			return writeUint16Bytes(enc.writer, 0xfc00) // float16 -Inf
		}
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
		if err := enc.encodeArgument(mtArray, uint64(cnt)); err != nil {
			return err
		}
		for n := range cnt {
			if err := enc.Encode(v[n]); err != nil {
				return err
			}
		}
		return nil
	}

	// Major type 4: An array of data items.

	v, ok := item.([]any)
	if ok {
		return writeAnyArray(v)
	}

	v, err := arrayToAnyArray(item)
	if err != nil {
		return err
	}
	return writeAnyArray(v)
}

// isMapKeySortRequired returns true if map keys must be sorted.
func (enc *Encoder) isMapKeySortRequired() bool {
	switch enc.encodeMode {
	case EncodeModeCoreDeterministic, EncodeModeLengthFirstDeterministic:
		return true
	case EncodeModePreferred, EncodeModeTypePreserving:
		return enc.MapSortEnabled
	}
	return enc.MapSortEnabled
}

// encodeToBytes returns the encoding of the specified item with the same configuration.
func (enc *Encoder) encodeToBytes(item any) ([]byte, error) {
	var buf bytes.Buffer
	sub := &Encoder{Config: enc.Config, writer: &buf}
	if err := sub.Encode(item); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (enc *Encoder) encodeAnyMap(m map[any]any) error {
	if err := enc.encodeArgument(mtMap, uint64(len(m))); err != nil {
		return err
	}

	if !enc.isMapKeySortRequired() {
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

	// 4.2.1. Core Deterministic Encoding Requirements: the keys are sorted by their encodings.
	type encodedPair struct {
		key []byte
		val any
	}
	pairs := make([]encodedPair, 0, len(m))
	for k, v := range m {
		key, err := enc.encodeToBytes(k)
		if err != nil {
			return err
		}
		pairs = append(pairs, encodedPair{key: key, val: v})
	}
	lengthFirst := enc.encodeMode == EncodeModeLengthFirstDeterministic
	sort.Slice(pairs, func(i, j int) bool {
		// 4.2.3. Length-First Map Key Ordering
		if lengthFirst && len(pairs[i].key) != len(pairs[j].key) {
			return len(pairs[i].key) < len(pairs[j].key)
		}
		return bytes.Compare(pairs[i].key, pairs[j].key) < 0
	})
	for n, pair := range pairs {
		// Different Go keys, such as int8(1) and int(1), can have the same encoding,
		// which would produce a map with duplicate keys (RFC 8949 Section 5.6).
		if 0 < n && bytes.Equal(pairs[n-1].key, pair.key) {
			return newErrorEncodeDuplicateMapKey(pair.key)
		}
		if err := writeBytes(enc.writer, pair.key); err != nil {
			return err
		}
		if err := enc.Encode(pair.val); err != nil {
			return err
		}
	}
	return nil
}

func (enc *Encoder) encodeMap(item any) error {
	// Major type 5: A map of pairs of data items.

	v, ok := item.(map[any]any)
	if ok {
		return enc.encodeAnyMap(v)
	}

	v, err := mapToAnyMap(item)
	if err != nil {
		return err
	}
	return enc.encodeAnyMap(v)
}

// encodeInt writes a signed integer in the shortest form.
func (enc *Encoder) encodeInt(v int64) error {
	if 0 <= v {
		return enc.encodeArgument(mtUint, uint64(v))
	}
	// A negative integer n is encoded as -1 - n.
	return enc.encodeArgument(mtNInt, uint64(-(v + 1)))
}

// encodePreferredFloat writes a floating-point value in the shortest form that preserves the value (RFC 8949 Section 4.1).
func (enc *Encoder) encodePreferredFloat(v float64) error {
	writeFloat16 := func(bits uint16) error {
		if err := writeHeader(enc.writer, mtFloat, fpnFloat16); err != nil {
			return err
		}
		return writeUint16Bytes(enc.writer, bits)
	}
	switch {
	case math.IsNaN(v):
		return writeFloat16(float16NaN)
	case math.IsInf(v, 1):
		return writeFloat16(float16PositiveInf)
	case math.IsInf(v, -1):
		return writeFloat16(float16NegativeInf)
	}
	if bits, ok := float64ToFloat16(v); ok {
		return writeFloat16(bits)
	}
	if f32 := float32(v); float64(f32) == v {
		if err := writeHeader(enc.writer, mtFloat, fpnFloat32); err != nil {
			return err
		}
		return writeFloat32Bytes(enc.writer, f32)
	}
	if err := writeHeader(enc.writer, mtFloat, fpnFloat64); err != nil {
		return err
	}
	return writeFloat64Bytes(enc.writer, v)
}

// encodeArgument writes the initial byte and the argument in the shortest form (RFC 8949 Section 3).
func (enc *Encoder) encodeArgument(mt majorType, n uint64) error {
	switch {
	case n < uint64(aiOneByte):
		return writeHeader(enc.writer, mt, majorInfo(n))
	case n <= math.MaxUint8:
		if err := writeHeader(enc.writer, mt, aiOneByte); err != nil {
			return err
		}
		return writeUint8Bytes(enc.writer, uint8(n))
	case n <= math.MaxUint16:
		if err := writeHeader(enc.writer, mt, aiTwoByte); err != nil {
			return err
		}
		return writeUint16Bytes(enc.writer, uint16(n))
	case n <= math.MaxUint32:
		if err := writeHeader(enc.writer, mt, aiFourByte); err != nil {
			return err
		}
		return writeUint32Bytes(enc.writer, uint32(n))
	}
	if err := writeHeader(enc.writer, mt, aiEightByte); err != nil {
		return err
	}
	return writeUint64Bytes(enc.writer, n)
}

// encodeSimpleValue writes a simple value (RFC 8949 Section 3.3).
func (enc *Encoder) encodeSimpleValue(v SimpleValue) error {
	switch {
	case v < SimpleValue(simpOneByte):
		return writeHeader(enc.writer, mtFloat, majorInfo(v))
	case v < simpMinOneByte:
		return newErrorEncodeReservedSimpleValue(v)
	}
	if err := writeHeader(enc.writer, mtFloat, simpOneByte); err != nil {
		return err
	}
	return writeUint8Bytes(enc.writer, uint8(v))
}

// encodeBigInt writes an integer as major type 0 or 1 if it fits in 64 bits,
// otherwise as a bignum (tag 2 or 3) as defined in RFC 8949 Section 3.4.3.
func (enc *Encoder) encodeBigInt(v *big.Int) error {
	if 0 <= v.Sign() {
		if v.IsUint64() {
			return enc.encodeArgument(mtUint, v.Uint64())
		}
		return enc.encodeTag(tagPositiveBignum, v.Bytes())
	}
	// A negative integer n is encoded as -1 - n.
	n := new(big.Int).Neg(v)
	n.Sub(n, big.NewInt(1))
	if n.IsUint64() {
		return enc.encodeArgument(mtNInt, n.Uint64())
	}
	return enc.encodeTag(tagNegativeBignum, n.Bytes())
}

// encodeTag writes a tagged data item (RFC 8949 Section 3.4).
func (enc *Encoder) encodeTag(number uint64, content any) error {
	if err := enc.encodeArgument(mtTag, number); err != nil {
		return err
	}
	return enc.Encode(content)
}

func (enc *Encoder) encodeStdStruct(item any) error {
	switch v := item.(type) {
	case time.Time:
		if err := enc.encodeArgument(mtTag, tagStdDateTime); err != nil {
			return err
		}
		// RFC3339Nano keeps fractional seconds and omits them when they are zero.
		return enc.encodeTextString(v.Format(time.RFC3339Nano))
	default:
		return newErrorNotSupportedNativeType(item)
	}
}

// nolint: exhaustive
func (enc *Encoder) encodeStruct(item any) error {
	var itemStruct reflect.Value
	switch reflect.TypeOf(item).Kind() {
	case reflect.Struct:
		itemStruct = reflect.ValueOf(item)
	case reflect.Pointer:
		itemStruct = reflect.ValueOf(item).Elem()
		if itemStruct.Type().Kind() != reflect.Struct {
			return newErrorNotSupportedNativeType(item)
		}
	default:
		return newErrorNotSupportedNativeType(item)
	}

	// Unexported fields are skipped like encoding/json, because their values cannot be read through reflection.
	fields := make([]int, 0, itemStruct.NumField())
	for n := range itemStruct.NumField() {
		if itemStruct.Type().Field(n).IsExported() {
			fields = append(fields, n)
		}
	}

	if enc.isMapKeySortRequired() {
		structMap := map[any]any{}
		for _, n := range fields {
			typeField := itemStruct.Type().Field(n)
			structMap[typeField.Name] = itemStruct.Field(n).Interface()
		}
		return enc.encodeMap(structMap)
	}

	// Encode the fields in the declaration order so that the output is deterministic.
	if err := enc.encodeArgument(mtMap, uint64(len(fields))); err != nil {
		return err
	}
	for _, n := range fields {
		if err := enc.encodeTextString(itemStruct.Type().Field(n).Name); err != nil {
			return err
		}
		if err := enc.Encode(itemStruct.Field(n).Interface()); err != nil {
			return err
		}
	}
	return nil
}
