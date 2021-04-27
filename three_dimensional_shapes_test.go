/*
 * Copyright (c) 2021.
 * Marc Concepcion
 * marcanthonyconcepcion@gmail.com
 */
package MarcGoLangTestDrivenDevelopmentDemo

import (
	"testing"
)

type TwoDimensionalShapeMock struct{}

var getAreaMock func() float64

func (twoDimensionalShapeMock TwoDimensionalShapeMock) get_area() float64 {
	return getAreaMock()
}

var getPerimeterMock func() float64

func (twoDimensionalShapeMock TwoDimensionalShapeMock) get_perimeter() float64 {
	return getPerimeterMock()
}

func TestPrism(t *testing.T) {
	length := float64(40)
	width := float64(50.35)
	height := float64(60)
	base_rectangle_area := length * width
	base_rectangle_perimeter := 2*length + width
	shape := TwoDimensionalShapeMock{}
	getAreaMock = func() float64 {
		return base_rectangle_area
	}
	getPerimeterMock = func() float64 {
		return base_rectangle_perimeter
	}
	prism := makePrism(height, shape)
	actual_volume := prism.get_volume()
	expected_volume := height * base_rectangle_area
	if expected_volume != actual_volume {
		t.Errorf("Expected prism volume should be %f. However, actual prism volume is wrongly computed as %f.", expected_volume, actual_volume)
	}
	actual_surface_area := prism.get_surface_area()
	expected_surface_area := 2*base_rectangle_area + height*base_rectangle_perimeter
	if expected_surface_area != actual_surface_area {
		t.Errorf("Expected prism surface area should be %f. However, actual rectangle area is wrongly computed as %f.", expected_surface_area, actual_surface_area)
	}
}
