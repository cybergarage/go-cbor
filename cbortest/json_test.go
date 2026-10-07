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
	"encoding/json"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/cybergarage/go-cbor/cbor"
)

func toJSONHex(t *testing.T, h string) (string, error) {
	t.Helper()
	b, err := hex.DecodeString(h)
	if err != nil {
		t.Fatal(err)
	}
	j, err := cbor.ToJSON(b)
	return string(j), err
}

// RFC 8949 Section 6.1 with the examples in Appendix A.
func TestToJSON(t *testing.T) {
	tests := []struct {
		cbor string
		json string
	}{
		// Integers
		{"00", "0"},
		{"17", "23"},
		{"1903e8", "1000"},
		{"1bffffffffffffffff", "18446744073709551615"},
		{"20", "-1"},
		{"3863", "-100"},
		{"3bffffffffffffffff", "-18446744073709551616"},
		// Bignums
		{"c249010000000000000000", `"AQAAAAAAAAAA"`},
		{"c349010000000000000000", `"~AQAAAAAAAAAA"`},
		// Floating-point values
		{"f90000", "0.0"},
		{"f98000", "-0.0"},
		{"f93c00", "1.0"},
		{"fb3ff199999999999a", "1.1"},
		{"f93e00", "1.5"},
		{"f97bff", "65504.0"},
		{"fa47c35000", "100000.0"},
		{"fa7f7fffff", "3.4028234663852886e+38"},
		{"fb7e37e43c8800759c", "1e+300"},
		{"f90001", "5.960464477539063e-8"},
		{"f90400", "0.00006103515625"},
		{"f9c400", "-4.0"},
		{"fbc010666666666666", "-4.1"},
		{"fa3dcccccd", "0.10000000149011612"},
		{"f97c00", "null"},
		{"f97e00", "null"},
		{"f9fc00", "null"},
		{"fa7f800000", "null"},
		{"fb7ff8000000000000", "null"},
		// Simple values
		{"f4", "false"},
		{"f5", "true"},
		{"f6", "null"},
		{"f7", "null"},
		{"f0", "null"},
		{"f8ff", "null"},
		// Tags
		{"c074323031332d30332d32315432303a30343a30305a", `"2013-03-21T20:04:00Z"`},
		{"c11a514b67b0", "1363896240"},
		{"c1fb41d452d9ec200000", "1363896240.5"},
		{"d74401020304", `"01020304"`},
		{"d818456449455446", `"ZElFVEY"`},
		{"d82076687474703a2f2f7777772e6578616d706c652e636f6d", `"http://www.example.com"`},
		{"d9d9f700", "0"},
		// Strings
		{"40", `""`},
		{"4401020304", `"AQIDBA"`},
		{"60", `""`},
		{"6161", `"a"`},
		{"6449455446", `"IETF"`},
		{"62225c", `"\"\\"`},
		{"62c3bc", `"ü"`},
		{"63e6b0b4", "\"\u6c34\""},
		{"64f0908591", `"𐅑"`},
		{"63" + "0a017f", "\"\\n\\u0001\x7f\""},
		{"63" + "3c263e", `"<&>"`},
		{"63" + "e280a8", "\" \""},
		// Arrays and maps
		{"80", "[]"},
		{"83010203", "[1,2,3]"},
		{"8301820203820405", "[1,[2,3],[4,5]]"},
		{"98190102030405060708090a0b0c0d0e0f101112131415161718181819", "[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25]"},
		{"a0", "{}"},
		{"a201020304", `{"1":2,"3":4}`},
		{"a26161016162820203", `{"a":1,"b":[2,3]}`},
		{"826161a161626163", `["a",{"b":"c"}]`},
		{"a56161614161626142616361436164614461656145", `{"a":"A","b":"B","c":"C","d":"D","e":"E"}`},
		{"a2616201616102", `{"b":1,"a":2}`},
		{"a4f401f502f603fb3ff800000000000004", `{"false":1,"true":2,"null":3,"1.5":4}`},
		{"a1f93c0000", `{"1.0":0}`},
		{"fa61b03030", "406262404035697970000.0"},
		{"a1410100", `{"AQ":0}`},
		// Indefinite-length items
		{"5f42010243030405ff", `"AQIDBAU"`},
		{"7f657374726561646d696e67ff", `"streaming"`},
		{"9fff", "[]"},
		{"9f018202039f0405ffff", "[1,[2,3],[4,5]]"},
		{"bf61610161629f0203ffff", `{"a":1,"b":[2,3]}`},
		{"826161bf61626163ff", `["a",{"b":"c"}]`},
		{"bf6346756ef563416d7421ff", `{"Fun":true,"Amt":-2}`},
	}
	for _, test := range tests {
		t.Run(test.cbor, func(t *testing.T) {
			got, err := toJSONHex(t, test.cbor)
			if err != nil {
				t.Fatal(err)
			}
			if got != test.json {
				t.Errorf("got %s, want %s", got, test.json)
			}
			if !json.Valid([]byte(got)) {
				t.Errorf("%s is not valid JSON", got)
			}
		})
	}
}

