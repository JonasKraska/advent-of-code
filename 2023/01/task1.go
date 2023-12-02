package day_01

import (
	"bufio"
	"slices"
	"unicode"

	"github.com/JonasKraska/advent-of-code/2023/helper"
)

func Task1() int {
	input := helper.OpenFile("01/input.txt")
	defer input.Close()

	result := 0
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()

		first := findFirstDigit(line)
		last := findFirstDigit(revserseLine(line))

		result += first*10 + last
	}

	return result
}

func findFirstDigit(input string) int {
	for _, ch := range input {
		if unicode.IsDigit(ch) {
			return int(ch - '0')
		}
	}

	panic("No digit found!")
}

func revserseLine(line string) string {
	chars := []rune(line)
	slices.Reverse(chars)
	return string(chars)
}
