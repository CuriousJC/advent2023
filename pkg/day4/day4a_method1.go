package day4

import (
	"fmt"
	"github.com/curiousjc/advent2023/internal/inputs"
	"slices"
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
ANSWER: 8 + 2 + 2 + 1 + 0 + 0 = 13 points total
*/
func day4aMethod1() (total int) {
	fmt.Println("running day4a method1...")
	lines, err := inputs.Get("static/day4.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		//fmt.Println(line)

		//Break apart our string into the card numbers and the winning numbers
		card, winners := crackCardAndWinners(line)

		//get our score, add to our total
		total += scoreCard(card, winners)

	}

	return total
}

func crackCardAndWinners(line string) (card []int, winners []int) {
	//TODO regex crack instead of substring?

	// Find the index of the colon and pipe
	colonIndex := strings.Index(line, ":")
	pipeIndex := strings.Index(line, "|")

	// Extract the substrings between the colon and pipe, and after the pipe
	firstPart := strings.TrimSpace(line[colonIndex+1 : pipeIndex])
	secondPart := strings.TrimSpace(line[pipeIndex+1:])

	// Split these substrings by whitespace
	firstSlice := strings.Fields(firstPart)
	secondSlice := strings.Fields(secondPart)

	// Convert the slices of strings to slices of integers
	firstInts := make([]int, len(firstSlice))
	secondInts := make([]int, len(secondSlice))

	for i, v := range firstSlice {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return
		}
		firstInts[i] = num
	}

	for i, v := range secondSlice {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return
		}
		secondInts[i] = num
	}

	return firstInts, secondInts

}

func scoreCard(card []int, winners []int) int {
	slices.Sort(card)
	slices.Sort(winners)

	var winnerIndex int
	var cardIndex int
	var winnerCount int
	for cardIndex < len(card) && winnerIndex < len(winners) {

		//fmt.Println(cardIndex, " ", winnerIndex, " ", card[cardIndex], " ", winners[winnerIndex])
		if card[cardIndex] == winners[winnerIndex] {
			winnerCount++
			winnerIndex++
			cardIndex++
			continue
		}
		if card[cardIndex] > winners[winnerIndex] {
			winnerIndex++
			continue
		}
		if card[cardIndex] < winners[winnerIndex] {
			cardIndex++
			continue
		}

	}

	var score int
	if winnerCount > 0 {
		score = 1
		for i := 1; i < winnerCount; i++ {
			score = score * 2
		}
	}

	return score

}
