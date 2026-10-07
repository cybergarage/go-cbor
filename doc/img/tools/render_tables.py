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
"""Rasterize render.pdf (one table per page) and crop each page to a table image.

Usage: render_tables.py render.pdf out_dir
"""
import os
import subprocess
import sys
import tempfile

from PIL import Image, ImageOps

# The page order matches the sheet order in build_tables.py.
NAMES = ["conv_table_from", "conv_table_to", "unmarshal_table_to", "unmarshal_table_to_basic",
         "unmarshal_table_to_special"]
DPI = 128
PADDING = 4


def main():
    pdf, out_dir = sys.argv[1], sys.argv[2]
    with tempfile.TemporaryDirectory() as work:
        subprocess.run(["pdftoppm", "-r", str(DPI), "-png", pdf, os.path.join(work, "page")], check=True)
        pages = sorted(f for f in os.listdir(work) if f.endswith(".png"))
        if len(pages) != len(NAMES):
            sys.exit("expected %d pages, got %d" % (len(NAMES), len(pages)))
        for name, page in zip(NAMES, pages):
            im = Image.open(os.path.join(work, page)).convert("RGB")
            left, top, right, bottom = ImageOps.invert(im.convert("L")).getbbox()
            box = (max(left - PADDING, 0), max(top - PADDING, 0),
                   min(right + PADDING, im.width), min(bottom + PADDING, im.height))
            path = os.path.join(out_dir, name + ".png")
            im.crop(box).save(path, optimize=True)
            print(path)


if __name__ == "__main__":
    main()
