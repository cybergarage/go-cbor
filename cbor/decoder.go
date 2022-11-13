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
	// Inner utility functions

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

	readNumberOfBytes := func(mt majorType, ai addInfo) (uint, error) {
		if ai < aiOneByte {
			return uint(ai), nil
		}
		switch ai {
		case aiOneByte:
			v, err := readUint8Bytes(dec.reader)
			if err != nil {
				return 0, err
			}
			return uint(v), nil
		case aiTwoByte:
			v, err := readUint16Bytes(dec.reader)
			if err != nil {
				return 0, err
			}
			return uint(v), nil
		case aiFourByte:
			v, err := readUint32Bytes(dec.reader)
			if err != nil {
				return 0, err
			}
			return uint(v), nil
		case aiEightByte:
			v, err := readUint64Bytes(dec.reader)
			if err != nil {
				return 0, err
			}
			return uint(v), nil
		}
		return 0, newErrorNotSupportedAddInfo(mt, ai)
	}

	// 3. Specification of the CBOR Encoding.

	if _, err := io.ReadFull(dec.reader, dec.header); err != nil {
		return nil, err
	}

	majorType := majorType(dec.header[0] & majorTypeMask)
	addInfo := addInfo(dec.header[0] & addInfoMask)

	switch majorType {
	case Uint:
		if addInfo < aiOneByte {
			return returnDecordedUint8(uint8(addInfo)), nil
		}
		switch addInfo {
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
		return nil, newErrorNotSupportedAddInfo(Uint, addInfo)
	case NInt:
		if addInfo < aiOneByte {
			return -int8(addInfo + 1), nil
		}
		switch addInfo {
		case aiOneByte:
			return readNint8Bytes(dec.reader)
		case aiTwoByte:
			return readNint16Bytes(dec.reader)
		case aiFourByte:
			return readNint32Bytes(dec.reader)
		case aiEightByte:
			return readNint64Bytes(dec.reader)
		}
		return nil, newErrorNotSupportedAddInfo(NInt, addInfo)
	case Bytes:
		return nil, newErrorNotSupportedMajorType(majorType)
	case Text:
		n, err := readNumberOfBytes(Text, addInfo)
		if err != nil {
			return nil, err
		}
		bytes, err := readBytes(dec.reader, n)
		if err != nil {
			return nil, err
		}
		return string(bytes), nil
	case Array:
		return nil, newErrorNotSupportedMajorType(majorType)
	case Map:
		return nil, newErrorNotSupportedMajorType(majorType)
	case Tag:
		// switch addInfo {
		// case tagStdDateTime:
		// }
		return nil, newErrorNotSupportedMajorType(majorType)
	case Float:
		switch addInfo {
		case simpFalse:
			return false, nil
		case simpTrue:
			return true, nil
		case simpNull:
			return nil, nil
		case fpnFloat16:
			return nil, newErrorNotSupportedAddInfo(Float, addInfo)
		case fpnFloat32:
			return readFloat32Bytes(dec.reader)
		case fpnFloat64:
			return readFloat64Bytes(dec.reader)
		}
		return nil, newErrorNotSupportedAddInfo(Float, addInfo)
	}

	return nil, newErrorNotSupportedMajorType(majorType)
}
