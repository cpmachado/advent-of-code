package day06

import (
	"bufio"
	"fmt"
	"io"
)

func Process(r io.Reader, part2 bool) (int, error) {
	scanner := bufio.NewScanner(r)

	sum := 0

	// parse ranges
	for scanner.Scan() {
		token := scanner.Text()
		fmt.Println(token)
	}

	return sum, nil
}
