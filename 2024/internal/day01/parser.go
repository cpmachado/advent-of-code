package day01

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// ParseFile parses puzzle input
func ParseFile(filename string) (LocationIDList, LocationIDList) {
	listA := LocationIDList{}
	listB := LocationIDList{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		cols := strings.Fields(line)
		a, _ := strconv.Atoi(cols[0])
		b, _ := strconv.Atoi(cols[1])
		listA = append(listA, a)
		listB = append(listB, b)
	}
	return listA, listB
}
