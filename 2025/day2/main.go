package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

// s is made of repeated substrings if and only if:
// s is a substring of (s+s) with the first and last character removed
func invalid2(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	}
	ss := (s + s)[1 : 2*n-1]
	return strings.Contains(ss, s)
}

func invalid1(s string) bool {
	mid := len(s) / 2
	return s[:mid] == s[mid:]
}

func endpoint(interval string) (int, int) {
	values := strings.Split(interval, "-")
	start, _ := strconv.Atoi(values[0])
	end, _ := strconv.Atoi(values[1])
	return start, end
}

func solve(intervals []string, invalid func(s string) bool) int {
	res := 0

	for _, interval := range intervals {
		start, end := endpoint(interval)
		for i := start; i <= end; i++ {
			if invalid(strconv.Itoa(i)) {
				res += i
			}
		}
	}

	return res
}

func main() {
	ranges := utils.ReadInput("input.txt", ",")

	res1 := solve(ranges, invalid1)
	fmt.Printf("Part 1: %d\n", res1)

	res2 := solve(ranges, invalid2)
	fmt.Printf("Part 2: %d\n", res2)
}
