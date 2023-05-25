package chapter6_methods

import "math"

type Point struct {
	X, Y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance2(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
