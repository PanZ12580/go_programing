package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	launch3()
}

func launch1() {
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown >= 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	fmt.Println("Rocket launch!")
}

func launch2() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	select {
	case <-time.After(10 * time.Second):
		fmt.Println("Rocket launch!")
	case <-abort:
		fmt.Println("Terminating launch!")
	}
}

func launch3() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for countdown := 10; countdown > 0; countdown-- {
		select {
		case <-abort:
			fmt.Println("Terminating launch!")
			return
		case <-ticker.C:
			fmt.Println(countdown)
		}
	}
	fmt.Println("Rocket launch!")
}
