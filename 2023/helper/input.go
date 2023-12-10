package helper

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func OpenFile(path string) *os.File {
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func ParseNumbers(input string) []int {
	parts := strings.Split(input, ":")
	elem := strings.Split(parts[1], " ")

	numbers := make([]int, 0)
	for _, e := range elem {
		if n := strings.TrimSpace(e); n != "" {
			numbers = append(numbers, StringToNumber(n))
		}
	}

	return numbers
}
