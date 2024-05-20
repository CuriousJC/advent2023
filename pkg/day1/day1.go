package day1

import (
	"fmt"
	"time"
)

func Day1() {
	fmt.Println("------")
	fmt.Println("Exercise 1a Answer is supposed to be 55002...")

	start := time.Now()
	fmt.Println("using the regex method the answer is: ", regexAnswer())
	fmt.Println("method execution took", time.Since(start))

	start = time.Now()
	fmt.Println("using the char method the answer is: ", charAnswer())
	fmt.Println("method execution took", time.Since(start))

	start = time.Now()
	fmt.Println("using the unicode method the answer is: ", unicodeAnswer())
	fmt.Println("method execution took", time.Since(start))

	fmt.Println("------")
	fmt.Println("Exercise 1b Answer is 55093")
	start = time.Now()
	fmt.Println("using the FAILED char method the answer is: ", charAnswerb())
	fmt.Println("method execution took", time.Since(start))

	start = time.Now()
	fmt.Println("using the char method the answer is: ", charAnswer1b())
	fmt.Println("method execution took", time.Since(start))

	start = time.Now()
	fmt.Println("using the speedy char method the answer is: ", charAnswer1bSpeedy())
	fmt.Println("method execution took", time.Since(start))

	start = time.Now()
	fmt.Println("using the speedy speedy char method the answer is: ", charAnswer1bSpeedySpeedy())
	fmt.Println("method execution took", time.Since(start))

	start = time.Now()
	fmt.Println("UtilSolve answer is: ", UtilSolve())
	fmt.Println("method execution took", time.Since(start))

}
