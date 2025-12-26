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

	validator := isInvalid
	if part2 {
		validator = isInvalidPart2
	}

	for scanner.Scan() {
		token := scanner.Text()
		s := strings.Split(strings.TrimSpace(token), "-")
		a, err := strconv.Atoi(s[0])
		if err != nil {
			return 0, err
		}
		b, err := strconv.Atoi(s[1])
		if err != nil {
			return 0, err
		}
		for a <= b {
			if validator(a) {
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

func isInvalidPart2(a int) bool {
	sa := strconv.Itoa(a)
	n := len(sa)

	for _, f := range factorsCanonic(n) {
		found := true
		prev := sa[:f]
		for i := f; i < n; i += f {
			curr := sa[i : i+f]
			if prev != curr {
				found = false
				break
			}
			prev = curr
		}
		if found {
			return true
		}
	}

	return false
}

func factorsCanonic(k int) []int {
	var fset []int

	for i := 1; i < k; i++ {
		if k%i == 0 {
			fset = append(fset, i)
		}
	}

	return fset
}
