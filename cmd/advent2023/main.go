package main

import (
	"fmt"
	"github.com/curiousjc/advent2023/internal/learning"
	"github.com/curiousjc/advent2023/pkg/day1"
	"github.com/curiousjc/advent2023/pkg/day2"
	"github.com/curiousjc/advent2023/pkg/day3"
	"github.com/curiousjc/advent2023/pkg/day4"
)

func main() {
	var ignoreMe bool = false

	fmt.Println("-----------------")

	if ignoreMe {
		learning.Print()
		fmt.Println("-----------------")
		day1.Day1()
		fmt.Println("-----------------")
		day2.Day2()
		fmt.Println("-----------------")
		day3.Day3()
		fmt.Println("-----------------")
	} else {
		day4.Day4()
	}

}
