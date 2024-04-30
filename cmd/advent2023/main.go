package main

import (
	"fmt"
	"github.com/curiousjc/advent2023/internal/learning"
	"github.com/curiousjc/advent2023/pkg/day1"
)

// todo - commandline to print only some answers?
func main() {
	learning.Print()
	fmt.Println("-----------------")
	day1.Answer1a()
	day1.Answer1b()
}
