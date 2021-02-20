package main

import (
	"flag"
	"fmt"
	"go_programing/chapter8/dutool/du"
	"os"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()
	fileSizes := make(chan int64)
	var wg sync.WaitGroup
	roots := flag.Args()
	var done = make(chan struct{})

	if len(roots) == 0 {
		roots = []string{"."}
	}


	for _, dir := range roots {
		wg.Add(1)
		go du.WalkDir(dir, fileSizes, &wg, done)
	}

	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	var nfiles, nbytes int64
	/*	for size := range fileSizes {
			nfiles++;
			nbytes += size
		}
		printDiskUsage(nfiles, nbytes)*/
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

loop:
	for {
		select {
		case <-done:
			for range fileSizes {}
			return
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files, %.1fGB\n", nfiles, float64(nbytes)/1e9)
}
