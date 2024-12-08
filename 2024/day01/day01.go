/* Copyright © 2024 cpmachado <cpmachado@protonmail> */
package day01

import (
	"slices"

	"github.com/cpmachado/advent-of-code/2024/common/util"
)

// Sorts provided slices, and computes the sum of differences
func ListDiff(lst, other []int) int {
	slices.Sort(lst)
	slices.Sort(other)
	sum := 0

	for i, v := range lst {
		sum += util.IntAbs(v - other[i])
	}

	return sum
}

// Computes SimilarityScore between 2 lists
func ListSimilarityScore(lst, other []int) int {
	counter := make(map[int]int)

	// Compute count of second list
	for _, v := range other {
		counter[v]++
	}
	sum := 0

	for _, v := range lst {
		sum += int(v) * counter[v]
	}

	return sum
}
