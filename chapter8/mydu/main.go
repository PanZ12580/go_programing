package main

import (
	"flag"
	"fmt"
	"go_programing/chapter8/mydu/du"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()
	root := flag.Arg(0)
	fmt.Println(root)
	dirMap := make(map[string]chan int64)
	printMap := make(map[string]int64)

	var wg sync.WaitGroup

	if len(root) == 0 {
		root = "."
	}
	for _, entry := range du.Dirents(root) {
		if entry.IsDir() {
			wg.Add(1)
			fileSizes := make(chan int64)
			dirMap[entry.Name()] = fileSizes
			printMap[entry.Name()] = 0

			go func(entry os.FileInfo) {
				du.WalkDir(filepath.Join(root, entry.Name()), &wg, fileSizes)
			}(entry)
		}
	}

	go func() {
		wg.Wait()
		for _, c := range dirMap {
			close(c)
		}
	}()

	tick := make(<-chan time.Time)
	if *verbose {
		tick = time.Tick(1 * time.Second)
	}

loop:
	for {
		select {
		case <-tick:
			for k, v := range printMap {
				printDirUsage(k, v)
			}
			fmt.Println("===============================================")
		default:
			closeCount := 0
			for k, c := range dirMap {
				select {
				case size, ok := <-c:
					if !ok {
						closeCount++
					} else {
						printMap[k] += size
					}
				default:
					break
				}
			}
			if closeCount == len(dirMap) {
				break loop
			}
		}
	}
	var nfiles, nbytes int64
	for k, v := range printMap {
		nfiles++
		nbytes += v
		printDirUsage(k, v)
	}
	fmt.Printf("directory count: %d, total size: %.1fGB\n", nfiles, float64(nbytes)/1e9)
}

func printDirUsage(dirName string, size int64) {
	fmt.Printf("%s: %.2fMB\n", dirName, float64(size) / 1e6)
}
