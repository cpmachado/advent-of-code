package day03

import (
	"bufio"
	"io"
)

func Process(r io.Reader, part2 bool) (int, error) {
	scanner := bufio.NewScanner(r)

	sum := 0

	var paperMap []string

	for scanner.Scan() {
		token := scanner.Text()
		paperMap = append(paperMap, token)
	}

	return sum, nil
}
