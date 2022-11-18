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

func FuzzIntData(f *testing.F) {
	f.Add(int(0))
	f.Add(int(math.MinInt))
	f.Add(int(math.MaxInt))
	f.Fuzz(func(t *testing.T, v int) {
		fuzzTest(t, v)
	})
}

func FuzzInt8Data(f *testing.F) {
	f.Add(int8(0))
	f.Add(int8(math.MinInt8))
	f.Add(int8(math.MaxInt8))
	f.Fuzz(func(t *testing.T, v int8) {
		fuzzTest(t, v)
	})
}

func FuzzInt16Data(f *testing.F) {
	f.Add(int16(0))
	f.Add(int16(math.MinInt16))
	f.Add(int16(math.MaxInt16))
	f.Fuzz(func(t *testing.T, v int16) {
		fuzzTest(t, v)
	})
}

func FuzzInt32Data(f *testing.F) {
	f.Add(int32(0))
	f.Add(int32(math.MinInt32))
	f.Add(int32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, v int32) {
		fuzzTest(t, v)
	})
}

func FuzzInt64Data(f *testing.F) {
	f.Add(int64(0))
	f.Add(int64(math.MinInt64))
	f.Add(int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, v int64) {
		fuzzTest(t, v)
	})
}

func FuzzUintData(f *testing.F) {
	f.Add(uint(0))
	f.Add(uint(math.MaxUint))
	f.Fuzz(func(t *testing.T, v uint) {
		fuzzTest(t, v)
	})
}

func FuzzUint8Data(f *testing.F) {
	f.Add(uint8(0))
	f.Add(uint8(math.MaxUint8))
	f.Fuzz(func(t *testing.T, v uint8) {
		fuzzTest(t, v)
	})
}

func FuzzUint16Data(f *testing.F) {
	f.Add(uint16(0))
	f.Add(uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, v uint16) {
		fuzzTest(t, v)
	})
}

func FuzzUint32Data(f *testing.F) {
	f.Add(uint32(0))
	f.Add(uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, v uint32) {
		fuzzTest(t, v)
	})
}

func FuzzUint64Data(f *testing.F) {
	f.Add(uint64(0))
	f.Add(uint64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, v uint64) {
		fuzzTest(t, v)
	})
}

func FuzzFloat32Data(f *testing.F) {
	f.Add(float32(-math.MaxFloat32))
	f.Add(float32(0))
	f.Add(float32(math.MaxFloat32))
	f.Fuzz(func(t *testing.T, v float32) {
		fuzzTest(t, v)
	})
}

func FuzzFloat64Data(f *testing.F) {
	f.Add(float64(-math.MaxFloat32))
	f.Add(float64(0))
	f.Add(float64(math.MaxFloat64))
	f.Fuzz(func(t *testing.T, v float64) {
		fuzzTest(t, v)
	})
}

func FuzzBoolData(f *testing.F) {
	f.Add(bool(true))
	f.Add(bool(false))
	f.Fuzz(func(t *testing.T, v bool) {
		fuzzTest(t, v)
	})
}

func FuzzStringData(f *testing.F) {
	f.Add(string("a"))
	f.Add(string("ab"))
	f.Add(string("abc"))
	f.Fuzz(func(t *testing.T, v string) {
		fuzzTest(t, v)
	})
}

func FuzzByteData(f *testing.F) {
	f.Add([]byte("x"))
	f.Add([]byte("xy"))
	f.Add([]byte("xyz"))
	f.Fuzz(func(t *testing.T, v []byte) {
		fuzzTest(t, v)
	})
}

func FuzzIntArray(f *testing.F) {
	f.Add(int(0))
	f.Add(int(math.MinInt))
	f.Add(int(math.MaxInt))
	f.Fuzz(func(t *testing.T, v int) {
		va := []int{}
		fuzzTest(t, va)
		va = []int{v}
		fuzzTest(t, va)
		va = []int{v, v}
		fuzzTest(t, va)
		va = []int{v, v, v}
		fuzzTest(t, va)
	})
}

func FuzzInt8Array(f *testing.F) {
	f.Add(int8(0))
	f.Add(int8(math.MinInt8))
	f.Add(int8(math.MaxInt8))
	f.Fuzz(func(t *testing.T, v int8) {
		va := []int8{}
		fuzzTest(t, va)
		va = []int8{v}
		fuzzTest(t, va)
		va = []int8{v, v}
		fuzzTest(t, va)
		va = []int8{v, v, v}
		fuzzTest(t, va)
	})
}

func FuzzInt16Array(f *testing.F) {
	f.Add(int16(0))
	f.Add(int16(math.MinInt16))
	f.Add(int16(math.MaxInt16))
	f.Fuzz(func(t *testing.T, v int16) {
		va := []int16{}
		fuzzTest(t, va)
		va = []int16{v}
		fuzzTest(t, va)
		va = []int16{v, v}
		fuzzTest(t, va)
		va = []int16{v, v, v}
		fuzzTest(t, va)
	})
}

func FuzzInt32Array(f *testing.F) {
	f.Add(int32(0))
	f.Add(int32(math.MinInt32))
	f.Add(int32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, v int32) {
		va := []int32{}
		fuzzTest(t, va)
		va = []int32{v}
		fuzzTest(t, va)
		va = []int32{v, v}
		fuzzTest(t, va)
		va = []int32{v, v, v}
		fuzzTest(t, va)
	})
}

func FuzzInt64Array(f *testing.F) {
	f.Add(int64(0))
	f.Add(int64(math.MinInt64))
	f.Add(int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, v int64) {
		va := []int64{}
		fuzzTest(t, va)
		va = []int64{v}
		fuzzTest(t, va)
		va = []int64{v, v}
		fuzzTest(t, va)
		va = []int64{v, v, v}
		fuzzTest(t, va)
	})
}

func FuzzUintArray(f *testing.F) {
	f.Add(uint(0))
	f.Add(uint(math.MaxUint))
	f.Fuzz(func(t *testing.T, v uint) {
		va := []uint{}
		fuzzTest(t, va)
		va = []uint{v}
		fuzzTest(t, va)
		va = []uint{v, v}
		fuzzTest(t, va)
		va = []uint{v, v, v}
		fuzzTest(t, va)
	})
}

func FuzzUint8Array(f *testing.F) {
	f.Add(uint8(0))
	f.Add(uint8(math.MaxUint8))
	f.Fuzz(func(t *testing.T, v uint8) {
		va := []uint8{}
		fuzzTest(t, va)
		va = []uint8{v}
		fuzzTest(t, va)
		va = []uint8{v, v}
		fuzzTest(t, va)
		va = []uint8{v, v, v}
		fuzzTest(t, va)
	})
}

func FuzzUint16Array(f *testing.F) {
	f.Add(uint16(0))
	f.Add(uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, v uint16) {
		va := []uint16{}
		fuzzTest(t, va)
		va = []uint16{v}
		fuzzTest(t, va)
		va = []uint16{v, v}
		fuzzTest(t, va)
		va = []uint16{v, v, v}
		fuzzTest(t, va)
	})
}

func FuzzUint32Array(f *testing.F) {
	f.Add(uint32(0))
	f.Add(uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, v uint32) {
		va := []uint32{}
		fuzzTest(t, va)
		va = []uint32{v}
		fuzzTest(t, va)
		va = []uint32{v, v}
		fuzzTest(t, va)
		va = []uint32{v, v, v}
		fuzzTest(t, va)
	})
}

func FuzzUint64Array(f *testing.F) {
	f.Add(uint64(0))
	f.Add(uint64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, v uint64) {
		va := []uint64{}
		fuzzTest(t, va)
		va = []uint64{v}
		fuzzTest(t, va)
		va = []uint64{v, v}
		fuzzTest(t, va)
		va = []uint64{v, v, v}
		fuzzTest(t, va)
	})
}

func FuzzFloat32Array(f *testing.F) {
	f.Add(float32(-math.MaxFloat32))
	f.Add(float32(0))
	f.Add(float32(math.MaxFloat32))
	f.Fuzz(func(t *testing.T, v float32) {
		va := []float32{}
		fuzzTest(t, va)
		va = []float32{v}
		fuzzTest(t, va)
		va = []float32{v, v}
		fuzzTest(t, va)
		va = []float32{v, v, v}
		fuzzTest(t, va)
	})
}

func FuzzFloat64Array(f *testing.F) {
	f.Add(float64(-math.MaxFloat32))
	f.Add(float64(0))
	f.Add(float64(math.MaxFloat64))
	f.Fuzz(func(t *testing.T, v float64) {
		va := []float64{}
		fuzzTest(t, va)
		va = []float64{v}
		fuzzTest(t, va)
		va = []float64{v, v}
		fuzzTest(t, va)
		va = []float64{v, v, v}
		fuzzTest(t, va)
	})
}

func FuzzBoolArray(f *testing.F) {
	f.Add(bool(true))
	f.Add(bool(false))
	f.Fuzz(func(t *testing.T, v bool) {
		va := []bool{}
		fuzzTest(t, va)
		va = []bool{v}
		fuzzTest(t, va)
		va = []bool{v, v}
		fuzzTest(t, va)
		va = []bool{v, v, v}
		fuzzTest(t, va)
	})
}

func FuzzStringArray(f *testing.F) {
	f.Add(string("a"))
	f.Add(string("ab"))
	f.Add(string("abc"))
	f.Fuzz(func(t *testing.T, v string) {
		va := []string{}
		fuzzTest(t, va)
		va = []string{v}
		fuzzTest(t, va)
		va = []string{v, v}
		fuzzTest(t, va)
		va = []string{v, v, v}
		fuzzTest(t, va)
	})
}

func FuzzByteArray(f *testing.F) {
	f.Add([]byte("x"))
	f.Add([]byte("xy"))
	f.Add([]byte("xyz"))
	f.Fuzz(func(t *testing.T, v []byte) {
		va := [][]byte{}
		fuzzTest(t, va)
		va = [][]byte{v}
		fuzzTest(t, va)
		va = [][]byte{v, v}
		fuzzTest(t, va)
		va = [][]byte{v, v, v}
		fuzzTest(t, va)
	})
}

