/* Copyright © 2024 cpmachado <cpmachado@protonmail> */
package day02

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// Parses the input file for day02
func ParseFile(filename string) []Report {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	result := []Report{}
	for scanner.Scan() {
		line := scanner.Text()
		cols := strings.Fields(line)
		report := Report{}

		for _, col := range cols {
			level, _ := strconv.Atoi(col)
			report = append(report, Level(level))
		}

		result = append(result, report)
	}
	return result
}
