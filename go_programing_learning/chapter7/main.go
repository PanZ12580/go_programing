/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"strings"
	"unicode"
)

type Reader struct {
	s        string
	i        int64 // current reading index
	prevRune int   // index of previous rune; or < 0
}

func (r *Reader) Read(b []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	r.prevRune = -1
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return
}

type Counter struct {
	count, rows int
	writer io.Writer
}

type WordCounter int

func (wc *WordCounter) Write(p []byte) (n int, err error) {
	bytes.TrimSpace(p)
	var flag  = true
	for _, b := range p {
		if unicode.IsSpace(rune(b)) {
			flag = true
			continue
		} else if flag {
			*wc++
			flag = false
		}
	}
	return int(*wc), nil
}

func (c *Counter) Write(p []byte) (n int, err error) {
	var flag = true
	if len(p) > 0 {
		c.rows += 1
	}
	for _, r := range p {
		if r == '\n' {
			c.rows++
		}
		if unicode.IsSpace(rune(r)) {
			flag = true
			continue
		} else if flag {
			c.count++
			flag = false
		}
	}
	return c.count, nil
}

func main() {
/*	ad, tk, err := bufio.ScanWords([]byte("hello world  hah"), false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ad)
	for _, r := range tk {
		fmt.Printf("%c\t", r)
	}*/
	words := `hello world, 
hello goland hahahha   are u ok?`
	var c Counter
	var wc WordCounter
	fmt.Fprint(&c, words)
	fmt.Println(c.count, ", ", c.rows)
	fmt.Println(wc.Write([]byte(words)))
	s := "<p><a href='http://www.baidu.com'>click</a>, hello world!</p>"
	doc, err := html.Parse(NewReader(s))
	if err != nil {
		log.Fatal(err)
	}
	for s := doc.FirstChild; s != nil; s = s.NextSibling {
		fmt.Println(s.Data)
	}

	fmt.Println("=====================================")
	reader := LimitReader(strings.NewReader(s), 10)
	fmt.Println(reader.Read([]byte(s)))

}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var c Counter
	c.writer = w
	res := int64(c.count)
	return &c, &res
}

func NewReader(s string) io.Reader {
	return &Reader{s, 0, -1}
}

type LimitRead struct {
	r io.Reader
	s string
	n int64
}

func (lr *LimitRead) Read(b []byte) (n int, err error) {
	n, err = lr.r.Read(b[:lr.n])
	lr.s = string(b[:lr.n])
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitRead{r, "", n}
}
