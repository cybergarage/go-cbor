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
	"fmt"
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

func readInt8Bytes(src []byte) (int8, []byte, error) {
	srcLen := len(src)
	if srcLen < 1 {
		return 0, nil, fmt.Errorf(errorInvalidIntegerBytes, src)
	}
	return int8(src[0]), src[1:], nil
}

func writeInt8Bytes(w io.Writer, val int8) error {
	_, err := w.Write([]byte{byte(val)})
	return err
}

////////////////////////////////////////////////////////////
// uint8
////////////////////////////////////////////////////////////

func readUint8Bytes(src []byte) (uint8, []byte, error) {
	srcLen := len(src)
	if srcLen < 1 {
		return 0, nil, fmt.Errorf(errorInvalidIntegerBytes, src)
	}
	return src[0], src[1:], nil
}

func appendUint8Bytes(buf []byte, val uint8) []byte {
	return append(buf,
		val,
	)
}

////////////////////////////////////////////////////////////
// int16
////////////////////////////////////////////////////////////

func readInt16Bytes(src []byte) (int16, []byte, error) {
	srcLen := len(src)
	if srcLen < 2 {
		return 0, nil, fmt.Errorf(errorInvalidIntegerBytes, src)
	}
	return (int16(src[0])<<8 | int16(src[1])), src[2:], nil
}

func appendInt16Bytes(buf []byte, val int16) []byte {
	return append(buf,
		byte(val>>8),
		byte(val),
	)
}

////////////////////////////////////////////////////////////
// uint16
////////////////////////////////////////////////////////////

func readUint16Bytes(src []byte) (uint16, []byte, error) {
	srcLen := len(src)
	if srcLen < 2 {
		return 0, nil, fmt.Errorf(errorInvalidIntegerBytes, src)
	}
	return (uint16(src[0])<<8 | uint16(src[1])), src[2:], nil
}

func appendUint16Bytes(buf []byte, val uint16) []byte {
	return append(buf,
		byte(val>>8),
		byte(val),
	)
}

////////////////////////////////////////////////////////////
// int32
////////////////////////////////////////////////////////////

func readInt32Bytes(src []byte) (int32, []byte, error) {
	srcLen := len(src)
	if srcLen < 4 {
		return 0, nil, fmt.Errorf(errorInvalidIntegerBytes, src)
	}
	return (int32(src[0])<<24 | int32(src[1])<<16 | int32(src[2])<<8 | int32(src[3])), src[4:], nil
}

func appendInt32Bytes(buf []byte, val int32) []byte {
	return append(buf,
		byte(val>>24),
		byte(val>>16),
		byte(val>>8),
		byte(val),
	)
}

////////////////////////////////////////////////////////////
// uint32
////////////////////////////////////////////////////////////

func readUint32Bytes(src []byte) (uint32, []byte, error) {
	srcLen := len(src)
	if srcLen < 4 {
		return 0, nil, fmt.Errorf(errorInvalidIntegerBytes, src)
	}
	return (uint32(src[0])<<24 | uint32(src[1])<<16 | uint32(src[2])<<8 | uint32(src[3])), src[4:], nil
}

func appendUint32Bytes(buf []byte, val uint32) []byte {
	return append(buf,
		byte(val>>24),
		byte(val>>16),
		byte(val>>8),
		byte(val),
	)
}

////////////////////////////////////////////////////////////
// uint64
////////////////////////////////////////////////////////////

func readUint64Bytes(src []byte) (uint64, []byte, error) {
	srcLen := len(src)
	if srcLen < 8 {
		return 0, nil, fmt.Errorf(errorInvalidIntegerBytes, src)
	}
	return (uint64(src[0])<<56 | uint64(src[1])<<48 | uint64(src[2])<<40 | uint64(src[3])<<32 | uint64(src[4])<<24 | uint64(src[5])<<16 | uint64(src[6])<<8 | uint64(src[7])), src[8:], nil
}

func appendUint64Bytes(buf []byte, val uint64) []byte {
	return append(buf,
		byte(val>>56),
		byte(val>>48),
		byte(val>>40),
		byte(val>>32),
		byte(val>>24),
		byte(val>>16),
		byte(val>>8),
		byte(val),
	)
}
