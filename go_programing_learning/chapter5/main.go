package main

import (
	"fmt"
	"go_programing/chapter5/links"
	"golang.org/x/net/html"
	"log"
	"strings"
)

var deep int

func main() {
	url := "http://gopl.io"
	f, err := links.Fetch(url)
	if err != nil {
		log.Fatal(err)
	}
	doc, err := html.Parse(strings.NewReader(f))
	if err != nil {
		log.Fatal(err)
	}
	for _, l := range visit([]string{}, doc) {
		fmt.Println(l)
	}
	links.ForEachNode(doc, startElement, endElement)
}

func visit(link []string, node *html.Node) (links []string) {
	if node.Type == html.ElementNode {
		link = append(link, node.Data)
	}

	for s := node.FirstChild; s != nil; s = s.NextSibling {
		link = visit(link, s)
	}

	return link
}

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", deep*2, "", n.Data)
		deep++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		deep--
		fmt.Printf("%*s<%s>\n", deep*2, "", n.Data)
	}
}
