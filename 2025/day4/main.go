package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

var neighbors = [8][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func inBounds(mat [][]string, row, col int) bool {
	return row >= 0 && row < len(mat) && col >= 0 && col < len(mat[0])
}

func accessible(mat [][]string, row, col int) bool {
	count := 0
	for _, neighbor := range neighbors {
		nr, nc := row+neighbor[0], col+neighbor[1]
		if inBounds(mat, nr, nc) && mat[nr][nc] == "@" {
			count++
		}
	}
	return count < 4
}

func countRolls(mat [][]string) int {
	count := 0
	for i := range len(mat) {
		for j := range len(mat[0]) {
			if mat[i][j] == "@" {
				count++
			}
		}
	}
	return count
}

// returns all accessible locations
func part1(mat [][]string) [][2]int {
	var locations [][2]int
	for i := range len(mat) {
		for j := range len(mat[0]) {
			if mat[i][j] == "@" && accessible(mat, i, j) {
				locations = append(locations, [2]int{i, j})
			}
		}
	}
	return locations
}

func part2(mat [][]string) int {
	before := countRolls(mat)
	queue := part1(mat)

	for len(queue) > 0 {
		row, col := queue[0][0], queue[0][1]
		queue = queue[1:]

		if mat[row][col] != "@" {
			continue
		}

		mat[row][col] = "."

		for _, neighbor := range neighbors {
			nr, nc := row+neighbor[0], col+neighbor[1]
			if inBounds(mat, nr, nc) && mat[nr][nc] == "@" && accessible(mat, nr, nc) {
				queue = append(queue, [2]int{nr, nc})
			}
		}
	}

	after := countRolls(mat)
	return before - after
}

func main() {
	lines := utils.ReadLines("input.txt")
	mat := make([][]string, len(lines))
	for i, line := range lines {
		mat[i] = strings.Split(line, "")
	}

	locations := part1(mat)
	res1 := len(locations)
	fmt.Printf("Part 1: %d\n", res1)

	res2 := part2(mat)
	fmt.Printf("Part 2: %d\n", res2)
}
