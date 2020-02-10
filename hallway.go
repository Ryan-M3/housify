package main

import (
	data "housify/dataStructures"
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

// Return a list of rooms which are adjacent to the provided backbone.
func RoomsAdjToBackbone(house []*data.Room, backbone data.Graph) map[*data.Room][]Side {
	adj := make(map[*data.Room][]Side, 0)
	for _, room := range house {
		a, b, c, d := data.RectToLines(room.Rect)
		for _, ln := range unique(data.GraphToLines(backbone)) {
			for i, roomLn := range []data.Line{a, b, c, d} {
				if data.Intersects(ln, roomLn) {
					side := []Side{top, right, btm, left}[i]
					if !hasSide(adj[room], side) {
						adj[room] = append(adj[room], side)
					}
				}
			}
		}
	}
	return adj
}

func InsertHallway(roomMap map[*data.Room][]Side, width float64) {
	for room, sides := range roomMap {
		if room.Rect.Label == "Living" {
			continue
		}
		for _, side := range sides {
			switch side {
			case top:
				room.Rect.SetHeightTop(room.Rect.Height() - width/2.0)
				if len(room.Doors) == 0 {
					room.Doors = append(room.Doors, &data.Door{data.N, 0.50})
				}
			case right:
				room.Rect.SetWidthR(room.Rect.Width() - width/2.0)
				if len(room.Doors) == 0 {
					room.Doors = append(room.Doors, &data.Door{data.E, 0.50})
				}
			case btm:
				room.Rect.SetHeightBtm(room.Rect.Height() - width/2.0)
				if len(room.Doors) == 0 {
					room.Doors = append(room.Doors, &data.Door{data.S, 0.50})
				}
			case left:
				room.Rect.SetWidthL(room.Rect.Width() - width/2.0)
				if len(room.Doors) == 0 {
					room.Doors = append(room.Doors, &data.Door{data.W, 0.50})
				}
			}
		}
	}
}
