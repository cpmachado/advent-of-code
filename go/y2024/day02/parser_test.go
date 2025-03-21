package day02

import (
	"slices"
	"testing"
)

const exampleFile = "example_file.txt"

func TestParseFile(t *testing.T) {
	got := ParseFile(exampleFile)
	// Slice of Reports
	want := []Report{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	for i, report := range want {
		if !slices.Equal(got[i], report) {
			t.Fatalf("Expected '%v' but got '%v'", want, got)
		}
	}
}
