package day1

import (
	"fmt"
	"github.com/curiousjc/advent2023/internal/inputs"
)

func charAnswer() (total int) {
	lines, err := inputs.Input1()
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		value := extractValueByChar(line)
		fmt.Print(line, " value is: ", value, "\n")
		total += value
	}
	return total
}

func extractValueByChar(input string) (value int) {
	foundFirst := false
	first := 0
	last := 0

	//need to extract first integer character and then last integer character
	for _, char := range input {
		if char >= '0' && char <= '9' {
			if foundFirst {
				last = 1
			}

		}
	}

	value = (first * 10) + last

	return value
}
