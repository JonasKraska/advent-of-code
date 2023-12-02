package helper

import (
	"log"
	"os"
	"path/filepath"
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
