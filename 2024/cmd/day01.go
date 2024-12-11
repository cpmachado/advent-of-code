/* Copyright © 2024 cpmachado <cpmachado@protonmail> */
package cmd

import (
	"fmt"

	"github.com/cpmachado/advent-of-code/2024/internal/day01"
	"github.com/spf13/cobra"
)

// day01Cmd represents the day01 command
var day01Cmd = &cobra.Command{
	Use:   "day01",
	Short: "Command to parse and solve day01 of AOC 2024",
	Long:  "Command to parse and solve day01 of AOC 2024",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		a, b := day01.ParseFile(filename)
		if Second {
			fmt.Printf("Similarity Score is %d\n", a.SimilarityScore(b))
		} else {
			fmt.Printf("Difference is %d\n", a.DifferenceScore(b))
		}
	},
}

func init() {
	rootCmd.AddCommand(day01Cmd)
}
