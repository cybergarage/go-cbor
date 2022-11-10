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
	"fmt"
	"io"
)

// An Encoder writes CBOR values to an output stream.
type Encoder struct {
	writer io.Writer
}

// NewEncoder returns a new encoder that writes to the specified writer.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		writer: w,
	}
}

// Encode writes the specified object to the specified writer.
func (enc *Encoder) Encode(item any) error {
	// 3. Specification of the CBOR Encoding.
	switch v := item.(type) {
	case uint8:
		header := byte(Uint)
		if v < 24 {
			header |= v
			return writeByte(enc.writer, header)
		}
		header |= byte(uIntOneByte)
		if err := writeByte(enc.writer, header); err != nil {
			return err
		}
		return writeByte(enc.writer, v)
	case int:
		return nil
	case string:
		if _, err := io.WriteString(enc.writer, v); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf(errorNotSupportedDataType, item)
}
