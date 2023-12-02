package day_02

import (
	"bufio"

	"github.com/JonasKraska/advent-of-code/2023/helper"
)

func Task2() int {
	input := helper.OpenFile("02/input.txt")
	defer input.Close()

	result := 0
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		_, rounds := parseGame(scanner.Text())

		result += calculatePower(rounds)
	}

	return result
}

func calculatePower(rounds []CubeSet) int {
	needed := CubeSet{Red: 1, Green: 1, Blue: 1}

	for _, r := range rounds {
		if r.Red > needed.Red {
			needed.Red = r.Red
		}
		if r.Green > needed.Green {
			needed.Green = r.Green
		}
		if r.Blue > needed.Blue {
			needed.Blue = r.Blue
		}
	}

	return needed.Red * needed.Green * needed.Blue
}
