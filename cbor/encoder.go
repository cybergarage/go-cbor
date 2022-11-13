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

// nolint: gocyclo, maintidx
// Encode writes the specified object to the specified writer.
func (enc *Encoder) Encode(item any) error {
	encodeNull := func() error {
		return writeByte(enc.writer, byte(Float)|byte(simpNull))
	}

	encodeBool := func(v bool) error {
		header := byte(Float)
		if v {
			header |= byte(simpTrue)
		} else {
			header |= byte(simpFalse)
		}
		return writeByte(enc.writer, header)
	}

	encodeUint8 := func(v uint8) error {
		header := byte(Uint)
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
		if err := writeHeader(enc.writer, Uint, aiTwoByte); err != nil {
			return err
		}
		return writeUint16Bytes(enc.writer, v)
	}

	encodeUint32 := func(v uint32) error {
		if err := writeHeader(enc.writer, Uint, aiFourByte); err != nil {
			return err
		}
		return writeUint32Bytes(enc.writer, v)
	}

	encodeUint64 := func(v uint64) error {
		if err := writeHeader(enc.writer, Uint, aiEightByte); err != nil {
			return err
		}
		return writeUint64Bytes(enc.writer, v)
	}

	encodeNumberOfBytes := func(mt majorType, n int) error {
		header := byte(mt)
		switch {
		case n < 24:
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
		return writeByte(enc.writer, header)
	}

	writeByteString := func(v []byte) error {
		n := len(v)
		if err := encodeNumberOfBytes(Bytes, n); err != nil {
			return err
		}
		return writeBytes(enc.writer, v)
	}

	writeTextString := func(v string) error {
		n := len(v)
		if err := encodeNumberOfBytes(Text, n); err != nil {
			return err
		}
		return writeString(enc.writer, v)
	}

	writeAnyArray := func(v []any) error {
		cnt := len(v)
		if err := encodeNumberOfBytes(Array, cnt); err != nil {
			return err
		}
		for n := 0; n < cnt; n++ {
			if err := enc.Encode(v[n]); err != nil {
				return err
			}
		}
		return nil
	}

	toAnyArray := func(v []int8) []any {
		a := make([]any, len(v))
		for n, t := range v {
			a[n] = t
		}
		return a
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
		header := byte(NInt)
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
		if err := writeHeader(enc.writer, NInt, aiTwoByte); err != nil {
			return err
		}
		return writeNint16Bytes(enc.writer, v)
	case int32:
		if 0 <= v {
			return encodeUint32(uint32(v))
		}
		if err := writeHeader(enc.writer, NInt, aiFourByte); err != nil {
			return err
		}
		return writeNint32Bytes(enc.writer, v)
	case int64:
		if 0 <= v {
			return encodeUint64(uint64(v))
		}
		if err := writeHeader(enc.writer, NInt, aiEightByte); err != nil {
			return err
		}
		return writeNint64Bytes(enc.writer, v)
	case int:
		if 0 <= v {
			return encodeUint64(uint64(v))
		}
		if err := writeHeader(enc.writer, NInt, aiEightByte); err != nil {
			return err
		}
		return writeNint64Bytes(enc.writer, int64(v))
	case float32:
		if err := writeHeader(enc.writer, Float, fpnFloat32); err != nil {
			return err
		}
		return writeFloat32Bytes(enc.writer, v)
	case float64:
		if err := writeHeader(enc.writer, Float, fpnFloat64); err != nil {
			return err
		}
		return writeFloat64Bytes(enc.writer, v)
	case bool:
		return encodeBool(v)
	case nil:
		return encodeNull()
	case []byte:
		return writeByteString(v)
	case string:
		return writeTextString(v)
	case time.Time:
		if err := writeHeader(enc.writer, Tag, tagStdDateTime); err != nil {
			return err
		}
		return writeTextString(v.Format(time.RFC3339))
	case []int8:
		return writeAnyArray(toAnyArray(v))
	case []any:
		return writeAnyArray(v)
	}
	return newErrorNotSupportedNativeType(item)
}
