package main

import (
	"2024/utils"
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

func part2(lines []string) int {
	text := strings.Join(lines, "")
	r := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\)`)
	tokens := r.FindAllString(text, -1)
	total := 0
	mode := true // true=do, false=don't
	for _, token := range tokens {
		switch token {
		case "do()":
			mode = true
		case "don't()":
			mode = false
		default:
			if mode {
				token = strings.TrimPrefix(token, "mul(")
				token = strings.TrimSuffix(token, ")")
				parts := strings.Split(token, ",")
				num1, _ := strconv.Atoi(parts[0])
				num2, _ := strconv.Atoi(parts[1])
				total += num1 * num2
			}
		}
	}
	return total
}

func main() {
	lines := utils.ReadLines("input.txt")
	res1 := part1(lines)
	fmt.Printf("Part 1: %d\n", res1)

	res2 := part2(lines)
	fmt.Printf("Part 2: %d\n", res2)
}
