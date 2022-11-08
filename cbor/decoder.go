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

import (
	"io"
)

// An Decoder reads CBOR values from an output stream.
type Decoder struct {
	reader io.Reader
	header []byte
}

// NewDecoder returns a new decoder that reads from the specified writer.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		reader: r,
		header: make([]byte, 1),
	}
}

// Decode returns the next item if available, otherwise returns EOF or error.
func (dec *Decoder) Decode() (interface{}, error) {
	if _, err := io.ReadFull(dec.reader, dec.header); err != nil {
		return nil, err
	}

	majorType := MajorType(dec.header[0] & majorTypeMask)
	addInfo := AddInfo(dec.header[0] & addInfoMask)

	switch majorType {
	case Uint:
		// 3. Specification of the CBOR Encoding.
		var v uint64
		switch {
		case addInfo < uIntImmdiate:
			v = uint64(addInfo)
		case addInfo == uIntOneByte:
			v = uint64(addInfo)
		}
		return v, nil
	case Negint:
		return 1, nil
	case Bytes:
		return 1, nil
	case Text:
		return 1, nil
	case Array:
		return 1, nil
	case Map:
		return 1, nil
	case Tag:
		return 1, nil
	case Float:
		return 1, nil
	}

	return nil, io.EOF
}
