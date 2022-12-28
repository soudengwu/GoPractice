package arrays

import (
	"fmt"
	"testing"
)

func TestArraySum(t *testing.T) {
	t.Run("test slice sum", func(t *testing.T) {
		input := []int{1, 2, 3}
		got := ArraySum(input)
		expected := 6

		if got != expected {
			t.Errorf("The input of the slice is %v, we got %d and expecting %d", input, got, expected)
		}
	})
}

func BenchmarkArraySum(b *testing.B) {
	input := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		ArraySum(input)
	}
}

func ExampleArraySum() {
	input := []int{1, 2, 3, 4, 5}
	sum := ArraySum(input)
	fmt.Println(sum)
	// Output: 15
}
