// Copyright (c) Konrad Gądek
// SPDX-License-Identifier: MPL-2.0

package cmd

import (
	"fmt"
	"io"
	"os"
	"shamir-vault/shamir"
	"strings"

	"github.com/spf13/cobra"
)

var Parts int
var Threshold int

// splitCmd represents the split command
var splitCmd = &cobra.Command{
	Use: "split",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}

		result, err := shamir.Split(data, Parts, Threshold)
		if err != nil {
			panic(err)
		}

		for _, key := range result {
			result := strings.Builder{}
			for _, num := range key {
				_, err := result.WriteString(fmt.Sprintf("%02x", num))
				if err != nil {
					panic(err)
				}
			}
			fmt.Println(result.String())
		}
	},
}

func init() {
	rootCmd.AddCommand(splitCmd)

	splitCmd.Flags().IntVarP(&Parts, "parts", "p", -1, "Parts")
	splitCmd.MarkFlagRequired("parts")

	splitCmd.Flags().IntVarP(&Threshold, "threshold", "t", -1, "Threshold")
	splitCmd.MarkFlagRequired("threshold")
}
