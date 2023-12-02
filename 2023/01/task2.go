package day_01

import (
	"bufio"
	"strings"
	"unicode"

	"github.com/JonasKraska/advent-of-code/2023/helper"
)

func Task2() int {
	input := helper.OpenFile("01/input.txt")
	defer input.Close()

	result := 0
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()

		first := findFirstNumber(line)
		last := findFLasttNumber(line)

		result += first*10 + last
	}

	return result
}

func findFirstNumber(input string) int {
	chars := []rune(input)
	for i := 0; i < len(input); i++ {
		if unicode.IsDigit(chars[i]) {
			return int(chars[i] - '0')
		}

		if strings.HasPrefix(string(chars[i:]), "one") {
			return 1
		}
		if strings.HasPrefix(string(chars[i:]), "two") {
			return 2
		}
		if strings.HasPrefix(string(chars[i:]), "three") {
			return 3
		}
		if strings.HasPrefix(string(chars[i:]), "four") {
			return 4
		}
		if strings.HasPrefix(string(chars[i:]), "five") {
			return 5
		}
		if strings.HasPrefix(string(chars[i:]), "six") {
			return 6
		}
		if strings.HasPrefix(string(chars[i:]), "seven") {
			return 7
		}
		if strings.HasPrefix(string(chars[i:]), "eight") {
			return 8
		}
		if strings.HasPrefix(string(chars[i:]), "nine") {
			return 9
		}
	}

	panic("No digit found!")
}

func findFLasttNumber(input string) int {
	chars := []rune(input)
	for i := len(input) - 1; i >= 0; i-- {
		if unicode.IsDigit(chars[i]) {
			return int(chars[i] - '0')
		}

		if strings.HasSuffix(string(chars[0:i+1]), "one") {
			return 1
		}
		if strings.HasSuffix(string(chars[0:i+1]), "two") {
			return 2
		}
		if strings.HasSuffix(string(chars[0:i+1]), "three") {
			return 3
		}
		if strings.HasSuffix(string(chars[0:i+1]), "four") {
			return 4
		}
		if strings.HasSuffix(string(chars[0:i+1]), "five") {
			return 5
		}
		if strings.HasSuffix(string(chars[0:i+1]), "six") {
			return 6
		}
		if strings.HasSuffix(string(chars[0:i+1]), "seven") {
			return 7
		}
		if strings.HasSuffix(string(chars[0:i+1]), "eight") {
			return 8
		}
		if strings.HasSuffix(string(chars[0:i+1]), "nine") {
			return 9
		}
	}

	panic("No digit found!")
}
