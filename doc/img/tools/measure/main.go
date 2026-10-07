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

// measure decodes a representative CBOR data item for each row of the conversion tables,
// unmarshals it into each destination type, and prints the results as JSON for build_tables.py.
//
// Usage: go run ./doc/img/tools/measure > measured.json
package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"net/url"
	"os"
	"time"

	"github.com/cybergarage/go-cbor/cbor"
)

// row is a row of the decoding and unmarshaling tables.
type row struct {
	MT    string          `json:"mt"`
	Info  string          `json:"info"`
	Cond  string          `json:"cond"`
	Hex   string          `json:"hex"`
	To    string          `json:"to"`
	Dests map[string]bool `json:"dests"`
}

// dest is a destination type of Decoder::Unmarshal() and UnmarshalTo().
// A destination is supported when the value can be unmarshaled into any of the types.
type dest struct {
	name  string
	types []func() any
}

// rows lists a representative data item for each row. Conditions on integers are tested
// with values on the specified side of the boundary.
func rows() []row {
	return []row{
		{MT: "0 (unsigned integer)", Info: "24 or <24", Cond: ">  MaxInt8", Hex: "18c8"},
		{MT: "0 (unsigned integer)", Info: "24 or <24", Cond: "<= MaxInt8", Hex: "05"},
		{MT: "0 (unsigned integer)", Info: "25", Cond: ">  MaxInt16", Hex: "199c40"},
		{MT: "0 (unsigned integer)", Info: "25", Cond: "<= MaxInt16", Hex: "1903e8"},
		{MT: "0 (unsigned integer)", Info: "26", Cond: ">  MaxInt32", Hex: "1ab2d05e00"},
		{MT: "0 (unsigned integer)", Info: "26", Cond: "<= MaxInt32", Hex: "1a000f4240"},
		{MT: "0 (unsigned integer)", Info: "27", Cond: ">  MaxInt64", Hex: "1b8000000000000000"},
		{MT: "0 (unsigned integer)", Info: "27", Cond: "<= MaxInt64", Hex: "1b000000e8d4a51000"},
		{MT: "1 (negative integer)", Info: "24 or <24", Cond: ">= MinInt8", Hex: "3863"},
		{MT: "1 (negative integer)", Info: "24 or <24", Cond: "<  MinInt8", Hex: "38c7"},
		{MT: "1 (negative integer)", Info: "25", Cond: ">= MinInt16", Hex: "3903e7"},
		{MT: "1 (negative integer)", Info: "25", Cond: "<  MinInt16", Hex: "399c3f"},
		{MT: "1 (negative integer)", Info: "26", Cond: ">= MinInt32", Hex: "3a000f423f"},
		{MT: "1 (negative integer)", Info: "26", Cond: "<  MinInt32", Hex: "3ab2d05dff"},
		{MT: "1 (negative integer)", Info: "27", Cond: ">= MinInt64", Hex: "3b000000e8d4a50fff"},
		{MT: "1 (negative integer)", Info: "27", Cond: "<  MinInt64", Hex: "3b8000000000000000"},
		{MT: "2 (byte string)", Info: "-", Cond: "-", Hex: "43010203"},
		{MT: "3 (text string)", Info: "-", Cond: "-", Hex: "6131"},
		{MT: "4 (array)", Info: "-", Cond: "-", Hex: "8101"},
		{MT: "5 (map)", Info: "-", Cond: "-", Hex: "a1614101"},
		{MT: "6 (tag)", Info: "0 (Date/Time)", Cond: "-", Hex: "c074323031332d30332d32315432303a30343a30305a"},
		{MT: "6 (tag)", Info: "1 (Epoch)", Cond: "-", Hex: "c11a514b67b0"},
		{MT: "6 (tag)", Info: "2 (Bignum)", Cond: "-", Hex: "c249010000000000000000"},
		{MT: "6 (tag)", Info: "3 (Bignum)", Cond: "-", Hex: "c349010000000000000000"},
		{MT: "6 (tag)", Info: "32 (URI)", Cond: "-", Hex: "d82076687474703a2f2f7777772e6578616d706c652e636f6d"},
		{MT: "6 (tag)", Info: "55799 (CBOR)", Cond: "-", Hex: "d9d9f701", To: "(content)"},
		{MT: "6 (tag)", Info: "others", Cond: "-", Hex: "d81a01"},
		{MT: "7 (simple)", Info: "<20", Cond: "-", Hex: "f0"},
		{MT: "7 (simple)", Info: "20 (false)", Cond: "-", Hex: "f4"},
		{MT: "7 (simple)", Info: "21 (true)", Cond: "-", Hex: "f5"},
		{MT: "7 (simple)", Info: "22 (null)", Cond: "-", Hex: "f6"},
		{MT: "7 (simple)", Info: "23 (undefined)", Cond: "-", Hex: "f7"},
		{MT: "7 (simple)", Info: "24", Cond: "-", Hex: "f8ff"},
		{MT: "7 (floating-point)", Info: "25 (float16)", Cond: "-", Hex: "f93e00"},
		{MT: "7 (floating-point)", Info: "26 (float32)", Cond: "-", Hex: "fa47c35000"},
		{MT: "7 (floating-point)", Info: "27 (float64)", Cond: "-", Hex: "fb3ff199999999999a"},
	}
}

