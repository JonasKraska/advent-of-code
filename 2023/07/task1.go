package day07

import (
	"bufio"
	"sort"
	"strings"

	"github.com/JonasKraska/advent-of-code/2023/helper"
)

var cardValue = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

var hands []*hand

func Task1() int {
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

		calculateValue(h)
		hands = append(hands, h)
	}

	sort.Slice(hands, sortHands)

	result := 0
	for i, h := range hands {
		result += (i + 1) * h.bid
	}

	return result
}

func sortHands(i, j int) bool {
	hand1 := hands[i]
	hand2 := hands[j]
	if hand1.value == hand2.value {
		for c := range hand1.cards {
			if hand1.cards[c] != hand2.cards[c] {
				return cardValue[hand1.cards[c]] < cardValue[hand2.cards[c]]
			}
		}
	}

	return hand1.value < hand2.value
}

func calculateValue(h *hand) {
	cards := make(map[string]int)
	for _, c := range h.cards {
		if _, ok := cards[c]; ok {
			cards[c]++
		} else {
			cards[c] = 1
		}
	}

	foundTwo := 0
	foundThree := false
	for _, amount := range cards {
		if amount == 5 {
			h.value = FIVE_SAME
			return
		}

		if amount == 4 {
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

	if foundTwo == 1 && foundThree {
		h.value = FULL_HOUSE
		return
	}

	if foundTwo == 0 && foundThree {
		h.value = THREE_SAME
		return
	}

	if foundTwo == 2 {
		h.value = TWO_PAIR
		return
	}

	if foundTwo == 1 {
		h.value = ONE_PAIR
		return
	}

	h.value = HIGH_CARD
}
