package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
)

func part1(lines []string) int {
	password := 0
	dial := 50
	for _, line := range lines {
		dist, _ := strconv.Atoi(line[1:])
		dir := 1
		if string(line[0]) == "L" {
			dir = -1
		}
		dial += (dir * dist) % 100
		if dial < 0 {
			dial += 100
		} else if dial > 99 {
			dial -= 100
		}
		if dial == 0 {
			password++
		}
	}
	return password
}

func part2(lines []string) int {
	password := 0
	dial := 50
	for _, line := range lines {
		dist, _ := strconv.Atoi(line[1:])
		dir := 1
		if string(line[0]) == "L" {
			dir = -1
		}
		password += dist / 100
		next := dial + ((dir * dist) % 100)
		if dial*next < 0 || (dial < 100 && next > 100) {
			password++
		}
		dial = next
		if dial < 0 {
			dial += 100
		} else if dial > 99 {
			dial -= 100
		}
		if dial == 0 {
			password++
		}
	}
	return password
}

func main() {
	lines := utils.ReadLines("input.txt")
	res1 := part1(lines)
	fmt.Printf("Part 1: %d\n", res1)

	res2 := part2(lines)
	fmt.Printf("Part 2: %d\n", res2)
}
