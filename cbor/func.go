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
	"errors"
	"io"
	"math"
	"math/big"
	"reflect"
)

////////////////////////////////////////////////////////////
// byte
////////////////////////////////////////////////////////////

func writeByte(w io.Writer, val byte) error {
	_, err := w.Write([]byte{val})
	return err
}

func writeBytes(w io.Writer, val []byte) error {
	_, err := w.Write(val)
	return err
}

func writeString(w io.Writer, val string) error {
	return writeBytes(w, []byte(val))
}

// maxPreallocBytes is the largest byte string length that readBytes allocates up front.
// Longer strings are read incrementally so that a forged length in the header cannot
// force a huge allocation before the data is actually available (RFC 8949 Section 10).
const maxPreallocBytes = 64 * 1024

func readBytes(r io.Reader, n int) ([]byte, error) {
	if n <= maxPreallocBytes {
		buf := make([]byte, n)
		if _, err := io.ReadFull(r, buf); err != nil {
			return nil, err
		}
		return buf, nil
	}
	var buf bytes.Buffer
	if _, err := io.CopyN(&buf, r, int64(n)); err != nil {
		if errors.Is(err, io.EOF) {
			return nil, io.ErrUnexpectedEOF
		}
		return nil, err
	}
	return buf.Bytes(), nil
}

////////////////////////////////////////////////////////////
// header
////////////////////////////////////////////////////////////

func writeHeader(w io.Writer, m majorType, i majorInfo) error {
	header := byte(m)
	header |= byte(i)
	return writeByte(w, header)
}

////////////////////////////////////////////////////////////
// int8
////////////////////////////////////////////////////////////

func readInt8Bytes(r io.Reader) (int8, error) {
	buf := []byte{0}
	if _, err := io.ReadFull(r, buf); err != nil {
		return 0, err
	}
	return int8(buf[0]), nil
}

func writeInt8Bytes(w io.Writer, v int8) error {
	_, err := w.Write([]byte{byte(v)})
	return err
}

////////////////////////////////////////////////////////////
// uint8
////////////////////////////////////////////////////////////

func readUint8Bytes(r io.Reader) (uint8, error) {
	buf := []byte{0}
	if _, err := io.ReadFull(r, buf); err != nil {
		return 0, err
	}
	return buf[0], nil
}

func writeUint8Bytes(w io.Writer, v uint8) error {
	_, err := w.Write([]byte{v})
	return err
}

////////////////////////////////////////////////////////////
// nint8 (CBOR)
////////////////////////////////////////////////////////////

func writeNint8Bytes(w io.Writer, v int8) error {
	return writeUint8Bytes(w, uint8(-(v + 1)))
}

////////////////////////////////////////////////////////////
// int16
////////////////////////////////////////////////////////////

func readInt16Bytes(r io.Reader) (int16, error) {
	buf := []byte{0, 0}
	if _, err := io.ReadFull(r, buf); err != nil {
		return 0, err
	}
	return (int16(buf[0])<<8 | int16(buf[1])), nil
}

func writeInt16Bytes(w io.Writer, v int16) error {
	_, err := w.Write([]byte{
		byte(v >> 8),
		byte(v)})
	return err
}

////////////////////////////////////////////////////////////
// uint16
////////////////////////////////////////////////////////////

func readUint16Bytes(r io.Reader) (uint16, error) {
	buf := []byte{0, 0}
	if _, err := io.ReadFull(r, buf); err != nil {
		return 0, err
	}
	return (uint16(buf[0])<<8 | uint16(buf[1])), nil
}

func writeUint16Bytes(w io.Writer, v uint16) error {
	_, err := w.Write([]byte{
		byte(v >> 8),
		byte(v)})
	return err
}

////////////////////////////////////////////////////////////
// nint16 (CBOR)
////////////////////////////////////////////////////////////

func writeNint16Bytes(w io.Writer, v int16) error {
	return writeUint16Bytes(w, uint16(-(v + 1)))
}

////////////////////////////////////////////////////////////
// int32
////////////////////////////////////////////////////////////

func readInt32Bytes(r io.Reader) (int32, error) {
	buf := []byte{0, 0, 0, 0}
	if _, err := io.ReadFull(r, buf); err != nil {
		return 0, err
	}
	return (int32(buf[0])<<24 | int32(buf[1])<<16 | int32(buf[2])<<8 | int32(buf[3])), nil
}

func writeInt32Bytes(w io.Writer, v int32) error {
	_, err := w.Write([]byte{
		byte(v >> 24),
		byte(v >> 16),
		byte(v >> 8),
		byte(v)})
	return err
}

////////////////////////////////////////////////////////////
// uint32
////////////////////////////////////////////////////////////

func readUint32Bytes(r io.Reader) (uint32, error) {
	buf := []byte{0, 0, 0, 0}
	if _, err := io.ReadFull(r, buf); err != nil {
		return 0, err
	}
	return (uint32(buf[0])<<24 | uint32(buf[1])<<16 | uint32(buf[2])<<8 | uint32(buf[3])), nil
}

