package structsmethodsinterfaces

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{5.00, 10.00}
	got := Perimeter(rectangle)
	expected := 30.00

	if expected != got {
		t.Errorf("Expected %.2f, but got %.2f instead", expected, got)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{10.00, 3.00}
	got := Area(rectangle)
	expected := 30.00

	if expected != got {
		t.Errorf("Expected %.2f, but got %.2f instead", expected, got)
	}
}
