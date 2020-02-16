package main

import (
	"math"
)

// Shape is a interface which can be implemented by our Rectangle or Circle
type Shape interface {
	Area() float64
}

// Rectangle defines a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Area returns the area of a rectangle r
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle defines a circle
type Circle struct {
	Radius float64
}

// Area returns the area of a circle c
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Height * t.Base) / 2
}

// Perimeter returns the perimeter of a rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area returns the area of a rectangle
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
