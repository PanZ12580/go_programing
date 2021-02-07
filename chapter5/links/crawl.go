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

var flag bool = true

func Bfs(f func(string) []string, urls []string) {
	visited := make(map[string]bool)
	for len(urls) > 0 {
		items := urls
		urls = nil
		for _, item := range items {
			if !visited[item] {
				visited[item] = true
				urls = append(urls, f(item)...)
			}
		}
	}
}

func Crawl(url string) []string {
	fmt.Println(url)
	links, err := Extract(url)
	if flag {
		saveHTML(links)
	}
	flag = false
	if err != nil {
		log.Fatal(err)
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
