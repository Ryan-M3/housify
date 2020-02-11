package main

import (
	data "housify/dataStructures"
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

func getDoors(room *data.Room, visited map[*data.Room]bool) ([]*data.Door, map[*data.Room]bool) {
	visited[room] = true
	doors := room.Doors
	for _, adj := range room.Adj {
		if _, ok := visited[adj]; !ok {
			newDoors, _ := getDoors(adj, visited)
			doors = append(doors, newDoors...)
		}
	}
	return doors, visited
}

func hasEscape(room *data.Room) bool {
	// If an adjacent room has a door which is not in any other adjacent room,
	// then that room has access to a hallway or outside.
	doors, _ := getDoors(room, make(map[*data.Room]bool, 0))
	doorCount := make(map[*data.Door]int, 0)
	for _, door := range doors {
		doorCount[door]++
	}
	for _, v := range doorCount {
		if v == 1 {
			return true
		}
	}
	return false
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
				if !hasEscape(room) && len(room.Doors) < 2 {
					room.Doors = append(room.Doors, &data.Door{data.N, 0.50})
				}
			case right:
				room.Rect.SetWidthR(room.Rect.Width() - width/2.0)
				if !hasEscape(room) && len(room.Doors) < 2 {
					room.Doors = append(room.Doors, &data.Door{data.E, 0.50})
				}
			case btm:
				room.Rect.SetHeightBtm(room.Rect.Height() - width/2.0)
				if !hasEscape(room) && len(room.Doors) < 2 {
					room.Doors = append(room.Doors, &data.Door{data.S, 0.50})
				}
			case left:
				room.Rect.SetWidthL(room.Rect.Width() - width/2.0)
				if !hasEscape(room) && len(room.Doors) < 2 {
					room.Doors = append(room.Doors, &data.Door{data.W, 0.50})
				}
			}
		}
	}
}
