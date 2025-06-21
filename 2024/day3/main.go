package main

import (
	utils "2024"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func sum(line string) int {
	r := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	occurences := r.FindAllString(line, -1)
	total := 0
	for _, o := range occurences {
		o = strings.TrimPrefix(o, "mul(")
		o = strings.TrimSuffix(o, ")")
		parts := strings.Split(o, ",")
		num1, _ := strconv.Atoi(parts[0])
		num2, _ := strconv.Atoi(parts[1])
		total += num1 * num2
	}
	return total
}

func part1(lines []string) int {
	total := 0
	for _, line := range lines {
		total += sum(line)
	}
	return total
}

func main() {
	lines := utils.ReadLines("input.txt")
	res1 := part1(lines)
	fmt.Printf("Part 1: %d\n", res1)
}
