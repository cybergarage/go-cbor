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
	"net/url"
	"reflect"
	"testing"

	"github.com/cybergarage/go-cbor/cbor"
)

func textHex(s string) string {
	b, _ := cbor.Marshal(s)
	return hex.EncodeToString(b)
}

// RFC 8949 Section 3.4.5.1.
func TestEncodedCBORDataItemTag(t *testing.T) {
	valid := []struct {
		encoded  string
		expected any
	}{
		{"d818456449455446", cbor.Tag{Number: 24, Content: []byte("dIETF")}},
		{"d8184100", cbor.Tag{Number: 24, Content: []byte{0x00}}},
		{"d818439f01ff", cbor.Tag{Number: 24, Content: []byte{0x9f, 0x01, 0xff}}},
	}
	for _, test := range valid {
		v, err := decodeHex(t, test.encoded)
		if err != nil {
			t.Fatalf("%s: %v", test.encoded, err)
		}
		if !reflect.DeepEqual(v, test.expected) {
			t.Errorf("%s: got %#v, want %#v", test.encoded, v, test.expected)
		}
	}

	invalid := []string{
		"d81801",       // not a byte string
		"d81840",       // empty
		"d8184118",     // truncated item
		"d818420001",   // two items
		"d81841ff",     // break
		"d8184162c328", // truncated text
		"d81843d81800", // embedded tag 24 whose content is not a byte string
		"d8184362c328", // embedded text string with invalid UTF-8
	}
	for _, test := range invalid {
		if _, err := decodeHex(t, test); !errors.Is(err, cbor.ErrDecode) {
			t.Errorf("%s: got err=%v, want %v", test, err, cbor.ErrDecode)
		}
	}

	// The nesting depth of embedded data items counts toward the limit.
	embedded := func(levels int) string {
		h := "00"
		for range levels {
			b, _ := hex.DecodeString(h)
			bs, _ := cbor.Marshal(b)
			h = "d818" + hex.EncodeToString(bs)
		}
		return h
	}
	if _, err := decodeHex(t, embedded(cbor.DefaultMaxNestedLevels)); err != nil {
		t.Errorf("depth %d: %v", cbor.DefaultMaxNestedLevels, err)
	}
	if _, err := decodeHex(t, embedded(cbor.DefaultMaxNestedLevels+1)); !errors.Is(err, cbor.ErrDecode) {
		t.Errorf("depth %d: got err=%v, want %v", cbor.DefaultMaxNestedLevels+1, err, cbor.ErrDecode)
	}
}

// RFC 8949 Section 3.4.5.2: the expected conversions are not validated.
func TestExpectedLaterEncodingTags(t *testing.T) {
	tests := []struct {
		encoded  string
		expected any
	}{
		{"d54401020304", cbor.Tag{Number: 21, Content: []byte{1, 2, 3, 4}}},
		{"d6820141ff", cbor.Tag{Number: 22, Content: []any{int8(1), []byte{0xff}}}},
		{"d74401020304", cbor.Tag{Number: 23, Content: []byte{1, 2, 3, 4}}},
		{"d76161", cbor.Tag{Number: 23, Content: "a"}},
	}
	for _, test := range tests {
		v, err := decodeHex(t, test.encoded)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(v, test.expected) {
			t.Errorf("%s: got %#v, want %#v", test.encoded, v, test.expected)
		}
	}
}

