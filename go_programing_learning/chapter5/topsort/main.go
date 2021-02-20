package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus": {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems, computer organization"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, c := range topsort(prereqs) {
		fmt.Printf("%d. %s\n", i + 1, c)
	}
/*	for k, _ := range topSortMap(prereqs) {
		fmt.Printf("%s\n", k)
	}*/
	fmt.Printf("exist ring? %t\n", detectRing(prereqs))
}

func topsort(m map[string][]string) []string {
	visited := make(map[string]bool)
	res := make([]string, 0)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !visited[item] {
				visited[item] = true
				visitAll(m[item])
				res = append(res, item)
			}
		}
	}

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	visitAll(keys)

	return res
}

func topSortMap(m map[string][]string) map[string]bool {
	visited := make(map[string]bool)
	res := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !visited[item] {
				visited[item] = true
				visitAll(m[item])
				res[item] = true
			}
		}
	}

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	visitAll(keys)

	return res
}

func detectRing(m map[string][]string) bool {
	startSet := make(map[string]bool)
	visited := make(map[string]bool)

	var dfs func(items []string) bool
	dfs = func(items []string) bool {
		for _, item := range items {
			if startSet[item] {
				return true
			}
			if !visited[item] {
				visited[item] = true
				startSet[item] = true
				if dfs(m[item]) {
					return true
				}
				startSet[item] = false
			}
		}
		return false
	}

	var keys []string
	for k, _ := range m {
		keys = append(keys, k)
	}
	return dfs(keys)
}
