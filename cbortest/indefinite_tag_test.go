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

package cbortest

import (
	"bytes"
	"encoding/hex"
	"errors"
	"io"
	"reflect"
	"testing"

	"github.com/cybergarage/go-cbor/cbor"
)

// RFC 8949 Section 3.2 and the indefinite-length examples in Appendix A.
func TestRFC8949IndefiniteLength(t *testing.T) {
	nested := []any{int8(1), []any{int8(2), int8(3)}, []any{int8(4), int8(5)}}
	oneTo25 := make([]any, 0)
	for n := 1; n <= 25; n++ {
		oneTo25 = append(oneTo25, int8(n))
	}
	tests := []struct {
		encoded  string
		expected any
	}{
		{"5fff", []byte{}},
		{"5f4100ff", []byte{0x00}},
		{"5f42010243030405ff", []byte{0x01, 0x02, 0x03, 0x04, 0x05}},
		{"5f40420102ff", []byte{0x01, 0x02}},
		{"7fff", ""},
		{"7f657374726561646d696e67ff", "streaming"},
		{"7f62c3a160ff", "á"},
		{"9fff", []any{}},
		{"9f018202039f0405ffff", nested},
		{"9f01820203820405ff", nested},
		{"83018202039f0405ff", nested},
		{"83019f0203ff820405", nested},
		{"9f0102030405060708090a0b0c0d0e0f101112131415161718181819ff", oneTo25},
		{"bfff", map[any]any{}},
		{"bf61610161629f0203ffff", map[any]any{"a": int8(1), "b": []any{int8(2), int8(3)}}},
		{"826161bf61626163ff", []any{"a", map[any]any{"b": "c"}}},
		{"bf6346756ef563416d7421ff", map[any]any{"Fun": true, "Amt": int8(-2)}},
		{"a161619fff", map[any]any{"a": []any{}}},
	}
	for _, test := range tests {
		t.Run(test.encoded, func(t *testing.T) {
			v, err := decodeHex(t, test.encoded)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(v, test.expected) {
				t.Errorf("%s: got %#v, want %#v", test.encoded, v, test.expected)
			}
		})
	}
}

// RFC 8949 Section 3.2 and Appendix F: misplaced "break" stop codes and invalid chunks are not well-formed.
func TestRFC8949IndefiniteLengthMalformed(t *testing.T) {
	notWellFormed := []string{
		"ff",           // break outside an indefinite-length item
		"81ff",         // break in a definite-length array
		"a1ff",         // break as a key of a definite-length map
		"a101ff",       // break as a value of a definite-length map
		"bf01ff",       // break instead of a map value
		"c0ff",         // break as tag content
		"5f6161ff",     // text chunk in an indefinite-length byte string
		"7f4161ff",     // byte chunk in an indefinite-length text string
		"5f00ff",       // integer chunk
		"5f5f4100ffff", // indefinite-length chunk
		"7f7f6161ffff", // indefinite-length chunk
		"bf9fff01ff",   // array key in an indefinite-length map (not a valid Go map key)
	}
	for _, test := range notWellFormed {
		t.Run(test, func(t *testing.T) {
			_, err := decodeHex(t, test)
			if !errors.Is(err, cbor.ErrDecode) {
				t.Errorf("%s: got err=%v, want %v", test, err, cbor.ErrDecode)
			}
		})
	}

	truncated := []string{"5f", "5f4101", "5f41", "7f", "7f6161", "9f", "9f01", "bf", "bf01", "bf0101", "9f9fff"}
	for _, test := range truncated {
		t.Run(test, func(t *testing.T) {
			_, err := decodeHex(t, test)
			if !errors.Is(err, io.ErrUnexpectedEOF) {
				t.Errorf("%s: got err=%v, want %v", test, err, io.ErrUnexpectedEOF)
			}
		})
	}

	// Additional information 31 is not allowed for major types 0, 1, and 6.
	for _, test := range []string{"1f", "3f", "df00"} {
		t.Run(test, func(t *testing.T) {
			if _, err := decodeHex(t, test); err == nil {
				t.Errorf("%s: expected an error", test)
			}
		})
	}
}

func TestRFC8949IndefiniteLengthStream(t *testing.T) {
	b, _ := hex.DecodeString("9f01ff7f6161ff02")
	dec := cbor.NewDecoder(bytes.NewReader(b))
	for _, want := range []any{[]any{int8(1)}, "a", int8(2)} {
		got, err := dec.Decode()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v, want %#v", got, want)
		}
	}
	if _, err := dec.Decode(); err != io.EOF { //nolint:errorlint
		t.Errorf("got err=%v, want io.EOF", err)
	}
}

