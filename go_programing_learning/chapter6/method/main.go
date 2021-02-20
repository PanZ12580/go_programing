package main

import (
	"fmt"
	"math"
	"sync"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Sqrt((p.X - q.X) * (p.X - q.X) + (p.Y - q.Y) * (p.Y - q.Y))
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p *Point) Add(offset Point) Point {
	return Point{p.X + offset.X, p.Y + offset.Y}
}

func (p *Point) Sub(offset Point) Point {
	return Point{p.X - offset.X, p.Y - offset.Y}
}

type Path []Point

func (path *Path) TranslateBy(p Point, add bool) {
	var op func(p *Point, q Point) Point
	if add {
		op = (*Point).Add
	} else {
		op = (*Point).Sub
	}

	for i := range *path {
		(*path)[i] = op(&(*path)[i], p)
	}
}

type cache struct {
	sync.Mutex
	mapping map[string]string
}

func (c *cache) put(k string, v string) {
	c.Lock()
	c.mapping[k] = v
	defer c.Unlock()
}

func (c *cache) get(k string) string {
	c.Lock()
	v := c.mapping[k]
	c.Unlock()
	return v
}

func main() {
/*	p := Point{3, 4}
	fmt.Println(p.Distance(Point{0, 0}))

	s := (*Point).ScaleBy
	d := Point.Distance
	s(&p, 2.5)
	fmt.Println(d(p, Point{0, 0}))
	fmt.Printf("p{X: %.2f, Y: %.2f}\n", p.X, p.Y)*/
/*	c := &cache{
		mapping: make(map[string]string),
	}
	c.put("k1", "v1")
	c.put("k1", "v1111")
	fmt.Println(c.get("k1"))*/

	p := Path{
		{0, 0},
		{1, 1},
		{2, 4},
		{3, 6},
	}
	p.TranslateBy(Point{10, 20}, false)
	for _, q := range p{
		fmt.Printf("Point{X: %.2f, Y: %.2f}\n", q.X, q.Y)
	}
}
