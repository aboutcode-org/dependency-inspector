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

func pypiCmd() *cobra.Command {
	pipFreezeFile := "requirements.deplock"
	lockFiles := []string{"Pipfile.lock", pipFreezeFile}
	lockGenCommand := []string{"pip", "freeze"}
	forced := false

	pypiCmd := &cobra.Command{
		Use:   "pypi [path]",
		Short: "Generate lockfile for Python project",
		Long: `Create lockfile for Python project if it doesn't exist in the specified [path].
If no path is provided, the command defaults to the current directory.`,
		Args: cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			internal.CreateLockFile(
				lockFiles,
				args,
				lockGenCommand,
				pipFreezeFile,
				forced,
			)
		},
	}

	pypiCmd.Flags().BoolVarP(&forced, "force", "f", false, "Generate lockfile forcibly, ignoring existing lockfiles")

	return pypiCmd
}
