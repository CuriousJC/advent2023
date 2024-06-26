package day1

import (
	"github.com/curiousjc/advent2023/internal/inputs"
)

func charAnswer() (total int) {
	lines, err := inputs.Get("static/day1.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		value := extractValueByChar(line)
		//fmt.Print(line, " value is: ", value, "\n")
		total += value
	}
	return total
}

func extractValueByChar(input string) (value int) {
	foundFirst := false
	first := -1
	last := -1

	//need to extract first integer character and then last integer character
	for _, char := range input {
		if char >= '0' && char <= '9' {
			if foundFirst {
				last = int(char - '0')
			} else {
				first = int(char - '0')
				foundFirst = true
			}
		}
	}

	if last == -1 {
		last = first
	}

	value = (first * 10) + last

	return value
}
