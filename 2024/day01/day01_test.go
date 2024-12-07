package day01

import "testing"

func TestListDiff(t *testing.T) {
	// List of Ids
	listA := []int{3, 4, 2, 1, 3, 3}
	// Other List of Ids
	listB := []int{4, 3, 5, 3, 9, 3}
	got := ListDiff(listA, listB)
	want := 11

	if got != want {
		t.Fatalf("Expected difference to be '%d' but got '%d'", want, got)
	}
}
