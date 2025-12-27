package y2025

import (
	"github.com/cpmachado/advent-of-code/go/util"
	"github.com/cpmachado/advent-of-code/go/y2025/day01"
	"github.com/cpmachado/advent-of-code/go/y2025/day02"
	"github.com/cpmachado/advent-of-code/go/y2025/day03"
	"github.com/cpmachado/advent-of-code/go/y2025/day04"
)

var DayCommands = map[int]util.Command{
	1: day01.Command,
	2: day02.Command,
	3: day03.Command,
	4: day04.Command,
}
