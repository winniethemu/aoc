package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

type Point [2]int
type Vec2 struct {
	start Point
	end   Point
}

var reds []Point

func intersecting(v1, v2 Vec2) bool {
	return false
}

func getBoundingBox(polygon []Point) (Point, Point) {
	minX, minY := polygon[0][0], polygon[0][1]
	maxX, maxY := polygon[0][0], polygon[0][1]
	for _, point := range polygon {
		minX = min(minX, point[0])
		minY = min(minY, point[1])
		maxX = max(maxX, point[0])
		maxY = max(maxY, point[1])
	}
	return Point{minX, minY}, Point{maxX, maxY}
}

// ray casting
func inside(point Point, polygon []Point) bool {
	// points outside of bounding box are def not in the polygon
	topleft, bottomright := getBoundingBox(polygon)
	if point[0] < topleft[0] || point[0] > bottomright[0] || point[1] < topleft[1] || point[1] > bottomright[1] {
		return false
	}

	sides := make([]Vec2, 0, len(polygon))
	for i := 1; i < len(polygon); i++ {
		sides = append(sides, Vec2{start: polygon[i-1], end: polygon[i]})
	}

	ray := Vec2{start: point, end: Point{0, 0}} // assuming (0, 0) is outside of polygon
	intersections := 0
	for _, side := range sides {
		if intersecting(ray, side) {
			intersections++
		}
	}
	return intersections%2 == 1
}

// check if the rectangle specified by the given two points are fully contained
// by the polygon
func valid(pos1, pos2 Point) bool {
	// pos1, pos2 are red tiles and therefore must be inside the polygon
	// check the other two vertices
	p1 := Point{pos1[0], pos2[1]}
	p2 := Point{pos2[0], pos1[1]}

	return inside(p1, reds) && inside(p2, reds)
}

func area(pos1, pos2 Point) int {
	return (utils.AbsInt(pos1[0]-pos2[0]) + 1) * (utils.AbsInt(pos1[1]-pos2[1]) + 1)
}

func solve(validator func(p1, p2 Point) bool) int {
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

func parse(lines []string) []Point {
	res := make([]Point, len(lines))
	for i, line := range lines {
		xy := strings.Split(line, ",")
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])
		res[i] = Point{x, y}
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
