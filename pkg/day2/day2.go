package day2

import (
	"fmt"
	"time"
)

func Day2() {
	fmt.Println("------")
	fmt.Println("Day2a Answer is supposed to be 2679...")

	start := time.Now()
	fmt.Println("method1 answer: ", day2aMethod1())
	fmt.Println("method execution took", time.Since(start))

	fmt.Println("------")
	fmt.Println("Day2b Answer is supposed to be 77607...")

	start = time.Now()
	fmt.Println("method1 answer: ", day2bMethod1())
	fmt.Println("method execution took", time.Since(start))

}
