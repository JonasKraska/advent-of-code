package day02

import (
	"bufio"
	"log"
	"strconv"
	"strings"

	"github.com/JonasKraska/advent-of-code/2023/helper"
)

type CubeSet struct {
	Red   int
	Green int
	Blue  int
}

func Task1() int {
	input := helper.OpenFile("02/input.txt")
	defer input.Close()

	result := 0
	configuration := CubeSet{Red: 12, Green: 13, Blue: 14}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		gameId, rounds := parseGame(scanner.Text())

		if allRoundsValid(rounds, configuration) {
			result += gameId
		}
	}

	return result
}

func parseGame(line string) (int, []CubeSet) {
	parts := strings.Split(line, ":")
	id, err := strconv.Atoi(parts[0][5:])
	if err != nil {
		log.Fatal(err)
	}

	rounds := make([]CubeSet, 0)
	for _, r := range strings.Split(parts[1], ";") {
		set := CubeSet{}

		for _, c := range strings.Split(r, ",") {
			v := strings.Split(strings.TrimSpace(c), " ")
			amount, err := strconv.Atoi(v[0])
			if err != nil {
				log.Fatal(err)
			}

			if v[1] == "red" {
				set.Red = amount
			}

			if v[1] == "green" {
				set.Green = amount
			}

			if v[1] == "blue" {
				set.Blue = amount
			}
		}

		rounds = append(rounds, set)
	}

	return id, rounds
}

func allRoundsValid(rounds []CubeSet, configuration CubeSet) bool {
	for _, round := range rounds {
		if round.Red > configuration.Red {
			return false
		}
		if round.Green > configuration.Green {
			return false
		}
		if round.Blue > configuration.Blue {
			return false
		}
	}

	return true
}
