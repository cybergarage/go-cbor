#!/usr/bin/env python3
# Copyright (C) 2022 The go-cbor Authors All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
"""Build doc/img/conv_table.ods and a render-only ODS for the table images.

Usage: build_tables.py measured.json out_dir
  out_dir/conv_table.ods   editable spreadsheet with the sheets conv_table and unmarshal_table
  out_dir/render.ods       one table per sheet, on a large page without header/footer

The encoding table (Go to CBOR) is defined in enc_rows() below. The decoding and unmarshaling
tables are built from measured.json, which is printed by `go run ./doc/img/tools/measure`.
"""
import json
import sys

from odf.opendocument import OpenDocumentSpreadsheet
from odf.style import (FontFace, MasterPage, PageLayout, PageLayoutProperties, ParagraphProperties, Style,
                       TableCellProperties, TableColumnProperties, TableProperties, TableRowProperties,
                       TextProperties)
from odf.table import CoveredTableCell, Table, TableCell, TableColumn, TableRow
from odf.text import P

FONT = "Source Code Pro"
CHAR_WIDTH_IN = 0.0835  # 10pt Source Code Pro (0.6em)
PADDING_IN = 0.22

measured = json.load(open(sys.argv[1]))
out_dir = sys.argv[2]

# ---------------------------------------------------------------------------
# Table definitions
#
# A table is a dict: name, header rows, body rows, merge columns, left-aligned columns.
# A header cell is (text, colspan). A body row is a list of cell texts, and an optional
# group key per row: vertically adjacent cells with the same text in a merge column are
# merged when they have the same group key in that column.
# ---------------------------------------------------------------------------

UINT_INFO = "<24 - 27 *1"
FLOAT_INFO = "25 - 27 *2"


def enc_rows():
    rows = []  # (cells, groups)
    for t in ["uint8", "uint16", "uint32", "uint64", "uint"]:
        rows.append(([t, "-", "0 (unsigned integer)", UINT_INFO], [t, t, "uint", "uint"]))
    for t in ["int8", "int16", "int32", "int64", "int"]:
        rows.append(([t, ">= 0", "0 (unsigned integer)", UINT_INFO], [t, t, t + "+", t]))
        rows.append(([t, "<  0", "1 (negative integer)", UINT_INFO], [t, t, t + "-", t]))
    big = "big.Int"
    rows += [
        ([big, "0 <= n < 2^64", "0 (unsigned integer)", "<24 - 27"], [big, "b1", "b1", "b1"]),
        ([big, "2^64 <= n", "6 (tag)", "2 (Bignum)"], [big, "b2", "b2", "b2"]),
        ([big, "-2^64 <= n < 0", "1 (negative integer)", "<24 - 27"], [big, "b3", "b3", "b3"]),
        ([big, "n < -2^64", "6 (tag)", "3 (Bignum)"], [big, "b4", "b4", "b4"]),
        (["[]byte", "-", "2 (byte string)", "-"], ["bytes", "bytes", "bytes", "bytes"]),
        (["string", "-", "3 (text string)", "-"], ["text", "text", "text", "text"]),
        (["array, slice", "-", "4 (array)", "-"], ["array", "array", "array", "array"]),
        (["struct", "-", "5 (map)", "-"], ["struct", "struct", "map", "struct"]),
        (["map", "-", "5 (map)", "-"], ["map", "map", "map", "map"]),
        (["time.Time", "year 0-9999", "6 (tag)", "0 (Date/Time)"], ["time", "t0", "tag6", "t0"]),
        (["time.Time", "others", "6 (tag)", "1 (Epoch)"], ["time", "t1", "tag6", "t1"]),
        (["url.URL", "-", "6 (tag)", "32 (URI)"], ["url", "url", "tag6", "url"]),
        (["cbor.Tag", "-", "6 (tag)", "Number"], ["tag", "tag", "tag6", "tag"]),
        (["bool", "false", "7 (simple)", "20"], ["bool", "f", "simple", "f"]),
        (["bool", "true", "7 (simple)", "21"], ["bool", "t", "simple", "t"]),
        (["nil", "-", "7 (simple)", "22"], ["nil", "nil", "simple", "nil"]),
        (["cbor.SimpleValue", "< 24", "7 (simple)", "Value"], ["sv", "sv1", "simple", "sv1"]),
        (["cbor.SimpleValue", ">= 32", "7 (simple)", "24"], ["sv", "sv2", "simple", "sv2"]),
        (["float32", "-", "7 (floating-point)", FLOAT_INFO], ["f32", "f32", "float", "float"]),
        (["float64", "-", "7 (floating-point)", FLOAT_INFO], ["f64", "f64", "float", "float"]),
    ]
    return rows


