package day03

import (
	"regexp"
	"strconv"
	"strings"
)

// Processes the multiplications from corrupted input program
func ProgramInterpreter(input string) int {
	sum := 0
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	cleanRe := regexp.MustCompile("[mul()]")
	for _, v := range re.FindAllString(input, -1) {
		cv := strings.Split(cleanRe.ReplaceAllString(v, ""), ",")
		a, _ := strconv.Atoi(cv[0])
		b, _ := strconv.Atoi(cv[1])
		sum += a * b
	}
	return sum
}
