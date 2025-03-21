package day01

import "fmt"

func Command(args []string, second bool) error {
	if len(args) != 1 {
		return fmt.Errorf("day01 requires a filename")
	}
	filename := args[0]
	a, b := ParseFile(filename)
	if second {
		fmt.Printf("Similarity Score is %d\n", a.SimilarityScore(b))
	} else {
		fmt.Printf("Difference is %d\n", a.DifferenceScore(b))
	}
	return nil
}
