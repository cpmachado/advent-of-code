package day04

import (
	"strings"
	"testing"
)

const sample = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
`

func TestPart1(t *testing.T) {
	// given
	r := strings.NewReader(sample)
	want := 13

	// when
	got, err := Process(r, false)
	// then
	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatalf("Expected %d, but got %d\n", want, got)
	}
}

func TestPart2(t *testing.T) {
	// given
	r := strings.NewReader(sample)
	want := 43

	// when
	got, err := Process(r, true)
	// then
	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatalf("Expected %d, but got %d\n", want, got)
	}
}
