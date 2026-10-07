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

// SimpleValue represents a CBOR simple value (major type 7) as defined in RFC 8949 Section 3.3.
// Decoder returns false, true, and nil for the simple values 20, 21, and 22,
// and SimpleValue for the other simple values, such as Undefined (23).
// Simple values 24 to 31 are reserved and cannot be encoded.
type SimpleValue uint8

// Undefined represents the CBOR "undefined" simple value (23).
const Undefined SimpleValue = 23
