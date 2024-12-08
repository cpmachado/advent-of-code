/* Copyright © 2024 cpmachado <cpmachado@protonmail> */
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/cpmachado/advent-of-code/2024/day03"
	"github.com/spf13/cobra"
)

// day03Cmd represents the day03 command
var day03Cmd = &cobra.Command{
	Use:   "day03",
	Short: "Command to parse and solve day03 of AOC 2024",
	Long:  "Command to parse and solve day03 of AOC 2024",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		input, err := ioutil.ReadFile(filename)

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
