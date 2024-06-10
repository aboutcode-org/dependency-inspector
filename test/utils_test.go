/*

Copyright (c) nexB Inc. and others. All rights reserved.
ScanCode is a trademark of nexB Inc.
SPDX-License-Identifier: Apache-2.0
See http://www.apache.org/licenses/LICENSE-2.0 for the license text.
See https://github.com/nexB/dependency-inspector for support or download.
See https://aboutcode.org for more information about nexB OSS projects.

*/

package main

import (
	"path/filepath"
	"testing"

	"github.com/nexB/dependency-inspector/internal"
)

func TestDoesFileExists(t *testing.T) {
	dataDir, err := filepath.Abs("data")
	if err != nil {
		t.Fatalf("Error getting absolute path for data directory: %v", err)
	}

	filePath := filepath.Join(dataDir, "package-lock.json")
	exists := internal.DoesFileExists(filePath)
	if !exists {
		t.Errorf("Expected file to not exist, but it does")
	}

	filePath = filepath.Join(dataDir, "tXN6iXJlTf.txt")
	exists = internal.DoesFileExists(filePath)
	if exists {
		t.Errorf("Expected file not to exist, but it does.")
	}
}
