package day01

import "fmt"

func Command(args []string, second bool) {
	filename := args[0]
	a, b := ParseFile(filename)
	if second {
		fmt.Printf("Similarity Score is %d\n", a.SimilarityScore(b))
	} else {
		fmt.Printf("Difference is %d\n", a.DifferenceScore(b))
	}
}
