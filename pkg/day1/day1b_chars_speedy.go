package day1

import (
	//"fmt"
	"github.com/curiousjc/advent2023/internal/inputs"
	"strings"
)

func charAnswer1bSpeedy() (total int) {
	lines, err := inputs.Input1()
	if err != nil {
		panic(err)
	}

	stringInt := make(map[string]rune)
	stringInt["zero"] = '0'
	stringInt["one"] = '1'
	stringInt["two"] = '2'
	stringInt["three"] = '3'
	stringInt["four"] = '4'
	stringInt["five"] = '5'
	stringInt["six"] = '6'
	stringInt["seven"] = '7'
	stringInt["eight"] = '8'
	stringInt["nine"] = '9'

	for _, line := range lines {
		foundFirst := false
		stringFound := false
		charValue := 'n'
		first := -1
		last := -1

		//need to extract first integer character and then last integer character
		for index, char := range line {
			if char >= '0' && char <= '9' {
				if foundFirst {
					last = int(char - '0')
				} else {
					first = int(char - '0')
					foundFirst = true
				}
			} else {

				line = strings.ToLower(line)

				for key, value := range stringInt {
					if len(line) >= index+len(key) {
						if line[index:index+len(key)] == key {
							stringFound = true
							charValue = value
						}
					}
				}

				if stringFound {
					if charValue >= '0' && charValue <= '9' {
						if foundFirst {
							last = int(charValue - '0')
						} else {
							first = int(charValue - '0')
							foundFirst = true
						}
					}
				}
			}
		}

		if last == -1 {
			last = first
		}

		total += (first * 10) + last
	}
	return total
}
