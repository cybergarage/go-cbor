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
	"fmt"
	"io"
	"math"
	"math/big"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/cybergarage/go-cbor/cbor"
)

func decodeHexWith(t *testing.T, h string, setup func(*cbor.Decoder)) (any, error) {
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
	dec := cbor.NewDecoder(bytes.NewReader(b))
	if setup != nil {
		setup(dec)
	}
	return dec.Decode()
}

func deterministicDecoder(dec *cbor.Decoder) {
	dec.SetDecodeMode(cbor.DecodeModeCoreDeterministic)
}

func TestDecoderConfigDefaults(t *testing.T) {
	config := cbor.NewConfig()
	if config.DecodeMode() != cbor.DecodeModeLenient {
		t.Errorf("default decode mode is %s", config.DecodeMode())
	}
	if config.DuplicateMapKeyMode() != cbor.DuplicateMapKeyAllowed {
		t.Errorf("default duplicate map key mode is %d", config.DuplicateMapKeyMode())
	}
	if config.MaxNestedLevels() != cbor.DefaultMaxNestedLevels {
		t.Errorf("default max nested levels is %d", config.MaxNestedLevels())
	}
	if !config.IsUTF8ValidationEnabled() {
		t.Error("UTF-8 validation is disabled by default")
	}
	config.SetMaxNestedLevels(0)
	if config.MaxNestedLevels() != cbor.DefaultMaxNestedLevels {
		t.Errorf("SetMaxNestedLevels(0) set %d", config.MaxNestedLevels())
	}
	if cbor.DecodeModeLenient.String() != "Lenient" || cbor.DecodeModeCoreDeterministic.String() != "CoreDeterministic" {
		t.Error("unexpected decode mode names")
	}
}

// RFC 8949 Section 5.3.1 and Section 3.2.3.
func TestUTF8Validation(t *testing.T) {
	invalid := []string{
		"62c328",       // invalid continuation byte
		"61ff",         // invalid byte
		"63eda080",     // surrogate (U+D800)
		"7f61c361a9ff", // a code point split between chunks
		"a162c32801",   // map key
	}
	for _, test := range invalid {
		t.Run(test, func(t *testing.T) {
			_, err := decodeHexWith(t, test, nil)
			if !errors.Is(err, cbor.ErrDecode) || !errors.Is(err, cbor.ErrInvalidUTF8) {
				t.Errorf("%s: got err=%v, want %v and %v", test, err, cbor.ErrDecode, cbor.ErrInvalidUTF8)
			}
		})
	}

	valid := map[string]string{
		"62c3a9":       "é",
		"7f62c3a9ff":   "é",
		"7f61616162ff": "ab",
		"64f0908591":   "\U00010151",
	}
	for h, want := range valid {
		v, err := decodeHexWith(t, h, nil)
		if err != nil || v != want {
			t.Errorf("%s: got %v, %v; want %q", h, v, err, want)
		}
	}

	// Validation can be disabled.
	v, err := decodeHexWith(t, "62c328", func(dec *cbor.Decoder) { dec.SetUTF8ValidationEnabled(false) })
	if err != nil || v != "\xc3(" {
		t.Errorf("got %q, %v", v, err)
	}
}

func TestUTF8ValidationEncode(t *testing.T) {
	invalid := []any{
		string([]byte{0xc3, 0x28}),
		[]string{"a", "\xff"},
		map[string]int{"\xff": 1},
	}
	for _, v := range invalid {
		_, err := cbor.Marshal(v)
		if !errors.Is(err, cbor.ErrEncode) || !errors.Is(err, cbor.ErrInvalidUTF8) {
			t.Errorf("%q: got err=%v, want %v and %v", v, err, cbor.ErrEncode, cbor.ErrInvalidUTF8)
		}
	}

	var buf bytes.Buffer
	enc := cbor.NewEncoder(&buf)
	enc.SetUTF8ValidationEnabled(false)
	if err := enc.Encode(string([]byte{0xc3, 0x28})); err != nil {
		t.Fatal(err)
	}
	if got := hex.EncodeToString(buf.Bytes()); got != "62c328" {
		t.Errorf("got %s, want 62c328", got)
	}
}

