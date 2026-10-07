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
	"math"
	"reflect"
	"strings"
	"testing"

	"github.com/cybergarage/go-cbor/cbor"
)

func encodeWithMode(t *testing.T, mode cbor.EncodeMode, sort bool, v any) (string, error) {
	t.Helper()
	var buf bytes.Buffer
	enc := cbor.NewEncoder(&buf)
	enc.SetEncodeMode(mode)
	enc.SetMapSortEnabled(sort)
	if err := enc.Encode(v); err != nil {
		return "", err
	}
	return hex.EncodeToString(buf.Bytes()), nil
}

func TestEncodeModeDefault(t *testing.T) {
	if mode := cbor.NewConfig().EncodeMode(); mode != cbor.EncodeModePreferred {
		t.Errorf("default encode mode is %s, want %s", mode, cbor.EncodeModePreferred)
	}
	for mode, name := range map[cbor.EncodeMode]string{
		cbor.EncodeModePreferred:                "Preferred",
		cbor.EncodeModeTypePreserving:           "TypePreserving",
		cbor.EncodeModeCoreDeterministic:        "CoreDeterministic",
		cbor.EncodeModeLengthFirstDeterministic: "LengthFirstDeterministic",
	} {
		if mode.String() != name {
			t.Errorf("got %s, want %s", mode.String(), name)
		}
	}
}

// RFC 8949 Section 4.1 and the integer examples in Appendix A.
func TestPreferredSerializationIntegers(t *testing.T) {
	tests := []struct {
		value    any
		expected string
	}{
		{uint8(0), "00"},
		{int(1), "01"},
		{int64(10), "0a"},
		{uint16(23), "17"},
		{int(24), "1818"},
		{int32(25), "1819"},
		{int(100), "1864"},
		{uint(1000), "1903e8"},
		{int(1000000), "1a000f4240"},
		{int64(1000000000000), "1b000000e8d4a51000"},
		{uint64(math.MaxUint64), "1bffffffffffffffff"},
		{int(-1), "20"},
		{int8(-10), "29"},
		{int(-100), "3863"},
		{int16(-1000), "3903e7"},
		{int64(math.MinInt64), "3b7fffffffffffffff"},
		{[]int{1, 2, 3}, "83010203"},
	}
	for _, test := range tests {
		for _, mode := range []cbor.EncodeMode{cbor.EncodeModePreferred, cbor.EncodeModeCoreDeterministic, cbor.EncodeModeLengthFirstDeterministic} {
			got, err := encodeWithMode(t, mode, false, test.value)
			if err != nil {
				t.Fatal(err)
			}
			if got != test.expected {
				t.Errorf("%s: %v (%T): got %s, want %s", mode, test.value, test.value, got, test.expected)
			}
		}
	}
}

// RFC 8949 Section 4.1 and the floating-point examples in Appendix A.
func TestPreferredSerializationFloats(t *testing.T) {
	tests := []struct {
		value    any
		expected string
	}{
		{float64(0.0), "f90000"},
		{math.Copysign(0, -1), "f98000"},
		{float64(1.0), "f93c00"},
		{float64(1.1), "fb3ff199999999999a"},
		{float64(1.5), "f93e00"},
		{float64(65504.0), "f97bff"},
		{float64(100000.0), "fa47c35000"},
		{float64(3.4028234663852886e+38), "fa7f7fffff"},
		{float64(1.0e+300), "fb7e37e43c8800759c"},
		{float64(5.960464477539063e-8), "f90001"},
		{float64(0.00006103515625), "f90400"},
		{float64(-4.0), "f9c400"},
		{float64(-4.1), "fbc010666666666666"},
		{math.Inf(1), "f97c00"},
		{math.NaN(), "f97e00"},
		{math.Inf(-1), "f9fc00"},
		{float32(1.5), "f93e00"},
		{float32(100000.0), "fa47c35000"},
		{float32(-4.1), "fac0833333"},
		{float32(math.NaN()), "f97e00"},
		{float64(65505.0), "fa477fe100"}, // not exact in float16
		{float64(1.0 / (1 << 25)), "fa33000000"},
	}
	for _, test := range tests {
		got, err := encodeWithMode(t, cbor.EncodeModePreferred, false, test.value)
		if err != nil {
			t.Fatal(err)
		}
		if got != test.expected {
			t.Errorf("%v (%T): got %s, want %s", test.value, test.value, got, test.expected)
		}
	}
}

