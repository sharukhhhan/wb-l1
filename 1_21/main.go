package main

import "fmt"

type Shape interface {
	Area() string
}

type Rectangle struct {
	Width, Height int
}

func (r *Rectangle) CalculateArea() string {
	return fmt.Sprintf("The area: %d", r.Width*r.Height)
}

type RectangleAdapter struct {
	Rectangle *Rectangle
}

func (a *RectangleAdapter) Area() string {
	return a.Rectangle.CalculateArea()
}

func main() {
	rect := &Rectangle{Width: 10, Height: 5}
	adapter := &RectangleAdapter{Rectangle: rect}

	var shape Shape = adapter
	fmt.Println(shape.Area())
}
