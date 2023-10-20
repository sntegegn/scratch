package structs_interfaces

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func Perimeter(r Rectangle) float64 {
	return 2.0 * (r.width + r.height)
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type Triangle struct {
	base   float64
	height float64
	side1  float64
	side2  float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

func (t Triangle) Perimeter() float64 {
	return t.side1 + t.side2 + t.base
}
