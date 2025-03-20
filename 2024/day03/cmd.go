package day03

import (
	"fmt"
	"log"
	"os"
)

func Command(args []string, second bool) {
	filename := args[0]
	input, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	str := string(input)

	if !Second {
		fmt.Printf("Program resulted in: %d\n", ProgramInterpreter(str))
	} else {
		fmt.Printf("Program resulted in: %d\n", ProgramInterpreterImproved(str))
	}
}
