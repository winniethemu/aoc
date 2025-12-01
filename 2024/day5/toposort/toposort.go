package toposort

import (
	"fmt"
	"strconv"
	"strings"

	"aoc/utils"
)

// check if `pages` is a subsequence of `rule`
func validate(pages []string, rule []string) bool {
	i := 0 // current index in rule
	j := 0 // current index in pages

	for ; i < len(rule); i++ {
		if rule[i] == pages[j] {
			j++
			if j == len(pages) {
				return true
			}
		}
	}

	return false
}

func toposort(edges [][]string) []string {
	res := make([]string, 0)
	vertices := utils.NewSet[string]()
	indegree := make(map[string]utils.Set[string])

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		if indegree[v] == nil {
			indegree[v] = utils.NewSet[string]()
		}
		indegree[v].Add(u)
		vertices.Add(u)
		vertices.Add(v)
	}

	Q := make([]string, 0)
	for vertex := range vertices {
		_, found := indegree[vertex]
		if !found {
			Q = append(Q, vertex)
		}
	}

	for len(Q) > 0 {
		node := Q[0]
		Q = Q[1:]
		res = append(res, node)
		for k, v := range indegree {
			if v.Contains(node) {
				v.Remove(node)
				if v.Size() == 0 {
					Q = append(Q, k)
					delete(indegree, k)
				}
			}
		}
	}

	return res
}

func Run() {
	lines := utils.ReadLines("example.txt")

	idx := 0
	edges := make([][]string, 0)
	for _, line := range lines {
		idx++
		if strings.Contains(line, "|") {
			pair := strings.Split(line, "|")
			edges = append(edges, pair)
		} else {
			break
		}
	}
	rule := toposort(edges)

	total1 := 0
	for ; idx < len(lines); idx++ {
		pages := strings.Split(lines[idx], ",")
		if validate(pages, rule) {
			midpage, _ := strconv.Atoi(pages[len(pages)/2])
			total1 += midpage
		}
	}

	fmt.Println(total1)
}
