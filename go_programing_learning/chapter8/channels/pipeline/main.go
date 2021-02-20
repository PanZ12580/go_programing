package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)
/*	go func() {
		for x := 0; x < 10; x++ {
			time.Sleep(100 * time.Millisecond)
			natruals <- x
		}
		close(natruals)
	}()

	go func() {*/
		/*for {
			x, ok := <- natruals
			if ok {
				squarer <- x * x
			} else {
				close(squarer)
				break
			}
		}*/
/*		for x := range natruals {
			squarer <- x * x
		}
		close(squarer)
	}()*/

	/*for {
		x, ok := <- squarer
		if ok {
			fmt.Println(x)
		} else {
			break
		}
	}*/
/*	for x := range squarer {
		fmt.Println(x)
	}*/
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

func counter(out  chan<- int) {
	for x := 0; x < 10; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}