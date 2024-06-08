package day4

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// break our line into the card numbers and the winning numbers
func crackCardAndWinners(line string) (card []int, winners []int) {
	// Find the index of the colon and pipe
	colonIndex := strings.Index(line, ":")
	pipeIndex := strings.Index(line, "|")

	// Extract the substrings between the colon and pipe, and after the pipe
	cardNums := strings.TrimSpace(line[colonIndex+1 : pipeIndex])
	winNums := strings.TrimSpace(line[pipeIndex+1:])

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

	return card, winners
}

// given the card and the winners get the score
func scoreCard(card []int, winners []int) int {
	slices.Sort(card)
	slices.Sort(winners)

	var winnerIndex int
	var cardIndex int
	var winnerCount int
	for cardIndex < len(card) && winnerIndex < len(winners) {
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

	//double the score for each winning number
	var score int
	if winnerCount > 0 {
		score = 1
		for i := 1; i < winnerCount; i++ {
			score = score * 2
		}
	}

	return score
}