def dec_groups(r):
    # Merge Major Type over the same type, Info within the same type.
    return [r["mt"], r["mt"] + "/" + r["info"], None, None]


DESTS_BASIC = ["&string", "&int(bit)?", "&uint(bit)?", "&float(bit)", "&bool", "&[]byte"]
DESTS_SPECIAL = ["map[type]type", "&[]type", "&struct", "&time.Time", "&big.Int", "&url.URL", "&cbor.Tag",
                 "&cbor.SimpleValue", "&any"]


def unmarshal_table(name, dests):
    rows = []
    for r in measured["rows"]:
        cells = [r["mt"], r["info"], r["cond"]] + ["O" if r["dests"][d] else "-" for d in dests]
        rows.append((cells, [r["mt"], r["mt"] + "/" + r["info"]] + [None] * (1 + len(dests))))
    return {
        "name": name,
        "header": [[("CBOR", 3), ("Decoder::Unmarshal(), UnmarshalTo()", len(dests))],
                   [("Major Type", 1), ("Info", 1), ("Cond", 1)] + [(d, 1) for d in dests]],
        "rows": rows,
        "merge": {0, 1},
        "left": {0},
    }


TABLES = {
    "conv_table_from": {
        "name": "conv_table_from",
        "header": [[("Go", 2), ("CBOR", 2)], [("From", 1), ("Cond", 1), ("Major Type", 1), ("Info", 1)]],
        "rows": enc_rows(),
        "merge": {0, 1, 2, 3},
        "left": {2},
    },
    "conv_table_to": {
        "name": "conv_table_to",
        "header": [[("CBOR", 2), ("Go", 2)], [("Major Type", 1), ("Info", 1), ("Cond", 1), ("To", 1)]],
        "rows": [([r["mt"], r["info"], r["cond"], r["to"]], dec_groups(r)) for r in measured["rows"]],
        "merge": {0, 1},
        "left": {0},
    },
    "unmarshal_table_to": unmarshal_table("unmarshal_table_to", DESTS_BASIC + DESTS_SPECIAL),
    "unmarshal_table_to_basic": unmarshal_table("unmarshal_table_to_basic", DESTS_BASIC),
    "unmarshal_table_to_special": unmarshal_table("unmarshal_table_to_special", DESTS_SPECIAL),
}


# ---------------------------------------------------------------------------
# Layout: compute cells with spans
# ---------------------------------------------------------------------------

def layout(table):
    """Return (ncols, grid) where grid[r][c] = (text, colspan, rowspan, kind) or None for covered cells."""
    ncols = sum(span for _, span in table["header"][0])
    grid = []
    for hrow in table["header"]:
        row = []
        for text, span in hrow:
            row.append((text, span, 1, "header"))
            row += [None] * (span - 1)
        grid.append(row)
    body = table["rows"]
    nbody = len(body)
    covered = [[False] * ncols for _ in range(nbody)]
    body_grid = [[None] * ncols for _ in range(nbody)]
    for r in range(nbody):
        cells, groups = body[r]
        for c in range(ncols):
            if covered[r][c]:
                continue
            span = 1
            if c in table["merge"] and groups[c] is not None:
                while (r + span < nbody and body[r + span][0][c] == cells[c]
                       and body[r + span][1][c] == groups[c]):
                    span += 1
            kind = "left" if c in table["left"] else "center"
            body_grid[r][c] = (cells[c], 1, span, kind)
            for k in range(1, span):
                covered[r + k][c] = True
    return ncols, grid + body_grid


def column_widths(tables):
    widths = {}
    for table in tables:
        ncols, grid = layout(table)
        for row in grid:
            for c, cell in enumerate(row):
                if cell is None or cell[1] != 1:
                    continue
                w = len(cell[0]) * CHAR_WIDTH_IN + PADDING_IN
                widths[c] = max(widths.get(c, 0), w)
    return widths


# ---------------------------------------------------------------------------
# ODS writer
# ---------------------------------------------------------------------------

