package shapes

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return 0
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 0
}

// Perimeter returns the perimeter of a rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area returns the area of a rectangle
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
