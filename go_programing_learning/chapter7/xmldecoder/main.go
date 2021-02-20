package main

import (
	"encoding/xml"
	"fmt"
	"go_programing/chapter7/xmldecoder/fetch"
	"io"
	"os"
	"strings"
)

func main() {
	s := fetch.Fetch("http://www.w3.org/TR/2006/REC-xml11-20060816")
	decoder := xml.NewDecoder(strings.NewReader(s))
	target := []string{"div", "div", "h2"}
	targetCls := []string{"copyright"}
	var stack []string
	var classStack []string
	for {
		tk, err := decoder.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tk := tk.(type) {
		case xml.StartElement:
			stack = append(stack, tk.Name.Local)
			for _, a := range tk.Attr {
				if a.Name.Local == "class" {
					classStack = append(classStack, a.Value)
				}
			}
		case xml.EndElement:
			if len(stack) > 0 {
				stack = stack[:len(stack) - 1]
			}
			if len(classStack) > 0 {
				classStack = classStack[:len(classStack) - 1]
			}
		case xml.CharData:
			if len(classStack) > 0 && containsAll(strings.Split(classStack[len(classStack) - 1], " "), targetCls) {
				fmt.Printf("%s %s\n", classStack[len(classStack) - 1], tk)
				classStack = classStack[:len(classStack) - 1]
			}
			if containsAll(stack, target) {
				fmt.Printf("%s %s\n", strings.Join(stack, " "), tk)
			}
		}
	}
}

func containsAll(x, y []string) bool {
	for len(x) >= len(y) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
