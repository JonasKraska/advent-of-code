package day03

import (
	"bufio"

	"github.com/JonasKraska/advent-of-code/2023/helper"
)

var gears [][][]*number

func Task2() int {
	input := helper.OpenFile("03/input.txt")
	defer input.Close()

	rows := make([][]rune, 0)
	gears = make([][][]*number, 0)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		row := []rune(scanner.Text())
		rows = append(rows, row)
		gears = append(gears, make([][]*number, len(row)))
	}

	result := 0

	for r := range rows {
		analyzeForGears(r, rows)
	}

	for y := range gears {
		for x := range gears[y] {
			if len(gears[y][x]) > 1 {
				product := 1

				for _, n := range gears[y][x] {
					product *= n.toInt()
				}

				result += product
			}
		}
	}

	return result
}

func analyzeForGears(row int, rows [][]rune) {
	var number *number
	number = nil

	for i, ch := range rows[row] {
		if helper.IsDigit(ch) && number == nil {
			number = newNumber(i, ch)
			continue
		}

		if helper.IsDigit(ch) && number != nil {
			number.append(ch)
			continue
		}

		if !helper.IsDigit(ch) && number != nil {
			number.finish(i - 1)
			isGear(number, row, rows)

			number = nil
		}
	}

	if number != nil {
		number.finish(len(rows[row]) - 1)
		isGear(number, row, rows)
	}
}

func isGear(n *number, row int, rows [][]rune) {
	for y := row - 1; y <= row+1; y++ {
		if y < 0 || y > len(rows)-1 {
			continue
		}

		for x := n.start - 1; x <= n.end+1; x++ {
			if x < 0 || x > len(rows[y])-1 {
				continue
			}

			ch := rows[y][x]

			if string(ch) == "*" {
				gears[y][x] = append(gears[y][x], n)
			}
		}
	}
}