// Every half-precision value must be encoded back to the same half-precision bits.
func TestPreferredSerializationAllFloat16(t *testing.T) {
	for bits := range 0x10000 {
		in := []byte{0xF9, byte(bits >> 8), byte(bits)}
		v, err := cbor.Unmarshal(in)
		if err != nil {
			t.Fatal(err)
		}
		f := v.(float64)
		want := hex.EncodeToString(in)
		if math.IsNaN(f) {
			want = "f97e00"
		}
		got, err := encodeWithMode(t, cbor.EncodeModePreferred, false, f)
		if err != nil {
			t.Fatal(err)
		}
		if got != want {
			t.Fatalf("float16 0x%04x (%v): got %s, want %s", bits, f, got, want)
		}
	}
}

// The preferred serialization uses the shortest length for strings, arrays, and maps.
func TestPreferredSerializationLengths(t *testing.T) {
	tests := []struct {
		length int
		header string
	}{
		{23, "77"},
		{24, "7818"},
		{255, "78ff"},
		{256, "790100"},
		{65535, "79ffff"},
		{65536, "7a00010000"},
	}
	for _, test := range tests {
		for _, mode := range []cbor.EncodeMode{cbor.EncodeModePreferred, cbor.EncodeModeTypePreserving} {
			got, err := encodeWithMode(t, mode, false, strings.Repeat("a", test.length))
			if err != nil {
				t.Fatal(err)
			}
			if !strings.HasPrefix(got, test.header) || len(got) != len(test.header)+2*test.length {
				t.Errorf("%s: length %d: got header %s, want %s", mode, test.length, got[:len(test.header)], test.header)
			}
		}
	}
}

func TestTypePreservingMode(t *testing.T) {
	tests := []struct {
		value    any
		expected string
	}{
		{int(1000), "1b00000000000003e8"},
		{int(-1000), "3b00000000000003e7"},
		{uint16(5), "190005"},
		{int8(5), "05"},
		{float64(1.5), "fb3ff8000000000000"},
		{float32(1.5), "fa3fc00000"},
		{[]int{1, 2, 3}, "831b00000000000000011b00000000000000021b0000000000000003"},
	}
	for _, test := range tests {
		got, err := encodeWithMode(t, cbor.EncodeModeTypePreserving, false, test.value)
		if err != nil {
			t.Fatal(err)
		}
		if got != test.expected {
			t.Errorf("%v (%T): got %s, want %s", test.value, test.value, got, test.expected)
		}
	}

	// Fixed-width Go types are decoded back to the same types.
	for _, v := range []any{int64(1000), int32(-1000), int16(5), uint16(40000), float64(1.5), float32(1.5)} {
		got, err := encodeWithMode(t, cbor.EncodeModeTypePreserving, false, v)
		if err != nil {
			t.Fatal(err)
		}
		b, _ := hex.DecodeString(got)
		r, err := cbor.Unmarshal(b)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(r, v) {
			t.Errorf("got %v (%T), want %v (%T)", r, r, v, v)
		}
	}
}

// The map in the examples of RFC 8949 Section 4.2.1 and Section 4.2.3.
func rfc8949OrderingMap() map[any]any {
	return map[any]any{
		int(10):     uint8(0),
		int(100):    uint8(1),
		int(-1):     uint8(2),
		"z":         uint8(3),
		"aa":        uint8(4),
		[1]int{100}: uint8(5),
		[1]int{-1}:  uint8(6),
		false:       uint8(7),
	}
}

