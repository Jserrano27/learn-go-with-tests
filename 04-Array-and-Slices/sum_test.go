package arrayslices

import (
	"fmt"
	"reflect"
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

func TestSumAll(t *testing.T) {
	t.Run("with two slices", func(t *testing.T) {

		got := SumAll([]int{1, 2, 3}, []int{1, 0, 1})
		expected := []int{6, 2}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Expected %v, but got %v instead", expected, got)
		}
	})

	t.Run("with a single slice", func(t *testing.T) {
		numbers := []int{5, 5}

		got := SumAll(numbers)
		expected := []int{10}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Expected %v, but got %v instead", expected, got)
		}
	})
}

func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(Sum(numbers))

	// Output: 15
}
