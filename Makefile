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
	$(GOCMD) get -u golang.org/x/tools/cmd/goimports

valid:
	$(GOFMT) -w .
	$(GOIMPORTS) -w .

.PHONY: build clean test dev valid
