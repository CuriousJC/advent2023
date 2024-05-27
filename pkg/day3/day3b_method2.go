package day3

//Adapted from https://github.com/dbut2/advent-of-code

import (
	"fmt"
	"github.com/curiousjc/advent2023/internal/inputs"
	"github.com/curiousjc/advent2023/pkg/dbut2/sets"
	"github.com/curiousjc/advent2023/pkg/dbut2/space"
	"github.com/curiousjc/advent2023/pkg/dbut2/utils"
)

func day3bMethod2() (total int) {
	fmt.Println("running day3b method2...")
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

	grid2 := space.NewGridFromInput(s)
	for cell, char := range grid2.Cells() {
		if *char != '*' {
			continue
		}

		neighbours := sets.Set[*int]{}
		for _, neighbour := range grid.Surrounding(cell) {
			if *neighbour == nil {
				continue
			}
			neighbours.Add(*neighbour)
		}

		if len(neighbours) == 2 {
			slice := neighbours.Slice()
			total += *slice[0] * *slice[1]
		}
	}

	return total
}
