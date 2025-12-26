package day02

import (
	"bufio"
	"io"
)

func Process(r io.Reader, part2 bool) (int, error) {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
	}

	return 0, scanner.Err()
}
