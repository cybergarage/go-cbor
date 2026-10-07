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
	"reflect"
	"testing"

	"github.com/cybergarage/go-cbor/cbor"
)

type indefiniteStep func(enc *cbor.Encoder) error

func startBytes(enc *cbor.Encoder) error { return enc.StartIndefiniteByteString() }
func startText(enc *cbor.Encoder) error  { return enc.StartIndefiniteTextString() }
func startArray(enc *cbor.Encoder) error { return enc.StartIndefiniteArray() }
func startMap(enc *cbor.Encoder) error   { return enc.StartIndefiniteMap() }
func end(enc *cbor.Encoder) error        { return enc.EndIndefinite() }

func item(v any) indefiniteStep {
	return func(enc *cbor.Encoder) error { return enc.Encode(v) }
}

func runSteps(mode cbor.EncodeMode, steps ...indefiniteStep) (string, *cbor.Encoder, error) {
	var buf bytes.Buffer
	enc := cbor.NewEncoder(&buf)
	enc.SetEncodeMode(mode)
	for _, step := range steps {
		if err := step(enc); err != nil {
			return hex.EncodeToString(buf.Bytes()), enc, err
		}
	}
	return hex.EncodeToString(buf.Bytes()), enc, nil
}

// RFC 8949 Section 3.2 and the indefinite-length examples in Appendix A.
func TestIndefiniteLengthEncoding(t *testing.T) {
	tests := []struct {
		name     string
		steps    []indefiniteStep
		expected string
	}{
		{"(_ h'0102', h'030405')", []indefiniteStep{startBytes, item([]byte{1, 2}), item([]byte{3, 4, 5}), end}, "5f42010243030405ff"},
		{"(_ \"strea\", \"ming\")", []indefiniteStep{startText, item("strea"), item("ming"), end}, "7f657374726561646d696e67ff"},
		{"(_ )", []indefiniteStep{startBytes, end}, "5fff"},
		{"[_ ]", []indefiniteStep{startArray, end}, "9fff"},
		{"[_ 1, [2, 3], [_ 4, 5]]", []indefiniteStep{startArray, item(uint8(1)), item([]int{2, 3}), startArray, item(4), item(5), end, end}, "9f018202039f0405ffff"},
		{"[1, [2, 3], [_ 4, 5]] items", []indefiniteStep{startArray, item(1), item([]int{2, 3}), startArray, item(4), item(5), end, end}, "9f018202039f0405ffff"},
		{"[_ 1, 2, ..., 25]", func() []indefiniteStep {
			steps := []indefiniteStep{startArray}
			for n := 1; n <= 25; n++ {
				steps = append(steps, item(n))
			}
			return append(steps, end)
		}(), "9f0102030405060708090a0b0c0d0e0f101112131415161718181819ff"},
		{"{_ \"a\": 1, \"b\": [_ 2, 3]}", []indefiniteStep{startMap, item("a"), item(1), item("b"), startArray, item(2), item(3), end, end}, "bf61610161629f0203ffff"},
		{"{_ \"Fun\": true, \"Amt\": -2}", []indefiniteStep{startMap, item("Fun"), item(true), item("Amt"), item(-2), end}, "bf6346756ef563416d7421ff"},
		{"{_ [_ ]: (_ \"a\")}", []indefiniteStep{startMap, startArray, end, startText, item("a"), end, end}, "bf9fff7f6161ffff"},
		{"tag in array", []indefiniteStep{startArray, item(cbor.NewTag(23, []byte{1})), end}, "9fd74101ff"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, mode := range []cbor.EncodeMode{cbor.EncodeModePreferred, cbor.EncodeModeTypePreserving} {
				got, enc, err := runSteps(mode, test.steps...)
				if err != nil {
					t.Fatalf("%s: %v", mode, err)
				}
				if mode == cbor.EncodeModePreferred && got != test.expected {
					t.Errorf("got %s, want %s", got, test.expected)
				}
				if enc.IndefiniteDepth() != 0 {
					t.Errorf("IndefiniteDepth() = %d", enc.IndefiniteDepth())
				}
			}
		})
	}
}

// The indefinite-length encoding decodes to the same values as the definite-length encoding.
func TestIndefiniteLengthEncodingRoundTrip(t *testing.T) {
	got, _, err := runSteps(cbor.EncodeModePreferred,
		startMap,
		item("bytes"), startBytes, item([]byte("ab")), item([]byte{}), item([]byte("c")), end,
		item("text"), startText, item("é"), item("x"), end,
		item("list"), startArray, item(1), item(map[string]int{"k": 2}), end,
		end)
	if err != nil {
		t.Fatal(err)
	}
	v, err := decodeHex(t, got)
	if err != nil {
		t.Fatal(err)
	}
	want := map[any]any{
		"bytes": []byte("abc"),
		"text":  "éx",
		"list":  []any{int8(1), map[any]any{"k": int8(2)}},
	}
	if !reflect.DeepEqual(v, want) {
		t.Errorf("got %#v, want %#v", v, want)
	}
}

func TestIndefiniteLengthEncodingErrors(t *testing.T) {
	tests := []struct {
		name  string
		mode  cbor.EncodeMode
		steps []indefiniteStep
	}{
		{"end without start", cbor.EncodeModePreferred, []indefiniteStep{end}},
		{"text chunk in byte string", cbor.EncodeModePreferred, []indefiniteStep{startBytes, item("a")}},
		{"byte chunk in text string", cbor.EncodeModePreferred, []indefiniteStep{startText, item([]byte("a"))}},
		{"integer chunk", cbor.EncodeModePreferred, []indefiniteStep{startBytes, item(1)}},
		{"nested indefinite byte string", cbor.EncodeModePreferred, []indefiniteStep{startBytes, startBytes}},
		{"nested indefinite text string", cbor.EncodeModePreferred, []indefiniteStep{startText, startText}},
		{"array in string", cbor.EncodeModePreferred, []indefiniteStep{startText, startArray}},
		{"key without value", cbor.EncodeModePreferred, []indefiniteStep{startMap, item("a"), end}},
		{"invalid UTF-8 chunk", cbor.EncodeModePreferred, []indefiniteStep{startText, item("\xff")}},
		{"core deterministic", cbor.EncodeModeCoreDeterministic, []indefiniteStep{startArray}},
		{"length-first deterministic", cbor.EncodeModeLengthFirstDeterministic, []indefiniteStep{startMap}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if _, _, err := runSteps(test.mode, test.steps...); !errors.Is(err, cbor.ErrEncode) {
				t.Errorf("got err=%v, want %v", err, cbor.ErrEncode)
			}
		})
	}

	// A map whose end failed because of a missing value can be completed.
	var buf bytes.Buffer
	enc := cbor.NewEncoder(&buf)
	if err := enc.StartIndefiniteMap(); err != nil {
		t.Fatal(err)
	}
	if err := enc.Encode("a"); err != nil {
		t.Fatal(err)
	}
	if err := enc.EndIndefinite(); !errors.Is(err, cbor.ErrEncode) {
		t.Fatalf("got err=%v, want %v", err, cbor.ErrEncode)
	}
	if err := enc.Encode(1); err != nil {
		t.Fatal(err)
	}
	if err := enc.EndIndefinite(); err != nil {
		t.Fatal(err)
	}
	if got := hex.EncodeToString(buf.Bytes()); got != "bf616101ff" {
		t.Errorf("got %s, want bf616101ff", got)
	}
}
