package day1

import (
	"fmt"
	"time"
)

func Day1a() {
	fmt.Println("Exercise 1 Answer is supposed to be 55002...")

	start := time.Now()
	fmt.Println("using the regex method the answer is: ", regexAnswer())
	fmt.Println("method execution took", time.Since(start))

	start = time.Now()
	fmt.Println("using the char method the answer is: ", charAnswer())
	fmt.Println("method execution took", time.Since(start))

	start = time.Now()
	fmt.Println("using the unicode method the answer is: ", unicodeAnswer())
	fmt.Println("method execution took", time.Since(start))

}
