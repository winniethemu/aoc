package utils

import (
	"os"
	"strings"
)

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ReadLines(filename string) []string {
	return readInput(filename, "\n")
}

func readInput(filename, delimiter string) []string {
	dat, err := os.ReadFile(filename)
	if err != nil {
		panic("error reading the file")
	}
	return strings.Split(string(dat), delimiter)
}
