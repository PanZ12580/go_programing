package main

import (
	"fmt"
	"go_programing/chapter8/ptest"
	"time"
)
var s = []string{"1", "2", "3"}

func main() {
/*	go spinner(100 * time.Millisecond)
	res := fib(45)
	fmt.Println(res)*/
	s := []int{5}
	ptest.PrintVar(s)
	fmt.Println(s)
}

func test(count int) {
	if count < 0 {
		return
	}
	fmt.Printf("%p\n", &s)
	test(count - 1)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `/\|/` {
			fmt.Printf("%c", r)
		}
		time.Sleep(delay)
	}
}

func fib(n int) int {
	if n < 2 {
		return 1
	}
	return fib(n - 1) + fib(n - 2)
}
