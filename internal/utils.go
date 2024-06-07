/*

Copyright (c) nexB Inc. and others. All rights reserved.
ScanCode is a trademark of nexB Inc.
SPDX-License-Identifier: Apache-2.0
See http://www.apache.org/licenses/LICENSE-2.0 for the license text.
See https://github.com/nexB/dependency-inspector for support or download.
See https://aboutcode.org for more information about nexB OSS projects.

*/

package internal

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func CreateLockFile(lockFiles []string, cmdArgs []string, lockGenCmd []string) {

	path := "."
	if len(cmdArgs) > 0 {
		path = cmdArgs[0]
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Failed to retrieve absolute path: %v\n", err)
		os.Exit(1)
	}

	for _, lockFile := range lockFiles {
		lockFileAbsPath := filepath.Join(absPath, lockFile)

		if res := DoesFileExists(lockFileAbsPath); res {
			continue
		}
		break

	}
	genLock(lockGenCmd, absPath)

}

func DoesFileExists(absPath string) bool {
	if _, err := os.Stat(absPath); err == nil {

		cwd, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}

		relPath, err := filepath.Rel(cwd, absPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}

		fmt.Printf("Lockfile '%s' already present.", relPath)
		return true
	}
	return false
}

func genLock(lockGenCmd []string, absPath string) {
	fmt.Printf("Generating lockfile using '%s'\n", lockGenCmd)

	// #nosec G204
	command := exec.Command(lockGenCmd[0], lockGenCmd[1:]...)
	command.Dir = absPath
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error: Failed to generate lockfile: ", err)
		return
	}

	fmt.Println("Lock file generated successfully.")
}
