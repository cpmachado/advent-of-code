package day05

import (
	"bufio"
	"io"
	"slices"
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

	// ---------------------------------------------------------------------
	// Sort and coalesce sets
	slices.SortFunc(sets, func(a Set, b Set) int {
		if a.A == b.A {
			return a.B - b.B
		}
		return a.A - b.A
	})

	oldSets := sets
	sets = nil
	var acc Set
	lidx := -1

	for i, curr := range oldSets {
		if i == 0 {
			acc = curr
		} else if curr.Contains(acc.B) {
			acc.B = curr.B
		} else {
			sets = append(sets, acc)
			acc = curr
			lidx = i
		}
	}
	if lidx+1 < len(oldSets) {
		sets = append(sets, acc)
	}

	// ---------------------------------------------------------------------

	if part2 {
		for _, s := range sets {
			count += s.B - s.A + 1
		}
	} else {
		// parse igredients
		for scanner.Scan() {
			token := scanner.Text()

			i, _ := strconv.Atoi(token)

			if isFresh(sets, i) {
				count++
			}
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
