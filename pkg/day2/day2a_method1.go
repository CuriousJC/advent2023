package day2

import (
	"fmt"
	"github.com/curiousjc/advent2023/internal/inputs"
	"regexp"
	"strconv"
)

// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
// Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
// Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
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
			total += 0 //impossible game
		} else {
			total += getGameNumber(line)
		}
	}
	return total
}

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

func getGameNumber(line string) int {

	//result: Game 53
	pattern := regexp.MustCompile(`^[^:]+`)
	gameLabel := pattern.FindString(line)

	//result 53
	pattern = regexp.MustCompile(`\d{1,3}`)
	gameNumber := pattern.FindString(gameLabel)

	gameInt, err := strconv.Atoi(gameNumber)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		panic(err)
	}
	return gameInt
}
