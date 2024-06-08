package day3

import (
	"fmt"
	"github.com/curiousjc/advent2023/internal/inputs"
	"slices"
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

		//compose a list of every column that "matters" for the previous line, the processed line, and the subsequent line
		var capturedCols []int
		prevCols := getSymbolColumns(prevLine)
		capturedCols = append(capturedCols, prevCols...)
		procCols := getSymbolColumns(line)
		capturedCols = append(capturedCols, procCols...)
		nextCols := getSymbolColumns(nextLine)
		capturedCols = append(capturedCols, nextCols...)
		distinctCols := expandSymbolColumns(capturedCols)

		//establish our processing line separately because may be changing if there's more than one part number
		procLine := line

		//sequence will matter
		slices.Sort(distinctCols)

		//process only the columns that mattered, sequentially, looking for a digit that signals a part number
		for _, validColumn := range distinctCols {
			char := rune(procLine[validColumn])

			//when we have found a digit on our column we have found a piece of a part number
			if unicode.IsDigit(char) {
				digits, updatedLine := extractDigitsFromIndex(procLine, validColumn)
				total += digits
				procLine = updatedLine
			}
		}
	}

	return total
}

func getSymbolColumns(input string) []int {
	var symbolCols []int

	//anything that isn't a digit or a period should be captured
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

	for num := range uniqueMap {
		distinct = append(distinct, num)
	}

	return distinct
}
