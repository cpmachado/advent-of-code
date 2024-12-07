/* Copyright © 2024 cpmachado <cpmachado@protonmail> */
package day01

import (
	"log"
)

type LocationId int

// Absolute number for integers
func iAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// Sum of LocationIds
func iSum(s []LocationId) int {
	sum := 0
	for _, v := range s {
		sum += int(v)
	}
	return sum
}

// Solves first question, as multiplication is commutative, one can simply add
// up all, compute difference and get the absolute value of it.
// Computes the difference of the sum between two lists
func ListDiff(lst, other []LocationId) int {
	if len(lst) != len(other) {
		log.Fatal("There's an assumption of the lists being the same size\n")
	}

	return iAbs(iSum(lst) - iSum(other))
}
