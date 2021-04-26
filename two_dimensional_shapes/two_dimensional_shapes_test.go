/*
 * Copyright (c) 2021.
 * Marc Concepcion
 * marcanthonyconcepcion@gmail.com
 */
package two_dimensional_shapes

import (
	"math"
	"testing"
)

func TestRectangle(t *testing.T) {
	length := float64(10.45)
	width := float64(20)
	rectangle := Rectangle{length: length, width: width}

	expected_area := length * width
	actual_area := rectangle.get_area()
	if expected_area != actual_area {
		t.Errorf("Expected rectangle area should be %f. However, actual rectangle area is wrongly computed as %f.", expected_area, actual_area)
	}

	expected_perimeter := 2*length + 2*width
	actual_perimeter := rectangle.get_perimeter()
	if expected_perimeter != actual_perimeter {
		t.Errorf("Expected rectangle perimeter should be %f. However, actual rectangle perimeter is wrongly computed as %f.", expected_area, actual_area)
	}
}

func TestCircle(t *testing.T) {
	radius := float64(10)
	circle := Circle{radius: radius}

	expected_area := math.Pi * math.Pow(radius, 2)
	actual_area := circle.get_area()
	if expected_area != actual_area {
		t.Errorf("Expected circle area should be %f. However, actual circle area is wrongly computed as %f.", expected_area, actual_area)
	}

	expected_perimeter := 2 * math.Pi * radius
	actual_perimeter := circle.get_perimeter()
	if expected_perimeter != actual_perimeter {
		t.Errorf("Expected circle perimeter should be %f. However, actual circle perimeter is wrongly computed as %f.", expected_area, actual_area)
	}
}

func TestSquare(t *testing.T) {
	side := float64(10)
	square := makeSquare(side)

	expected_area := math.Pow(side, 2)
	actual_area := square.get_area()
	if expected_area != actual_area {
		t.Errorf("Expected square area should be %f. However, actual square area is wrongly computed as %f.", expected_area, actual_area)
	}

	expected_perimeter := 4 * side
	actual_perimeter := square.get_perimeter()
	if expected_perimeter != actual_perimeter {
		t.Errorf("Expected square perimeter should be %f. However, actual square perimeter is wrongly computed as %f.", expected_area, actual_area)
	}
}
