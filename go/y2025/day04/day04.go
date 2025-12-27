package day04

import (
	"bufio"
	"io"
	"strings"
)

func Process(r io.Reader, part2 bool) (int, error) {
	scanner := bufio.NewScanner(r)

	count := 0

	var papermap []string

	for scanner.Scan() {
		token := scanner.Text()
		papermap = append(papermap, token)
	}

	count, papermap = countAndClean(papermap)

	if part2 {
		k := count
		for k > 0 {
			k, papermap = countAndClean(papermap)
			count += k
		}
	}

	return count, nil
}

func countAndClean(papermap []string) (int, []string) {
	count := 0
	var clean []string

	for i, row := range papermap {
		nrow := &strings.Builder{}
		for j, r := range row {
			if r == '@' {
				if k := countAdjacentRolls(i, j, papermap); k < 4 {
					count++
					nrow.WriteRune('.')
					continue
				}
			}
			nrow.WriteRune(r)
		}

		clean = append(clean, nrow.String())
	}

	return count, clean
}

func countAdjacentRolls(x, y int, papermap []string) int {
	count := 0
	for i := -1; i < 2; i++ {
		if x+i < 0 || x+i >= len(papermap) {
			continue
		}
		row := papermap[x+i]
		ni := len(row)
		for j := -1; j < 2; j++ {
			if y+j < 0 || y+j >= ni || i == 0 && j == 0 {
				continue
			}
			if row[y+j] == '@' {
				count++
			}
		}
	}
	return count
}
