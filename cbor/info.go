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
	aiOneByte   addInfo = 24
	aiTwoByte   addInfo = 25
	aiFourByte  addInfo = 26
	aiEightByte addInfo = 27
	// 3.3. Floating-Point Numbers and Values with No Content.
	fpnFloat16 addInfo = 25
	fpnFloat32 addInfo = 26
	fpnFloat64 addInfo = 27
	simpFalse  addInfo = 20
	simpTrue   addInfo = 21
	simpNull   addInfo = 22
	// 3.4. Tagging of Items.
	tagStdDateTime   = 0
	tagEpochDateTime = 1
)
