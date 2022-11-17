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
	"math"
	"testing"

	"github.com/cybergarage/go-cbor/cbor"
)

func fuzzTest(t *testing.T, v any) {
	t.Helper()
	b, err := cbor.Marshal(v)
	if err != nil {
		t.Errorf("Marshal(%v) : %s", v, err)
		return
	}
	r, err := cbor.Unmarshal(b)
	if err != nil {
		t.Errorf("Unmarshal(%v => %s) : %s", v, hex.EncodeToString(b), err)
		return
	}

	err = DeepEqual(v, r)
	if err != nil {
		t.Error(err)
		return
	}
}

func fuzzPrimitiveTest[T comparable](t *testing.T, v T) {
	t.Helper()
	fuzzTest(t, v)
}

func FuzzInt(f *testing.F) {
	f.Add(int(0))
	f.Add(int(math.MinInt))
	f.Add(int(math.MaxInt))
	f.Fuzz(func(t *testing.T, v int) {
		fuzzPrimitiveTest(t, v)
	})
}

func FuzzInt8(f *testing.F) {
	f.Add(int8(0))
	f.Add(int8(math.MinInt8))
	f.Add(int8(math.MaxInt8))
	f.Fuzz(func(t *testing.T, v int8) {
		fuzzPrimitiveTest(t, v)
	})
}

func FuzzInt16(f *testing.F) {
	f.Add(int16(0))
	f.Add(int16(math.MinInt16))
	f.Add(int16(math.MaxInt16))
	f.Fuzz(func(t *testing.T, v int16) {
		fuzzPrimitiveTest(t, v)
	})
}

func FuzzInt32(f *testing.F) {
	f.Add(int32(0))
	f.Add(int32(math.MinInt32))
	f.Add(int32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, v int32) {
		fuzzPrimitiveTest(t, v)
	})
}

func FuzzInt64(f *testing.F) {
	f.Add(int64(0))
	f.Add(int64(math.MinInt64))
	f.Add(int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, v int64) {
		fuzzPrimitiveTest(t, v)
	})
}

func FuzzUint(f *testing.F) {
	f.Add(uint(0))
	f.Add(uint(math.MaxUint))
	f.Fuzz(func(t *testing.T, v uint) {
		fuzzPrimitiveTest(t, v)
	})
}

func FuzzUint8(f *testing.F) {
	f.Add(uint8(0))
	f.Add(uint8(math.MaxUint8))
	f.Fuzz(func(t *testing.T, v uint8) {
		fuzzPrimitiveTest(t, v)
	})
}

func FuzzUint16(f *testing.F) {
	f.Add(uint16(0))
	f.Add(uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, v uint16) {
		fuzzPrimitiveTest(t, v)
	})
}

func FuzzUint32(f *testing.F) {
	f.Add(uint32(0))
	f.Add(uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, v uint32) {
		fuzzPrimitiveTest(t, v)
	})
}

func FuzzUint64(f *testing.F) {
	f.Add(uint64(0))
	f.Add(uint64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, v uint64) {
		fuzzPrimitiveTest(t, v)
	})
}

func FuzzFloat32(f *testing.F) {
	f.Add(float32(math.MaxFloat32))
	f.Fuzz(func(t *testing.T, v float32) {
		fuzzPrimitiveTest(t, v)
	})
}

func FuzzFloat64(f *testing.F) {
	f.Add(float64(math.MaxFloat64))
	f.Fuzz(func(t *testing.T, v float64) {
		fuzzPrimitiveTest(t, v)
	})
}

func FuzzBool(f *testing.F) {
	f.Add(bool(true))
	f.Add(bool(false))
	f.Fuzz(func(t *testing.T, v bool) {
		fuzzPrimitiveTest(t, v)
	})
}

func FuzzString(f *testing.F) {
	f.Add(string("abc"))
	f.Add(string("xyz"))
	f.Fuzz(func(t *testing.T, v string) {
		fuzzPrimitiveTest(t, v)
	})
}

func FuzzIntArray(f *testing.F) {
	f.Add(int(0))
	f.Add(int(math.MinInt))
	f.Add(int(math.MaxInt))
	f.Fuzz(func(t *testing.T, v int) {
		va := []int{v, v, v, v, v}
		fuzzTest(t, va)
	})
}

func FuzzInt8Array(f *testing.F) {
	f.Add(int8(0))
	f.Add(int8(math.MinInt8))
	f.Add(int8(math.MaxInt8))
	f.Fuzz(func(t *testing.T, v int8) {
		va := []int8{v, v, v, v, v}
		fuzzTest(t, va)
	})
}

func FuzzInt16Array(f *testing.F) {
	f.Add(int16(0))
	f.Add(int16(math.MinInt16))
	f.Add(int16(math.MaxInt16))
	f.Fuzz(func(t *testing.T, v int16) {
		va := []int16{v, v, v, v, v}
		fuzzTest(t, va)
	})
}

func FuzzInt32Array(f *testing.F) {
	f.Add(int32(0))
	f.Add(int32(math.MinInt32))
	f.Add(int32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, v int32) {
		va := []int32{v, v, v, v, v}
		fuzzTest(t, va)
	})
}

func FuzzInt64Array(f *testing.F) {
	f.Add(int64(0))
	f.Add(int64(math.MinInt64))
	f.Add(int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, v int64) {
		va := []int64{v, v, v, v, v}
		fuzzTest(t, va)
	})
}

func FuzzUintArray(f *testing.F) {
	f.Add(uint(0))
	f.Add(uint(math.MaxUint))
	f.Fuzz(func(t *testing.T, v uint) {
		va := []uint{v, v, v, v, v}
		fuzzTest(t, va)
	})
}

func FuzzUint8Array(f *testing.F) {
	f.Add(uint8(0))
	f.Add(uint8(math.MaxUint8))
	f.Fuzz(func(t *testing.T, v uint8) {
		va := []uint8{v, v, v, v, v}
		fuzzTest(t, va)
	})
}

func FuzzUint16Array(f *testing.F) {
	f.Add(uint16(0))
	f.Add(uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, v uint16) {
		va := []uint16{v, v, v, v, v}
		fuzzTest(t, va)
	})
}

func FuzzUint32Array(f *testing.F) {
	f.Add(uint32(0))
	f.Add(uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, v uint32) {
		va := []uint32{v, v, v, v, v}
		fuzzTest(t, va)
	})
}

func FuzzUint64Array(f *testing.F) {
	f.Add(uint64(0))
	f.Add(uint64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, v uint64) {
		va := []uint64{v, v, v, v, v}
		fuzzTest(t, va)
	})
}

func FuzzFloat32Array(f *testing.F) {
	f.Add(float32(math.MaxFloat32))
	f.Fuzz(func(t *testing.T, v float32) {
		va := []float32{v, v, v, v, v}
		fuzzTest(t, va)
	})
}

func FuzzFloat64Array(f *testing.F) {
	f.Add(float64(math.MaxFloat64))
	f.Fuzz(func(t *testing.T, v float64) {
		va := []float64{v, v, v, v, v}
		fuzzTest(t, va)
	})
}

func FuzzBoolArray(f *testing.F) {
	f.Add(bool(true))
	f.Add(bool(false))
	f.Fuzz(func(t *testing.T, v bool) {
		va := []bool{v, v, v, v, v}
		fuzzTest(t, va)
	})
}

func FuzzStringArray(f *testing.F) {
	f.Add(string("abc"))
	f.Add(string("xyz"))
	f.Fuzz(func(t *testing.T, v string) {
		va := []string{v, v, v, v, v}
		fuzzTest(t, va)
	})
}

