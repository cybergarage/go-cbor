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

// breakStopCode is returned internally by readItem when the "break" stop code (0xFF) is read.
type breakStopCode struct{}

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
	item, err := dec.readItem()
	if err != nil {
		return nil, err
	}
	if _, ok := item.(breakStopCode); ok {
		return nil, newErrorDecodeUnexpectedBreak()
	}
	return item, nil
}

// readItem reads a next data item, or returns breakStopCode if the "break" stop code is read.
func (dec *Decoder) readItem() (any, error) {
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

// readArgument reads the argument of the data item following the initial byte (RFC 8949 Section 3).
func (dec *Decoder) readArgument(mt majorType, ai majorInfo) (uint64, error) {
	switch {
	case ai < aiOneByte:
		return uint64(ai), nil
	case ai == aiOneByte:
		v, err := readUint8Bytes(dec.reader)
		return uint64(v), err
	case ai == aiTwoByte:
		v, err := readUint16Bytes(dec.reader)
		return uint64(v), err
	case ai == aiFourByte:
		v, err := readUint32Bytes(dec.reader)
		return uint64(v), err
	case ai == aiEightByte:
		return readUint64Bytes(dec.reader)
	}
	return 0, newErrorNotSupportedAddInfo(mt, ai)
}

// readLength reads the length (or the number of items) of a definite-length data item.
func (dec *Decoder) readLength(mt majorType, ai majorInfo) (int, error) {
	n, err := dec.readArgument(mt, ai)
	if err != nil {
		return 0, err
	}
	// Reject lengths that cannot be represented as a Go int instead of overflowing (RFC 8949 Section 10).
	if uint64(math.MaxInt) < n {
		return 0, newErrorDecodeLengthTooLarge(n)
	}
	return int(n), nil
}

// readString reads a byte or text string, which may be an indefinite-length string (RFC 8949 Section 3.2.3).
func (dec *Decoder) readString(mt majorType, ai majorInfo) ([]byte, error) {
	if ai != aiIndefinite {
		n, err := dec.readLength(mt, ai)
		if err != nil {
			return nil, err
		}
		return readBytes(dec.reader, n)
	}
	// An indefinite-length string is a sequence of definite-length strings of the same major type,
	// called chunks, terminated by the "break" stop code.
	str := make([]byte, 0)
	for {
		header, err := readUint8Bytes(dec.reader)
		if err != nil {
			return nil, err
		}
		if header == breakCode {
			return str, nil
		}
		chunkType := majorType(header & majorTypeMask)
		chunkInfo := majorInfo(header & majorInfoMask)
		if chunkType != mt || chunkInfo == aiIndefinite {
			return nil, newErrorDecodeInvalidChunk(mt, header)
		}
		n, err := dec.readLength(chunkType, chunkInfo)
		if err != nil {
			return nil, err
		}
		chunk, err := readBytes(dec.reader, n)
		if err != nil {
			return nil, err
		}
		str = append(str, chunk...)
	}
}

// readArray reads an array, which may be an indefinite-length array (RFC 8949 Section 3.2.2).
func (dec *Decoder) readArray(ai majorInfo) ([]any, error) {
	itemArray := make([]any, 0)
	if ai == aiIndefinite {
		for {
			item, err := dec.readItem()
			if err != nil {
				return nil, err
			}
			if _, ok := item.(breakStopCode); ok {
				return itemArray, nil
			}
			itemArray = append(itemArray, item)
		}
	}
	itemCount, err := dec.readLength(mtArray, ai)
	if err != nil {
		return nil, err
	}
	for range itemCount {
		item, err := dec.Decode()
		if err != nil {
			return nil, err
		}
		itemArray = append(itemArray, item)
	}
	return itemArray, nil
}

// readMap reads a map, which may be an indefinite-length map (RFC 8949 Section 3.2.2).
func (dec *Decoder) readMap(ai majorInfo) (map[any]any, error) {
	itemMap := map[any]any{}
	readPair := func(key any) error {
		// RFC 8949 allows any data item as a map key, but Go maps panic on
		// non-comparable keys such as slices and maps.
		if !isHashableKey(key) {
			return newErrorDecodeUnhashableKey(key)
		}
		val, err := dec.Decode()
		if err != nil {
			return err
		}
		itemMap[key] = val
		return nil
	}
	if ai == aiIndefinite {
		for {
			key, err := dec.readItem()
			if err != nil {
				return nil, err
			}
			if _, ok := key.(breakStopCode); ok {
				return itemMap, nil
			}
			if err := readPair(key); err != nil {
				return nil, err
			}
		}
	}
	itemCount, err := dec.readLength(mtMap, ai)
	if err != nil {
		return nil, err
	}
	for range itemCount {
		key, err := dec.Decode()
		if err != nil {
			return nil, err
		}
		if err := readPair(key); err != nil {
			return nil, err
		}
	}
	return itemMap, nil
}

// readTag reads a tagged data item (RFC 8949 Section 3.4).
func (dec *Decoder) readTag(ai majorInfo) (any, error) {
	tagNumber, err := dec.readArgument(mtTag, ai)
	if err != nil {
		return nil, err
	}
	content, err := dec.Decode()
	if err != nil {
		return nil, err
	}
	switch tagNumber {
	case tagStdDateTime:
		dateTimeStr, ok := content.(string)
		if !ok {
			return nil, newErrorNotSupportedAddInfo(mtTag, majorInfo(tagNumber))
		}
		return time.Parse(time.RFC3339, dateTimeStr)
	case tagSelfDescribed:
		// 3.4.6. Self-Described CBOR: the tag only marks the data as CBOR and does not change its meaning.
		return content, nil
	}
	// Tags that are not interpreted by the decoder are returned with their content as is,
	// so that generic decoders can pass them through (RFC 8949 Section 3.4).
	return Tag{Number: tagNumber, Content: content}, nil
}

func (dec *Decoder) decodeItem(header byte) (any, error) { //nolint:gocyclo,exhaustive
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

	if header == breakCode {
		return breakStopCode{}, nil
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
		return dec.readString(mtBytes, majorInfo)
	case mtText:
		str, err := dec.readString(mtText, majorInfo)
		if err != nil {
			return nil, err
		}
		return string(str), nil
	case mtArray:
		return dec.readArray(majorInfo)
	case mtMap:
		return dec.readMap(majorInfo)
	case mtTag:
		return dec.readTag(majorInfo)
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

// isHashableKey returns true if the specified decoded item can be used as a Go map key.
func isHashableKey(key any) bool {
	switch k := key.(type) {
	case nil:
		return true
	case Tag:
		return isHashableKey(k.Content)
	}
	return reflect.TypeOf(key).Comparable()
}
