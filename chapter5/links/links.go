package links

import (
	"golang.org/x/net/html"
	"strings"
)

func Extract(url string) ([]string, error) {
	f, err := Fetch(url)
	if err != nil {
		return nil, err
	}
	doc, err := html.Parse(strings.NewReader(f))
	if err != nil {
		return nil, err
	}
	var link []string
	var visitAll func(n *html.Node)
	visitAll = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					link = append(link, a.Val)
					break
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visitAll(c)
		}
	}

	ForEachNode(doc, visitAll, nil)

	return link, nil
}

// forEachNode针对每个结点x，都会调用pre(x)和post(x)。
// pre和post都是可选的。
// 遍历孩子结点之前，pre被调用
// 遍历孩子结点之后，post被调用
func ForEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ForEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
