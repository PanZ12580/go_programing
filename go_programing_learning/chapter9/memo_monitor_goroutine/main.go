package main

import (
	"fmt"
	"go_programing/chapter9/memo_monitor_goroutine/cache"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	memo := cache.New(httpGet)
	urls := []string {
		"http://www.baidu.com",
		"http://www.panzvor.com",
		"https://www.hnu.edu.cn/",
		"http://www.baidu.com",
		"http://www.panzvor.com",
		"https://www.hnu.edu.cn/",
	}
	var n sync.WaitGroup

	for i, url := range urls {
		n.Add(1)
		go func(i int, url string) {
			start := time.Now()
			defer n.Done()
			value, err := memo.Get(url, make(chan struct{}))
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Printf("%d -> %s %s %d\n", i, url, time.Since(start), len(value.([]byte)))
		}(i, url)
	}
	n.Wait()
}

func httpGet(url string) (interface{}, error) {
	fmt.Println("request")
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	return ioutil.ReadAll(resp.Body)
}