// nolint: dupl
func FuzzIntIntMap(f *testing.F) {
	f.Add(int(0), int(0))
	f.Add(int(0), int(math.MinInt))
	f.Add(int(0), int(math.MaxInt))
	f.Add(int(math.MinInt), int(0))
	f.Add(int(math.MinInt), int(math.MinInt))
	f.Add(int(math.MinInt), int(math.MaxInt))
	f.Add(int(math.MaxInt), int(0))
	f.Add(int(math.MaxInt), int(math.MinInt))
	f.Add(int(math.MaxInt), int(math.MaxInt))
	f.Fuzz(func(t *testing.T, k int, v int) {
		vm := map[int]int{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzIntInt8Map(f *testing.F) {
	f.Add(int(0), int8(0))
	f.Add(int(0), int8(math.MinInt8))
	f.Add(int(0), int8(math.MaxInt8))
	f.Add(int(math.MinInt), int8(0))
	f.Add(int(math.MinInt), int8(math.MinInt8))
	f.Add(int(math.MinInt), int8(math.MaxInt8))
	f.Add(int(math.MaxInt), int8(0))
	f.Add(int(math.MaxInt), int8(math.MinInt8))
	f.Add(int(math.MaxInt), int8(math.MaxInt8))
	f.Fuzz(func(t *testing.T, k int, v int8) {
		vm := map[int]int8{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzIntInt16Map(f *testing.F) {
	f.Add(int(0), int16(0))
	f.Add(int(0), int16(math.MinInt16))
	f.Add(int(0), int16(math.MaxInt16))
	f.Add(int(math.MinInt), int16(0))
	f.Add(int(math.MinInt), int16(math.MinInt16))
	f.Add(int(math.MinInt), int16(math.MaxInt16))
	f.Add(int(math.MaxInt), int16(0))
	f.Add(int(math.MaxInt), int16(math.MinInt16))
	f.Add(int(math.MaxInt), int16(math.MaxInt16))
	f.Fuzz(func(t *testing.T, k int, v int16) {
		vm := map[int]int16{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzIntInt32Map(f *testing.F) {
	f.Add(int(0), int32(0))
	f.Add(int(0), int32(math.MinInt32))
	f.Add(int(0), int32(math.MaxInt32))
	f.Add(int(math.MinInt), int32(0))
	f.Add(int(math.MinInt), int32(math.MinInt32))
	f.Add(int(math.MinInt), int32(math.MaxInt32))
	f.Add(int(math.MaxInt), int32(0))
	f.Add(int(math.MaxInt), int32(math.MinInt32))
	f.Add(int(math.MaxInt), int32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, k int, v int32) {
		vm := map[int]int32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzIntInt64Map(f *testing.F) {
	f.Add(int(0), int64(0))
	f.Add(int(0), int64(math.MinInt64))
	f.Add(int(0), int64(math.MaxInt64))
	f.Add(int(math.MinInt), int64(0))
	f.Add(int(math.MinInt), int64(math.MinInt64))
	f.Add(int(math.MinInt), int64(math.MaxInt64))
	f.Add(int(math.MaxInt), int64(0))
	f.Add(int(math.MaxInt), int64(math.MinInt64))
	f.Add(int(math.MaxInt), int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k int, v int64) {
		vm := map[int]int64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzIntUintMap(f *testing.F) {
	f.Add(int(0), uint(0))
	f.Add(int(0), uint(math.MaxUint))
	f.Add(int(math.MinInt), uint(0))
	f.Add(int(math.MinInt), uint(math.MaxUint))
	f.Add(int(math.MaxInt), uint(0))
	f.Add(int(math.MaxInt), uint(math.MaxUint))
	f.Fuzz(func(t *testing.T, k int, v uint) {
		vm := map[int]uint{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzIntUint8Map(f *testing.F) {
	f.Add(int(0), uint8(0))
	f.Add(int(0), uint8(math.MaxUint8))
	f.Add(int(math.MinInt), uint8(0))
	f.Add(int(math.MinInt), uint8(math.MaxUint8))
	f.Add(int(math.MaxInt), uint8(0))
	f.Add(int(math.MaxInt), uint8(math.MaxUint8))
	f.Fuzz(func(t *testing.T, k int, v uint8) {
		vm := map[int]uint8{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzIntUint16Map(f *testing.F) {
	f.Add(int(0), uint16(0))
	f.Add(int(0), uint16(math.MaxUint16))
	f.Add(int(math.MinInt), uint16(0))
	f.Add(int(math.MinInt), uint16(math.MaxUint16))
	f.Add(int(math.MaxInt), uint16(0))
	f.Add(int(math.MaxInt), uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, k int, v uint16) {
		vm := map[int]uint16{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzIntUint32Map(f *testing.F) {
	f.Add(int(0), uint32(0))
	f.Add(int(0), uint32(math.MaxUint32))
	f.Add(int(math.MinInt), uint32(0))
	f.Add(int(math.MinInt), uint32(math.MaxUint32))
	f.Add(int(math.MaxInt), uint32(0))
	f.Add(int(math.MaxInt), uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, k int, v uint32) {
		vm := map[int]uint32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzIntUint64Map(f *testing.F) {
	f.Add(int(0), uint64(0))
	f.Add(int(0), uint64(math.MaxInt64))
	f.Add(int(math.MinInt), uint64(0))
	f.Add(int(math.MinInt), uint64(math.MaxInt64))
	f.Add(int(math.MaxInt), uint64(0))
	f.Add(int(math.MaxInt), uint64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k int, v uint64) {
		vm := map[int]uint64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzIntFloat32Map(f *testing.F) {
	f.Add(int(0), float32(math.MaxFloat32))
	f.Add(int(math.MinInt), float32(math.MaxFloat32))
	f.Add(int(math.MaxInt), float32(math.MaxFloat32))
	f.Fuzz(func(t *testing.T, k int, v float32) {
		vm := map[int]float32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzIntFloat64Map(f *testing.F) {
	f.Add(int(0), float64(math.MaxFloat64))
	f.Add(int(math.MinInt), float64(math.MaxFloat64))
	f.Add(int(math.MaxInt), float64(math.MaxFloat64))
	f.Fuzz(func(t *testing.T, k int, v float64) {
		vm := map[int]float64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzIntBoolMap(f *testing.F) {
	f.Add(int(0), bool(true))
	f.Add(int(0), bool(false))
	f.Add(int(math.MinInt), bool(true))
	f.Add(int(math.MinInt), bool(false))
	f.Add(int(math.MaxInt), bool(true))
	f.Add(int(math.MaxInt), bool(false))
	f.Fuzz(func(t *testing.T, k int, v bool) {
		vm := map[int]bool{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzIntStringMap(f *testing.F) {
	f.Add(int(0), string("abc"))
	f.Add(int(0), string("xyz"))
	f.Add(int(math.MinInt), string("abc"))
	f.Add(int(math.MinInt), string("xyz"))
	f.Add(int(math.MaxInt), string("abc"))
	f.Add(int(math.MaxInt), string("xyz"))
	f.Fuzz(func(t *testing.T, k int, v string) {
		vm := map[int]string{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt8IntMap(f *testing.F) {
	f.Add(int8(0), int(0))
	f.Add(int8(0), int(math.MinInt))
	f.Add(int8(0), int(math.MaxInt))
	f.Add(int8(math.MinInt8), int(0))
	f.Add(int8(math.MinInt8), int(math.MinInt))
	f.Add(int8(math.MinInt8), int(math.MaxInt))
	f.Add(int8(math.MaxInt8), int(0))
	f.Add(int8(math.MaxInt8), int(math.MinInt))
	f.Add(int8(math.MaxInt8), int(math.MaxInt))
	f.Fuzz(func(t *testing.T, k int8, v int) {
		vm := map[int8]int{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt8Int8Map(f *testing.F) {
	f.Add(int8(0), int8(0))
	f.Add(int8(0), int8(math.MinInt8))
	f.Add(int8(0), int8(math.MaxInt8))
	f.Add(int8(math.MinInt8), int8(0))
	f.Add(int8(math.MinInt8), int8(math.MinInt8))
	f.Add(int8(math.MinInt8), int8(math.MaxInt8))
	f.Add(int8(math.MaxInt8), int8(0))
	f.Add(int8(math.MaxInt8), int8(math.MinInt8))
	f.Add(int8(math.MaxInt8), int8(math.MaxInt8))
	f.Fuzz(func(t *testing.T, k int8, v int8) {
		vm := map[int8]int8{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt8Int16Map(f *testing.F) {
	f.Add(int8(0), int16(0))
	f.Add(int8(0), int16(math.MinInt16))
	f.Add(int8(0), int16(math.MaxInt16))
	f.Add(int8(math.MinInt8), int16(0))
	f.Add(int8(math.MinInt8), int16(math.MinInt16))
	f.Add(int8(math.MinInt8), int16(math.MaxInt16))
	f.Add(int8(math.MaxInt8), int16(0))
	f.Add(int8(math.MaxInt8), int16(math.MinInt16))
	f.Add(int8(math.MaxInt8), int16(math.MaxInt16))
	f.Fuzz(func(t *testing.T, k int8, v int16) {
		vm := map[int8]int16{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt8Int32Map(f *testing.F) {
	f.Add(int8(0), int32(0))
	f.Add(int8(0), int32(math.MinInt32))
	f.Add(int8(0), int32(math.MaxInt32))
	f.Add(int8(math.MinInt8), int32(0))
	f.Add(int8(math.MinInt8), int32(math.MinInt32))
	f.Add(int8(math.MinInt8), int32(math.MaxInt32))
	f.Add(int8(math.MaxInt8), int32(0))
	f.Add(int8(math.MaxInt8), int32(math.MinInt32))
	f.Add(int8(math.MaxInt8), int32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, k int8, v int32) {
		vm := map[int8]int32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt8Int64Map(f *testing.F) {
	f.Add(int8(0), int64(0))
	f.Add(int8(0), int64(math.MinInt64))
	f.Add(int8(0), int64(math.MaxInt64))
	f.Add(int8(math.MinInt8), int64(0))
	f.Add(int8(math.MinInt8), int64(math.MinInt64))
	f.Add(int8(math.MinInt8), int64(math.MaxInt64))
	f.Add(int8(math.MaxInt8), int64(0))
	f.Add(int8(math.MaxInt8), int64(math.MinInt64))
	f.Add(int8(math.MaxInt8), int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k int8, v int64) {
		vm := map[int8]int64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt8UintMap(f *testing.F) {
	f.Add(int8(0), uint(0))
	f.Add(int8(0), uint(math.MaxUint))
	f.Add(int8(math.MinInt8), uint(0))
	f.Add(int8(math.MinInt8), uint(math.MaxUint))
	f.Add(int8(math.MaxInt8), uint(0))
	f.Add(int8(math.MaxInt8), uint(math.MaxUint))
	f.Fuzz(func(t *testing.T, k int8, v uint) {
		vm := map[int8]uint{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt8Uint8Map(f *testing.F) {
	f.Add(int8(0), uint8(0))
	f.Add(int8(0), uint8(math.MaxUint8))
	f.Add(int8(math.MinInt8), uint8(0))
	f.Add(int8(math.MinInt8), uint8(math.MaxUint8))
	f.Add(int8(math.MaxInt8), uint8(0))
	f.Add(int8(math.MaxInt8), uint8(math.MaxUint8))
	f.Fuzz(func(t *testing.T, k int8, v uint8) {
		vm := map[int8]uint8{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt8Uint16Map(f *testing.F) {
	f.Add(int8(0), uint16(0))
	f.Add(int8(0), uint16(math.MaxUint16))
	f.Add(int8(math.MinInt8), uint16(0))
	f.Add(int8(math.MinInt8), uint16(math.MaxUint16))
	f.Add(int8(math.MaxInt8), uint16(0))
	f.Add(int8(math.MaxInt8), uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, k int8, v uint16) {
		vm := map[int8]uint16{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt8Uint32Map(f *testing.F) {
	f.Add(int8(0), uint32(0))
	f.Add(int8(0), uint32(math.MaxUint32))
	f.Add(int8(math.MinInt8), uint32(0))
	f.Add(int8(math.MinInt8), uint32(math.MaxUint32))
	f.Add(int8(math.MaxInt8), uint32(0))
	f.Add(int8(math.MaxInt8), uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, k int8, v uint32) {
		vm := map[int8]uint32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt8Uint64Map(f *testing.F) {
	f.Add(int8(0), uint64(0))
	f.Add(int8(0), uint64(math.MaxInt64))
	f.Add(int8(math.MinInt8), uint64(0))
	f.Add(int8(math.MinInt8), uint64(math.MaxInt64))
	f.Add(int8(math.MaxInt8), uint64(0))
	f.Add(int8(math.MaxInt8), uint64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k int8, v uint64) {
		vm := map[int8]uint64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt8Float32Map(f *testing.F) {
	f.Add(int8(0), float32(math.MaxFloat32))
	f.Add(int8(math.MinInt8), float32(math.MaxFloat32))
	f.Add(int8(math.MaxInt8), float32(math.MaxFloat32))
	f.Fuzz(func(t *testing.T, k int8, v float32) {
		vm := map[int8]float32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt8Float64Map(f *testing.F) {
	f.Add(int8(0), float64(math.MaxFloat64))
	f.Add(int8(math.MinInt8), float64(math.MaxFloat64))
	f.Add(int8(math.MaxInt8), float64(math.MaxFloat64))
	f.Fuzz(func(t *testing.T, k int8, v float64) {
		vm := map[int8]float64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt8BoolMap(f *testing.F) {
	f.Add(int8(0), bool(true))
	f.Add(int8(0), bool(false))
	f.Add(int8(math.MinInt8), bool(true))
	f.Add(int8(math.MinInt8), bool(false))
	f.Add(int8(math.MaxInt8), bool(true))
	f.Add(int8(math.MaxInt8), bool(false))
	f.Fuzz(func(t *testing.T, k int8, v bool) {
		vm := map[int8]bool{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt8StringMap(f *testing.F) {
	f.Add(int8(0), string("abc"))
	f.Add(int8(0), string("xyz"))
	f.Add(int8(math.MinInt8), string("abc"))
	f.Add(int8(math.MinInt8), string("xyz"))
	f.Add(int8(math.MaxInt8), string("abc"))
	f.Add(int8(math.MaxInt8), string("xyz"))
	f.Fuzz(func(t *testing.T, k int8, v string) {
		vm := map[int8]string{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt16IntMap(f *testing.F) {
	f.Add(int16(0), int(0))
	f.Add(int16(0), int(math.MinInt))
	f.Add(int16(0), int(math.MaxInt))
	f.Add(int16(math.MinInt16), int(0))
	f.Add(int16(math.MinInt16), int(math.MinInt))
	f.Add(int16(math.MinInt16), int(math.MaxInt))
	f.Add(int16(math.MaxInt16), int(0))
	f.Add(int16(math.MaxInt16), int(math.MinInt))
	f.Add(int16(math.MaxInt16), int(math.MaxInt))
	f.Fuzz(func(t *testing.T, k int16, v int) {
		vm := map[int16]int{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt16Int8Map(f *testing.F) {
	f.Add(int16(0), int8(0))
	f.Add(int16(0), int8(math.MinInt8))
	f.Add(int16(0), int8(math.MaxInt8))
	f.Add(int16(math.MinInt16), int8(0))
	f.Add(int16(math.MinInt16), int8(math.MinInt8))
	f.Add(int16(math.MinInt16), int8(math.MaxInt8))
	f.Add(int16(math.MaxInt16), int8(0))
	f.Add(int16(math.MaxInt16), int8(math.MinInt8))
	f.Add(int16(math.MaxInt16), int8(math.MaxInt8))
	f.Fuzz(func(t *testing.T, k int16, v int8) {
		vm := map[int16]int8{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt16Int16Map(f *testing.F) {
	f.Add(int16(0), int16(0))
	f.Add(int16(0), int16(math.MinInt16))
	f.Add(int16(0), int16(math.MaxInt16))
	f.Add(int16(math.MinInt16), int16(0))
	f.Add(int16(math.MinInt16), int16(math.MinInt16))
	f.Add(int16(math.MinInt16), int16(math.MaxInt16))
	f.Add(int16(math.MaxInt16), int16(0))
	f.Add(int16(math.MaxInt16), int16(math.MinInt16))
	f.Add(int16(math.MaxInt16), int16(math.MaxInt16))
	f.Fuzz(func(t *testing.T, k int16, v int16) {
		vm := map[int16]int16{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt16Int32Map(f *testing.F) {
	f.Add(int16(0), int32(0))
	f.Add(int16(0), int32(math.MinInt32))
	f.Add(int16(0), int32(math.MaxInt32))
	f.Add(int16(math.MinInt16), int32(0))
	f.Add(int16(math.MinInt16), int32(math.MinInt32))
	f.Add(int16(math.MinInt16), int32(math.MaxInt32))
	f.Add(int16(math.MaxInt16), int32(0))
	f.Add(int16(math.MaxInt16), int32(math.MinInt32))
	f.Add(int16(math.MaxInt16), int32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, k int16, v int32) {
		vm := map[int16]int32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt16Int64Map(f *testing.F) {
	f.Add(int16(0), int64(0))
	f.Add(int16(0), int64(math.MinInt64))
	f.Add(int16(0), int64(math.MaxInt64))
	f.Add(int16(math.MinInt16), int64(0))
	f.Add(int16(math.MinInt16), int64(math.MinInt64))
	f.Add(int16(math.MinInt16), int64(math.MaxInt64))
	f.Add(int16(math.MaxInt16), int64(0))
	f.Add(int16(math.MaxInt16), int64(math.MinInt64))
	f.Add(int16(math.MaxInt16), int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k int16, v int64) {
		vm := map[int16]int64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt16UintMap(f *testing.F) {
	f.Add(int16(0), uint(0))
	f.Add(int16(0), uint(math.MaxUint))
	f.Add(int16(math.MinInt16), uint(0))
	f.Add(int16(math.MinInt16), uint(math.MaxUint))
	f.Add(int16(math.MaxInt16), uint(0))
	f.Add(int16(math.MaxInt16), uint(math.MaxUint))
	f.Fuzz(func(t *testing.T, k int16, v uint) {
		vm := map[int16]uint{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt16Uint8Map(f *testing.F) {
	f.Add(int16(0), uint8(0))
	f.Add(int16(0), uint8(math.MaxUint8))
	f.Add(int16(math.MinInt16), uint8(0))
	f.Add(int16(math.MinInt16), uint8(math.MaxUint8))
	f.Add(int16(math.MaxInt16), uint8(0))
	f.Add(int16(math.MaxInt16), uint8(math.MaxUint8))
	f.Fuzz(func(t *testing.T, k int16, v uint8) {
		vm := map[int16]uint8{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt16Uint16Map(f *testing.F) {
	f.Add(int16(0), uint16(0))
	f.Add(int16(0), uint16(math.MaxUint16))
	f.Add(int16(math.MinInt16), uint16(0))
	f.Add(int16(math.MinInt16), uint16(math.MaxUint16))
	f.Add(int16(math.MaxInt16), uint16(0))
	f.Add(int16(math.MaxInt16), uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, k int16, v uint16) {
		vm := map[int16]uint16{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt16Uint32Map(f *testing.F) {
	f.Add(int16(0), uint32(0))
	f.Add(int16(0), uint32(math.MaxUint32))
	f.Add(int16(math.MinInt16), uint32(0))
	f.Add(int16(math.MinInt16), uint32(math.MaxUint32))
	f.Add(int16(math.MaxInt16), uint32(0))
	f.Add(int16(math.MaxInt16), uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, k int16, v uint32) {
		vm := map[int16]uint32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt16Uint64Map(f *testing.F) {
	f.Add(int16(0), uint64(0))
	f.Add(int16(0), uint64(math.MaxInt64))
	f.Add(int16(math.MinInt16), uint64(0))
	f.Add(int16(math.MinInt16), uint64(math.MaxInt64))
	f.Add(int16(math.MaxInt16), uint64(0))
	f.Add(int16(math.MaxInt16), uint64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k int16, v uint64) {
		vm := map[int16]uint64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt16Float32Map(f *testing.F) {
	f.Add(int16(0), float32(math.MaxFloat32))
	f.Add(int16(math.MinInt16), float32(math.MaxFloat32))
	f.Add(int16(math.MaxInt16), float32(math.MaxFloat32))
	f.Fuzz(func(t *testing.T, k int16, v float32) {
		vm := map[int16]float32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt16Float64Map(f *testing.F) {
	f.Add(int16(0), float64(math.MaxFloat64))
	f.Add(int16(math.MinInt16), float64(math.MaxFloat64))
	f.Add(int16(math.MaxInt16), float64(math.MaxFloat64))
	f.Fuzz(func(t *testing.T, k int16, v float64) {
		vm := map[int16]float64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt16BoolMap(f *testing.F) {
	f.Add(int16(0), bool(true))
	f.Add(int16(0), bool(false))
	f.Add(int16(math.MinInt16), bool(true))
	f.Add(int16(math.MinInt16), bool(false))
	f.Add(int16(math.MaxInt16), bool(true))
	f.Add(int16(math.MaxInt16), bool(false))
	f.Fuzz(func(t *testing.T, k int16, v bool) {
		vm := map[int16]bool{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt16StringMap(f *testing.F) {
	f.Add(int16(0), string("abc"))
	f.Add(int16(0), string("xyz"))
	f.Add(int16(math.MinInt16), string("abc"))
	f.Add(int16(math.MinInt16), string("xyz"))
	f.Add(int16(math.MaxInt16), string("abc"))
	f.Add(int16(math.MaxInt16), string("xyz"))
	f.Fuzz(func(t *testing.T, k int16, v string) {
		vm := map[int16]string{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt32IntMap(f *testing.F) {
	f.Add(int32(0), int(0))
	f.Add(int32(0), int(math.MinInt))
	f.Add(int32(0), int(math.MaxInt))
	f.Add(int32(math.MinInt32), int(0))
	f.Add(int32(math.MinInt32), int(math.MinInt))
	f.Add(int32(math.MinInt32), int(math.MaxInt))
	f.Add(int32(math.MaxInt32), int(0))
	f.Add(int32(math.MaxInt32), int(math.MinInt))
	f.Add(int32(math.MaxInt32), int(math.MaxInt))
	f.Fuzz(func(t *testing.T, k int32, v int) {
		vm := map[int32]int{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt32Int8Map(f *testing.F) {
	f.Add(int32(0), int8(0))
	f.Add(int32(0), int8(math.MinInt8))
	f.Add(int32(0), int8(math.MaxInt8))
	f.Add(int32(math.MinInt32), int8(0))
	f.Add(int32(math.MinInt32), int8(math.MinInt8))
	f.Add(int32(math.MinInt32), int8(math.MaxInt8))
	f.Add(int32(math.MaxInt32), int8(0))
	f.Add(int32(math.MaxInt32), int8(math.MinInt8))
	f.Add(int32(math.MaxInt32), int8(math.MaxInt8))
	f.Fuzz(func(t *testing.T, k int32, v int8) {
		vm := map[int32]int8{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt32Int16Map(f *testing.F) {
	f.Add(int32(0), int16(0))
	f.Add(int32(0), int16(math.MinInt16))
	f.Add(int32(0), int16(math.MaxInt16))
	f.Add(int32(math.MinInt32), int16(0))
	f.Add(int32(math.MinInt32), int16(math.MinInt16))
	f.Add(int32(math.MinInt32), int16(math.MaxInt16))
	f.Add(int32(math.MaxInt32), int16(0))
	f.Add(int32(math.MaxInt32), int16(math.MinInt16))
	f.Add(int32(math.MaxInt32), int16(math.MaxInt16))
	f.Fuzz(func(t *testing.T, k int32, v int16) {
		vm := map[int32]int16{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt32Int32Map(f *testing.F) {
	f.Add(int32(0), int32(0))
	f.Add(int32(0), int32(math.MinInt32))
	f.Add(int32(0), int32(math.MaxInt32))
	f.Add(int32(math.MinInt32), int32(0))
	f.Add(int32(math.MinInt32), int32(math.MinInt32))
	f.Add(int32(math.MinInt32), int32(math.MaxInt32))
	f.Add(int32(math.MaxInt32), int32(0))
	f.Add(int32(math.MaxInt32), int32(math.MinInt32))
	f.Add(int32(math.MaxInt32), int32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, k int32, v int32) {
		vm := map[int32]int32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt32Int64Map(f *testing.F) {
	f.Add(int32(0), int64(0))
	f.Add(int32(0), int64(math.MinInt64))
	f.Add(int32(0), int64(math.MaxInt64))
	f.Add(int32(math.MinInt32), int64(0))
	f.Add(int32(math.MinInt32), int64(math.MinInt64))
	f.Add(int32(math.MinInt32), int64(math.MaxInt64))
	f.Add(int32(math.MaxInt32), int64(0))
	f.Add(int32(math.MaxInt32), int64(math.MinInt64))
	f.Add(int32(math.MaxInt32), int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k int32, v int64) {
		vm := map[int32]int64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt32UintMap(f *testing.F) {
	f.Add(int32(0), uint(0))
	f.Add(int32(0), uint(math.MaxUint))
	f.Add(int32(math.MinInt32), uint(0))
	f.Add(int32(math.MinInt32), uint(math.MaxUint))
	f.Add(int32(math.MaxInt32), uint(0))
	f.Add(int32(math.MaxInt32), uint(math.MaxUint))
	f.Fuzz(func(t *testing.T, k int32, v uint) {
		vm := map[int32]uint{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt32Uint8Map(f *testing.F) {
	f.Add(int32(0), uint8(0))
	f.Add(int32(0), uint8(math.MaxUint8))
	f.Add(int32(math.MinInt32), uint8(0))
	f.Add(int32(math.MinInt32), uint8(math.MaxUint8))
	f.Add(int32(math.MaxInt32), uint8(0))
	f.Add(int32(math.MaxInt32), uint8(math.MaxUint8))
	f.Fuzz(func(t *testing.T, k int32, v uint8) {
		vm := map[int32]uint8{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt32Uint16Map(f *testing.F) {
	f.Add(int32(0), uint16(0))
	f.Add(int32(0), uint16(math.MaxUint16))
	f.Add(int32(math.MinInt32), uint16(0))
	f.Add(int32(math.MinInt32), uint16(math.MaxUint16))
	f.Add(int32(math.MaxInt32), uint16(0))
	f.Add(int32(math.MaxInt32), uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, k int32, v uint16) {
		vm := map[int32]uint16{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt32Uint32Map(f *testing.F) {
	f.Add(int32(0), uint32(0))
	f.Add(int32(0), uint32(math.MaxUint32))
	f.Add(int32(math.MinInt32), uint32(0))
	f.Add(int32(math.MinInt32), uint32(math.MaxUint32))
	f.Add(int32(math.MaxInt32), uint32(0))
	f.Add(int32(math.MaxInt32), uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, k int32, v uint32) {
		vm := map[int32]uint32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt32Uint64Map(f *testing.F) {
	f.Add(int32(0), uint64(0))
	f.Add(int32(0), uint64(math.MaxInt64))
	f.Add(int32(math.MinInt32), uint64(0))
	f.Add(int32(math.MinInt32), uint64(math.MaxInt64))
	f.Add(int32(math.MaxInt32), uint64(0))
	f.Add(int32(math.MaxInt32), uint64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k int32, v uint64) {
		vm := map[int32]uint64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt32Float32Map(f *testing.F) {
	f.Add(int32(0), float32(math.MaxFloat32))
	f.Add(int32(math.MinInt32), float32(math.MaxFloat32))
	f.Add(int32(math.MaxInt32), float32(math.MaxFloat32))
	f.Fuzz(func(t *testing.T, k int32, v float32) {
		vm := map[int32]float32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt32Float64Map(f *testing.F) {
	f.Add(int32(0), float64(math.MaxFloat64))
	f.Add(int32(math.MinInt32), float64(math.MaxFloat64))
	f.Add(int32(math.MaxInt32), float64(math.MaxFloat64))
	f.Fuzz(func(t *testing.T, k int32, v float64) {
		vm := map[int32]float64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt32BoolMap(f *testing.F) {
	f.Add(int32(0), bool(true))
	f.Add(int32(0), bool(false))
	f.Add(int32(math.MinInt32), bool(true))
	f.Add(int32(math.MinInt32), bool(false))
	f.Add(int32(math.MaxInt32), bool(true))
	f.Add(int32(math.MaxInt32), bool(false))
	f.Fuzz(func(t *testing.T, k int32, v bool) {
		vm := map[int32]bool{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt32StringMap(f *testing.F) {
	f.Add(int32(0), string("abc"))
	f.Add(int32(0), string("xyz"))
	f.Add(int32(math.MinInt32), string("abc"))
	f.Add(int32(math.MinInt32), string("xyz"))
	f.Add(int32(math.MaxInt32), string("abc"))
	f.Add(int32(math.MaxInt32), string("xyz"))
	f.Fuzz(func(t *testing.T, k int32, v string) {
		vm := map[int32]string{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt64IntMap(f *testing.F) {
	f.Add(int64(0), int(0))
	f.Add(int64(0), int(math.MinInt))
	f.Add(int64(0), int(math.MaxInt))
	f.Add(int64(math.MinInt64), int(0))
	f.Add(int64(math.MinInt64), int(math.MinInt))
	f.Add(int64(math.MinInt64), int(math.MaxInt))
	f.Add(int64(math.MaxInt64), int(0))
	f.Add(int64(math.MaxInt64), int(math.MinInt))
	f.Add(int64(math.MaxInt64), int(math.MaxInt))
	f.Fuzz(func(t *testing.T, k int64, v int) {
		vm := map[int64]int{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt64Int8Map(f *testing.F) {
	f.Add(int64(0), int8(0))
	f.Add(int64(0), int8(math.MinInt8))
	f.Add(int64(0), int8(math.MaxInt8))
	f.Add(int64(math.MinInt64), int8(0))
	f.Add(int64(math.MinInt64), int8(math.MinInt8))
	f.Add(int64(math.MinInt64), int8(math.MaxInt8))
	f.Add(int64(math.MaxInt64), int8(0))
	f.Add(int64(math.MaxInt64), int8(math.MinInt8))
	f.Add(int64(math.MaxInt64), int8(math.MaxInt8))
	f.Fuzz(func(t *testing.T, k int64, v int8) {
		vm := map[int64]int8{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt64Int16Map(f *testing.F) {
	f.Add(int64(0), int16(0))
	f.Add(int64(0), int16(math.MinInt16))
	f.Add(int64(0), int16(math.MaxInt16))
	f.Add(int64(math.MinInt64), int16(0))
	f.Add(int64(math.MinInt64), int16(math.MinInt16))
	f.Add(int64(math.MinInt64), int16(math.MaxInt16))
	f.Add(int64(math.MaxInt64), int16(0))
	f.Add(int64(math.MaxInt64), int16(math.MinInt16))
	f.Add(int64(math.MaxInt64), int16(math.MaxInt16))
	f.Fuzz(func(t *testing.T, k int64, v int16) {
		vm := map[int64]int16{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt64Int32Map(f *testing.F) {
	f.Add(int64(0), int32(0))
	f.Add(int64(0), int32(math.MinInt32))
	f.Add(int64(0), int32(math.MaxInt32))
	f.Add(int64(math.MinInt64), int32(0))
	f.Add(int64(math.MinInt64), int32(math.MinInt32))
	f.Add(int64(math.MinInt64), int32(math.MaxInt32))
	f.Add(int64(math.MaxInt64), int32(0))
	f.Add(int64(math.MaxInt64), int32(math.MinInt32))
	f.Add(int64(math.MaxInt64), int32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, k int64, v int32) {
		vm := map[int64]int32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt64Int64Map(f *testing.F) {
	f.Add(int64(0), int64(0))
	f.Add(int64(0), int64(math.MinInt64))
	f.Add(int64(0), int64(math.MaxInt64))
	f.Add(int64(math.MinInt64), int64(0))
	f.Add(int64(math.MinInt64), int64(math.MinInt64))
	f.Add(int64(math.MinInt64), int64(math.MaxInt64))
	f.Add(int64(math.MaxInt64), int64(0))
	f.Add(int64(math.MaxInt64), int64(math.MinInt64))
	f.Add(int64(math.MaxInt64), int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k int64, v int64) {
		vm := map[int64]int64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt64UintMap(f *testing.F) {
	f.Add(int64(0), uint(0))
	f.Add(int64(0), uint(math.MaxUint))
	f.Add(int64(math.MinInt64), uint(0))
	f.Add(int64(math.MinInt64), uint(math.MaxUint))
	f.Add(int64(math.MaxInt64), uint(0))
	f.Add(int64(math.MaxInt64), uint(math.MaxUint))
	f.Fuzz(func(t *testing.T, k int64, v uint) {
		vm := map[int64]uint{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt64Uint8Map(f *testing.F) {
	f.Add(int64(0), uint8(0))
	f.Add(int64(0), uint8(math.MaxUint8))
	f.Add(int64(math.MinInt64), uint8(0))
	f.Add(int64(math.MinInt64), uint8(math.MaxUint8))
	f.Add(int64(math.MaxInt64), uint8(0))
	f.Add(int64(math.MaxInt64), uint8(math.MaxUint8))
	f.Fuzz(func(t *testing.T, k int64, v uint8) {
		vm := map[int64]uint8{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt64Uint16Map(f *testing.F) {
	f.Add(int64(0), uint16(0))
	f.Add(int64(0), uint16(math.MaxUint16))
	f.Add(int64(math.MinInt64), uint16(0))
	f.Add(int64(math.MinInt64), uint16(math.MaxUint16))
	f.Add(int64(math.MaxInt64), uint16(0))
	f.Add(int64(math.MaxInt64), uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, k int64, v uint16) {
		vm := map[int64]uint16{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt64Uint32Map(f *testing.F) {
	f.Add(int64(0), uint32(0))
	f.Add(int64(0), uint32(math.MaxUint32))
	f.Add(int64(math.MinInt64), uint32(0))
	f.Add(int64(math.MinInt64), uint32(math.MaxUint32))
	f.Add(int64(math.MaxInt64), uint32(0))
	f.Add(int64(math.MaxInt64), uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, k int64, v uint32) {
		vm := map[int64]uint32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt64Uint64Map(f *testing.F) {
	f.Add(int64(0), uint64(0))
	f.Add(int64(0), uint64(math.MaxInt64))
	f.Add(int64(math.MinInt64), uint64(0))
	f.Add(int64(math.MinInt64), uint64(math.MaxInt64))
	f.Add(int64(math.MaxInt64), uint64(0))
	f.Add(int64(math.MaxInt64), uint64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k int64, v uint64) {
		vm := map[int64]uint64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt64Float32Map(f *testing.F) {
	f.Add(int64(0), float32(math.MaxFloat32))
	f.Add(int64(math.MinInt64), float32(math.MaxFloat32))
	f.Add(int64(math.MaxInt64), float32(math.MaxFloat32))
	f.Fuzz(func(t *testing.T, k int64, v float32) {
		vm := map[int64]float32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt64Float64Map(f *testing.F) {
	f.Add(int64(0), float64(math.MaxFloat64))
	f.Add(int64(math.MinInt64), float64(math.MaxFloat64))
	f.Add(int64(math.MaxInt64), float64(math.MaxFloat64))
	f.Fuzz(func(t *testing.T, k int64, v float64) {
		vm := map[int64]float64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt64BoolMap(f *testing.F) {
	f.Add(int64(0), bool(true))
	f.Add(int64(0), bool(false))
	f.Add(int64(math.MinInt64), bool(true))
	f.Add(int64(math.MinInt64), bool(false))
	f.Add(int64(math.MaxInt64), bool(true))
	f.Add(int64(math.MaxInt64), bool(false))
	f.Fuzz(func(t *testing.T, k int64, v bool) {
		vm := map[int64]bool{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzInt64StringMap(f *testing.F) {
	f.Add(int64(0), string("abc"))
	f.Add(int64(0), string("xyz"))
	f.Add(int64(math.MinInt64), string("abc"))
	f.Add(int64(math.MinInt64), string("xyz"))
	f.Add(int64(math.MaxInt64), string("abc"))
	f.Add(int64(math.MaxInt64), string("xyz"))
	f.Fuzz(func(t *testing.T, k int64, v string) {
		vm := map[int64]string{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUintIntMap(f *testing.F) {
	f.Add(uint(0), int(0))
	f.Add(uint(0), int(math.MinInt))
	f.Add(uint(0), int(math.MaxInt))
	f.Add(uint(math.MaxUint), int(0))
	f.Add(uint(math.MaxUint), int(math.MinInt))
	f.Add(uint(math.MaxUint), int(math.MaxInt))
	f.Fuzz(func(t *testing.T, k uint, v int) {
		vm := map[uint]int{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUintInt8Map(f *testing.F) {
	f.Add(uint(0), int8(0))
	f.Add(uint(0), int8(math.MinInt8))
	f.Add(uint(0), int8(math.MaxInt8))
	f.Add(uint(math.MaxUint), int8(0))
	f.Add(uint(math.MaxUint), int8(math.MinInt8))
	f.Add(uint(math.MaxUint), int8(math.MaxInt8))
	f.Fuzz(func(t *testing.T, k uint, v int8) {
		vm := map[uint]int8{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUintInt16Map(f *testing.F) {
	f.Add(uint(0), int16(0))
	f.Add(uint(0), int16(math.MinInt16))
	f.Add(uint(0), int16(math.MaxInt16))
	f.Add(uint(math.MaxUint), int16(0))
	f.Add(uint(math.MaxUint), int16(math.MinInt16))
	f.Add(uint(math.MaxUint), int16(math.MaxInt16))
	f.Fuzz(func(t *testing.T, k uint, v int16) {
		vm := map[uint]int16{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUintInt32Map(f *testing.F) {
	f.Add(uint(0), int32(0))
	f.Add(uint(0), int32(math.MinInt32))
	f.Add(uint(0), int32(math.MaxInt32))
	f.Add(uint(math.MaxUint), int32(0))
	f.Add(uint(math.MaxUint), int32(math.MinInt32))
	f.Add(uint(math.MaxUint), int32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, k uint, v int32) {
		vm := map[uint]int32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUintInt64Map(f *testing.F) {
	f.Add(uint(0), int64(0))
	f.Add(uint(0), int64(math.MinInt64))
	f.Add(uint(0), int64(math.MaxInt64))
	f.Add(uint(math.MaxUint), int64(0))
	f.Add(uint(math.MaxUint), int64(math.MinInt64))
	f.Add(uint(math.MaxUint), int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k uint, v int64) {
		vm := map[uint]int64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUintUintMap(f *testing.F) {
	f.Add(uint(0), uint(0))
	f.Add(uint(0), uint(math.MaxUint))
	f.Add(uint(math.MaxUint), uint(0))
	f.Add(uint(math.MaxUint), uint(math.MaxUint))
	f.Fuzz(func(t *testing.T, k uint, v uint) {
		vm := map[uint]uint{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUintUint8Map(f *testing.F) {
	f.Add(uint(0), uint8(0))
	f.Add(uint(0), uint8(math.MaxUint8))
	f.Add(uint(math.MaxUint), uint8(0))
	f.Add(uint(math.MaxUint), uint8(math.MaxUint8))
	f.Fuzz(func(t *testing.T, k uint, v uint8) {
		vm := map[uint]uint8{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUintUint16Map(f *testing.F) {
	f.Add(uint(0), uint16(0))
	f.Add(uint(0), uint16(math.MaxUint16))
	f.Add(uint(math.MaxUint), uint16(0))
	f.Add(uint(math.MaxUint), uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, k uint, v uint16) {
		vm := map[uint]uint16{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUintUint32Map(f *testing.F) {
	f.Add(uint(0), uint32(0))
	f.Add(uint(0), uint32(math.MaxUint32))
	f.Add(uint(math.MaxUint), uint32(0))
	f.Add(uint(math.MaxUint), uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, k uint, v uint32) {
		vm := map[uint]uint32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUintUint64Map(f *testing.F) {
	f.Add(uint(0), uint64(0))
	f.Add(uint(0), uint64(math.MaxInt64))
	f.Add(uint(math.MaxUint), uint64(0))
	f.Add(uint(math.MaxUint), uint64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k uint, v uint64) {
		vm := map[uint]uint64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUintFloat32Map(f *testing.F) {
	f.Add(uint(0), float32(math.MaxFloat32))
	f.Add(uint(math.MaxUint), float32(math.MaxFloat32))
	f.Fuzz(func(t *testing.T, k uint, v float32) {
		vm := map[uint]float32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUintFloat64Map(f *testing.F) {
	f.Add(uint(0), float64(math.MaxFloat64))
	f.Add(uint(math.MaxUint), float64(math.MaxFloat64))
	f.Fuzz(func(t *testing.T, k uint, v float64) {
		vm := map[uint]float64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUintBoolMap(f *testing.F) {
	f.Add(uint(0), bool(true))
	f.Add(uint(0), bool(false))
	f.Add(uint(math.MaxUint), bool(true))
	f.Add(uint(math.MaxUint), bool(false))
	f.Fuzz(func(t *testing.T, k uint, v bool) {
		vm := map[uint]bool{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUintStringMap(f *testing.F) {
	f.Add(uint(0), string("abc"))
	f.Add(uint(0), string("xyz"))
	f.Add(uint(math.MaxUint), string("abc"))
	f.Add(uint(math.MaxUint), string("xyz"))
	f.Fuzz(func(t *testing.T, k uint, v string) {
		vm := map[uint]string{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint8IntMap(f *testing.F) {
	f.Add(uint8(0), int(0))
	f.Add(uint8(0), int(math.MinInt))
	f.Add(uint8(0), int(math.MaxInt))
	f.Add(uint8(math.MaxUint8), int(0))
	f.Add(uint8(math.MaxUint8), int(math.MinInt))
	f.Add(uint8(math.MaxUint8), int(math.MaxInt))
	f.Fuzz(func(t *testing.T, k uint8, v int) {
		vm := map[uint8]int{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint8Int8Map(f *testing.F) {
	f.Add(uint8(0), int8(0))
	f.Add(uint8(0), int8(math.MinInt8))
	f.Add(uint8(0), int8(math.MaxInt8))
	f.Add(uint8(math.MaxUint8), int8(0))
	f.Add(uint8(math.MaxUint8), int8(math.MinInt8))
	f.Add(uint8(math.MaxUint8), int8(math.MaxInt8))
	f.Fuzz(func(t *testing.T, k uint8, v int8) {
		vm := map[uint8]int8{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint8Int16Map(f *testing.F) {
	f.Add(uint8(0), int16(0))
	f.Add(uint8(0), int16(math.MinInt16))
	f.Add(uint8(0), int16(math.MaxInt16))
	f.Add(uint8(math.MaxUint8), int16(0))
	f.Add(uint8(math.MaxUint8), int16(math.MinInt16))
	f.Add(uint8(math.MaxUint8), int16(math.MaxInt16))
	f.Fuzz(func(t *testing.T, k uint8, v int16) {
		vm := map[uint8]int16{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint8Int32Map(f *testing.F) {
	f.Add(uint8(0), int32(0))
	f.Add(uint8(0), int32(math.MinInt32))
	f.Add(uint8(0), int32(math.MaxInt32))
	f.Add(uint8(math.MaxUint8), int32(0))
	f.Add(uint8(math.MaxUint8), int32(math.MinInt32))
	f.Add(uint8(math.MaxUint8), int32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, k uint8, v int32) {
		vm := map[uint8]int32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint8Int64Map(f *testing.F) {
	f.Add(uint8(0), int64(0))
	f.Add(uint8(0), int64(math.MinInt64))
	f.Add(uint8(0), int64(math.MaxInt64))
	f.Add(uint8(math.MaxUint8), int64(0))
	f.Add(uint8(math.MaxUint8), int64(math.MinInt64))
	f.Add(uint8(math.MaxUint8), int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k uint8, v int64) {
		vm := map[uint8]int64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint8UintMap(f *testing.F) {
	f.Add(uint8(0), uint(0))
	f.Add(uint8(0), uint(math.MaxUint))
	f.Add(uint8(math.MaxUint8), uint(0))
	f.Add(uint8(math.MaxUint8), uint(math.MaxUint))
	f.Fuzz(func(t *testing.T, k uint8, v uint) {
		vm := map[uint8]uint{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint8Uint8Map(f *testing.F) {
	f.Add(uint8(0), uint8(0))
	f.Add(uint8(0), uint8(math.MaxUint8))
	f.Add(uint8(math.MaxUint8), uint8(0))
	f.Add(uint8(math.MaxUint8), uint8(math.MaxUint8))
	f.Fuzz(func(t *testing.T, k uint8, v uint8) {
		vm := map[uint8]uint8{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint8Uint16Map(f *testing.F) {
	f.Add(uint8(0), uint16(0))
	f.Add(uint8(0), uint16(math.MaxUint16))
	f.Add(uint8(math.MaxUint8), uint16(0))
	f.Add(uint8(math.MaxUint8), uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, k uint8, v uint16) {
		vm := map[uint8]uint16{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint8Uint32Map(f *testing.F) {
	f.Add(uint8(0), uint32(0))
	f.Add(uint8(0), uint32(math.MaxUint32))
	f.Add(uint8(math.MaxUint8), uint32(0))
	f.Add(uint8(math.MaxUint8), uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, k uint8, v uint32) {
		vm := map[uint8]uint32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint8Uint64Map(f *testing.F) {
	f.Add(uint8(0), uint64(0))
	f.Add(uint8(0), uint64(math.MaxInt64))
	f.Add(uint8(math.MaxUint8), uint64(0))
	f.Add(uint8(math.MaxUint8), uint64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k uint8, v uint64) {
		vm := map[uint8]uint64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint8Float32Map(f *testing.F) {
	f.Add(uint8(0), float32(math.MaxFloat32))
	f.Add(uint8(math.MaxUint8), float32(math.MaxFloat32))
	f.Fuzz(func(t *testing.T, k uint8, v float32) {
		vm := map[uint8]float32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint8Float64Map(f *testing.F) {
	f.Add(uint8(0), float64(math.MaxFloat64))
	f.Add(uint8(math.MaxUint8), float64(math.MaxFloat64))
	f.Fuzz(func(t *testing.T, k uint8, v float64) {
		vm := map[uint8]float64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint8BoolMap(f *testing.F) {
	f.Add(uint8(0), bool(true))
	f.Add(uint8(0), bool(false))
	f.Add(uint8(math.MaxUint8), bool(true))
	f.Add(uint8(math.MaxUint8), bool(false))
	f.Fuzz(func(t *testing.T, k uint8, v bool) {
		vm := map[uint8]bool{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint8StringMap(f *testing.F) {
	f.Add(uint8(0), string("abc"))
	f.Add(uint8(0), string("xyz"))
	f.Add(uint8(math.MaxUint8), string("abc"))
	f.Add(uint8(math.MaxUint8), string("xyz"))
	f.Fuzz(func(t *testing.T, k uint8, v string) {
		vm := map[uint8]string{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint16IntMap(f *testing.F) {
	f.Add(uint16(0), int(0))
	f.Add(uint16(0), int(math.MinInt))
	f.Add(uint16(0), int(math.MaxInt))
	f.Add(uint16(math.MaxUint16), int(0))
	f.Add(uint16(math.MaxUint16), int(math.MinInt))
	f.Add(uint16(math.MaxUint16), int(math.MaxInt))
	f.Fuzz(func(t *testing.T, k uint16, v int) {
		vm := map[uint16]int{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint16Int8Map(f *testing.F) {
	f.Add(uint16(0), int8(0))
	f.Add(uint16(0), int8(math.MinInt8))
	f.Add(uint16(0), int8(math.MaxInt8))
	f.Add(uint16(math.MaxUint16), int8(0))
	f.Add(uint16(math.MaxUint16), int8(math.MinInt8))
	f.Add(uint16(math.MaxUint16), int8(math.MaxInt8))
	f.Fuzz(func(t *testing.T, k uint16, v int8) {
		vm := map[uint16]int8{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint16Int16Map(f *testing.F) {
	f.Add(uint16(0), int16(0))
	f.Add(uint16(0), int16(math.MinInt16))
	f.Add(uint16(0), int16(math.MaxInt16))
	f.Add(uint16(math.MaxUint16), int16(0))
	f.Add(uint16(math.MaxUint16), int16(math.MinInt16))
	f.Add(uint16(math.MaxUint16), int16(math.MaxInt16))
	f.Fuzz(func(t *testing.T, k uint16, v int16) {
		vm := map[uint16]int16{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint16Int32Map(f *testing.F) {
	f.Add(uint16(0), int32(0))
	f.Add(uint16(0), int32(math.MinInt32))
	f.Add(uint16(0), int32(math.MaxInt32))
	f.Add(uint16(math.MaxUint16), int32(0))
	f.Add(uint16(math.MaxUint16), int32(math.MinInt32))
	f.Add(uint16(math.MaxUint16), int32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, k uint16, v int32) {
		vm := map[uint16]int32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint16Int64Map(f *testing.F) {
	f.Add(uint16(0), int64(0))
	f.Add(uint16(0), int64(math.MinInt64))
	f.Add(uint16(0), int64(math.MaxInt64))
	f.Add(uint16(math.MaxUint16), int64(0))
	f.Add(uint16(math.MaxUint16), int64(math.MinInt64))
	f.Add(uint16(math.MaxUint16), int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k uint16, v int64) {
		vm := map[uint16]int64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint16UintMap(f *testing.F) {
	f.Add(uint16(0), uint(0))
	f.Add(uint16(0), uint(math.MaxUint))
	f.Add(uint16(math.MaxUint16), uint(0))
	f.Add(uint16(math.MaxUint16), uint(math.MaxUint))
	f.Fuzz(func(t *testing.T, k uint16, v uint) {
		vm := map[uint16]uint{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint16Uint8Map(f *testing.F) {
	f.Add(uint16(0), uint8(0))
	f.Add(uint16(0), uint8(math.MaxUint8))
	f.Add(uint16(math.MaxUint16), uint8(0))
	f.Add(uint16(math.MaxUint16), uint8(math.MaxUint8))
	f.Fuzz(func(t *testing.T, k uint16, v uint8) {
		vm := map[uint16]uint8{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint16Uint16Map(f *testing.F) {
	f.Add(uint16(0), uint16(0))
	f.Add(uint16(0), uint16(math.MaxUint16))
	f.Add(uint16(math.MaxUint16), uint16(0))
	f.Add(uint16(math.MaxUint16), uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, k uint16, v uint16) {
		vm := map[uint16]uint16{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint16Uint32Map(f *testing.F) {
	f.Add(uint16(0), uint32(0))
	f.Add(uint16(0), uint32(math.MaxUint32))
	f.Add(uint16(math.MaxUint16), uint32(0))
	f.Add(uint16(math.MaxUint16), uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, k uint16, v uint32) {
		vm := map[uint16]uint32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint16Uint64Map(f *testing.F) {
	f.Add(uint16(0), uint64(0))
	f.Add(uint16(0), uint64(math.MaxInt64))
	f.Add(uint16(math.MaxUint16), uint64(0))
	f.Add(uint16(math.MaxUint16), uint64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k uint16, v uint64) {
		vm := map[uint16]uint64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint16Float32Map(f *testing.F) {
	f.Add(uint16(0), float32(math.MaxFloat32))
	f.Add(uint16(math.MaxUint16), float32(math.MaxFloat32))
	f.Fuzz(func(t *testing.T, k uint16, v float32) {
		vm := map[uint16]float32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint16Float64Map(f *testing.F) {
	f.Add(uint16(0), float64(math.MaxFloat64))
	f.Add(uint16(math.MaxUint16), float64(math.MaxFloat64))
	f.Fuzz(func(t *testing.T, k uint16, v float64) {
		vm := map[uint16]float64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint16BoolMap(f *testing.F) {
	f.Add(uint16(0), bool(true))
	f.Add(uint16(0), bool(false))
	f.Add(uint16(math.MaxUint16), bool(true))
	f.Add(uint16(math.MaxUint16), bool(false))
	f.Fuzz(func(t *testing.T, k uint16, v bool) {
		vm := map[uint16]bool{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint16StringMap(f *testing.F) {
	f.Add(uint16(0), string("abc"))
	f.Add(uint16(0), string("xyz"))
	f.Add(uint16(math.MaxUint16), string("abc"))
	f.Add(uint16(math.MaxUint16), string("xyz"))
	f.Fuzz(func(t *testing.T, k uint16, v string) {
		vm := map[uint16]string{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint32IntMap(f *testing.F) {
	f.Add(uint32(0), int(0))
	f.Add(uint32(0), int(math.MinInt))
	f.Add(uint32(0), int(math.MaxInt))
	f.Add(uint32(math.MaxUint32), int(0))
	f.Add(uint32(math.MaxUint32), int(math.MinInt))
	f.Add(uint32(math.MaxUint32), int(math.MaxInt))
	f.Fuzz(func(t *testing.T, k uint32, v int) {
		vm := map[uint32]int{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint32Int8Map(f *testing.F) {
	f.Add(uint32(0), int8(0))
	f.Add(uint32(0), int8(math.MinInt8))
	f.Add(uint32(0), int8(math.MaxInt8))
	f.Add(uint32(math.MaxUint32), int8(0))
	f.Add(uint32(math.MaxUint32), int8(math.MinInt8))
	f.Add(uint32(math.MaxUint32), int8(math.MaxInt8))
	f.Fuzz(func(t *testing.T, k uint32, v int8) {
		vm := map[uint32]int8{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint32Int16Map(f *testing.F) {
	f.Add(uint32(0), int16(0))
	f.Add(uint32(0), int16(math.MinInt16))
	f.Add(uint32(0), int16(math.MaxInt16))
	f.Add(uint32(math.MaxUint32), int16(0))
	f.Add(uint32(math.MaxUint32), int16(math.MinInt16))
	f.Add(uint32(math.MaxUint32), int16(math.MaxInt16))
	f.Fuzz(func(t *testing.T, k uint32, v int16) {
		vm := map[uint32]int16{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint32Int32Map(f *testing.F) {
	f.Add(uint32(0), int32(0))
	f.Add(uint32(0), int32(math.MinInt32))
	f.Add(uint32(0), int32(math.MaxInt32))
	f.Add(uint32(math.MaxUint32), int32(0))
	f.Add(uint32(math.MaxUint32), int32(math.MinInt32))
	f.Add(uint32(math.MaxUint32), int32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, k uint32, v int32) {
		vm := map[uint32]int32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint32Int64Map(f *testing.F) {
	f.Add(uint32(0), int64(0))
	f.Add(uint32(0), int64(math.MinInt64))
	f.Add(uint32(0), int64(math.MaxInt64))
	f.Add(uint32(math.MaxUint32), int64(0))
	f.Add(uint32(math.MaxUint32), int64(math.MinInt64))
	f.Add(uint32(math.MaxUint32), int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k uint32, v int64) {
		vm := map[uint32]int64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint32UintMap(f *testing.F) {
	f.Add(uint32(0), uint(0))
	f.Add(uint32(0), uint(math.MaxUint))
	f.Add(uint32(math.MaxUint32), uint(0))
	f.Add(uint32(math.MaxUint32), uint(math.MaxUint))
	f.Fuzz(func(t *testing.T, k uint32, v uint) {
		vm := map[uint32]uint{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint32Uint8Map(f *testing.F) {
	f.Add(uint32(0), uint8(0))
	f.Add(uint32(0), uint8(math.MaxUint8))
	f.Add(uint32(math.MaxUint32), uint8(0))
	f.Add(uint32(math.MaxUint32), uint8(math.MaxUint8))
	f.Fuzz(func(t *testing.T, k uint32, v uint8) {
		vm := map[uint32]uint8{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint32Uint16Map(f *testing.F) {
	f.Add(uint32(0), uint16(0))
	f.Add(uint32(0), uint16(math.MaxUint16))
	f.Add(uint32(math.MaxUint32), uint16(0))
	f.Add(uint32(math.MaxUint32), uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, k uint32, v uint16) {
		vm := map[uint32]uint16{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint32Uint32Map(f *testing.F) {
	f.Add(uint32(0), uint32(0))
	f.Add(uint32(0), uint32(math.MaxUint32))
	f.Add(uint32(math.MaxUint32), uint32(0))
	f.Add(uint32(math.MaxUint32), uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, k uint32, v uint32) {
		vm := map[uint32]uint32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint32Uint64Map(f *testing.F) {
	f.Add(uint32(0), uint64(0))
	f.Add(uint32(0), uint64(math.MaxInt64))
	f.Add(uint32(math.MaxUint32), uint64(0))
	f.Add(uint32(math.MaxUint32), uint64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k uint32, v uint64) {
		vm := map[uint32]uint64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint32Float32Map(f *testing.F) {
	f.Add(uint32(0), float32(math.MaxFloat32))
	f.Add(uint32(math.MaxUint32), float32(math.MaxFloat32))
	f.Fuzz(func(t *testing.T, k uint32, v float32) {
		vm := map[uint32]float32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint32Float64Map(f *testing.F) {
	f.Add(uint32(0), float64(math.MaxFloat64))
	f.Add(uint32(math.MaxUint32), float64(math.MaxFloat64))
	f.Fuzz(func(t *testing.T, k uint32, v float64) {
		vm := map[uint32]float64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint32BoolMap(f *testing.F) {
	f.Add(uint32(0), bool(true))
	f.Add(uint32(0), bool(false))
	f.Add(uint32(math.MaxUint32), bool(true))
	f.Add(uint32(math.MaxUint32), bool(false))
	f.Fuzz(func(t *testing.T, k uint32, v bool) {
		vm := map[uint32]bool{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint32StringMap(f *testing.F) {
	f.Add(uint32(0), string("abc"))
	f.Add(uint32(0), string("xyz"))
	f.Add(uint32(math.MaxUint32), string("abc"))
	f.Add(uint32(math.MaxUint32), string("xyz"))
	f.Fuzz(func(t *testing.T, k uint32, v string) {
		vm := map[uint32]string{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint64IntMap(f *testing.F) {
	f.Add(uint64(0), int(0))
	f.Add(uint64(0), int(math.MinInt))
	f.Add(uint64(0), int(math.MaxInt))
	f.Add(uint64(math.MaxInt64), int(0))
	f.Add(uint64(math.MaxInt64), int(math.MinInt))
	f.Add(uint64(math.MaxInt64), int(math.MaxInt))
	f.Fuzz(func(t *testing.T, k uint64, v int) {
		vm := map[uint64]int{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint64Int8Map(f *testing.F) {
	f.Add(uint64(0), int8(0))
	f.Add(uint64(0), int8(math.MinInt8))
	f.Add(uint64(0), int8(math.MaxInt8))
	f.Add(uint64(math.MaxInt64), int8(0))
	f.Add(uint64(math.MaxInt64), int8(math.MinInt8))
	f.Add(uint64(math.MaxInt64), int8(math.MaxInt8))
	f.Fuzz(func(t *testing.T, k uint64, v int8) {
		vm := map[uint64]int8{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint64Int16Map(f *testing.F) {
	f.Add(uint64(0), int16(0))
	f.Add(uint64(0), int16(math.MinInt16))
	f.Add(uint64(0), int16(math.MaxInt16))
	f.Add(uint64(math.MaxInt64), int16(0))
	f.Add(uint64(math.MaxInt64), int16(math.MinInt16))
	f.Add(uint64(math.MaxInt64), int16(math.MaxInt16))
	f.Fuzz(func(t *testing.T, k uint64, v int16) {
		vm := map[uint64]int16{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint64Int32Map(f *testing.F) {
	f.Add(uint64(0), int32(0))
	f.Add(uint64(0), int32(math.MinInt32))
	f.Add(uint64(0), int32(math.MaxInt32))
	f.Add(uint64(math.MaxInt64), int32(0))
	f.Add(uint64(math.MaxInt64), int32(math.MinInt32))
	f.Add(uint64(math.MaxInt64), int32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, k uint64, v int32) {
		vm := map[uint64]int32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint64Int64Map(f *testing.F) {
	f.Add(uint64(0), int64(0))
	f.Add(uint64(0), int64(math.MinInt64))
	f.Add(uint64(0), int64(math.MaxInt64))
	f.Add(uint64(math.MaxInt64), int64(0))
	f.Add(uint64(math.MaxInt64), int64(math.MinInt64))
	f.Add(uint64(math.MaxInt64), int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k uint64, v int64) {
		vm := map[uint64]int64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint64UintMap(f *testing.F) {
	f.Add(uint64(0), uint(0))
	f.Add(uint64(0), uint(math.MaxUint))
	f.Add(uint64(math.MaxInt64), uint(0))
	f.Add(uint64(math.MaxInt64), uint(math.MaxUint))
	f.Fuzz(func(t *testing.T, k uint64, v uint) {
		vm := map[uint64]uint{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint64Uint8Map(f *testing.F) {
	f.Add(uint64(0), uint8(0))
	f.Add(uint64(0), uint8(math.MaxUint8))
	f.Add(uint64(math.MaxInt64), uint8(0))
	f.Add(uint64(math.MaxInt64), uint8(math.MaxUint8))
	f.Fuzz(func(t *testing.T, k uint64, v uint8) {
		vm := map[uint64]uint8{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint64Uint16Map(f *testing.F) {
	f.Add(uint64(0), uint16(0))
	f.Add(uint64(0), uint16(math.MaxUint16))
	f.Add(uint64(math.MaxInt64), uint16(0))
	f.Add(uint64(math.MaxInt64), uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, k uint64, v uint16) {
		vm := map[uint64]uint16{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint64Uint32Map(f *testing.F) {
	f.Add(uint64(0), uint32(0))
	f.Add(uint64(0), uint32(math.MaxUint32))
	f.Add(uint64(math.MaxInt64), uint32(0))
	f.Add(uint64(math.MaxInt64), uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, k uint64, v uint32) {
		vm := map[uint64]uint32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint64Uint64Map(f *testing.F) {
	f.Add(uint64(0), uint64(0))
	f.Add(uint64(0), uint64(math.MaxInt64))
	f.Add(uint64(math.MaxInt64), uint64(0))
	f.Add(uint64(math.MaxInt64), uint64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k uint64, v uint64) {
		vm := map[uint64]uint64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint64Float32Map(f *testing.F) {
	f.Add(uint64(0), float32(math.MaxFloat32))
	f.Add(uint64(math.MaxInt64), float32(math.MaxFloat32))
	f.Fuzz(func(t *testing.T, k uint64, v float32) {
		vm := map[uint64]float32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint64Float64Map(f *testing.F) {
	f.Add(uint64(0), float64(math.MaxFloat64))
	f.Add(uint64(math.MaxInt64), float64(math.MaxFloat64))
	f.Fuzz(func(t *testing.T, k uint64, v float64) {
		vm := map[uint64]float64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint64BoolMap(f *testing.F) {
	f.Add(uint64(0), bool(true))
	f.Add(uint64(0), bool(false))
	f.Add(uint64(math.MaxInt64), bool(true))
	f.Add(uint64(math.MaxInt64), bool(false))
	f.Fuzz(func(t *testing.T, k uint64, v bool) {
		vm := map[uint64]bool{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzUint64StringMap(f *testing.F) {
	f.Add(uint64(0), string("abc"))
	f.Add(uint64(0), string("xyz"))
	f.Add(uint64(math.MaxInt64), string("abc"))
	f.Add(uint64(math.MaxInt64), string("xyz"))
	f.Fuzz(func(t *testing.T, k uint64, v string) {
		vm := map[uint64]string{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzFloat32IntMap(f *testing.F) {
	f.Add(float32(math.MaxFloat32), int(0))
	f.Add(float32(math.MaxFloat32), int(math.MinInt))
	f.Add(float32(math.MaxFloat32), int(math.MaxInt))
	f.Fuzz(func(t *testing.T, k float32, v int) {
		vm := map[float32]int{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzFloat32Int8Map(f *testing.F) {
	f.Add(float32(math.MaxFloat32), int8(0))
	f.Add(float32(math.MaxFloat32), int8(math.MinInt8))
	f.Add(float32(math.MaxFloat32), int8(math.MaxInt8))
	f.Fuzz(func(t *testing.T, k float32, v int8) {
		vm := map[float32]int8{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzFloat32Int16Map(f *testing.F) {
	f.Add(float32(math.MaxFloat32), int16(0))
	f.Add(float32(math.MaxFloat32), int16(math.MinInt16))
	f.Add(float32(math.MaxFloat32), int16(math.MaxInt16))
	f.Fuzz(func(t *testing.T, k float32, v int16) {
		vm := map[float32]int16{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzFloat32Int32Map(f *testing.F) {
	f.Add(float32(math.MaxFloat32), int32(0))
	f.Add(float32(math.MaxFloat32), int32(math.MinInt32))
	f.Add(float32(math.MaxFloat32), int32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, k float32, v int32) {
		vm := map[float32]int32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzFloat32Int64Map(f *testing.F) {
	f.Add(float32(math.MaxFloat32), int64(0))
	f.Add(float32(math.MaxFloat32), int64(math.MinInt64))
	f.Add(float32(math.MaxFloat32), int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k float32, v int64) {
		vm := map[float32]int64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzFloat32UintMap(f *testing.F) {
	f.Add(float32(math.MaxFloat32), uint(0))
	f.Add(float32(math.MaxFloat32), uint(math.MaxUint))
	f.Fuzz(func(t *testing.T, k float32, v uint) {
		vm := map[float32]uint{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzFloat32Uint8Map(f *testing.F) {
	f.Add(float32(math.MaxFloat32), uint8(0))
	f.Add(float32(math.MaxFloat32), uint8(math.MaxUint8))
	f.Fuzz(func(t *testing.T, k float32, v uint8) {
		vm := map[float32]uint8{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzFloat32Uint16Map(f *testing.F) {
	f.Add(float32(math.MaxFloat32), uint16(0))
	f.Add(float32(math.MaxFloat32), uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, k float32, v uint16) {
		vm := map[float32]uint16{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzFloat32Uint32Map(f *testing.F) {
	f.Add(float32(math.MaxFloat32), uint32(0))
	f.Add(float32(math.MaxFloat32), uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, k float32, v uint32) {
		vm := map[float32]uint32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzFloat32Uint64Map(f *testing.F) {
	f.Add(float32(math.MaxFloat32), uint64(0))
	f.Add(float32(math.MaxFloat32), uint64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k float32, v uint64) {
		vm := map[float32]uint64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzFloat32Float32Map(f *testing.F) {
	f.Add(float32(math.MaxFloat32), float32(math.MaxFloat32))
	f.Fuzz(func(t *testing.T, k float32, v float32) {
		vm := map[float32]float32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzFloat32Float64Map(f *testing.F) {
	f.Add(float32(math.MaxFloat32), float64(math.MaxFloat64))
	f.Fuzz(func(t *testing.T, k float32, v float64) {
		vm := map[float32]float64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzFloat32BoolMap(f *testing.F) {
	f.Add(float32(math.MaxFloat32), bool(true))
	f.Add(float32(math.MaxFloat32), bool(false))
	f.Fuzz(func(t *testing.T, k float32, v bool) {
		vm := map[float32]bool{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzFloat32StringMap(f *testing.F) {
	f.Add(float32(math.MaxFloat32), string("abc"))
	f.Add(float32(math.MaxFloat32), string("xyz"))
	f.Fuzz(func(t *testing.T, k float32, v string) {
		vm := map[float32]string{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzFloat64IntMap(f *testing.F) {
	f.Add(float64(math.MaxFloat64), int(0))
	f.Add(float64(math.MaxFloat64), int(math.MinInt))
	f.Add(float64(math.MaxFloat64), int(math.MaxInt))
	f.Fuzz(func(t *testing.T, k float64, v int) {
		vm := map[float64]int{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzFloat64Int8Map(f *testing.F) {
	f.Add(float64(math.MaxFloat64), int8(0))
	f.Add(float64(math.MaxFloat64), int8(math.MinInt8))
	f.Add(float64(math.MaxFloat64), int8(math.MaxInt8))
	f.Fuzz(func(t *testing.T, k float64, v int8) {
		vm := map[float64]int8{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzFloat64Int16Map(f *testing.F) {
	f.Add(float64(math.MaxFloat64), int16(0))
	f.Add(float64(math.MaxFloat64), int16(math.MinInt16))
	f.Add(float64(math.MaxFloat64), int16(math.MaxInt16))
	f.Fuzz(func(t *testing.T, k float64, v int16) {
		vm := map[float64]int16{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzFloat64Int32Map(f *testing.F) {
	f.Add(float64(math.MaxFloat64), int32(0))
	f.Add(float64(math.MaxFloat64), int32(math.MinInt32))
	f.Add(float64(math.MaxFloat64), int32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, k float64, v int32) {
		vm := map[float64]int32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzFloat64Int64Map(f *testing.F) {
	f.Add(float64(math.MaxFloat64), int64(0))
	f.Add(float64(math.MaxFloat64), int64(math.MinInt64))
	f.Add(float64(math.MaxFloat64), int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k float64, v int64) {
		vm := map[float64]int64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzFloat64UintMap(f *testing.F) {
	f.Add(float64(math.MaxFloat64), uint(0))
	f.Add(float64(math.MaxFloat64), uint(math.MaxUint))
	f.Fuzz(func(t *testing.T, k float64, v uint) {
		vm := map[float64]uint{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzFloat64Uint8Map(f *testing.F) {
	f.Add(float64(math.MaxFloat64), uint8(0))
	f.Add(float64(math.MaxFloat64), uint8(math.MaxUint8))
	f.Fuzz(func(t *testing.T, k float64, v uint8) {
		vm := map[float64]uint8{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzFloat64Uint16Map(f *testing.F) {
	f.Add(float64(math.MaxFloat64), uint16(0))
	f.Add(float64(math.MaxFloat64), uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, k float64, v uint16) {
		vm := map[float64]uint16{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzFloat64Uint32Map(f *testing.F) {
	f.Add(float64(math.MaxFloat64), uint32(0))
	f.Add(float64(math.MaxFloat64), uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, k float64, v uint32) {
		vm := map[float64]uint32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzFloat64Uint64Map(f *testing.F) {
	f.Add(float64(math.MaxFloat64), uint64(0))
	f.Add(float64(math.MaxFloat64), uint64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k float64, v uint64) {
		vm := map[float64]uint64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzFloat64Float32Map(f *testing.F) {
	f.Add(float64(math.MaxFloat64), float32(math.MaxFloat32))
	f.Fuzz(func(t *testing.T, k float64, v float32) {
		vm := map[float64]float32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzFloat64Float64Map(f *testing.F) {
	f.Add(float64(math.MaxFloat64), float64(math.MaxFloat64))
	f.Fuzz(func(t *testing.T, k float64, v float64) {
		vm := map[float64]float64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzFloat64BoolMap(f *testing.F) {
	f.Add(float64(math.MaxFloat64), bool(true))
	f.Add(float64(math.MaxFloat64), bool(false))
	f.Fuzz(func(t *testing.T, k float64, v bool) {
		vm := map[float64]bool{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzFloat64StringMap(f *testing.F) {
	f.Add(float64(math.MaxFloat64), string("abc"))
	f.Add(float64(math.MaxFloat64), string("xyz"))
	f.Fuzz(func(t *testing.T, k float64, v string) {
		vm := map[float64]string{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzBoolIntMap(f *testing.F) {
	f.Add(bool(true), int(0))
	f.Add(bool(true), int(math.MinInt))
	f.Add(bool(true), int(math.MaxInt))
	f.Add(bool(false), int(0))
	f.Add(bool(false), int(math.MinInt))
	f.Add(bool(false), int(math.MaxInt))
	f.Fuzz(func(t *testing.T, k bool, v int) {
		vm := map[bool]int{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzBoolInt8Map(f *testing.F) {
	f.Add(bool(true), int8(0))
	f.Add(bool(true), int8(math.MinInt8))
	f.Add(bool(true), int8(math.MaxInt8))
	f.Add(bool(false), int8(0))
	f.Add(bool(false), int8(math.MinInt8))
	f.Add(bool(false), int8(math.MaxInt8))
	f.Fuzz(func(t *testing.T, k bool, v int8) {
		vm := map[bool]int8{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzBoolInt16Map(f *testing.F) {
	f.Add(bool(true), int16(0))
	f.Add(bool(true), int16(math.MinInt16))
	f.Add(bool(true), int16(math.MaxInt16))
	f.Add(bool(false), int16(0))
	f.Add(bool(false), int16(math.MinInt16))
	f.Add(bool(false), int16(math.MaxInt16))
	f.Fuzz(func(t *testing.T, k bool, v int16) {
		vm := map[bool]int16{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzBoolInt32Map(f *testing.F) {
	f.Add(bool(true), int32(0))
	f.Add(bool(true), int32(math.MinInt32))
	f.Add(bool(true), int32(math.MaxInt32))
	f.Add(bool(false), int32(0))
	f.Add(bool(false), int32(math.MinInt32))
	f.Add(bool(false), int32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, k bool, v int32) {
		vm := map[bool]int32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzBoolInt64Map(f *testing.F) {
	f.Add(bool(true), int64(0))
	f.Add(bool(true), int64(math.MinInt64))
	f.Add(bool(true), int64(math.MaxInt64))
	f.Add(bool(false), int64(0))
	f.Add(bool(false), int64(math.MinInt64))
	f.Add(bool(false), int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k bool, v int64) {
		vm := map[bool]int64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzBoolUintMap(f *testing.F) {
	f.Add(bool(true), uint(0))
	f.Add(bool(true), uint(math.MaxUint))
	f.Add(bool(false), uint(0))
	f.Add(bool(false), uint(math.MaxUint))
	f.Fuzz(func(t *testing.T, k bool, v uint) {
		vm := map[bool]uint{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzBoolUint8Map(f *testing.F) {
	f.Add(bool(true), uint8(0))
	f.Add(bool(true), uint8(math.MaxUint8))
	f.Add(bool(false), uint8(0))
	f.Add(bool(false), uint8(math.MaxUint8))
	f.Fuzz(func(t *testing.T, k bool, v uint8) {
		vm := map[bool]uint8{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzBoolUint16Map(f *testing.F) {
	f.Add(bool(true), uint16(0))
	f.Add(bool(true), uint16(math.MaxUint16))
	f.Add(bool(false), uint16(0))
	f.Add(bool(false), uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, k bool, v uint16) {
		vm := map[bool]uint16{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzBoolUint32Map(f *testing.F) {
	f.Add(bool(true), uint32(0))
	f.Add(bool(true), uint32(math.MaxUint32))
	f.Add(bool(false), uint32(0))
	f.Add(bool(false), uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, k bool, v uint32) {
		vm := map[bool]uint32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzBoolUint64Map(f *testing.F) {
	f.Add(bool(true), uint64(0))
	f.Add(bool(true), uint64(math.MaxInt64))
	f.Add(bool(false), uint64(0))
	f.Add(bool(false), uint64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k bool, v uint64) {
		vm := map[bool]uint64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzBoolFloat32Map(f *testing.F) {
	f.Add(bool(true), float32(math.MaxFloat32))
	f.Add(bool(false), float32(math.MaxFloat32))
	f.Fuzz(func(t *testing.T, k bool, v float32) {
		vm := map[bool]float32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzBoolFloat64Map(f *testing.F) {
	f.Add(bool(true), float64(math.MaxFloat64))
	f.Add(bool(false), float64(math.MaxFloat64))
	f.Fuzz(func(t *testing.T, k bool, v float64) {
		vm := map[bool]float64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzBoolBoolMap(f *testing.F) {
	f.Add(bool(true), bool(true))
	f.Add(bool(true), bool(false))
	f.Add(bool(false), bool(true))
	f.Add(bool(false), bool(false))
	f.Fuzz(func(t *testing.T, k bool, v bool) {
		vm := map[bool]bool{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzBoolStringMap(f *testing.F) {
	f.Add(bool(true), string("abc"))
	f.Add(bool(true), string("xyz"))
	f.Add(bool(false), string("abc"))
	f.Add(bool(false), string("xyz"))
	f.Fuzz(func(t *testing.T, k bool, v string) {
		vm := map[bool]string{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzStringIntMap(f *testing.F) {
	f.Add(string("abc"), int(0))
	f.Add(string("abc"), int(math.MinInt))
	f.Add(string("abc"), int(math.MaxInt))
	f.Add(string("xyz"), int(0))
	f.Add(string("xyz"), int(math.MinInt))
	f.Add(string("xyz"), int(math.MaxInt))
	f.Fuzz(func(t *testing.T, k string, v int) {
		vm := map[string]int{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzStringInt8Map(f *testing.F) {
	f.Add(string("abc"), int8(0))
	f.Add(string("abc"), int8(math.MinInt8))
	f.Add(string("abc"), int8(math.MaxInt8))
	f.Add(string("xyz"), int8(0))
	f.Add(string("xyz"), int8(math.MinInt8))
	f.Add(string("xyz"), int8(math.MaxInt8))
	f.Fuzz(func(t *testing.T, k string, v int8) {
		vm := map[string]int8{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzStringInt16Map(f *testing.F) {
	f.Add(string("abc"), int16(0))
	f.Add(string("abc"), int16(math.MinInt16))
	f.Add(string("abc"), int16(math.MaxInt16))
	f.Add(string("xyz"), int16(0))
	f.Add(string("xyz"), int16(math.MinInt16))
	f.Add(string("xyz"), int16(math.MaxInt16))
	f.Fuzz(func(t *testing.T, k string, v int16) {
		vm := map[string]int16{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzStringInt32Map(f *testing.F) {
	f.Add(string("abc"), int32(0))
	f.Add(string("abc"), int32(math.MinInt32))
	f.Add(string("abc"), int32(math.MaxInt32))
	f.Add(string("xyz"), int32(0))
	f.Add(string("xyz"), int32(math.MinInt32))
	f.Add(string("xyz"), int32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, k string, v int32) {
		vm := map[string]int32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzStringInt64Map(f *testing.F) {
	f.Add(string("abc"), int64(0))
	f.Add(string("abc"), int64(math.MinInt64))
	f.Add(string("abc"), int64(math.MaxInt64))
	f.Add(string("xyz"), int64(0))
	f.Add(string("xyz"), int64(math.MinInt64))
	f.Add(string("xyz"), int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k string, v int64) {
		vm := map[string]int64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzStringUintMap(f *testing.F) {
	f.Add(string("abc"), uint(0))
	f.Add(string("abc"), uint(math.MaxUint))
	f.Add(string("xyz"), uint(0))
	f.Add(string("xyz"), uint(math.MaxUint))
	f.Fuzz(func(t *testing.T, k string, v uint) {
		vm := map[string]uint{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzStringUint8Map(f *testing.F) {
	f.Add(string("abc"), uint8(0))
	f.Add(string("abc"), uint8(math.MaxUint8))
	f.Add(string("xyz"), uint8(0))
	f.Add(string("xyz"), uint8(math.MaxUint8))
	f.Fuzz(func(t *testing.T, k string, v uint8) {
		vm := map[string]uint8{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzStringUint16Map(f *testing.F) {
	f.Add(string("abc"), uint16(0))
	f.Add(string("abc"), uint16(math.MaxUint16))
	f.Add(string("xyz"), uint16(0))
	f.Add(string("xyz"), uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, k string, v uint16) {
		vm := map[string]uint16{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzStringUint32Map(f *testing.F) {
	f.Add(string("abc"), uint32(0))
	f.Add(string("abc"), uint32(math.MaxUint32))
	f.Add(string("xyz"), uint32(0))
	f.Add(string("xyz"), uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, k string, v uint32) {
		vm := map[string]uint32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzStringUint64Map(f *testing.F) {
	f.Add(string("abc"), uint64(0))
	f.Add(string("abc"), uint64(math.MaxInt64))
	f.Add(string("xyz"), uint64(0))
	f.Add(string("xyz"), uint64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k string, v uint64) {
		vm := map[string]uint64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzStringFloat32Map(f *testing.F) {
	f.Add(string("abc"), float32(math.MaxFloat32))
	f.Add(string("xyz"), float32(math.MaxFloat32))
	f.Fuzz(func(t *testing.T, k string, v float32) {
		vm := map[string]float32{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzStringFloat64Map(f *testing.F) {
	f.Add(string("abc"), float64(math.MaxFloat64))
	f.Add(string("xyz"), float64(math.MaxFloat64))
	f.Fuzz(func(t *testing.T, k string, v float64) {
		vm := map[string]float64{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzStringBoolMap(f *testing.F) {
	f.Add(string("abc"), bool(true))
	f.Add(string("abc"), bool(false))
	f.Add(string("xyz"), bool(true))
	f.Add(string("xyz"), bool(false))
	f.Fuzz(func(t *testing.T, k string, v bool) {
		vm := map[string]bool{k: v}
		fuzzTest(t, vm)
	})
}

// nolint: dupl
func FuzzStringStringMap(f *testing.F) {
	f.Add(string("abc"), string("abc"))
	f.Add(string("abc"), string("xyz"))
	f.Add(string("xyz"), string("abc"))
	f.Add(string("xyz"), string("xyz"))
	f.Fuzz(func(t *testing.T, k string, v string) {
		vm := map[string]string{k: v}
		fuzzTest(t, vm)
	})
}
