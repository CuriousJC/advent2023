package day3

import (
	"fmt"
	"github.com/curiousjc/advent2023/internal/inputs"
	"strconv"
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

		var capturedCols []int

		prevCols := getSymbolColumns(prevLine)
		capturedCols = append(capturedCols, prevCols...)

		procCols := getSymbolColumns(procLine)
		capturedCols = append(capturedCols, procCols...)

		nextCols := getSymbolColumns(nextLine)
		capturedCols = append(capturedCols, nextCols...)

		distinctCols := expandSymbolColumns(capturedCols)

		for _, col := range distinctCols {
			fmt.Println(strconv.Itoa(col))
		}

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

	numbers := []int{1, 2, 3, 4, 5}

	return numbers

}

func expandSymbolColumns(input []int) []int {
	// Create a map to store unique elements
	uniqueMap := make(map[int]bool)

	// Iterate over the input slice and add each element to the map, growing left and right by one column to capture diagonals
	for _, num := range input {
		uniqueMap[num-1] = true
		uniqueMap[num] = true
		uniqueMap[num+1] = true
	}

	fmt.Println("map length is", len(uniqueMap))

	// Create a new slice to store distinct elements
	distinct := make([]int, 0, len(uniqueMap))

	// Iterate over the map keys and add them to the distinct slice
	for num := range uniqueMap {
		distinct = append(distinct, num)
	}

	return distinct
}
