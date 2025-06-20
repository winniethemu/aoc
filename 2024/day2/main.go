package main

import (
	utils "2024"
	"fmt"
	"strconv"
	"strings"
)

func validate(levels []int) bool {
	monotonicity := 0
	if levels[0] < levels[1] {
		monotonicity = 1
	} else if levels[0] > levels[1] {
		monotonicity = -1
	} else {
		return false
	}

	for i := 1; i < len(levels); i++ {
		prev, curr := levels[i-1], levels[i]
		if utils.AbsInt(curr-prev) < 1 || utils.AbsInt(curr-prev) > 3 {
			return false
		} else if (curr-prev)/utils.AbsInt(curr-prev) != monotonicity {
			return false
		}
	}

	return true
}

func convert(s []string) []int {
	res := make([]int, len(s))
	for idx, item := range s {
		num, err := strconv.Atoi(item)
		if err != nil {
			panic("error converting string to number, bad input")
		}
		res[idx] = num
	}
	return res
}

func part1(lines []string) int {
	safe := 0
	for _, line := range lines {
		sslice := strings.Split(line, " ")
		islice := convert(sslice)
		if validate(islice) {
			safe++
		}
	}
	return safe
}

func main() {
	lines := utils.ReadLines("input.txt")
	res1 := part1(lines)
	fmt.Printf("Part 1: %d\n", res1)
}
