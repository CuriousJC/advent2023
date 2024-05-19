package day1

import (
	"github.com/curiousjc/advent2023/internal/inputs"
	"regexp"
	"strconv"
)

func regexAnswer() (total int) {
	lines, err := inputs.Get("static/day1.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		value := extractValueByRegex(line)
		//fmt.Print(line, "value is: ", value, "\n")
		total += value
	}
	return total
}

func extractValueByRegex(input string) (value int) {
	regex := regexp.MustCompile(`\d`)

	first, err := strconv.Atoi(regex.FindString(input))
	if err != nil {
		panic(err)
	}

	all := regex.FindAllString(input, -1)
	lastIndex := len(all) - 1
	last, err := strconv.Atoi(all[lastIndex])
	if err != nil {
		panic(err)
	}

	value = (first * 10) + last

	return value
}
