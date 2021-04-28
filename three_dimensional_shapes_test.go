/*
 * Copyright (c) 2021.
 * Marc Concepcion
 * marcanthonyconcepcion@gmail.com
 */
package MarcGoLangTestDrivenDevelopmentDemo

import (
	"math"
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
		t.Errorf("Expected prism surface area should be %f. However, actual prism surface area is wrongly computed as %f.", expected_surface_area, actual_surface_area)
	}
}

func TestRectangularPrism(t *testing.T) {
	length := float64(10)
	width := float64(20)
	height := float64(30)

	base_rectangle_area := length * width
	base_rectangle_perimeter := 2 * (length + width)

	rectangular_prism := makeRectangularPrism(length, width, height)
	actual_volume := rectangular_prism.get_volume()
	expected_volume := height * base_rectangle_area
	if expected_volume != actual_volume {
		t.Errorf("Expected rectangular prism volume should be %f. However, actual rectangular prism volume is wrongly computed as %f.", expected_volume, actual_volume)
	}
	actual_surface_area := rectangular_prism.get_surface_area()
	expected_surface_area := 2*base_rectangle_area + height*base_rectangle_perimeter
	if expected_surface_area != actual_surface_area {
		t.Errorf("Expected rectangular prism surface area should be %f. However, actual rectangular prism surface area is wrongly computed as %f.", expected_surface_area, actual_surface_area)
	}
}

func TestTriangularPrism(t *testing.T) {
	a := float64(20)
	b := float64(30)
	c := float64(40)
	height := float64(30)
	base_triangle_perimeter := a + b + c
	base_triangle_half_perimeter := base_triangle_perimeter / 2
	base_triangle_area := math.Sqrt(base_triangle_half_perimeter * (base_triangle_half_perimeter - a) * (base_triangle_half_perimeter - b) * (base_triangle_half_perimeter - c))

	triangular_prism := makeTriangularPrism(a, b, c, height)
	actual_volume := triangular_prism.get_volume()
	expected_volume := height * base_triangle_area
	if expected_volume != actual_volume {
		t.Errorf("Expected triangular prism volume should be %f. However, actual triangular prism volume is wrongly computed as %f.", expected_volume, actual_volume)
	}
	actual_surface_area := triangular_prism.get_surface_area()
	expected_surface_area := 2*base_triangle_area + height*base_triangle_perimeter
	if expected_surface_area != actual_surface_area {
		t.Errorf("Expected triangular prism surface area should be %f. However, actual triangular prism surface area is wrongly computed as %f.", expected_surface_area, actual_surface_area)
	}
}

func TestCylinder(t *testing.T) {
	radius := float64(10)
	height := float64(30)
	base_circle_area := math.Pi * math.Pow(radius, 2)
	base_circle_perimeter := 2 * math.Pi * radius

	cylinder := makeCylinder(height, radius)
	actual_volume := cylinder.get_volume()
	expected_prism_volume := height * base_circle_area
	if expected_prism_volume != actual_volume {
		t.Errorf("Expected cylinder volume should be %f. However, actual cylinder volume is wrongly computed as %f.", expected_prism_volume, actual_volume)
	}
	actual_surface_area := cylinder.get_surface_area()
	expected_prism_surface_area := 2*base_circle_area + height*base_circle_perimeter
	if expected_prism_surface_area != actual_surface_area {
		t.Errorf("Expected cylinder surface area should be %f. However, actual cylinder surface area is wrongly computed as %f.", expected_prism_surface_area, actual_surface_area)
	}
	expected_calculated_volume := math.Pi * math.Pow(radius, 2) * height
	if expected_calculated_volume != actual_volume {
		t.Errorf("Expected cylinder volume should be %f. However, actual cylinder volume is wrongly computed as %f.", expected_calculated_volume, actual_volume)
	}
	expected_calculated_surface_area := 2*math.Pi*radius*height + 2*math.Pi*math.Pow(radius, 2)
	if expected_calculated_surface_area != actual_surface_area {
		t.Errorf("Expected cylinder surface area should be %f. However, actual cylinder surface area is wrongly computed as %f.", expected_calculated_surface_area, actual_surface_area)
	}
}

func TestSphere(t *testing.T) {
	radius := float64(10)
	sphere := Sphere{radius}
	actual_volume := sphere.get_volume()
	expected_volume := 4 * math.Pi * math.Pow(radius, 3) / 3
	if expected_volume != actual_volume {
		t.Errorf("Expected sphere volume should be %f. However, actual sphere volume is wrongly computed as %f.", expected_volume, actual_volume)
	}
	actual_surface_area := sphere.get_surface_area()
	expected_surface_area := 4 * math.Pi * math.Pow(radius, 2)
	if expected_surface_area != actual_surface_area {
		t.Errorf("Expected sphere surface area should be %f. However, actual sphere surface area is wrongly computed as %f.", expected_surface_area, actual_surface_area)
	}
}

func TestCube(t *testing.T) {
	side := float64(10)
	base_square_area := math.Pow(side, 2)
	base_square_perimeter := 4 * side

	cube := makeCube(side)
	actual_volume := cube.get_volume()
	expected_prism_volume := side * base_square_area
	if expected_prism_volume != actual_volume {
		t.Errorf("Expected cube volume should be %f. However, actual cube volume is wrongly computed as %f.", expected_prism_volume, actual_volume)
	}
	actual_surface_area := cube.get_surface_area()
	expected_prism_surface_area := 2*base_square_area + side*base_square_perimeter
	if expected_prism_surface_area != actual_surface_area {
		t.Errorf("Expected cube surface area should be %f. However, actual cube surface area is wrongly computed as %f.", expected_prism_surface_area, actual_surface_area)
	}
	expected_calculated_volume := math.Pow(side, 3)
	if expected_calculated_volume != actual_volume {
		t.Errorf("Expected cube volume should be %f. However, actual cube volume is wrongly computed as %f.", expected_calculated_volume, actual_volume)
	}
	expected_calculated_surface_area := 6 * math.Pow(side, 2)
	if expected_calculated_surface_area != actual_surface_area {
		t.Errorf("Expected cube surface area should be %f. However, actual cube surface area is wrongly computed as %f.", expected_calculated_surface_area, actual_surface_area)
	}
}
