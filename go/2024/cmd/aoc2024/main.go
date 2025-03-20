package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/cpmachado/advent-of-code/2024/day01"
	"github.com/cpmachado/advent-of-code/2024/day02"
	"github.com/cpmachado/advent-of-code/2024/day03"
)

type Command func(args []string, second bool) error

var cmds = []Command{
	day01.Command,
	day02.Command,
	day03.Command,
}

func main() {
	day, second := 1, false

	flag.IntVar(&day, "day", day, "Day to be run")
	flag.BoolVar(&second, "second", second, "Run part 2, instead of part 1")
	flag.Parse()

	if day > 25 {
		slog.Error("Invalid day it should be in [1, 25]")
		flag.Usage()
		os.Exit(1)
	}

	if day > len(cmds) {
		fmt.Println("Not Implemented")
		os.Exit(0)
	}

	cmd := cmds[day-1]
	if err := cmd(flag.Args(), second); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
