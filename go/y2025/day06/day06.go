package day06

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Mode int

const (
	ModeNumbers Mode = iota
	ModeOps
)

func Process(r io.Reader, part2 bool) (int, error) {
	scanner := bufio.NewScanner(r)

	sum := 0

	var numbers [][]int
	var ops []string

	mode := ModeNumbers

	// parse ranges
	for scanner.Scan() {
		tokenRow := scanner.Text()
		var nrow []int
		tokens := strings.SplitSeq(tokenRow, " ")

		for token := range tokens {
			token = strings.TrimSpace(token)
			if token == "" {
				continue
			}
			if mode == ModeOps {
				ops = append(ops, token)
			} else {
				n, err := strconv.Atoi(token)
				if err != nil {
					mode = ModeOps
					ops = append(ops, token)
					continue
				}
				nrow = append(nrow, n)
			}
		}

		if mode == ModeNumbers {
			numbers = append(numbers, nrow)
		}
	}

	if len(numbers) == 0 {
		return 0, fmt.Errorf("empty numbers")
	}

	var oper func(int, int) int

	for i := range len(numbers[0]) {
		acc := 0
		switch ops[i] {
		case "+":
			oper = add
		case "*":
			oper = mul
			acc = 1
		}
		for _, row := range numbers {
			acc = oper(acc, row[i])
		}
		sum += acc
	}

	return sum, nil
}

func add(a int, b int) int {
	return a + b
}

func mul(a int, b int) int {
	return a * b
}
