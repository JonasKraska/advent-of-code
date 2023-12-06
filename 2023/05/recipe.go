package day05

import (
	"log"
)

type recipe struct {
	destinationStart int
	sourceStart      int
	length           int
}

func newRecipe(destinationStart, sourceStart, length int) *recipe {
	return &recipe{
		destinationStart: destinationStart,
		sourceStart:      sourceStart,
		length:           length,
	}
}

func (r *recipe) contains(id int) bool {
	return r.sourceStart <= id && r.sourceStart+r.length > id
}

func (r *recipe) calculate(id int) int {
	if !r.contains(id) {
		log.Fatal("Seed is not contained in given atlas map!")
	}

	diff := id - r.sourceStart

	return r.destinationStart + diff
}
