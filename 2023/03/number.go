package day03

import (
	"log"

	"github.com/JonasKraska/advent-of-code/2023/helper"
)

type number struct {
	start int
	end   int
	value string
}

func newNumber(start int, ch rune) *number {
	if !helper.IsDigit(ch) {
		log.Fatal("Can not create number from non numeric character")
	}

	return &number{
		start: start,
		end:   -1,
		value: string(ch),
	}
}

func (n *number) isFinished() bool {
	return n.end >= 0
}

func (n *number) append(ch rune) {
	if n.isFinished() {
		log.Fatal("Can not append to finsied number")
	}

	if !helper.IsDigit(ch) {
		log.Fatal("Can not create number from non numeric character")
	}

	n.value = n.value + string(ch)
}

func (n *number) finish(end int) {
	n.end = end
}

func (n *number) toInt() int {
	if !n.isFinished() {
		log.Fatal("Can not wrap up an unfinished number")
	}

	return helper.StringToNumber(n.value)
}
