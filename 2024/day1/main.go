package main

import (
	utils "2024"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func part1(lines []string) int {
	left := make([]int, 0)
	right := make([]int, 0)

	for _, line := range lines {
		parts := strings.Split(line, "   ")
		num1, _ := strconv.Atoi(parts[0])
		num2, _ := strconv.Atoi(parts[1])
		left = append(left, num1)
		right = append(right, num2)
	}

	sort.Ints(left)
	sort.Ints(right)

	res := 0
	for i := range left {
		res += utils.AbsInt(left[i] - right[i])
	}

	return res
}

func part2(lines []string) int {
	left := make([]int, 0)
	right := make(map[int]int)

	for _, line := range lines {
		parts := strings.Split(line, "   ")
		num1, _ := strconv.Atoi(parts[0])
		num2, _ := strconv.Atoi(parts[1])
		left = append(left, num1)
		right[num2]++
	}

	res := 0
	for i := range left {
		res += left[i] * right[left[i]]
	}
	return res
}

func main() {
	lines := utils.ReadLines("input.txt")

	res1 := part1(lines)
	fmt.Printf("Part 1: %d\n", res1)

	res2 := part2(lines)
	fmt.Printf("Part 2: %d\n", res2)
}
