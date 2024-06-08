package day4

import (
	"slices"
)

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
