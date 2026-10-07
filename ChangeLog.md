# Changelog

## Unreleased

### Breaking changes
- Encoder::Encode() and Marshal() now use the preferred serialization of RFC 8949 Section 4.1 by default: integers, lengths, and floating-point values use the shortest form, so decoded values may have smaller Go types (for example int(1) is decoded as int8(1)). Use EncodeModeTypePreserving for the previous behavior
- MapSortEnabled now sorts map keys in the bytewise lexicographic order of their encodings (RFC 8949 Section 4.2.1) instead of the string order of fmt.Sprintf("%v")

### Changes
- Added EncodeMode and Config::SetEncodeMode() with EncodeModePreferred, EncodeModeTypePreserving, EncodeModeCoreDeterministic (RFC 8949 Section 4.2.1), and EncodeModeLengthFirstDeterministic (RFC 8949 Section 4.2.3)
- Fixed Encoder::Encode() to encode lengths of 255, 65535, and 4294967295 in the shortest form
- Fixed Encoder::Encode() to return ErrEncode instead of writing duplicate map keys when different Go keys have the same encoding in sorted maps
- Fixed Decoder::Unmarshal() and UnmarshalTo() panicking when unmarshaling nested arrays such as [][]byte
- Fixed Decoder::Decode() to decode negative integers (major type 1) correctly for all argument values, returning *big.Int for values less than math.MinInt64
- Fixed Decoder::Decode() to return io.ErrUnexpectedEOF for truncated data items instead of silently returning wrong values
- Fixed Decoder::Decode() to return an error instead of panicking for map keys that cannot be used as Go map keys (arrays, maps, and byte strings)
- Fixed Decoder::Decode() to reject lengths that overflow int and to avoid preallocating memory for untrusted byte and text string lengths
- Added indefinite-length decoding for byte strings, text strings, arrays, and maps (RFC 8949 Section 3.2)
- Added Tag type; Decoder::Decode() returns Tag for tag numbers it does not interpret instead of an error, and Encoder::Encode() encodes Tag and *Tag (RFC 8949 Section 3.4)
- Added support for tag numbers encoded with 1 to 8 byte arguments
- Updated Decoder::Decode() to skip the self-described CBOR tag (55799)
- Updated Decoder::Decode() to return ErrDecode for a misplaced break stop code (0xFF) or an invalid chunk in an indefinite-length string
- Added SimpleValue type and Undefined; Decoder::Decode() decodes all simple values (RFC 8949 Section 3.3) and Encoder::Encode() encodes SimpleValue
- Updated Decoder::Decode() to decode epoch-based date/time (tag 1) into time.Time in UTC (RFC 8949 Section 3.4.2)
- Updated Decoder::Decode() to decode bignums (tags 2 and 3) into *big.Int, and Encoder::Encode() to encode big.Int and *big.Int (RFC 8949 Section 3.4.3)
- Updated Encoder::Encode() to keep fractional seconds when encoding time.Time as tag 0
- Fixed Encoder::Encode() to encode struct fields in declaration order instead of random order
- Fixed Encoder::Encode() to skip unexported struct fields instead of panicking

## v1.3.3 (2026-02-03)
- Updated go-safecast package from v1.3.4 to v1.3.5
- Fix golangci-lint issues

## v1.3.2 (2025-08-08)
- Updated go-safecast package from v1.3.3 to v1.3.4
- Fix golangci-lint issues
- Improved test coverage

## v1.3.1 (2023-05-22)
- Updated Decoder::Unmarshal() and UnmarshalTo() to unmarshal map fields to structure fields when conversion is possible

## v1.3.0 (2023-05-20)
- Updated Decoder::Unmarshal() and UnmarshalTo() to support nested destination structures
- Updated Decoder::Unmarshal() and UnmarshalTo() to convert as flexibly as possible to destination structures
- Fixed Decoder::Unmarshal() and UnmarshalTo() to not panic when an invalid destination struct is passed

## v1.2.1 (2023-05-08)
- Added Config for Encoder and Decoder
- Added MapSortEnabled config for testing
- Update go version to 1.20 for generics

## v1.2.0 (2022-11-20)
- Updated Decoder::Unmarshal() and UnmarshalTo() to unmarshal decoded objects into basic primitive data types.

## v1.1.1 (2022-11-18)
- Improved performance of Decoder::Unmarshal() when a map is specified as the unmarshal object.
- Improved Decoder::Unmarshal() to expand the slice capacity automatically if the specified array is shorter than the decorded array.
- Added fuzzing tests

## v1.1.0 (2022-11-16)
- Updated Encoder::Encode() and Marshal() to support any user-defined maps, arrays, and structures
- Added Decoder::Unmarshal() and UnmarshalTo() to unmarshal decoded objects into any user-defined maps and structures
###  Supported
- Go
  - struct

## v1.0.0 (2022-11-13)
- Initial release  
###  Supported
- CBOR
  - 0 (unsigned integer)
  - 1 (negative integer)
  - 2 (byte string)
  - 3 (text string)
  - 4 (array)
  - 5 (map)
  - 6 (tag)
    - 0 (Date/Time)
  - 7 (Simple)
    - 20 (false), 21 (true), 22 (null)
  - 7 (Floating-point)
    - 26 (IEEE 754 Single-Precision)
    - 27 (IEEE 754 Double-Precision)
- Go
  - int, int8, int16, int32, int64
  - uint, uint8, uint16, uint32, uint64
  - []byte
  - string
  - floag32, float64
  - bool
  - nil
  - array ([]any)
  - map (map[any]any)
  - time.Time
 
