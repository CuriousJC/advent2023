package day3

import (
	"fmt"
	"github.com/curiousjc/advent2023/internal/inputs"
	"strings"
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
	lines, err := inputs.Get("static/day3_validate.txt")
	if err != nil {
		panic(err)
	}

	for index, line := range lines {
		prevLine := ""
		procLine := ""
		nextLine := ""

		//handle processing first and last lines
		if index == 0 {
			prevLine = strings.Repeat(".", len(line))
		}
		if index == len(lines)-1 {
			nextLine = strings.Repeat(".", len(line))
		}

		prevCols := getSymbolColumns(prevLine)
		procCols := getSymbolColumns(procLine)
		nextCols := getSymbolColumns(nextLine)

		//TODO: CaptureCols = combine all the columns into a distinct list of columns that matter and also capture the column before and after each meaningful column

		//TODO: for each byte of the procLine
		//if its column is a CaptureCol then grow forwards and backwards until meeting a not-integer, then capture that number, then wipe number and update procLine and continue

	}

	for _, line := range lines {
		fmt.Println(line)
	}
	return total
}

func getSymbolColumns(input string) []int {

	//TODO: Take a string and identify each column that has a symble, return back those column numbers in the slice
}
