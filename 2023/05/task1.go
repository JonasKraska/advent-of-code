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

var seeds []*seed
var recipes [][]*recipe

func Task1() int {
	input := helper.OpenFile("05/input.txt")
	defer input.Close()

	section := 0
	seeds = make([]*seed, 0)
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
			ids := strings.Split(strings.TrimSpace(parts[1]), " ")

			for _, i := range ids {
				seeds = append(seeds, newSeed(helper.StringToNumber(i)))
			}
		} else {
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
	}

	lowest := math.MaxInt32
	for _, s := range seeds {
		analyseSeed(s)

		if s.data[LOCATION] < lowest {
			lowest = s.data[LOCATION]
		}
	}

	return lowest
}

func analyseSeed(s *seed) {
	id := s.seed
	for r := range recipes {
		found := false
		for _, rec := range recipes[r] {
			if rec.contains(id) {
				id = rec.calculate(id)
				s.data[r] = id
				found = true
				break
			}
		}

		if !found {
			s.data[r] = id
		}
	}
}
