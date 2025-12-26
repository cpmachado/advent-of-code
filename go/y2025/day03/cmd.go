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

	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	sum, err := Process(file, second)
	if err != nil {
		return err
	}

	fmt.Printf("Sum: %d\n", sum)
	return nil
}
