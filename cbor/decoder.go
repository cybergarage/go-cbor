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
	if _, err := io.ReadFull(dec.reader, dec.header); err != nil {
		return nil, err
	}

	// 3. Specification of the CBOR Encoding.

	majorType := majorType(dec.header[0] & majorTypeMask)
	addInfo := AddInfo(dec.header[0] & addInfoMask)

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

	switch majorType {
	case Uint:
		if addInfo < uIntOneByte {
			return returnDecordedUint8(uint8(addInfo)), nil
		}
		switch addInfo {
		case uIntOneByte:
			v, err := readUint8Bytes(dec.reader)
			if err != nil {
				return 0, err
			}
			return returnDecordedUint8(v), nil
		case uIntTwoByte:
			v, err := readUint16Bytes(dec.reader)
			if err != nil {
				return 0, err
			}
			return returnDecordedUint16(v), nil
		case uIntFourByte:
			v, err := readUint32Bytes(dec.reader)
			if err != nil {
				return 0, err
			}
			return returnDecordedUint32(v), nil
		case uIntEightByte:
			v, err := readUint64Bytes(dec.reader)
			if err != nil {
				return 0, err
			}
			return returnDecordedUint64(v), nil
		}
		return nil, newErrorNotSupportedAddInfo(Uint, addInfo)
	case NInt:
		if addInfo < uIntOneByte {
			return -int8(addInfo + 1), nil
		}
		switch addInfo {
		case uIntOneByte:
			return readNint8Bytes(dec.reader)
		case uIntTwoByte:
			return readNint16Bytes(dec.reader)
		case uIntFourByte:
			return readNint32Bytes(dec.reader)
		case uIntEightByte:
			return readNint64Bytes(dec.reader)
		}
		return nil, newErrorNotSupportedAddInfo(NInt, addInfo)
	case Bytes:
		return 1, nil
	case Text:
		return 1, nil
	case Array:
		return 1, nil
	case Map:
		return 1, nil
	case Tag:
		return 1, nil
	case Float:
		switch addInfo {
		case False:
			return false, nil
		case True:
			return true, nil
		case Null:
			return nil, nil
		case Float16:
			return nil, newErrorNotSupportedAddInfo(Float, addInfo)
		case Float32:
			return readFloat32Bytes(dec.reader)
		case Float64:
			return readFloat64Bytes(dec.reader)
		}
		return nil, newErrorNotSupportedAddInfo(Float, addInfo)
	}

	return nil, io.EOF
}
