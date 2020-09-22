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
		shape    Shape
		expected float64
	}{
		{Rectangle{10.00, 3.00}, 30.00},
		{Circle{5.00}, 78.53981633974483},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()

		if tt.expected != got {
			t.Errorf("Expected %g, but got %g instead", tt.expected, got)
		}
	}
}
