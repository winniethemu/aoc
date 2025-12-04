package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func accessible(mat [][]string, row, col int) bool {
	n, m := len(mat), len(mat[0])
	neighbors := [8][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	count := 0
	for _, neighbor := range neighbors {
		dr, dc := neighbor[0], neighbor[1]
		if 0 <= dr+row && dr+row < n && 0 <= dc+col && dc+col < m {
			if mat[dr+row][dc+col] == "@" {
				count++
			}
		}
	}
	return count < 4
}

func part1(mat [][]string) int {
	res := 0
	n, m := len(mat), len(mat[0])
	for i := range n {
		for j := range m {
			if mat[i][j] == "@" && accessible(mat, i, j) {
				res++
			}
		}
	}
	return res
}

func main() {
	lines := utils.ReadLines("input.txt")
	mat := make([][]string, len(lines))
	for i := 0; i < len(lines); i++ {
		chars := strings.Split(lines[i], "")
		mat[i] = chars
	}
	res1 := part1(mat)
	fmt.Printf("Part 1: %d\n", res1)
}
