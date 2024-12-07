/* Copyright © 2024 cpmachado <cpmachado@protonmail> */
package day01

import (
	"log"
	"math"
	"slices"
)

type LocationId int

// Solves first question, as sum is commutative, one can simply add
// up all, compute difference and get the absolute value of it.
// Computes the difference of the sum between two lists
func ListDiff(lst, other []LocationId) int {
	if len(lst) != len(other) {
		log.Fatal("There's an assumption of the lists being the same size\n")
	}
	slices.Sort(lst)
	slices.Sort(other)
	sum := 0

	for i, v := range lst {
		sum += int(math.Abs(float64(v - other[i])))
	}

	return sum
}
