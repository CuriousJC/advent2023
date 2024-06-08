package day4

import (
	"fmt"
	"time"
)

func Day4() {
	fmt.Println("Exercise 4a Answer is supposed to be 21138...")

	start := time.Now()
	fmt.Println("Day4a method 1 answer is: ", day4aMethod1())
	fmt.Println("method execution took", time.Since(start))

	start = time.Now()
	fmt.Println("Day4a method 2 answer is: ", day4aMethod2())
	fmt.Println("method execution took", time.Since(start))

	fmt.Println("Exercise 4b Answer is supposed to be xxxxx...")
	// start = time.Now()
	// fmt.Println("Day4b method 1 answer is: ", day4bMethod1())
	// fmt.Println("method execution took", time.Since(start))

}
