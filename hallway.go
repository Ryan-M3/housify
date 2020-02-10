package main

import (
	"fmt"
	data "housify/data_structures"
)

type Side int

const (
	top   Side = iota
	right Side = iota
	btm   Side = iota
	left  Side = iota
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

func unique(lines []data.Line) []data.Line {
	var output []data.Line
	added := make(map[data.Line]bool, 0)
	for _, ln := range lines {
		if _, ok := added[ln]; !ok {
			output = append(output, ln)
			added[ln] = true
		}
	}
	return output
}

func hasSide(sides []Side, query Side) bool {
	for _, side := range sides {
		if side == query {
			return true
		}
	}
	return false
}

func RoomsAdjToBackbone_(house *data.RTree, backbone data.Graph) map[*data.Rect][]Side {
	adj := make(map[*data.Rect][]Side, 0)
	for _, room := range house.Leafs() {
		a, b, c, d := data.RectToLines(&room.Value)
		for _, ln := range unique(data.GraphToLines(backbone)) {
			for i, roomLn := range []data.Line{a, b, c, d} {
				if data.Intersects(ln, roomLn) {
					side := []Side{top, right, btm, left}[i]
					if !hasSide(adj[&room.Value], side) {
						adj[&room.Value] = append(adj[&room.Value], side)
					}
				}
			}
		}
	}
	return adj
}

func RoomsAdjToBackbone(house *data.RTree, backbone data.Graph) map[*data.Rect][]Side {
	adj := make(map[*data.Rect][]Side, 0)
	for _, ln := range unique(data.GraphToLines(backbone)) {
		a, b := data.LineToPt(ln)
		rooms := append(house.FindNearestPt(a), house.FindNearestPt(b)...)
		//rooms := InBoth(house.FindNearestPt(a), house.FindNearestPt(b))
		if data.Horz(ln) {
			for _, r := range rooms {
				if sides, ok := adj[r]; ok {
					continue
				} else if r.AboveLine(ln) {
					if !hasSide(sides, btm) {
						adj[r] = append(adj[r], btm)
					}
				} else {
					if !hasSide(sides, top) {
						adj[r] = append(adj[r], top)
					}
				}
			}
		} else {
			for _, r := range rooms {
				if sides, ok := adj[r]; ok {
					continue
				} else if r.RightOfLine(ln) {
					if !hasSide(sides, left) {
						adj[r] = append(adj[r], left)
					}
				} else if r.LeftOfLine(ln) {
					if !hasSide(sides, right) {
						adj[r] = append(adj[r], right)
					}
				}
			}
		}
	}
	return adj
}

func InsertHallway(roomMap map[*data.Rect][]Side, width float64) {
	for room, sides := range roomMap {
		for _, side := range sides {
			fmt.Println(room)
			switch side {
			case top:
				room.SetHeightTop(room.Height() - width/2.0)
			case right:
				room.SetWidthR(room.Width() - width/2.0)
			case btm:
				room.SetHeightBtm(room.Height() - width/2.0)
			case left:
				room.SetWidthL(room.Width() - width/2.0)
			}
		}
	}
}

//func InsertHallway(house *data.RTree, backbone data.Graph, width float64) {
//	movedHorz := make(map[*data.Rect]bool, 0)
//	movedVert := make(map[*data.Rect]bool, 0)
//	for _, ln := range unique(data.GraphToLines(backbone)) {
//		a, b := data.LineToPt(ln)
//		rooms := append(house.FindNearestPt(a), house.FindNearestPt(b)...)
//		if data.Horz(ln) {
//			for _, r := range rooms {
//				if _, ok := movedHorz[r]; ok {
//					continue
//				} else {
//					movedHorz[r] = true
//				}
//				if r.AboveLine(ln) {
//					r.SetHeightBtm(r.Height() - width/2.0)
//				} else if r.BelowLine(ln) {
//					r.SetHeightTop(r.Height() - width/2.0)
//				} else {
//					panic("Invalid input encounterd in function call to InsertHallway()")
//				}
//			}
//		} else {
//			for _, r := range rooms {
//				if _, ok := movedVert[r]; ok {
//					continue
//				} else {
//					movedVert[r] = true
//				}
//				if r.RightOfLine(ln) {
//					r.SetWidthL(r.Width() - width/2.0)
//				} else if r.LeftOfLine(ln) {
//					r.SetWidthR(r.Width() - width/2.0)
//				} else {
//					panic("Invalid input encounterd in function call to InsertHallway()")
//				}
//			}
//		}
//	}
//}
