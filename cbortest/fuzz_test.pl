#!/usr/bin/perl
# Copyright (C) 2022 The go-cbor Authors All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# #  limitations under the License.

print<<HEADER;
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
HEADER

# Go Fuzzing - The Go Programming Language
# https://go.dev/security/fuzz/

my @priTypes = (
	"byte", 
	"int", 
	"int8", 
	"int16", 
	"int32", 
	"int64",
	"uint", 
	"uint8", 
	"uint16", 
	"uint32", 
	"uint64",
	"float32",
	"float64",
	"bool",
	"string",
	);

foreach my $priType (@priTypes) {
	printf("\n");
	printf("func Fuzz%s(f *testing.F) {\n", ucfirst($priType));
	printf("\tf.Fuzz(func(t *testing.T, v %s) {\n", $priType);
	printf("\t\tt.Run(fmt.Sprintf(\"%%v\", v), func(t *testing.T) {\n");
	printf("\t\t\tfuzzPrimitiveTest(t, v)\n");
	printf("\t\t})\n");
	printf("\t})\n");
	printf("}\n");
}
