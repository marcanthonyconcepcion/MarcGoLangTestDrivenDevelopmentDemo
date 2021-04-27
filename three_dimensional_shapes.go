/*
 * Copyright (c) 2021.
 * Marc Concepcion
 * marcanthonyconcepcion@gmail.com
 */
package MarcGoLangTestDrivenDevelopmentDemo

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
