package cmd

import (
	"github.com/cpmachado/advent-of-code/2024/day02"
	"github.com/spf13/cobra"
)

var day02Cmd = &cobra.Command{
	Use:   "day02",
	Short: "Command to parse and solve day02 of AOC 2024",
	Long:  "Command to parse and solve day02 of AOC 2024",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		day02.Command(args, Second)
	},
}

func init() {
	rootCmd.AddCommand(day02Cmd)
}
