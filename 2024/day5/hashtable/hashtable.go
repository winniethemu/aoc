package hashtable

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"aoc/utils"
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

func reorder(pages []string, rules map[string]utils.Set[string]) []string {
	slices.SortFunc(pages, func(a, b string) int {
		if rules[a].Contains(b) {
			return -1
		}
		if rules[b].Contains(a) {
			return 1
		}
		return 0
	})
	return pages
}

func Run() {
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

	total1 := 0
	total2 := 0
	for ; idx < len(lines); idx++ {
		line := lines[idx]
		pages := strings.Split(line, ",")
		if validate(pages, rules) {
			midpage, _ := strconv.Atoi(pages[len(pages)/2])
			total1 += midpage
		} else {
			pages = reorder(pages, rules)
			midpage, _ := strconv.Atoi(pages[len(pages)/2])
			total2 += midpage
		}
	}

	fmt.Println(total1, total2)
}
