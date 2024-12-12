// Package cmd hosts the various subcommands associated with each day of the
// Advent of Code day
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// Second is a flag if the program should consider it's the second part
var Second bool

var rootCmd = &cobra.Command{
	Use:   "aoc2024",
	Short: "Command to run AOC 2024 solutions",
	Long:  "Command to run AOC 2024 solutions",
}

// Execute executes the root command
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Second, "second", "s", false, "second part")
}
