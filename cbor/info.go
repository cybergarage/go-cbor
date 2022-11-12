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

type addInfo byte

const (
	// 3. Specification of the CBOR Encoding.
	uIntOneByte   addInfo = 24
	uIntTwoByte   addInfo = 25
	uIntFourByte  addInfo = 26
	uIntEightByte addInfo = 27
	// 3.3. Floating-Point Numbers and Values with No Content.
	Float16   addInfo = 25
	Float32   addInfo = 26
	Float64   addInfo = 27
	False     addInfo = 20
	True      addInfo = 21
	Null      addInfo = 22
	Undefined addInfo = 23
)
