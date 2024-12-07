/* Copyright © 2024 cpmachado <cpmachado@protonmail> */
package day01

import "testing"

func TestListDiff(t *testing.T) {
	// List of Ids
	listA := []LocationId{3, 4, 2, 1, 3, 3}
	// Other List of Ids
	listB := []LocationId{4, 3, 5, 3, 9, 3}
	got := ListDiff(listA, listB)
	want := 11

	if got != want {
		t.Fatalf("Expected difference to be '%d' but got '%d'", want, got)
	}
}

func TestListSimilarityScore(t *testing.T) {
	// List of Ids
	listA := []LocationId{3, 4, 2, 1, 3, 3}
	// Other List of Ids
	listB := []LocationId{4, 3, 5, 3, 9, 3}
	got := ListSimilarityScore(listA, listB)
	want := 31

	if got != want {
		t.Fatalf("Expected difference to be '%d' but got '%d'", want, got)
	}
}
