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

// 3.2. Indefinite Lengths for Some Major Types
//
// The indefinite-length encoding is used when the length of a string, an array, or a map is not
// known when its encoding starts, such as when streaming data. The deterministic encode modes
// do not allow indefinite-length items (RFC 8949 Section 4.2.1).

// StartIndefiniteByteString starts an indefinite-length byte string.
// Each following Encode([]byte) writes a chunk until EndIndefinite is called.
func (enc *Encoder) StartIndefiniteByteString() error {
	return enc.startIndefinite(mtBytes)
}

// StartIndefiniteTextString starts an indefinite-length text string.
// Each following Encode(string) writes a chunk until EndIndefinite is called.
func (enc *Encoder) StartIndefiniteTextString() error {
	return enc.startIndefinite(mtText)
}

// StartIndefiniteArray starts an indefinite-length array.
// Each following Encode writes an element until EndIndefinite is called.
func (enc *Encoder) StartIndefiniteArray() error {
	return enc.startIndefinite(mtArray)
}

// StartIndefiniteMap starts an indefinite-length map.
// The following Encode calls write keys and values alternately until EndIndefinite is called.
func (enc *Encoder) StartIndefiniteMap() error {
	return enc.startIndefinite(mtMap)
}

// EndIndefinite ends the innermost indefinite-length item by writing the "break" stop code.
func (enc *Encoder) EndIndefinite() error {
	n := len(enc.indefinite)
	if n == 0 {
		return newErrorEncodeIndefinite("no indefinite-length item to end")
	}
	item := enc.indefinite[n-1]
	if item.majorType == mtMap && item.elements%2 != 0 {
		return newErrorEncodeIndefinite("the indefinite-length map has a key without a value")
	}
	if err := writeByte(enc.writer, breakCode); err != nil {
		return err
	}
	enc.indefinite = enc.indefinite[:n-1]
	return enc.afterIndefiniteElement()
}

// IndefiniteDepth returns the number of indefinite-length items that are started and not ended yet.
func (enc *Encoder) IndefiniteDepth() int {
	return len(enc.indefinite)
}

func (enc *Encoder) startIndefinite(mt majorType) error {
	switch enc.EncodeMode() {
	case EncodeModeCoreDeterministic, EncodeModeLengthFirstDeterministic:
		return newErrorEncodeIndefinite("indefinite-length items are not allowed in the deterministic encode modes")
	case EncodeModePreferred, EncodeModeTypePreserving:
	}
	// Chunks of indefinite-length strings must be definite-length strings (RFC 8949 Section 3.2.3).
	if err := enc.checkIndefiniteElement(); err != nil {
		return err
	}
	if err := writeHeader(enc.writer, mt, aiIndefinite); err != nil {
		return err
	}
	enc.indefinite = append(enc.indefinite, &indefiniteItem{majorType: mt, elements: 0})
	return nil
}

// checkIndefiniteElement returns an error if a new indefinite-length item cannot be started
// in the innermost indefinite-length item.
func (enc *Encoder) checkIndefiniteElement() error {
	n := len(enc.indefinite)
	if n == 0 {
		return nil
	}
	switch enc.indefinite[n-1].majorType { //nolint:exhaustive
	case mtBytes, mtText:
		return newErrorEncodeIndefinite("a chunk of an indefinite-length string must be a definite-length string")
	}
	return nil
}

// beforeIndefiniteElement checks the specified item as an element of the innermost indefinite-length item.
func (enc *Encoder) beforeIndefiniteElement(item any) error {
	n := len(enc.indefinite)
	if n == 0 {
		return nil
	}
	switch enc.indefinite[n-1].majorType { //nolint:exhaustive
	case mtBytes:
		if _, ok := item.([]byte); !ok {
			return newErrorEncodeIndefinite("a chunk of an indefinite-length byte string must be []byte")
		}
	case mtText:
		if _, ok := item.(string); !ok {
			return newErrorEncodeIndefinite("a chunk of an indefinite-length text string must be string")
		}
	}
	enc.indefinite[n-1].elements++
	return nil
}

// afterIndefiniteElement counts an ended indefinite-length item as an element of its parent.
func (enc *Encoder) afterIndefiniteElement() error {
	if n := len(enc.indefinite); 0 < n {
		enc.indefinite[n-1].elements++
	}
	return nil
}
