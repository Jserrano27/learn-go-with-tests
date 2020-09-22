package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	actual := Repeat("a", 7)
	expected := "aaaaaaa"
	if expected != actual {
		t.Errorf("Expected %q, but got %q instead", expected, actual)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	s := "a"
	times := 5
	fmt.Println(Repeat(s, times))

	// Output: aaaaa
}
