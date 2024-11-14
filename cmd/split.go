// Copyright (c) Konrad Gądek
// SPDX-License-Identifier: MPL-2.0

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// splitCmd represents the split command
var splitCmd = &cobra.Command{
	Use: "split",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("split called")
	},
}

func init() {
	rootCmd.AddCommand(splitCmd)
}
