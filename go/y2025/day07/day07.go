package day07

import (
	"bufio"
	"io"
)

func Process(r io.Reader, part2 bool) (int, error) {
	scanner := bufio.NewScanner(r)

	count := 0

	var rays []int

	// parse ranges
	for scanner.Scan() {
		token := scanner.Text()
		if rays == nil {
			rays = make([]int, len(token))
		}
		for i, r := range token {
			switch r {
			case 'S':
				rays[i] = 1
			case '^':
				if rays[i] > 0 {
					if !part2 {
						rays[i-1] = 1
						rays[i+1] = 1
						rays[i] = 0
						count++
					} else {
						rays[i-1] += rays[i]
						rays[i+1] += rays[i]
						rays[i] = 0
					}
				}
			}
		}

	}

	if part2 {
		for _, ray := range rays {
			count += ray
		}
	}

	return count, nil
}
