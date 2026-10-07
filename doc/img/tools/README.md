# Conversion Table Tools

These tools regenerate the conversion tables in `doc/img` from the current behavior of go-cbor:

| File | Description |
|---|---|
| `conv_table.ods` | Editable source spreadsheet (sheets `conv_table` and `unmarshal_table`) |
| `conv_table_from.png` | Encoding: Go to CBOR |
| `conv_table_to.png` | Decoding: CBOR to Go |
| `unmarshal_table_to.png` | Unmarshaling: all destination types (not referenced from the documents) |
| `unmarshal_table_to_basic.png` | Unmarshaling: basic destination types |
| `unmarshal_table_to_special.png` | Unmarshaling: other destination types |

## Usage

```
./doc/img/tools/update.sh
```

`update.sh` runs the following steps and overwrites the files in `doc/img`:

1. `measure/main.go` decodes a representative data item for each row of the decoding and unmarshaling tables, unmarshals it into each destination type, and prints the results as JSON.
2. `build_tables.py` builds `conv_table.ods` and a render-only spreadsheet with one table per sheet from the JSON. The encoding table is defined in `enc_rows()` in this script, because it describes the encoder rather than measured results.
3. LibreOffice converts the render-only spreadsheet to PDF, and `render_tables.py` rasterizes each page at 128 DPI and crops it to the table.

The output is deterministic: running `update.sh` without behavior changes reproduces the same images. `conv_table.ods` may differ only in its metadata, so revert it if its content did not change.

## Requirements

- Go (see `go.mod`)
- Python 3 with `odfpy` and `Pillow`
- LibreOffice (`soffice`) and poppler (`pdftoppm`)
- The [Source Code Pro](https://github.com/adobe-fonts/source-code-pro) font (SIL Open Font License), which is also available as the npm package `source-code-pro`

## Updating the tables

- When the decoder or the unmarshaler changes, update the rows in `measure/main.go` (representative data items) or the destination types in `dests()`.
- When the encoder changes, update `enc_rows()` in `build_tables.py`.
- The footnotes (`*1`, `*2`) are explained in the text under the images in `README.md` and `doc/conversion.md`; update them together.
