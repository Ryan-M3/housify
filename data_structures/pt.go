package data_structures

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
