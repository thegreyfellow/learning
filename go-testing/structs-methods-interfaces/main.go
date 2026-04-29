// Package main for explaining structs methods and interfaces
package main

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Height float64
	Width  float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (rec Rectangle) Perimeter() float64 {
	return 2 * (rec.Height + rec.Width)
}

func (rec Rectangle) Area() float64 {
	return rec.Height * rec.Width
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