// RFC 8949 Section 3.4.5.3: tag 32.
func TestURITag(t *testing.T) {
	valid := []string{
		"http://www.example.com",
		"https://user@example.com:8080/a/b?q=1&r=%2F#frag",
		"urn:ietf:rfc:8949",
		"mailto:info@example.com",
		"/relative/path?x",
		"../a",
		"#only-fragment",
		"",
		"http://[2001:db8::1]/",
	}
	for _, uri := range valid {
		v, err := decodeHex(t, "d820"+textHex(uri))
		if err != nil {
			t.Errorf("%q: %v", uri, err)
			continue
		}
		u, ok := v.(*url.URL)
		if !ok {
			t.Errorf("%q: got %T, want *url.URL", uri, v)
			continue
		}
		if u.String() != uri {
			t.Errorf("%q: got %q", uri, u.String())
		}
	}

	invalid := []string{
		"http://example.com/a b",
		"http://example.com/%zz",
		"http://example.com/%2",
		"http://example.com/a#b#c",
		"http://example.com/é",
		"http://example.com/<>",
		"http://example.com/\\",
		"http://[::1/",
	}
	for _, uri := range invalid {
		if _, err := decodeHex(t, "d820"+textHex(uri)); !errors.Is(err, cbor.ErrDecode) {
			t.Errorf("%q: got err=%v, want %v", uri, err, cbor.ErrDecode)
		}
	}
	if _, err := decodeHex(t, "d82001"); !errors.Is(err, cbor.ErrDecode) {
		t.Errorf("got err=%v, want %v", err, cbor.ErrDecode)
	}

	// Encoding
	u, _ := url.Parse("http://www.example.com")
	for _, v := range []any{u, *u} {
		b, err := cbor.Marshal(v)
		if err != nil {
			t.Fatal(err)
		}
		if got, want := hex.EncodeToString(b), "d82076687474703a2f2f7777772e6578616d706c652e636f6d"; got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}
	b, _ := cbor.Marshal((*url.URL)(nil))
	if got := hex.EncodeToString(b); got != "f6" {
		t.Errorf("got %s, want f6", got)
	}
}

// RFC 8949 Section 3.4.5.3: tags 33 and 34.
func TestBase64TextTags(t *testing.T) {
	tests := []struct {
		tag   string
		text  string
		valid bool
	}{
		// base64url without padding
		{"d821", "", true},
		{"d821", "AQID", true},
		{"d821", "AQI", true},
		{"d821", "AQ", true},
		{"d821", "-_8", true},
		{"d821", "AQ==", false}, // padding
		{"d821", "A", false},    // only 1 character in the last block
		{"d821", "AR", false},   // non-zero padding bits
		{"d821", "AQI+", false}, // base64 alphabet
		{"d821", "AQ\nID", false},
		// base64 with padding
		{"d822", "", true},
		{"d822", "AQID", true},
		{"d822", "AQI=", true},
		{"d822", "AQ==", true},
		{"d822", "+/8=", true},
		{"d822", "AQ", false},   // missing padding
		{"d822", "AQ=", false},  // wrong padding
		{"d822", "AR==", false}, // non-zero padding bits
		{"d822", "-_8=", false}, // base64url alphabet
		{"d822", "AQID\r\n", false},
	}
	for _, test := range tests {
		v, err := decodeHex(t, test.tag+textHex(test.text))
		if test.valid {
			want := cbor.Tag{Number: 33, Content: test.text}
			if test.tag == "d822" {
				want.Number = 34
			}
			if err != nil || !reflect.DeepEqual(v, want) {
				t.Errorf("%s %q: got %#v, %v", test.tag, test.text, v, err)
			}
			continue
		}
		if !errors.Is(err, cbor.ErrDecode) {
			t.Errorf("%s %q: got err=%v, want %v", test.tag, test.text, err, cbor.ErrDecode)
		}
	}
	for _, test := range []string{"d82101", "d82241ff"} {
		if _, err := decodeHex(t, test); !errors.Is(err, cbor.ErrDecode) {
			t.Errorf("%s: got err=%v, want %v", test, err, cbor.ErrDecode)
		}
	}
}

// RFC 8949 Section 3.4.5.3 and RFC 7049: tags 35 and 36 require text strings.
func TestRegexpAndMIMETags(t *testing.T) {
	for _, test := range []string{"d823" + textHex("^a+$"), "d824" + textHex("Content-Type: text/plain\r\n\r\nhello")} {
		v, err := decodeHex(t, test)
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := v.(cbor.Tag); !ok {
			t.Errorf("%s: got %#v", test, v)
		}
	}
	for _, test := range []string{"d82301", "d8244161"} {
		if _, err := decodeHex(t, test); !errors.Is(err, cbor.ErrDecode) {
			t.Errorf("%s: got err=%v, want %v", test, err, cbor.ErrDecode)
		}
	}
}
