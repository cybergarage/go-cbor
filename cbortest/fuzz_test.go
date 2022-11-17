// Copyright (C) 2022 The go-cbor Authors All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//b
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cbortest

import (
	"testing"

	"github.com/cybergarage/go-cbor/cbor"
)

func fuzzPrimitiveTest[T comparable](t *testing.T, v T) {
	t.Helper()
	bytes, err := cbor.Marshal(v)
	if err != nil {
		t.Errorf("%v : %s", v, err)
		return
	}
	r, err := cbor.Unmarshal(bytes)
	if err != nil {
		t.Errorf("%v : %s", v, err)
		return
	}

	err = DeepEqual(v, r)
	if err != nil {
		t.Error(err)
		return
	}
}

func FuzzPrimitiveByte(f *testing.F) {
	f.Fuzz(func(t *testing.T, v byte) {
		fuzzPrimitiveTest(t, v)
	})
}

func FuzzPrimitiveInt(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int) {
		fuzzPrimitiveTest(t, v)
	})
}

func FuzzPrimitiveInt8(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int8) {
		fuzzPrimitiveTest(t, v)
	})
}

func FuzzPrimitiveInt16(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int16) {
		fuzzPrimitiveTest(t, v)
	})
}

func FuzzPrimitiveInt32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int32) {
		fuzzPrimitiveTest(t, v)
	})
}

func FuzzPrimitiveInt64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int64) {
		fuzzPrimitiveTest(t, v)
	})
}

func FuzzPrimitiveUint(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint) {
		fuzzPrimitiveTest(t, v)
	})
}

func FuzzPrimitiveUint8(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint8) {
		fuzzPrimitiveTest(t, v)
	})
}

func FuzzPrimitiveUint16(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint16) {
		fuzzPrimitiveTest(t, v)
	})
}

func FuzzPrimitiveUint32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint32) {
		fuzzPrimitiveTest(t, v)
	})
}

func FuzzPrimitiveUint64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint64) {
		fuzzPrimitiveTest(t, v)
	})
}

func FuzzPrimitiveFloat32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v float32) {
		fuzzPrimitiveTest(t, v)
	})
}

func FuzzPrimitiveFloat64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v float64) {
		fuzzPrimitiveTest(t, v)
	})
}

func FuzzPrimitiveBool(f *testing.F) {
	f.Fuzz(func(t *testing.T, v bool) {
		fuzzPrimitiveTest(t, v)
	})
}

func FuzzPrimitiveString(f *testing.F) {
	f.Fuzz(func(t *testing.T, v string) {
		fuzzPrimitiveTest(t, v)
	})
}
