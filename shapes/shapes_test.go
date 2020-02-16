package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	// areaTests is a slice of structs
	areaTests := []struct {
		// and the structs that will go in it have a Shape and a float64
		shape Shape
		want  float64
	}{
		// And here's a list of the ad-hoc structs that go in our slice.
		{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		// {shape: Triangle{Base: 12, Height: 6}, want: 36.0},
		{shape: Triangle{Base: 12, Height: 6}, want: 10.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %.2f want %.2f", tt, got, tt.want)
		}
	}

}
