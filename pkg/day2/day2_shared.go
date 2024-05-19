package day2

import (
	"fmt"
	"regexp"
	"strconv"
)

func findHighest(line string, color string) int {

	//break out all our pulls for the color we're processing.  results: '13 red', '5 red', '2 red'
	colorPattern := fmt.Sprintf(`\b\d{1,2}\s+%s\b`, color)
	pattern := regexp.MustCompile(colorPattern)
	colorPulls := pattern.FindAllString(line, -1)

	//extract number pulled. results: 13, 5, 2
	var pulls []int
	pulls = append(pulls, -1) //ensure at least one item
	for _, colorPull := range colorPulls {
		pattern = regexp.MustCompile(`\d{1,2}`)
		pullCount := pattern.FindString(colorPull)
		pullInt, err := strconv.Atoi(pullCount)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			panic(err)
		}
		pulls = append(pulls, pullInt)
	}

	//find highest number pulled
	max := pulls[0]
	for _, value := range pulls {
		if value > max {
			max = value
		}
	}
	return max
}
