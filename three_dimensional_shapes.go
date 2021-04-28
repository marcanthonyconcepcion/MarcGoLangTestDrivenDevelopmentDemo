/*
 * Copyright (c) 2021.
 * Marc Concepcion
 * marcanthonyconcepcion@gmail.com
 */
package MarcGoLangTestDrivenDevelopmentDemo

import "math"

type ThreeDimensionalShape interface {
	get_volume() float64
	get_surface_area() float64
}

type Prism struct {
	height float64
	TwoDimensionalShape
}

func makePrism(height float64, shape TwoDimensionalShape) Prism {
	return Prism{height, shape}
}
func (prism Prism) get_volume() float64 {
	return prism.height * prism.TwoDimensionalShape.get_area()
}
func (prism Prism) get_surface_area() float64 {
	return 2*prism.TwoDimensionalShape.get_area() + prism.height*prism.TwoDimensionalShape.get_perimeter()
}

type RectangularPrism struct {
	Prism
}

func makeRectangularPrism(length float64, width float64, height float64) RectangularPrism {
	return RectangularPrism{makePrism(height, Rectangle{length, width})}
}
func (rectangularPrism RectangularPrism) get_volume() float64 {
	return rectangularPrism.height * rectangularPrism.TwoDimensionalShape.get_area()
}
func (rectangularPrism RectangularPrism) get_surface_area() float64 {
	return 2*rectangularPrism.TwoDimensionalShape.get_area() + rectangularPrism.height*rectangularPrism.TwoDimensionalShape.get_perimeter()
}

type TriangularPrism struct {
	Prism
}

func makeTriangularPrism(a float64, b float64, c float64, height float64) TriangularPrism {
	return TriangularPrism{makePrism(height, Triangle{a, b, c})}
}
func (triangularPrism TriangularPrism) get_volume() float64 {
	return triangularPrism.height * triangularPrism.TwoDimensionalShape.get_area()
}
func (triangularPrism TriangularPrism) get_surface_area() float64 {
	return 2*triangularPrism.TwoDimensionalShape.get_area() + triangularPrism.height*triangularPrism.TwoDimensionalShape.get_perimeter()
}

type Cylinder struct {
	Prism
}

func makeCylinder(height float64, radius float64) Cylinder {
	return Cylinder{makePrism(height, Circle{radius})}
}
func (cylinder Cylinder) get_volume() float64 {
	return cylinder.height * cylinder.TwoDimensionalShape.get_area()
}
func (cylinder Cylinder) get_surface_area() float64 {
	return 2*cylinder.TwoDimensionalShape.get_area() + cylinder.height*cylinder.TwoDimensionalShape.get_perimeter()
}

type Sphere struct {
	radius float64
}

func (sphere Sphere) get_volume() float64 {
	return 4 * math.Pi * math.Pow(sphere.radius, 3) / 3
}
func (sphere Sphere) get_surface_area() float64 {
	return 4 * math.Pi * math.Pow(sphere.radius, 2)
}

type Cube struct {
	RectangularPrism
}

func makeCube(side float64) Cube {
	return Cube{makeRectangularPrism(side, side, side)}
}
func (cube Cube) get_volume() float64 {
	return cube.RectangularPrism.get_volume()
}
func (cube Cube) get_surface_area() float64 {
	return cube.RectangularPrism.get_surface_area()
}
