package day1

import (
	//"fmt"
	"github.com/curiousjc/advent2023/internal/inputs"
	"strings"
)

func charAnswer1bSpeedy() (total int) {
	lines, err := inputs.Get("static/day1.txt")
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
			stringFound = false
			if char >= '0' && char <= '9' {
				if foundFirst {
					//fmt.Println("Setting Last: ", string(char))
					last = int(char - '0')
				} else {
					first = int(char - '0')
					//fmt.Println("Setting First: ", string(char))
					foundFirst = true
				}
			} else {

				line = strings.ToLower(line)

				for key, value := range stringInt {
					if len(line) >= index+len(key) {
						if line[index:index+len(key)] == key {
							//fmt.Println("Setting stringFound...")
							stringFound = true
							charValue = value
						}
					}
				}

				if stringFound {
					if charValue >= '0' && charValue <= '9' {
						if foundFirst {
							last = int(charValue - '0')
							//fmt.Println("Setting Last String: ", string(charValue))
						} else {
							first = int(charValue - '0')
							//fmt.Println("Setting First String: ", string(charValue))
							foundFirst = true
						}
					}
				}
			}
		}

		if last == -1 {
			last = first
		}

		value := (first * 10) + last
		//fmt.Print(line, " value is: ", value, "\n")
		total += value
	}
	return total
}
