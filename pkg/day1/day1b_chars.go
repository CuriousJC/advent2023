package day1

import (
	//"fmt"
	"github.com/curiousjc/advent2023/internal/inputs"
	"strings"
)

func charAnswer1b() (total int) {
	lines, err := inputs.Input1()
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		value := extractValueByChar1b(line)
		//fmt.Print(line, " value is: ", value, "\n")
		total += value
	}
	return total
}

func extractValueByChar1b(input string) (value int) {
	foundFirst := false
	first := -1
	last := -1

	//need to extract first integer character and then last integer character
	for index, char := range input {
		if char >= '0' && char <= '9' {
			if foundFirst {
				last = int(char - '0')
			} else {
				first = int(char - '0')
				foundFirst = true
			}
		} else {
			stringFound, charValue := checkForString(index, input)
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

	value = (first * 10) + last

	return value
}

func checkForString(index int, input string) (stringFound bool, charValue rune) {
	input = strings.ToLower(input)

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

	for key, value := range stringInt {
		if len(input) >= index+len(key) {
			if input[index:index+len(key)] == key {
				return true, value
			}
		}
	}

	return false, ';'

}
