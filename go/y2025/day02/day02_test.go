package day02

import (
	"strings"
	"testing"
)

const sample = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

func TestPart1(t *testing.T) {
	// given
	r := strings.NewReader(sample)
	want := 1227775554

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
	want := 4174379265

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
