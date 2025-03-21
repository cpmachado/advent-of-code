package day03

import "testing"

func TestProgramInterpreter(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	got := ProgramInterpreter(input)
	want := 161

	if got != want {
		t.Fatalf("Wanted '%d' and got '%d'\n", want, got)
	}
}

func TestProgramInterpreterImproved(t *testing.T) {
	input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	got := ProgramInterpreterImproved(input)
	want := 48

	if got != want {
		t.Fatalf("Wanted '%d' and got '%d'\n", want, got)
	}
}
