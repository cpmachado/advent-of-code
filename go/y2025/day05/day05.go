package day05

import (
	"bufio"
	"fmt"
	"io"
)

func Process(r io.Reader, part2 bool) (int, error) {
	scanner := bufio.NewScanner(r)

	count := 0

	var papermap []string

	for scanner.Scan() {
		token := scanner.Text()
		papermap = append(papermap, token)
		fmt.Printf("%q\n", token)
	}

	return count, nil
}