// RFC 8949 Section 10: the nesting depth is limited.
func TestMaxNestedLevels(t *testing.T) {
	nestedArray := func(n int) string { return strings.Repeat("81", n) + "00" }
	nestedIndefinite := func(n int) string { return strings.Repeat("9f", n) + strings.Repeat("ff", n) }
	nestedTag := func(n int) string { return strings.Repeat("d81a", n) + "00" }
	nestedMap := func(n int) string { return strings.Repeat("a100", n) + "00" }

	for _, nested := range []func(int) string{nestedArray, nestedIndefinite, nestedTag, nestedMap} {
		if _, err := decodeHexWith(t, nested(cbor.DefaultMaxNestedLevels), nil); err != nil {
			t.Errorf("depth %d: %v", cbor.DefaultMaxNestedLevels, err)
		}
		if _, err := decodeHexWith(t, nested(cbor.DefaultMaxNestedLevels+1), nil); !errors.Is(err, cbor.ErrDecode) {
			t.Errorf("depth %d: got err=%v, want %v", cbor.DefaultMaxNestedLevels+1, err, cbor.ErrDecode)
		}
	}

	limit2 := func(dec *cbor.Decoder) { dec.SetMaxNestedLevels(2) }
	if _, err := decodeHexWith(t, "818100", limit2); err != nil {
		t.Error(err)
	}
	if _, err := decodeHexWith(t, "81818100", limit2); !errors.Is(err, cbor.ErrDecode) {
		t.Errorf("got err=%v, want %v", err, cbor.ErrDecode)
	}

	// A very deep input fails without exhausting the stack.
	if _, err := decodeHexWith(t, nestedArray(1000000), nil); !errors.Is(err, cbor.ErrDecode) {
		t.Errorf("got err=%v, want %v", err, cbor.ErrDecode)
	}

	// The depth is restored after each item.
	b, _ := hex.DecodeString(strings.Repeat(nestedArray(cbor.DefaultMaxNestedLevels), 3))
	dec := cbor.NewDecoder(bytes.NewReader(b))
	for range 3 {
		if _, err := dec.Decode(); err != nil {
			t.Fatal(err)
		}
	}
}

// RFC 8949 Section 5.6.
func TestDuplicateMapKeys(t *testing.T) {
	v, err := decodeHexWith(t, "a201020103", nil)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(v, map[any]any{int8(1): int8(3)}) {
		t.Errorf("got %v", v)
	}

	reject := func(dec *cbor.Decoder) { dec.SetDuplicateMapKeyMode(cbor.DuplicateMapKeyRejected) }
	for _, test := range []string{
		"a201020103",
		"bf01020103ff",
		"a2616101616102",
		"a201021801" + "03", // 0x1801 is also decoded as int8(1)
		"a2c10102c10103",    // tag keys
	} {
		if _, err := decodeHexWith(t, test, reject); !errors.Is(err, cbor.ErrDecode) {
			t.Errorf("%s: got err=%v, want %v", test, err, cbor.ErrDecode)
		}
	}
	if _, err := decodeHexWith(t, "a201020203", reject); err != nil {
		t.Error(err)
	}
}

// RFC 8949 Section 4.2.1.
func TestCoreDeterministicDecoding(t *testing.T) {
	valid := []string{
		"00", "17", "1818", "18ff", "190100", "1a00010000", "1b0000000100000000",
		"20", "37", "3818", "38ff", "390100",
		"40", "4100", "60", "6161", "80", "a0",
		"f90000", "f93e00", "f97e00", "f97c00", "fa47c35000", "fb3ff199999999999a",
		"f4", "f5", "f6", "f7", "f820",
		"c11a514b67b0", "d81ad81a01",
		"c249010000000000000000",
		// The map in RFC 8949 Section 4.2.1 without the array keys, which cannot be Go map keys.
		"a6" + "0a00" + "186401" + "2002" + "617a03" + "62616104" + "f407",
		"a101a201020304",
		"a2616101616202",
	}
	for _, test := range valid {
		t.Run(test, func(t *testing.T) {
			got, err := decodeHexWith(t, test, deterministicDecoder)
			if err != nil {
				t.Fatalf("%s: %v", test, err)
			}
			want, _ := decodeHexWith(t, test, nil)
			// Compare the printed values, since NaN is not equal to itself.
			if fmt.Sprintf("%#v", got) != fmt.Sprintf("%#v", want) {
				t.Errorf("%s: got %v, want %v", test, got, want)
			}
		})
	}

	invalid := []string{
		// Non-shortest arguments
		"1817", "1900ff", "1a0000ffff", "1b00000000ffffffff",
		"3817", "390017", "5800", "5900014100", "7800", "9800", "b800", "d81700", "d9001a00",
		// Non-shortest floating-point values
		"fa3fc00000", "fb3ff8000000000000", "fb40f86a0000000000", "fa7fc00000", "fb7ff8000000000000", "fa7f800000",
		// Indefinite-length items
		"5fff", "7fff", "9fff", "bfff",
		// Unsorted or duplicate keys
		"a2616201616102", "a2010201 03", "a2" + "2002" + "0a00", "a2" + "626161" + "00" + "617a" + "01",
		// Non-preferred bignums
		"c24100", "c240", "c34101", "c249000100000000000000",
		// Nested
		"81fa3fc00000", "a10181" + "1817",
	}
	for _, test := range invalid {
		test = strings.ReplaceAll(test, " ", "")
		t.Run(test, func(t *testing.T) {
			if _, err := decodeHexWith(t, test, deterministicDecoder); !errors.Is(err, cbor.ErrDecode) {
				t.Errorf("%s: got err=%v, want %v", test, err, cbor.ErrDecode)
			}
			// The lenient mode accepts them, except for duplicate keys which are allowed.
			if _, err := decodeHexWith(t, test, nil); err != nil {
				t.Errorf("%s: lenient: %v", test, err)
			}
		})
	}
}

