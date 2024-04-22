package ex1

import (
	"fmt"
	"github.com/curiousjc/advent2023/internal/inputs"
)

//Define the lines I'm working with
//Break out and process each line asynchronously
// --> need to handle single numbers as doubled up
// --> need to identify digits only, strip off the strings?
//Add all the answers together to form a single number
//plugin answer here: https://adventofcode.com/2023/day/1

func Answer() {

	fmt.Println("Example 1 Answer is...")

	lines, err := inputs.Input1()
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		fmt.Println(line)
	}

}