// RFC 8949 Section 3.4.5.2.
func TestToJSONExpectedEncodings(t *testing.T) {
	tests := []struct {
		cbor string
		json string
	}{
		{"d54401020304", `"AQIDBA"`},
		{"d64401020304", `"AQIDBA=="`},
		{"d74401020304", `"01020304"`},
		{"d540", `""`},
		{"d640", `""`},
		{"d740", `""`},
		{"d7d64401020304", `"AQIDBA=="`},
		// The hint applies to all byte strings in the tag content.
		{"d78241ab41cd", `["AB","CD"]`},
		{"d7a1616141ab", `{"a":"AB"}`},
		{"d7a141ab00", `{"AB":0}`},
		{"d75f41ab41cdff", `"ABCD"`},
		// except for those in a nested data item tagged with another hint.
		{"d68241ffd741ff", `["/w==","FF"]`},
		// Bignums are always base64url.
		{"d7c249010000000000000000", `"AQAAAAAAAAAA"`},
		// Text strings are not affected.
		{"d76161", `"a"`},
		// Unknown tags keep the hint.
		{"d7d9100041ab", `"AB"`},
	}
	for _, test := range tests {
		got, err := toJSONHex(t, test.cbor)
		if err != nil {
			t.Fatalf("%s: %v", test.cbor, err)
		}
		if got != test.json {
			t.Errorf("%s: got %s, want %s", test.cbor, got, test.json)
		}
	}
}

func TestToJSONErrors(t *testing.T) {
	jsonErrors := []string{
		"a1800000",   // array key
		"a1a00000",   // map key
		"a201006131", // 1 and "1" collide
		"a2" + "0100" + "6131" + "00",
		"a2f600646e756c6c00",
	}
	for _, test := range jsonErrors {
		if _, err := toJSONHex(t, test); !errors.Is(err, cbor.ErrJSON) {
			t.Errorf("%s: got err=%v, want %v", test, err, cbor.ErrJSON)
		}
	}
	decodeErrors := []string{
		"ff", "81ff", "a1ff", "0001", "62c328", "c201", "c26161", "f800", "1c",
	}
	for _, test := range decodeErrors {
		if _, err := toJSONHex(t, test); err == nil {
			t.Errorf("%s: expected an error", test)
		}
	}
	for _, test := range []string{"81", "a101", "5f41", "9f01", "c2"} {
		if _, err := toJSONHex(t, test); err == nil {
			t.Errorf("%s: expected an error", test)
		}
	}

	// Nesting depth
	if _, err := toJSONHex(t, strings.Repeat("81", cbor.DefaultMaxNestedLevels)+"00"); err != nil {
		t.Error(err)
	}
	if _, err := toJSONHex(t, strings.Repeat("81", cbor.DefaultMaxNestedLevels+1)+"00"); !errors.Is(err, cbor.ErrDecode) {
		t.Errorf("got err=%v, want %v", err, cbor.ErrDecode)
	}

	// Decode mode
	for _, test := range []string{"1817", "9fff", "fa3fc00000", "a2616201616102", "c24101"} {
		b, _ := hex.DecodeString(test)
		dec := cbor.NewDecoder(bytes.NewReader(b))
		dec.SetDecodeMode(cbor.DecodeModeCoreDeterministic)
		if _, err := dec.DecodeJSON(); !errors.Is(err, cbor.ErrDecode) {
			t.Errorf("%s: got err=%v, want %v", test, err, cbor.ErrDecode)
		}
		if _, err := toJSONHex(t, test); err != nil {
			t.Errorf("%s: lenient: %v", test, err)
		}
	}
}

func TestDecodeJSONSequence(t *testing.T) {
	b, _ := hex.DecodeString("01a161610282f5f6")
	dec := cbor.NewDecoder(bytes.NewReader(b))
	var got []string
	for {
		j, err := dec.DecodeJSON()
		if err != nil {
			if !errors.Is(err, io.EOF) {
				t.Fatal(err)
			}
			break
		}
		got = append(got, string(j))
	}
	if want := []string{"1", `{"a":2}`, "[true,null]"}; strings.Join(got, " ") != strings.Join(want, " ") {
		t.Errorf("got %v, want %v", got, want)
	}
}