// RFC 8949 Section 3.4: tags that are not interpreted by the decoder are passed through.
func TestRFC8949Tags(t *testing.T) {
	tests := []struct {
		encoded  string
		expected any
	}{
		{"c48221196ab3", cbor.Tag{Number: 4, Content: []any{int8(-2), int16(27315)}}},
		{"d74401020304", cbor.Tag{Number: 23, Content: []byte{0x01, 0x02, 0x03, 0x04}}},
		{"d818456449455446", cbor.Tag{Number: 24, Content: []byte("dIETF")}},
		{"d82076687474703a2f2f7777772e6578616d706c652e636f6d", cbor.Tag{Number: 32, Content: "http://www.example.com"}},
		{"d903e801", cbor.Tag{Number: 1000, Content: int8(1)}},
		{"da0001000001", cbor.Tag{Number: 65536, Content: int8(1)}},
		{"dbffffffffffffffff01", cbor.Tag{Number: 18446744073709551615, Content: int8(1)}},
		{"d8189f01ff", cbor.Tag{Number: 24, Content: []any{int8(1)}}},
		{"d81ad81b01", cbor.Tag{Number: 26, Content: cbor.Tag{Number: 27, Content: int8(1)}}},
		// 3.4.6. Self-Described CBOR
		{"d9d9f700", int8(0)},
		{"d9d9f783010203", []any{int8(1), int8(2), int8(3)}},
	}

	for _, test := range tests {
		t.Run(test.encoded, func(t *testing.T) {
			v, err := decodeHex(t, test.encoded)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(v, test.expected) {
				t.Errorf("%s: got %#v, want %#v", test.encoded, v, test.expected)
			}
		})
	}

	for _, test := range []string{"c0", "d8", "d818", "d903"} {
		t.Run(test, func(t *testing.T) {
			if _, err := decodeHex(t, test); !errors.Is(err, io.ErrUnexpectedEOF) {
				t.Errorf("%s: got err=%v, want %v", test, err, io.ErrUnexpectedEOF)
			}
		})
	}
}

func TestRFC8949TagMapKeys(t *testing.T) {
	v, err := decodeHex(t, "a1d81b0102")
	if err != nil {
		t.Fatal(err)
	}
	want := map[any]any{cbor.Tag{Number: 27, Content: int8(1)}: int8(2)}
	if !reflect.DeepEqual(v, want) {
		t.Errorf("got %#v, want %#v", v, want)
	}

	// A tag enclosing an array cannot be a Go map key.
	if _, err := decodeHex(t, "a1d8188201020304"); !errors.Is(err, cbor.ErrDecode) {
		t.Errorf("got err=%v, want %v", err, cbor.ErrDecode)
	}
}

func TestRFC8949TagEncode(t *testing.T) {
	tests := []struct {
		tag      any
		expected string
	}{
		{cbor.NewTag(32, "http://www.example.com"), "d82076687474703a2f2f7777772e6578616d706c652e636f6d"},
		{cbor.NewTag(23, []byte{0x01, 0x02, 0x03, 0x04}), "d74401020304"},
		{cbor.NewTag(24, uint8(1)), "d81801"},
		{cbor.NewTag(1000, uint8(1)), "d903e801"},
		{cbor.NewTag(65536, uint8(1)), "da0001000001"},
		{cbor.NewTag(4294967296, uint8(1)), "db000000010000000001"},
		{&cbor.Tag{Number: 2, Content: []byte{0x01}}, "c24101"},
		{(*cbor.Tag)(nil), "f6"},
		{cbor.NewTag(1, cbor.NewTag(2, uint8(1))), "c1c201"},
	}
	for _, test := range tests {
		b, err := cbor.Marshal(test.tag)
		if err != nil {
			t.Fatal(err)
		}
		if got := hex.EncodeToString(b); got != test.expected {
			t.Errorf("%#v: got %s, want %s", test.tag, got, test.expected)
		}
	}

	// Round trip
	tag := cbor.NewTag(32, "http://www.example.com")
	b, _ := cbor.Marshal(tag)
	v, err := cbor.Unmarshal(b)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(v, tag) {
		t.Errorf("got %#v, want %#v", v, tag)
	}
}
