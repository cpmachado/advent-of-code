package day03

import (
	"bufio"
	"io"
)

func Process(r io.Reader, part2 bool) (int, error) {
	scanner := bufio.NewScanner(r)

	count := 0

	var papermap []string

	for scanner.Scan() {
		token := scanner.Text()
		papermap = append(papermap, token)
	}

	for i, row := range papermap {
		for j, r := range row {
			if r == '@' {
				if k := countAdjacentRolls(i, j, papermap); k < 4 {
					count++
				}
			}
		}
	}

	return count, nil
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
