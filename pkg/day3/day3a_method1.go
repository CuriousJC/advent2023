package day3

import (
	"fmt"
	"github.com/curiousjc/advent2023/internal/inputs"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

/*
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
*/
func day3aMethod1() (total int) {
	fmt.Println("running day3a method1...")
	lines, err := inputs.Get("static/day3.txt")
	if err != nil {
		panic(err)
	}

	for index, line := range lines {
		prevLine := ""
		nextLine := ""

		//handle processing first and last lines by creating a fully "...." line
		if index == 0 {
			prevLine = strings.Repeat(".", len(line))
		} else {
			prevLine = lines[index-1]
		}

		if index == len(lines)-1 {
			nextLine = strings.Repeat(".", len(line))
		} else {
			nextLine = lines[index+1]
		}

		//compose a list of every column that "matters"
		var capturedCols []int
		prevCols := getSymbolColumns(prevLine)
		capturedCols = append(capturedCols, prevCols...)
		procCols := getSymbolColumns(line)
		capturedCols = append(capturedCols, procCols...)
		nextCols := getSymbolColumns(nextLine)
		capturedCols = append(capturedCols, nextCols...)
		distinctCols := expandSymbolColumns(capturedCols)

		//establish our processing line separately because it will be changing
		procLine := line

		//sequence will matter
		slices.Sort(distinctCols)

		//process only the columns that mattered, sequentially, looking for a digit that signals a part number
		for _, validColumn := range distinctCols {
			char := rune(procLine[validColumn])

			//when we have found a digit on our column we have found a piece of a part
			if unicode.IsDigit(char) {
				digits, updatedLine := extractDigitsFromIndex(procLine, validColumn)
				total += digits
				procLine = updatedLine
			}
		}
	}

	return total
}

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

func getSymbolColumns(input string) []int {
	var symbolCols []int

	//anything that isn't a digit or a period should be captured, and we need to increment index to change zero-index to one-index columns
	for index, char := range input {
		if !unicode.IsDigit(char) {
			if char != '.' {
				symbolCols = append(symbolCols, index)
			}
		}
	}

	return symbolCols
}

func expandSymbolColumns(input []int) []int {
	//key the map on column number
	uniqueMap := make(map[int]bool)

	//with diagonals we need to include the preceding and succeeding column
	for _, num := range input {
		uniqueMap[num-1] = true
		uniqueMap[num] = true
		uniqueMap[num+1] = true
	}

	distinct := make([]int, 0, len(uniqueMap))

	//TODO: It's possible this would be faster if we just leave it as a map?
	for num := range uniqueMap {
		distinct = append(distinct, num)
	}

	return distinct
}
