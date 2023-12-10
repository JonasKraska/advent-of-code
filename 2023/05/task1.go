package day05

import (
	"bufio"
	"math"
	"strings"

	"github.com/JonasKraska/advent-of-code/2023/helper"
)

const (
	SOIL = iota
	FERTILIZER
	WATER
	LIGHT
	TEMPERATURE
	HUMIDITY
	LOCATION
)

var seedIds []int
var seedAnalysed map[int]bool
var recipes [][]*recipe

func Task1() int {
	input := helper.OpenFile("05/input.txt")
	defer input.Close()

	section := 0
	seedIds = make([]int, 0)
	seedAnalysed = make(map[int]bool)
	recipes = make([][]*recipe, 7)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			section++
			continue
		}

		if section == 0 {
			parts := strings.Split(line, ":")
			numbers := strings.Split(strings.TrimSpace(parts[1]), " ")

			for _, id := range numbers {
				seedIds = append(seedIds, helper.StringToNumber(id))
			}

			continue
		}

		if strings.Contains(line, ":") {
			recipes[section-1] = make([]*recipe, 0)
			continue
		}

		parts := strings.Split(line, " ")
		recipes[section-1] = append(recipes[section-1], newRecipe(
			helper.StringToNumber(parts[0]),
			helper.StringToNumber(parts[1]),
			helper.StringToNumber(parts[2]),
		))
	}

	lowest := math.MaxInt32
	for _, id := range seedIds {
		if !seedAnalysed[id] {
			seedAnalysed[id] = true

			location := analyseSeedId(id)

			if location < lowest {
				lowest = location
			}
		}
	}

	return lowest
}

func analyseSeedId(id int) int {
	for r := range recipes {
		for _, rec := range recipes[r] {
			if rec.contains(id) {
				id = rec.calculate(id)
				break
			}
		}
	}

	return id
}
