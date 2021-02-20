package main

import (
	"go_programing/chapter8/concurrent_crawl/links"
)

func main() {
	links.Bfs(links.Crawl, []string{"http://gopl.io"})
}
