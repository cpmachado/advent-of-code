package y2025

import (
	"github.com/cpmachado/advent-of-code/go/util"
	"github.com/cpmachado/advent-of-code/go/y2025/day01"
	"github.com/cpmachado/advent-of-code/go/y2025/day02"
)

var DayCommands = map[int]util.Command{
	1: day01.Command,
	2: day02.Command,
}
