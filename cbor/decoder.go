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
	"bytes"
	"io"
	"math"
	"math/big"
	"reflect"
	"time"
	"unicode/utf8"
)

// breakStopCode is returned internally by readItem when the "break" stop code (0xFF) is read.
type breakStopCode struct{}

// An Decoder reads CBOR values from an output stream.
type Decoder struct {
	*Config

	reader io.Reader
	header []byte
	depth  int
}

// recordingReader records the bytes read through it, to compare the encodings of map keys.
type recordingReader struct {
	reader io.Reader
	buf    []byte
}

func (r *recordingReader) Read(p []byte) (int, error) {
	n, err := r.reader.Read(p)
	r.buf = append(r.buf, p[:n]...)
	return n, err
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

// isDeterministic returns true if only the core deterministic encoding is accepted.
func (dec *Decoder) isDeterministic() bool {
	return dec.DecodeMode() == DecodeModeCoreDeterministic
}

// enterNested increments the nesting depth and returns an error if it exceeds the maximum (RFC 8949 Section 10).
// The caller must call leaveNested whether or not an error is returned.
func (dec *Decoder) enterNested() error {
	dec.depth++
	if maxLevels := dec.MaxNestedLevels(); maxLevels < dec.depth {
		return newErrorDecodeTooDeep(maxLevels)
	}
	return nil
}

func (dec *Decoder) leaveNested() {
	dec.depth--
}

// readArgument reads the argument of the data item following the initial byte (RFC 8949 Section 3).
func (dec *Decoder) readArgument(mt majorType, ai majorInfo) (uint64, error) {
	var v uint64
	var err error
	var shortest bool
	switch {
	case ai < aiOneByte:
		return uint64(ai), nil
	case ai == aiOneByte:
		var v8 uint8
		v8, err = readUint8Bytes(dec.reader)
		v = uint64(v8)
		shortest = uint64(aiOneByte) <= v
	case ai == aiTwoByte:
		var v16 uint16
		v16, err = readUint16Bytes(dec.reader)
		v = uint64(v16)
		shortest = math.MaxUint8 < v
	case ai == aiFourByte:
		var v32 uint32
		v32, err = readUint32Bytes(dec.reader)
		v = uint64(v32)
		shortest = math.MaxUint16 < v
	case ai == aiEightByte:
		v, err = readUint64Bytes(dec.reader)
		shortest = math.MaxUint32 < v
	default:
		return 0, newErrorNotSupportedAddInfo(mt, ai)
	}
	if err != nil {
		return 0, err
	}
	// 4.2.1. Arguments must be as short as possible.
	if !shortest && dec.isDeterministic() {
		return 0, newErrorDecodeNotDeterministic("non-shortest argument")
	}
	return v, nil
}

// argumentSize returns the size in bytes of the argument for the specified additional information,
// treating the arguments in the initial byte as one byte.
func argumentSize(ai majorInfo) int {
	switch ai {
	case aiTwoByte:
		return 2
	case aiFourByte:
		return 4
	case aiEightByte:
		return 8
	}
	return 1
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
	// 5.3.1. A text string must be a valid UTF-8 string, and so must be each chunk of
	// an indefinite-length text string (3.2.3).
	validateUTF8 := mt == mtText && dec.IsUTF8ValidationEnabled()
	if ai != aiIndefinite {
		n, err := dec.readLength(mt, ai)
		if err != nil {
			return nil, err
		}
		str, err := readBytes(dec.reader, n)
		if err != nil {
			return nil, err
		}
		if validateUTF8 && !utf8.Valid(str) {
			return nil, newErrorDecodeInvalidUTF8()
		}
		return str, nil
	}
	if dec.isDeterministic() {
		return nil, newErrorDecodeNotDeterministic("indefinite-length item")
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
		if validateUTF8 && !utf8.Valid(chunk) {
			return nil, newErrorDecodeInvalidUTF8()
		}
		str = append(str, chunk...)
	}
}

// readArray reads an array, which may be an indefinite-length array (RFC 8949 Section 3.2.2).
func (dec *Decoder) readArray(ai majorInfo) ([]any, error) {
	defer dec.leaveNested()
	if err := dec.enterNested(); err != nil {
		return nil, err
	}
	itemArray := make([]any, 0)
	if ai == aiIndefinite {
		if dec.isDeterministic() {
			return nil, newErrorDecodeNotDeterministic("indefinite-length item")
		}
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
	defer dec.leaveNested()
	if err := dec.enterNested(); err != nil {
		return nil, err
	}
	rejectDuplicates := dec.isDeterministic() || dec.DuplicateMapKeyMode() == DuplicateMapKeyRejected
	itemMap := map[any]any{}
	readPair := func(key any) error {
		// RFC 8949 allows any data item as a map key, but Go maps panic on
		// non-comparable keys such as slices and maps.
		if !isHashableKey(key) {
			return newErrorDecodeUnhashableKey(key)
		}
		// 5.6. Maps with duplicate keys are not valid.
		if _, ok := itemMap[key]; ok && rejectDuplicates {
			return newErrorDecodeDuplicateMapKey(key)
		}
		val, err := dec.Decode()
		if err != nil {
			return err
		}
		itemMap[key] = val
		return nil
	}
	if ai == aiIndefinite {
		if dec.isDeterministic() {
			return nil, newErrorDecodeNotDeterministic("indefinite-length item")
		}
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
	var prevKey []byte
	for range itemCount {
		if !dec.isDeterministic() {
			key, err := dec.Decode()
			if err != nil {
				return nil, err
			}
			if err := readPair(key); err != nil {
				return nil, err
			}
			continue
		}
		// 4.2.1. The keys must be sorted in the bytewise lexicographic order of their encodings,
		// which also rejects duplicate keys.
		recorder := &recordingReader{reader: dec.reader, buf: nil}
		dec.reader = recorder
		key, err := dec.Decode()
		dec.reader = recorder.reader
		if err != nil {
			return nil, err
		}
		if prevKey != nil && 0 <= bytes.Compare(prevKey, recorder.buf) {
			return nil, newErrorDecodeNotDeterministic("unsorted or duplicate map key")
		}
		prevKey = recorder.buf
		if err := readPair(key); err != nil {
			return nil, err
		}
	}
	return itemMap, nil
}

// readTag reads a tagged data item (RFC 8949 Section 3.4).
func (dec *Decoder) readTag(ai majorInfo) (any, error) {
	defer dec.leaveNested()
	if err := dec.enterNested(); err != nil {
		return nil, err
	}
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
		// 3.4.1. Standard Date/Time String
		dateTimeStr, ok := content.(string)
		if !ok {
			return nil, newErrorDecodeInvalidTagContent(tagNumber, content)
		}
		return time.Parse(time.RFC3339, dateTimeStr)
	case tagEpochDateTime:
		return epochToTime(content)
	case tagPositiveBignum, tagNegativeBignum:
		// 3.4.3. Bignums
		b, ok := content.([]byte)
		if !ok {
			return nil, newErrorDecodeInvalidTagContent(tagNumber, content)
		}
		// The preferred serialization of a bignum has no leading zeroes, and values that fit
		// in 64 bits are encoded as major type 0 or 1 instead.
		if dec.isDeterministic() && (len(b) <= 8 || b[0] == 0) {
			return nil, newErrorDecodeNotDeterministic("non-preferred bignum")
		}
		v := new(big.Int).SetBytes(b)
		if tagNumber == tagNegativeBignum {
			// The value of a negative bignum is -1 - n.
			v.Add(v, big.NewInt(1))
			v.Neg(v)
		}
		return v, nil
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
		v, err := dec.readArgument(mtUint, majorInfo)
		if err != nil {
			return nil, err
		}
		switch argumentSize(majorInfo) {
		case 1:
			return returnDecordedUint8(uint8(v)), nil
		case 2:
			return returnDecordedUint16(uint16(v)), nil
		case 4:
			return returnDecordedUint32(uint32(v)), nil
		}
		return returnDecordedUint64(v), nil
	case mtNInt:
		v, err := dec.readArgument(mtNInt, majorInfo)
		if err != nil {
			return nil, err
		}
		return nintValue(v, argumentSize(majorInfo)), nil
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
		// 3.3. Simple values 0..19 are unassigned and 23 is "undefined".
		if majorInfo < simpFalse || SimpleValue(majorInfo) == Undefined {
			return SimpleValue(majorInfo), nil
		}
		switch majorInfo {
		case simpFalse:
			return false, nil
		case simpTrue:
			return true, nil
		case simpNull:
			return nil, nil
		case simpOneByte:
			v, err := readUint8Bytes(dec.reader)
			if err != nil {
				return nil, err
			}
			// 3.3. Simple values 0..31 must use the one-byte form, so "f8 00".."f8 1f" is not well-formed.
			if v < simpMinOneByte {
				return nil, newErrorDecodeInvalidSimpleValue(v)
			}
			return SimpleValue(v), nil
		case fpnFloat16:
			return readFloat16Bytes(dec.reader)
		case fpnFloat32:
			v, err := readFloat32Bytes(dec.reader)
			if err != nil {
				return nil, err
			}
			// 4.2.1. Floating-point values must use the shortest form that preserves the value.
			if dec.isDeterministic() && !isShortestFloat32(v) {
				return nil, newErrorDecodeNotDeterministic("non-shortest floating-point value")
			}
			return v, nil
		case fpnFloat64:
			v, err := readFloat64Bytes(dec.reader)
			if err != nil {
				return nil, err
			}
			if dec.isDeterministic() && !isShortestFloat64(v) {
				return nil, newErrorDecodeNotDeterministic("non-shortest floating-point value")
			}
			return v, nil
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

// epochToTime converts the content of an epoch-based date/time (tag 1) to time.Time in UTC (RFC 8949 Section 3.4.2).
func epochToTime(content any) (time.Time, error) {
	invalidContent := func() (time.Time, error) {
		return time.Time{}, newErrorDecodeInvalidTagContent(tagEpochDateTime, content)
	}
	var secs float64
	switch v := content.(type) {
	case int8:
		return time.Unix(int64(v), 0).UTC(), nil
	case int16:
		return time.Unix(int64(v), 0).UTC(), nil
	case int32:
		return time.Unix(int64(v), 0).UTC(), nil
	case int64:
		return time.Unix(v, 0).UTC(), nil
	case uint8:
		return time.Unix(int64(v), 0).UTC(), nil
	case uint16:
		return time.Unix(int64(v), 0).UTC(), nil
	case uint32:
		return time.Unix(int64(v), 0).UTC(), nil
	case uint64:
		if math.MaxInt64 < v {
			return invalidContent()
		}
		return time.Unix(int64(v), 0).UTC(), nil
	case float32:
		secs = float64(v)
	case float64:
		secs = v
	default:
		return invalidContent()
	}
	if math.IsNaN(secs) || math.IsInf(secs, 0) {
		return invalidContent()
	}
	sec := math.Floor(secs)
	// float64(math.MaxInt64) rounds up to 2^63, so it is excluded.
	if sec < math.MinInt64 || float64(math.MaxInt64) <= sec {
		return invalidContent()
	}
	nsec := math.Round((secs - sec) * 1e9)
	return time.Unix(int64(sec), int64(nsec)).UTC(), nil
}

// isShortestFloat32 returns true if the single-precision value cannot be encoded as a half-precision value.
// NaN and infinities are encoded as half-precision values in the preferred serialization.
func isShortestFloat32(v float32) bool {
	f := float64(v)
	if math.IsNaN(f) || math.IsInf(f, 0) {
		return false
	}
	_, ok := float64ToFloat16(f)
	return !ok
}

// isShortestFloat64 returns true if the double-precision value cannot be encoded as a shorter value.
func isShortestFloat64(v float64) bool {
	if math.IsNaN(v) || math.IsInf(v, 0) {
		return false
	}
	return float64(float32(v)) != v
}
