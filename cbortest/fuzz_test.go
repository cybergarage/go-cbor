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
	"fmt"
	"math"
	"testing"

	"github.com/cybergarage/go-cbor/cbor"
)

func fuzzTes(t *testing.T, v any) {
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
	fuzzTes(t, v)
}

func FuzzInt(f *testing.F) {
	f.Add(int(0))
	f.Add(int(math.MinInt))
	f.Add(int(math.MaxInt))
	f.Fuzz(func(t *testing.T, v int) {
		t.Run(fmt.Sprintf("%v", v), func(t *testing.T) {
			fuzzPrimitiveTest(t, v)
		})
	})
}

func FuzzInt8(f *testing.F) {
	f.Add(int8(0))
	f.Add(int8(math.MinInt8))
	f.Add(int8(math.MaxInt8))
	f.Fuzz(func(t *testing.T, v int8) {
		t.Run(fmt.Sprintf("%v", v), func(t *testing.T) {
			fuzzPrimitiveTest(t, v)
		})
	})
}

func FuzzInt16(f *testing.F) {
	f.Add(int16(0))
	f.Add(int16(math.MinInt16))
	f.Add(int16(math.MaxInt16))
	f.Fuzz(func(t *testing.T, v int16) {
		t.Run(fmt.Sprintf("%v", v), func(t *testing.T) {
			fuzzPrimitiveTest(t, v)
		})
	})
}

func FuzzInt32(f *testing.F) {
	f.Add(int32(0))
	f.Add(int32(math.MinInt32))
	f.Add(int32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, v int32) {
		t.Run(fmt.Sprintf("%v", v), func(t *testing.T) {
			fuzzPrimitiveTest(t, v)
		})
	})
}

func FuzzInt64(f *testing.F) {
	f.Add(int64(0))
	f.Add(int64(math.MinInt64))
	f.Add(int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, v int64) {
		t.Run(fmt.Sprintf("%v", v), func(t *testing.T) {
			fuzzPrimitiveTest(t, v)
		})
	})
}

func FuzzUint(f *testing.F) {
	f.Add(uint(0))
	f.Add(uint(math.MaxUint))
	f.Fuzz(func(t *testing.T, v uint) {
		t.Run(fmt.Sprintf("%v", v), func(t *testing.T) {
			fuzzPrimitiveTest(t, v)
		})
	})
}

func FuzzUint8(f *testing.F) {
	f.Add(uint8(0))
	f.Add(uint8(math.MaxUint8))
	f.Fuzz(func(t *testing.T, v uint8) {
		t.Run(fmt.Sprintf("%v", v), func(t *testing.T) {
			fuzzPrimitiveTest(t, v)
		})
	})
}

func FuzzUint16(f *testing.F) {
	f.Add(uint16(0))
	f.Add(uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, v uint16) {
		t.Run(fmt.Sprintf("%v", v), func(t *testing.T) {
			fuzzPrimitiveTest(t, v)
		})
	})
}

func FuzzUint32(f *testing.F) {
	f.Add(uint32(0))
	f.Add(uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, v uint32) {
		t.Run(fmt.Sprintf("%v", v), func(t *testing.T) {
			fuzzPrimitiveTest(t, v)
		})
	})
}

func FuzzUint64(f *testing.F) {
	f.Add(uint64(0))
	f.Add(uint64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, v uint64) {
		t.Run(fmt.Sprintf("%v", v), func(t *testing.T) {
			fuzzPrimitiveTest(t, v)
		})
	})
}

func FuzzFloat32(f *testing.F) {
	f.Add(float32(math.MaxFloat32))
	f.Fuzz(func(t *testing.T, v float32) {
		t.Run(fmt.Sprintf("%v", v), func(t *testing.T) {
			fuzzPrimitiveTest(t, v)
		})
	})
}

func FuzzFloat64(f *testing.F) {
	f.Add(float64(math.MaxFloat64))
	f.Fuzz(func(t *testing.T, v float64) {
		t.Run(fmt.Sprintf("%v", v), func(t *testing.T) {
			fuzzPrimitiveTest(t, v)
		})
	})
}

func FuzzBool(f *testing.F) {
	f.Add(bool(true))
	f.Add(bool(false))
	f.Fuzz(func(t *testing.T, v bool) {
		t.Run(fmt.Sprintf("%v", v), func(t *testing.T) {
			fuzzPrimitiveTest(t, v)
		})
	})
}

func FuzzString(f *testing.F) {
	f.Add(string("abc"))
	f.Add(string("xyz"))
	f.Fuzz(func(t *testing.T, v string) {
		t.Run(fmt.Sprintf("%v", v), func(t *testing.T) {
			fuzzPrimitiveTest(t, v)
		})
	})
}
