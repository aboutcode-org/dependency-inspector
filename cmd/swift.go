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

func swiftCmd() *cobra.Command {
	deplockSwiftManifestDumpFile := "Package.swift.deplock"
	deplockSwiftShowDependenciesFile := "swift-show-dependencies.deplock"

	resolvedLockFiles := []string{"Package.resolved", ".package.resolved"}
	deplockManifestDumpFiles := []string{deplockSwiftManifestDumpFile}
	deplockSwiftShowDependenciesFiles := []string{deplockSwiftShowDependenciesFile}

	resolvedLockGenCommand := []string{"swift", "package", "resolve"}
	deplockManifestDumpGenCommand := []string{"swift", "package", "dump-package"}
	deplockShowDependenciesGenCommand := []string{"swift", "package", "show-dependencies", "--format", "json"}

	forced := false

	swiftCmd := &cobra.Command{
		Use:   "swift [path]",
		Short: "Generate lockfile for swift project",
		Long: `Create lockfile and JSON dump of manifest for swift project
if they doesn't already exist in the specified [path].
If no path is provided, the command defaults to the current directory.`,
		Args: cobra.MaximumNArgs(1),

		Run: func(cmd *cobra.Command, args []string) {

			internal.CreateLockFile(
				resolvedLockFiles,
				args,
				resolvedLockGenCommand,
				"",
				forced,
			)

			internal.CreateLockFile(
				deplockManifestDumpFiles,
				args,
				deplockManifestDumpGenCommand,
				deplockSwiftManifestDumpFile,
				forced,
			)

			internal.CreateLockFile(
				deplockSwiftShowDependenciesFiles,
				args,
				deplockShowDependenciesGenCommand,
				deplockSwiftShowDependenciesFile,
				forced,
			)

		},
	}

	swiftCmd.Flags().BoolVarP(&forced, "force", "f", false, "Generate lockfile forcibly, ignoring existing lockfiles")

	return swiftCmd
}