// dests lists the destination types in the order of the table columns.
func dests() []dest {
	return []dest{
		{"&string", []func() any{func() any { return new(string) }}},
		{"&int(bit)?", []func() any{
			func() any { return new(int8) }, func() any { return new(int16) }, func() any { return new(int32) },
			func() any { return new(int64) }, func() any { return new(int) },
		}},
		{"&uint(bit)?", []func() any{
			func() any { return new(uint8) }, func() any { return new(uint16) }, func() any { return new(uint32) },
			func() any { return new(uint64) }, func() any { return new(uint) },
		}},
		{"&float(bit)", []func() any{func() any { return new(float32) }, func() any { return new(float64) }}},
		{"&bool", []func() any{func() any { return new(bool) }}},
		{"&[]byte", []func() any{func() any { return new([]byte) }}},
		{"map[type]type", []func() any{func() any { return map[string]int{} }}},
		{"&[]type", []func() any{func() any { return &[]int{} }}},
		{"&struct", []func() any{func() any { return &struct{ A int }{A: 0} }}},
		{"&time.Time", []func() any{func() any { return &time.Time{} }}},
		{"&big.Int", []func() any{func() any { return new(big.Int) }}},
		{"&url.URL", []func() any{func() any { return &url.URL{} }}},
		{"&cbor.Tag", []func() any{func() any { return new(cbor.Tag) }}},
		{"&cbor.SimpleValue", []func() any{func() any { return new(cbor.SimpleValue) }}},
		{"&any", []func() any{func() any { return new(any) }}},
	}
}

// typeName returns the Go type name of a decoded value as shown in the tables.
func typeName(v any) string {
	switch x := v.(type) {
	case nil:
		return "nil"
	case []byte:
		return "[]byte"
	case []any:
		return "[]any"
	case map[any]any:
		return "map[any]any"
	case bool:
		return fmt.Sprintf("bool(%v)", x)
	case cbor.SimpleValue:
		if x == cbor.Undefined {
			return "cbor.Undefined"
		}
	}
	return fmt.Sprintf("%T", v)
}

// unmarshalTo returns true if the data item can be unmarshaled into the destination without errors or panics.
func unmarshalTo(b []byte, to any) bool {
	ok := false
	func() {
		defer func() {
			_ = recover()
		}()
		ok = cbor.UnmarshalTo(b, to) == nil
	}()
	return ok
}

func measure() ([]row, []string, error) {
	rs := rows()
	ds := dests()
	for n := range rs {
		b, err := hex.DecodeString(rs[n].Hex)
		if err != nil {
			return nil, nil, err
		}
		v, err := cbor.Unmarshal(b)
		if err != nil {
			return nil, nil, fmt.Errorf("%s: %w", rs[n].Hex, err)
		}
		if rs[n].To == "" {
			rs[n].To = typeName(v)
		}
		rs[n].Dests = map[string]bool{}
		for _, d := range ds {
			supported := false
			for _, newType := range d.types {
				if unmarshalTo(b, newType()) {
					supported = true
				}
			}
			rs[n].Dests[d.name] = supported
		}
	}
	names := make([]string, 0, len(ds))
	for _, d := range ds {
		names = append(names, d.name)
	}
	return rs, names, nil
}

func main() {
	rs, names, err := measure()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", " ")
	if err := enc.Encode(map[string]any{"rows": rs, "dests": names}); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
