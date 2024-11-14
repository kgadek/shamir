// Copyright (c) Konrad Gądek
// SPDX-License-Identifier: MPL-2.0

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var combineCmd = &cobra.Command{
	Use: "combine",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("combine called")
	},
}

func init() {
	rootCmd.AddCommand(combineCmd)
}
