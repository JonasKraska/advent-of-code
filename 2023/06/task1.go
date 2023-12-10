package day06

import (
	"bufio"

	"github.com/JonasKraska/advent-of-code/2023/helper"
)

func Task1() int {
	input := helper.OpenFile("06/input.txt")
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

func analyzeCombinations(time, distance int) int {
	combinations := 0

	for t := 0; t < time; t++ {
		timeMoving := time - t
		distanceReached := timeMoving * t

		if distanceReached > distance {
			combinations++
		}
	}

	return combinations
}