func writeUint32Bytes(w io.Writer, v uint32) error {
	_, err := w.Write([]byte{
		byte(v >> 24),
		byte(v >> 16),
		byte(v >> 8),
		byte(v)})
	return err
}

////////////////////////////////////////////////////////////
// nint32 (CBOR)
////////////////////////////////////////////////////////////

func writeNint32Bytes(w io.Writer, v int32) error {
	return writeUint32Bytes(w, uint32(-(v + 1)))
}

////////////////////////////////////////////////////////////
// int64
////////////////////////////////////////////////////////////

func readInt64Bytes(r io.Reader) (int64, error) {
	buf := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	if _, err := io.ReadFull(r, buf); err != nil {
		return 0, err
	}
	return (int64(buf[0])<<56 | int64(buf[1])<<48 | int64(buf[2])<<40 | int64(buf[3])<<32 | int64(buf[4])<<24 | int64(buf[5])<<16 | int64(buf[6])<<8 | int64(buf[7])), nil
}

func writeInt64Bytes(w io.Writer, v int64) error {
	_, err := w.Write([]byte{
		byte(v >> 56),
		byte(v >> 48),
		byte(v >> 40),
		byte(v >> 32),
		byte(v >> 24),
		byte(v >> 16),
		byte(v >> 8),
		byte(v)})
	return err
}

////////////////////////////////////////////////////////////
// uint64
////////////////////////////////////////////////////////////

func readUint64Bytes(r io.Reader) (uint64, error) {
	buf := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	if _, err := io.ReadFull(r, buf); err != nil {
		return 0, err
	}
	return (uint64(buf[0])<<56 | uint64(buf[1])<<48 | uint64(buf[2])<<40 | uint64(buf[3])<<32 | uint64(buf[4])<<24 | uint64(buf[5])<<16 | uint64(buf[6])<<8 | uint64(buf[7])), nil
}

func writeUint64Bytes(w io.Writer, v uint64) error {
	_, err := w.Write([]byte{
		byte(v >> 56),
		byte(v >> 48),
		byte(v >> 40),
		byte(v >> 32),
		byte(v >> 24),
		byte(v >> 16),
		byte(v >> 8),
		byte(v)})
	return err
}

////////////////////////////////////////////////////////////
// nint64 (CBOR)
////////////////////////////////////////////////////////////

func writeNint64Bytes(w io.Writer, v int64) error {
	return writeUint64Bytes(w, uint64(-(v + 1)))
}

////////////////////////////////////////////////////////////
// nint (CBOR major type 1)
////////////////////////////////////////////////////////////

// nintValue returns the value of a CBOR negative integer (major type 1) whose
// argument is n, that is -1 - n (RFC 8949 Section 3.1).
// The value is returned as the smallest Go signed integer type, no smaller than
// the argument size (in bytes), that can hold it. Values less than math.MinInt64
// (-2^64 <= value < -2^63) are returned as *big.Int.
func nintValue(n uint64, size int) any {
	switch {
	case size <= 1 && n <= math.MaxInt8:
		return -int8(n) - 1
	case size <= 2 && n <= math.MaxInt16:
		return -int16(n) - 1
	case size <= 4 && n <= math.MaxInt32:
		return -int32(n) - 1
	case n <= math.MaxInt64:
		return -int64(n) - 1
	}
	v := new(big.Int).SetUint64(n)
	v.Add(v, big.NewInt(1))
	return v.Neg(v)
}

func readNint8Bytes(r io.Reader) (any, error) {
	v, err := readUint8Bytes(r)
	if err != nil {
		return nil, err
	}
	return nintValue(uint64(v), 1), nil
}

func readNint16Bytes(r io.Reader) (any, error) {
	v, err := readUint16Bytes(r)
	if err != nil {
		return nil, err
	}
	return nintValue(uint64(v), 2), nil
}

func readNint32Bytes(r io.Reader) (any, error) {
	v, err := readUint32Bytes(r)
	if err != nil {
		return nil, err
	}
	return nintValue(uint64(v), 4), nil
}

func readNint64Bytes(r io.Reader) (any, error) {
	v, err := readUint64Bytes(r)
	if err != nil {
		return nil, err
	}
	return nintValue(v, 8), nil
}

////////////////////////////////////////////////////////////
// float16 (IEEE 754 half-precision, read-only)
////////////////////////////////////////////////////////////

func float16ToFloat64(bits uint16) float64 {
	sign := uint64((bits >> 15) & 0x1)
	exp := uint64((bits >> 10) & 0x1F)
	mant := uint64(bits & 0x3FF)

	var f64bits uint64
	switch {
	case exp == 0 && mant == 0:
		// ±zero
		f64bits = sign << 63
	case exp == 0:
		// Subnormal float16 → normalized float64
		e := uint64(1023 - 14) // bias adjustment: float64_bias - float16_min_exp
		for (mant & 0x400) == 0 {
			mant <<= 1
			e--
		}
		mant &= 0x3FF // remove implicit leading 1
		f64bits = (sign << 63) | (e << 52) | (mant << 42)
	case exp == 31 && mant == 0:
		// ±Inf
		f64bits = (sign << 63) | (0x7FF << 52)
	case exp == 31:
		// NaN (preserve mantissa payload)
		f64bits = (sign << 63) | (0x7FF << 52) | (mant << 42)
	default:
		// Normal: rebias exponent from 15 to 1023 (offset = 1008)
		f64bits = (sign << 63) | ((exp + 1008) << 52) | (mant << 42)
	}

	return math.Float64frombits(f64bits)
}

