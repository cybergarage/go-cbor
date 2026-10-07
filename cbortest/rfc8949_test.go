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
	"math"
	"math/big"
	"reflect"
	"testing"

	"github.com/cybergarage/go-cbor/cbor"
)

func decodeHex(t *testing.T, h string) (any, error) {
	t.Helper()
	b, err := hex.DecodeString(h)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("%s: decoder panicked: %v", h, r)
		}
	}()
	return cbor.NewDecoder(bytes.NewReader(b)).Decode()
}

// RFC 8949 Section 3.1 (major type 1) and Appendix A.
func TestRFC8949NegativeIntegers(t *testing.T) {
	bigNeg := func(s string) *big.Int {
		v, _ := new(big.Int).SetString(s, 10)
		return v
	}
	tests := []struct {
		encoded  string
		expected any
	}{
		{"20", int8(-1)},
		{"37", int8(-24)},
		{"3817", int8(-24)},
		{"3818", int8(-25)},
		{"387f", int8(-128)},
		{"3880", int16(-129)},
		{"38c8", int16(-201)},
		{"38ff", int16(-256)},
		{"3900ff", int16(-256)},
		{"397fff", int16(math.MinInt16)},
		{"398000", int32(-32769)},
		{"39ffff", int32(-65536)},
		{"3a7fffffff", int32(math.MinInt32)},
		{"3a80000000", int64(-2147483649)},
		{"3affffffff", int64(-4294967296)},
		{"3b7fffffffffffffff", int64(math.MinInt64)},
		{"3b8000000000000000", bigNeg("-9223372036854775809")},
		{"3bffffffffffffffff", bigNeg("-18446744073709551616")},
	}
	for _, test := range tests {
		t.Run(test.encoded, func(t *testing.T) {
			v, err := decodeHex(t, test.encoded)
			if err != nil {
				t.Fatal(err)
			}
			if expected, ok := test.expected.(*big.Int); ok {
				got, ok := v.(*big.Int)
				if !ok || got.Cmp(expected) != 0 {
					t.Errorf("%s: got %v (%T), want %v (*big.Int)", test.encoded, v, v, expected)
				}
				return
			}
			if !reflect.DeepEqual(v, test.expected) {
				t.Errorf("%s: got %v (%T), want %v (%T)", test.encoded, v, v, test.expected, test.expected)
			}
		})
	}
}

func TestRFC8949NegativeIntegerRoundTrip(t *testing.T) {
	values := []any{
		int8(math.MinInt8), int8(-1),
		int16(math.MinInt16), int16(-129),
		int32(math.MinInt32), int32(-32769),
		int64(math.MinInt64), int64(-2147483649),
	}
	for _, v := range values {
		b, err := cbor.Marshal(v)
		if err != nil {
			t.Fatal(err)
		}
		got, err := cbor.Unmarshal(b)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, v) {
			t.Errorf("%v (%T): got %v (%T)", v, v, got, got)
		}
	}
}

// RFC 8949 Appendix F: an incomplete data item is not well-formed.
func TestRFC8949TruncatedInput(t *testing.T) {
	tests := []string{
		"18", "1901", "1a0102", "1b01020304050607",
		"38", "3901", "3a0102", "3b01020304050607",
		"5901", "4449", "6449", "7a00000001",
		"83", "8301", "830102", "9801",
		"a1", "a101", "a20102",
		"c0", "c074",
		"f901", "fa0102", "fb01020304050607",
	}
	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			_, err := decodeHex(t, test)
			if !errors.Is(err, io.ErrUnexpectedEOF) {
				t.Errorf("%s: got err=%v, want %v", test, err, io.ErrUnexpectedEOF)
			}
		})
	}
}

func TestRFC8949DecodeStreamEOF(t *testing.T) {
	b, _ := hex.DecodeString("0102")
	dec := cbor.NewDecoder(bytes.NewReader(b))
	for _, want := range []any{int8(1), int8(2)} {
		got, err := dec.Decode()
		if err != nil || got != want {
			t.Fatalf("got %v, %v; want %v", got, err, want)
		}
	}
	if _, err := dec.Decode(); err != io.EOF { //nolint:errorlint
		t.Errorf("got err=%v, want io.EOF at the end of the stream", err)
	}
}

// RFC 8949 Section 3.1 allows any data item as a map key.
// Keys that cannot be Go map keys must be reported as errors instead of panicking.
func TestRFC8949UnhashableMapKeys(t *testing.T) {
	tests := []string{
		"a182010203", // {[1, 2]: 3}
		"a1a1010203", // {{1: 2}: 3}
		"a142010203", // {h'0102': 3}
		"a201028201020304",
	}
	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			_, err := decodeHex(t, test)
			if !errors.Is(err, cbor.ErrDecode) {
				t.Errorf("%s: got err=%v, want %v", test, err, cbor.ErrDecode)
			}
		})
	}
}

// RFC 8949 Section 10: lengths in the input must not cause overflows or unbounded allocation.
func TestRFC8949HugeLengths(t *testing.T) {
	tooLarge := []string{
		"5bffffffffffffffff",
		"7bffffffffffffffff",
		"9bffffffffffffffff",
		"bbffffffffffffffff",
		"5b8000000000000000",
	}
	for _, test := range tooLarge {
		t.Run(test, func(t *testing.T) {
			_, err := decodeHex(t, test)
			if !errors.Is(err, cbor.ErrDecode) {
				t.Errorf("%s: got err=%v, want %v", test, err, cbor.ErrDecode)
			}
		})
	}

	// Lengths that fit in an int but exceed the available data must fail without allocating the declared size.
	shortData := []string{
		"5b7fffffffffffffff00",
		"7b7fffffffffffffff00",
		"9b7fffffffffffffff00",
		"bb7fffffffffffffff0000",
		"5affffffff00",
	}
	for _, test := range shortData {
		t.Run(test, func(t *testing.T) {
			_, err := decodeHex(t, test)
			if !errors.Is(err, io.ErrUnexpectedEOF) {
				t.Errorf("%s: got err=%v, want %v", test, err, io.ErrUnexpectedEOF)
			}
		})
	}
}

func TestRFC8949LargeByteString(t *testing.T) {
	data := bytes.Repeat([]byte{0xAB}, 200*1024)
	b, err := cbor.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	v, err := cbor.Unmarshal(b)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(v.([]byte), data) {
		t.Error("large byte string was not decoded correctly")
	}
}
