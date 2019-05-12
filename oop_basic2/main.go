package main

import "fmt"

// Shape interface
type Shape interface {
	Area()
}

// Triangle struct
type Triangle struct {
	Height float32
	Base   float32
}

// Area returns area of a triangle
func (t Triangle) Area() {
	area := (t.Base * t.Height / 0.50)
	fmt.Printf("Area of %T is %f\n", t, area)
}

// Rectangle struct
type Rectangle struct {
	Height float32
	Width  float32
}

// Area return area of a rectangle
func (r Rectangle) Area() {
	area := r.Height * r.Width
	fmt.Printf("Area of %T is %f\n", r, area)
}

func main() {
	segitigaSatu := &Triangle{
		Height: 10.00,
		Base:   15.50,
	}

	persegiA := &Rectangle{
		Height: 30.00,
		Width:  80.00,
	}

	printArea(segitigaSatu)
	printArea(persegiA)

}

func printArea(s Shape) {
	s.Area()
}
