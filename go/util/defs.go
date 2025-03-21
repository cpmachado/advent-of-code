package util

type Command func(args []string, second bool) error

type YearCmds map[int]Command
