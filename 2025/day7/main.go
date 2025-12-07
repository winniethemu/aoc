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
	res := 0
	n, m := len(manifold), len(manifold[0])

	dp := make([][]int, 0)
	for i := range n {
		dp = append(dp, make([]int, m))
		for j := range m {
			dp[i][j] = 0
		}
	}

	r0, c0 := 0, slices.Index(manifold[0], "S")

	queue := make([][2]int, 0)
	queue = append(queue, [2]int{r0, c0})
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		r, c := node[0], node[1]
		if r < 0 || r > n-1 || c < 0 || c > m-1 {
			continue
		}
		switch manifold[r][c] {
		case "^":
			queue = append(queue, [2]int{r, c - 1}, [2]int{r, c + 1})
		case "S":
			dp[r][c] = 1
			queue = append(queue, [2]int{r + 1, c})
		case ".":
			dp[r][c] = dp[r-1][c]
			if c > 0 && manifold[r][c-1] == "^" {
				dp[r][c] += dp[r-1][c-1]
			}
			if c < m-1 && manifold[r][c+1] == "^" {
				dp[r][c] += dp[r-1][c+1]
			}
			queue = append(queue, [2]int{r + 1, c})
		}
	}

	for i := range n {
		fmt.Println(dp[i])
	}

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
	lines := utils.ReadLines("example.txt")
	manifold := parse(lines)

	res1 := part1(manifold)
	fmt.Printf("Part 1: %d\n", res1)

	res2 := part2(manifold)
	fmt.Printf("Part 2: %d\n", res2)
}
