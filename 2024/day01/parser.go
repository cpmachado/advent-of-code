package day01

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ParseFile(filename string) ([]LocationId, []LocationId) {
	listA := []LocationId{}
	listB := []LocationId{}

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
		listA = append(listA, LocationId(a))
		listB = append(listB, LocationId(b))

	}
	return listA, listB
}
