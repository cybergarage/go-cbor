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

func readBytes(r io.Reader, n int) ([]byte, error) {
	buf := make([]byte, n)
	if _, err := io.ReadFull(r, buf); err != nil {
		return nil, err
	}
	return buf, nil
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
	if _, err := r.Read(buf); err != nil {
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
	if _, err := r.Read(buf); err != nil {
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

func readNint8Bytes(r io.Reader) (int8, error) {
	v, err := readUint8Bytes(r)
	if err != nil {
		return 0, err
	}
	return -int8(v + 1), nil
}

func writeNint8Bytes(w io.Writer, v int8) error {
	return writeUint8Bytes(w, uint8(-(v + 1)))
}

////////////////////////////////////////////////////////////
// int16
////////////////////////////////////////////////////////////

func readInt16Bytes(r io.Reader) (int16, error) {
	buf := []byte{0, 0}
	if _, err := r.Read(buf); err != nil {
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
	if _, err := r.Read(buf); err != nil {
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

func readNint16Bytes(r io.Reader) (int16, error) {
	v, err := readUint16Bytes(r)
	if err != nil {
		return 0, err
	}
	return -int16(v + 1), nil
}

func writeNint16Bytes(w io.Writer, v int16) error {
	return writeUint16Bytes(w, uint16(-(v + 1)))
}

////////////////////////////////////////////////////////////
// int32
////////////////////////////////////////////////////////////

func readInt32Bytes(r io.Reader) (int32, error) {
	buf := []byte{0, 0, 0, 0}
	if _, err := r.Read(buf); err != nil {
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
	if _, err := r.Read(buf); err != nil {
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

func readNint32Bytes(r io.Reader) (int32, error) {
	v, err := readUint32Bytes(r)
	if err != nil {
		return 0, err
	}
	return -int32(v + 1), nil
}

func writeNint32Bytes(w io.Writer, v int32) error {
	return writeUint32Bytes(w, uint32(-(v + 1)))
}

////////////////////////////////////////////////////////////
// int64
////////////////////////////////////////////////////////////

func readInt64Bytes(r io.Reader) (int64, error) {
	buf := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	if _, err := r.Read(buf); err != nil {
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
	if _, err := r.Read(buf); err != nil {
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

func readNint64Bytes(r io.Reader) (int64, error) {
	v, err := readUint64Bytes(r)
	if err != nil {
		return 0, err
	}
	return -int64(v + 1), nil
}

func writeNint64Bytes(w io.Writer, v int64) error {
	return writeUint64Bytes(w, uint64(-(v + 1)))
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
