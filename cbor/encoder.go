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
	encodeUint8 := func(v uint8) error {
		header := byte(Uint)
		if v < 24 {
			header |= v
			return writeByte(enc.writer, header)
		}
		header |= byte(uIntOneByte)
		if err := writeByte(enc.writer, header); err != nil {
			return err
		}
		return writeUint8Bytes(enc.writer, v)
	}

	encodeUint16 := func(v uint16) error {
		header := byte(Uint)
		header |= byte(uIntTwoByte)
		if err := writeByte(enc.writer, header); err != nil {
			return err
		}
		return writeUint16Bytes(enc.writer, v)
	}

	encodeUint32 := func(v uint32) error {
		header := byte(Uint)
		header |= byte(uIntFourByte)
		if err := writeByte(enc.writer, header); err != nil {
			return err
		}
		return writeUint32Bytes(enc.writer, v)
	}

	encodeUint64 := func(v uint64) error {
		header := byte(Uint)
		header |= byte(uIntEightByte)
		if err := writeByte(enc.writer, header); err != nil {
			return err
		}
		return writeUint64Bytes(enc.writer, v)
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
		if 0 < v {
			return encodeUint8(uint8(v))
		}
		header := byte(NInt)
		if (-v) < 24 {
			header |= uint8(-v) - 1
			return writeByte(enc.writer, header)
		}
		header |= byte(uIntOneByte)
		if err := writeByte(enc.writer, header); err != nil {
			return err
		}
		return writeNint8Bytes(enc.writer, v)
	case int16:
		if 0 < v {
			return encodeUint16(uint16(v))
		}
		header := byte(NInt)
		header |= byte(uIntTwoByte)
		if err := writeByte(enc.writer, header); err != nil {
			return err
		}
		return writeNint16Bytes(enc.writer, v)
	case int32:
		if 0 < v {
			return encodeUint32(uint32(v))
		}
		header := byte(NInt)
		header |= byte(uIntFourByte)
		if err := writeByte(enc.writer, header); err != nil {
			return err
		}
		return writeNint32Bytes(enc.writer, v)
	case int64:
		if 0 < v {
			return encodeUint64(uint64(v))
		}
		header := byte(NInt)
		header |= byte(uIntEightByte)
		if err := writeByte(enc.writer, header); err != nil {
			return err
		}
		return writeNint64Bytes(enc.writer, v)
	case int:
		if 0 < v {
			return encodeUint64(uint64(v))
		}
		header := byte(NInt)
		header |= byte(uIntEightByte)
		if err := writeByte(enc.writer, header); err != nil {
			return err
		}
		return writeNint64Bytes(enc.writer, int64(v))
	case string:
		if _, err := io.WriteString(enc.writer, v); err != nil {
			return err
		}
		return nil
	}
	return newErrorNotSupportedNativeType(item)
}
