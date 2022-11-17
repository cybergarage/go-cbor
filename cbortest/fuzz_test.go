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
	"fmt"
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

func FuzzByte(f *testing.F) {
	f.Fuzz(func(t *testing.T, v byte) {
		t.Run(fmt.Sprintf("%v", v), func(t *testing.T) {
			fuzzPrimitiveTest(t, v)
		})
	})
}

func FuzzInt(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int) {
		t.Run(fmt.Sprintf("%v", v), func(t *testing.T) {
			fuzzPrimitiveTest(t, v)
		})
	})
}

func FuzzInt8(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int8) {
		t.Run(fmt.Sprintf("%v", v), func(t *testing.T) {
			fuzzPrimitiveTest(t, v)
		})
	})
}

func FuzzInt16(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int16) {
		t.Run(fmt.Sprintf("%v", v), func(t *testing.T) {
			fuzzPrimitiveTest(t, v)
		})
	})
}

func FuzzInt32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int32) {
		t.Run(fmt.Sprintf("%v", v), func(t *testing.T) {
			fuzzPrimitiveTest(t, v)
		})
	})
}

func FuzzInt64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int64) {
		t.Run(fmt.Sprintf("%v", v), func(t *testing.T) {
			fuzzPrimitiveTest(t, v)
		})
	})
}

func FuzzUint(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint) {
		t.Run(fmt.Sprintf("%v", v), func(t *testing.T) {
			fuzzPrimitiveTest(t, v)
		})
	})
}

func FuzzUint8(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint8) {
		t.Run(fmt.Sprintf("%v", v), func(t *testing.T) {
			fuzzPrimitiveTest(t, v)
		})
	})
}

func FuzzUint16(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint16) {
		t.Run(fmt.Sprintf("%v", v), func(t *testing.T) {
			fuzzPrimitiveTest(t, v)
		})
	})
}

func FuzzUint32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint32) {
		t.Run(fmt.Sprintf("%v", v), func(t *testing.T) {
			fuzzPrimitiveTest(t, v)
		})
	})
}

func FuzzUint64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint64) {
		t.Run(fmt.Sprintf("%v", v), func(t *testing.T) {
			fuzzPrimitiveTest(t, v)
		})
	})
}

func FuzzFloat32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v float32) {
		t.Run(fmt.Sprintf("%v", v), func(t *testing.T) {
			fuzzPrimitiveTest(t, v)
		})
	})
}

func FuzzFloat64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v float64) {
		t.Run(fmt.Sprintf("%v", v), func(t *testing.T) {
			fuzzPrimitiveTest(t, v)
		})
	})
}

func FuzzBool(f *testing.F) {
	f.Fuzz(func(t *testing.T, v bool) {
		t.Run(fmt.Sprintf("%v", v), func(t *testing.T) {
			fuzzPrimitiveTest(t, v)
		})
	})
}

func FuzzString(f *testing.F) {
	f.Fuzz(func(t *testing.T, v string) {
		t.Run(fmt.Sprintf("%v", v), func(t *testing.T) {
			fuzzPrimitiveTest(t, v)
		})
	})
}
