package day05

import (
	"bufio"
	"fmt"
	"math"
	"slices"
	"strings"
	"sync"

	"github.com/JonasKraska/advent-of-code/2023/helper"
)

var seedData [][]int
var mu sync.Mutex

func Task2() int {
	input := helper.OpenFile("05/input.txt")
	defer input.Close()

	section := 0
	mu = sync.Mutex{}
	seedData = make([][]int, 0)
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

			for i := 0; i < len(numbers); i += 2 {
				seedData = append(seedData, []int{
					helper.StringToNumber(numbers[i]),
					helper.StringToNumber(numbers[i+1]),
				})
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
	slices.Reverse(seedData)
	for d, data := range seedData {
		fmt.Printf("Dataset %d/%d\n", d+1, len(seedData))

		start := data[0]
		length := data[1]
		end := start + length

		for id := start; id < end; id++ {
			current := id - start + 1
			fmt.Printf("\r%d/%d (%.2f %%)", current, length, (float32(current) / float32(length) * 100))

			if !seedAnalysed[id] {
				seedAnalysed[id] = true

				go func(id int) {
					location := analyseSeedId(id)

					mu.Lock()
					if location < lowest {
						lowest = location
						fmt.Println("\nNew Lowest Location:", lowest)
					}
					mu.Unlock()
				}(id)
			}
		}

		fmt.Print("\n")
	}

	return lowest
}
