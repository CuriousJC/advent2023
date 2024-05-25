package day3

import (
	"fmt"
	"time"
)

func Day3() {
	fmt.Println("Exercise 3a Answer is supposed to be 556057...")

	start := time.Now()
	fmt.Println("Day3a method 1 answer is: ", day3aMethod1())
	fmt.Println("method execution took", time.Since(start))

	start = time.Now()
	fmt.Println("Day3a method 2 answer is: ", day3aMethod2())
	fmt.Println("method execution took", time.Since(start))

	fmt.Println("Exercise 3b Answer is supposed to be 82824352...")
	start = time.Now()
	fmt.Println("Day3b method 1 answer is: ", day3bMethod1())
	fmt.Println("method execution took", time.Since(start))

	start = time.Now()
	fmt.Println("Day3b method 2 answer is: ", day3bMethod2())
	fmt.Println("method execution took", time.Since(start))
}
