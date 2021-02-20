package main

import (
	"fmt"
	"sort"
)

type seq []byte

func (s *seq) Len() int {
	return len(*s)
}

func (s *seq) Less(i, j int) bool {
	return (*s)[i] < (*s)[j]
}

func (s *seq) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func main() {
	s := "abcdcba"
	test := seq(s)
	fmt.Println(IsPalindrome(&test))
}

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len() - 1; i < j; i, j = i + 1, j - 1 {
		if !s.Less(i, j) && !s.Less(j, i) {
			continue
		} else {
			return false
		}
	}
	return true
}
