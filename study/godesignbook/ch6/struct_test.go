package ch6

import "testing"

type Point struct {
	x int
	y int
}

type Circle struct {
	Point
	r int
}

func TestStruct(t *testing.T) {
	c := Circle{Point{1, 2}, 3}
	t.Log(c.x)
	t.Log(c.Point.y)
}
