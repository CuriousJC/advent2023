package day4

import (
	"fmt"
	"github.com/curiousjc/advent2023/internal/inputs"
	"regexp"
	"strconv"
	"strings"
)

/*
Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
*/
func day4aMethod2() (total int) {
	fmt.Println("running day4b method1...")

	lines, err := inputs.Get("static/day4.txt")
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		//Crack our string into the card numbers and the winning numbers
		card, winners := crackCardAndWinnersReggie(line)

		//get our score, add to our total
		total += scoreCard(card, winners)

	}

	return total
}

func crackCardAndWinnersReggie(line string) (card []int, winners []int) {
	// Define the regular expression pattern to match the substrings between colon and pipe
	//re := regexp.MustCompile(`:(.*?)\|`)
	//re := regexp.MustCompile(`:(.*?)\| (.*?)$`)
	re := regexp.MustCompile(`:(.*?)(?:\| )?(.*?)$`)

	// Find all matches of the pattern in the input line
	matches := re.FindAllStringSubmatch(line, -1)

	//TODO rewrite this crap.  capture everything between colon and pipe as one set, then capture all the stuff after the pipe as the second

	fmt.Println(line)
	fmt.Println(matches)

	// Extract the matched substrings
	if len(matches) == 1 && len(matches[0]) == 2 {
		cardNums := strings.TrimSpace(matches[0][1])
		winNums := strings.TrimSpace(matches[0][1])

		// Split these substrings by whitespace
		cardNumStrings := strings.Fields(cardNums)
		winNumStrings := strings.Fields(winNums)

		// Convert the slices of strings to slices of integers
		card = make([]int, len(cardNumStrings))
		winners = make([]int, len(winNumStrings))

		for i, v := range cardNumStrings {
			num, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return
			}
			card[i] = num
		}

		for i, v := range winNumStrings {
			num, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return
			}
			winners[i] = num
		}
	}

	return card, winners

}
