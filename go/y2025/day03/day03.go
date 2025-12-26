package day03

import (
	"bufio"
	"io"
)

func Process(r io.Reader, part2 bool) (int, error) {
	scanner := bufio.NewScanner(r)

	sum := 0

	for scanner.Scan() {
	}

	return sum, nil
}
