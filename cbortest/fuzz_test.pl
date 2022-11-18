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
	"encoding/hex"
	"math"
	"testing"

	"github.com/cybergarage/go-cbor/cbor"
)

func fuzzUnmarshalTest(t *testing.T, v any) {
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

	err = deepEqual(v, r)
	if err != nil {
		t.Error(err)
		return
	}
}

HEADER
# Go Fuzzing - The Go Programming Language
# https://go.dev/security/fuzz/

my @types = (
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
	"[]byte",
	);

my @seeds = (
	["0", "math.MinInt", "math.MaxInt"],
	["0", "math.MinInt8", "math.MaxInt8"],
	["0", "math.MinInt16", "math.MaxInt16"],
	["0", "math.MinInt32", "math.MaxInt32"],
	["0", "math.MinInt64", "math.MaxInt64"],
	["0", "math.MaxUint"],
	["0", "math.MaxUint8"],
	["0", "math.MaxUint16"],
	["0", "math.MaxUint32"],
	["0", "math.MaxInt64"],
	["-math.MaxFloat32", "0", "math.MaxFloat32"],
	["-math.MaxFloat32", "0", "math.MaxFloat64"],
	["true", "false"],
	["\"a\"", "\"ab\"", "\"abc\""],
	["\"x\"", "\"xy\"", "\"xyz\""],
	);

sub to_fuzz_name {
	my ($fuzz_name) = @_;
	$fuzz_name =~ s/\[\]//g;
	$fuzz_name = ucfirst($fuzz_name);
	return $fuzz_name;
}

sub is_exclude_from_map {
	my ($type_name) = @_;
	if ($type_name == "[]byte") {
		return 1;
	}
	return 0;
}

########################################
# Primitive data type tests
########################################

for (my $i = 0; $i <= $#types; $i++){
	printf("\n");
	my $type = $types[$i];
	printf("func Fuzz%sData(f *testing.F) {\n", to_fuzz_name($type));
	for ($j = 0; $j < @{$seeds[$i]}; $j++) {
		printf("\tf.Add(%s(%s))\n", $type, $seeds[$i]->[$j]);
    }
	printf("\tf.Fuzz(func(t *testing.T, v %s) {\n", $type);
	printf("\t\tfuzzUnmarshalTest(t, v)\n");
	printf("\t})\n");
	printf("}\n");
}

########################################
# Array tests
########################################

for (my $i = 0; $i <= $#types; $i++){
	printf("\n");
	my $type = $types[$i];
	printf("func Fuzz%sArray(f *testing.F) {\n", to_fuzz_name($type));
	for ($j = 0; $j < @{$seeds[$i]}; $j++) {
		printf("\tf.Add(%s(%s))\n", $type, $seeds[$i]->[$j]);
    }
	printf("\tf.Fuzz(func(t *testing.T, v %s) {\n", $type);
	printf("\t\tva := []%s{}\n", $type);
	printf("\t\tfuzzUnmarshalTest(t, va)\n");
	printf("\t\tva = []%s{v}\n", $type);
	printf("\t\tfuzzUnmarshalTest(t, va)\n");
	printf("\t\tva = []%s{v, v}\n", $type);
	printf("\t\tfuzzUnmarshalTest(t, va)\n");
	printf("\t\tva = []%s{v, v, v}\n", $type);
	printf("\t\tfuzzUnmarshalTest(t, va)\n");
	printf("\t})\n");
	printf("}\n");
}

########################################
# Map tests
########################################

for (my $i = 0; $i <= $#types; $i++){
	my $itype = $types[$i];
	if (is_exclude_from_map($itype)) {
		next;
	}
	for (my $j = 0; $j <= $#types; $j++){
		my $jtype = $types[$j];
		if (is_exclude_from_map($jtype)) {
			next;
		}
		printf("\n");
		printf("// nolint: dupl\n");
		printf("func Fuzz%s%sMap(f *testing.F) {\n", to_fuzz_name($itype), to_fuzz_name($jtype));
		for (my $n = 0; $n < @{$seeds[$i]}; $n++) {
			for (my $m = 0; $m < @{$seeds[$j]}; $m++) {
				printf("\tf.Add(%s(%s), %s(%s))\n", $itype, $seeds[$i]->[$n], $jtype, $seeds[$j]->[$m]);
	    	}
    	}
		printf("\tf.Fuzz(func(t *testing.T, k1 %s, v1 %s) {\n", $itype, $jtype);
		# printf("\tf.Fuzz(func(t *testing.T, k1 %s, k2 %s, k3 %s, v1 %s, v2 %s, v3 %s) {\n", $itype, $itype, $itype, $jtype, $jtype, $jtype);
		printf("\t\tvm := map[%s]%s{k1: v1}\n", $itype, $jtype);
		printf("\t\tfuzzUnmarshalTest(t, vm)\n");
		# printf("\t\tvm = map[%s]%s{k1: v1, k2: v2}\n", $itype, $jtype);
		# printf("\t\tfuzzUnmarshalTest(t, vm)\n");
		# printf("\t\tvm = map[%s]%s{k1: v1, k2: v2, k3: v3}\n", $itype, $jtype);
		# printf("\t\tfuzzUnmarshalTest(t, vm)\n");
		printf("\t})\n");
		printf("}\n");
	}
}

########################################
# Struct tests
########################################

for (my $i = 0; $i <= $#types; $i++){
	my $itype = $types[$i];
	for (my $j = 0; $j <= $#types; $j++){
		my $jtype = $types[$j];
		printf("\n");
		printf("// nolint: dupl, maligned\n");
		printf("func Fuzz%s%sStruct(f *testing.F) {\n", to_fuzz_name($itype), to_fuzz_name($jtype));
		for (my $n = 0; $n < @{$seeds[$i]}; $n++) {
			for (my $m = 0; $m < @{$seeds[$j]}; $m++) {
				printf("\tf.Add(%s(%s), %s(%s))\n", $itype, $seeds[$i]->[$n], $jtype, $seeds[$j]->[$m]);
	    	}
    	}
		printf("\tf.Fuzz(func(t *testing.T, k %s, v %s) {\n", $itype, $jtype);
		printf("\t\tvs1 := struct {\n");
		printf("\t\t\tElem1 %s\n", $itype);
		printf("\t\t}{\n");
		printf("\t\t\tElem1: k,\n");
		printf("\t\t}\n");
		printf("\t\tfuzzUnmarshalTest(t, vs1)\n");
		printf("\t\tvs2 := struct {\n");
		printf("\t\t\tElem1 %s\n", $itype);
		printf("\t\t\tElem2 %s\n", $jtype);
		printf("\t\t}{\n");
		printf("\t\t\tElem1: k, Elem2: v,\n");
		printf("\t\t}\n");
		printf("\t\tfuzzUnmarshalTest(t, vs2)\n");
		printf("\t\tvs3 := struct {\n");
		printf("\t\t\tElem1 %s\n", $itype);
		printf("\t\t\tElem2 %s\n", $jtype);
		printf("\t\t\tElem3 %s\n", $itype);
		printf("\t\t}{\n");
		printf("\t\t\tElem1: k, Elem2: v, Elem3: k,\n");
		printf("\t\t}\n");
		printf("\t\tfuzzUnmarshalTest(t, vs3)\n");
		printf("\t})\n");
		printf("}\n");
	}
}
