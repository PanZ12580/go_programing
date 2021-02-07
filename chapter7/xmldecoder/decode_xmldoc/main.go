package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"go_programing/chapter7/xmldecoder/fetch"
	"io"
	"log"
	"strings"
)

type Node interface{}

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func main() {
	s := fetch.Fetch("http://www.w3.org/TR/2006/REC-xml11-20060816")
	decoder := xml.NewDecoder(strings.NewReader(s))
	var nodeList []Node
	var stack []*Element
	for {
		tk, err := decoder.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		switch tk := tk.(type) {
		case xml.StartElement:
			var node Element
			node.Type = tk.Name
			node.Attr = tk.Attr
			if len(stack) > 0 {
				tmp := stack[len(stack) - 1]
				(*tmp).Children = append((*tmp).Children, &node)
			} else {
				nodeList = append(nodeList, &node)
			}
			stack = append(stack, &node)
		case xml.EndElement:
			stack = stack[:len(stack) - 1]
		case xml.CharData:
			var dataNode CharData
			dataNode = CharData(tk)
			if len(stack) > 0 {
				n := stack[len(stack) - 1]
				(*n).Children = append((*n).Children, &dataNode)
			} else {
				nodeList = append(nodeList, &dataNode)
			}
		}
	}

	for _, e := range nodeList {
		print(e, 0)
	}
}

func print(n Node, level int) {
	switch n := n.(type) {
	case *CharData:
		fmt.Println(*n)
	case *Element:
		var buf bytes.Buffer
		for _, a := range n.Attr {
			buf.WriteRune(' ')
			buf.WriteString(fmt.Sprintf("%s=%q", a.Name.Local, a.Value))
		}
		fmt.Printf("<%s%s>\n", n.Type.Local, buf.String())
		for _, c := range n.Children {
			for i := level + 1; i > 0; i-- {
				fmt.Print("\t")
			}
			print(c, level + 1)
		}
		for i := level; i > 0; i-- {
			fmt.Print("\t")
		}
		fmt.Printf("</%s>\n", n.Type.Local)
	}
}
