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
