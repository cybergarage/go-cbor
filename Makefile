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

SHELL := bash

MODULE_ROOT=github.com/cybergarage/go-cbor

PKG_NAME=cbor
PKG_ID=${MODULE_ROOT}/${PKG_NAME}
PKG_SRC_DIR=${PKG_NAME}
PKG_SRCS=\
        ${PKG_SRC_DIR}
PKGS=\
	${PKG_ID}

TEST_PKG_NAME=${PKG_NAME}test
TEST_PKG_ID=${MODULE_ROOT}/${TEST_PKG_NAME}
TEST_PKG_DIR=${TEST_PKG_NAME}
TEST_PKG_SRCS=\
	${TEST_PKG_DIR}
TEST_PKGS=\
	${TEST_PKG_ID}

.PHONY: version format vet lint clean

all: test

format: version
	gofmt -s -w ${PKG_SRC_DIR} ${BIN_DIR} ${TEST_PKG_DIR}

vet: format
	go vet ${PKG_ID} ${TEST_PKG_ID} ${BINS}

lint: vet
	golangci-lint run ${PKG_SRCS} ${BIN_SRCS} ${TEST_PKG_SRCS}

build: fuzz
	go build -v ${PKGS}

test: lint
	go test -v -coverpkg=${PKG_ID} -timeout 60s ${PKGS} ${TEST_PKGS}

fuzz: test
	pushd ${TEST_PKG_DIR} && make && popd
	pushd ${TEST_PKG_DIR} && ./fuzz && popd

clean:
	go clean -i ${PKGS}
