package day07

import (
	"bufio"
	"sort"
	"strings"

	"github.com/JonasKraska/advent-of-code/2023/helper"
)

var cardValue2 = map[string]int{
	"J": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}

func Task2() int {
	input := helper.OpenFile("07/input.txt")
	defer input.Close()

	hands = make([]*hand, 0)
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		h := newHand(
			strings.Split(parts[0], ""),
			helper.StringToNumber(parts[1]),
		)

		calculateValue2(h)
		hands = append(hands, h)
	}

	sort.Slice(hands, sortHands2)

	result := 0
	for i, h := range hands {
		result += (i + 1) * h.bid
	}

	return result
}

func sortHands2(i, j int) bool {
	hand1 := hands[i]
	hand2 := hands[j]
	if hand1.value == hand2.value {
		for c := range hand1.cards {
			if hand1.cards[c] != hand2.cards[c] {
				return cardValue2[hand1.cards[c]] < cardValue2[hand2.cards[c]]
			}
		}
	}

	return hand1.value < hand2.value
}

func calculateValue2(h *hand) {
	cards := make(map[string]int)
	for _, c := range h.cards {
		if _, ok := cards[c]; ok {
			cards[c]++
		} else {
			cards[c] = 1
		}
	}

	foundJokers := 0
	if j, ok := cards["J"]; ok {
		foundJokers = j
	}
	delete(cards, "J")

	if foundJokers == 5 {
		h.value = FIVE_SAME
		return
	}

	foundTwo := 0
	foundThree := false
	for _, amount := range cards {
		if amount+foundJokers == 5 {
			h.value = FIVE_SAME
			return
		}

		if amount+foundJokers == 4 {
			h.value = FOUR_SAME
			return
		}

		if amount == 3 {
			foundThree = true
		}

		if amount == 2 {
			foundTwo++
		}
	}

	if (foundTwo == 2 && foundJokers == 1) || (foundTwo == 1 && foundThree) {
		h.value = FULL_HOUSE
		return
	}

	if foundJokers == 2 || (foundTwo == 1 && foundJokers == 1) || foundThree {
		h.value = THREE_SAME
		return
	}

	if foundTwo == 2 {
		h.value = TWO_PAIR
		return
	}

	if foundTwo == 1 || foundJokers == 1 {
		h.value = ONE_PAIR
		return
	}

	h.value = HIGH_CARD
}
