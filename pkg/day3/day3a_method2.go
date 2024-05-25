package day3

import (
	"fmt"
	"github.com/curiousjc/advent2023/internal/inputs"
	"github.com/curiousjc/advent2023/pkg/dbut2/sets"
	"github.com/curiousjc/advent2023/pkg/dbut2/space"
	"github.com/curiousjc/advent2023/pkg/dbut2/utils"
)

func day3aMethod2() (total int) {
	fmt.Println("running day3a method2...")
	input, err := inputs.GetFullFile("static/day3.txt")
	if err != nil {
		panic(err)
	}

	s := utils.ParseInput(input)

	grid := space.Grid[*int]{}

	// parse the input and store a pointer to the buffer for each number found
	// all digits in a sequence that make up a number will all point to the same number in grid
	buffer := new(int)
	bufferCells := space.Cells{}
	for j, line := range s {
		for i, char := range line {
			if char >= '0' && char <= '9' {
				*buffer *= 10
				*buffer += int(char - '0')
				bufferCells = append(bufferCells, space.Cell{i, j})
			} else {
				for _, cell := range bufferCells {
					grid.Set(cell, buffer)
				}
				buffer = new(int)
				bufferCells = space.Cells{}
			}
		}

		for _, cell := range bufferCells {
			grid.Set(cell, buffer)
		}
		buffer = new(int)
		bufferCells = space.Cells{}
	}

	numbers := sets.Set[*int]{}
	grid2 := space.NewGridFromInput(s)
	for cell, char := range grid2.Cells() {
		if (*char == '.') || (*char >= '0' && *char <= '9') {
			continue
		}

		for _, neighbourCell := range grid.Surrounding(cell) {
			if *neighbourCell == nil {
				continue
			}
			numbers.Add(*neighbourCell)
		}
	}

	for num := range numbers {
		total += *num
	}
	return total
}
