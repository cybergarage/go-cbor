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

package cbor

const (
	majorTypeMask = 0xE0
	addInfoMask   = 0x1F
)

type majorType byte

const (
	// 3.1. Major Types.
	Uint  majorType = 0x00
	NInt  majorType = 0x20
	Bytes majorType = 0x40
	Text  majorType = 0x60
	Array majorType = 0x80
	Map   majorType = 0xA0
	Tag   majorType = 0xC0
	Float majorType = 0xE0
)

type majorInfo byte

const (
	// 3. Specification of the CBOR Encoding.
	aiOneByte   majorInfo = 24
	aiTwoByte   majorInfo = 25
	aiFourByte  majorInfo = 26
	aiEightByte majorInfo = 27
	// 3.3. Floating-Point Numbers and Values with No Content.
	fpnFloat16 majorInfo = 25
	fpnFloat32 majorInfo = 26
	fpnFloat64 majorInfo = 27
	simpFalse  majorInfo = 20
	simpTrue   majorInfo = 21
	simpNull   majorInfo = 22
	// 3.4. Tagging of Items.
	tagStdDateTime   majorInfo = 0
	tagEpochDateTime majorInfo = 1
)
