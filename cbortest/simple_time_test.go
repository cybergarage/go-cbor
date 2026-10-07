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
	"encoding/hex"
	"errors"
	"math/big"
	"reflect"
	"testing"
	"time"

	"github.com/cybergarage/go-cbor/cbor"
)

// RFC 8949 Section 3.3 and the simple value examples in Appendix A.
func TestRFC8949SimpleValues(t *testing.T) {
	tests := []struct {
		encoded  string
		expected any
	}{
		{"f4", false},
		{"f5", true},
		{"f6", nil},
		{"f7", cbor.Undefined},
		{"e0", cbor.SimpleValue(0)},
		{"f0", cbor.SimpleValue(16)},
		{"f3", cbor.SimpleValue(19)},
		{"f820", cbor.SimpleValue(32)},
		{"f8ff", cbor.SimpleValue(255)},
		{"83f7f6f0", []any{cbor.Undefined, nil, cbor.SimpleValue(16)}},
		{"a1f701", map[any]any{cbor.Undefined: int8(1)}},
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

	// 3.3. Simple values 0..31 must not be encoded in the following byte.
	for _, test := range []string{"f800", "f813", "f814", "f817", "f818", "f81f"} {
		t.Run(test, func(t *testing.T) {
			if _, err := decodeHex(t, test); !errors.Is(err, cbor.ErrDecode) {
				t.Errorf("%s: got err=%v, want %v", test, err, cbor.ErrDecode)
			}
		})
	}
}

func TestRFC8949SimpleValueEncode(t *testing.T) {
	tests := []struct {
		value    cbor.SimpleValue
		expected string
	}{
		{cbor.SimpleValue(0), "e0"},
		{cbor.SimpleValue(16), "f0"},
		{cbor.SimpleValue(20), "f4"},
		{cbor.SimpleValue(21), "f5"},
		{cbor.SimpleValue(22), "f6"},
		{cbor.Undefined, "f7"},
		{cbor.SimpleValue(32), "f820"},
		{cbor.SimpleValue(255), "f8ff"},
	}
	for _, test := range tests {
		b, err := cbor.Marshal(test.value)
		if err != nil {
			t.Fatal(err)
		}
		if got := hex.EncodeToString(b); got != test.expected {
			t.Errorf("%d: got %s, want %s", test.value, got, test.expected)
		}
	}

	b, err := cbor.Marshal([]any{cbor.Undefined, uint8(1)})
	if err != nil {
		t.Fatal(err)
	}
	if got := hex.EncodeToString(b); got != "82f701" {
		t.Errorf("got %s, want 82f701", got)
	}

	// 24..31 are reserved.
	for v := 24; v < 32; v++ {
		if _, err := cbor.Marshal(cbor.SimpleValue(v)); !errors.Is(err, cbor.ErrEncode) {
			t.Errorf("%d: got err=%v, want %v", v, err, cbor.ErrEncode)
		}
	}
}

// RFC 8949 Section 3.4.2 and Appendix A.
func TestRFC8949EpochDateTime(t *testing.T) {
	tests := []struct {
		encoded  string
		expected time.Time
	}{
		{"c11a514b67b0", time.Date(2013, 3, 21, 20, 4, 0, 0, time.UTC)},
		{"c1fb41d452d9ec200000", time.Date(2013, 3, 21, 20, 4, 0, 500000000, time.UTC)},
		{"c100", time.Unix(0, 0).UTC()},
		{"c120", time.Unix(-1, 0).UTC()},
		{"c13a7fffffff", time.Unix(-2147483648, 0).UTC()},
		{"c11b0000000100000000", time.Unix(4294967296, 0).UTC()},
		{"c1fbbff8000000000000", time.Unix(-2, 500000000).UTC()}, // -1.5
		{"c1f93c00", time.Unix(1, 0).UTC()},                      // float16 1.0
		{"c1fa47c35000", time.Unix(100000, 0).UTC()},             // float32 100000.0
	}
	for _, test := range tests {
		t.Run(test.encoded, func(t *testing.T) {
			v, err := decodeHex(t, test.encoded)
			if err != nil {
				t.Fatal(err)
			}
			got, ok := v.(time.Time)
			if !ok {
				t.Fatalf("%s: got %#v (%T), want time.Time", test.encoded, v, v)
			}
			if !got.Equal(test.expected) || got.Location() != time.UTC {
				t.Errorf("%s: got %v, want %v", test.encoded, got, test.expected)
			}
		})
	}

	invalid := []string{
		"c16161",               // text string
		"c1f6",                 // null
		"c180",                 // array
		"c1f97e00",             // NaN
		"c1f97c00",             // +Infinity
		"c1f9fc00",             // -Infinity
		"c11bffffffffffffffff", // out of the int64 range
		"c13bffffffffffffffff", // out of the int64 range (big.Int)
		"c1fb7e37e43c8800759c", // 1e300
		"c001",                 // tag 0 with an integer
	}
	for _, test := range invalid {
		t.Run(test, func(t *testing.T) {
			if _, err := decodeHex(t, test); !errors.Is(err, cbor.ErrDecode) {
				t.Errorf("%s: got err=%v, want %v", test, err, cbor.ErrDecode)
			}
		})
	}
}

func TestRFC8949EpochDateTimeUnmarshal(t *testing.T) {
	b, _ := hex.DecodeString("c1fb41d452d9ec200000")
	var v time.Time
	if err := cbor.UnmarshalTo(b, &v); err != nil {
		t.Fatal(err)
	}
	if want := time.Date(2013, 3, 21, 20, 4, 0, 500000000, time.UTC); !v.Equal(want) {
		t.Errorf("got %v, want %v", v, want)
	}
}

// RFC 8949 Section 3.4.1: fractional seconds are kept when time.Time is encoded as tag 0.
func TestRFC8949StdDateTimeFraction(t *testing.T) {
	src := time.Date(2013, 3, 21, 20, 4, 0, 500000000, time.UTC)
	b, err := cbor.Marshal(src)
	if err != nil {
		t.Fatal(err)
	}
	// 0("2013-03-21T20:04:00.5Z")
	if got, want := hex.EncodeToString(b), "c076323031332d30332d32315432303a30343a30302e355a"; got != want {
		t.Errorf("got %s, want %s", got, want)
	}
	v, err := cbor.Unmarshal(b)
	if err != nil {
		t.Fatal(err)
	}
	if got, ok := v.(time.Time); !ok || !got.Equal(src) {
		t.Errorf("got %v, want %v", v, src)
	}
}

// RFC 8949 Section 3.4.3 and the bignum examples in Appendix A.
func TestRFC8949Bignums(t *testing.T) {
	bigInt := func(s string) *big.Int {
		v, ok := new(big.Int).SetString(s, 10)
		if !ok {
			t.Fatalf("invalid number: %s", s)
		}
		return v
	}
	tests := []struct {
		encoded  string
		expected *big.Int
	}{
		{"c249010000000000000000", bigInt("18446744073709551616")},
		{"c349010000000000000000", bigInt("-18446744073709551617")},
		{"c240", bigInt("0")},
		{"c340", bigInt("-1")},
		{"c24101", bigInt("1")},
		{"c25f41014100ff", bigInt("256")}, // indefinite-length byte string
	}
	for _, test := range tests {
		t.Run(test.encoded, func(t *testing.T) {
			v, err := decodeHex(t, test.encoded)
			if err != nil {
				t.Fatal(err)
			}
			got, ok := v.(*big.Int)
			if !ok || got.Cmp(test.expected) != 0 {
				t.Errorf("%s: got %v (%T), want %v", test.encoded, v, v, test.expected)
			}
		})
	}

	for _, test := range []string{"c201", "c36161", "c2f6"} {
		t.Run(test, func(t *testing.T) {
			if _, err := decodeHex(t, test); !errors.Is(err, cbor.ErrDecode) {
				t.Errorf("%s: got err=%v, want %v", test, err, cbor.ErrDecode)
			}
		})
	}
}

func TestRFC8949BigIntEncode(t *testing.T) {
	bigInt := func(s string) *big.Int {
		v, _ := new(big.Int).SetString(s, 10)
		return v
	}
	tests := []struct {
		value    any
		expected string
	}{
		{bigInt("0"), "00"},
		{bigInt("23"), "17"},
		{bigInt("1000"), "1903e8"},
		{bigInt("18446744073709551615"), "1bffffffffffffffff"},
		{bigInt("18446744073709551616"), "c249010000000000000000"},
		{bigInt("-1"), "20"},
		{bigInt("-1000"), "3903e7"},
		{bigInt("-18446744073709551616"), "3bffffffffffffffff"},
		{bigInt("-18446744073709551617"), "c349010000000000000000"},
		{*bigInt("1000"), "1903e8"},
		{(*big.Int)(nil), "f6"},
	}
	for _, test := range tests {
		b, err := cbor.Marshal(test.value)
		if err != nil {
			t.Fatal(err)
		}
		if got := hex.EncodeToString(b); got != test.expected {
			t.Errorf("%v: got %s, want %s", test.value, got, test.expected)
		}
	}

	// Values decoded as *big.Int can be encoded again.
	for _, h := range []string{"3b8000000000000000", "3bffffffffffffffff", "c249010000000000000000", "c349010000000000000000"} {
		v, err := decodeHex(t, h)
		if err != nil {
			t.Fatal(err)
		}
		b, err := cbor.Marshal(v)
		if err != nil {
			t.Fatal(err)
		}
		if got := hex.EncodeToString(b); got != h {
			t.Errorf("round trip: got %s, want %s", got, h)
		}
	}
}

func TestStructUnexportedFields(t *testing.T) {
	type withUnexported struct {
		Name   string
		hidden int
		Value  uint8
	}
	b, err := cbor.Marshal(withUnexported{Name: "a", hidden: 1, Value: 2})
	if err != nil {
		t.Fatal(err)
	}
	// {"Name": "a", "Value": 2}
	if got, want := hex.EncodeToString(b), "a2644e616d6561616556616c756502"; got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
