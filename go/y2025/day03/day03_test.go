package day03

import (
	"strings"
	"testing"
)

const sample = `987654321111111
811111111111119
234234234234278
818181911112111
`

func TestPart1(t *testing.T) {
	// given
	r := strings.NewReader(sample)
	want := 357

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
