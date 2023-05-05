package calculate

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	r float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

type Triangle struct {
	Base   float64
	Height float64
}

func (c Triangle) Area() float64 {
	return (c.Base * c.Height) * 0.5
}

func Perimeter(rec Rectangle) (peri float64) {
	peri = 2 * (rec.Width + rec.Height)
	return
}

func Area(rec Rectangle) (area float64) {
	return rec.Width * rec.Height
}
