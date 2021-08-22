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

func TestRectangle(t *testing.T) {
	length := rand.Float64() * float64(rand.Int())
	width := rand.Float64() * float64(rand.Int())
	rectangle := Rectangle{length: length, width: width}

	expectedArea := length * width
	actualArea := rectangle.getArea()
	if expectedArea != actualArea {
		t.Errorf("Expected rectangle area should be %f. However, actual rectangle area is wrongly computed as %f.", expectedArea, actualArea)
	}

	expectedPerimeter := 2*length + 2*width
	actualPerimeter := rectangle.getPerimeter()
	if expectedPerimeter != actualPerimeter {
		t.Errorf("Expected rectangle perimeter should be %f. However, actual rectangle perimeter is wrongly computed as %f.", expectedPerimeter, actualPerimeter)
	}
}

func TestCircle(t *testing.T) {
	radius := rand.Float64() * float64(rand.Int())
	circle := Circle{radius: radius}

	expectedArea := math.Pi * math.Pow(radius, 2)
	actualArea := circle.getArea()
	if expectedArea != actualArea {
		t.Errorf("Expected circle area should be %f. However, actual circle area is wrongly computed as %f.", expectedArea, actualArea)
	}

	expectedPerimeter := 2 * math.Pi * radius
	actualPerimeter := circle.getPerimeter()
	if expectedPerimeter != actualPerimeter {
		t.Errorf("Expected circle perimeter should be %f. However, actual circle perimeter is wrongly computed as %f.", expectedPerimeter, actualPerimeter)
	}
}

func TestSquare(t *testing.T) {
	side := rand.Float64() * float64(rand.Int())
	square := makeSquare(side)

	expectedArea := math.Pow(side, 2)
	actualArea := square.getArea()
	if expectedArea != actualArea {
		t.Errorf("Expected square area should be %f. However, actual square area is wrongly computed as %f.", expectedArea, actualArea)
	}

	expectedPerimeter := 4 * side
	actualPerimeter := square.getPerimeter()
	if expectedPerimeter != actualPerimeter {
		t.Errorf("Expected square perimeter should be %f. However, actual square perimeter is wrongly computed as %f.", expectedPerimeter, actualPerimeter)
	}
}

func TestTriangle(t *testing.T) {
	a := float64(70) + rand.Float64()
	b := float64(100) + rand.Float64()
	c := float64(50) + rand.Float64()
	halfPerimeter := (a + b + c) / 2

	triangle, _ := makeTriangle(a, b, c)
	actualArea := triangle.getArea()
	expectedArea := math.Sqrt(halfPerimeter * (halfPerimeter - a) * (halfPerimeter - b) * (halfPerimeter - c))
	if expectedArea != actualArea {
		t.Errorf("Expected triangle area should be %f. However, actual triangle area is wrongly computed as %f.", expectedArea, actualArea)
	}

	actualPerimeter := triangle.getPerimeter()
	expectedPerimeter := a + b + c
	if expectedPerimeter != actualPerimeter {
		t.Errorf("Expected triangle perimeter should be %f. However, actual triangle perimeter is wrongly computed as %f.", expectedPerimeter, actualPerimeter)
	}
}

func TestInvalidTriangles(t *testing.T) {
	a1 := float64(5) + rand.Float64()
	b1 := float64(20) + rand.Float64()
	c1 := float64(30) + rand.Float64()
	_, error1 := makeTriangle(a1, b1, c1)
	if error1 == nil {
		t.Errorf("Error has not been raised. Please check your codes on making this triangle.")
	}

	a2 := float64(7) + rand.Float64()
	b2 := float64(31) + rand.Float64()
	c2 := float64(20) + rand.Float64()
	_, error2 := makeTriangle(a2, b2, c2)
	if error2 == nil {
		t.Errorf("Error has not been raised. Please check your codes on making this triangle.")
	}

	a3 := float64(40) + rand.Float64()
	b3 := float64(11) + rand.Float64()
	c3 := float64(22) + rand.Float64()
	_, error3 := makeTriangle(a3, b3, c3)
	if error3 == nil {
		t.Errorf("Error has not been raised. Please check your codes on making this triangle.")
	}

	a4 := float64(10)
	b4 := float64(20)
	c4 := float64(30)
	_, error4 := makeTriangle(a4, b4, c4)
	if error4 == nil {
		t.Errorf("Error has not been raised. Please check your codes on making this triangle.")
	}
}
