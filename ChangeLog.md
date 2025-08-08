# Changelog

## v1.3.2 (2025-08-08)
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
 
