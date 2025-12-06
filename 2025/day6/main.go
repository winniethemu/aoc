package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func parse(str string) []string {
	res := make([]string, 0)
	items := strings.Split(str, " ")
	for _, item := range items {
		if item != "" {
			res = append(res, item)
		}
	}
	return res
}

func convert(ss []string) []int {
	res := make([]int, 0)
	for i := range len(ss) {
		val, _ := strconv.Atoi(ss[i])
		res = append(res, val)
	}
	return res
}

func part1(dataInput []string, opsInput string) int {
	res := 0

	data := make([][]int, 0)
	for _, dat := range dataInput {
		nums := convert(parse(dat))
		data = append(data, nums)
	}

	ops := parse(opsInput)

	for j := 0; j < len(ops); j++ {
		op := ops[j]

		total := 0
		if op == "*" {
			total = 1
		}

		for i := 0; i < len(data); i++ {
			if op == "+" {
				total += data[i][j]
			} else {
				total *= data[i][j]
			}
		}

		res += total
	}

	return res
}

func resetSubtotal(op string) int {
	if op == "*" {
		return 1
	}
	return 0
}

func part2(data []string, opsline string) int {
	res := 0
	ops := parse(opsline)

	columns := len(data[0])
	problem := 0
	subtotal := resetSubtotal(ops[problem])

	for j := range columns {
		item := ""
		for i := range data {
			item += string(data[i][j])
		}

		if strings.TrimSpace(item) == "" {
			res += subtotal
			problem++
			subtotal = resetSubtotal(ops[problem])
		} else {
			num, _ := strconv.Atoi(strings.TrimSpace(item))
			if ops[problem] == "+" {
				subtotal += num
			} else {
				subtotal *= num
			}
		}
	}

	res += subtotal
	return res
}

func main() {
	lines := utils.ReadLines("input.txt")

	data := lines[:len(lines)-1]
	ops := lines[len(lines)-1]

	res1 := part1(data, ops)
	fmt.Printf("Part 1: %d\n", res1)

	res2 := part2(data, ops)
	fmt.Printf("Part 2: %d\n", res2)
}
