// Package day03 contains code relating to AOC day03 of 2024
package day03

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	mulRegex  = `mul\(\d{1,3},\d{1,3}\)`
	doRegex   = `do\(\)`
	dontRegex = `don't\(\)`
)

// ProgramInterpreter processes a corrupted input program
func ProgramInterpreter(input string) int {
	sum := 0
	re := regexp.MustCompile(mulRegex)
	cleanRe := regexp.MustCompile("[mul()]")
	for _, v := range re.FindAllString(input, -1) {
		cv := strings.Split(cleanRe.ReplaceAllString(v, ""), ",")
		a, _ := strconv.Atoi(cv[0])
		b, _ := strconv.Atoi(cv[1])
		sum += a * b
	}
	return sum
}

// ProgramInterpreterImproved processes thoroughly a corrupted input program
func ProgramInterpreterImproved(input string) int {
	re := regexp.MustCompile(dontRegex)
	dore := regexp.MustCompile(doRegex)
	sum := 0

	for i, v := range re.Split(input, -1) {
		if i == 0 {
			sum += ProgramInterpreter(v)
		} else {
			for _, k := range dore.Split(v, -1)[1:] {
				sum += ProgramInterpreter(k)
			}
		}
	}
	return sum
}
