// Package MarcGoLangTestDrivenDevelopmentDemo
// Copyright (c) 2021.
// Marc Concepcion
// marcanthonyconcepcion@gmail.com
package MarcGoLangTestDrivenDevelopmentDemo

import "math"

type ThreeDimensionalShape interface {
	getVolume() float64
	getSurfaceArea() float64
}

type Prism struct {
	height float64
	TwoDimensionalShape
}

func makePrism(height float64, shape TwoDimensionalShape) Prism {
	return Prism{height, shape}
}
func (prism Prism) getVolume() float64 {
	return prism.height * prism.TwoDimensionalShape.getArea()
}
func (prism Prism) getSurfaceArea() float64 {
	return 2*prism.TwoDimensionalShape.getArea() + prism.height*prism.TwoDimensionalShape.getPerimeter()
}

type RectangularPrism struct {
	Prism
}

func makeRectangularPrism(length float64, width float64, height float64) RectangularPrism {
	return RectangularPrism{makePrism(height, Rectangle{length, width})}
}
func (rectangularPrism RectangularPrism) getVolume() float64 {
	return rectangularPrism.height * rectangularPrism.TwoDimensionalShape.getArea()
}
func (rectangularPrism RectangularPrism) getSurfaceArea() float64 {
	return 2*rectangularPrism.TwoDimensionalShape.getArea() + rectangularPrism.height*rectangularPrism.TwoDimensionalShape.getPerimeter()
}

type TriangularPrism struct {
	Prism
}

func makeTriangularPrism(a float64, b float64, c float64, height float64) TriangularPrism {
	return TriangularPrism{makePrism(height, Triangle{a, b, c})}
}
func (triangularPrism TriangularPrism) getVolume() float64 {
	return triangularPrism.height * triangularPrism.TwoDimensionalShape.getArea()
}
func (triangularPrism TriangularPrism) getSurfaceArea() float64 {
	return 2*triangularPrism.TwoDimensionalShape.getArea() + triangularPrism.height*triangularPrism.TwoDimensionalShape.getPerimeter()
}

type Cylinder struct {
	Prism
}

func makeCylinder(height float64, radius float64) Cylinder {
	return Cylinder{makePrism(height, Circle{radius})}
}
func (cylinder Cylinder) getVolume() float64 {
	return cylinder.height * cylinder.TwoDimensionalShape.getArea()
}
func (cylinder Cylinder) getSurfaceArea() float64 {
	return 2*cylinder.TwoDimensionalShape.getArea() + cylinder.height*cylinder.TwoDimensionalShape.getPerimeter()
}

type Sphere struct {
	radius float64
}

func (sphere Sphere) getVolume() float64 {
	return 4 * math.Pi * math.Pow(sphere.radius, 3) / 3
}
func (sphere Sphere) getSurfaceArea() float64 {
	return 4 * math.Pi * math.Pow(sphere.radius, 2)
}

type Cube struct {
	RectangularPrism
}

func makeCube(side float64) Cube {
	return Cube{makeRectangularPrism(side, side, side)}
}
func (cube Cube) getVolume() float64 {
	return cube.RectangularPrism.getVolume()
}
func (cube Cube) getSurfaceArea() float64 {
	return cube.RectangularPrism.getSurfaceArea()
}
