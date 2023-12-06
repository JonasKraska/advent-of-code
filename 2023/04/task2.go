package day04

import (
	"bufio"
	"strings"

	"github.com/JonasKraska/advent-of-code/2023/helper"
)

func Task2() int {
	input := helper.OpenFile("04/input.txt")
	defer input.Close()

	result := 0
	cards := make([]string, 0)
	amount := make([]int, 0)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		cards = append(cards, scanner.Text())
		amount = append(amount, 1)
	}

	for i, c := range cards {
		current := amount[i]

		card := strings.Split(c, ":")
		parts := strings.Split(card[1], "|")

		winners := strings.Split(strings.TrimSpace(parts[0]), " ")
		numbers := strings.Split(strings.TrimSpace(parts[1]), " ")

		similar := countSimilars(numbers, winners)

		for c := 1; c <= similar; c++ {
			amount[i+c] += current
		}
	}

	for _, a := range amount {
		result += a
	}

	return result
}
