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
	Type  RotateType
	Steps int
}

func Process(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	dial := 50
	count := 0

	for scanner.Scan() {
		t, err := lexer(scanner.Text())
		if err != nil {
			return 0, nil
		}

		switch t.Type {
		case RotateRight:
			dial = (dial + t.Steps) % 100
		case RotateLeft:
			dial = (100 + (dial-t.Steps)%100) % 100
		}
		if dial == 0 {
			count++
		}
	}
	return count, nil
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
		Type:  rtype,
		Steps: n,
	}, nil
}
