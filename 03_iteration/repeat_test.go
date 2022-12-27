package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("test repeat iteration", func(t *testing.T) {
		repeated := Repeat("b", 10)
		expected := "bbbbbbbbbb"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	repeatChar := Repeat("a", 5)
	fmt.Println(repeatChar)
	// Output: aaaaa

}
