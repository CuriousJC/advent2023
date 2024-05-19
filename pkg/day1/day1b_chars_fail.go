package day1

//This fails!  It fails because it will process something like 'plckgsixeight6eight95bnrfonetwonej' and return back 62 instead of 61
// or for 'hrh34sixfourqqng7eightwot' it will return back 38 instead of 32 - because it flips that "eight" to an 8 and then there isn't a "two" afterwards anymore
// leaving to analyze compared to a different approach

import (
	//"fmt"
	"github.com/curiousjc/advent2023/internal/inputs"
	"strings"
)

func charAnswerb() (total int) {
	lines, err := inputs.Get("static/day1.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		value := extractValueByCharb(line)
		//fmt.Print(line, " value is: ", value, "\n")
		total += value
	}
	return total
}

func extractValueByCharb(input string) (value int) {
	foundFirst := false
	first := -1
	last := -1

	//FAILURE: flip most words to integers - but it's processing sequentially so if it flips a word that ends with a character that makes a subsequent word then it messes it up
	input = wordsToInts(input)

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

func wordsToInts(input string) string {
	input = strings.ToLower(input)

	stringInt := make(map[string]string)
	stringInt["zero"] = "0"
	stringInt["one"] = "1"
	stringInt["two"] = "2"
	stringInt["three"] = "3"
	stringInt["four"] = "4"
	stringInt["five"] = "5"
	stringInt["six"] = "6"
	stringInt["seven"] = "7"
	stringInt["eight"] = "8"
	stringInt["nine"] = "9"

	//fmt.Println("The string before processing is: ", input)

	for i, l := 0, len(input); i <= l; i++ {
		//fmt.Println("processing index", i, "input is: ", input)
		for key, value := range stringInt {
			//fmt.Println("searching...", key, "input is...", input[i:])
			if len(input) >= i+len(key) {
				if input[i:i+len(key)] == key {
					//fmt.Println("Or panic here?")
					input = strings.Replace(input, key, string(value), 1)
					i = 0
					l = len(input)
				}
			}
		}
	}

	//fmt.Println("The string after processing is: ", input)
	return input
}
