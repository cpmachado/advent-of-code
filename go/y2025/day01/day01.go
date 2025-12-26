package day01

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
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

func Process(r io.Reader, part2 bool) (int, error) {
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
			s := dial + t.Steps
			q, r := s/100, s%100
			dial = r
			if part2 {
				count += q
			}
		case RotateLeft:
			s := dial - t.Steps
			q, r := s/100, s%100
			dial = (100 + r)
			if part2 {
				count += q
				if r < 0 {
					count++
				}
			}
		}
		dial %= 100
		if !part2 && dial == 0 {
			count++
		}
		slog.Info("check", slog.Any("t", t), slog.Int("dial", dial), slog.Int("count", count))
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
