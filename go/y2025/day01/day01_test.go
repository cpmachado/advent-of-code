package day01

import (
	"strings"
	"testing"
)

const sample = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`

func TestPart1(t *testing.T) {
	// given
	r := strings.NewReader(sample)

	// when
	count, err := Process(r)
	// then
	if err != nil {
		t.Fatal(err)
	}

	if count != 3 {
		t.Fatalf("Expected 3, but got %d\n", count)
	}
}
