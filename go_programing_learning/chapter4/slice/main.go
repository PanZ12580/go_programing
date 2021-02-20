package main

import (
	"fmt"
	"go_programing/chapter4/slice/function"
)

func main() {
//  test array reverse
	arr := [5]int{1, 2, 3, 4, 5}
	function.Reverse(&arr)
	fmt.Println(arr)
//  test function rotate
	s := []int{0, 1, 2, 3, 4, 5}
	res := function.Rotate(s, 2)
	fmt.Println(res)
//	test eliminate repeat string
	str := []string{"ag", "ag", "bbbb", "t", "b", "a", "l", "l"}
	function.EliminateRepeat(&str)
	fmt.Println(str)
//	test eliminate repeat space byte
	b := []byte{'1', 'a', '\t', ' ', ' ', '\t', 'r', 'h', 'z'}
	fmt.Println(len(b))
	function.EliminateRepeatSpace(&b)
	fmt.Printf("%c\tlen(b) = %d\n", b, len(b))
//	test reverse byte array
	b2 := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i'}
	function.ReverseByteArr(&b2)
	fmt.Printf("%c\n", b2)
}
