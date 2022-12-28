package arrays

import (
	"fmt"
	"reflect"
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

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	// Reflect.DeepEqual is not type safe, so be careful.
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("we got %v and expecting %v", got, expected)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		// Reflect.DeepEqual is not type safe, so be careful.
		if !reflect.DeepEqual(got, want) {
			t.Errorf("we got %v and expecting %v", got, want)
		}
	}
	t.Run("make sum of all slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		checkSums(t, got, expected)
	})

	t.Run("safely sum empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}
		checkSums(t, got, expected)
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
