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

	checkSums := func(t *testing.T, got []int, expected []int) {
		t.Helper()

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Expected %v, but got %v instead", expected, got)
		}
	}

	t.Run("with two slices", func(t *testing.T) {

		got := SumAll([]int{1, 2, 3}, []int{1, 0, 1})
		expected := []int{6, 2}

		checkSums(t, got, expected)
	})

	t.Run("with a single slice", func(t *testing.T) {
		numbers := []int{5, 5}

		got := SumAll(numbers)
		expected := []int{10}

		checkSums(t, got, expected)
	})

	t.Run("with an empty slice", func(t *testing.T) {
		got := SumAll([]int{})
		expected := []int{0}

		checkSums(t, got, expected)
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got []int, expected []int) {
		t.Helper()

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Expected %v, but got %v instead", expected, got)
		}
	}

	t.Run("with two slices", func(t *testing.T) {

		got := SumAllTails([]int{1, 2, 3, 4}, []int{0, 1, 1})
		expected := []int{9, 2}

		checkSums(t, got, expected)
	})

	t.Run("with a single slice", func(t *testing.T) {
		numbers := []int{2, 5}

		got := SumAllTails(numbers)
		expected := []int{5}

		checkSums(t, got, expected)
	})

	t.Run("with an empty slice", func(t *testing.T) {

		got := SumAllTails([]int{}, []int{0, 1, 1})
		expected := []int{0, 2}

		checkSums(t, got, expected)
	})
}

func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(Sum(numbers))

	// Output: 15
}
