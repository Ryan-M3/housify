package main

import (
	data "housify/data_structures"
)

func InBoth(xs, ys []*data.Rect) []*data.Rect {
	var output []*data.Rect
	// This looks inefficient, but each list can be at most 4 items long, and
	// will usually have items in common, which skips to the next iteration of
	// the loop.
	for _, x := range xs {
		for _, y := range ys {
			if x == y {
				output = append(output, x)
				continue
			}
		}
	}
	return output
}

func InsertHallway(house *data.RTree, backbone data.Graph, width float64) {
	for _, ln := range data.GraphToLines(backbone) {
		a, b := data.LineToPt(ln)
		rooms := InBoth(house.FindNearestPt(a), house.FindNearestPt(b))
		if data.Horz(ln) {
			for _, r := range rooms {
				if r.AboveLine(ln) {
					r.SetHeightBtm(r.Height() - width/2.0)
				} else if r.BelowLine(ln) {
					r.SetHeightTop(r.Height() - width/2.0)
				} else {
					panic("Invalid input encounterd in function call to InsertHallway()")
				}
			}
		} else {
			for _, r := range rooms {
				if r.RightOfLine(ln) {
					r.SetWidthL(r.Width() - width/2.0)
				} else if r.LeftOfLine(ln) {
					r.SetWidthR(r.Width() - width/2.0)
				} else {
					panic("Invalid input encounterd in function call to InsertHallway()")
				}
			}
		}
	}
}
