# Converting Data between Go and CBOR

`go-cobor` was developed as a seamless serializer for the memory representation of any data types in Go like `encodiong/json`. `go-cobor` provides the optimized encoder and decoder to convert between CBOR and Go data models easily.

![](img/concept.png)

This section describes how go-cobor` converts data model between Go and CBOR in more detail.

## Converting from Go to CBOR

`Encoder::Encode()` and `Marshal()` convert from the specified data model of Go into the equivalent data model of CBOR as the following.

![](img/conv_table_from.png)

- \*1 The shortest form (RFC 8949 Section 4.1). With `EncodeModeTypePreserving`, integers use the width of their Go types: `24 or <24` for 8 bits, `25` for 16 bits, `26` for 32 bits, and `27` for 64 bits, `int`, and `uint`.
- \*2 The shortest of `25` (float16), `26` (float32), and `27` (float64) that preserves the value; NaN and infinities use `25`. With `EncodeModeTypePreserving`, `float32` uses `26` and `float64` uses `27`.
- `big.Int`, `url.URL`, and `cbor.Tag` can also be passed as pointers. `struct` encodes exported fields only.

By default, `go-cobor` encodes data items with the preferred serialization of RFC 8949, so decoded values may have smaller Go types than the encoded ones. Use `EncodeModeTypePreserving` to keep the widths of the Go types.

## Converting from CBOR to Go

`Decoder::Decode()` and `Unmarshal()` convert from the specified data model of CBOR into the equivalent data model of Go as the following.

![](img/conv_table_to.png)

- Integers are decoded to the smallest Go type of the argument width that can hold the value.
- Indefinite-length strings, arrays, and maps (additional information 31) are decoded to the same types as the definite-length ones.
- Tag 55799 (self-described CBOR) is skipped, and its content is decoded. Tags 21-24 and 33-36 are returned as `cbor.Tag` after their content is validated.

## Unmarshaling from CBOR to Go

To unmarshal to a user-defined struct, `go-cbor` offers `Decoder::Unmarshal()` and `UnmarshalTo()`. The unmarshal functions try to convert from an encoded bytes of CBOR into the specified basic data types of Go as the following.

![](img/unmarshal_table_to_basic.png)

In addition to the basic standard data types of Go, the unmarshal functions support any user-defined maps and structs, standard types such as `time.Time`, `big.Int`, and `url.URL`, and the go-cbor types `cbor.Tag` and `cbor.SimpleValue` as the following.

![](img/unmarshal_table_to_special.png)

- `O`: the value is converted when it is in the range of the destination type. For example, text strings are converted to numbers and booleans only when they represent them.
- `(bit)?` means `8`, `16`, `32`, `64`, or none, as in `int8` or `int`.
- Values that cannot be converted return `ErrUnmarshal`.
