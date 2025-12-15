package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

var reds [][2]int

func leftDiagonal(pos1, pos2 [2]int) bool {
	return pos1[0] < pos2[0] && pos1[1] > pos2[1]
}

func diagonals(pos1, pos2 [2]int) bool {
	for _, red := range reds {
		if leftDiagonal(pos1, pos2) {
			if (red[0] <= pos1[0] && red[1] <= pos2[1]) ||
				(red[0] >= pos2[0] && red[1] >= pos1[1]) {
				return true
			}
		} else {
			if (red[0] <= pos1[0] && red[1] >= pos2[1]) ||
				(red[0] >= pos2[0] && red[1] <= pos1[1]) {
				return true
			}
		}
	}
	return false
}

func inbetween(pos1, pos2 [2]int) bool {
	for _, red := range reds {
		if pos1[0] < red[0] && red[0] < pos2[0] &&
			pos1[1] < red[1] && red[1] < pos2[1] {
			return true
		}
	}
	return false

}

func valid(pos1, pos2 [2]int) bool {
	if inbetween(pos1, pos2) {
		return false
	}

	if !diagonals(pos1, pos2) {
		return false
	}

	return true
}

func area(pos1, pos2 [2]int) int {
	return (utils.AbsInt(pos1[0]-pos2[0]) + 1) * (utils.AbsInt(pos1[1]-pos2[1]) + 1)
}

func solve(validator func(p1, p2 [2]int) bool) int {
	res := 0
	n := len(reds)
	for i := range n {
		for j := i + 1; j < n; j++ {
			valid := true
			if validator != nil {
				valid = validator(reds[i], reds[j])
			}
			if valid {
				a := area(reds[i], reds[j])
				if res < a {
					res = a
				}
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
	lines := utils.ReadLines("example.txt")
	reds = parse(lines)

	res1 := solve(nil)
	fmt.Printf("Part 1: %d\n", res1)

	res2 := solve(valid)
	fmt.Printf("Part 2: %d\n", res2)
}
