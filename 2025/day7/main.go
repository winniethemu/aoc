package main

import (
	"aoc/utils"
	"fmt"
	"slices"
	"strings"
)

func part1(manifold [][]string) int {
	res := 0
	n, m := len(manifold), len(manifold[0])
	src := slices.Index(manifold[0], "S")
	visited := make(map[string]bool)

	var traverse func(r, c int)
	traverse = func(row, col int) {
		if row < 0 || row > n-1 || col < 0 || col > m-1 {
			return
		}
		key := fmt.Sprintf("%d:%d", row, col)
		if _, exists := visited[key]; exists {
			return
		}
		visited[key] = true
		if manifold[row][col] == "^" {
			res++
			traverse(row, col-1)
			traverse(row, col+1)
		} else {
			traverse(row+1, col)
		}
	}
	traverse(0, src)

	return res
}

func part2() int {
	res := 0
	return res
}

func parse(lines []string) [][]string {
	manifold := make([][]string, 0)
	for _, line := range lines {
		cells := strings.Split(line, "")
		manifold = append(manifold, cells)
	}
	return manifold
}

func main() {
	lines := utils.ReadLines("example.txt")
	manifold := parse(lines)

	res1 := part1(manifold)
	fmt.Printf("Part 1: %d\n", res1)

	res2 := part2()
	fmt.Printf("Part 2: %d\n", res2)
}
