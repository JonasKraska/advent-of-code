package day04

import (
	"bufio"
	"math"
	"strings"

	"github.com/JonasKraska/advent-of-code/2023/helper"
)

func Task1() int {
	input := helper.OpenFile("04/input.txt")
	defer input.Close()

	result := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		card := strings.Split(line, ":")
		parts := strings.Split(card[1], "|")

		winners := strings.Split(strings.TrimSpace(parts[0]), " ")
		numbers := strings.Split(strings.TrimSpace(parts[1]), " ")

		similar := countSimilars(numbers, winners)

		if similar > 0 {
			result += int(math.Pow(2, float64(similar-1)))
		}
	}

	return result
}

func countSimilars(numbers, winners []string) int {
	counter := 0

	for _, n := range numbers {
		for _, w := range winners {
			if n != "" && n == w {
				counter++
			}
		}
	}

	return counter
}
