package y2024

import (
	"github.com/cpmachado/advent-of-code/go/util"
	"github.com/cpmachado/advent-of-code/go/y2024/day01"
	"github.com/cpmachado/advent-of-code/go/y2024/day02"
	"github.com/cpmachado/advent-of-code/go/y2024/day03"
)

var DayCommands = map[int]util.Command{
	1: day01.Command,
	2: day02.Command,
	3: day03.Command,
}