class Writer:
    def __init__(self, page=None):
        self.doc = OpenDocumentSpreadsheet()
        self.doc.fontfacedecls.addElement(FontFace(name=FONT, fontfamily="'%s'" % FONT,
                                                   fontfamilygeneric="modern", fontpitch="fixed"))
        self.styles = {}
        border = "1.5pt solid #000000"
        for kind, align, bg in [("header", "center", "#eeeeee"), ("center", "center", None), ("left", "start", None)]:
            st = Style(name="ce_" + kind, family="table-cell")
            props = {"border": border, "verticalalign": "middle"}
            if bg:
                props["backgroundcolor"] = bg
            st.addElement(TableCellProperties(**props))
            st.addElement(ParagraphProperties(textalign=align, marginleft="0.04in" if align == "start" else "0in"))
            st.addElement(TextProperties(fontname=FONT, fontsize="10pt", fontnameasian=FONT, fontnamecomplex=FONT))
            self.doc.automaticstyles.addElement(st)
            self.styles[kind] = st
        row_style = Style(name="ro1", family="table-row")
        row_style.addElement(TableRowProperties(rowheight="0.178in", useoptimalrowheight="true"))
        self.doc.automaticstyles.addElement(row_style)
        self.row_style = row_style
        self.table_style = None
        if page:
            layout_style = PageLayout(name="pm_big")
            layout_style.addElement(PageLayoutProperties(pagewidth=page[0], pageheight=page[1], margin="0.2in",
                                                         printorientation="landscape"))
            self.doc.automaticstyles.addElement(layout_style)
            master = MasterPage(name="Big", pagelayoutname=layout_style)
            self.doc.masterstyles.addElement(master)
            self.table_style = Style(name="ta_big", family="table", masterpagename="Big")
            self.table_style.addElement(TableProperties(display="true", writingmode="lr-tb"))
            self.doc.automaticstyles.addElement(self.table_style)
        self.col_styles = {}

    def col_style(self, width):
        key = "%.3f" % width
        if key not in self.col_styles:
            st = Style(name="co_%d" % len(self.col_styles), family="table-column")
            st.addElement(TableColumnProperties(columnwidth="%sin" % key))
            self.doc.automaticstyles.addElement(st)
            self.col_styles[key] = st
        return self.col_styles[key]

    def add_sheet(self, name, tables, margin=True, gap=2):
        """Add a sheet with the specified tables stacked vertically."""
        kwargs = {"name": name}
        if self.table_style is not None:
            kwargs["stylename"] = self.table_style
        sheet = Table(**kwargs)
        widths = column_widths(tables)
        offset = 1 if margin else 0
        if margin:
            sheet.addElement(TableColumn(stylename=self.col_style(0.2146)))
        for c in sorted(widths):
            sheet.addElement(TableColumn(stylename=self.col_style(widths[c])))
        if margin:
            sheet.addElement(TableRow(stylename=self.row_style))
        for n, table in enumerate(tables):
            if 0 < n:
                for _ in range(gap):
                    sheet.addElement(TableRow(stylename=self.row_style))
            ncols, grid = layout(table)
            for row in grid:
                tr = TableRow(stylename=self.row_style)
                if offset:
                    tr.addElement(TableCell())
                for cell in row:
                    if cell is None:
                        tr.addElement(CoveredTableCell())
                        continue
                    text, colspan, rowspan, kind = cell
                    tc = TableCell(stylename=self.styles[kind], valuetype="string",
                                   numbercolumnsspanned=colspan, numberrowsspanned=rowspan)
                    tc.addElement(P(text=text))
                    tr.addElement(tc)
                sheet.addElement(tr)
        self.doc.spreadsheet.addElement(sheet)

    def save(self, path):
        self.doc.save(path)


editable = Writer()
editable.add_sheet("conv_table", [TABLES["conv_table_from"], TABLES["conv_table_to"]])
editable.add_sheet("unmarshal_table", [TABLES["unmarshal_table_to"], TABLES["unmarshal_table_to_basic"],
                                       TABLES["unmarshal_table_to_special"]])
editable.save(out_dir + "/conv_table.ods")

render = Writer(page=("40in", "20in"))
for name in ["conv_table_from", "conv_table_to", "unmarshal_table_to", "unmarshal_table_to_basic",
             "unmarshal_table_to_special"]:
    render.add_sheet(name, [TABLES[name]], margin=False)
render.save(out_dir + "/render.ods")
print("ok")
