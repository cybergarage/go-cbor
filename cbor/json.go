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
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"math"
	"math/big"
	"strconv"
	"strings"
)

// jsonSubstituteValue is the substitute value for CBOR values that cannot be represented in JSON
// (RFC 8949 Section 6.1).
const jsonSubstituteValue = "null"

// noEncodingHint means that byte strings are not enclosed in an expected later encoding tag.
const noEncodingHint uint64 = 0

// ToJSON converts the specified CBOR data item to JSON as suggested in RFC 8949 Section 6.1.
// ToJSON returns ErrDecode if extraneous data follows the data item.
func ToJSON(cborBytes []byte) ([]byte, error) {
	reader := bytes.NewReader(cborBytes)
	jsonBytes, err := NewDecoder(reader).DecodeJSON()
	if err != nil {
		return nil, err
	}
	if 0 < reader.Len() {
		return nil, newErrorDecodeExtraneousData(reader.Len())
	}
	return jsonBytes, nil
}

// FromJSON converts the specified JSON text to CBOR as suggested in RFC 8949 Section 6.2,
// using the preferred serialization.
func FromJSON(jsonBytes []byte) ([]byte, error) {
	var buf bytes.Buffer
	if err := NewEncoder(&buf).EncodeJSON(jsonBytes); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// DecodeJSON reads a next data item and converts it to JSON as suggested in RFC 8949 Section 6.1:
//   - Integers become JSON numbers with all their digits.
//   - Byte strings become base64url strings without padding, or base64 or base16 (uppercase) strings
//     when they are enclosed in tag 22 or 23 (Section 3.4.5.2).
//   - Text strings, arrays, false, true, and null become the corresponding JSON values, and
//     indefinite-length items are made definite.
//   - Maps become JSON objects in the order of the encoded keys. Keys that are numbers, false, true,
//     or null are converted to strings of their JSON representations, such as "1". Other keys, and
//     keys that collide after the conversion, return ErrJSON.
//   - Finite floating-point values become JSON numbers with a fractional part or an exponent, such
//     as 1.0, so that they are converted back to floating-point values. NaN, infinities, and the
//     other simple values, such as undefined, become the substitute value null.
//   - Bignums (tags 2 and 3) become base64url strings of their byte strings, with "~" for tag 3.
//   - For the other tags, the tag content is converted and the tag number is ignored.
//
// The decoder configuration, such as MaxNestedLevels and DecodeMode, applies to the data item.
func (dec *Decoder) DecodeJSON() ([]byte, error) {
	if _, err := io.ReadFull(dec.reader, dec.header); err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err := dec.convertToJSON(&buf, dec.header[0], noEncodingHint)
	if err == io.EOF { //nolint:errorlint // io.EOF is returned unwrapped by io.Reader
		return nil, io.ErrUnexpectedEOF
	}
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// readJSONItem reads a next data item and converts it to JSON. It returns true if the "break" stop code is read.
func (dec *Decoder) readJSONItem(buf *bytes.Buffer, hint uint64) (bool, error) {
	header, err := readUint8Bytes(dec.reader)
	if err != nil {
		return false, err
	}
	if header == breakCode {
		return true, nil
	}
	return false, dec.convertToJSON(buf, header, hint)
}

// readJSONElement reads a next data item that must not be the "break" stop code and converts it to JSON.
func (dec *Decoder) readJSONElement(buf *bytes.Buffer, hint uint64) error {
	isBreak, err := dec.readJSONItem(buf, hint)
	if err != nil {
		return err
	}
	if isBreak {
		return newErrorDecodeUnexpectedBreak()
	}
	return nil
}

//nolint:gocyclo,exhaustive
func (dec *Decoder) convertToJSON(buf *bytes.Buffer, header byte, hint uint64) error {
	if header == breakCode {
		return newErrorDecodeUnexpectedBreak()
	}
	mt := majorType(header & majorTypeMask)
	ai := majorInfo(header & majorInfoMask)
	switch mt {
	case mtUint:
		v, err := dec.readArgument(mt, ai)
		if err != nil {
			return err
		}
		buf.WriteString(strconv.FormatUint(v, 10))
		return nil
	case mtNInt:
		v, err := dec.readArgument(mt, ai)
		if err != nil {
			return err
		}
		// The value is -1 - v, which is -2^64 for the largest argument.
		n := new(big.Int).SetUint64(v)
		buf.WriteString(n.Neg(n.Add(n, big.NewInt(1))).String())
		return nil
	case mtBytes:
		b, err := dec.readString(mt, ai)
		if err != nil {
			return err
		}
		writeJSONString(buf, encodeBytesForJSON(b, hint))
		return nil
	case mtText:
		str, err := dec.readString(mt, ai)
		if err != nil {
			return err
		}
		writeJSONString(buf, string(str))
		return nil
	case mtArray:
		return dec.convertArrayToJSON(buf, ai, hint)
	case mtMap:
		return dec.convertMapToJSON(buf, ai, hint)
	case mtTag:
		return dec.convertTagToJSON(buf, ai, hint)
	case mtFloat:
		return dec.convertSimpleToJSON(buf, ai)
	}
	return newErrorNotSupportedMajorType(mt)
}

func (dec *Decoder) convertArrayToJSON(buf *bytes.Buffer, ai majorInfo, hint uint64) error {
	defer dec.leaveNested()
	if err := dec.enterNested(); err != nil {
		return err
	}
	buf.WriteByte('[')
	if ai == aiIndefinite {
		if dec.isDeterministic() {
			return newErrorDecodeNotDeterministic("indefinite-length item")
		}
		for n := 0; ; n++ {
			var elem bytes.Buffer
			isBreak, err := dec.readJSONItem(&elem, hint)
			if err != nil {
				return err
			}
			if isBreak {
				break
			}
			if 0 < n {
				buf.WriteByte(',')
			}
			buf.Write(elem.Bytes())
		}
		buf.WriteByte(']')
		return nil
	}
	count, err := dec.readLength(mtArray, ai)
	if err != nil {
		return err
	}
	for n := range count {
		if 0 < n {
			buf.WriteByte(',')
		}
		if err := dec.readJSONElement(buf, hint); err != nil {
			return err
		}
	}
	buf.WriteByte(']')
	return nil
}

func (dec *Decoder) convertMapToJSON(buf *bytes.Buffer, ai majorInfo, hint uint64) error {
	defer dec.leaveNested()
	if err := dec.enterNested(); err != nil {
		return err
	}
	indefinite := ai == aiIndefinite
	count := 0
	if indefinite {
		if dec.isDeterministic() {
			return newErrorDecodeNotDeterministic("indefinite-length item")
		}
	} else {
		n, err := dec.readLength(mtMap, ai)
		if err != nil {
			return err
		}
		count = n
	}

	keys := map[string]bool{}
	var prevKey []byte
	buf.WriteByte('{')
	for n := 0; indefinite || n < count; n++ {
		var key bytes.Buffer
		recorder := &recordingReader{reader: dec.reader, buf: nil}
		dec.reader = recorder
		isBreak, err := dec.readJSONItem(&key, hint)
		dec.reader = recorder.reader
		if err != nil {
			return err
		}
		if isBreak {
			if !indefinite {
				return newErrorDecodeUnexpectedBreak()
			}
			break
		}
		// 4.2.1. The keys must be sorted in the bytewise lexicographic order of their encodings.
		if dec.isDeterministic() {
			if prevKey != nil && 0 <= bytes.Compare(prevKey, recorder.buf) {
				return newErrorDecodeNotDeterministic("unsorted or duplicate map key")
			}
			prevKey = recorder.buf
		}
		keyStr, err := jsonObjectKey(key.Bytes())
		if err != nil {
			return err
		}
		if keys[keyStr] {
			return newErrorJSONKeyCollision(keyStr)
		}
		keys[keyStr] = true
		if 0 < n {
			buf.WriteByte(',')
		}
		writeJSONString(buf, keyStr)
		buf.WriteByte(':')
		if err := dec.readJSONElement(buf, hint); err != nil {
			return err
		}
	}
	buf.WriteByte('}')
	return nil
}

// jsonObjectKey returns the JSON object key for the specified JSON value of a map key.
func jsonObjectKey(value []byte) (string, error) {
	switch {
	case len(value) == 0:
		return "", newErrorJSONUnsupportedKey(string(value))
	case value[0] == '"':
		var str string
		if err := json.Unmarshal(value, &str); err != nil {
			return "", newErrorJSON(err)
		}
		return str, nil
	case value[0] == '[' || value[0] == '{':
		// 6.1. Maps can be converted directly only if all keys are text strings.
		return "", newErrorJSONUnsupportedKey(string(value))
	}
	// Numbers, false, true, and null are converted to strings of their JSON representations.
	return string(value), nil
}

func (dec *Decoder) convertTagToJSON(buf *bytes.Buffer, ai majorInfo, hint uint64) error {
	defer dec.leaveNested()
	if err := dec.enterNested(); err != nil {
		return err
	}
	tagNumber, err := dec.readArgument(mtTag, ai)
	if err != nil {
		return err
	}
	switch tagNumber {
	case tagPositiveBignum, tagNegativeBignum:
		// 6.1. A bignum is represented by encoding its byte string in base64url without padding,
		// and "~" is inserted before the base-encoded value of a negative bignum.
		header, err := readUint8Bytes(dec.reader)
		if err != nil {
			return err
		}
		if majorType(header&majorTypeMask) != mtBytes {
			return newErrorDecodeInvalidTagContent(tagNumber, header)
		}
		b, err := dec.readString(mtBytes, majorInfo(header&majorInfoMask))
		if err != nil {
			return err
		}
		if dec.isDeterministic() && (len(b) <= 8 || b[0] == 0) {
			return newErrorDecodeNotDeterministic("non-preferred bignum")
		}
		prefix := ""
		if tagNumber == tagNegativeBignum {
			prefix = "~"
		}
		writeJSONString(buf, prefix+base64.RawURLEncoding.EncodeToString(b))
		return nil
	case tagExpectedBase64URL, tagExpectedBase64, tagExpectedBase16:
		// 3.4.5.2. The hint applies to all byte strings in the tag content, except for those
		// in a nested data item tagged with another hint.
		return dec.readJSONElement(buf, tagNumber)
	}
	// 6.1. For all other tags, the tag content is represented as a JSON value; the tag number is ignored.
	return dec.readJSONElement(buf, hint)
}

func (dec *Decoder) convertSimpleToJSON(buf *bytes.Buffer, ai majorInfo) error {
	writeFloat := func(v float64) {
		// 6.1. Non-finite values are represented by the substitute value.
		if math.IsNaN(v) || math.IsInf(v, 0) {
			buf.WriteString(jsonSubstituteValue)
			return
		}
		buf.WriteString(formatJSONNumber(v))
	}
	switch ai {
	case simpFalse:
		buf.WriteString("false")
	case simpTrue:
		buf.WriteString("true")
	case simpNull:
		buf.WriteString("null")
	case simpOneByte:
		v, err := readUint8Bytes(dec.reader)
		if err != nil {
			return err
		}
		if v < simpMinOneByte {
			return newErrorDecodeInvalidSimpleValue(v)
		}
		buf.WriteString(jsonSubstituteValue)
	case fpnFloat16:
		v, err := readFloat16Bytes(dec.reader)
		if err != nil {
			return err
		}
		writeFloat(v)
	case fpnFloat32:
		v, err := readFloat32Bytes(dec.reader)
		if err != nil {
			return err
		}
		if dec.isDeterministic() && !isShortestFloat32(v) {
			return newErrorDecodeNotDeterministic("non-shortest floating-point value")
		}
		writeFloat(float64(v))
	case fpnFloat64:
		v, err := readFloat64Bytes(dec.reader)
		if err != nil {
			return err
		}
		if dec.isDeterministic() && !isShortestFloat64(v) {
			return newErrorDecodeNotDeterministic("non-shortest floating-point value")
		}
		writeFloat(v)
	default:
		if ai < aiOneByte {
			// 6.1. Any other simple value, such as undefined, is represented by the substitute value.
			buf.WriteString(jsonSubstituteValue)
			return nil
		}
		return newErrorNotSupportedAddInfo(mtFloat, ai)
	}
	return nil
}

// formatJSONNumber formats a finite floating-point value as a JSON number like encoding/json, except
// that integral values have a fractional part, such as 1.0. The shortest representation of the
// binary64 value is used for all precisions, so that JSON decoders that read numbers as binary64
// (RFC 8949 Section 6.2) get the exact value.
func formatJSONNumber(v float64) string {
	abs := math.Abs(v)
	format := byte('f')
	if abs != 0 && (abs < 1e-6 || 1e21 <= abs) {
		format = 'e'
	}
	str := strconv.FormatFloat(v, format, -1, 64)
	// Keep a fractional part so that the value is converted back to a floating-point value,
	// not an integer (RFC 8949 Section 6.2), and -0.0 keeps its sign.
	if format == 'f' && !strings.Contains(str, ".") {
		str += ".0"
	}
	if format == 'e' {
		// Clean up e-09 to e-9 like encoding/json.
		if n := len(str); 4 <= n && str[n-4] == 'e' && str[n-3] == '-' && str[n-2] == '0' {
			str = str[:n-2] + str[n-1:]
		}
	}
	return str
}

// encodeBytesForJSON encodes a byte string as suggested by the expected later encoding (RFC 8949 Section 3.4.5.2).
func encodeBytesForJSON(b []byte, hint uint64) string {
	switch hint {
	case tagExpectedBase64:
		return base64.StdEncoding.EncodeToString(b)
	case tagExpectedBase16:
		return strings.ToUpper(hex.EncodeToString(b))
	}
	return base64.RawURLEncoding.EncodeToString(b)
}

// writeJSONString writes a JSON string, escaping only the quotation mark, the reverse solidus,
// and the C0 control characters (RFC 8949 Section 6.1).
func writeJSONString(buf *bytes.Buffer, str string) {
	const hexDigits = "0123456789abcdef"
	buf.WriteByte('"')
	for n := range len(str) {
		c := str[n]
		switch c {
		case '"':
			buf.WriteString(`\"`)
		case '\\':
			buf.WriteString(`\\`)
		case '\n':
			buf.WriteString(`\n`)
		case '\r':
			buf.WriteString(`\r`)
		case '\t':
			buf.WriteString(`\t`)
		case '\b':
			buf.WriteString(`\b`)
		case '\f':
			buf.WriteString(`\f`)
		default:
			if c < 0x20 {
				buf.WriteString(`\u00`)
				buf.WriteByte(hexDigits[c>>4])
				buf.WriteByte(hexDigits[c&0xF])
				continue
			}
			buf.WriteByte(c)
		}
	}
	buf.WriteByte('"')
}

////////////////////////////////////////////////////////////
// JSON to CBOR
////////////////////////////////////////////////////////////

// jsonObject represents a JSON object whose members are kept in order.
type jsonObject struct {
	keys   []string
	values []any
}

// EncodeJSON converts the specified JSON text to CBOR as suggested in RFC 8949 Section 6.2 and writes it:
//   - Numbers without fractional parts or exponents become integers in the shortest form, or bignums
//     (tags 2 and 3) when they do not fit in 64 bits.
//   - The other numbers are converted to binary64 with roundTiesToEven and are encoded as
//     floating-point values; the preferred serialization uses the shortest form that preserves them.
//   - Strings, arrays, false, true, and null become the corresponding CBOR data items, and
//     objects become maps in the order of the JSON members.
//
// The encoder configuration, such as EncodeMode and MapSortEnabled, applies to the output.
// EncodeJSON returns ErrJSON for invalid JSON, duplicate object member names, numbers that overflow
// binary64, and nesting deeper than MaxNestedLevels.
func (enc *Encoder) EncodeJSON(jsonBytes []byte) error {
	dec := json.NewDecoder(bytes.NewReader(jsonBytes))
	dec.UseNumber()
	v, err := enc.parseJSONValue(dec, 0)
	if err != nil {
		return err
	}
	if _, err := dec.Token(); !errors.Is(err, io.EOF) {
		return newErrorJSON(errors.New("extraneous data after the JSON value"))
	}
	if err := enc.beforeIndefiniteElement(v); err != nil {
		return err
	}
	return enc.encodeJSONValue(v)
}

func (enc *Encoder) parseJSONValue(dec *json.Decoder, depth int) (any, error) {
	token, err := dec.Token()
	if err != nil {
		if errors.Is(err, io.EOF) {
			err = io.ErrUnexpectedEOF
		}
		return nil, newErrorJSON(err)
	}
	switch v := token.(type) {
	case json.Delim:
		if enc.MaxNestedLevels() <= depth {
			return nil, newErrorJSON(newErrorDecodeTooDeep(enc.MaxNestedLevels()))
		}
		if v == '[' {
			array := make([]any, 0)
			for dec.More() {
				elem, err := enc.parseJSONValue(dec, depth+1)
				if err != nil {
					return nil, err
				}
				array = append(array, elem)
			}
			if _, err := dec.Token(); err != nil {
				return nil, newErrorJSON(err)
			}
			return array, nil
		}
		object := &jsonObject{keys: nil, values: nil}
		names := map[string]bool{}
		for dec.More() {
			keyToken, err := dec.Token()
			if err != nil {
				return nil, newErrorJSON(err)
			}
			key, _ := keyToken.(string)
			// 5.6. Maps with duplicate keys are not valid.
			if names[key] {
				return nil, newErrorJSONKeyCollision(key)
			}
			names[key] = true
			value, err := enc.parseJSONValue(dec, depth+1)
			if err != nil {
				return nil, err
			}
			object.keys = append(object.keys, key)
			object.values = append(object.values, value)
		}
		if _, err := dec.Token(); err != nil {
			return nil, newErrorJSON(err)
		}
		return object, nil
	case json.Number:
		return parseJSONNumber(v)
	}
	// string, bool, and nil
	return token, nil
}

// parseJSONNumber converts a JSON number to an integer (*big.Int) or a floating-point value (float64).
func parseJSONNumber(n json.Number) (any, error) {
	str := n.String()
	if !strings.ContainsAny(str, ".eE") {
		v, ok := new(big.Int).SetString(str, 10)
		if !ok {
			return nil, newErrorJSON(errors.New("invalid number: " + str))
		}
		return v, nil
	}
	v, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return nil, newErrorJSON(err)
	}
	return v, nil
}

func (enc *Encoder) encodeJSONValue(v any) error {
	switch v := v.(type) {
	case *big.Int:
		return enc.encodeBigInt(v)
	case float64:
		if enc.EncodeMode() == EncodeModeTypePreserving {
			return enc.encode(v)
		}
		return enc.encodePreferredFloat(v)
	case []any:
		if err := enc.encodeArgument(mtArray, uint64(len(v))); err != nil {
			return err
		}
		for _, elem := range v {
			if err := enc.encodeJSONValue(elem); err != nil {
				return err
			}
		}
		return nil
	case *jsonObject:
		return enc.encodeJSONObject(v)
	}
	// string, bool, and nil
	return enc.encode(v)
}

func (enc *Encoder) encodeJSONObject(object *jsonObject) error {
	if enc.isMapKeySortRequired() {
		// Encode the members to sort them by the encodings of the keys.
		m := make(map[any]any, len(object.keys))
		for n, key := range object.keys {
			var buf bytes.Buffer
			sub := &Encoder{Config: enc.Config, writer: &buf, indefinite: nil}
			if err := sub.encodeJSONValue(object.values[n]); err != nil {
				return err
			}
			m[key] = rawCBOR(buf.Bytes())
		}
		return enc.encodeAnyMap(m)
	}
	if err := enc.encodeArgument(mtMap, uint64(len(object.keys))); err != nil {
		return err
	}
	for n, key := range object.keys {
		if err := enc.encodeTextString(key); err != nil {
			return err
		}
		if err := enc.encodeJSONValue(object.values[n]); err != nil {
			return err
		}
	}
	return nil
}

// rawCBOR is an already encoded data item that is written as is.
type rawCBOR []byte
