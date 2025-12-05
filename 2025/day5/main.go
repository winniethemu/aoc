package main

import (
	"aoc/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func fresh(intervals [][2]int, id int) bool {
	for _, interval := range intervals {
		if interval[0] <= id && id <= interval[1] {
			return true
		}
	}
	return false
}

func part1(intervals [][2]int, ingredients []int) int {
	res := 0

	for _, ingredient := range ingredients {
		if fresh(intervals, ingredient) {
			res++
		}
	}

	return res
}

func part2(intervals [][2]int) int {
	res := 0

	// merge intervals
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := make([][2]int, 0)
	result = append(result, intervals[0])
	for i := 1; i < len(intervals); i++ {
		prev := result[len(result)-1]
		curr := intervals[i]
		if prev[1] < curr[0] {
			result = append(result, curr)
		} else {
			start := min(prev[0], curr[0])
			end := max(prev[1], curr[1])
			result = result[:len(result)-1]
			result = append(result, [2]int{start, end})
		}
	}

	for _, item := range result {
		res += item[1] - item[0] + 1
	}

	return res
}

func parseInterval(str string) [2]int {
	parts := strings.Split(str, "-")
	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])
	return [2]int{start, end}
}

func main() {
	lines := utils.ReadLines("input.txt")

	intervals := make([][2]int, 0)
	ingredients := make([]int, 0)
	readingInterval := true
	for _, line := range lines {
		if line == "" {
			readingInterval = false
			continue
		}

		if readingInterval {
			intervals = append(intervals, parseInterval(line))
		} else {
			ingredient, _ := strconv.Atoi(line)
			ingredients = append(ingredients, ingredient)
		}
	}

	res1 := part1(intervals, ingredients)
	fmt.Printf("Part 1: %d\n", res1)

	res2 := part2(intervals)
	fmt.Printf("Part 2: %d\n", res2)
}
