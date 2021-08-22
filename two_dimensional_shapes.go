// Package MarcGoLangTestDrivenDevelopmentDemo
// Copyright (c) 2021.
// Marc Concepcion
// marcanthonyconcepcion@gmail.com
package MarcGoLangTestDrivenDevelopmentDemo

import (
	"errors"
	"math"
)

type TwoDimensionalShape interface {
	getArea() float64
	getPerimeter() float64
}

type Rectangle struct {
	length float64
	width  float64
}

func (rectangle Rectangle) getArea() float64 {
	return rectangle.length * rectangle.width
}
func (rectangle Rectangle) getPerimeter() float64 {
	return 2*rectangle.length + 2*rectangle.width
}

type Circle struct {
	radius float64
}

func (circle Circle) getArea() float64 {
	return math.Pi * math.Pow(circle.radius, 2)
}
func (circle Circle) getPerimeter() float64 {
	return 2 * math.Pi * circle.radius
}

type Square struct {
	side float64
	Rectangle
}

func makeSquare(side float64) Square {
	return Square{side, Rectangle{length: side, width: side}}
}
func (square Square) getArea() float64 {
	return square.Rectangle.getArea()
}
func (square Square) getPerimeter() float64 {
	return square.Rectangle.getPerimeter()
}

type Triangle struct {
	a float64
	b float64
	c float64
}

func makeTriangle(a float64, b float64, c float64) (*Triangle, error) {
	if !(a+b > c && a+c > b && b+c > a) {
		return nil, errors.New("invalid triangle")
	}
	return &Triangle{a, b, c}, nil
}

func (triangle Triangle) getArea() float64 {
	halfPerimeter := triangle.getPerimeter() / 2
	return math.Sqrt(halfPerimeter * (halfPerimeter - triangle.a) * (halfPerimeter - triangle.b) * (halfPerimeter - triangle.c))
}
func (triangle Triangle) getPerimeter() float64 {
	return triangle.a + triangle.b + triangle.c
}
