package day4

import (
	"fmt"
	"github.com/curiousjc/advent2023/internal/inputs"
	"slices"
)

/*
Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
1 instance of card 1,
2 instances of card 2,
4 instances of card 3,
8 instances of card 4,
14 instances of card 5, and
1 instance of card 6. In total, this example pile of scratchcards causes you to ultimately have 30 scratchcards!
*/
func day4bMethod1() (total int) {
	fmt.Println("running day4a method1...")
	lines, err := inputs.Get("static/day4.txt")
	if err != nil {
		panic(err)
	}

	//create a slice of our cards and the count of copies of that card we have, initially starting with one
	cardCounts := make([]int, len(lines))
	for cardCountIndex := range cardCounts {
		cardCounts[cardCountIndex] = 1
	}

	//for each of our cards we're going to process it, paying attention to the number of that card we have
	for cardCountIndex, cardCountNumber := range cardCounts {

		//identify how many wins the card has
		winners := procCard(lines[cardCountIndex])

		//for however many copies we have we will need to add our wins to subsequent cards that many times
		for i := 1; i <= cardCountNumber; i++ {

			//for each win increment the next card
			for w := 1; w <= winners; w++ {
				nextCardIndex := cardCountIndex + w
				cardCounts[nextCardIndex] += 1
			}
		}
	}

	//count up how many cards total
	for _, cardCount := range cardCounts {
		total = total + cardCount
	}

	return total
}

// given a card line get the number of wins
func procCard(line string) int {

	//Crack our string into the card numbers and the winning numbers
	card, winners := crackCardAndWinners(line)

	//get our wins
	return countWins(card, winners)
}

// given the card numbers and the winning numbers, count the wins
func countWins(card []int, winners []int) int {
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

	return winnerCount
}
