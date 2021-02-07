package main

import (
	"go_programing/chapter5/links"
)

func main() {
	links.Bfs(links.Crawl, []string{"http://gopl.io"})
}
