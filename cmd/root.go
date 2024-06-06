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
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var ecosystems = []func() *cobra.Command{
	pnpmCmd,
	npmCmd,
	yarnCmd,
}

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "deplock",
		Short: "DepLock: Dependency Locking CLI",
	}

	initConfig(rootCmd)
	for _, subCmd := range ecosystems {
		rootCmd.AddCommand(subCmd())
	}

	return rootCmd
}

func initConfig(rootCmd *cobra.Command) {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
}

func Execute() {
	rootCmd := NewRootCmd()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
