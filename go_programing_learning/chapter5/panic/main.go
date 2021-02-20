package main

import (
	"fmt"
	"os"
	"runtime"
)

type num struct {
	val int
}

var out = &num{}

func main() {
	/*	defer printStack()
		f(3)*/
	//fmt.Print(testReturn())
	defer func() {
		switch p := recover(); p {
		case nil:
		case num{}:
			fmt.Println(out)
		default:
			panic(p)
		}
	}()

	panicAndRecover()
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer => %d\n", x)
	f(x - 1)
}

func printStack() {
	var b [4096]byte
	n := runtime.Stack(b[:], false)
	os.Stdout.Write(b[:n])
}

func testReturn() (num int) {
	return num
}

func panicAndRecover() {
	out.val = 100
	panic(num{})
}
