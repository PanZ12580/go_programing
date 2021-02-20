/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package du

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func WalkDir(dir string, group *sync.WaitGroup, fileSizes chan int64) {
	defer group.Done()
	for _, entry := range Dirents(dir) {
		if entry.IsDir() {
			group.Add(1)
			go func(entry os.FileInfo) {
				WalkDir(filepath.Join(dir, entry.Name()), group, fileSizes)
			}(entry)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

var sema = make(chan struct{}, 20)

func Dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Printf("du: %v\n", err)
		return nil
	}
	return entries
}
