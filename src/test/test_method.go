package testing

import (
	"fmt"
	"math"
)

type Point struct{ x, y float64 }

func Distance(p, q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func TestDis() {
	p1 := Point{1, 2}
	p2 := Point{3, 4}
	fmt.Println(Distance(p1, p2))
	fmt.Println(p1.Distance(p2))
}

func TestDisPath() {
	p := Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}
	fmt.Println(p.Distance())
}

//------------------------------------------------

func (p *Point) ScaleBy(factor float64) {
	p.x *= factor
	p.y *= factor
}

func (p Point) Add() float64 {
	return p.x + p.y
}

func TestPointer() {
	p := Point{1, 2}
	p.ScaleBy(10)
	fmt.Println(p)

	p1 := Point{3, 4}
	p1ptr := &p1
	p1.ScaleBy(100)
	fmt.Println(*p1ptr)

	fmt.Println(p1ptr.Add())
}
