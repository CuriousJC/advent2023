package day2

import (
	"fmt"
	"github.com/curiousjc/advent2023/internal/inputs"
	"regexp"
	"strconv"
)

func method1() (total int) {
	fmt.Println("running day2 method1...")
	lines, err := inputs.GetInput("static/input2.txt")
	if err != nil {
		panic(err)
	}

	maxRed := 12
	maxGreen := 13
	maxBlue := 14

	for _, line := range lines {
		highestRed := findHighest(line, "red")
		highestGreen := findHighest(line, "green")
		highestBlue := findHighest(line, "blue")

		if highestRed > maxRed || highestGreen > maxGreen || highestBlue > maxBlue {
			total += 0
		} else {
			total += getGameNumber(line)
		}
	}
	return total
}

func findHighest(line string, color string) int {

	//break out all our pulls for the color we're processing
	colorPattern := fmt.Sprintf(`\b\d{1,2}\s+%s\b`, color)
	pattern := regexp.MustCompile(colorPattern)
	colorPulls := pattern.FindAllString(line, -1)

	//extract number pulled
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

func getGameNumber(line string) int {

	pattern := regexp.MustCompile(`^[^:]+`)
	gameLabel := pattern.FindString(line)

	pattern = regexp.MustCompile(`\d{1,3}`)
	gameNumber := pattern.FindString(gameLabel)

	gameInt, err := strconv.Atoi(gameNumber)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		panic(err)
	}

	return gameInt
}
