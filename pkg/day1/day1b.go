package day1

import (
	"fmt"
	"time"
)

func Answer1b() {
	fmt.Println("Exercise 1b Answer is 55093")

	start := time.Now()
	fmt.Println("using the char method the answer is: ", charAnswerb())
	fmt.Println("method execution took", time.Since(start))

	start = time.Now()
	fmt.Println("using the char method the answer is: ", charAnswer1b())
	fmt.Println("method execution took", time.Since(start))

	start = time.Now()
	fmt.Println("UtilSovle answer is: ", UtilSolve())
	fmt.Println("method execution took", time.Since(start))

}
