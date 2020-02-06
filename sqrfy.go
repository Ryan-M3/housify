package main

import (
	data "housify/data_structures"
	"math"
)

// This formula is the heart of the squarified treemap algorithm; otherwise, it
// would just be a regular treemapping algorithm. It estimates the worst aspect
// ratio if you were to squarify the areas. It works because it makes the
// simplifying assumption that the squarification algorithm produces areas that
// are approximately square; that is, it generally works *because* it assumes
// it will genrally work!
func worst(row []float64, w float64) float64 {
	var bad float64
	s := data.Sum(row)
	for _, r := range row {
		bad = math.Max(math.Max(w*w*r/(s*s), s*s/(w*w*r)), bad)
	}
	return bad
}

// Given a list of areas and the length of the row we're fitting them into,
// determine which areas to fit into the row that will produce the squarishest
// treemap. The function assumes the area list is already in descending order,
// and therefore partitions the list into two segments where the first
// partition needs to be fit into the row, and the second partition needs to
// have the squarified algorithm recursively applied to it.
func determineSplit(areas []float64, w float64) int {
	for i := 1; i < len(areas); i++ {
		if worst(areas[:i], w) <= worst(areas[i+1:], w) {
			return i
		}
	}
	return len(areas)
}

func Squarify(bounds data.Rect, areas *data.FTree) *data.RTree {
	parent := data.RTree{Value: bounds, Cnx: nil}
	if len(areas.Cnx) == 0 {
		return &parent
	}

	w := math.Max(bounds.Width(), bounds.Height())
	i := determineSplit(areas.CnxValues(), w)

	if i == len(areas.Cnx) {
		// Recursively squarify each fitted rectangle.
		for i, r := range data.FitInto(&bounds, areas) {
			subtree := Squarify(*r, areas.Cnx[i])
			parent.Cnx = append(parent.Cnx, subtree)
		}
		return &parent
	} else {
		var remL data.Rect
		var remR data.Rect
		numer := data.Sum(areas.CnxValues()[:i])
		denom := data.Sum(areas.CnxValues())
		ratio := numer / denom
		if bounds.Width() > bounds.Height() {
			remL, remR = data.SplitDownTheMiddle(&bounds, ratio)
		} else {
			remL, remR = data.SplitLeftToRight(&bounds, ratio)
		}

		cnxL, cnxR := areas.Cnx[:i], areas.Cnx[i:]
		areas.Cnx = cnxL
		for i, r := range data.FitInto(&remL, areas) {
			subtree := Squarify(*r, areas.Cnx[i])
			parent.Cnx = append(parent.Cnx, subtree)
		}

		areas.Cnx = cnxR
		for _, branch := range Squarify(remR, areas).Cnx {
			parent.Cnx = append(parent.Cnx, branch)
		}
	}

	return &parent
}
