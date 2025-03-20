// Package day01 contains code relating to AOC day01 of 2024
package day01

import "slices"

// LocationIDList represents a list of LocationID
type LocationIDList []int

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// DifferenceScore computes the DifferenceScore between 2 lists
func (lst LocationIDList) DifferenceScore(other LocationIDList) int {
	slices.Sort(lst)
	slices.Sort(other)
	sum := 0

	for i, v := range lst {
		sum += intAbs(v - other[i])
	}

	return sum
}

// SimilarityScore computes SimilarityScore between 2 lists
func (lst LocationIDList) SimilarityScore(other LocationIDList) int {
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
