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
	"math"
	"testing"

	"github.com/cybergarage/go-cbor/cbor"
)

func unmarshalProfile(v any) error {
	b, err := cbor.Marshal(v)
	if err != nil {
		return err
	}
	r, err := cbor.Unmarshal(b)
	if err != nil {
		return err
	}

	err = deepEqual(v, r)
	if err != nil {
		return err
	}
	return nil
}

func unmarshalToProfile(v any, to any) error {
	b, err := cbor.Marshal(v)
	if err != nil {
		return err
	}
	err = cbor.UnmarshalTo(b, to)
	if err != nil {
		return err
	}

	err = deepEqual(v, to)
	if err != nil {
		return err
	}
	return nil
}

func BenchmarkData(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var v int
		v = int(0)
		unmarshalProfile(v)
		v = int(math.MinInt)
		unmarshalProfile(v)
		v = int(math.MaxInt)
		unmarshalProfile(v)
	}
	for n := 0; n < b.N; n++ {
		var v int8
		v = int8(0)
		unmarshalProfile(v)
		v = int8(math.MinInt8)
		unmarshalProfile(v)
		v = int8(math.MaxInt8)
		unmarshalProfile(v)
	}
	for n := 0; n < b.N; n++ {
		var v int16
		v = int16(0)
		unmarshalProfile(v)
		v = int16(math.MinInt16)
		unmarshalProfile(v)
		v = int16(math.MaxInt16)
		unmarshalProfile(v)
	}
	for n := 0; n < b.N; n++ {
		var v int32
		v = int32(0)
		unmarshalProfile(v)
		v = int32(math.MinInt32)
		unmarshalProfile(v)
		v = int32(math.MaxInt32)
		unmarshalProfile(v)
	}
	for n := 0; n < b.N; n++ {
		var v int64
		v = int64(0)
		unmarshalProfile(v)
		v = int64(math.MinInt64)
		unmarshalProfile(v)
		v = int64(math.MaxInt64)
		unmarshalProfile(v)
	}
	for n := 0; n < b.N; n++ {
		var v uint
		v = uint(0)
		unmarshalProfile(v)
		v = uint(math.MaxUint)
		unmarshalProfile(v)
	}
	for n := 0; n < b.N; n++ {
		var v uint8
		v = uint8(0)
		unmarshalProfile(v)
		v = uint8(math.MaxUint8)
		unmarshalProfile(v)
	}
	for n := 0; n < b.N; n++ {
		var v uint16
		v = uint16(0)
		unmarshalProfile(v)
		v = uint16(math.MaxUint16)
		unmarshalProfile(v)
	}
	for n := 0; n < b.N; n++ {
		var v uint32
		v = uint32(0)
		unmarshalProfile(v)
		v = uint32(math.MaxUint32)
		unmarshalProfile(v)
	}
	for n := 0; n < b.N; n++ {
		var v uint64
		v = uint64(0)
		unmarshalProfile(v)
		v = uint64(math.MaxInt64)
		unmarshalProfile(v)
	}
	for n := 0; n < b.N; n++ {
		var v float32
		v = float32(-math.MaxFloat32)
		unmarshalProfile(v)
		v = float32(0)
		unmarshalProfile(v)
		v = float32(math.MaxFloat32)
		unmarshalProfile(v)
	}
	for n := 0; n < b.N; n++ {
		var v float64
		v = float64(-math.MaxFloat32)
		unmarshalProfile(v)
		v = float64(0)
		unmarshalProfile(v)
		v = float64(math.MaxFloat64)
		unmarshalProfile(v)
	}
	for n := 0; n < b.N; n++ {
		var v bool
		v = bool(true)
		unmarshalProfile(v)
		v = bool(false)
		unmarshalProfile(v)
	}
	for n := 0; n < b.N; n++ {
		var v string
		v = string("a")
		unmarshalProfile(v)
		v = string("ab")
		unmarshalProfile(v)
		v = string("abc")
		unmarshalProfile(v)
	}
	for n := 0; n < b.N; n++ {
		var v []byte
		v = []byte("x")
		unmarshalProfile(v)
		v = []byte("xy")
		unmarshalProfile(v)
		v = []byte("xyz")
		unmarshalProfile(v)
	}
}
