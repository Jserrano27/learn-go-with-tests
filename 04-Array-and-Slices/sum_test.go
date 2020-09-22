package arrayslices

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{2, 4, 6, 8}

		got := Sum(numbers)
		expected := 20

		if got != expected {
			t.Errorf("Expected %d, but got %d instead. Given was %v", expected, got, numbers)
		}
	})
}

func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(Sum(numbers))

	// Output: 15
}
