package main

import (
	"aoc/utils"
	"container/heap"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Position struct {
	ID int // for grouping
	X  int
	Y  int
	Z  int
}

type JunctionBoxPair struct {
	Priority int // Euclidian distance squared
	Pos1     Position
	Pos2     Position
}

type PriorityQueue []*JunctionBoxPair

func (pq *PriorityQueue) Len() int {
	return len(*pq)
}

// Use a max-heap to kick out large values when heap is full
func (pq *PriorityQueue) Less(i, j int) bool {
	return (*pq)[i].Priority > (*pq)[j].Priority
}

func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PriorityQueue) Push(item any) {
	pair := item.(*JunctionBoxPair)
	*pq = append(*pq, pair)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	old[n-1] = nil
	return *item
}

type DSU struct {
	parents []int
	sizes   []int // size of each component (only valid for roots)
}

func NewDSU(size int) *DSU {
	parents := make([]int, size)
	sizes := make([]int, size)
	for i := range size {
		parents[i] = i
		sizes[i] = 1 // each element starts as its own component of size 1
	}

	return &DSU{
		parents: parents,
		sizes:   sizes,
	}
}

func (dsu *DSU) Union(x, y int) {
	rx := dsu.Find(x)
	ry := dsu.Find(y)
	if rx == ry {
		return
	}
	// Union by size: attach smaller tree to larger tree
	if dsu.sizes[rx] < dsu.sizes[ry] {
		rx, ry = ry, rx
	}
	dsu.parents[ry] = rx
	dsu.sizes[rx] += dsu.sizes[ry]
}

func (dsu *DSU) Find(x int) int {
	if dsu.parents[x] == x {
		return x
	}
	// Path compression
	dsu.parents[x] = dsu.Find(dsu.parents[x])
	return dsu.parents[x]
}

func part1(positions []Position, nConn int) int {
	pq := make(PriorityQueue, 0)
	for i := range len(positions) {
		for j := i + 1; j < len(positions); j++ {
			pair := JunctionBoxPair{
				Priority: dist2(positions[i], positions[j]),
				Pos1:     positions[i],
				Pos2:     positions[j],
			}
			heap.Push(&pq, &pair)
			if pq.Len() > nConn {
				heap.Pop(&pq)
			}
		}
	}
	// Now we end up with a max-heap of the smallest values

	pairs := make([]JunctionBoxPair, nConn)
	for i := nConn - 1; i >= 0; i-- {
		pairs[i] = heap.Pop(&pq).(JunctionBoxPair)
	}

	// Union-Find
	dsu := NewDSU(len(positions))
	for _, pair := range pairs {
		pos1, pos2 := pair.Pos1, pair.Pos2
		dsu.Union(pos1.ID, pos2.ID)
	}

	res := 1

	roots := make(map[int]int)
	for i := range dsu.parents {
		r := dsu.Find(i)
		roots[r] = dsu.sizes[r]
	}

	sizes := make([]int, 0)
	for _, v := range roots {
		sizes = append(sizes, v)
	}
	slices.SortFunc(sizes, func(a, b int) int {
		return b - a
	})

	// Count the three largest circuits
	for i := range 3 {
		res *= sizes[i]
	}

	return res
}

func part2(positions []Position) int {
	pq := make(PriorityQueue, 0)
	for i := range len(positions) {
		for j := i + 1; j < len(positions); j++ {
			pair := JunctionBoxPair{
				Priority: dist2(positions[i], positions[j]),
				Pos1:     positions[i],
				Pos2:     positions[j],
			}
			heap.Push(&pq, &pair)
		}
	}

	pairs := make([]JunctionBoxPair, pq.Len())
	for i := pq.Len() - 1; i >= 0; i-- {
		pairs[i] = heap.Pop(&pq).(JunctionBoxPair)
	}

	// Union-Find
	dsu := NewDSU(len(positions))
	for _, pair := range pairs {
		pos1, pos2 := pair.Pos1, pair.Pos2
		dsu.Union(pos1.ID, pos2.ID)
		if slices.Index(dsu.sizes, len(positions)) > -1 {
			return pos1.X * pos2.X
		}
	}

	panic("should not get here")
}

// Distance squared
func dist2(pos1, pos2 Position) int {
	dx := pos1.X - pos2.X
	dy := pos1.Y - pos2.Y
	dz := pos1.Z - pos2.Z
	return dx*dx + dy*dy + dz*dz
}

func parse(lines []string) []Position {
	positions := make([]Position, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		positions[i] = Position{ID: i, X: x, Y: y, Z: z}
	}
	return positions
}

func main() {
	lines := utils.ReadLines("input.txt")

	res1 := part1(parse(lines), 1000)
	fmt.Printf("Part 1: %d\n", res1)

	res2 := part2(parse(lines))
	fmt.Printf("Part 2: %d\n", res2)
}
