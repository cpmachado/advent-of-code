package cmd

import (
	"github.com/cpmachado/advent-of-code/2024/day03"
	"github.com/spf13/cobra"
)

var day03Cmd = &cobra.Command{
	Use:   "day03",
	Short: "Command to parse and solve day03 of AOC 2024",
	Long:  "Command to parse and solve day03 of AOC 2024",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		day03.Command(args, Second)
	},
}

func init() {
	rootCmd.AddCommand(day03Cmd)
}
