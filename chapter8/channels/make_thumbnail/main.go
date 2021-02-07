package main

import (
	"gopl.io/ch8/thumbnail"
	"log"
	"os"
	"sync"
)

func main() {
	
}

func makeThumbnails1(filenames []string) error {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go thumbnail.ImageFile(f)
	}
}

func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			thumbnail.ImageFile(f)
			ch <- struct{}{}
		}(f)
	}

	for range filenames {
		<- ch
	}
}

func makeThumbnails4(filenames []string) error {
	errors := make(chan error)
	for _, f := range  filenames {
		go func(f string) {
			_, err := thumbnail.ImageFile(f)
			if err != nil {
				errors <- err
			}
		}(f)
	}

	for range filenames {
		if err := <- errors; err != nil {
			return err
		}
	}
	return nil
}

func makeThumbnails5(filenames []string) (thumbfile []string, err error) {
	type item struct {
		thumbfile string
		err error
	}

	ch := make(chan item, len(filenames))
	for _, f := range  filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <- ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfile = append(thumbfile, it.thumbfile)
	}
	return
}

func makeThumbnails6(filenames []string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for _, f := range filenames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			thumbfile, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumbfile)
			sizes <- info.Size()
		}(f)
	}

	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for s := range sizes {
		total += s
	}
	return total
}
