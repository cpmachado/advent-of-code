package day02

import "fmt"

func Command(args []string, second bool) {
	filename := args[0]
	reports := ParseFile(filename)
	fmt.Printf("Safe report count: %d\n", CountSafeReports(reports, second))
}
