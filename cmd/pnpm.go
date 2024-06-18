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

func pnpmCmd() *cobra.Command {
	lockFiles := []string{"pnpm-lock.yaml", ".pnpm-lock.yaml"}
	lockGenCommand := []string{"pnpm", "i", "--lockfile-only"}
	forced := false

	pnpmCmd := &cobra.Command{
		Use:   "pnpm [path]",
		Short: "Generate lockfile for pnpm project",
		Long: `Create lockfile for pnpm project if it doesn't exist in the specified [path].
If no path is provided, the command defaults to the current directory.`,
		Args: cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			internal.CreateLockFile(
				lockFiles,
				args,
				lockGenCommand,
				"",
				forced,
			)
		},
	}

	pnpmCmd.Flags().BoolVarP(&forced, "force", "f", false, "Generate lockfile forcibly, ignoring existing lockfiles")

	return pnpmCmd
}
