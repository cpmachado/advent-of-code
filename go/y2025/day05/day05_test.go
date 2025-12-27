package day05

import (
	"strings"
	"testing"
)

const sample = `3-5
10-14
16-20
12-18

1
5
8
11
17
32
`

func TestPart1(t *testing.T) {
	// given
	r := strings.NewReader(sample)
	want := 3

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
	want := 14

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
