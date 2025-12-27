package day06

import (
	"bufio"
	"io"
	"slices"
)

func Process(r io.Reader, part2 bool) (int, error) {
	scanner := bufio.NewScanner(r)

	count := 0

	var rays []int

	// parse ranges
	for scanner.Scan() {
		token := scanner.Text()
		for i, r := range token {
			switch r {
			case 'S':
				rays = append(rays, i)
			case '^':
				if idx := slices.Index(rays, i); idx >= 0 {
					rays = slices.Delete(rays, idx, idx+1)
					if !slices.Contains(rays, i-1) {
						rays = append(rays, i-1)
					}
					if !slices.Contains(rays, i+1) {
						rays = append(rays, i+1)
					}
					count++
				}
			}
		}

	}

	return count, nil
}
