package day03

import (
	"bufio"
	"io"
)

func Process(r io.Reader, part2 bool) (int, error) {
	scanner := bufio.NewScanner(r)

	sum := 0

	for scanner.Scan() {
		token := scanner.Text()
		mj := retrieveMaxJoltage(token)
		sum += mj
	}

	return sum, nil
}

func retrieveMaxJoltage(s string) int {
	mdidx := 0
	ba := '0'
	n := len(s)

	for i, r := range s {
		if r > ba && i+1 < n {
			mdidx = i
			ba = r
		}
	}

	bb := '0'

	for i, r := range s[mdidx+1:] {
		if r > bb && i+1 < n {
			bb = r
		}
	}

	a := int(ba - '0')
	b := int(bb - '0')
	return a*10 + b
}
