package day1

//Adapted from https://github.com/dbut2/advent-of-code

import (
	//"fmt"
	"github.com/curiousjc/advent2023/internal/inputs"
)

func UtilSolve() int {
	s, err := inputs.Input1()
	if err != nil {
		panic(err)
	}

	words := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	total := 0
	for _, line := range s {
		firstDigit := -1
		lastDigit := 0

		for i, char := range line {
			var dig int

			found := false
			if IsNum(char) {
				dig = NumVal(char)
				found = true
			}
			if !found {
				dig, found = checkWords(line, i, words)
			}
			if !found {
				continue
			}

			if firstDigit == -1 {
				firstDigit = dig
			}
			lastDigit = dig
		}

		value := (firstDigit * 10) + lastDigit
		//fmt.Print(line, " value is: ", value, "\n")
		total += value

	}

	return total
}

func checkWords(line string, i int, words []string) (int, bool) {
	for j, word := range words {
		if checkWord(line, i, word) {
			return j, true
		}
	}
	return 0, false
}

func checkWord(line string, i int, word string) bool {
	if i+len(word) > len(line) {
		return false
	}
	return line[i:i+len(word)] == word
}

func IsNum(c rune) bool {
	return c >= '0' && c <= '9'
}

func IsLower(c rune) bool {
	return c >= 'a' && c <= 'z'
}

func IsCapital(c rune) bool {
	return c >= 'A' && c <= 'Z'
}

func IsLetter(c rune) bool {
	return IsLower(c) || IsCapital(c)
}

func NumVal(c rune) int {
	return int(c - '0')
}
