package day01

import (
	"slices"
	"testing"
)

const exampleFile = "example_file.txt"

func TestParseFile(t *testing.T) {
	gotA, gotB := ParseFile(exampleFile)
	// List of Ids
	wantA := []LocationId{3, 4, 2, 1, 3, 3}
	// Other List of Ids
	wantB := []LocationId{4, 3, 5, 3, 9, 3}

	if !slices.Equal(gotA, wantA) {
		t.Fatalf("Expected '%q' but got '%q'", wantA, gotA)
	}

	if !slices.Equal(gotB, wantB) {
		t.Fatalf("Expected '%q' but got '%q'", wantB, gotB)
	}

}
