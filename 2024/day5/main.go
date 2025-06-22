package main

import (
	"2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func validate(pages []string, rules map[string]utils.Set[string]) bool {
	for i := 1; i < len(pages); i++ {
		prev, curr := pages[i-1], pages[i]
		successors := rules[prev]
		if !successors.Contains(curr) {
			return false
		}
	}
	return true
}

func main() {
	rules := make(map[string]utils.Set[string])

	idx := 0
	lines := utils.ReadLines("input.txt")
	for _, line := range lines {
		idx++
		if strings.Contains(line, "|") {
			pair := strings.Split(line, "|")
			p1, p2 := pair[0], pair[1]
			_, found := rules[p1]
			if !found {
				rules[p1] = utils.NewSet[string]()
			}
			rules[p1].Add(p2)
		} else {
			break
		}
	}

	total := 0
	for ; idx < len(lines); idx++ {
		line := lines[idx]
		pages := strings.Split(line, ",")
		if validate(pages, rules) {
			midpage, _ := strconv.Atoi(pages[len(pages)/2])
			total += midpage
		}
	}

	fmt.Println(total)
}
