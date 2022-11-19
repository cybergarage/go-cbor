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

sub is_exclude_from_unmarshalto {
	my ($type_name) = @_;
	if ($type_name == "int") {
		return 0;
	}
	return 1;
}

########################################
# Primitive benchmarking
########################################

printf("func Benchmark%sData(b *testing.B) {\n", to_fuzz_name($type));
for (my $i = 0; $i <= $#types; $i++){
	my $type = $types[$i];
	printf("\tfor n:= 0; n < b.N; n++ {\n");
	printf("\t\tvar v %s\n", $type);
	for ($j = 0; $j < @{$seeds[$i]}; $j++) {
		printf("\t\tv = %s(%s)\n", $type, $seeds[$i]->[$j]);
	printf("\t\tunmarshalProfile(v)\n");
    }
	printf("\t}\n");
}
printf("}\n");
