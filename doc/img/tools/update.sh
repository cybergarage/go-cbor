#!/bin/bash
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
#
# Regenerates doc/img/conv_table.ods and the table images from the current behavior of go-cbor.
# Requirements: go, python3 with odfpy and Pillow, LibreOffice (soffice), pdftoppm (poppler),
# and the Source Code Pro font.
set -euo pipefail

TOOLS_DIR=$(cd "$(dirname "$0")" && pwd)
ROOT_DIR=$(cd "${TOOLS_DIR}/../../.." && pwd)
IMG_DIR="${ROOT_DIR}/doc/img"

if ! fc-match "Source Code Pro" | grep -q "Source Code Pro"; then
	echo "warning: Source Code Pro is not installed; the images will use a fallback font" >&2
fi

WORK_DIR=$(mktemp -d)
trap 'rm -rf "${WORK_DIR}"' EXIT

(cd "${ROOT_DIR}" && go run ./doc/img/tools/measure) > "${WORK_DIR}/measured.json"
python3 "${TOOLS_DIR}/build_tables.py" "${WORK_DIR}/measured.json" "${WORK_DIR}"
(cd "${WORK_DIR}" && soffice --headless --convert-to pdf render.ods > /dev/null)
python3 "${TOOLS_DIR}/render_tables.py" "${WORK_DIR}/render.pdf" "${IMG_DIR}"
cp "${WORK_DIR}/conv_table.ods" "${IMG_DIR}/conv_table.ods"
echo "${IMG_DIR}/conv_table.ods"
