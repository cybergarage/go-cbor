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
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/cybergarage/go-cbor/cbor"
)

func unmarshalHex(t *testing.T, h string, to any) error {
	t.Helper()
	b, err := hex.DecodeString(h)
	if err != nil {
		t.Fatal(err)
	}
	return cbor.UnmarshalTo(b, to)
}

func TestUnmarshalToAny(t *testing.T) {
	for _, h := range []string{"05", "6161", "8101", "a1616101", "c249010000000000000000", "f7", "d81a01", "f6"} {
		var v any
		if err := unmarshalHex(t, h, &v); err != nil {
			t.Fatalf("%s: %v", h, err)
		}
		want, _ := decodeHex(t, h)
		if !reflect.DeepEqual(v, want) {
			t.Errorf("%s: got %#v, want %#v", h, v, want)
		}
	}
}

func TestUnmarshalToNewTypes(t *testing.T) {
	// *big.Int
	var b big.Int
	if err := unmarshalHex(t, "c349010000000000000000", &b); err != nil || b.String() != "-18446744073709551617" {
		t.Errorf("got %v, %v", b.String(), err)
	}
	if err := unmarshalHex(t, "3b8000000000000000", &b); err != nil || b.String() != "-9223372036854775809" {
		t.Errorf("got %v, %v", b.String(), err)
	}
	var pb *big.Int
	if err := unmarshalHex(t, "c249010000000000000000", &pb); err != nil || pb.String() != "18446744073709551616" {
		t.Errorf("got %v, %v", pb, err)
	}
	var s string
	if err := unmarshalHex(t, "c249010000000000000000", &s); err != nil || s != "18446744073709551616" {
		t.Errorf("got %v, %v", s, err)
	}
	var i64 int64
	if err := unmarshalHex(t, "c24101", &i64); err != nil || i64 != 1 {
		t.Errorf("got %v, %v", i64, err)
	}
	if err := unmarshalHex(t, "c249010000000000000000", &i64); err == nil {
		t.Error("expected an out of range error")
	}

	// *url.URL
	var u url.URL
	if err := unmarshalHex(t, "d82076687474703a2f2f7777772e6578616d706c652e636f6d", &u); err != nil || u.String() != "http://www.example.com" {
		t.Errorf("got %v, %v", u.String(), err)
	}
	var pu *url.URL
	if err := unmarshalHex(t, "d82076687474703a2f2f7777772e6578616d706c652e636f6d", &pu); err != nil || pu.Host != "www.example.com" {
		t.Errorf("got %v, %v", pu, err)
	}
	if err := unmarshalHex(t, "d82076687474703a2f2f7777772e6578616d706c652e636f6d", &s); err != nil || s != "http://www.example.com" {
		t.Errorf("got %v, %v", s, err)
	}

	// cbor.Tag and cbor.SimpleValue
	var tag cbor.Tag
	if err := unmarshalHex(t, "d81a01", &tag); err != nil || !reflect.DeepEqual(tag, cbor.Tag{Number: 26, Content: int8(1)}) {
		t.Errorf("got %v, %v", tag, err)
	}
	var sv cbor.SimpleValue
	if err := unmarshalHex(t, "f7", &sv); err != nil || sv != cbor.Undefined {
		t.Errorf("got %v, %v", sv, err)
	}

	// time.Time keeps fractional seconds in strings.
	if err := unmarshalHex(t, "c1fb41d452d9ec200000", &s); err != nil || s != "2013-03-21T20:04:00.5Z" {
		t.Errorf("got %v, %v", s, err)
	}
}

func TestUnmarshalNull(t *testing.T) {
	p := new(int)
	list := []int{1}
	m := map[string]int{"a": 1}
	var a any = 1
	for _, to := range []any{&p, &list, &m, &a} {
		if err := unmarshalHex(t, "f6", to); err != nil {
			t.Fatalf("%T: %v", to, err)
		}
	}
	if p != nil || list != nil || m != nil || a != nil {
		t.Errorf("got %v, %v, %v, %v", p, list, m, a)
	}
	var n int
	if err := unmarshalHex(t, "f6", &n); err == nil {
		t.Error("expected an error for null into int")
	}
}

// Values that cannot be converted to the destination type return errors instead of being ignored.
func TestUnmarshalUnsupportedDestinations(t *testing.T) {
	tests := []struct {
		hex string
		to  any
	}{
		{"43010203", new(int)},
		{"43010203", new(bool)},
		{"43010203", new(time.Time)},
		{"c074323031332d30332d32315432303a30343a30305a", new(int)},
		{"c11a514b67b0", new(bool)},
		{"c11a514b67b0", new([]byte)},
		{"d82076687474703a2f2f7777772e6578616d706c652e636f6d", new(int)},
		{"c249010000000000000000", new(time.Time)},
		{"f7", new(int)},
		{"d81a01", new(string)},
	}
	for _, test := range tests {
		if err := unmarshalHex(t, test.hex, test.to); !errors.Is(err, cbor.ErrUnmarshal) {
			t.Errorf("%s -> %T: got err=%v, want %v", test.hex, test.to, err, cbor.ErrUnmarshal)
		}
	}
}