const (
	float16NaN         uint16 = 0x7E00
	float16PositiveInf uint16 = 0x7C00
	float16NegativeInf uint16 = 0xFC00
)

// float64ToFloat16 returns the IEEE 754 half-precision bits of the specified finite value
// if the value can be represented exactly in half precision.
func float64ToFloat16(v float64) (uint16, bool) {
	f32 := float32(v)
	if float64(f32) != v || math.IsInf(v, 0) || math.IsNaN(v) {
		return 0, false
	}
	bits := math.Float32bits(f32)
	sign := uint16(bits>>16) & 0x8000
	exp := int((bits>>23)&0xFF) - 127
	mant := bits & 0x7FFFFF

	switch {
	case exp == -127:
		// ±zero, or a float32 subnormal that is too small for float16
		if mant == 0 {
			return sign, true
		}
		return 0, false
	case 15 < exp:
		return 0, false
	case -14 <= exp:
		// Normal float16 with a 10-bit mantissa
		if mant&0x1FFF != 0 {
			return 0, false
		}
		return sign | uint16(exp+15)<<10 | uint16(mant>>13), true
	case -24 <= exp:
		// Subnormal float16: the value is m * 2^-24
		full := mant | 0x800000
		shift := uint(-(exp + 1))
		if full&((1<<shift)-1) != 0 {
			return 0, false
		}
		return sign | uint16(full>>shift), true
	}
	return 0, false
}

func readFloat16Bytes(r io.Reader) (float64, error) {
	v, err := readUint16Bytes(r)
	if err != nil {
		return 0, err
	}
	return float16ToFloat64(v), nil
}

////////////////////////////////////////////////////////////
// float32
////////////////////////////////////////////////////////////

func readFloat32Bytes(r io.Reader) (float32, error) {
	v, err := readUint32Bytes(r)
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(v), nil
}

func writeFloat32Bytes(w io.Writer, v float32) error {
	return writeUint32Bytes(w, math.Float32bits(v))
}

////////////////////////////////////////////////////////////
// float64
////////////////////////////////////////////////////////////

func readFloat64Bytes(r io.Reader) (float64, error) {
	v, err := readUint64Bytes(r)
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(v), nil
}

func writeFloat64Bytes(w io.Writer, v float64) error {
	return writeUint64Bytes(w, math.Float64bits(v))
}

////////////////////////////////////////////////////////////
// Array
////////////////////////////////////////////////////////////

// nolint: exhaustive
func arrayToAnyArray(fromArray any) ([]any, error) {
	fromArrayVal := reflect.ValueOf(fromArray)
	fromArrayType := fromArrayVal.Type()
	switch fromArrayType.Kind() {
	case reflect.Array:
	case reflect.Slice:
	default:
		return nil, newErrorUnmarshalCastTypes(fromArray, make([]any, 0))
	}

	fromArrayLen := fromArrayVal.Len()
	toArray := make([]any, fromArrayLen)

	toArrayVal := reflect.ValueOf(toArray)
	toArrayType := toArrayVal.Type()
	toArrayElemType := toArrayType.Elem()
	for n := range fromArrayLen {
		fromArrayIndex := fromArrayVal.Index(n)
		toArrayIndex := toArrayVal.Index(n)
		toArrayIndex.Set(fromArrayIndex.Convert(toArrayElemType))
	}

	return toArray, nil
}

////////////////////////////////////////////////////////////
// Map
////////////////////////////////////////////////////////////

func mapToAnyMap(fromMap any) (map[any]any, error) {
	toMap := map[any]any{}

	fromMapVal := reflect.ValueOf(fromMap)
	fromMapType := fromMapVal.Type()
	if fromMapType.Kind() != reflect.Map {
		return nil, newErrorUnmarshalCastTypes(fromMap, toMap)
	}

	toMapVal := reflect.ValueOf(toMap)
	toMapType := toMapVal.Type()
	toMapKeyType := toMapType.Key()
	toMapElemType := toMapType.Elem()

	fromMapIter := fromMapVal.MapRange()
	for fromMapIter.Next() {
		fromMapKeyVal := fromMapIter.Key()
		if !fromMapKeyVal.CanConvert(toMapKeyType) {
			return nil, newErrorUnmarshalCastTypes(fromMap, toMap)
		}
		fromMapElemVal := fromMapIter.Value()
		if !fromMapElemVal.CanConvert(toMapElemType) {
			return nil, newErrorUnmarshalCastTypes(fromMap, toMap)
		}
		toMapVal.SetMapIndex(fromMapKeyVal.Convert(toMapKeyType), fromMapElemVal.Convert(toMapElemType))
	}
	return toMap, nil
}
