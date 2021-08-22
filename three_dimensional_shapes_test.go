// Copyright (c) 2021.
// Marc Concepcion
// marcanthonyconcepcion@gmail.com
package MarcGoLangTestDrivenDevelopmentDemo

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type TwoDimensionalShapeMock struct{}

var getAreaMock func() float64

func (twoDimensionalShapeMock TwoDimensionalShapeMock) getArea() float64 {
	return getAreaMock()
}

var getPerimeterMock func() float64

func (twoDimensionalShapeMock TwoDimensionalShapeMock) getPerimeter() float64 {
	return getPerimeterMock()
}

func TestPrism(t *testing.T) {
	length := rand.Float64() * float64(rand.Int())
	width := rand.Float64() * float64(rand.Int())
	height := rand.Float64() * float64(rand.Int())
	baseRectangleArea := length * width
	baseRectanglePerimeter := 2*length + width
	shape := TwoDimensionalShapeMock{}
	getAreaMock = func() float64 {
		return baseRectangleArea
	}
	getPerimeterMock = func() float64 {
		return baseRectanglePerimeter
	}
	prism := makePrism(height, shape)
	actualVolume := prism.getVolume()
	expectedVolume := height * baseRectangleArea
	if expectedVolume != actualVolume {
		t.Errorf("Expected prism volume should be %f. However, actual prism volume is wrongly computed as %f.", expectedVolume, actualVolume)
	}
	actualSurfaceArea := prism.getSurfaceArea()
	expectedSurfaceArea := 2*baseRectangleArea + height*baseRectanglePerimeter
	if expectedSurfaceArea != actualSurfaceArea {
		t.Errorf("Expected prism surface area should be %f. However, actual prism surface area is wrongly computed as %f.", expectedSurfaceArea, actualSurfaceArea)
	}
}

func TestRectangularPrism(t *testing.T) {
	length := rand.Float64() * float64(rand.Int())
	width := rand.Float64() * float64(rand.Int())
	height := rand.Float64() * float64(rand.Int())

	baseRectangleArea := length * width
	baseRectanglePerimeter := 2 * (length + width)

	rectangularPrism := makeRectangularPrism(length, width, height)
	actualVolume := rectangularPrism.getVolume()
	expectedVolume := height * baseRectangleArea
	if expectedVolume != actualVolume {
		t.Errorf("Expected rectangular prism volume should be %f. However, actual rectangular prism volume is wrongly computed as %f.", expectedVolume, actualVolume)
	}
	actualSurfaceArea := rectangularPrism.getSurfaceArea()
	expectedSurfaceArea := 2*baseRectangleArea + height*baseRectanglePerimeter
	if expectedSurfaceArea != actualSurfaceArea {
		t.Errorf("Expected rectangular prism surface area should be %f. However, actual rectangular prism surface area is wrongly computed as %f.", expectedSurfaceArea, actualSurfaceArea)
	}
}

func TestTriangularPrism(t *testing.T) {
	a := float64(20) + rand.Float64()
	b := float64(30) + rand.Float64()
	c := float64(40) + rand.Float64()
	height := rand.Float64() * float64(rand.Int())
	baseTrianglePerimeter := a + b + c
	baseTriangleHalfPerimeter := baseTrianglePerimeter / 2
	baseTriangleArea := math.Sqrt(baseTriangleHalfPerimeter * (baseTriangleHalfPerimeter - a) * (baseTriangleHalfPerimeter - b) * (baseTriangleHalfPerimeter - c))

	triangularPrism := makeTriangularPrism(a, b, c, height)
	actualVolume := triangularPrism.getVolume()
	expectedVolume := height * baseTriangleArea
	if expectedVolume != actualVolume {
		t.Errorf("Expected triangular prism volume should be %f. However, actual triangular prism volume is wrongly computed as %f.", expectedVolume, actualVolume)
	}
	actualSurfaceArea := triangularPrism.getSurfaceArea()
	expectedSurfaceArea := 2*baseTriangleArea + height*baseTrianglePerimeter
	if expectedSurfaceArea != actualSurfaceArea {
		t.Errorf("Expected triangular prism surface area should be %f. However, actual triangular prism surface area is wrongly computed as %f.", expectedSurfaceArea, actualSurfaceArea)
	}
}

