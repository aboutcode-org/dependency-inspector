/*

Copyright (c) nexB Inc. and others. All rights reserved.
ScanCode is a trademark of nexB Inc.
SPDX-License-Identifier: Apache-2.0
See http://www.apache.org/licenses/LICENSE-2.0 for the license text.
See https://github.com/nexB/dependency-inspector for support or download.
See https://aboutcode.org for more information about nexB OSS projects.

*/

package cmd

import (
	"github.com/nexB/dependency-inspector/internal"
	"github.com/spf13/cobra"
)

func swiftCmd() *cobra.Command {
	lockFiles := [][]string{
		{"Package.resolved", ".package.resolved"},
		{"Package.swift.json"},
	}
	lockGenCommands := [][]string{
		{"swift", "package", "resolve"},
		{"swift", "package", "dump-package"},
	}
	commandOutput := []string{
		"",
		"Package.swift.json",
	}
	forced := false

	swiftCmd := &cobra.Command{
		Use:   "swift [path]",
		Short: "Generate lockfile for swift project",
		Long: `Create lockfile and JSON dump of manifest for swift project
if they doesn't already exist in the specified [path].
If no path is provided, the command defaults to the current directory.`,
		Args: cobra.MaximumNArgs(1),

		Run: func(cmd *cobra.Command, args []string) {

			for i := range lockFiles {
				internal.CreateLockFile(
					lockFiles[i],
					args,
					lockGenCommands[i],
					commandOutput[i],
					forced,
				)
			}

		},
	}

	swiftCmd.Flags().BoolVarP(&forced, "force", "f", false, "Generate lockfile forcibly, ignoring existing lockfiles")

	return swiftCmd
}
