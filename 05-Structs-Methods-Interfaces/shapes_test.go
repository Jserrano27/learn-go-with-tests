package structsmethodsinterfaces

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{5.00, 10.00}
	got := rectangle.Perimeter()
	expected := 30.00

	if expected != got {
		t.Errorf("Expected %.2f, but got %.2f instead", expected, got)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "Rectangle", shape: Rectangle{Height: 10.00, Width: 3.00}, expected: 30.00},
		{name: "Circle", shape: Circle{Radius: 5.00}, expected: 78.53981633974483},
		{name: "Triangle", shape: Triangle{Base: 3.00, Height: 6.00}, expected: 9.00},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if tt.expected != got {
				t.Errorf("%#v expected %g, but got %g instead", tt.shape, tt.expected, got)
			}
		})
	}
}
