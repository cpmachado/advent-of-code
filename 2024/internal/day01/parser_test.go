package day01

import (
	"slices"
	"testing"
)

const exampleFile = "example_file.txt"

func TestParseFile(t *testing.T) {
	a, b := ParseFile(exampleFile)
	got := []LocationIDList{a, b}
	want := []LocationIDList{
		{3, 4, 2, 1, 3, 3},
		{4, 3, 5, 3, 9, 3},
	}

	for i, v := range got {
		if !slices.Equal(v, want[i]) {
			t.Fatalf("Expected '%q' but got '%q'", want[i], v)
		}
	}
}
