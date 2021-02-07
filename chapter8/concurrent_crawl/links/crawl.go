/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package links

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	flag = true
	n = 1
)

func Bfs(f func(string) []string, urls []string) {
	worklist := make(chan []string)
	unVisitedLink := make(chan string)
	depth := 0

	go func() {
		worklist <- urls
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for l := range unVisitedLink {
				links := f(l)
				go func() {
					if depth < 3 {
						worklist <- links
						depth++
					}
				}()
			}
		}()
	}

	visited := make(map[string]bool)
	for ; n > 0; n-- {
		for list := range worklist {
			for _, link := range list {
				if !visited[link] {
					unVisitedLink <- link
					visited[link] = true
					n++
				}
			}
		}
	}
}

func Crawl(url string) []string {
	fmt.Println(url)
	links, err := Extract(url)
/*	if flag {
		go saveHTML(links)
	}
	flag = false*/
	if err != nil {
		log.Println(err)
		return nil
	}
	return links
}

func saveHTML(links []string) {
	for i, link := range links {
		resp, err := http.Get(link)
		if err != nil {
			log.Fatal(err)
		}
		f, err := os.Create("./" + strconv.Itoa(i) + ".html")
		if err != nil {
			log.Fatal(err)
		}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		f.Write(b)
		f.Close()
	}
}
