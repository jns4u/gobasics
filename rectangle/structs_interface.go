package rectangle

import (
	 "math"
)

func distance(X1, Y1, X2, Y2 float64) float64 {
	a := X2 - X1
	b := Y2 - Y1
	return math.Sqrt(a*a + b*b)
}

// struct
type Rectangle struct {
	X1, Y1, X2, Y2 float64
}

func (r *Rectangle) Area() float64 {
	l := distance(r.X1, r.Y1, r.X1, r.Y2)
	w := distance(r.X1, r.Y1, r.X2, r.Y1)
	return l * w
}
