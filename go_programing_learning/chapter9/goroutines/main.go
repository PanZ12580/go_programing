package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 1)
	var num = 0
	var wg sync.WaitGroup

	ch <- num

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			num := <-ch
			ch <- num + 1
		}()
	}

	wg.Wait()
	fmt.Println(<-ch)

	//pingPongTest()
}

func pingPongTest() {
	var ch = make(chan string)
	count := 0
	var done = make(chan struct{})

	go func() {
	loop:
		for {
			select {
			case <-done:
				break loop
			default:
				go func() {
					fmt.Println("ping")
					ch <- "pong"
				}()
				go func() {
					count++
					fmt.Println(<-ch)
				}()
			}
		}
	}()

	time.Sleep(1 * time.Second)
	close(done)
	fmt.Printf("count: %d\n", count)
}
