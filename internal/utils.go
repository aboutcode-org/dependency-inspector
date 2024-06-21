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

	"github.com/bmatcuk/doublestar/v4"
)

func CreateLockFile(lockFiles []string, cmdArgs []string, lockGenCmd []string, outputFileName string, forced bool) {
	absPath := getAbsPath(cmdArgs)
	if absPath == "" {
		return
	}

	if !forced {
		// If any lockfile is present already then skip lockfile generation.
		for _, lockFile := range lockFiles {
			lockFileAbsPath := filepath.Join(absPath, lockFile)

			if res := DoesFileExists(lockFileAbsPath); !res {
				continue
			}
			return
		}
	}
	genLock(lockGenCmd, absPath, outputFileName)

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

		fmt.Printf("Lockfile '%s' already present.\n", relPath)
		return true
	}
	return false
}

func getAbsPath(cmdArgs []string) string {
	path := "."
	if len(cmdArgs) > 0 {
		path = cmdArgs[0]
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: Failed to retrieve absolute path: ", err)
		return ""
	}

	return absPath
}

func genLock(lockGenCmd []string, absPath string, outputFileName string) {
	fmt.Printf("Generating lockfile at '%s' using '%s'\n", absPath, lockGenCmd)

	// #nosec G204
	command := exec.Command(lockGenCmd[0], lockGenCmd[1:]...)
	command.Dir = absPath
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout
	if outputFileName != "" {

		outputPath := filepath.Join(absPath, outputFileName)

		// #nosec G304
		outputFile, err := os.Create(outputPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: failed to create output file: ", err)
			os.Exit(1)
		}
		defer outputFile.Close()

		command.Stdout = outputFile
	}

	if err := command.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error: Failed to generate lockfile: ", err)
		return
	}

	fmt.Println("Lock file generated successfully.")
}

func CreateLockFileNuGet(cmdArgs []string, force bool) {
	nuGetLockFileName := "packages.lock.json"
	nuGetLockFileGenCmd := []string{"dotnet", "restore", "--use-lock-file"}

	project_path := getAbsPath(cmdArgs)
	if project_path == "" {
		return
	}

	fs := os.DirFS(project_path)
	csproj_pattern := "**/*.csproj"

	csproj_files, _ := doublestar.Glob(fs, csproj_pattern)
	if len(csproj_files) == 0 {
		fmt.Fprintln(os.Stderr, "Error: Path does not contain a NuGet project")
		return
	}

	// Generate lockfile for all NuGet projects
	for _, file := range csproj_files {
		fullPath := filepath.Join(project_path, file)
		dir := filepath.Dir(fullPath)

		lockFile := filepath.Join(dir, nuGetLockFileName)
		if force || !DoesFileExists(lockFile) {
			genLock(nuGetLockFileGenCmd, dir, "")
		}

	}

}