func TestCoreDeterministicMapKeyOrder(t *testing.T) {
	// 4.2.1: 10, 100, -1, "z", "aa", [100], [-1], false
	want := "a8" + "0a00" + "186401" + "2002" + "617a03" + "62616104" + "81186405" + "812006" + "f407"
	for range 20 {
		got, err := encodeWithMode(t, cbor.EncodeModeCoreDeterministic, false, rfc8949OrderingMap())
		if err != nil {
			t.Fatal(err)
		}
		if got != want {
			t.Fatalf("got %s, want %s", got, want)
		}
		// MapSortEnabled uses the same bytewise lexicographic order in the other modes.
		for _, mode := range []cbor.EncodeMode{cbor.EncodeModePreferred, cbor.EncodeModeTypePreserving} {
			got, err := encodeWithMode(t, mode, true, map[any]any{"z": uint8(3), "aa": uint8(4), false: uint8(7), int8(-1): uint8(2), int8(10): uint8(0)})
			if err != nil {
				t.Fatal(err)
			}
			if want := "a5" + "0a00" + "2002" + "617a03" + "62616104" + "f407"; got != want {
				t.Fatalf("%s with MapSortEnabled: got %s, want %s", mode, got, want)
			}
		}
	}
}

func TestLengthFirstDeterministicMapKeyOrder(t *testing.T) {
	// 4.2.3: 10, -1, false, 100, "z", [-1], "aa", [100]
	want := "a8" + "0a00" + "2002" + "f407" + "186401" + "617a03" + "812006" + "62616104" + "81186405"
	for range 20 {
		got, err := encodeWithMode(t, cbor.EncodeModeLengthFirstDeterministic, false, rfc8949OrderingMap())
		if err != nil {
			t.Fatal(err)
		}
		if got != want {
			t.Fatalf("got %s, want %s", got, want)
		}
	}
}

func TestDeterministicStructAndNestedMaps(t *testing.T) {
	type inner struct {
		Zz uint8
		A  uint8
	}
	type outer struct {
		Map   map[string]uint8
		Inner inner
	}
	v := outer{Map: map[string]uint8{"b": 2, "a": 1}, Inner: inner{Zz: 1, A: 2}}
	// {"Map": {"a": 1, "b": 2}, "Inner": {"A": 2, "Zz": 1}} with keys sorted bytewise:
	// "Map" (634d6170) < "Inner" (65496e6e6572)
	want := "a2" + "634d6170" + "a2" + "616101" + "616202" + "65496e6e6572" + "a2" + "614102" + "625a7a01"
	got, err := encodeWithMode(t, cbor.EncodeModeCoreDeterministic, false, v)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

// Different Go keys that have the same encoding must not produce duplicate map keys (RFC 8949 Section 5.6).
func TestDeterministicDuplicateKeys(t *testing.T) {
	m := map[any]any{int8(1): "a", int(1): "b"}
	for _, mode := range []cbor.EncodeMode{cbor.EncodeModeCoreDeterministic, cbor.EncodeModeLengthFirstDeterministic} {
		if _, err := encodeWithMode(t, mode, false, m); !errors.Is(err, cbor.ErrEncode) {
			t.Errorf("%s: got err=%v, want %v", mode, err, cbor.ErrEncode)
		}
	}
	if _, err := encodeWithMode(t, cbor.EncodeModePreferred, true, m); !errors.Is(err, cbor.ErrEncode) {
		t.Errorf("Preferred with MapSortEnabled: got err=%v, want %v", err, cbor.ErrEncode)
	}
	// The keys are different in the type-preserving mode.
	if _, err := encodeWithMode(t, cbor.EncodeModeTypePreserving, true, m); err != nil {
		t.Errorf("TypePreserving with MapSortEnabled: %v", err)
	}
}
