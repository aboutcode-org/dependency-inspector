/*

Copyright (c) nexB Inc. and others. All rights reserved.
ScanCode is a trademark of nexB Inc.
SPDX-License-Identifier: Apache-2.0
See http://www.apache.org/licenses/LICENSE-2.0 for the license text.
See https://github.com/aboutcode-org/dependency-inspector for support or download.
See https://aboutcode.org for more information about nexB OSS projects.

*/

package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/aboutcode-org/dependency-inspector/internal"
	"github.com/spf13/cobra"
)

func goCmd() *cobra.Command {
	forced := false

	goCmd := &cobra.Command{
		Use:   "go [path]",
		Short: "Generate lockfile for go project",
		Long: `Create lockfile (go.list.json) for Go project if it doesn't exist in the specified [path].
If no path is provided, the command defaults to the current directory.`,
		Args: cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			path := "."
			if len(args) > 0 {
				path = args[0]
			}

			absPath, err := filepath.Abs(path)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error: Failed to retrieve absolute path: ", err)
				return
			}

			outputFileName := "go.list.json"
			outputPath := filepath.Join(absPath, outputFileName)

			if !forced {
				if internal.DoesFileExists(outputPath) {
					return
				}
			}

			fmt.Printf("Generating lockfile at '%s' using '[go list -m -json all]'\n", absPath)

			// #nosec G204
			command := exec.Command("go", "list", "-m", "-json", "all")
			command.Dir = absPath
			command.Stderr = os.Stderr

			out, err := command.Output()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error: Failed to execute go list command: ", err)
				return
			}

			// Parse NDJSON to valid JSON Array
			// 'go list -m -json all' outputs objects separated by newline, e.g.:
			// { ... }
			// { ... }
			strOut := strings.TrimSpace(string(out))
			if strOut == "" {
				strOut = "[]"
			} else {
				// Replace "}\n{" with "},{"
				strOut = strings.ReplaceAll(strOut, "}\n{", "},{")
				strOut = "[" + strOut + "]"
			}

			err = os.WriteFile(outputPath, []byte(strOut), 0644)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error: failed to write output file: ", err)
				os.Exit(1)
			}

			fmt.Println("Lock file generated successfully.")
		},
	}

	goCmd.Flags().BoolVarP(&forced, "force", "f", false, "Generate lockfile forcibly, ignoring existing lockfiles")

	return goCmd
}
