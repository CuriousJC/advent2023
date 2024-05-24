package day3

import (
	"fmt"
	"github.com/curiousjc/advent2023/internal/inputs"
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
func day3bMethod1() (total int) {
	fmt.Println("running day3b method1...")
	lines, err := inputs.Get("static/day3.txt")
	if err != nil {
		panic(err)
	}

	for index, line := range lines {

		//don't care about lines unless they have a gear
		if strings.Contains(line, "*") {
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

			//find any columns with a gear in our line
			gearCols := getGearColumns(line)

			//process each found gear separately, looking for exactly 2 part numbers total that are adjacent above, on, or below the column that has the gear
			for _, col := range gearCols {
				var partNums []int
				prevPartNums := extractPartNums(prevLine, col)
				partNums = append(partNums, prevPartNums...)
				linePartNums := extractPartNums(line, col)
				partNums = append(partNums, linePartNums...)
				nextPartNums := extractPartNums(nextLine, col)
				partNums = append(partNums, nextPartNums...)

				if len(partNums) == 2 {
					gearRatio := partNums[0] * partNums[1]
					total += gearRatio
				}

			}
		}
	}

	return total
}

func getGearColumns(input string) []int {
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

// for a given line and a meaningful column - pull any part numbers that adjacent to that column
func extractPartNums(input string, column int) (partNums []int) {

	//we need to process the column before our column and the column after, so start one behind
	i := column
	if column > 0 {
		i = column - 1
	}

	procLine := input
	for i <= column+1 {
		char := rune(procLine[i])
		if unicode.IsDigit(char) {
			partNum, updatedLine := extractDigitsFromIndex(procLine, i)
			procLine = updatedLine
			partNums = append(partNums, partNum)
		}
		i++
	}

	return partNums
}
