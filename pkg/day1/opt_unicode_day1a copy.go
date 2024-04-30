package day1

import (
	"github.com/curiousjc/advent2023/internal/inputs"
	"unicode"
)

func unicodeAnswer() (total int) {
	lines, err := inputs.Input1()
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		value := extractByUnicode(line)
		total += value
	}
	return total
}

func extractByUnicode(input string) (value int) {
	foundFirst := false
	first := -1
	last := -1

	//need to extract first integer character and then last integer character
	for _, char := range input {
		if unicode.IsDigit(char) {
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
