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
	checkArea := func(t *testing.T, shape Shape, expected float64) {
		t.Helper()
		got := shape.Area()

		if expected != got {
			t.Errorf("Expected %g, but got %g instead", expected, got)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.00, 3.00}
		checkArea(t, rectangle, 30.00)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{5.00}
		checkArea(t, circle, 78.53981633974483)
	})

}