// nolint: dupl, maligned
func FuzzIntIntStruct(f *testing.F) {
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
		vs1 := struct {
			Elem1 int
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int
			Elem2 int
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int
			Elem2 int
			Elem3 int
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzIntInt8Struct(f *testing.F) {
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
		vs1 := struct {
			Elem1 int
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int
			Elem2 int8
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int
			Elem2 int8
			Elem3 int
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzIntInt16Struct(f *testing.F) {
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
		vs1 := struct {
			Elem1 int
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int
			Elem2 int16
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int
			Elem2 int16
			Elem3 int
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzIntInt32Struct(f *testing.F) {
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
		vs1 := struct {
			Elem1 int
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int
			Elem2 int32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int
			Elem2 int32
			Elem3 int
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzIntInt64Struct(f *testing.F) {
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
		vs1 := struct {
			Elem1 int
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int
			Elem2 int64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int
			Elem2 int64
			Elem3 int
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzIntUintStruct(f *testing.F) {
	f.Add(int(0), uint(0))
	f.Add(int(0), uint(math.MaxUint))
	f.Add(int(math.MinInt), uint(0))
	f.Add(int(math.MinInt), uint(math.MaxUint))
	f.Add(int(math.MaxInt), uint(0))
	f.Add(int(math.MaxInt), uint(math.MaxUint))
	f.Fuzz(func(t *testing.T, k int, v uint) {
		vs1 := struct {
			Elem1 int
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int
			Elem2 uint
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int
			Elem2 uint
			Elem3 int
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzIntUint8Struct(f *testing.F) {
	f.Add(int(0), uint8(0))
	f.Add(int(0), uint8(math.MaxUint8))
	f.Add(int(math.MinInt), uint8(0))
	f.Add(int(math.MinInt), uint8(math.MaxUint8))
	f.Add(int(math.MaxInt), uint8(0))
	f.Add(int(math.MaxInt), uint8(math.MaxUint8))
	f.Fuzz(func(t *testing.T, k int, v uint8) {
		vs1 := struct {
			Elem1 int
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int
			Elem2 uint8
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int
			Elem2 uint8
			Elem3 int
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzIntUint16Struct(f *testing.F) {
	f.Add(int(0), uint16(0))
	f.Add(int(0), uint16(math.MaxUint16))
	f.Add(int(math.MinInt), uint16(0))
	f.Add(int(math.MinInt), uint16(math.MaxUint16))
	f.Add(int(math.MaxInt), uint16(0))
	f.Add(int(math.MaxInt), uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, k int, v uint16) {
		vs1 := struct {
			Elem1 int
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int
			Elem2 uint16
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int
			Elem2 uint16
			Elem3 int
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzIntUint32Struct(f *testing.F) {
	f.Add(int(0), uint32(0))
	f.Add(int(0), uint32(math.MaxUint32))
	f.Add(int(math.MinInt), uint32(0))
	f.Add(int(math.MinInt), uint32(math.MaxUint32))
	f.Add(int(math.MaxInt), uint32(0))
	f.Add(int(math.MaxInt), uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, k int, v uint32) {
		vs1 := struct {
			Elem1 int
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int
			Elem2 uint32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int
			Elem2 uint32
			Elem3 int
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzIntUint64Struct(f *testing.F) {
	f.Add(int(0), uint64(0))
	f.Add(int(0), uint64(math.MaxInt64))
	f.Add(int(math.MinInt), uint64(0))
	f.Add(int(math.MinInt), uint64(math.MaxInt64))
	f.Add(int(math.MaxInt), uint64(0))
	f.Add(int(math.MaxInt), uint64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k int, v uint64) {
		vs1 := struct {
			Elem1 int
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int
			Elem2 uint64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int
			Elem2 uint64
			Elem3 int
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzIntFloat32Struct(f *testing.F) {
	f.Add(int(0), float32(-math.MaxFloat32))
	f.Add(int(0), float32(0))
	f.Add(int(0), float32(math.MaxFloat32))
	f.Add(int(math.MinInt), float32(-math.MaxFloat32))
	f.Add(int(math.MinInt), float32(0))
	f.Add(int(math.MinInt), float32(math.MaxFloat32))
	f.Add(int(math.MaxInt), float32(-math.MaxFloat32))
	f.Add(int(math.MaxInt), float32(0))
	f.Add(int(math.MaxInt), float32(math.MaxFloat32))
	f.Fuzz(func(t *testing.T, k int, v float32) {
		vs1 := struct {
			Elem1 int
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int
			Elem2 float32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int
			Elem2 float32
			Elem3 int
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzIntFloat64Struct(f *testing.F) {
	f.Add(int(0), float64(-math.MaxFloat32))
	f.Add(int(0), float64(0))
	f.Add(int(0), float64(math.MaxFloat64))
	f.Add(int(math.MinInt), float64(-math.MaxFloat32))
	f.Add(int(math.MinInt), float64(0))
	f.Add(int(math.MinInt), float64(math.MaxFloat64))
	f.Add(int(math.MaxInt), float64(-math.MaxFloat32))
	f.Add(int(math.MaxInt), float64(0))
	f.Add(int(math.MaxInt), float64(math.MaxFloat64))
	f.Fuzz(func(t *testing.T, k int, v float64) {
		vs1 := struct {
			Elem1 int
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int
			Elem2 float64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int
			Elem2 float64
			Elem3 int
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzIntBoolStruct(f *testing.F) {
	f.Add(int(0), bool(true))
	f.Add(int(0), bool(false))
	f.Add(int(math.MinInt), bool(true))
	f.Add(int(math.MinInt), bool(false))
	f.Add(int(math.MaxInt), bool(true))
	f.Add(int(math.MaxInt), bool(false))
	f.Fuzz(func(t *testing.T, k int, v bool) {
		vs1 := struct {
			Elem1 int
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int
			Elem2 bool
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int
			Elem2 bool
			Elem3 int
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzIntStringStruct(f *testing.F) {
	f.Add(int(0), string("a"))
	f.Add(int(0), string("ab"))
	f.Add(int(0), string("abc"))
	f.Add(int(math.MinInt), string("a"))
	f.Add(int(math.MinInt), string("ab"))
	f.Add(int(math.MinInt), string("abc"))
	f.Add(int(math.MaxInt), string("a"))
	f.Add(int(math.MaxInt), string("ab"))
	f.Add(int(math.MaxInt), string("abc"))
	f.Fuzz(func(t *testing.T, k int, v string) {
		vs1 := struct {
			Elem1 int
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int
			Elem2 string
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int
			Elem2 string
			Elem3 int
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzIntByteStruct(f *testing.F) {
	f.Add(int(0), []byte("x"))
	f.Add(int(0), []byte("xy"))
	f.Add(int(0), []byte("xyz"))
	f.Add(int(math.MinInt), []byte("x"))
	f.Add(int(math.MinInt), []byte("xy"))
	f.Add(int(math.MinInt), []byte("xyz"))
	f.Add(int(math.MaxInt), []byte("x"))
	f.Add(int(math.MaxInt), []byte("xy"))
	f.Add(int(math.MaxInt), []byte("xyz"))
	f.Fuzz(func(t *testing.T, k int, v []byte) {
		vs1 := struct {
			Elem1 int
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int
			Elem2 []byte
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int
			Elem2 []byte
			Elem3 int
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt8IntStruct(f *testing.F) {
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
		vs1 := struct {
			Elem1 int8
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int8
			Elem2 int
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int8
			Elem2 int
			Elem3 int8
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt8Int8Struct(f *testing.F) {
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
		vs1 := struct {
			Elem1 int8
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int8
			Elem2 int8
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int8
			Elem2 int8
			Elem3 int8
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt8Int16Struct(f *testing.F) {
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
		vs1 := struct {
			Elem1 int8
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int8
			Elem2 int16
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int8
			Elem2 int16
			Elem3 int8
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt8Int32Struct(f *testing.F) {
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
		vs1 := struct {
			Elem1 int8
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int8
			Elem2 int32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int8
			Elem2 int32
			Elem3 int8
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt8Int64Struct(f *testing.F) {
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
		vs1 := struct {
			Elem1 int8
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int8
			Elem2 int64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int8
			Elem2 int64
			Elem3 int8
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt8UintStruct(f *testing.F) {
	f.Add(int8(0), uint(0))
	f.Add(int8(0), uint(math.MaxUint))
	f.Add(int8(math.MinInt8), uint(0))
	f.Add(int8(math.MinInt8), uint(math.MaxUint))
	f.Add(int8(math.MaxInt8), uint(0))
	f.Add(int8(math.MaxInt8), uint(math.MaxUint))
	f.Fuzz(func(t *testing.T, k int8, v uint) {
		vs1 := struct {
			Elem1 int8
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int8
			Elem2 uint
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int8
			Elem2 uint
			Elem3 int8
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt8Uint8Struct(f *testing.F) {
	f.Add(int8(0), uint8(0))
	f.Add(int8(0), uint8(math.MaxUint8))
	f.Add(int8(math.MinInt8), uint8(0))
	f.Add(int8(math.MinInt8), uint8(math.MaxUint8))
	f.Add(int8(math.MaxInt8), uint8(0))
	f.Add(int8(math.MaxInt8), uint8(math.MaxUint8))
	f.Fuzz(func(t *testing.T, k int8, v uint8) {
		vs1 := struct {
			Elem1 int8
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int8
			Elem2 uint8
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int8
			Elem2 uint8
			Elem3 int8
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt8Uint16Struct(f *testing.F) {
	f.Add(int8(0), uint16(0))
	f.Add(int8(0), uint16(math.MaxUint16))
	f.Add(int8(math.MinInt8), uint16(0))
	f.Add(int8(math.MinInt8), uint16(math.MaxUint16))
	f.Add(int8(math.MaxInt8), uint16(0))
	f.Add(int8(math.MaxInt8), uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, k int8, v uint16) {
		vs1 := struct {
			Elem1 int8
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int8
			Elem2 uint16
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int8
			Elem2 uint16
			Elem3 int8
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt8Uint32Struct(f *testing.F) {
	f.Add(int8(0), uint32(0))
	f.Add(int8(0), uint32(math.MaxUint32))
	f.Add(int8(math.MinInt8), uint32(0))
	f.Add(int8(math.MinInt8), uint32(math.MaxUint32))
	f.Add(int8(math.MaxInt8), uint32(0))
	f.Add(int8(math.MaxInt8), uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, k int8, v uint32) {
		vs1 := struct {
			Elem1 int8
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int8
			Elem2 uint32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int8
			Elem2 uint32
			Elem3 int8
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt8Uint64Struct(f *testing.F) {
	f.Add(int8(0), uint64(0))
	f.Add(int8(0), uint64(math.MaxInt64))
	f.Add(int8(math.MinInt8), uint64(0))
	f.Add(int8(math.MinInt8), uint64(math.MaxInt64))
	f.Add(int8(math.MaxInt8), uint64(0))
	f.Add(int8(math.MaxInt8), uint64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k int8, v uint64) {
		vs1 := struct {
			Elem1 int8
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int8
			Elem2 uint64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int8
			Elem2 uint64
			Elem3 int8
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt8Float32Struct(f *testing.F) {
	f.Add(int8(0), float32(-math.MaxFloat32))
	f.Add(int8(0), float32(0))
	f.Add(int8(0), float32(math.MaxFloat32))
	f.Add(int8(math.MinInt8), float32(-math.MaxFloat32))
	f.Add(int8(math.MinInt8), float32(0))
	f.Add(int8(math.MinInt8), float32(math.MaxFloat32))
	f.Add(int8(math.MaxInt8), float32(-math.MaxFloat32))
	f.Add(int8(math.MaxInt8), float32(0))
	f.Add(int8(math.MaxInt8), float32(math.MaxFloat32))
	f.Fuzz(func(t *testing.T, k int8, v float32) {
		vs1 := struct {
			Elem1 int8
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int8
			Elem2 float32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int8
			Elem2 float32
			Elem3 int8
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt8Float64Struct(f *testing.F) {
	f.Add(int8(0), float64(-math.MaxFloat32))
	f.Add(int8(0), float64(0))
	f.Add(int8(0), float64(math.MaxFloat64))
	f.Add(int8(math.MinInt8), float64(-math.MaxFloat32))
	f.Add(int8(math.MinInt8), float64(0))
	f.Add(int8(math.MinInt8), float64(math.MaxFloat64))
	f.Add(int8(math.MaxInt8), float64(-math.MaxFloat32))
	f.Add(int8(math.MaxInt8), float64(0))
	f.Add(int8(math.MaxInt8), float64(math.MaxFloat64))
	f.Fuzz(func(t *testing.T, k int8, v float64) {
		vs1 := struct {
			Elem1 int8
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int8
			Elem2 float64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int8
			Elem2 float64
			Elem3 int8
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt8BoolStruct(f *testing.F) {
	f.Add(int8(0), bool(true))
	f.Add(int8(0), bool(false))
	f.Add(int8(math.MinInt8), bool(true))
	f.Add(int8(math.MinInt8), bool(false))
	f.Add(int8(math.MaxInt8), bool(true))
	f.Add(int8(math.MaxInt8), bool(false))
	f.Fuzz(func(t *testing.T, k int8, v bool) {
		vs1 := struct {
			Elem1 int8
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int8
			Elem2 bool
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int8
			Elem2 bool
			Elem3 int8
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt8StringStruct(f *testing.F) {
	f.Add(int8(0), string("a"))
	f.Add(int8(0), string("ab"))
	f.Add(int8(0), string("abc"))
	f.Add(int8(math.MinInt8), string("a"))
	f.Add(int8(math.MinInt8), string("ab"))
	f.Add(int8(math.MinInt8), string("abc"))
	f.Add(int8(math.MaxInt8), string("a"))
	f.Add(int8(math.MaxInt8), string("ab"))
	f.Add(int8(math.MaxInt8), string("abc"))
	f.Fuzz(func(t *testing.T, k int8, v string) {
		vs1 := struct {
			Elem1 int8
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int8
			Elem2 string
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int8
			Elem2 string
			Elem3 int8
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt8ByteStruct(f *testing.F) {
	f.Add(int8(0), []byte("x"))
	f.Add(int8(0), []byte("xy"))
	f.Add(int8(0), []byte("xyz"))
	f.Add(int8(math.MinInt8), []byte("x"))
	f.Add(int8(math.MinInt8), []byte("xy"))
	f.Add(int8(math.MinInt8), []byte("xyz"))
	f.Add(int8(math.MaxInt8), []byte("x"))
	f.Add(int8(math.MaxInt8), []byte("xy"))
	f.Add(int8(math.MaxInt8), []byte("xyz"))
	f.Fuzz(func(t *testing.T, k int8, v []byte) {
		vs1 := struct {
			Elem1 int8
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int8
			Elem2 []byte
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int8
			Elem2 []byte
			Elem3 int8
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt16IntStruct(f *testing.F) {
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
		vs1 := struct {
			Elem1 int16
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int16
			Elem2 int
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int16
			Elem2 int
			Elem3 int16
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt16Int8Struct(f *testing.F) {
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
		vs1 := struct {
			Elem1 int16
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int16
			Elem2 int8
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int16
			Elem2 int8
			Elem3 int16
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt16Int16Struct(f *testing.F) {
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
		vs1 := struct {
			Elem1 int16
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int16
			Elem2 int16
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int16
			Elem2 int16
			Elem3 int16
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt16Int32Struct(f *testing.F) {
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
		vs1 := struct {
			Elem1 int16
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int16
			Elem2 int32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int16
			Elem2 int32
			Elem3 int16
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt16Int64Struct(f *testing.F) {
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
		vs1 := struct {
			Elem1 int16
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int16
			Elem2 int64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int16
			Elem2 int64
			Elem3 int16
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt16UintStruct(f *testing.F) {
	f.Add(int16(0), uint(0))
	f.Add(int16(0), uint(math.MaxUint))
	f.Add(int16(math.MinInt16), uint(0))
	f.Add(int16(math.MinInt16), uint(math.MaxUint))
	f.Add(int16(math.MaxInt16), uint(0))
	f.Add(int16(math.MaxInt16), uint(math.MaxUint))
	f.Fuzz(func(t *testing.T, k int16, v uint) {
		vs1 := struct {
			Elem1 int16
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int16
			Elem2 uint
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int16
			Elem2 uint
			Elem3 int16
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt16Uint8Struct(f *testing.F) {
	f.Add(int16(0), uint8(0))
	f.Add(int16(0), uint8(math.MaxUint8))
	f.Add(int16(math.MinInt16), uint8(0))
	f.Add(int16(math.MinInt16), uint8(math.MaxUint8))
	f.Add(int16(math.MaxInt16), uint8(0))
	f.Add(int16(math.MaxInt16), uint8(math.MaxUint8))
	f.Fuzz(func(t *testing.T, k int16, v uint8) {
		vs1 := struct {
			Elem1 int16
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int16
			Elem2 uint8
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int16
			Elem2 uint8
			Elem3 int16
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt16Uint16Struct(f *testing.F) {
	f.Add(int16(0), uint16(0))
	f.Add(int16(0), uint16(math.MaxUint16))
	f.Add(int16(math.MinInt16), uint16(0))
	f.Add(int16(math.MinInt16), uint16(math.MaxUint16))
	f.Add(int16(math.MaxInt16), uint16(0))
	f.Add(int16(math.MaxInt16), uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, k int16, v uint16) {
		vs1 := struct {
			Elem1 int16
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int16
			Elem2 uint16
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int16
			Elem2 uint16
			Elem3 int16
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt16Uint32Struct(f *testing.F) {
	f.Add(int16(0), uint32(0))
	f.Add(int16(0), uint32(math.MaxUint32))
	f.Add(int16(math.MinInt16), uint32(0))
	f.Add(int16(math.MinInt16), uint32(math.MaxUint32))
	f.Add(int16(math.MaxInt16), uint32(0))
	f.Add(int16(math.MaxInt16), uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, k int16, v uint32) {
		vs1 := struct {
			Elem1 int16
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int16
			Elem2 uint32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int16
			Elem2 uint32
			Elem3 int16
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt16Uint64Struct(f *testing.F) {
	f.Add(int16(0), uint64(0))
	f.Add(int16(0), uint64(math.MaxInt64))
	f.Add(int16(math.MinInt16), uint64(0))
	f.Add(int16(math.MinInt16), uint64(math.MaxInt64))
	f.Add(int16(math.MaxInt16), uint64(0))
	f.Add(int16(math.MaxInt16), uint64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k int16, v uint64) {
		vs1 := struct {
			Elem1 int16
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int16
			Elem2 uint64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int16
			Elem2 uint64
			Elem3 int16
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt16Float32Struct(f *testing.F) {
	f.Add(int16(0), float32(-math.MaxFloat32))
	f.Add(int16(0), float32(0))
	f.Add(int16(0), float32(math.MaxFloat32))
	f.Add(int16(math.MinInt16), float32(-math.MaxFloat32))
	f.Add(int16(math.MinInt16), float32(0))
	f.Add(int16(math.MinInt16), float32(math.MaxFloat32))
	f.Add(int16(math.MaxInt16), float32(-math.MaxFloat32))
	f.Add(int16(math.MaxInt16), float32(0))
	f.Add(int16(math.MaxInt16), float32(math.MaxFloat32))
	f.Fuzz(func(t *testing.T, k int16, v float32) {
		vs1 := struct {
			Elem1 int16
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int16
			Elem2 float32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int16
			Elem2 float32
			Elem3 int16
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt16Float64Struct(f *testing.F) {
	f.Add(int16(0), float64(-math.MaxFloat32))
	f.Add(int16(0), float64(0))
	f.Add(int16(0), float64(math.MaxFloat64))
	f.Add(int16(math.MinInt16), float64(-math.MaxFloat32))
	f.Add(int16(math.MinInt16), float64(0))
	f.Add(int16(math.MinInt16), float64(math.MaxFloat64))
	f.Add(int16(math.MaxInt16), float64(-math.MaxFloat32))
	f.Add(int16(math.MaxInt16), float64(0))
	f.Add(int16(math.MaxInt16), float64(math.MaxFloat64))
	f.Fuzz(func(t *testing.T, k int16, v float64) {
		vs1 := struct {
			Elem1 int16
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int16
			Elem2 float64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int16
			Elem2 float64
			Elem3 int16
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt16BoolStruct(f *testing.F) {
	f.Add(int16(0), bool(true))
	f.Add(int16(0), bool(false))
	f.Add(int16(math.MinInt16), bool(true))
	f.Add(int16(math.MinInt16), bool(false))
	f.Add(int16(math.MaxInt16), bool(true))
	f.Add(int16(math.MaxInt16), bool(false))
	f.Fuzz(func(t *testing.T, k int16, v bool) {
		vs1 := struct {
			Elem1 int16
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int16
			Elem2 bool
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int16
			Elem2 bool
			Elem3 int16
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt16StringStruct(f *testing.F) {
	f.Add(int16(0), string("a"))
	f.Add(int16(0), string("ab"))
	f.Add(int16(0), string("abc"))
	f.Add(int16(math.MinInt16), string("a"))
	f.Add(int16(math.MinInt16), string("ab"))
	f.Add(int16(math.MinInt16), string("abc"))
	f.Add(int16(math.MaxInt16), string("a"))
	f.Add(int16(math.MaxInt16), string("ab"))
	f.Add(int16(math.MaxInt16), string("abc"))
	f.Fuzz(func(t *testing.T, k int16, v string) {
		vs1 := struct {
			Elem1 int16
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int16
			Elem2 string
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int16
			Elem2 string
			Elem3 int16
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt16ByteStruct(f *testing.F) {
	f.Add(int16(0), []byte("x"))
	f.Add(int16(0), []byte("xy"))
	f.Add(int16(0), []byte("xyz"))
	f.Add(int16(math.MinInt16), []byte("x"))
	f.Add(int16(math.MinInt16), []byte("xy"))
	f.Add(int16(math.MinInt16), []byte("xyz"))
	f.Add(int16(math.MaxInt16), []byte("x"))
	f.Add(int16(math.MaxInt16), []byte("xy"))
	f.Add(int16(math.MaxInt16), []byte("xyz"))
	f.Fuzz(func(t *testing.T, k int16, v []byte) {
		vs1 := struct {
			Elem1 int16
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int16
			Elem2 []byte
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int16
			Elem2 []byte
			Elem3 int16
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt32IntStruct(f *testing.F) {
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
		vs1 := struct {
			Elem1 int32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int32
			Elem2 int
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int32
			Elem2 int
			Elem3 int32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt32Int8Struct(f *testing.F) {
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
		vs1 := struct {
			Elem1 int32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int32
			Elem2 int8
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int32
			Elem2 int8
			Elem3 int32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt32Int16Struct(f *testing.F) {
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
		vs1 := struct {
			Elem1 int32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int32
			Elem2 int16
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int32
			Elem2 int16
			Elem3 int32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt32Int32Struct(f *testing.F) {
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
		vs1 := struct {
			Elem1 int32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int32
			Elem2 int32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int32
			Elem2 int32
			Elem3 int32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt32Int64Struct(f *testing.F) {
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
		vs1 := struct {
			Elem1 int32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int32
			Elem2 int64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int32
			Elem2 int64
			Elem3 int32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt32UintStruct(f *testing.F) {
	f.Add(int32(0), uint(0))
	f.Add(int32(0), uint(math.MaxUint))
	f.Add(int32(math.MinInt32), uint(0))
	f.Add(int32(math.MinInt32), uint(math.MaxUint))
	f.Add(int32(math.MaxInt32), uint(0))
	f.Add(int32(math.MaxInt32), uint(math.MaxUint))
	f.Fuzz(func(t *testing.T, k int32, v uint) {
		vs1 := struct {
			Elem1 int32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int32
			Elem2 uint
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int32
			Elem2 uint
			Elem3 int32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt32Uint8Struct(f *testing.F) {
	f.Add(int32(0), uint8(0))
	f.Add(int32(0), uint8(math.MaxUint8))
	f.Add(int32(math.MinInt32), uint8(0))
	f.Add(int32(math.MinInt32), uint8(math.MaxUint8))
	f.Add(int32(math.MaxInt32), uint8(0))
	f.Add(int32(math.MaxInt32), uint8(math.MaxUint8))
	f.Fuzz(func(t *testing.T, k int32, v uint8) {
		vs1 := struct {
			Elem1 int32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int32
			Elem2 uint8
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int32
			Elem2 uint8
			Elem3 int32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt32Uint16Struct(f *testing.F) {
	f.Add(int32(0), uint16(0))
	f.Add(int32(0), uint16(math.MaxUint16))
	f.Add(int32(math.MinInt32), uint16(0))
	f.Add(int32(math.MinInt32), uint16(math.MaxUint16))
	f.Add(int32(math.MaxInt32), uint16(0))
	f.Add(int32(math.MaxInt32), uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, k int32, v uint16) {
		vs1 := struct {
			Elem1 int32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int32
			Elem2 uint16
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int32
			Elem2 uint16
			Elem3 int32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt32Uint32Struct(f *testing.F) {
	f.Add(int32(0), uint32(0))
	f.Add(int32(0), uint32(math.MaxUint32))
	f.Add(int32(math.MinInt32), uint32(0))
	f.Add(int32(math.MinInt32), uint32(math.MaxUint32))
	f.Add(int32(math.MaxInt32), uint32(0))
	f.Add(int32(math.MaxInt32), uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, k int32, v uint32) {
		vs1 := struct {
			Elem1 int32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int32
			Elem2 uint32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int32
			Elem2 uint32
			Elem3 int32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt32Uint64Struct(f *testing.F) {
	f.Add(int32(0), uint64(0))
	f.Add(int32(0), uint64(math.MaxInt64))
	f.Add(int32(math.MinInt32), uint64(0))
	f.Add(int32(math.MinInt32), uint64(math.MaxInt64))
	f.Add(int32(math.MaxInt32), uint64(0))
	f.Add(int32(math.MaxInt32), uint64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k int32, v uint64) {
		vs1 := struct {
			Elem1 int32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int32
			Elem2 uint64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int32
			Elem2 uint64
			Elem3 int32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt32Float32Struct(f *testing.F) {
	f.Add(int32(0), float32(-math.MaxFloat32))
	f.Add(int32(0), float32(0))
	f.Add(int32(0), float32(math.MaxFloat32))
	f.Add(int32(math.MinInt32), float32(-math.MaxFloat32))
	f.Add(int32(math.MinInt32), float32(0))
	f.Add(int32(math.MinInt32), float32(math.MaxFloat32))
	f.Add(int32(math.MaxInt32), float32(-math.MaxFloat32))
	f.Add(int32(math.MaxInt32), float32(0))
	f.Add(int32(math.MaxInt32), float32(math.MaxFloat32))
	f.Fuzz(func(t *testing.T, k int32, v float32) {
		vs1 := struct {
			Elem1 int32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int32
			Elem2 float32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int32
			Elem2 float32
			Elem3 int32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt32Float64Struct(f *testing.F) {
	f.Add(int32(0), float64(-math.MaxFloat32))
	f.Add(int32(0), float64(0))
	f.Add(int32(0), float64(math.MaxFloat64))
	f.Add(int32(math.MinInt32), float64(-math.MaxFloat32))
	f.Add(int32(math.MinInt32), float64(0))
	f.Add(int32(math.MinInt32), float64(math.MaxFloat64))
	f.Add(int32(math.MaxInt32), float64(-math.MaxFloat32))
	f.Add(int32(math.MaxInt32), float64(0))
	f.Add(int32(math.MaxInt32), float64(math.MaxFloat64))
	f.Fuzz(func(t *testing.T, k int32, v float64) {
		vs1 := struct {
			Elem1 int32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int32
			Elem2 float64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int32
			Elem2 float64
			Elem3 int32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt32BoolStruct(f *testing.F) {
	f.Add(int32(0), bool(true))
	f.Add(int32(0), bool(false))
	f.Add(int32(math.MinInt32), bool(true))
	f.Add(int32(math.MinInt32), bool(false))
	f.Add(int32(math.MaxInt32), bool(true))
	f.Add(int32(math.MaxInt32), bool(false))
	f.Fuzz(func(t *testing.T, k int32, v bool) {
		vs1 := struct {
			Elem1 int32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int32
			Elem2 bool
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int32
			Elem2 bool
			Elem3 int32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt32StringStruct(f *testing.F) {
	f.Add(int32(0), string("a"))
	f.Add(int32(0), string("ab"))
	f.Add(int32(0), string("abc"))
	f.Add(int32(math.MinInt32), string("a"))
	f.Add(int32(math.MinInt32), string("ab"))
	f.Add(int32(math.MinInt32), string("abc"))
	f.Add(int32(math.MaxInt32), string("a"))
	f.Add(int32(math.MaxInt32), string("ab"))
	f.Add(int32(math.MaxInt32), string("abc"))
	f.Fuzz(func(t *testing.T, k int32, v string) {
		vs1 := struct {
			Elem1 int32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int32
			Elem2 string
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int32
			Elem2 string
			Elem3 int32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt32ByteStruct(f *testing.F) {
	f.Add(int32(0), []byte("x"))
	f.Add(int32(0), []byte("xy"))
	f.Add(int32(0), []byte("xyz"))
	f.Add(int32(math.MinInt32), []byte("x"))
	f.Add(int32(math.MinInt32), []byte("xy"))
	f.Add(int32(math.MinInt32), []byte("xyz"))
	f.Add(int32(math.MaxInt32), []byte("x"))
	f.Add(int32(math.MaxInt32), []byte("xy"))
	f.Add(int32(math.MaxInt32), []byte("xyz"))
	f.Fuzz(func(t *testing.T, k int32, v []byte) {
		vs1 := struct {
			Elem1 int32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int32
			Elem2 []byte
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int32
			Elem2 []byte
			Elem3 int32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt64IntStruct(f *testing.F) {
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
		vs1 := struct {
			Elem1 int64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int64
			Elem2 int
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int64
			Elem2 int
			Elem3 int64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt64Int8Struct(f *testing.F) {
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
		vs1 := struct {
			Elem1 int64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int64
			Elem2 int8
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int64
			Elem2 int8
			Elem3 int64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt64Int16Struct(f *testing.F) {
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
		vs1 := struct {
			Elem1 int64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int64
			Elem2 int16
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int64
			Elem2 int16
			Elem3 int64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt64Int32Struct(f *testing.F) {
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
		vs1 := struct {
			Elem1 int64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int64
			Elem2 int32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int64
			Elem2 int32
			Elem3 int64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt64Int64Struct(f *testing.F) {
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
		vs1 := struct {
			Elem1 int64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int64
			Elem2 int64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int64
			Elem2 int64
			Elem3 int64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt64UintStruct(f *testing.F) {
	f.Add(int64(0), uint(0))
	f.Add(int64(0), uint(math.MaxUint))
	f.Add(int64(math.MinInt64), uint(0))
	f.Add(int64(math.MinInt64), uint(math.MaxUint))
	f.Add(int64(math.MaxInt64), uint(0))
	f.Add(int64(math.MaxInt64), uint(math.MaxUint))
	f.Fuzz(func(t *testing.T, k int64, v uint) {
		vs1 := struct {
			Elem1 int64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int64
			Elem2 uint
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int64
			Elem2 uint
			Elem3 int64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt64Uint8Struct(f *testing.F) {
	f.Add(int64(0), uint8(0))
	f.Add(int64(0), uint8(math.MaxUint8))
	f.Add(int64(math.MinInt64), uint8(0))
	f.Add(int64(math.MinInt64), uint8(math.MaxUint8))
	f.Add(int64(math.MaxInt64), uint8(0))
	f.Add(int64(math.MaxInt64), uint8(math.MaxUint8))
	f.Fuzz(func(t *testing.T, k int64, v uint8) {
		vs1 := struct {
			Elem1 int64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int64
			Elem2 uint8
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int64
			Elem2 uint8
			Elem3 int64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt64Uint16Struct(f *testing.F) {
	f.Add(int64(0), uint16(0))
	f.Add(int64(0), uint16(math.MaxUint16))
	f.Add(int64(math.MinInt64), uint16(0))
	f.Add(int64(math.MinInt64), uint16(math.MaxUint16))
	f.Add(int64(math.MaxInt64), uint16(0))
	f.Add(int64(math.MaxInt64), uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, k int64, v uint16) {
		vs1 := struct {
			Elem1 int64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int64
			Elem2 uint16
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int64
			Elem2 uint16
			Elem3 int64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt64Uint32Struct(f *testing.F) {
	f.Add(int64(0), uint32(0))
	f.Add(int64(0), uint32(math.MaxUint32))
	f.Add(int64(math.MinInt64), uint32(0))
	f.Add(int64(math.MinInt64), uint32(math.MaxUint32))
	f.Add(int64(math.MaxInt64), uint32(0))
	f.Add(int64(math.MaxInt64), uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, k int64, v uint32) {
		vs1 := struct {
			Elem1 int64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int64
			Elem2 uint32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int64
			Elem2 uint32
			Elem3 int64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt64Uint64Struct(f *testing.F) {
	f.Add(int64(0), uint64(0))
	f.Add(int64(0), uint64(math.MaxInt64))
	f.Add(int64(math.MinInt64), uint64(0))
	f.Add(int64(math.MinInt64), uint64(math.MaxInt64))
	f.Add(int64(math.MaxInt64), uint64(0))
	f.Add(int64(math.MaxInt64), uint64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k int64, v uint64) {
		vs1 := struct {
			Elem1 int64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int64
			Elem2 uint64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int64
			Elem2 uint64
			Elem3 int64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt64Float32Struct(f *testing.F) {
	f.Add(int64(0), float32(-math.MaxFloat32))
	f.Add(int64(0), float32(0))
	f.Add(int64(0), float32(math.MaxFloat32))
	f.Add(int64(math.MinInt64), float32(-math.MaxFloat32))
	f.Add(int64(math.MinInt64), float32(0))
	f.Add(int64(math.MinInt64), float32(math.MaxFloat32))
	f.Add(int64(math.MaxInt64), float32(-math.MaxFloat32))
	f.Add(int64(math.MaxInt64), float32(0))
	f.Add(int64(math.MaxInt64), float32(math.MaxFloat32))
	f.Fuzz(func(t *testing.T, k int64, v float32) {
		vs1 := struct {
			Elem1 int64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int64
			Elem2 float32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int64
			Elem2 float32
			Elem3 int64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt64Float64Struct(f *testing.F) {
	f.Add(int64(0), float64(-math.MaxFloat32))
	f.Add(int64(0), float64(0))
	f.Add(int64(0), float64(math.MaxFloat64))
	f.Add(int64(math.MinInt64), float64(-math.MaxFloat32))
	f.Add(int64(math.MinInt64), float64(0))
	f.Add(int64(math.MinInt64), float64(math.MaxFloat64))
	f.Add(int64(math.MaxInt64), float64(-math.MaxFloat32))
	f.Add(int64(math.MaxInt64), float64(0))
	f.Add(int64(math.MaxInt64), float64(math.MaxFloat64))
	f.Fuzz(func(t *testing.T, k int64, v float64) {
		vs1 := struct {
			Elem1 int64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int64
			Elem2 float64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int64
			Elem2 float64
			Elem3 int64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt64BoolStruct(f *testing.F) {
	f.Add(int64(0), bool(true))
	f.Add(int64(0), bool(false))
	f.Add(int64(math.MinInt64), bool(true))
	f.Add(int64(math.MinInt64), bool(false))
	f.Add(int64(math.MaxInt64), bool(true))
	f.Add(int64(math.MaxInt64), bool(false))
	f.Fuzz(func(t *testing.T, k int64, v bool) {
		vs1 := struct {
			Elem1 int64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int64
			Elem2 bool
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int64
			Elem2 bool
			Elem3 int64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt64StringStruct(f *testing.F) {
	f.Add(int64(0), string("a"))
	f.Add(int64(0), string("ab"))
	f.Add(int64(0), string("abc"))
	f.Add(int64(math.MinInt64), string("a"))
	f.Add(int64(math.MinInt64), string("ab"))
	f.Add(int64(math.MinInt64), string("abc"))
	f.Add(int64(math.MaxInt64), string("a"))
	f.Add(int64(math.MaxInt64), string("ab"))
	f.Add(int64(math.MaxInt64), string("abc"))
	f.Fuzz(func(t *testing.T, k int64, v string) {
		vs1 := struct {
			Elem1 int64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int64
			Elem2 string
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int64
			Elem2 string
			Elem3 int64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzInt64ByteStruct(f *testing.F) {
	f.Add(int64(0), []byte("x"))
	f.Add(int64(0), []byte("xy"))
	f.Add(int64(0), []byte("xyz"))
	f.Add(int64(math.MinInt64), []byte("x"))
	f.Add(int64(math.MinInt64), []byte("xy"))
	f.Add(int64(math.MinInt64), []byte("xyz"))
	f.Add(int64(math.MaxInt64), []byte("x"))
	f.Add(int64(math.MaxInt64), []byte("xy"))
	f.Add(int64(math.MaxInt64), []byte("xyz"))
	f.Fuzz(func(t *testing.T, k int64, v []byte) {
		vs1 := struct {
			Elem1 int64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 int64
			Elem2 []byte
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 int64
			Elem2 []byte
			Elem3 int64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUintIntStruct(f *testing.F) {
	f.Add(uint(0), int(0))
	f.Add(uint(0), int(math.MinInt))
	f.Add(uint(0), int(math.MaxInt))
	f.Add(uint(math.MaxUint), int(0))
	f.Add(uint(math.MaxUint), int(math.MinInt))
	f.Add(uint(math.MaxUint), int(math.MaxInt))
	f.Fuzz(func(t *testing.T, k uint, v int) {
		vs1 := struct {
			Elem1 uint
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint
			Elem2 int
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint
			Elem2 int
			Elem3 uint
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUintInt8Struct(f *testing.F) {
	f.Add(uint(0), int8(0))
	f.Add(uint(0), int8(math.MinInt8))
	f.Add(uint(0), int8(math.MaxInt8))
	f.Add(uint(math.MaxUint), int8(0))
	f.Add(uint(math.MaxUint), int8(math.MinInt8))
	f.Add(uint(math.MaxUint), int8(math.MaxInt8))
	f.Fuzz(func(t *testing.T, k uint, v int8) {
		vs1 := struct {
			Elem1 uint
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint
			Elem2 int8
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint
			Elem2 int8
			Elem3 uint
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUintInt16Struct(f *testing.F) {
	f.Add(uint(0), int16(0))
	f.Add(uint(0), int16(math.MinInt16))
	f.Add(uint(0), int16(math.MaxInt16))
	f.Add(uint(math.MaxUint), int16(0))
	f.Add(uint(math.MaxUint), int16(math.MinInt16))
	f.Add(uint(math.MaxUint), int16(math.MaxInt16))
	f.Fuzz(func(t *testing.T, k uint, v int16) {
		vs1 := struct {
			Elem1 uint
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint
			Elem2 int16
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint
			Elem2 int16
			Elem3 uint
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUintInt32Struct(f *testing.F) {
	f.Add(uint(0), int32(0))
	f.Add(uint(0), int32(math.MinInt32))
	f.Add(uint(0), int32(math.MaxInt32))
	f.Add(uint(math.MaxUint), int32(0))
	f.Add(uint(math.MaxUint), int32(math.MinInt32))
	f.Add(uint(math.MaxUint), int32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, k uint, v int32) {
		vs1 := struct {
			Elem1 uint
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint
			Elem2 int32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint
			Elem2 int32
			Elem3 uint
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUintInt64Struct(f *testing.F) {
	f.Add(uint(0), int64(0))
	f.Add(uint(0), int64(math.MinInt64))
	f.Add(uint(0), int64(math.MaxInt64))
	f.Add(uint(math.MaxUint), int64(0))
	f.Add(uint(math.MaxUint), int64(math.MinInt64))
	f.Add(uint(math.MaxUint), int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k uint, v int64) {
		vs1 := struct {
			Elem1 uint
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint
			Elem2 int64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint
			Elem2 int64
			Elem3 uint
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUintUintStruct(f *testing.F) {
	f.Add(uint(0), uint(0))
	f.Add(uint(0), uint(math.MaxUint))
	f.Add(uint(math.MaxUint), uint(0))
	f.Add(uint(math.MaxUint), uint(math.MaxUint))
	f.Fuzz(func(t *testing.T, k uint, v uint) {
		vs1 := struct {
			Elem1 uint
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint
			Elem2 uint
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint
			Elem2 uint
			Elem3 uint
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUintUint8Struct(f *testing.F) {
	f.Add(uint(0), uint8(0))
	f.Add(uint(0), uint8(math.MaxUint8))
	f.Add(uint(math.MaxUint), uint8(0))
	f.Add(uint(math.MaxUint), uint8(math.MaxUint8))
	f.Fuzz(func(t *testing.T, k uint, v uint8) {
		vs1 := struct {
			Elem1 uint
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint
			Elem2 uint8
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint
			Elem2 uint8
			Elem3 uint
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUintUint16Struct(f *testing.F) {
	f.Add(uint(0), uint16(0))
	f.Add(uint(0), uint16(math.MaxUint16))
	f.Add(uint(math.MaxUint), uint16(0))
	f.Add(uint(math.MaxUint), uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, k uint, v uint16) {
		vs1 := struct {
			Elem1 uint
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint
			Elem2 uint16
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint
			Elem2 uint16
			Elem3 uint
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUintUint32Struct(f *testing.F) {
	f.Add(uint(0), uint32(0))
	f.Add(uint(0), uint32(math.MaxUint32))
	f.Add(uint(math.MaxUint), uint32(0))
	f.Add(uint(math.MaxUint), uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, k uint, v uint32) {
		vs1 := struct {
			Elem1 uint
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint
			Elem2 uint32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint
			Elem2 uint32
			Elem3 uint
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUintUint64Struct(f *testing.F) {
	f.Add(uint(0), uint64(0))
	f.Add(uint(0), uint64(math.MaxInt64))
	f.Add(uint(math.MaxUint), uint64(0))
	f.Add(uint(math.MaxUint), uint64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k uint, v uint64) {
		vs1 := struct {
			Elem1 uint
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint
			Elem2 uint64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint
			Elem2 uint64
			Elem3 uint
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUintFloat32Struct(f *testing.F) {
	f.Add(uint(0), float32(-math.MaxFloat32))
	f.Add(uint(0), float32(0))
	f.Add(uint(0), float32(math.MaxFloat32))
	f.Add(uint(math.MaxUint), float32(-math.MaxFloat32))
	f.Add(uint(math.MaxUint), float32(0))
	f.Add(uint(math.MaxUint), float32(math.MaxFloat32))
	f.Fuzz(func(t *testing.T, k uint, v float32) {
		vs1 := struct {
			Elem1 uint
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint
			Elem2 float32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint
			Elem2 float32
			Elem3 uint
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUintFloat64Struct(f *testing.F) {
	f.Add(uint(0), float64(-math.MaxFloat32))
	f.Add(uint(0), float64(0))
	f.Add(uint(0), float64(math.MaxFloat64))
	f.Add(uint(math.MaxUint), float64(-math.MaxFloat32))
	f.Add(uint(math.MaxUint), float64(0))
	f.Add(uint(math.MaxUint), float64(math.MaxFloat64))
	f.Fuzz(func(t *testing.T, k uint, v float64) {
		vs1 := struct {
			Elem1 uint
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint
			Elem2 float64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint
			Elem2 float64
			Elem3 uint
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUintBoolStruct(f *testing.F) {
	f.Add(uint(0), bool(true))
	f.Add(uint(0), bool(false))
	f.Add(uint(math.MaxUint), bool(true))
	f.Add(uint(math.MaxUint), bool(false))
	f.Fuzz(func(t *testing.T, k uint, v bool) {
		vs1 := struct {
			Elem1 uint
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint
			Elem2 bool
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint
			Elem2 bool
			Elem3 uint
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUintStringStruct(f *testing.F) {
	f.Add(uint(0), string("a"))
	f.Add(uint(0), string("ab"))
	f.Add(uint(0), string("abc"))
	f.Add(uint(math.MaxUint), string("a"))
	f.Add(uint(math.MaxUint), string("ab"))
	f.Add(uint(math.MaxUint), string("abc"))
	f.Fuzz(func(t *testing.T, k uint, v string) {
		vs1 := struct {
			Elem1 uint
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint
			Elem2 string
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint
			Elem2 string
			Elem3 uint
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUintByteStruct(f *testing.F) {
	f.Add(uint(0), []byte("x"))
	f.Add(uint(0), []byte("xy"))
	f.Add(uint(0), []byte("xyz"))
	f.Add(uint(math.MaxUint), []byte("x"))
	f.Add(uint(math.MaxUint), []byte("xy"))
	f.Add(uint(math.MaxUint), []byte("xyz"))
	f.Fuzz(func(t *testing.T, k uint, v []byte) {
		vs1 := struct {
			Elem1 uint
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint
			Elem2 []byte
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint
			Elem2 []byte
			Elem3 uint
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint8IntStruct(f *testing.F) {
	f.Add(uint8(0), int(0))
	f.Add(uint8(0), int(math.MinInt))
	f.Add(uint8(0), int(math.MaxInt))
	f.Add(uint8(math.MaxUint8), int(0))
	f.Add(uint8(math.MaxUint8), int(math.MinInt))
	f.Add(uint8(math.MaxUint8), int(math.MaxInt))
	f.Fuzz(func(t *testing.T, k uint8, v int) {
		vs1 := struct {
			Elem1 uint8
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint8
			Elem2 int
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint8
			Elem2 int
			Elem3 uint8
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint8Int8Struct(f *testing.F) {
	f.Add(uint8(0), int8(0))
	f.Add(uint8(0), int8(math.MinInt8))
	f.Add(uint8(0), int8(math.MaxInt8))
	f.Add(uint8(math.MaxUint8), int8(0))
	f.Add(uint8(math.MaxUint8), int8(math.MinInt8))
	f.Add(uint8(math.MaxUint8), int8(math.MaxInt8))
	f.Fuzz(func(t *testing.T, k uint8, v int8) {
		vs1 := struct {
			Elem1 uint8
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint8
			Elem2 int8
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint8
			Elem2 int8
			Elem3 uint8
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint8Int16Struct(f *testing.F) {
	f.Add(uint8(0), int16(0))
	f.Add(uint8(0), int16(math.MinInt16))
	f.Add(uint8(0), int16(math.MaxInt16))
	f.Add(uint8(math.MaxUint8), int16(0))
	f.Add(uint8(math.MaxUint8), int16(math.MinInt16))
	f.Add(uint8(math.MaxUint8), int16(math.MaxInt16))
	f.Fuzz(func(t *testing.T, k uint8, v int16) {
		vs1 := struct {
			Elem1 uint8
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint8
			Elem2 int16
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint8
			Elem2 int16
			Elem3 uint8
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint8Int32Struct(f *testing.F) {
	f.Add(uint8(0), int32(0))
	f.Add(uint8(0), int32(math.MinInt32))
	f.Add(uint8(0), int32(math.MaxInt32))
	f.Add(uint8(math.MaxUint8), int32(0))
	f.Add(uint8(math.MaxUint8), int32(math.MinInt32))
	f.Add(uint8(math.MaxUint8), int32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, k uint8, v int32) {
		vs1 := struct {
			Elem1 uint8
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint8
			Elem2 int32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint8
			Elem2 int32
			Elem3 uint8
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint8Int64Struct(f *testing.F) {
	f.Add(uint8(0), int64(0))
	f.Add(uint8(0), int64(math.MinInt64))
	f.Add(uint8(0), int64(math.MaxInt64))
	f.Add(uint8(math.MaxUint8), int64(0))
	f.Add(uint8(math.MaxUint8), int64(math.MinInt64))
	f.Add(uint8(math.MaxUint8), int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k uint8, v int64) {
		vs1 := struct {
			Elem1 uint8
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint8
			Elem2 int64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint8
			Elem2 int64
			Elem3 uint8
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint8UintStruct(f *testing.F) {
	f.Add(uint8(0), uint(0))
	f.Add(uint8(0), uint(math.MaxUint))
	f.Add(uint8(math.MaxUint8), uint(0))
	f.Add(uint8(math.MaxUint8), uint(math.MaxUint))
	f.Fuzz(func(t *testing.T, k uint8, v uint) {
		vs1 := struct {
			Elem1 uint8
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint8
			Elem2 uint
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint8
			Elem2 uint
			Elem3 uint8
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint8Uint8Struct(f *testing.F) {
	f.Add(uint8(0), uint8(0))
	f.Add(uint8(0), uint8(math.MaxUint8))
	f.Add(uint8(math.MaxUint8), uint8(0))
	f.Add(uint8(math.MaxUint8), uint8(math.MaxUint8))
	f.Fuzz(func(t *testing.T, k uint8, v uint8) {
		vs1 := struct {
			Elem1 uint8
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint8
			Elem2 uint8
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint8
			Elem2 uint8
			Elem3 uint8
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint8Uint16Struct(f *testing.F) {
	f.Add(uint8(0), uint16(0))
	f.Add(uint8(0), uint16(math.MaxUint16))
	f.Add(uint8(math.MaxUint8), uint16(0))
	f.Add(uint8(math.MaxUint8), uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, k uint8, v uint16) {
		vs1 := struct {
			Elem1 uint8
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint8
			Elem2 uint16
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint8
			Elem2 uint16
			Elem3 uint8
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint8Uint32Struct(f *testing.F) {
	f.Add(uint8(0), uint32(0))
	f.Add(uint8(0), uint32(math.MaxUint32))
	f.Add(uint8(math.MaxUint8), uint32(0))
	f.Add(uint8(math.MaxUint8), uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, k uint8, v uint32) {
		vs1 := struct {
			Elem1 uint8
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint8
			Elem2 uint32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint8
			Elem2 uint32
			Elem3 uint8
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint8Uint64Struct(f *testing.F) {
	f.Add(uint8(0), uint64(0))
	f.Add(uint8(0), uint64(math.MaxInt64))
	f.Add(uint8(math.MaxUint8), uint64(0))
	f.Add(uint8(math.MaxUint8), uint64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k uint8, v uint64) {
		vs1 := struct {
			Elem1 uint8
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint8
			Elem2 uint64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint8
			Elem2 uint64
			Elem3 uint8
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint8Float32Struct(f *testing.F) {
	f.Add(uint8(0), float32(-math.MaxFloat32))
	f.Add(uint8(0), float32(0))
	f.Add(uint8(0), float32(math.MaxFloat32))
	f.Add(uint8(math.MaxUint8), float32(-math.MaxFloat32))
	f.Add(uint8(math.MaxUint8), float32(0))
	f.Add(uint8(math.MaxUint8), float32(math.MaxFloat32))
	f.Fuzz(func(t *testing.T, k uint8, v float32) {
		vs1 := struct {
			Elem1 uint8
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint8
			Elem2 float32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint8
			Elem2 float32
			Elem3 uint8
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint8Float64Struct(f *testing.F) {
	f.Add(uint8(0), float64(-math.MaxFloat32))
	f.Add(uint8(0), float64(0))
	f.Add(uint8(0), float64(math.MaxFloat64))
	f.Add(uint8(math.MaxUint8), float64(-math.MaxFloat32))
	f.Add(uint8(math.MaxUint8), float64(0))
	f.Add(uint8(math.MaxUint8), float64(math.MaxFloat64))
	f.Fuzz(func(t *testing.T, k uint8, v float64) {
		vs1 := struct {
			Elem1 uint8
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint8
			Elem2 float64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint8
			Elem2 float64
			Elem3 uint8
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint8BoolStruct(f *testing.F) {
	f.Add(uint8(0), bool(true))
	f.Add(uint8(0), bool(false))
	f.Add(uint8(math.MaxUint8), bool(true))
	f.Add(uint8(math.MaxUint8), bool(false))
	f.Fuzz(func(t *testing.T, k uint8, v bool) {
		vs1 := struct {
			Elem1 uint8
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint8
			Elem2 bool
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint8
			Elem2 bool
			Elem3 uint8
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint8StringStruct(f *testing.F) {
	f.Add(uint8(0), string("a"))
	f.Add(uint8(0), string("ab"))
	f.Add(uint8(0), string("abc"))
	f.Add(uint8(math.MaxUint8), string("a"))
	f.Add(uint8(math.MaxUint8), string("ab"))
	f.Add(uint8(math.MaxUint8), string("abc"))
	f.Fuzz(func(t *testing.T, k uint8, v string) {
		vs1 := struct {
			Elem1 uint8
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint8
			Elem2 string
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint8
			Elem2 string
			Elem3 uint8
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint8ByteStruct(f *testing.F) {
	f.Add(uint8(0), []byte("x"))
	f.Add(uint8(0), []byte("xy"))
	f.Add(uint8(0), []byte("xyz"))
	f.Add(uint8(math.MaxUint8), []byte("x"))
	f.Add(uint8(math.MaxUint8), []byte("xy"))
	f.Add(uint8(math.MaxUint8), []byte("xyz"))
	f.Fuzz(func(t *testing.T, k uint8, v []byte) {
		vs1 := struct {
			Elem1 uint8
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint8
			Elem2 []byte
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint8
			Elem2 []byte
			Elem3 uint8
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint16IntStruct(f *testing.F) {
	f.Add(uint16(0), int(0))
	f.Add(uint16(0), int(math.MinInt))
	f.Add(uint16(0), int(math.MaxInt))
	f.Add(uint16(math.MaxUint16), int(0))
	f.Add(uint16(math.MaxUint16), int(math.MinInt))
	f.Add(uint16(math.MaxUint16), int(math.MaxInt))
	f.Fuzz(func(t *testing.T, k uint16, v int) {
		vs1 := struct {
			Elem1 uint16
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint16
			Elem2 int
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint16
			Elem2 int
			Elem3 uint16
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint16Int8Struct(f *testing.F) {
	f.Add(uint16(0), int8(0))
	f.Add(uint16(0), int8(math.MinInt8))
	f.Add(uint16(0), int8(math.MaxInt8))
	f.Add(uint16(math.MaxUint16), int8(0))
	f.Add(uint16(math.MaxUint16), int8(math.MinInt8))
	f.Add(uint16(math.MaxUint16), int8(math.MaxInt8))
	f.Fuzz(func(t *testing.T, k uint16, v int8) {
		vs1 := struct {
			Elem1 uint16
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint16
			Elem2 int8
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint16
			Elem2 int8
			Elem3 uint16
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint16Int16Struct(f *testing.F) {
	f.Add(uint16(0), int16(0))
	f.Add(uint16(0), int16(math.MinInt16))
	f.Add(uint16(0), int16(math.MaxInt16))
	f.Add(uint16(math.MaxUint16), int16(0))
	f.Add(uint16(math.MaxUint16), int16(math.MinInt16))
	f.Add(uint16(math.MaxUint16), int16(math.MaxInt16))
	f.Fuzz(func(t *testing.T, k uint16, v int16) {
		vs1 := struct {
			Elem1 uint16
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint16
			Elem2 int16
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint16
			Elem2 int16
			Elem3 uint16
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint16Int32Struct(f *testing.F) {
	f.Add(uint16(0), int32(0))
	f.Add(uint16(0), int32(math.MinInt32))
	f.Add(uint16(0), int32(math.MaxInt32))
	f.Add(uint16(math.MaxUint16), int32(0))
	f.Add(uint16(math.MaxUint16), int32(math.MinInt32))
	f.Add(uint16(math.MaxUint16), int32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, k uint16, v int32) {
		vs1 := struct {
			Elem1 uint16
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint16
			Elem2 int32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint16
			Elem2 int32
			Elem3 uint16
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint16Int64Struct(f *testing.F) {
	f.Add(uint16(0), int64(0))
	f.Add(uint16(0), int64(math.MinInt64))
	f.Add(uint16(0), int64(math.MaxInt64))
	f.Add(uint16(math.MaxUint16), int64(0))
	f.Add(uint16(math.MaxUint16), int64(math.MinInt64))
	f.Add(uint16(math.MaxUint16), int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k uint16, v int64) {
		vs1 := struct {
			Elem1 uint16
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint16
			Elem2 int64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint16
			Elem2 int64
			Elem3 uint16
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint16UintStruct(f *testing.F) {
	f.Add(uint16(0), uint(0))
	f.Add(uint16(0), uint(math.MaxUint))
	f.Add(uint16(math.MaxUint16), uint(0))
	f.Add(uint16(math.MaxUint16), uint(math.MaxUint))
	f.Fuzz(func(t *testing.T, k uint16, v uint) {
		vs1 := struct {
			Elem1 uint16
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint16
			Elem2 uint
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint16
			Elem2 uint
			Elem3 uint16
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint16Uint8Struct(f *testing.F) {
	f.Add(uint16(0), uint8(0))
	f.Add(uint16(0), uint8(math.MaxUint8))
	f.Add(uint16(math.MaxUint16), uint8(0))
	f.Add(uint16(math.MaxUint16), uint8(math.MaxUint8))
	f.Fuzz(func(t *testing.T, k uint16, v uint8) {
		vs1 := struct {
			Elem1 uint16
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint16
			Elem2 uint8
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint16
			Elem2 uint8
			Elem3 uint16
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint16Uint16Struct(f *testing.F) {
	f.Add(uint16(0), uint16(0))
	f.Add(uint16(0), uint16(math.MaxUint16))
	f.Add(uint16(math.MaxUint16), uint16(0))
	f.Add(uint16(math.MaxUint16), uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, k uint16, v uint16) {
		vs1 := struct {
			Elem1 uint16
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint16
			Elem2 uint16
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint16
			Elem2 uint16
			Elem3 uint16
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint16Uint32Struct(f *testing.F) {
	f.Add(uint16(0), uint32(0))
	f.Add(uint16(0), uint32(math.MaxUint32))
	f.Add(uint16(math.MaxUint16), uint32(0))
	f.Add(uint16(math.MaxUint16), uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, k uint16, v uint32) {
		vs1 := struct {
			Elem1 uint16
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint16
			Elem2 uint32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint16
			Elem2 uint32
			Elem3 uint16
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint16Uint64Struct(f *testing.F) {
	f.Add(uint16(0), uint64(0))
	f.Add(uint16(0), uint64(math.MaxInt64))
	f.Add(uint16(math.MaxUint16), uint64(0))
	f.Add(uint16(math.MaxUint16), uint64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k uint16, v uint64) {
		vs1 := struct {
			Elem1 uint16
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint16
			Elem2 uint64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint16
			Elem2 uint64
			Elem3 uint16
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint16Float32Struct(f *testing.F) {
	f.Add(uint16(0), float32(-math.MaxFloat32))
	f.Add(uint16(0), float32(0))
	f.Add(uint16(0), float32(math.MaxFloat32))
	f.Add(uint16(math.MaxUint16), float32(-math.MaxFloat32))
	f.Add(uint16(math.MaxUint16), float32(0))
	f.Add(uint16(math.MaxUint16), float32(math.MaxFloat32))
	f.Fuzz(func(t *testing.T, k uint16, v float32) {
		vs1 := struct {
			Elem1 uint16
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint16
			Elem2 float32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint16
			Elem2 float32
			Elem3 uint16
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint16Float64Struct(f *testing.F) {
	f.Add(uint16(0), float64(-math.MaxFloat32))
	f.Add(uint16(0), float64(0))
	f.Add(uint16(0), float64(math.MaxFloat64))
	f.Add(uint16(math.MaxUint16), float64(-math.MaxFloat32))
	f.Add(uint16(math.MaxUint16), float64(0))
	f.Add(uint16(math.MaxUint16), float64(math.MaxFloat64))
	f.Fuzz(func(t *testing.T, k uint16, v float64) {
		vs1 := struct {
			Elem1 uint16
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint16
			Elem2 float64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint16
			Elem2 float64
			Elem3 uint16
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint16BoolStruct(f *testing.F) {
	f.Add(uint16(0), bool(true))
	f.Add(uint16(0), bool(false))
	f.Add(uint16(math.MaxUint16), bool(true))
	f.Add(uint16(math.MaxUint16), bool(false))
	f.Fuzz(func(t *testing.T, k uint16, v bool) {
		vs1 := struct {
			Elem1 uint16
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint16
			Elem2 bool
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint16
			Elem2 bool
			Elem3 uint16
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint16StringStruct(f *testing.F) {
	f.Add(uint16(0), string("a"))
	f.Add(uint16(0), string("ab"))
	f.Add(uint16(0), string("abc"))
	f.Add(uint16(math.MaxUint16), string("a"))
	f.Add(uint16(math.MaxUint16), string("ab"))
	f.Add(uint16(math.MaxUint16), string("abc"))
	f.Fuzz(func(t *testing.T, k uint16, v string) {
		vs1 := struct {
			Elem1 uint16
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint16
			Elem2 string
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint16
			Elem2 string
			Elem3 uint16
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint16ByteStruct(f *testing.F) {
	f.Add(uint16(0), []byte("x"))
	f.Add(uint16(0), []byte("xy"))
	f.Add(uint16(0), []byte("xyz"))
	f.Add(uint16(math.MaxUint16), []byte("x"))
	f.Add(uint16(math.MaxUint16), []byte("xy"))
	f.Add(uint16(math.MaxUint16), []byte("xyz"))
	f.Fuzz(func(t *testing.T, k uint16, v []byte) {
		vs1 := struct {
			Elem1 uint16
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint16
			Elem2 []byte
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint16
			Elem2 []byte
			Elem3 uint16
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint32IntStruct(f *testing.F) {
	f.Add(uint32(0), int(0))
	f.Add(uint32(0), int(math.MinInt))
	f.Add(uint32(0), int(math.MaxInt))
	f.Add(uint32(math.MaxUint32), int(0))
	f.Add(uint32(math.MaxUint32), int(math.MinInt))
	f.Add(uint32(math.MaxUint32), int(math.MaxInt))
	f.Fuzz(func(t *testing.T, k uint32, v int) {
		vs1 := struct {
			Elem1 uint32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint32
			Elem2 int
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint32
			Elem2 int
			Elem3 uint32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint32Int8Struct(f *testing.F) {
	f.Add(uint32(0), int8(0))
	f.Add(uint32(0), int8(math.MinInt8))
	f.Add(uint32(0), int8(math.MaxInt8))
	f.Add(uint32(math.MaxUint32), int8(0))
	f.Add(uint32(math.MaxUint32), int8(math.MinInt8))
	f.Add(uint32(math.MaxUint32), int8(math.MaxInt8))
	f.Fuzz(func(t *testing.T, k uint32, v int8) {
		vs1 := struct {
			Elem1 uint32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint32
			Elem2 int8
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint32
			Elem2 int8
			Elem3 uint32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint32Int16Struct(f *testing.F) {
	f.Add(uint32(0), int16(0))
	f.Add(uint32(0), int16(math.MinInt16))
	f.Add(uint32(0), int16(math.MaxInt16))
	f.Add(uint32(math.MaxUint32), int16(0))
	f.Add(uint32(math.MaxUint32), int16(math.MinInt16))
	f.Add(uint32(math.MaxUint32), int16(math.MaxInt16))
	f.Fuzz(func(t *testing.T, k uint32, v int16) {
		vs1 := struct {
			Elem1 uint32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint32
			Elem2 int16
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint32
			Elem2 int16
			Elem3 uint32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint32Int32Struct(f *testing.F) {
	f.Add(uint32(0), int32(0))
	f.Add(uint32(0), int32(math.MinInt32))
	f.Add(uint32(0), int32(math.MaxInt32))
	f.Add(uint32(math.MaxUint32), int32(0))
	f.Add(uint32(math.MaxUint32), int32(math.MinInt32))
	f.Add(uint32(math.MaxUint32), int32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, k uint32, v int32) {
		vs1 := struct {
			Elem1 uint32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint32
			Elem2 int32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint32
			Elem2 int32
			Elem3 uint32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint32Int64Struct(f *testing.F) {
	f.Add(uint32(0), int64(0))
	f.Add(uint32(0), int64(math.MinInt64))
	f.Add(uint32(0), int64(math.MaxInt64))
	f.Add(uint32(math.MaxUint32), int64(0))
	f.Add(uint32(math.MaxUint32), int64(math.MinInt64))
	f.Add(uint32(math.MaxUint32), int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k uint32, v int64) {
		vs1 := struct {
			Elem1 uint32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint32
			Elem2 int64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint32
			Elem2 int64
			Elem3 uint32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint32UintStruct(f *testing.F) {
	f.Add(uint32(0), uint(0))
	f.Add(uint32(0), uint(math.MaxUint))
	f.Add(uint32(math.MaxUint32), uint(0))
	f.Add(uint32(math.MaxUint32), uint(math.MaxUint))
	f.Fuzz(func(t *testing.T, k uint32, v uint) {
		vs1 := struct {
			Elem1 uint32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint32
			Elem2 uint
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint32
			Elem2 uint
			Elem3 uint32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint32Uint8Struct(f *testing.F) {
	f.Add(uint32(0), uint8(0))
	f.Add(uint32(0), uint8(math.MaxUint8))
	f.Add(uint32(math.MaxUint32), uint8(0))
	f.Add(uint32(math.MaxUint32), uint8(math.MaxUint8))
	f.Fuzz(func(t *testing.T, k uint32, v uint8) {
		vs1 := struct {
			Elem1 uint32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint32
			Elem2 uint8
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint32
			Elem2 uint8
			Elem3 uint32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint32Uint16Struct(f *testing.F) {
	f.Add(uint32(0), uint16(0))
	f.Add(uint32(0), uint16(math.MaxUint16))
	f.Add(uint32(math.MaxUint32), uint16(0))
	f.Add(uint32(math.MaxUint32), uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, k uint32, v uint16) {
		vs1 := struct {
			Elem1 uint32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint32
			Elem2 uint16
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint32
			Elem2 uint16
			Elem3 uint32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint32Uint32Struct(f *testing.F) {
	f.Add(uint32(0), uint32(0))
	f.Add(uint32(0), uint32(math.MaxUint32))
	f.Add(uint32(math.MaxUint32), uint32(0))
	f.Add(uint32(math.MaxUint32), uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, k uint32, v uint32) {
		vs1 := struct {
			Elem1 uint32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint32
			Elem2 uint32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint32
			Elem2 uint32
			Elem3 uint32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint32Uint64Struct(f *testing.F) {
	f.Add(uint32(0), uint64(0))
	f.Add(uint32(0), uint64(math.MaxInt64))
	f.Add(uint32(math.MaxUint32), uint64(0))
	f.Add(uint32(math.MaxUint32), uint64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k uint32, v uint64) {
		vs1 := struct {
			Elem1 uint32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint32
			Elem2 uint64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint32
			Elem2 uint64
			Elem3 uint32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint32Float32Struct(f *testing.F) {
	f.Add(uint32(0), float32(-math.MaxFloat32))
	f.Add(uint32(0), float32(0))
	f.Add(uint32(0), float32(math.MaxFloat32))
	f.Add(uint32(math.MaxUint32), float32(-math.MaxFloat32))
	f.Add(uint32(math.MaxUint32), float32(0))
	f.Add(uint32(math.MaxUint32), float32(math.MaxFloat32))
	f.Fuzz(func(t *testing.T, k uint32, v float32) {
		vs1 := struct {
			Elem1 uint32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint32
			Elem2 float32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint32
			Elem2 float32
			Elem3 uint32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint32Float64Struct(f *testing.F) {
	f.Add(uint32(0), float64(-math.MaxFloat32))
	f.Add(uint32(0), float64(0))
	f.Add(uint32(0), float64(math.MaxFloat64))
	f.Add(uint32(math.MaxUint32), float64(-math.MaxFloat32))
	f.Add(uint32(math.MaxUint32), float64(0))
	f.Add(uint32(math.MaxUint32), float64(math.MaxFloat64))
	f.Fuzz(func(t *testing.T, k uint32, v float64) {
		vs1 := struct {
			Elem1 uint32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint32
			Elem2 float64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint32
			Elem2 float64
			Elem3 uint32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint32BoolStruct(f *testing.F) {
	f.Add(uint32(0), bool(true))
	f.Add(uint32(0), bool(false))
	f.Add(uint32(math.MaxUint32), bool(true))
	f.Add(uint32(math.MaxUint32), bool(false))
	f.Fuzz(func(t *testing.T, k uint32, v bool) {
		vs1 := struct {
			Elem1 uint32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint32
			Elem2 bool
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint32
			Elem2 bool
			Elem3 uint32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint32StringStruct(f *testing.F) {
	f.Add(uint32(0), string("a"))
	f.Add(uint32(0), string("ab"))
	f.Add(uint32(0), string("abc"))
	f.Add(uint32(math.MaxUint32), string("a"))
	f.Add(uint32(math.MaxUint32), string("ab"))
	f.Add(uint32(math.MaxUint32), string("abc"))
	f.Fuzz(func(t *testing.T, k uint32, v string) {
		vs1 := struct {
			Elem1 uint32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint32
			Elem2 string
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint32
			Elem2 string
			Elem3 uint32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint32ByteStruct(f *testing.F) {
	f.Add(uint32(0), []byte("x"))
	f.Add(uint32(0), []byte("xy"))
	f.Add(uint32(0), []byte("xyz"))
	f.Add(uint32(math.MaxUint32), []byte("x"))
	f.Add(uint32(math.MaxUint32), []byte("xy"))
	f.Add(uint32(math.MaxUint32), []byte("xyz"))
	f.Fuzz(func(t *testing.T, k uint32, v []byte) {
		vs1 := struct {
			Elem1 uint32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint32
			Elem2 []byte
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint32
			Elem2 []byte
			Elem3 uint32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint64IntStruct(f *testing.F) {
	f.Add(uint64(0), int(0))
	f.Add(uint64(0), int(math.MinInt))
	f.Add(uint64(0), int(math.MaxInt))
	f.Add(uint64(math.MaxInt64), int(0))
	f.Add(uint64(math.MaxInt64), int(math.MinInt))
	f.Add(uint64(math.MaxInt64), int(math.MaxInt))
	f.Fuzz(func(t *testing.T, k uint64, v int) {
		vs1 := struct {
			Elem1 uint64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint64
			Elem2 int
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint64
			Elem2 int
			Elem3 uint64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint64Int8Struct(f *testing.F) {
	f.Add(uint64(0), int8(0))
	f.Add(uint64(0), int8(math.MinInt8))
	f.Add(uint64(0), int8(math.MaxInt8))
	f.Add(uint64(math.MaxInt64), int8(0))
	f.Add(uint64(math.MaxInt64), int8(math.MinInt8))
	f.Add(uint64(math.MaxInt64), int8(math.MaxInt8))
	f.Fuzz(func(t *testing.T, k uint64, v int8) {
		vs1 := struct {
			Elem1 uint64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint64
			Elem2 int8
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint64
			Elem2 int8
			Elem3 uint64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint64Int16Struct(f *testing.F) {
	f.Add(uint64(0), int16(0))
	f.Add(uint64(0), int16(math.MinInt16))
	f.Add(uint64(0), int16(math.MaxInt16))
	f.Add(uint64(math.MaxInt64), int16(0))
	f.Add(uint64(math.MaxInt64), int16(math.MinInt16))
	f.Add(uint64(math.MaxInt64), int16(math.MaxInt16))
	f.Fuzz(func(t *testing.T, k uint64, v int16) {
		vs1 := struct {
			Elem1 uint64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint64
			Elem2 int16
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint64
			Elem2 int16
			Elem3 uint64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint64Int32Struct(f *testing.F) {
	f.Add(uint64(0), int32(0))
	f.Add(uint64(0), int32(math.MinInt32))
	f.Add(uint64(0), int32(math.MaxInt32))
	f.Add(uint64(math.MaxInt64), int32(0))
	f.Add(uint64(math.MaxInt64), int32(math.MinInt32))
	f.Add(uint64(math.MaxInt64), int32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, k uint64, v int32) {
		vs1 := struct {
			Elem1 uint64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint64
			Elem2 int32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint64
			Elem2 int32
			Elem3 uint64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint64Int64Struct(f *testing.F) {
	f.Add(uint64(0), int64(0))
	f.Add(uint64(0), int64(math.MinInt64))
	f.Add(uint64(0), int64(math.MaxInt64))
	f.Add(uint64(math.MaxInt64), int64(0))
	f.Add(uint64(math.MaxInt64), int64(math.MinInt64))
	f.Add(uint64(math.MaxInt64), int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k uint64, v int64) {
		vs1 := struct {
			Elem1 uint64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint64
			Elem2 int64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint64
			Elem2 int64
			Elem3 uint64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint64UintStruct(f *testing.F) {
	f.Add(uint64(0), uint(0))
	f.Add(uint64(0), uint(math.MaxUint))
	f.Add(uint64(math.MaxInt64), uint(0))
	f.Add(uint64(math.MaxInt64), uint(math.MaxUint))
	f.Fuzz(func(t *testing.T, k uint64, v uint) {
		vs1 := struct {
			Elem1 uint64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint64
			Elem2 uint
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint64
			Elem2 uint
			Elem3 uint64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint64Uint8Struct(f *testing.F) {
	f.Add(uint64(0), uint8(0))
	f.Add(uint64(0), uint8(math.MaxUint8))
	f.Add(uint64(math.MaxInt64), uint8(0))
	f.Add(uint64(math.MaxInt64), uint8(math.MaxUint8))
	f.Fuzz(func(t *testing.T, k uint64, v uint8) {
		vs1 := struct {
			Elem1 uint64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint64
			Elem2 uint8
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint64
			Elem2 uint8
			Elem3 uint64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint64Uint16Struct(f *testing.F) {
	f.Add(uint64(0), uint16(0))
	f.Add(uint64(0), uint16(math.MaxUint16))
	f.Add(uint64(math.MaxInt64), uint16(0))
	f.Add(uint64(math.MaxInt64), uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, k uint64, v uint16) {
		vs1 := struct {
			Elem1 uint64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint64
			Elem2 uint16
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint64
			Elem2 uint16
			Elem3 uint64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint64Uint32Struct(f *testing.F) {
	f.Add(uint64(0), uint32(0))
	f.Add(uint64(0), uint32(math.MaxUint32))
	f.Add(uint64(math.MaxInt64), uint32(0))
	f.Add(uint64(math.MaxInt64), uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, k uint64, v uint32) {
		vs1 := struct {
			Elem1 uint64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint64
			Elem2 uint32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint64
			Elem2 uint32
			Elem3 uint64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint64Uint64Struct(f *testing.F) {
	f.Add(uint64(0), uint64(0))
	f.Add(uint64(0), uint64(math.MaxInt64))
	f.Add(uint64(math.MaxInt64), uint64(0))
	f.Add(uint64(math.MaxInt64), uint64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k uint64, v uint64) {
		vs1 := struct {
			Elem1 uint64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint64
			Elem2 uint64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint64
			Elem2 uint64
			Elem3 uint64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint64Float32Struct(f *testing.F) {
	f.Add(uint64(0), float32(-math.MaxFloat32))
	f.Add(uint64(0), float32(0))
	f.Add(uint64(0), float32(math.MaxFloat32))
	f.Add(uint64(math.MaxInt64), float32(-math.MaxFloat32))
	f.Add(uint64(math.MaxInt64), float32(0))
	f.Add(uint64(math.MaxInt64), float32(math.MaxFloat32))
	f.Fuzz(func(t *testing.T, k uint64, v float32) {
		vs1 := struct {
			Elem1 uint64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint64
			Elem2 float32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint64
			Elem2 float32
			Elem3 uint64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint64Float64Struct(f *testing.F) {
	f.Add(uint64(0), float64(-math.MaxFloat32))
	f.Add(uint64(0), float64(0))
	f.Add(uint64(0), float64(math.MaxFloat64))
	f.Add(uint64(math.MaxInt64), float64(-math.MaxFloat32))
	f.Add(uint64(math.MaxInt64), float64(0))
	f.Add(uint64(math.MaxInt64), float64(math.MaxFloat64))
	f.Fuzz(func(t *testing.T, k uint64, v float64) {
		vs1 := struct {
			Elem1 uint64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint64
			Elem2 float64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint64
			Elem2 float64
			Elem3 uint64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint64BoolStruct(f *testing.F) {
	f.Add(uint64(0), bool(true))
	f.Add(uint64(0), bool(false))
	f.Add(uint64(math.MaxInt64), bool(true))
	f.Add(uint64(math.MaxInt64), bool(false))
	f.Fuzz(func(t *testing.T, k uint64, v bool) {
		vs1 := struct {
			Elem1 uint64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint64
			Elem2 bool
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint64
			Elem2 bool
			Elem3 uint64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint64StringStruct(f *testing.F) {
	f.Add(uint64(0), string("a"))
	f.Add(uint64(0), string("ab"))
	f.Add(uint64(0), string("abc"))
	f.Add(uint64(math.MaxInt64), string("a"))
	f.Add(uint64(math.MaxInt64), string("ab"))
	f.Add(uint64(math.MaxInt64), string("abc"))
	f.Fuzz(func(t *testing.T, k uint64, v string) {
		vs1 := struct {
			Elem1 uint64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint64
			Elem2 string
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint64
			Elem2 string
			Elem3 uint64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzUint64ByteStruct(f *testing.F) {
	f.Add(uint64(0), []byte("x"))
	f.Add(uint64(0), []byte("xy"))
	f.Add(uint64(0), []byte("xyz"))
	f.Add(uint64(math.MaxInt64), []byte("x"))
	f.Add(uint64(math.MaxInt64), []byte("xy"))
	f.Add(uint64(math.MaxInt64), []byte("xyz"))
	f.Fuzz(func(t *testing.T, k uint64, v []byte) {
		vs1 := struct {
			Elem1 uint64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 uint64
			Elem2 []byte
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 uint64
			Elem2 []byte
			Elem3 uint64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzFloat32IntStruct(f *testing.F) {
	f.Add(float32(-math.MaxFloat32), int(0))
	f.Add(float32(-math.MaxFloat32), int(math.MinInt))
	f.Add(float32(-math.MaxFloat32), int(math.MaxInt))
	f.Add(float32(0), int(0))
	f.Add(float32(0), int(math.MinInt))
	f.Add(float32(0), int(math.MaxInt))
	f.Add(float32(math.MaxFloat32), int(0))
	f.Add(float32(math.MaxFloat32), int(math.MinInt))
	f.Add(float32(math.MaxFloat32), int(math.MaxInt))
	f.Fuzz(func(t *testing.T, k float32, v int) {
		vs1 := struct {
			Elem1 float32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 float32
			Elem2 int
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 float32
			Elem2 int
			Elem3 float32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzFloat32Int8Struct(f *testing.F) {
	f.Add(float32(-math.MaxFloat32), int8(0))
	f.Add(float32(-math.MaxFloat32), int8(math.MinInt8))
	f.Add(float32(-math.MaxFloat32), int8(math.MaxInt8))
	f.Add(float32(0), int8(0))
	f.Add(float32(0), int8(math.MinInt8))
	f.Add(float32(0), int8(math.MaxInt8))
	f.Add(float32(math.MaxFloat32), int8(0))
	f.Add(float32(math.MaxFloat32), int8(math.MinInt8))
	f.Add(float32(math.MaxFloat32), int8(math.MaxInt8))
	f.Fuzz(func(t *testing.T, k float32, v int8) {
		vs1 := struct {
			Elem1 float32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 float32
			Elem2 int8
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 float32
			Elem2 int8
			Elem3 float32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzFloat32Int16Struct(f *testing.F) {
	f.Add(float32(-math.MaxFloat32), int16(0))
	f.Add(float32(-math.MaxFloat32), int16(math.MinInt16))
	f.Add(float32(-math.MaxFloat32), int16(math.MaxInt16))
	f.Add(float32(0), int16(0))
	f.Add(float32(0), int16(math.MinInt16))
	f.Add(float32(0), int16(math.MaxInt16))
	f.Add(float32(math.MaxFloat32), int16(0))
	f.Add(float32(math.MaxFloat32), int16(math.MinInt16))
	f.Add(float32(math.MaxFloat32), int16(math.MaxInt16))
	f.Fuzz(func(t *testing.T, k float32, v int16) {
		vs1 := struct {
			Elem1 float32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 float32
			Elem2 int16
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 float32
			Elem2 int16
			Elem3 float32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzFloat32Int32Struct(f *testing.F) {
	f.Add(float32(-math.MaxFloat32), int32(0))
	f.Add(float32(-math.MaxFloat32), int32(math.MinInt32))
	f.Add(float32(-math.MaxFloat32), int32(math.MaxInt32))
	f.Add(float32(0), int32(0))
	f.Add(float32(0), int32(math.MinInt32))
	f.Add(float32(0), int32(math.MaxInt32))
	f.Add(float32(math.MaxFloat32), int32(0))
	f.Add(float32(math.MaxFloat32), int32(math.MinInt32))
	f.Add(float32(math.MaxFloat32), int32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, k float32, v int32) {
		vs1 := struct {
			Elem1 float32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 float32
			Elem2 int32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 float32
			Elem2 int32
			Elem3 float32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzFloat32Int64Struct(f *testing.F) {
	f.Add(float32(-math.MaxFloat32), int64(0))
	f.Add(float32(-math.MaxFloat32), int64(math.MinInt64))
	f.Add(float32(-math.MaxFloat32), int64(math.MaxInt64))
	f.Add(float32(0), int64(0))
	f.Add(float32(0), int64(math.MinInt64))
	f.Add(float32(0), int64(math.MaxInt64))
	f.Add(float32(math.MaxFloat32), int64(0))
	f.Add(float32(math.MaxFloat32), int64(math.MinInt64))
	f.Add(float32(math.MaxFloat32), int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k float32, v int64) {
		vs1 := struct {
			Elem1 float32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 float32
			Elem2 int64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 float32
			Elem2 int64
			Elem3 float32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzFloat32UintStruct(f *testing.F) {
	f.Add(float32(-math.MaxFloat32), uint(0))
	f.Add(float32(-math.MaxFloat32), uint(math.MaxUint))
	f.Add(float32(0), uint(0))
	f.Add(float32(0), uint(math.MaxUint))
	f.Add(float32(math.MaxFloat32), uint(0))
	f.Add(float32(math.MaxFloat32), uint(math.MaxUint))
	f.Fuzz(func(t *testing.T, k float32, v uint) {
		vs1 := struct {
			Elem1 float32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 float32
			Elem2 uint
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 float32
			Elem2 uint
			Elem3 float32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzFloat32Uint8Struct(f *testing.F) {
	f.Add(float32(-math.MaxFloat32), uint8(0))
	f.Add(float32(-math.MaxFloat32), uint8(math.MaxUint8))
	f.Add(float32(0), uint8(0))
	f.Add(float32(0), uint8(math.MaxUint8))
	f.Add(float32(math.MaxFloat32), uint8(0))
	f.Add(float32(math.MaxFloat32), uint8(math.MaxUint8))
	f.Fuzz(func(t *testing.T, k float32, v uint8) {
		vs1 := struct {
			Elem1 float32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 float32
			Elem2 uint8
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 float32
			Elem2 uint8
			Elem3 float32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzFloat32Uint16Struct(f *testing.F) {
	f.Add(float32(-math.MaxFloat32), uint16(0))
	f.Add(float32(-math.MaxFloat32), uint16(math.MaxUint16))
	f.Add(float32(0), uint16(0))
	f.Add(float32(0), uint16(math.MaxUint16))
	f.Add(float32(math.MaxFloat32), uint16(0))
	f.Add(float32(math.MaxFloat32), uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, k float32, v uint16) {
		vs1 := struct {
			Elem1 float32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 float32
			Elem2 uint16
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 float32
			Elem2 uint16
			Elem3 float32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzFloat32Uint32Struct(f *testing.F) {
	f.Add(float32(-math.MaxFloat32), uint32(0))
	f.Add(float32(-math.MaxFloat32), uint32(math.MaxUint32))
	f.Add(float32(0), uint32(0))
	f.Add(float32(0), uint32(math.MaxUint32))
	f.Add(float32(math.MaxFloat32), uint32(0))
	f.Add(float32(math.MaxFloat32), uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, k float32, v uint32) {
		vs1 := struct {
			Elem1 float32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 float32
			Elem2 uint32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 float32
			Elem2 uint32
			Elem3 float32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzFloat32Uint64Struct(f *testing.F) {
	f.Add(float32(-math.MaxFloat32), uint64(0))
	f.Add(float32(-math.MaxFloat32), uint64(math.MaxInt64))
	f.Add(float32(0), uint64(0))
	f.Add(float32(0), uint64(math.MaxInt64))
	f.Add(float32(math.MaxFloat32), uint64(0))
	f.Add(float32(math.MaxFloat32), uint64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k float32, v uint64) {
		vs1 := struct {
			Elem1 float32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 float32
			Elem2 uint64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 float32
			Elem2 uint64
			Elem3 float32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzFloat32Float32Struct(f *testing.F) {
	f.Add(float32(-math.MaxFloat32), float32(-math.MaxFloat32))
	f.Add(float32(-math.MaxFloat32), float32(0))
	f.Add(float32(-math.MaxFloat32), float32(math.MaxFloat32))
	f.Add(float32(0), float32(-math.MaxFloat32))
	f.Add(float32(0), float32(0))
	f.Add(float32(0), float32(math.MaxFloat32))
	f.Add(float32(math.MaxFloat32), float32(-math.MaxFloat32))
	f.Add(float32(math.MaxFloat32), float32(0))
	f.Add(float32(math.MaxFloat32), float32(math.MaxFloat32))
	f.Fuzz(func(t *testing.T, k float32, v float32) {
		vs1 := struct {
			Elem1 float32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 float32
			Elem2 float32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 float32
			Elem2 float32
			Elem3 float32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzFloat32Float64Struct(f *testing.F) {
	f.Add(float32(-math.MaxFloat32), float64(-math.MaxFloat32))
	f.Add(float32(-math.MaxFloat32), float64(0))
	f.Add(float32(-math.MaxFloat32), float64(math.MaxFloat64))
	f.Add(float32(0), float64(-math.MaxFloat32))
	f.Add(float32(0), float64(0))
	f.Add(float32(0), float64(math.MaxFloat64))
	f.Add(float32(math.MaxFloat32), float64(-math.MaxFloat32))
	f.Add(float32(math.MaxFloat32), float64(0))
	f.Add(float32(math.MaxFloat32), float64(math.MaxFloat64))
	f.Fuzz(func(t *testing.T, k float32, v float64) {
		vs1 := struct {
			Elem1 float32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 float32
			Elem2 float64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 float32
			Elem2 float64
			Elem3 float32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzFloat32BoolStruct(f *testing.F) {
	f.Add(float32(-math.MaxFloat32), bool(true))
	f.Add(float32(-math.MaxFloat32), bool(false))
	f.Add(float32(0), bool(true))
	f.Add(float32(0), bool(false))
	f.Add(float32(math.MaxFloat32), bool(true))
	f.Add(float32(math.MaxFloat32), bool(false))
	f.Fuzz(func(t *testing.T, k float32, v bool) {
		vs1 := struct {
			Elem1 float32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 float32
			Elem2 bool
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 float32
			Elem2 bool
			Elem3 float32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzFloat32StringStruct(f *testing.F) {
	f.Add(float32(-math.MaxFloat32), string("a"))
	f.Add(float32(-math.MaxFloat32), string("ab"))
	f.Add(float32(-math.MaxFloat32), string("abc"))
	f.Add(float32(0), string("a"))
	f.Add(float32(0), string("ab"))
	f.Add(float32(0), string("abc"))
	f.Add(float32(math.MaxFloat32), string("a"))
	f.Add(float32(math.MaxFloat32), string("ab"))
	f.Add(float32(math.MaxFloat32), string("abc"))
	f.Fuzz(func(t *testing.T, k float32, v string) {
		vs1 := struct {
			Elem1 float32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 float32
			Elem2 string
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 float32
			Elem2 string
			Elem3 float32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzFloat32ByteStruct(f *testing.F) {
	f.Add(float32(-math.MaxFloat32), []byte("x"))
	f.Add(float32(-math.MaxFloat32), []byte("xy"))
	f.Add(float32(-math.MaxFloat32), []byte("xyz"))
	f.Add(float32(0), []byte("x"))
	f.Add(float32(0), []byte("xy"))
	f.Add(float32(0), []byte("xyz"))
	f.Add(float32(math.MaxFloat32), []byte("x"))
	f.Add(float32(math.MaxFloat32), []byte("xy"))
	f.Add(float32(math.MaxFloat32), []byte("xyz"))
	f.Fuzz(func(t *testing.T, k float32, v []byte) {
		vs1 := struct {
			Elem1 float32
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 float32
			Elem2 []byte
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 float32
			Elem2 []byte
			Elem3 float32
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzFloat64IntStruct(f *testing.F) {
	f.Add(float64(-math.MaxFloat32), int(0))
	f.Add(float64(-math.MaxFloat32), int(math.MinInt))
	f.Add(float64(-math.MaxFloat32), int(math.MaxInt))
	f.Add(float64(0), int(0))
	f.Add(float64(0), int(math.MinInt))
	f.Add(float64(0), int(math.MaxInt))
	f.Add(float64(math.MaxFloat64), int(0))
	f.Add(float64(math.MaxFloat64), int(math.MinInt))
	f.Add(float64(math.MaxFloat64), int(math.MaxInt))
	f.Fuzz(func(t *testing.T, k float64, v int) {
		vs1 := struct {
			Elem1 float64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 float64
			Elem2 int
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 float64
			Elem2 int
			Elem3 float64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzFloat64Int8Struct(f *testing.F) {
	f.Add(float64(-math.MaxFloat32), int8(0))
	f.Add(float64(-math.MaxFloat32), int8(math.MinInt8))
	f.Add(float64(-math.MaxFloat32), int8(math.MaxInt8))
	f.Add(float64(0), int8(0))
	f.Add(float64(0), int8(math.MinInt8))
	f.Add(float64(0), int8(math.MaxInt8))
	f.Add(float64(math.MaxFloat64), int8(0))
	f.Add(float64(math.MaxFloat64), int8(math.MinInt8))
	f.Add(float64(math.MaxFloat64), int8(math.MaxInt8))
	f.Fuzz(func(t *testing.T, k float64, v int8) {
		vs1 := struct {
			Elem1 float64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 float64
			Elem2 int8
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 float64
			Elem2 int8
			Elem3 float64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzFloat64Int16Struct(f *testing.F) {
	f.Add(float64(-math.MaxFloat32), int16(0))
	f.Add(float64(-math.MaxFloat32), int16(math.MinInt16))
	f.Add(float64(-math.MaxFloat32), int16(math.MaxInt16))
	f.Add(float64(0), int16(0))
	f.Add(float64(0), int16(math.MinInt16))
	f.Add(float64(0), int16(math.MaxInt16))
	f.Add(float64(math.MaxFloat64), int16(0))
	f.Add(float64(math.MaxFloat64), int16(math.MinInt16))
	f.Add(float64(math.MaxFloat64), int16(math.MaxInt16))
	f.Fuzz(func(t *testing.T, k float64, v int16) {
		vs1 := struct {
			Elem1 float64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 float64
			Elem2 int16
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 float64
			Elem2 int16
			Elem3 float64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzFloat64Int32Struct(f *testing.F) {
	f.Add(float64(-math.MaxFloat32), int32(0))
	f.Add(float64(-math.MaxFloat32), int32(math.MinInt32))
	f.Add(float64(-math.MaxFloat32), int32(math.MaxInt32))
	f.Add(float64(0), int32(0))
	f.Add(float64(0), int32(math.MinInt32))
	f.Add(float64(0), int32(math.MaxInt32))
	f.Add(float64(math.MaxFloat64), int32(0))
	f.Add(float64(math.MaxFloat64), int32(math.MinInt32))
	f.Add(float64(math.MaxFloat64), int32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, k float64, v int32) {
		vs1 := struct {
			Elem1 float64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 float64
			Elem2 int32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 float64
			Elem2 int32
			Elem3 float64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzFloat64Int64Struct(f *testing.F) {
	f.Add(float64(-math.MaxFloat32), int64(0))
	f.Add(float64(-math.MaxFloat32), int64(math.MinInt64))
	f.Add(float64(-math.MaxFloat32), int64(math.MaxInt64))
	f.Add(float64(0), int64(0))
	f.Add(float64(0), int64(math.MinInt64))
	f.Add(float64(0), int64(math.MaxInt64))
	f.Add(float64(math.MaxFloat64), int64(0))
	f.Add(float64(math.MaxFloat64), int64(math.MinInt64))
	f.Add(float64(math.MaxFloat64), int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k float64, v int64) {
		vs1 := struct {
			Elem1 float64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 float64
			Elem2 int64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 float64
			Elem2 int64
			Elem3 float64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzFloat64UintStruct(f *testing.F) {
	f.Add(float64(-math.MaxFloat32), uint(0))
	f.Add(float64(-math.MaxFloat32), uint(math.MaxUint))
	f.Add(float64(0), uint(0))
	f.Add(float64(0), uint(math.MaxUint))
	f.Add(float64(math.MaxFloat64), uint(0))
	f.Add(float64(math.MaxFloat64), uint(math.MaxUint))
	f.Fuzz(func(t *testing.T, k float64, v uint) {
		vs1 := struct {
			Elem1 float64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 float64
			Elem2 uint
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 float64
			Elem2 uint
			Elem3 float64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzFloat64Uint8Struct(f *testing.F) {
	f.Add(float64(-math.MaxFloat32), uint8(0))
	f.Add(float64(-math.MaxFloat32), uint8(math.MaxUint8))
	f.Add(float64(0), uint8(0))
	f.Add(float64(0), uint8(math.MaxUint8))
	f.Add(float64(math.MaxFloat64), uint8(0))
	f.Add(float64(math.MaxFloat64), uint8(math.MaxUint8))
	f.Fuzz(func(t *testing.T, k float64, v uint8) {
		vs1 := struct {
			Elem1 float64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 float64
			Elem2 uint8
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 float64
			Elem2 uint8
			Elem3 float64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzFloat64Uint16Struct(f *testing.F) {
	f.Add(float64(-math.MaxFloat32), uint16(0))
	f.Add(float64(-math.MaxFloat32), uint16(math.MaxUint16))
	f.Add(float64(0), uint16(0))
	f.Add(float64(0), uint16(math.MaxUint16))
	f.Add(float64(math.MaxFloat64), uint16(0))
	f.Add(float64(math.MaxFloat64), uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, k float64, v uint16) {
		vs1 := struct {
			Elem1 float64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 float64
			Elem2 uint16
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 float64
			Elem2 uint16
			Elem3 float64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzFloat64Uint32Struct(f *testing.F) {
	f.Add(float64(-math.MaxFloat32), uint32(0))
	f.Add(float64(-math.MaxFloat32), uint32(math.MaxUint32))
	f.Add(float64(0), uint32(0))
	f.Add(float64(0), uint32(math.MaxUint32))
	f.Add(float64(math.MaxFloat64), uint32(0))
	f.Add(float64(math.MaxFloat64), uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, k float64, v uint32) {
		vs1 := struct {
			Elem1 float64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 float64
			Elem2 uint32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 float64
			Elem2 uint32
			Elem3 float64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzFloat64Uint64Struct(f *testing.F) {
	f.Add(float64(-math.MaxFloat32), uint64(0))
	f.Add(float64(-math.MaxFloat32), uint64(math.MaxInt64))
	f.Add(float64(0), uint64(0))
	f.Add(float64(0), uint64(math.MaxInt64))
	f.Add(float64(math.MaxFloat64), uint64(0))
	f.Add(float64(math.MaxFloat64), uint64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k float64, v uint64) {
		vs1 := struct {
			Elem1 float64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 float64
			Elem2 uint64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 float64
			Elem2 uint64
			Elem3 float64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzFloat64Float32Struct(f *testing.F) {
	f.Add(float64(-math.MaxFloat32), float32(-math.MaxFloat32))
	f.Add(float64(-math.MaxFloat32), float32(0))
	f.Add(float64(-math.MaxFloat32), float32(math.MaxFloat32))
	f.Add(float64(0), float32(-math.MaxFloat32))
	f.Add(float64(0), float32(0))
	f.Add(float64(0), float32(math.MaxFloat32))
	f.Add(float64(math.MaxFloat64), float32(-math.MaxFloat32))
	f.Add(float64(math.MaxFloat64), float32(0))
	f.Add(float64(math.MaxFloat64), float32(math.MaxFloat32))
	f.Fuzz(func(t *testing.T, k float64, v float32) {
		vs1 := struct {
			Elem1 float64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 float64
			Elem2 float32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 float64
			Elem2 float32
			Elem3 float64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzFloat64Float64Struct(f *testing.F) {
	f.Add(float64(-math.MaxFloat32), float64(-math.MaxFloat32))
	f.Add(float64(-math.MaxFloat32), float64(0))
	f.Add(float64(-math.MaxFloat32), float64(math.MaxFloat64))
	f.Add(float64(0), float64(-math.MaxFloat32))
	f.Add(float64(0), float64(0))
	f.Add(float64(0), float64(math.MaxFloat64))
	f.Add(float64(math.MaxFloat64), float64(-math.MaxFloat32))
	f.Add(float64(math.MaxFloat64), float64(0))
	f.Add(float64(math.MaxFloat64), float64(math.MaxFloat64))
	f.Fuzz(func(t *testing.T, k float64, v float64) {
		vs1 := struct {
			Elem1 float64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 float64
			Elem2 float64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 float64
			Elem2 float64
			Elem3 float64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzFloat64BoolStruct(f *testing.F) {
	f.Add(float64(-math.MaxFloat32), bool(true))
	f.Add(float64(-math.MaxFloat32), bool(false))
	f.Add(float64(0), bool(true))
	f.Add(float64(0), bool(false))
	f.Add(float64(math.MaxFloat64), bool(true))
	f.Add(float64(math.MaxFloat64), bool(false))
	f.Fuzz(func(t *testing.T, k float64, v bool) {
		vs1 := struct {
			Elem1 float64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 float64
			Elem2 bool
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 float64
			Elem2 bool
			Elem3 float64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzFloat64StringStruct(f *testing.F) {
	f.Add(float64(-math.MaxFloat32), string("a"))
	f.Add(float64(-math.MaxFloat32), string("ab"))
	f.Add(float64(-math.MaxFloat32), string("abc"))
	f.Add(float64(0), string("a"))
	f.Add(float64(0), string("ab"))
	f.Add(float64(0), string("abc"))
	f.Add(float64(math.MaxFloat64), string("a"))
	f.Add(float64(math.MaxFloat64), string("ab"))
	f.Add(float64(math.MaxFloat64), string("abc"))
	f.Fuzz(func(t *testing.T, k float64, v string) {
		vs1 := struct {
			Elem1 float64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 float64
			Elem2 string
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 float64
			Elem2 string
			Elem3 float64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzFloat64ByteStruct(f *testing.F) {
	f.Add(float64(-math.MaxFloat32), []byte("x"))
	f.Add(float64(-math.MaxFloat32), []byte("xy"))
	f.Add(float64(-math.MaxFloat32), []byte("xyz"))
	f.Add(float64(0), []byte("x"))
	f.Add(float64(0), []byte("xy"))
	f.Add(float64(0), []byte("xyz"))
	f.Add(float64(math.MaxFloat64), []byte("x"))
	f.Add(float64(math.MaxFloat64), []byte("xy"))
	f.Add(float64(math.MaxFloat64), []byte("xyz"))
	f.Fuzz(func(t *testing.T, k float64, v []byte) {
		vs1 := struct {
			Elem1 float64
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 float64
			Elem2 []byte
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 float64
			Elem2 []byte
			Elem3 float64
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzBoolIntStruct(f *testing.F) {
	f.Add(bool(true), int(0))
	f.Add(bool(true), int(math.MinInt))
	f.Add(bool(true), int(math.MaxInt))
	f.Add(bool(false), int(0))
	f.Add(bool(false), int(math.MinInt))
	f.Add(bool(false), int(math.MaxInt))
	f.Fuzz(func(t *testing.T, k bool, v int) {
		vs1 := struct {
			Elem1 bool
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 bool
			Elem2 int
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 bool
			Elem2 int
			Elem3 bool
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzBoolInt8Struct(f *testing.F) {
	f.Add(bool(true), int8(0))
	f.Add(bool(true), int8(math.MinInt8))
	f.Add(bool(true), int8(math.MaxInt8))
	f.Add(bool(false), int8(0))
	f.Add(bool(false), int8(math.MinInt8))
	f.Add(bool(false), int8(math.MaxInt8))
	f.Fuzz(func(t *testing.T, k bool, v int8) {
		vs1 := struct {
			Elem1 bool
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 bool
			Elem2 int8
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 bool
			Elem2 int8
			Elem3 bool
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzBoolInt16Struct(f *testing.F) {
	f.Add(bool(true), int16(0))
	f.Add(bool(true), int16(math.MinInt16))
	f.Add(bool(true), int16(math.MaxInt16))
	f.Add(bool(false), int16(0))
	f.Add(bool(false), int16(math.MinInt16))
	f.Add(bool(false), int16(math.MaxInt16))
	f.Fuzz(func(t *testing.T, k bool, v int16) {
		vs1 := struct {
			Elem1 bool
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 bool
			Elem2 int16
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 bool
			Elem2 int16
			Elem3 bool
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzBoolInt32Struct(f *testing.F) {
	f.Add(bool(true), int32(0))
	f.Add(bool(true), int32(math.MinInt32))
	f.Add(bool(true), int32(math.MaxInt32))
	f.Add(bool(false), int32(0))
	f.Add(bool(false), int32(math.MinInt32))
	f.Add(bool(false), int32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, k bool, v int32) {
		vs1 := struct {
			Elem1 bool
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 bool
			Elem2 int32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 bool
			Elem2 int32
			Elem3 bool
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzBoolInt64Struct(f *testing.F) {
	f.Add(bool(true), int64(0))
	f.Add(bool(true), int64(math.MinInt64))
	f.Add(bool(true), int64(math.MaxInt64))
	f.Add(bool(false), int64(0))
	f.Add(bool(false), int64(math.MinInt64))
	f.Add(bool(false), int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k bool, v int64) {
		vs1 := struct {
			Elem1 bool
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 bool
			Elem2 int64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 bool
			Elem2 int64
			Elem3 bool
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzBoolUintStruct(f *testing.F) {
	f.Add(bool(true), uint(0))
	f.Add(bool(true), uint(math.MaxUint))
	f.Add(bool(false), uint(0))
	f.Add(bool(false), uint(math.MaxUint))
	f.Fuzz(func(t *testing.T, k bool, v uint) {
		vs1 := struct {
			Elem1 bool
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 bool
			Elem2 uint
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 bool
			Elem2 uint
			Elem3 bool
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzBoolUint8Struct(f *testing.F) {
	f.Add(bool(true), uint8(0))
	f.Add(bool(true), uint8(math.MaxUint8))
	f.Add(bool(false), uint8(0))
	f.Add(bool(false), uint8(math.MaxUint8))
	f.Fuzz(func(t *testing.T, k bool, v uint8) {
		vs1 := struct {
			Elem1 bool
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 bool
			Elem2 uint8
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 bool
			Elem2 uint8
			Elem3 bool
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzBoolUint16Struct(f *testing.F) {
	f.Add(bool(true), uint16(0))
	f.Add(bool(true), uint16(math.MaxUint16))
	f.Add(bool(false), uint16(0))
	f.Add(bool(false), uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, k bool, v uint16) {
		vs1 := struct {
			Elem1 bool
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 bool
			Elem2 uint16
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 bool
			Elem2 uint16
			Elem3 bool
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzBoolUint32Struct(f *testing.F) {
	f.Add(bool(true), uint32(0))
	f.Add(bool(true), uint32(math.MaxUint32))
	f.Add(bool(false), uint32(0))
	f.Add(bool(false), uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, k bool, v uint32) {
		vs1 := struct {
			Elem1 bool
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 bool
			Elem2 uint32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 bool
			Elem2 uint32
			Elem3 bool
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzBoolUint64Struct(f *testing.F) {
	f.Add(bool(true), uint64(0))
	f.Add(bool(true), uint64(math.MaxInt64))
	f.Add(bool(false), uint64(0))
	f.Add(bool(false), uint64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k bool, v uint64) {
		vs1 := struct {
			Elem1 bool
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 bool
			Elem2 uint64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 bool
			Elem2 uint64
			Elem3 bool
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzBoolFloat32Struct(f *testing.F) {
	f.Add(bool(true), float32(-math.MaxFloat32))
	f.Add(bool(true), float32(0))
	f.Add(bool(true), float32(math.MaxFloat32))
	f.Add(bool(false), float32(-math.MaxFloat32))
	f.Add(bool(false), float32(0))
	f.Add(bool(false), float32(math.MaxFloat32))
	f.Fuzz(func(t *testing.T, k bool, v float32) {
		vs1 := struct {
			Elem1 bool
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 bool
			Elem2 float32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 bool
			Elem2 float32
			Elem3 bool
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzBoolFloat64Struct(f *testing.F) {
	f.Add(bool(true), float64(-math.MaxFloat32))
	f.Add(bool(true), float64(0))
	f.Add(bool(true), float64(math.MaxFloat64))
	f.Add(bool(false), float64(-math.MaxFloat32))
	f.Add(bool(false), float64(0))
	f.Add(bool(false), float64(math.MaxFloat64))
	f.Fuzz(func(t *testing.T, k bool, v float64) {
		vs1 := struct {
			Elem1 bool
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 bool
			Elem2 float64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 bool
			Elem2 float64
			Elem3 bool
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzBoolBoolStruct(f *testing.F) {
	f.Add(bool(true), bool(true))
	f.Add(bool(true), bool(false))
	f.Add(bool(false), bool(true))
	f.Add(bool(false), bool(false))
	f.Fuzz(func(t *testing.T, k bool, v bool) {
		vs1 := struct {
			Elem1 bool
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 bool
			Elem2 bool
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 bool
			Elem2 bool
			Elem3 bool
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzBoolStringStruct(f *testing.F) {
	f.Add(bool(true), string("a"))
	f.Add(bool(true), string("ab"))
	f.Add(bool(true), string("abc"))
	f.Add(bool(false), string("a"))
	f.Add(bool(false), string("ab"))
	f.Add(bool(false), string("abc"))
	f.Fuzz(func(t *testing.T, k bool, v string) {
		vs1 := struct {
			Elem1 bool
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 bool
			Elem2 string
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 bool
			Elem2 string
			Elem3 bool
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzBoolByteStruct(f *testing.F) {
	f.Add(bool(true), []byte("x"))
	f.Add(bool(true), []byte("xy"))
	f.Add(bool(true), []byte("xyz"))
	f.Add(bool(false), []byte("x"))
	f.Add(bool(false), []byte("xy"))
	f.Add(bool(false), []byte("xyz"))
	f.Fuzz(func(t *testing.T, k bool, v []byte) {
		vs1 := struct {
			Elem1 bool
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 bool
			Elem2 []byte
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 bool
			Elem2 []byte
			Elem3 bool
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzStringIntStruct(f *testing.F) {
	f.Add(string("a"), int(0))
	f.Add(string("a"), int(math.MinInt))
	f.Add(string("a"), int(math.MaxInt))
	f.Add(string("ab"), int(0))
	f.Add(string("ab"), int(math.MinInt))
	f.Add(string("ab"), int(math.MaxInt))
	f.Add(string("abc"), int(0))
	f.Add(string("abc"), int(math.MinInt))
	f.Add(string("abc"), int(math.MaxInt))
	f.Fuzz(func(t *testing.T, k string, v int) {
		vs1 := struct {
			Elem1 string
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 string
			Elem2 int
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 string
			Elem2 int
			Elem3 string
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzStringInt8Struct(f *testing.F) {
	f.Add(string("a"), int8(0))
	f.Add(string("a"), int8(math.MinInt8))
	f.Add(string("a"), int8(math.MaxInt8))
	f.Add(string("ab"), int8(0))
	f.Add(string("ab"), int8(math.MinInt8))
	f.Add(string("ab"), int8(math.MaxInt8))
	f.Add(string("abc"), int8(0))
	f.Add(string("abc"), int8(math.MinInt8))
	f.Add(string("abc"), int8(math.MaxInt8))
	f.Fuzz(func(t *testing.T, k string, v int8) {
		vs1 := struct {
			Elem1 string
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 string
			Elem2 int8
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 string
			Elem2 int8
			Elem3 string
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzStringInt16Struct(f *testing.F) {
	f.Add(string("a"), int16(0))
	f.Add(string("a"), int16(math.MinInt16))
	f.Add(string("a"), int16(math.MaxInt16))
	f.Add(string("ab"), int16(0))
	f.Add(string("ab"), int16(math.MinInt16))
	f.Add(string("ab"), int16(math.MaxInt16))
	f.Add(string("abc"), int16(0))
	f.Add(string("abc"), int16(math.MinInt16))
	f.Add(string("abc"), int16(math.MaxInt16))
	f.Fuzz(func(t *testing.T, k string, v int16) {
		vs1 := struct {
			Elem1 string
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 string
			Elem2 int16
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 string
			Elem2 int16
			Elem3 string
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzStringInt32Struct(f *testing.F) {
	f.Add(string("a"), int32(0))
	f.Add(string("a"), int32(math.MinInt32))
	f.Add(string("a"), int32(math.MaxInt32))
	f.Add(string("ab"), int32(0))
	f.Add(string("ab"), int32(math.MinInt32))
	f.Add(string("ab"), int32(math.MaxInt32))
	f.Add(string("abc"), int32(0))
	f.Add(string("abc"), int32(math.MinInt32))
	f.Add(string("abc"), int32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, k string, v int32) {
		vs1 := struct {
			Elem1 string
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 string
			Elem2 int32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 string
			Elem2 int32
			Elem3 string
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzStringInt64Struct(f *testing.F) {
	f.Add(string("a"), int64(0))
	f.Add(string("a"), int64(math.MinInt64))
	f.Add(string("a"), int64(math.MaxInt64))
	f.Add(string("ab"), int64(0))
	f.Add(string("ab"), int64(math.MinInt64))
	f.Add(string("ab"), int64(math.MaxInt64))
	f.Add(string("abc"), int64(0))
	f.Add(string("abc"), int64(math.MinInt64))
	f.Add(string("abc"), int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k string, v int64) {
		vs1 := struct {
			Elem1 string
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 string
			Elem2 int64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 string
			Elem2 int64
			Elem3 string
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzStringUintStruct(f *testing.F) {
	f.Add(string("a"), uint(0))
	f.Add(string("a"), uint(math.MaxUint))
	f.Add(string("ab"), uint(0))
	f.Add(string("ab"), uint(math.MaxUint))
	f.Add(string("abc"), uint(0))
	f.Add(string("abc"), uint(math.MaxUint))
	f.Fuzz(func(t *testing.T, k string, v uint) {
		vs1 := struct {
			Elem1 string
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 string
			Elem2 uint
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 string
			Elem2 uint
			Elem3 string
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzStringUint8Struct(f *testing.F) {
	f.Add(string("a"), uint8(0))
	f.Add(string("a"), uint8(math.MaxUint8))
	f.Add(string("ab"), uint8(0))
	f.Add(string("ab"), uint8(math.MaxUint8))
	f.Add(string("abc"), uint8(0))
	f.Add(string("abc"), uint8(math.MaxUint8))
	f.Fuzz(func(t *testing.T, k string, v uint8) {
		vs1 := struct {
			Elem1 string
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 string
			Elem2 uint8
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 string
			Elem2 uint8
			Elem3 string
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzStringUint16Struct(f *testing.F) {
	f.Add(string("a"), uint16(0))
	f.Add(string("a"), uint16(math.MaxUint16))
	f.Add(string("ab"), uint16(0))
	f.Add(string("ab"), uint16(math.MaxUint16))
	f.Add(string("abc"), uint16(0))
	f.Add(string("abc"), uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, k string, v uint16) {
		vs1 := struct {
			Elem1 string
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 string
			Elem2 uint16
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 string
			Elem2 uint16
			Elem3 string
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzStringUint32Struct(f *testing.F) {
	f.Add(string("a"), uint32(0))
	f.Add(string("a"), uint32(math.MaxUint32))
	f.Add(string("ab"), uint32(0))
	f.Add(string("ab"), uint32(math.MaxUint32))
	f.Add(string("abc"), uint32(0))
	f.Add(string("abc"), uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, k string, v uint32) {
		vs1 := struct {
			Elem1 string
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 string
			Elem2 uint32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 string
			Elem2 uint32
			Elem3 string
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzStringUint64Struct(f *testing.F) {
	f.Add(string("a"), uint64(0))
	f.Add(string("a"), uint64(math.MaxInt64))
	f.Add(string("ab"), uint64(0))
	f.Add(string("ab"), uint64(math.MaxInt64))
	f.Add(string("abc"), uint64(0))
	f.Add(string("abc"), uint64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k string, v uint64) {
		vs1 := struct {
			Elem1 string
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 string
			Elem2 uint64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 string
			Elem2 uint64
			Elem3 string
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzStringFloat32Struct(f *testing.F) {
	f.Add(string("a"), float32(-math.MaxFloat32))
	f.Add(string("a"), float32(0))
	f.Add(string("a"), float32(math.MaxFloat32))
	f.Add(string("ab"), float32(-math.MaxFloat32))
	f.Add(string("ab"), float32(0))
	f.Add(string("ab"), float32(math.MaxFloat32))
	f.Add(string("abc"), float32(-math.MaxFloat32))
	f.Add(string("abc"), float32(0))
	f.Add(string("abc"), float32(math.MaxFloat32))
	f.Fuzz(func(t *testing.T, k string, v float32) {
		vs1 := struct {
			Elem1 string
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 string
			Elem2 float32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 string
			Elem2 float32
			Elem3 string
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzStringFloat64Struct(f *testing.F) {
	f.Add(string("a"), float64(-math.MaxFloat32))
	f.Add(string("a"), float64(0))
	f.Add(string("a"), float64(math.MaxFloat64))
	f.Add(string("ab"), float64(-math.MaxFloat32))
	f.Add(string("ab"), float64(0))
	f.Add(string("ab"), float64(math.MaxFloat64))
	f.Add(string("abc"), float64(-math.MaxFloat32))
	f.Add(string("abc"), float64(0))
	f.Add(string("abc"), float64(math.MaxFloat64))
	f.Fuzz(func(t *testing.T, k string, v float64) {
		vs1 := struct {
			Elem1 string
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 string
			Elem2 float64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 string
			Elem2 float64
			Elem3 string
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzStringBoolStruct(f *testing.F) {
	f.Add(string("a"), bool(true))
	f.Add(string("a"), bool(false))
	f.Add(string("ab"), bool(true))
	f.Add(string("ab"), bool(false))
	f.Add(string("abc"), bool(true))
	f.Add(string("abc"), bool(false))
	f.Fuzz(func(t *testing.T, k string, v bool) {
		vs1 := struct {
			Elem1 string
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 string
			Elem2 bool
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 string
			Elem2 bool
			Elem3 string
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzStringStringStruct(f *testing.F) {
	f.Add(string("a"), string("a"))
	f.Add(string("a"), string("ab"))
	f.Add(string("a"), string("abc"))
	f.Add(string("ab"), string("a"))
	f.Add(string("ab"), string("ab"))
	f.Add(string("ab"), string("abc"))
	f.Add(string("abc"), string("a"))
	f.Add(string("abc"), string("ab"))
	f.Add(string("abc"), string("abc"))
	f.Fuzz(func(t *testing.T, k string, v string) {
		vs1 := struct {
			Elem1 string
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 string
			Elem2 string
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 string
			Elem2 string
			Elem3 string
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzStringByteStruct(f *testing.F) {
	f.Add(string("a"), []byte("x"))
	f.Add(string("a"), []byte("xy"))
	f.Add(string("a"), []byte("xyz"))
	f.Add(string("ab"), []byte("x"))
	f.Add(string("ab"), []byte("xy"))
	f.Add(string("ab"), []byte("xyz"))
	f.Add(string("abc"), []byte("x"))
	f.Add(string("abc"), []byte("xy"))
	f.Add(string("abc"), []byte("xyz"))
	f.Fuzz(func(t *testing.T, k string, v []byte) {
		vs1 := struct {
			Elem1 string
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 string
			Elem2 []byte
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 string
			Elem2 []byte
			Elem3 string
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzByteIntStruct(f *testing.F) {
	f.Add([]byte("x"), int(0))
	f.Add([]byte("x"), int(math.MinInt))
	f.Add([]byte("x"), int(math.MaxInt))
	f.Add([]byte("xy"), int(0))
	f.Add([]byte("xy"), int(math.MinInt))
	f.Add([]byte("xy"), int(math.MaxInt))
	f.Add([]byte("xyz"), int(0))
	f.Add([]byte("xyz"), int(math.MinInt))
	f.Add([]byte("xyz"), int(math.MaxInt))
	f.Fuzz(func(t *testing.T, k []byte, v int) {
		vs1 := struct {
			Elem1 []byte
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 []byte
			Elem2 int
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 []byte
			Elem2 int
			Elem3 []byte
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzByteInt8Struct(f *testing.F) {
	f.Add([]byte("x"), int8(0))
	f.Add([]byte("x"), int8(math.MinInt8))
	f.Add([]byte("x"), int8(math.MaxInt8))
	f.Add([]byte("xy"), int8(0))
	f.Add([]byte("xy"), int8(math.MinInt8))
	f.Add([]byte("xy"), int8(math.MaxInt8))
	f.Add([]byte("xyz"), int8(0))
	f.Add([]byte("xyz"), int8(math.MinInt8))
	f.Add([]byte("xyz"), int8(math.MaxInt8))
	f.Fuzz(func(t *testing.T, k []byte, v int8) {
		vs1 := struct {
			Elem1 []byte
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 []byte
			Elem2 int8
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 []byte
			Elem2 int8
			Elem3 []byte
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzByteInt16Struct(f *testing.F) {
	f.Add([]byte("x"), int16(0))
	f.Add([]byte("x"), int16(math.MinInt16))
	f.Add([]byte("x"), int16(math.MaxInt16))
	f.Add([]byte("xy"), int16(0))
	f.Add([]byte("xy"), int16(math.MinInt16))
	f.Add([]byte("xy"), int16(math.MaxInt16))
	f.Add([]byte("xyz"), int16(0))
	f.Add([]byte("xyz"), int16(math.MinInt16))
	f.Add([]byte("xyz"), int16(math.MaxInt16))
	f.Fuzz(func(t *testing.T, k []byte, v int16) {
		vs1 := struct {
			Elem1 []byte
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 []byte
			Elem2 int16
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 []byte
			Elem2 int16
			Elem3 []byte
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzByteInt32Struct(f *testing.F) {
	f.Add([]byte("x"), int32(0))
	f.Add([]byte("x"), int32(math.MinInt32))
	f.Add([]byte("x"), int32(math.MaxInt32))
	f.Add([]byte("xy"), int32(0))
	f.Add([]byte("xy"), int32(math.MinInt32))
	f.Add([]byte("xy"), int32(math.MaxInt32))
	f.Add([]byte("xyz"), int32(0))
	f.Add([]byte("xyz"), int32(math.MinInt32))
	f.Add([]byte("xyz"), int32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, k []byte, v int32) {
		vs1 := struct {
			Elem1 []byte
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 []byte
			Elem2 int32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 []byte
			Elem2 int32
			Elem3 []byte
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzByteInt64Struct(f *testing.F) {
	f.Add([]byte("x"), int64(0))
	f.Add([]byte("x"), int64(math.MinInt64))
	f.Add([]byte("x"), int64(math.MaxInt64))
	f.Add([]byte("xy"), int64(0))
	f.Add([]byte("xy"), int64(math.MinInt64))
	f.Add([]byte("xy"), int64(math.MaxInt64))
	f.Add([]byte("xyz"), int64(0))
	f.Add([]byte("xyz"), int64(math.MinInt64))
	f.Add([]byte("xyz"), int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k []byte, v int64) {
		vs1 := struct {
			Elem1 []byte
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 []byte
			Elem2 int64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 []byte
			Elem2 int64
			Elem3 []byte
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzByteUintStruct(f *testing.F) {
	f.Add([]byte("x"), uint(0))
	f.Add([]byte("x"), uint(math.MaxUint))
	f.Add([]byte("xy"), uint(0))
	f.Add([]byte("xy"), uint(math.MaxUint))
	f.Add([]byte("xyz"), uint(0))
	f.Add([]byte("xyz"), uint(math.MaxUint))
	f.Fuzz(func(t *testing.T, k []byte, v uint) {
		vs1 := struct {
			Elem1 []byte
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 []byte
			Elem2 uint
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 []byte
			Elem2 uint
			Elem3 []byte
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzByteUint8Struct(f *testing.F) {
	f.Add([]byte("x"), uint8(0))
	f.Add([]byte("x"), uint8(math.MaxUint8))
	f.Add([]byte("xy"), uint8(0))
	f.Add([]byte("xy"), uint8(math.MaxUint8))
	f.Add([]byte("xyz"), uint8(0))
	f.Add([]byte("xyz"), uint8(math.MaxUint8))
	f.Fuzz(func(t *testing.T, k []byte, v uint8) {
		vs1 := struct {
			Elem1 []byte
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 []byte
			Elem2 uint8
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 []byte
			Elem2 uint8
			Elem3 []byte
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzByteUint16Struct(f *testing.F) {
	f.Add([]byte("x"), uint16(0))
	f.Add([]byte("x"), uint16(math.MaxUint16))
	f.Add([]byte("xy"), uint16(0))
	f.Add([]byte("xy"), uint16(math.MaxUint16))
	f.Add([]byte("xyz"), uint16(0))
	f.Add([]byte("xyz"), uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, k []byte, v uint16) {
		vs1 := struct {
			Elem1 []byte
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 []byte
			Elem2 uint16
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 []byte
			Elem2 uint16
			Elem3 []byte
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzByteUint32Struct(f *testing.F) {
	f.Add([]byte("x"), uint32(0))
	f.Add([]byte("x"), uint32(math.MaxUint32))
	f.Add([]byte("xy"), uint32(0))
	f.Add([]byte("xy"), uint32(math.MaxUint32))
	f.Add([]byte("xyz"), uint32(0))
	f.Add([]byte("xyz"), uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, k []byte, v uint32) {
		vs1 := struct {
			Elem1 []byte
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 []byte
			Elem2 uint32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 []byte
			Elem2 uint32
			Elem3 []byte
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzByteUint64Struct(f *testing.F) {
	f.Add([]byte("x"), uint64(0))
	f.Add([]byte("x"), uint64(math.MaxInt64))
	f.Add([]byte("xy"), uint64(0))
	f.Add([]byte("xy"), uint64(math.MaxInt64))
	f.Add([]byte("xyz"), uint64(0))
	f.Add([]byte("xyz"), uint64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, k []byte, v uint64) {
		vs1 := struct {
			Elem1 []byte
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 []byte
			Elem2 uint64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 []byte
			Elem2 uint64
			Elem3 []byte
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzByteFloat32Struct(f *testing.F) {
	f.Add([]byte("x"), float32(-math.MaxFloat32))
	f.Add([]byte("x"), float32(0))
	f.Add([]byte("x"), float32(math.MaxFloat32))
	f.Add([]byte("xy"), float32(-math.MaxFloat32))
	f.Add([]byte("xy"), float32(0))
	f.Add([]byte("xy"), float32(math.MaxFloat32))
	f.Add([]byte("xyz"), float32(-math.MaxFloat32))
	f.Add([]byte("xyz"), float32(0))
	f.Add([]byte("xyz"), float32(math.MaxFloat32))
	f.Fuzz(func(t *testing.T, k []byte, v float32) {
		vs1 := struct {
			Elem1 []byte
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 []byte
			Elem2 float32
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 []byte
			Elem2 float32
			Elem3 []byte
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzByteFloat64Struct(f *testing.F) {
	f.Add([]byte("x"), float64(-math.MaxFloat32))
	f.Add([]byte("x"), float64(0))
	f.Add([]byte("x"), float64(math.MaxFloat64))
	f.Add([]byte("xy"), float64(-math.MaxFloat32))
	f.Add([]byte("xy"), float64(0))
	f.Add([]byte("xy"), float64(math.MaxFloat64))
	f.Add([]byte("xyz"), float64(-math.MaxFloat32))
	f.Add([]byte("xyz"), float64(0))
	f.Add([]byte("xyz"), float64(math.MaxFloat64))
	f.Fuzz(func(t *testing.T, k []byte, v float64) {
		vs1 := struct {
			Elem1 []byte
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 []byte
			Elem2 float64
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 []byte
			Elem2 float64
			Elem3 []byte
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzByteBoolStruct(f *testing.F) {
	f.Add([]byte("x"), bool(true))
	f.Add([]byte("x"), bool(false))
	f.Add([]byte("xy"), bool(true))
	f.Add([]byte("xy"), bool(false))
	f.Add([]byte("xyz"), bool(true))
	f.Add([]byte("xyz"), bool(false))
	f.Fuzz(func(t *testing.T, k []byte, v bool) {
		vs1 := struct {
			Elem1 []byte
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 []byte
			Elem2 bool
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 []byte
			Elem2 bool
			Elem3 []byte
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzByteStringStruct(f *testing.F) {
	f.Add([]byte("x"), string("a"))
	f.Add([]byte("x"), string("ab"))
	f.Add([]byte("x"), string("abc"))
	f.Add([]byte("xy"), string("a"))
	f.Add([]byte("xy"), string("ab"))
	f.Add([]byte("xy"), string("abc"))
	f.Add([]byte("xyz"), string("a"))
	f.Add([]byte("xyz"), string("ab"))
	f.Add([]byte("xyz"), string("abc"))
	f.Fuzz(func(t *testing.T, k []byte, v string) {
		vs1 := struct {
			Elem1 []byte
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 []byte
			Elem2 string
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 []byte
			Elem2 string
			Elem3 []byte
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}

// nolint: dupl, maligned
func FuzzByteByteStruct(f *testing.F) {
	f.Add([]byte("x"), []byte("x"))
	f.Add([]byte("x"), []byte("xy"))
	f.Add([]byte("x"), []byte("xyz"))
	f.Add([]byte("xy"), []byte("x"))
	f.Add([]byte("xy"), []byte("xy"))
	f.Add([]byte("xy"), []byte("xyz"))
	f.Add([]byte("xyz"), []byte("x"))
	f.Add([]byte("xyz"), []byte("xy"))
	f.Add([]byte("xyz"), []byte("xyz"))
	f.Fuzz(func(t *testing.T, k []byte, v []byte) {
		vs1 := struct {
			Elem1 []byte
		}{
			Elem1: k,
		}
		fuzzTest(t, vs1)
		vs2 := struct {
			Elem1 []byte
			Elem2 []byte
		}{
			Elem1: k, Elem2: v,
		}
		fuzzTest(t, vs2)
		vs3 := struct {
			Elem1 []byte
			Elem2 []byte
			Elem3 []byte
		}{
			Elem1: k, Elem2: v, Elem3: k,
		}
		fuzzTest(t, vs3)
	})
}
