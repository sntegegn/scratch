package structs_interfaces

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	testPerimeter := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "Rectangle", shape: Rectangle{width: 5.0, height: 4.0}, expected: 18.0},
		{name: "Circle", shape: Circle{radius: 3.0}, expected: math.Pi * 6.0},
		{name: "Triangle", shape: Triangle{base: 8.0, side1: 5.0, side2: 6.0}, expected: 19.0},
	}
	for _, pt := range testPerimeter {
		t.Run(pt.name, func(t *testing.T) {
			actual := pt.shape.Perimeter()
			if actual != pt.expected {
				t.Errorf("testing %#v expected %g got %g", pt.shape, pt.expected, actual)
			}
		})
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "Rectangle", shape: Rectangle{width: 4.0, height: 5.0}, expected: 20},
		{name: "Circle", shape: Circle{radius: 3.0}, expected: math.Pi * 3.0 * 3.0},
		{name: "Triangle", shape: Triangle{base: 2.0, height: 6.0}, expected: 6.0},
	}

	for _, at := range areaTests {
		t.Run(at.name, func(t *testing.T) {
			actual := at.shape.Area()
			if actual != at.expected {
				t.Errorf("%#v failed expected %.2f got %.2f", at.shape, at.expected, actual)
			}
		})

	}

}