func TestCylinder(t *testing.T) {
	radius := rand.Float64() * float64(rand.Int())
	height := rand.Float64() * float64(rand.Int())
	baseCircleArea := math.Pi * math.Pow(radius, 2)
	baseCirclePerimeter := 2 * math.Pi * radius

	cylinder := makeCylinder(height, radius)
	actualVolume := cylinder.getVolume()
	expectedPrismVolume := height * baseCircleArea
	if expectedPrismVolume != actualVolume {
		t.Errorf("Expected cylinder volume should be %f. However, actual cylinder volume is wrongly computed as %f.", expectedPrismVolume, actualVolume)
	}
	actualSurfaceArea := cylinder.getSurfaceArea()
	expectedPrismSurfaceArea := 2*baseCircleArea + height*baseCirclePerimeter
	if expectedPrismSurfaceArea != actualSurfaceArea {
		t.Errorf("Expected cylinder surface area should be %f. However, actual cylinder surface area is wrongly computed as %f.", expectedPrismSurfaceArea, actualSurfaceArea)
	}
	expectedCalculatedVolume := math.Pi * math.Pow(radius, 2) * height
	if expectedCalculatedVolume != actualVolume {
		t.Errorf("Expected cylinder volume should be %f. However, actual cylinder volume is wrongly computed as %f.", expectedCalculatedVolume, actualVolume)
	}
	expectedCalculatedSurfaceArea := 2*math.Pi*radius*height + 2*math.Pi*math.Pow(radius, 2)
	if expectedCalculatedSurfaceArea != actualSurfaceArea {
		t.Errorf("Expected cylinder surface area should be %f. However, actual cylinder surface area is wrongly computed as %f.", expectedCalculatedSurfaceArea, actualSurfaceArea)
	}
}

func TestSphere(t *testing.T) {
	radius := rand.Float64() * float64(rand.Int())
	sphere := Sphere{radius}
	actualVolume := sphere.getVolume()
	expectedVolume := 4 * math.Pi * math.Pow(radius, 3) / 3
	if expectedVolume != actualVolume {
		t.Errorf("Expected sphere volume should be %f. However, actual sphere volume is wrongly computed as %f.", expectedVolume, actualVolume)
	}
	actualSurfaceArea := sphere.getSurfaceArea()
	expectedSurfaceArea := 4 * math.Pi * math.Pow(radius, 2)
	if expectedSurfaceArea != actualSurfaceArea {
		t.Errorf("Expected sphere surface area should be %f. However, actual sphere surface area is wrongly computed as %f.", expectedSurfaceArea, actualSurfaceArea)
	}
}

func TestCube(t *testing.T) {
	side := rand.Float64() * float64(rand.Int())
	baseSquareArea := math.Pow(side, 2)
	baseSquarePerimeter := 4 * side

	cube := makeCube(side)
	actualVolume := cube.getVolume()
	expectedPrismVolume := side * baseSquareArea
	if expectedPrismVolume != actualVolume {
		t.Errorf("Expected cube volume should be %f. However, actual cube volume is wrongly computed as %f.", expectedPrismVolume, actualVolume)
	}
	actualSurfaceArea := cube.getSurfaceArea()
	expectedPrismSurfaceArea := 2*baseSquareArea + side*baseSquarePerimeter
	if expectedPrismSurfaceArea != actualSurfaceArea {
		t.Errorf("Expected cube surface area should be %f. However, actual cube surface area is wrongly computed as %f.", expectedPrismSurfaceArea, actualSurfaceArea)
	}
	expectedCalculatedVolume := math.Pow(side, 3)
	if expectedCalculatedVolume != actualVolume {
		t.Errorf("Expected cube volume should be %f. However, actual cube volume is wrongly computed as %f.", expectedCalculatedVolume, actualVolume)
	}
	expectedCalculatedSurfaceArea := 6 * math.Pow(side, 2)
	if expectedCalculatedSurfaceArea != actualSurfaceArea {
		t.Errorf("Expected cube surface area should be %f. However, actual cube surface area is wrongly computed as %f.", expectedCalculatedSurfaceArea, actualSurfaceArea)
	}
}
