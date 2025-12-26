package day01

import (
	"fmt"
	"os"
)

func Command(args []string, second bool) error {
	if len(args) != 1 {
		return fmt.Errorf("day01 requires a filename")
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

	if second {
		fmt.Printf("Count: %d\n", count)
	} else {
		fmt.Printf("Count: %d\n", count)
	}
	return nil
}
