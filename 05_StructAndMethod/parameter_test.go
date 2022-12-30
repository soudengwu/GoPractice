package parameter

import "testing"

func TestParameter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Parameter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

// func TestArea(t *testing.T) {

// 	checkArea := func(t testing.TB, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()

// 		if got != want {
// 			t.Errorf("got %g, want %g", got, want)
// 		}
// 	}

// 	t.Run("test for Rectangle", func(t *testing.T) {

// 		rectangle := Rectangle{12.0, 6.0}
// 		checkArea(t, rectangle, 72.0)

// 	})

// 	t.Run("test for Circle", func(t *testing.T) {

// 		circle := Circle{10}
// 		checkArea(t, circle, 314.1592653589793)
// 	})
// }

// Table Driven test
func TestArea(t *testing.T) {
	areaTest := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, tc := range areaTest {
		got := tc.shape.Area()
		if got != tc.hasArea {
			t.Errorf("%#v got %g, want %g", tc.shape, got, tc.hasArea)
		}
	}
}
