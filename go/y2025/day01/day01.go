package day01

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type RotateType int

const (
	RotateRight RotateType = iota
	RotateLeft
	RotateUnknown
)

type Rotation struct {
	Type      RotateType
	Rotations int
}

func Process(r io.Reader, part2 bool) (int, error) {
	scanner := bufio.NewScanner(r)

	dial := 50
	count := 0

	for scanner.Scan() {
		t, err := lexer(scanner.Text())
		if err != nil {
			return 0, err
		}

		rotations := t.Rotations
		q, r := rotations/100, rotations%100

		if part2 {
			count += q
		}

		switch t.Type {
		case RotateRight:
			if part2 && r > 0 && dial+r >= 100 {
				count++
			}
			dial = (dial + r) % 100

		case RotateLeft:
			if part2 && r > 0 && dial > 0 && r >= dial {
				count++
			}
			dial = (dial + 100 - r) % 100
		}

		if !part2 && dial == 0 {
			count++
		}
	}

	return count, scanner.Err()
}

func lexer(s string) (*Rotation, error) {
	rtype := RotateUnknown
	switch s[0] {
	case 'R':
		rtype = RotateRight
	case 'L':
		rtype = RotateLeft
	default:
		return nil, fmt.Errorf("unknown rotation")
	}
	n, err := strconv.Atoi(s[1:])
	if err != nil {
		return nil, err
	}

	return &Rotation{
		Type:      rtype,
		Rotations: n,
	}, nil
}
