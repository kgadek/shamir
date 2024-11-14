// Copyright (c) Konrad Gądek
// SPDX-License-Identifier: MPL-2.0

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "shamir-vault",
	Short: "Shamir secret sharing algorithm, as implemented by Hashicorp Vault",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
