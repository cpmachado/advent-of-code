package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/cpmachado/advent-of-code/2024/day03"
	"github.com/spf13/cobra"
)

var day03Cmd = &cobra.Command{
	Use:   "day03",
	Short: "Command to parse and solve day03 of AOC 2024",
	Long:  "Command to parse and solve day03 of AOC 2024",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		filename := args[0]
		input, err := os.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}

		str := string(input)

		if !Second {
			fmt.Printf("Program resulted in: %d\n", day03.ProgramInterpreter(str))
		} else {
			fmt.Printf("Program resulted in: %d\n", day03.ProgramInterpreterImproved(str))
		}
	},
}

func init() {
	rootCmd.AddCommand(day03Cmd)
}
