package day1

import (
	//"fmt"
	"github.com/curiousjc/advent2023/internal/inputs"
	"strings"
)

func charAnswer1bSpeedySpeedy() (total int) {
	lines, err := inputs.Get("static/day1.txt")
	if err != nil {
		panic(err)
	}

	words := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

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
					last = int(char - '0')
				} else {
					first = int(char - '0')
					foundFirst = true
				}
			} else {

				line = strings.ToLower(line)

				for key, value := range words {
					if len(line) >= index+len(value) {
						//fmt.Println(line[index : index+len(value)])
						if line[index:index+len(value)] == value {
							//fmt.Println("String found...")
							stringFound = true
							charValue = rune(key)
						}
					}
				}

				if stringFound {
					//fmt.Println("STring found value is", charValue)
					if charValue >= 0 && charValue <= 9 {
						if foundFirst {
							//fmt.Println("Setting last string...")
							last = int(charValue - 0)
						} else {
							first = int(charValue - 0)
							//fmt.Println("Setting first string...")
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
