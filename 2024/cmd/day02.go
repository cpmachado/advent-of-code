package cmd

import (
	"fmt"

	"github.com/cpmachado/advent-of-code/2024/day02"
	"github.com/spf13/cobra"
)

var day02Cmd = &cobra.Command{
	Use:   "day02",
	Short: "Command to parse and solve day02 of AOC 2024",
	Long:  "Command to parse and solve day02 of AOC 2024",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		filename := args[0]
		reports := day02.ParseFile(filename)
		fmt.Printf("Safe report count: %d\n", day02.CountSafeReports(reports, Second))
	},
}

func init() {
	rootCmd.AddCommand(day02Cmd)
}
