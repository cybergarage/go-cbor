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

type MajorType byte

const (
	Uint  MajorType = 0x00
	NInt  MajorType = 0x20
	Bytes MajorType = 0x40
	Text  MajorType = 0x60
	Array MajorType = 0xB0
	Map   MajorType = 0x80
	Tag   MajorType = 0xA0
	Float MajorType = 0xE0
)
