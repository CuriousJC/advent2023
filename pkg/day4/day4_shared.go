package day4

import (
	"strconv"
	"strings"
	"unicode"
)

// given the string and a "hit" index, backup to the start of the digits and move forwards to the last digit; that's our part number - grab it and replace with periods
func extractDigitsFromIndex(input string, startIndex int) (fullNumber int, cleansedInput string) {
	var digits string
	var firstDigit int
	var lastDigit int

	//move backwards until not a digit
	i := startIndex
	for i >= 0 && unicode.IsDigit(rune(input[i])) {
		digits = string(input[i]) + digits
		i--
	}
	firstDigit = i + 1

	//continue forwards until not a digit
	i = startIndex + 1
	for i < len(input) && unicode.IsDigit(rune(input[i])) {
		digits = digits + string(input[i])
		i++
	}
	lastDigit = i

	//use everything up to the first digit, flip the whole part number to ".", use everything after the last digit
	cleansedInput = input[:firstDigit] + strings.Repeat(".", len(digits)) + input[lastDigit:]

	// Convert extracted digits to integer
	fullNumber, _ = strconv.Atoi(digits)

	return fullNumber, cleansedInput
}
