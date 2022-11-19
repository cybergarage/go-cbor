# Quick Start

The section briefly describes how to converte data between Go and CBOR using `go-cbor`.

## Encoding

To convert data from Go to CBOR, `go-cbor` offers `Marshal()`. `Marshal()` converts from the specified data model of Go into the equivalent data model of CBOR as the following.

```
goObjs := []any{
    uint(1000),
    int(-1000),
    float32(100000.0),
    float64(-4.1),
    false,
    true,
    nil,
    []byte("IETF"),
    "IETF",
    []int{1, 2, 3},
    map[any]any{"a": "A", "b": "B", "c": "C"},
}
for _, goObj := range goObjs {
    cborBytes, _ := cbor.Marshal(goObj)
    fmt.Printf("%v => %s\n", goObj, hex.EncodeToString(cborBytes))
}
```

In addition to the basic Go data types, `go-cbor` supports additional tag major types such as `time.Time` as the following.

```
goObj, _ := time.Parse(time.RFC3339, "2013-03-21T20:04:00Z")
cborBytes, _ := cbor.Marshal(goObj)
fmt.Printf("%v => %s\n", goObj, hex.EncodeToString(cborBytes))
```

## Decoding

To convert data from CBOR to Go, `go-cbor` offers `Unmarshal()`. `Unmarshal()` converts from an encoded bytes of CBOR into the equivalent data model of Go as the following.

```
cborObjs := []string{
    "0a",
    "1903e8",
    "3903e7",
    "fb3ff199999999999a",
    "f90001",
    "f4",
    "f5",
    "f6",
    "c074323031332d30332d32315432303a30343a30305a",
    "4449455446",
    "6449455446",
    "83010203",
    "a201020304",
}
for _, cborObj := range cborObjs {
    cborBytes, _ := hex.DecodeString(cborObj)
    goObj, _ := cbor.Unmarshal(cborBytes)
    fmt.Printf("%s => %v\n", cborObj, goObj)
}
```

To unmarshal to a user-defined struct, `go-cbor` offers `UnmarshalTo()`. `Unmarshal()To` tries to convert from an encoded bytes of CBOR into the specified user-defined struct or map as the following.

```
fromObjs := []any{
    struct {
        Key   string
        Value string
    }{
        Key: "hello", Value: "world",
    },
    map[string]int{"one": 1, "two": 2},
}

toObjs := []any{
    &struct {
        Key   string
        Value string
    }{},
    map[string]int{},
}

for n, fromObj := range fromObjs {
    toObj := toObjs[n]
    encBytes, _ := cbor.Marshal(fromObj)
    cbor.UnmarshalTo(encBytes, toObj)
}
```
