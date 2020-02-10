package data_structures

import (
	"math"
)

type Pt struct {
	X, Y float64
}

func HasPt(things []Pt, thing Pt) bool {
	for _, t := range things {
		if t == thing {
			return true
		}
	}
	return false
}

func Distance(a, b Pt) float64 {
	return math.Sqrt(math.Abs(a.X-b.X) + math.Abs(a.Y-b.Y))
}
