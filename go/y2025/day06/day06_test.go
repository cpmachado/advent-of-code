package day06

import (
	"strings"
	"testing"
)

const sample = `123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   +
`

func TestPart1(t *testing.T) {
	// given
	r := strings.NewReader(sample)
	want := 4277556

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
	want := 3263827

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
