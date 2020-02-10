package dataStructures

import (
	"math"
)

type Pt struct {
	X, Y float64
}

func (pt *Pt) Mult(amt float64) Pt {
	return Pt{pt.X * amt, pt.Y * amt}
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

func Lerp(t float64, a, b Pt) Pt {
	return Pt{
		(a.X + (b.X-a.X)*t),
		(a.Y + (b.Y-a.Y)*t),
	}
}
