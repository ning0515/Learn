package calculate

import "testing"

func TestPerimeter(t *testing.T) {
	rec := Rectangle{10.0, 10.0}
	got := Perimeter(rec)
	want := 40.0

	if got != want {
		t.Errorf("want %.2f but got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {
	areaTest := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, want: 72.0},
		{name: "Circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, want: 36.0},
	}
	for _, tt := range areaTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %.2f want %.2f ", tt.shape, got, tt.want)
			}
		})

	}
	// checkArea := func(t *testing.T, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()
	// 	if got != want {
	// 		t.Errorf("got %.2f want %.2f", got, want)
	// 	}
	// }
	// t.Run("test rectangle", func(t *testing.T) {
	// 	rec := Rectangle{10.0, 10.0}
	// 	checkArea(t, rec, 100.0)
	// })

	// t.Run("test circles", func(t *testing.T) {
	// 	circle := Circle{10}
	// 	checkArea(t, circle, 314.1592653589793)
	// })

}
