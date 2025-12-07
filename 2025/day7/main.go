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

func part2(manifold [][]string) int {
	n, m := len(manifold), len(manifold[0])

	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, m)
	}

	// start position has 1 timeline
	src := slices.Index(manifold[0], "S")
	dp[0][src] = 1

	for r := range n {
		// first pass: handle splitters
		for c := range m {
			if manifold[r][c] == "^" && dp[r][c] > 0 {
				// split timelines to left and right on the same row
				if c > 0 {
					dp[r][c-1] += dp[r][c]
				}
				if c < m-1 {
					dp[r][c+1] += dp[r][c]
				}
				// clear the splitter cell - timelines have been split
				dp[r][c] = 0
			}
		}

		// second pass: propagate all timelines downward
		if r < n-1 {
			for c := range m {
				if dp[r][c] > 0 {
					dp[r+1][c] += dp[r][c]
				}
			}
		}
	}

	res := 0
	for i := range m {
		res += dp[n-1][i]
	}

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
	lines := utils.ReadLines("input.txt")
	manifold := parse(lines)

	res1 := part1(manifold)
	fmt.Printf("Part 1: %d\n", res1)

	res2 := part2(manifold)
	fmt.Printf("Part 2: %d\n", res2)
}
