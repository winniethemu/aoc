package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func findMax(batteries []string) (int, int) {
	mx, _ := strconv.Atoi(batteries[0])
	idx := 0
	for i := 1; i < len(batteries); i++ {
		val, _ := strconv.Atoi(batteries[i])
		if mx < val {
			mx = val
			idx = i
		}
	}
	return mx, idx
}

func part1(banks []string) int {
	res := 0
	for _, bank := range banks {
		batteries := strings.Split(bank, "")
		b1, i := findMax(batteries[:len(batteries)-1])
		b2, _ := findMax(batteries[i+1:])
		res += b1*10 + b2
	}
	return res
}

func part2(banks []string, size int) int {
	res := 0
	for _, bank := range banks {
		batteries := strings.Split(bank, "")
		sum := 0
		for sz := size - 1; sz >= 0; sz-- {
			battery, i := findMax(batteries[:len(batteries)-sz])
			sum *= 10
			sum += battery
			batteries = batteries[i+1:]
		}
		res += sum
	}
	return res
}

func main() {
	lines := utils.ReadLines("input.txt")

	res1 := part1(lines)
	fmt.Printf("Part 1: %d\n", res1)

	res2 := part2(lines, 12)
	fmt.Printf("Part 2: %d\n", res2)
}
