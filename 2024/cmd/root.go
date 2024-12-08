/* Copyright © 2024 cpmachado <cpmachado@protonmail> */
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var Second bool

var rootCmd = &cobra.Command{
	Use:   "aoc2024",
	Short: "Command to run AOC 2024 solutions",
	Long:  "Command to run AOC 2024 solutions",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Second, "second", "s", false, "second part")
}