// RFC 8949 Section 6.2.
func TestFromJSON(t *testing.T) {
	tests := []struct {
		json string
		cbor string
	}{
		{"0", "00"},
		{"-0", "00"},
		{"23", "17"},
		{"1000", "1903e8"},
		{"-1", "20"},
		{"-1000", "3903e7"},
		{"18446744073709551615", "1bffffffffffffffff"},
		{"-18446744073709551616", "3bffffffffffffffff"},
		{"18446744073709551616", "c249010000000000000000"},
		{"-18446744073709551617", "c349010000000000000000"},
		{"0.0", "f90000"},
		{"-0.0", "f98000"},
		{"1.0", "f93c00"},
		{"1.1", "fb3ff199999999999a"},
		{"1.5", "f93e00"},
		{"65504.0", "f97bff"},
		{"100000.0", "fa47c35000"},
		{"3.4028234663852886e+38", "fa7f7fffff"},
		{"1e300", "fb7e37e43c8800759c"},
		{"1E2", "f95640"},
		{"5.960464477539063e-8", "f90001"},
		{"-4.1", "fbc010666666666666"},
		{"true", "f5"},
		{"false", "f4"},
		{"null", "f6"},
		{`""`, "60"},
		{`"a"`, "6161"},
		{`"ü"`, "62c3bc"},
		{`"\"\\"`, "62225c"},
		{`"𐅑"`, "64f0908591"},
		{"[]", "80"},
		{"[1,2,3]", "83010203"},
		{"[1,[2,3],[4,5]]", "8301820203820405"},
		{"{}", "a0"},
		{`{"a":1,"b":[2,3]}`, "a26161016162820203"},
		{`["a",{"b":"c"}]`, "826161a161626163"},
		{`{"b":1,"a":2}`, "a2616201616102"},
		{" [ 1 , {\"a\" : null} ]\n", "8201a16161f6"},
	}
	for _, test := range tests {
		b, err := cbor.FromJSON([]byte(test.json))
		if err != nil {
			t.Errorf("%s: %v", test.json, err)
			continue
		}
		if got := hex.EncodeToString(b); got != test.cbor {
			t.Errorf("%s: got %s, want %s", test.json, got, test.cbor)
		}
	}
}

func TestEncodeJSONModes(t *testing.T) {
	encode := func(mode cbor.EncodeMode, j string) string {
		var buf bytes.Buffer
		enc := cbor.NewEncoder(&buf)
		enc.SetEncodeMode(mode)
		if err := enc.EncodeJSON([]byte(j)); err != nil {
			t.Fatalf("%s: %v", j, err)
		}
		return hex.EncodeToString(buf.Bytes())
	}
	tests := []struct {
		mode cbor.EncodeMode
		json string
		cbor string
	}{
		{cbor.EncodeModeCoreDeterministic, `{"b":1,"a":2}`, "a2616102616201"},
		{cbor.EncodeModeCoreDeterministic, `{"b":{"z":1.5,"aa":2},"a":[1]}`, "a2616181016162a2617af93e0062616102"},
		{cbor.EncodeModeLengthFirstDeterministic, `{"aa":1,"b":2}`, "a2616202626161" + "01"},
		{cbor.EncodeModeTypePreserving, "1.5", "fb3ff8000000000000"},
		{cbor.EncodeModeTypePreserving, "1000", "1903e8"},
	}
	for _, test := range tests {
		if got := encode(test.mode, test.json); got != test.cbor {
			t.Errorf("%s %s: got %s, want %s", test.mode, test.json, got, test.cbor)
		}
	}

	// In an indefinite-length array
	var buf bytes.Buffer
	enc := cbor.NewEncoder(&buf)
	if err := enc.StartIndefiniteArray(); err != nil {
		t.Fatal(err)
	}
	for _, j := range []string{"1", `{"a":true}`} {
		if err := enc.EncodeJSON([]byte(j)); err != nil {
			t.Fatal(err)
		}
	}
	if err := enc.EndIndefinite(); err != nil {
		t.Fatal(err)
	}
	if got := hex.EncodeToString(buf.Bytes()); got != "9f01a16161f5ff" {
		t.Errorf("got %s, want 9f01a16161f5ff", got)
	}
}

func TestFromJSONErrors(t *testing.T) {
	tests := []string{
		"",
		"[1,",
		"{\"a\"}",
		"1 2",
		"[1]]",
		`{"a":1,"a":2}`,
		"1e400",
		"-1e400",
		"NaN",
		strings.Repeat("[", cbor.DefaultMaxNestedLevels+1) + strings.Repeat("]", cbor.DefaultMaxNestedLevels+1),
	}
	for _, test := range tests {
		if _, err := cbor.FromJSON([]byte(test)); !errors.Is(err, cbor.ErrJSON) {
			t.Errorf("%q: got err=%v, want %v", test, err, cbor.ErrJSON)
		}
	}
	deep := strings.Repeat("[", cbor.DefaultMaxNestedLevels) + strings.Repeat("]", cbor.DefaultMaxNestedLevels)
	if _, err := cbor.FromJSON([]byte(deep)); err != nil {
		t.Error(err)
	}
}

// JSON -> CBOR -> JSON keeps the JSON values.
func TestJSONRoundTrip(t *testing.T) {
	tests := []string{
		`{"name":"go-cbor","version":1.5,"tags":["cbor","json"],"nested":{"empty":{},"list":[],"null":null}}`,
		`[0,-1,18446744073709551615,-18446744073709551616,1.1,1e+300,5e-324,1.0,-0.0,100000.0,true,false,"\u0000\u001f\"\\"]`,
		`{"b":1,"a":2,"c":{"z":[1,[2,[3]]]}}`,
	}
	for _, test := range tests {
		c, err := cbor.FromJSON([]byte(test))
		if err != nil {
			t.Fatal(err)
		}
		got, err := cbor.ToJSON(c)
		if err != nil {
			t.Fatal(err)
		}
		if string(got) != test {
			t.Errorf("got %s, want %s", got, test)
		}
	}
}
