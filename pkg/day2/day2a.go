package day2

import (
	"fmt"
	"time"
)

func Day2() {
	fmt.Println("Day2a Answer is supposed to be 2679...")

	start := time.Now()
	fmt.Println("method1 answer: ", method1())
	fmt.Println("method execution took", time.Since(start))

}
