package day03

import (
	"bufio"

	"github.com/JonasKraska/advent-of-code/2023/helper"
)

func Task1() int {
	input := helper.OpenFile("03/input.txt")
	defer input.Close()

	rows := make([][]rune, 0)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		row := []rune(scanner.Text())
		rows = append(rows, row)
	}

	result := 0

	for r := range rows {
		result += analyzeForParts(r, rows)
	}

	return result
}

func analyzeForParts(row int, rows [][]rune) int {
	sum := 0

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

			if isPart(number, row, rows) {
				sum += number.toInt()
			}

			number = nil
		}
	}

	if number != nil {
		number.finish(len(rows[row]) - 1)

		if isPart(number, row, rows) {
			sum += number.toInt()
		}
	}

	return sum
}

func isPart(n *number, row int, rows [][]rune) bool {
	for y := row - 1; y <= row+1; y++ {
		if y < 0 || y > len(rows)-1 {
			continue
		}

		for x := n.start - 1; x <= n.end+1; x++ {
			if x < 0 || x > len(rows[y])-1 {
				continue
			}

			ch := rows[y][x]

			if !helper.IsDigit(ch) && string(ch) != "." {
				return true
			}
		}
	}

	return false
}
