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
	"time"
)

// An Decoder reads CBOR values from an output stream.
type Decoder struct {
	*Config

	reader io.Reader
	header []byte
}

// NewDecoder returns a new decoder that reads from the specified writer.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		Config: NewConfig(),
		reader: r,
		header: make([]byte, 1),
	}
}

// Decode returns a next decoded item from the specified reader if available, otherwise returns EOF or another error.
// Decode returns io.EOF only when no more data is available at the beginning of an item.
// If the input ends in the middle of an item, Decode returns io.ErrUnexpectedEOF (RFC 8949 Appendix F).
func (dec *Decoder) Decode() (any, error) {
	// 3. Specification of the CBOR Encoding.
	if _, err := io.ReadFull(dec.reader, dec.header); err != nil {
		return nil, err
	}
	item, err := dec.decodeItem(dec.header[0])
	if err == io.EOF { //nolint:errorlint // io.EOF is returned unwrapped by io.Reader
		return nil, io.ErrUnexpectedEOF
	}
	return item, err
}

func (dec *Decoder) decodeItem(header byte) (any, error) { //nolint:gocyclo,maintidx,exhaustive
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

	readNumberOfItems := func(mt majorType, ai majorInfo) (int, error) {
		var n uint64
		switch {
		case ai < aiOneByte:
			n = uint64(ai)
		case ai == aiOneByte:
			v, err := readUint8Bytes(dec.reader)
			if err != nil {
				return 0, err
			}
			n = uint64(v)
		case ai == aiTwoByte:
			v, err := readUint16Bytes(dec.reader)
			if err != nil {
				return 0, err
			}
			n = uint64(v)
		case ai == aiFourByte:
			v, err := readUint32Bytes(dec.reader)
			if err != nil {
				return 0, err
			}
			n = uint64(v)
		case ai == aiEightByte:
			v, err := readUint64Bytes(dec.reader)
			if err != nil {
				return 0, err
			}
			n = v
		default:
			return 0, newErrorNotSupportedAddInfo(mt, ai)
		}
		// Reject lengths that cannot be represented as a Go int instead of overflowing (RFC 8949 Section 10).
		if uint64(math.MaxInt) < n {
			return 0, newErrorDecodeLengthTooLarge(n)
		}
		return int(n), nil
	}

	readByteString := func(m majorType, i majorInfo) ([]byte, error) {
		n, err := readNumberOfItems(m, i)
		if err != nil {
			return nil, err
		}
		return readBytes(dec.reader, n)
	}

	readTextString := func(m majorType, i majorInfo) (string, error) {
		bytes, err := readByteString(m, i)
		if err != nil {
			return "", err
		}
		return string(bytes), nil
	}

	majorType := majorType(header & majorTypeMask)
	majorInfo := majorInfo(header & majorInfoMask)

	switch majorType {
	case mtUint:
		if majorInfo < aiOneByte {
			return returnDecordedUint8(uint8(majorInfo)), nil
		}
		switch majorInfo {
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
		return nil, newErrorNotSupportedAddInfo(mtUint, majorInfo)
	case mtNInt:
		if majorInfo < aiOneByte {
			return -int8(majorInfo + 1), nil
		}
		switch majorInfo {
		case aiOneByte:
			return readNint8Bytes(dec.reader)
		case aiTwoByte:
			return readNint16Bytes(dec.reader)
		case aiFourByte:
			return readNint32Bytes(dec.reader)
		case aiEightByte:
			return readNint64Bytes(dec.reader)
		}
		return nil, newErrorNotSupportedAddInfo(mtNInt, majorInfo)
	case mtBytes:
		return readByteString(mtBytes, majorInfo)
	case mtText:
		return readTextString(mtText, majorInfo)
	case mtArray:
		itemCount, err := readNumberOfItems(mtArray, majorInfo)
		if err != nil {
			return nil, err
		}
		itemArray := make([]any, 0)
		for range itemCount {
			item, err := dec.Decode()
			if err != nil {
				return nil, err
			}
			itemArray = append(itemArray, item)
		}
		return itemArray, nil
	case mtMap:
		itemCount, err := readNumberOfItems(mtMap, majorInfo)
		if err != nil {
			return nil, err
		}
		itemMap := map[any]any{}
		for range itemCount {
			key, err := dec.Decode()
			if err != nil {
				return nil, err
			}
			// RFC 8949 allows any data item as a map key, but Go maps panic on
			// non-comparable keys such as slices and maps.
			if key != nil && !reflect.TypeOf(key).Comparable() {
				return nil, newErrorDecodeUnhashableKey(key)
			}
			val, err := dec.Decode()
			if err != nil {
				return nil, err
			}
			itemMap[key] = val
		}
		return itemMap, nil
	case mtTag:
		switch majorInfo {
		case tagStdDateTime:
			dateTime, err := dec.Decode()
			if err != nil {
				return nil, err
			}
			dateTimeStr, ok := dateTime.(string)
			if !ok {
				return nil, newErrorNotSupportedAddInfo(mtTag, majorInfo)
			}
			return time.Parse(time.RFC3339, dateTimeStr)
		case tagEpochDateTime:
		}
		return nil, newErrorNotSupportedMajorType(majorType)
	case mtFloat:
		switch majorInfo {
		case simpFalse:
			return false, nil
		case simpTrue:
			return true, nil
		case simpNull:
			return nil, nil
		case fpnFloat16:
			return readFloat16Bytes(dec.reader)
		case fpnFloat32:
			return readFloat32Bytes(dec.reader)
		case fpnFloat64:
			return readFloat64Bytes(dec.reader)
		}
		return nil, newErrorNotSupportedAddInfo(mtFloat, majorInfo)
	}

	return nil, newErrorNotSupportedMajorType(majorType)
}
