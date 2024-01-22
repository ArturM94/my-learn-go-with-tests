package structs

import "math"

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	SideA  float64
	SideB  float64
	Base   float64
	Height float64
}

func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.Base
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
