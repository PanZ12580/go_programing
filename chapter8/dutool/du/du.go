/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package du

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

var sema = make(chan struct{}, 20)

func WalkDir(directory string, fileSizes chan<- int64, group *sync.WaitGroup, done chan struct{}) {
	if cancelled(done) {
		return
	}
	defer group.Done()
	for _, entry := range dirents(directory, done) {
		if entry.IsDir() {
			group.Add(1)
			subDir := filepath.Join(directory, entry.Name())
			WalkDir(subDir, fileSizes, group, done)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(directory string, done chan struct{}) []os.FileInfo {
	select {
	case sema <- struct{}{}:
	case <-done:
		return nil
	}
	defer func() {
		<-sema
	}()
	entries, err := ioutil.ReadDir(directory)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}

func cancelled(done chan struct{}) bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}



