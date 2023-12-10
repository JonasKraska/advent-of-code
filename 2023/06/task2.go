package day06

import (
	"bufio"

	"github.com/JonasKraska/advent-of-code/2023/helper"
)

func Task2() int {
	input := helper.OpenFile("06/input2.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)

	scanner.Scan()
	times := helper.ParseNumbers(scanner.Text())

	scanner.Scan()
	distances := helper.ParseNumbers(scanner.Text())

	result := 1

	for i := range times {
		result *= analyzeCombinations(times[i], distances[i])
	}

	return result
}
