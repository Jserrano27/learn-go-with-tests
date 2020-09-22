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

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.00, 3.00}
		got := rectangle.Area()
		expected := 30.00

		if expected != got {
			t.Errorf("Expected %.2f, but got %.2f instead", expected, got)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{5.00}
		got := circle.Area()
		expected := 78.53981633974483

		if expected != got {
			t.Errorf("Expected %g, but got %g instead", expected, got)
		}
	})

}
