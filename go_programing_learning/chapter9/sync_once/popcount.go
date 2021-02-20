package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			weight := hammingWeight2(i * 2 + i * 24)
			fmt.Println(weight)
		}()
	}
	wg.Wait()
}

var (
	pc [256]byte
	isInit bool
	mu sync.RWMutex
	once sync.Once
)

func initP() {
	fmt.Println("init pc")
	for i := 0; i < 256; i++ {
		pc[i] = pc[i/2] + byte(i & 1)
	}
	isInit = true
}

func hammingWeight(n int) int {
	/*pc := func() (pc [256]byte) {
		for i := 0; i < 256; i++ {
			pc[i] = pc[i/2] + byte(i & 1)
		}
		return
	}()*/
	mu.RLock()
	if isInit {
		res := 0
		for i := 0; i < 8; i++ {
			res += int(pc[byte(n >> (8 * i))])
		}
		mu.RUnlock()
		return res
	}
	mu.RUnlock()

	mu.Lock()
	defer mu.Unlock()
	if !isInit {
		initP()
	}
	res := 0
	for i := 0; i < 8; i++ {
		res += int(pc[byte(n >> (8 * i))])
	}
	return res
}

func hammingWeight2(n int) int {
	once.Do(initP)
	res := 0
	for i := 0; i < 8; i++ {
		res += int(pc[byte(n >> (8 * i))])
	}
	return res
}
