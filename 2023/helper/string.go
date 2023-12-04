package helper

import (
	"log"
	"strconv"
	"unicode"
)

func IsDigit(r rune) bool {
	return unicode.IsDigit(r)
}

func RuneToInt(r rune) int {
	return int(r - '0')
}

func StringToNumber(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		log.Fatal(err)
	}

	return i
}
