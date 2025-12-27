package day07

import (
	"strings"
	"testing"
)

const sample = `.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............
`

func TestProcess(t *testing.T) {
	tests := []struct {
		description string
		want        int
		part2       bool
	}{
		{description: "part1", want: 21},
		{description: "part2", want: 40, part2: true},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			// given
			want := test.want
			part2 := test.part2
			r := strings.NewReader(sample)

			// when
			got, err := Process(r, part2)
			// then
			if err != nil {
				t.Fatal(err)
			}

			if got != want {
				t.Fatalf("Expected %d, but got %d\n", want, got)
			}
		})
	}
}
