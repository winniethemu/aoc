package utils

import (
	"os"
	"strings"
)

func ReadLines(filename string) []string {
	dat, err := os.ReadFile(filename)
	if err != nil {
		panic("error reading the file")
	}
	return strings.Split(string(dat), "\n")
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
