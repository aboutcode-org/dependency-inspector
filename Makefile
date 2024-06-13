# SPDX-License-Identifier: Apache-2.0
#
# Copyright (c) nexB Inc. and others. All rights reserved.
# ScanCode is a trademark of nexB Inc.
# SPDX-License-Identifier: Apache-2.0
# See http://www.apache.org/licenses/LICENSE-2.0 for the license text.
# See https://github.com/nexB/dependency-inspector for support or download.
# See https://aboutcode.org for more information about nexB OSS projects.
#

GOCMD=go
GOFMT=gofmt
GOIMPORTS=goimports
GOLINT=golangci-lint
GOFMT_CMD = $(GOFMT) -l .
GOIMPORTS_CMD = $(GOIMPORTS) -l .
GOSEC=gosec
BINARY_NAME=deplock
BUILD_DIR=./build
GOOS := $(shell $(GOCMD)  env GOOS)
GOARCH := $(shell $(GOCMD) env GOARCH)
BINARY_OUTPUT=$(BUILD_DIR)/$(BINARY_NAME)-$(GOOS)-$(GOARCH)
BUILD_FLAGS=
STRIP_FLAG=

strip ?= no

# Omit symbol table, debug info and DWARF symbol table
# https://github.com/golang/go/blob/ca5ba146da7a9d4e2a8cbe1715a78be42b45a745/src/cmd/link/doc.go#L115-L123
ifeq ($(strip), yes)
	BUILD_FLAGS=-s -w
	STRIP_FLAG=--strip
endif


build:
	@echo "Building binary to $(BINARY_OUTPUT)"
	$(GOCMD) build -ldflags "$(BUILD_FLAGS)" -o $(BINARY_OUTPUT)

build-all:
	./scripts/build-all.sh $(STRIP_FLAG)

clean:
	$(GOCMD) clean
	rm -rf $(BUILD_DIR)

test:
	$(GOCMD) test -v ./test

# Install dev dependency, makes sure  $HOME/go/bin is in PATH
dev:
	$(GOCMD) install golang.org/x/tools/cmd/goimports@latest
	$(GOCMD) install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.0
	$(GOCMD) install github.com/securego/gosec/v2/cmd/gosec@latest
	$(GOCMD) mod tidy

gofmt:
	@echo "-> Apply gofmt code formatter"
	$(GOFMT) -w .

goimports:
	@echo "-> Apply goimports changes to ensure proper imports ordering"
	$(GOIMPORTS) -w .

valid: goimports gofmt

check-gofmt:
	@echo "-> Running gofmt for code formatting validation..."
	@files=$$($(GOFMT_CMD)); \
	if [ -n "$$files" ]; then \
		echo "The following files are not properly formatted:"; \
		echo "$$files"; \
		exit 1; \
	fi

check-goimports:
	@echo "-> Running goimports for import ordering validation..."
	@files=$$($(GOIMPORTS_CMD)); \
	if [ -n "$$files" ]; then \
		echo "The following files have incorrect imports:"; \
		echo "$$files"; \
		exit 1; \
	fi

check: check-gofmt check-goimports
	@echo "\n-> Running golangci-lint for linting..."
	$(GOLINT) run --issues-exit-code=1 ./...
	@echo "\n-> Running gosec for security checks..."
	$(GOSEC) ./...

.PHONY: build build-all clean test dev gofmt goimports valid check-gofmt check-goimports check