// Data encoded with EncodeModeCoreDeterministic is accepted by DecodeModeCoreDeterministic.
func TestCoreDeterministicRoundTrip(t *testing.T) {
	type inner struct {
		Zz float64
		A  []any
	}
	values := []any{
		map[any]any{int(10): 0, int(100): 1, int(-1): 2, "z": 3, "aa": 4, false: 7},
		map[string]any{"b": []any{1.5, 1.1, 100000.0, math.Inf(-1)}, "a": inner{Zz: 0.1, A: []any{"x", nil, true}}},
		map[any]any{int64(math.MinInt64): uint64(math.MaxUint64), "": []byte{}, 1.0: "f"},
		[]any{bigInt(t, "18446744073709551616"), bigInt(t, "-18446744073709551617"), cbor.Undefined},
	}
	for _, v := range values {
		var buf bytes.Buffer
		enc := cbor.NewEncoder(&buf)
		enc.SetEncodeMode(cbor.EncodeModeCoreDeterministic)
		if err := enc.Encode(v); err != nil {
			t.Fatal(err)
		}
		h := hex.EncodeToString(buf.Bytes())
		if _, err := decodeHexWith(t, h, deterministicDecoder); err != nil {
			t.Errorf("%s: %v", h, err)
		}
	}
}

// Unmarshal and UnmarshalTo decode a single data item, so extraneous data is an error.
func TestUnmarshalExtraneousData(t *testing.T) {
	b, _ := hex.DecodeString("0102")
	if _, err := cbor.Unmarshal(b); !errors.Is(err, cbor.ErrDecode) {
		t.Errorf("got err=%v, want %v", err, cbor.ErrDecode)
	}
	var v int
	if err := cbor.UnmarshalTo(b, &v); !errors.Is(err, cbor.ErrDecode) {
		t.Errorf("got err=%v, want %v", err, cbor.ErrDecode)
	}
	// Decoder reads a sequence of data items.
	dec := cbor.NewDecoder(bytes.NewReader(b))
	for range 2 {
		if _, err := dec.Decode(); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := dec.Decode(); !errors.Is(err, io.EOF) {
		t.Errorf("got err=%v, want io.EOF", err)
	}
}

func bigInt(t *testing.T, s string) *big.Int {
	t.Helper()
	v, ok := new(big.Int).SetString(s, 10)
	if !ok {
		t.Fatalf("invalid number: %s", s)
	}
	return v
}

// RFC 3339 cannot represent years outside 0000..9999, so such times are encoded as tag 1.
func TestTimeOutsideRFC3339Years(t *testing.T) {
	tests := []struct {
		time     time.Time
		expected string
	}{
		{time.Date(10000, 1, 1, 0, 0, 0, 0, time.UTC), "c11b0000003afff44180"},
		{time.Date(-1, 1, 1, 0, 0, 0, 0, time.UTC), "c13b0000000e7b55af7f"},
		{time.Date(10000, 1, 1, 0, 0, 0, 500000000, time.UTC), "c1fb424d7ffa20c04000"},
	}
	for _, test := range tests {
		b, err := cbor.Marshal(test.time)
		if err != nil {
			t.Fatal(err)
		}
		if got := hex.EncodeToString(b); got != test.expected {
			t.Errorf("%v: got %s, want %s", test.time, got, test.expected)
		}
		v, err := cbor.Unmarshal(b)
		if err != nil {
			t.Fatal(err)
		}
		if got, ok := v.(time.Time); !ok || !got.Equal(test.time) {
			t.Errorf("got %v, want %v", v, test.time)
		}
	}
}
