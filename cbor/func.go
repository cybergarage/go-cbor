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
