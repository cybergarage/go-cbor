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

// Decode returns the next item if available, otherwise returns EOF or error.
func (dec *Decoder) Decode() (any, error) {
	if _, err := io.ReadFull(dec.reader, dec.header); err != nil {
		return nil, err
	}

	// 3. Specification of the CBOR Encoding.

	majorType := MajorType(dec.header[0] & majorTypeMask)
	addInfo := AddInfo(dec.header[0] & addInfoMask)

	switch majorType {
	case Uint:
		if addInfo < uIntOneByte {
			return uint8(addInfo), nil
		}
		switch addInfo {
		case uIntOneByte:
			return readUint8Bytes(dec.reader)
		case uIntTwoByte:
			return readUint16Bytes(dec.reader)
		case uIntFourByte:
			return readUint32Bytes(dec.reader)
		case uIntEightByte:
			return readUint64Bytes(dec.reader)
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
		case FloatSimple:
			return nil, newErrorNotSupportedAddInfo(Float, addInfo)
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
