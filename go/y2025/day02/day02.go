package day02

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Process(r io.Reader, part2 bool) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF {
			return len(data), data, io.EOF
		}

		for i, r := range data {
			if r == ',' {
				return i + 1, data[:i], nil
			}
		}
		return len(data), data, nil
	})

	sum := 0

	for scanner.Scan() {
		s := strings.Split(strings.TrimSpace(scanner.Text()), "-")
		a, err := strconv.Atoi(s[0])
		if err != nil {
			return 0, err
		}
		b, err := strconv.Atoi(s[1])
		if err != nil {
			return 0, err
		}
		for a <= b {

			if isInvalid(a) {
				sum += a
			}
			a++
		}
	}

	return sum, nil
}

func isInvalid(a int) bool {
	sa := strconv.Itoa(a)
	k := len(sa)
	mid := k / 2

	return k%2 == 0 && sa[:mid] == sa[mid:]
}
