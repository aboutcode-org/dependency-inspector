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
	"github.com/aboutcode-org/dependency-inspector/internal"
	"github.com/spf13/cobra"
)

func nugetCmd() *cobra.Command {
	forced := false

	nugetCmd := &cobra.Command{
		Use:   "nuget [path]",
		Short: "Generate lockfile for NuGet project",
		Long: `Create lockfile for NuGet project if it doesn't exist in the specified [path].
If no path is provided, the command defaults to the current directory.`,
		Args: cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			internal.CreateLockFileNuGet(
				args,
				forced,
			)
		},
	}

	nugetCmd.Flags().BoolVarP(&forced, "force", "f", false, "Generate lockfile forcibly, ignoring existing lockfiles")

	return nugetCmd
}
