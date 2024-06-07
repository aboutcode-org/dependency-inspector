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
GOSEC=gosec
BINARY_NAME=deplock
BUILD_DIR=./build
BINARY_OUTPUT=$(BUILD_DIR)/$(BINARY_NAME)


build:
	$(GOCMD) build -o $(BINARY_OUTPUT) -v

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

gofmt:
	@echo "-> Apply gofmt code formatter"
	$(GOFMT) -w .

goimports:
	@echo "-> Apply goimports changes to ensure proper imports ordering"
	$(GOIMPORTS) -w .

valid: goimports gofmt

check:
	@echo "-> Running goimports for import ordering validation..."
	$(GOIMPORTS) -d -e .
	@echo "\n-> Running gofmt for code formatting validation..."
	$(GOFMT) -d -e .
	@echo "\n-> Running golangci-lint for linting..."
	$(GOLINT) run ./...
	@echo "\n-> Running gosec for security checks..."
	$(GOSEC) ./...

.PHONY: build clean test dev gofmt goimports valid check
