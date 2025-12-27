package day07

import (
	"fmt"
	"os"
)

func Command(args []string, second bool) error {
	if len(args) != 1 {
		return fmt.Errorf("day07 requires a filename")
	}
	filename := args[0]

	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	count, err := Process(file, second)
	if err != nil {
		return err
	}

	fmt.Printf("Count: %d\n", count)
	return nil
}
