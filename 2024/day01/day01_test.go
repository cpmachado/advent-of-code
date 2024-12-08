/* Copyright © 2024 cpmachado <cpmachado@protonmail> */
package day01

import "testing"

func TestDifferenceScore(t *testing.T) {
	// List of Ids
	lst := LocationIdList{3, 4, 2, 1, 3, 3}
	// Other List of Ids
	other := LocationIdList{4, 3, 5, 3, 9, 3}
	got := lst.DifferenceScore(other)
	want := 11

	if got != want {
		t.Fatalf("Expected difference to be '%d' but got '%d'", want, got)
	}
}

func TestListSimilarityScore(t *testing.T) {
	// List of Ids
	lst := LocationIdList{3, 4, 2, 1, 3, 3}
	// Other List of Ids
	other := LocationIdList{4, 3, 5, 3, 9, 3}
	got := lst.SimilarityScore(other)
	want := 31

	if got != want {
		t.Fatalf("Expected difference to be '%d' but got '%d'", want, got)
	}
}
