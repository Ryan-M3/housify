package main

import (
	data "housify/dataStructures"
)

// Intersection of two lists of rectangles.
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

func hasDir(sides []data.Dir, query data.Dir) bool {
	for _, side := range sides {
		if side == query {
			return true
		}
	}
	return false
}

// Return a list of rooms which are adjacent to the provided backbone.
func RoomsAdjToBackbone(house []*data.Room, backbone data.Graph) map[*data.Room][]data.Dir {
	adj := make(map[*data.Room][]data.Dir, 0)
	for _, room := range house {
		a, b, c, d := data.RectToLines(room.Rect)
		for _, ln := range unique(data.GraphToLines(backbone)) {
			for i, roomLn := range []data.Line{a, b, c, d} {
				if data.Intersects(ln, roomLn) {
					side := []data.Dir{data.N, data.E, data.S, data.W}[i]
					if !hasDir(adj[room], side) {
						adj[room] = append(adj[room], side)
					}
				}
			}
		}
	}
	return adj
}

// List of all doors in a room or a room connected to that room.
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

// Whether or not a room has a door that isn't to an adjacent room (i.e. to
// outside or to a hallway).
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

// Resize rooms so that a hallway can fit inside it.
func InsertHallway(roomMap map[*data.Room][]data.Dir, width float64) {
	for room, sides := range roomMap {
		if room.Rect.Label == "Living" {
			continue
		}
		for _, side := range sides {
			switch side {
			case data.N:
				room.Rect.SetHeightTop(room.Rect.Height() - width/2.0)
				if !hasEscape(room) && len(room.Doors) < 2 {
					room.Doors = append(room.Doors, &data.Door{data.N, 0.50})
				}
			case data.E:
				room.Rect.SetWidthR(room.Rect.Width() - width/2.0)
				if !hasEscape(room) && len(room.Doors) < 2 {
					room.Doors = append(room.Doors, &data.Door{data.E, 0.50})
				}
			case data.S:
				room.Rect.SetHeightBtm(room.Rect.Height() - width/2.0)
				if !hasEscape(room) && len(room.Doors) < 2 {
					room.Doors = append(room.Doors, &data.Door{data.S, 0.50})
				}
			case data.W:
				room.Rect.SetWidthL(room.Rect.Width() - width/2.0)
				if !hasEscape(room) && len(room.Doors) < 2 {
					room.Doors = append(room.Doors, &data.Door{data.W, 0.50})
				}
			}
		}
	}
}
