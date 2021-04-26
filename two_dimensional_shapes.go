/*
 * Copyright (c) 2021.
 * Marc Concepcion
 * marcanthonyconcepcion@gmail.com
 */
package two_dimensional_shapes

import (
	"math"
)

type TwoDimensionalShape interface {
	get_area() float64
	get_perimeter() float64
}

type Rectangle struct {
	length float64
	width  float64
}

func (rectangle Rectangle) get_area() float64 {
	return rectangle.length * rectangle.width
}
func (rectangle Rectangle) get_perimeter() float64 {
	return 2*rectangle.length + 2*rectangle.width
}

type Circle struct {
	radius float64
}

func (circle Circle) get_area() float64 {
	return math.Pi * math.Pow(circle.radius, 2)
}
func (circle Circle) get_perimeter() float64 {
	return 2 * math.Pi * circle.radius
}

type Square struct {
	side float64
	Rectangle
}

func makeSquare(side float64) Square {
	return Square{side, Rectangle{length: side, width: side}}
}
func (square Square) get_area() float64 {
	return square.Rectangle.get_area()
}
func (square Square) get_perimeter() float64 {
	return square.Rectangle.get_perimeter()
}
