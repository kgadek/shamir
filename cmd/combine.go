// Copyright (c) Konrad Gądek
// SPDX-License-Identifier: MPL-2.0

package cmd

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"shamir-vault/shamir"
	"strings"

	"github.com/spf13/cobra"
)

var combineCmd = &cobra.Command{
	Use: "combine",
	Run: func(cmd *cobra.Command, args []string) {
		dataB, err := io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}

		dataS := string(dataB)
		lines := strings.Split(dataS, "\n")

		data := [][]byte{}

		for _, line := range lines {
			if line == "" {
				continue
			}

			decoded, err := hex.DecodeString(line)
			if err != nil {
				panic(err)
			}
			data = append(data, decoded)
		}

		decoded, err := shamir.Combine(data)
		if err != nil {
			panic(err)
		}

		result := strings.Builder{}
		result.Write(decoded)
		fmt.Print(result.String())
	},
}

func init() {
	rootCmd.AddCommand(combineCmd)
}
