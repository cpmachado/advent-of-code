package day05

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Set struct {
	A, B int
}

func (s *Set) Contains(k int) bool {
	return s.A <= k && k <= s.B
}

func Process(r io.Reader, part2 bool) (int, error) {
	scanner := bufio.NewScanner(r)

	count := 0
	var sets []Set

	// parse ranges
	for scanner.Scan() {
		token := scanner.Text()

		if token == "" {
			break
		}

		ns := strings.Split(token, "-")
		r := Set{}

		r.A, _ = strconv.Atoi(ns[0])
		r.B, _ = strconv.Atoi(ns[1])

		sets = append(sets, r)
	}

	// parse igredients
	for scanner.Scan() {
		token := scanner.Text()

		i, _ := strconv.Atoi(token)

		if isFresh(sets, i) {
			count++
		}
	}

	return count, nil
}

func isFresh(sets []Set, i int) bool {
	for _, s := range sets {
		if s.Contains(i) {
			return true
		}
	}
	return false
}
