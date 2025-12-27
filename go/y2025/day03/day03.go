package day03

import (
	"bufio"
	"io"
)

func Process(r io.Reader, part2 bool) (int, error) {
	scanner := bufio.NewScanner(r)

	sum := 0

	k := 2

	if part2 {
		k = 12
	}

	for scanner.Scan() {
		token := scanner.Text()
		mj := retrieveMaxJoltage(token, k)
		sum += mj
	}

	return sum, nil
}

func retrieveMaxJoltage(s string, k int) int {
	n := len(s)
	b := 0
	sum := 0

	for k > 0 {
		f := n - k - b + 1
		t := '0'
		ti := 0
		for i, r := range s[b : b+f] {
			if r > t {
				ti = i
				t = r
			}
		}
		sum = sum*10 + int(t-'0')
		b += ti
		b++
		k--
	}

	return sum
}
