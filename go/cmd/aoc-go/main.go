package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/cpmachado/advent-of-code/go/util"
	"github.com/cpmachado/advent-of-code/go/y2024"
	"github.com/cpmachado/advent-of-code/go/y2025"
)

var cmds = map[int]util.YearCmds{
	2024: y2024.DayCommands,
	2025: y2025.DayCommands,
}

func main() {
	year, day, second := 2025, 1, false

	flag.IntVar(&day, "day", day, "Day to be run")
	flag.BoolVar(&second, "second", second, "Run part 2, instead of part 1")
	flag.Parse()

	earliestValidYear := 2015
	lastValidYear := time.Now().Year()
	if time.Now().Month() < 12 {
		lastValidYear -= 1
	}
	switch {
	case day > 25:
		slog.Error("Invalid day it should be in [1, 25]")
		flag.Usage()
		os.Exit(1)
	case year < earliestValidYear || year > lastValidYear:
		slog.Error(fmt.Sprintf("Year should be in [%d, %d]", earliestValidYear, lastValidYear))
		flag.Usage()
		os.Exit(1)
	}

	ycmd, found := cmds[year]

	if !found {
		fmt.Println("Year not Implemented")
		os.Exit(1)
	}

	cmd, found := ycmd[day]
	if !found {
		fmt.Println("Not implemented yet")
		os.Exit(0)
	}

	if err := cmd(flag.Args(), second); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
