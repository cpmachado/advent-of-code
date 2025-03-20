package day03

import (
	"fmt"
	"os"
)

func Command(args []string, second bool) error {
	if len(args) != 1 {
		return fmt.Errorf("day03 requires a filename")
	}
	filename := args[0]
	input, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	str := string(input)

	if !second {
		fmt.Printf("Program resulted in: %d\n", ProgramInterpreter(str))
	} else {
		fmt.Printf("Program resulted in: %d\n", ProgramInterpreterImproved(str))
	}
	return nil
}
