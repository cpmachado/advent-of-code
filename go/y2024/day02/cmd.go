package day02

import "fmt"

func Command(args []string, second bool) error {
	if len(args) != 1 {
		return fmt.Errorf("day02 requires a filename")
	}
	filename := args[0]
	reports := ParseFile(filename)
	fmt.Printf("Safe report count: %d\n", CountSafeReports(reports, second))
	return nil
}
