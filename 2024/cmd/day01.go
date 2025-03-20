package cmd

import (
	"github.com/cpmachado/advent-of-code/2024/day01"
	"github.com/spf13/cobra"
)

var day01Cmd = &cobra.Command{
	Use:   "day01",
	Short: "Command to parse and solve day01 of AOC 2024",
	Long:  "Command to parse and solve day01 of AOC 2024",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		day01.Command(args, Second)
	},
}

func init() {
	rootCmd.AddCommand(day01Cmd)
}
