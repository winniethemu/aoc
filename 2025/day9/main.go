package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func area(pos1, pos2 [2]int) int {
	return (utils.AbsInt(pos1[0]-pos2[0]) + 1) * (utils.AbsInt(pos1[1]-pos2[1]) + 1)
}

func part1(positions [][2]int) int {
	res := 0
	n := len(positions)
	for i := range n {
		for j := i + 1; j < n; j++ {
			a := area(positions[i], positions[j])
			if res < a {
				res = a
			}
		}
	}
	return res
}

func parse(lines []string) [][2]int {
	res := make([][2]int, len(lines))
	for i, line := range lines {
		xy := strings.Split(line, ",")
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])
		res[i] = [2]int{x, y}
	}
	return res
}

func main() {
	lines := utils.ReadLines("input.txt")

	res1 := part1(parse(lines))
	fmt.Printf("Part 1: %d\n", res1)
}
