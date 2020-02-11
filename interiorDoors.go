package main

import (
	data "housify/dataStructures"
	"sort"
)

type lineAndRoom struct {
	line data.Line
	room *data.Room
}

func toLineAndRoom(rooms []*data.Room) []lineAndRoom {
	var output []lineAndRoom
	for _, r := range rooms {
		a, b, c, d := data.RectToLines(r.Rect)
		sides := []data.Line{a, b, c, d}
		for _, side := range sides {
			output = append(output, lineAndRoom{line: side, room: r})
		}
	}
	return output
}

func hasRoom(rooms []*data.Room, query *data.Room) bool {
	for _, room := range rooms {
		if room == query {
			return true
		}
	}
	return false
}

// Find overlapping lines from a list of lines.
func Sweep(line data.Line, rooms []*data.Room) []*data.Room {
	var output []*data.Room
	lars := toLineAndRoom(rooms)
	sort.SliceStable(lars, func(i, j int) bool {
		a, _ := data.LineToPt(lars[i].line)
		b, _ := data.LineToPt(lars[j].line)
		if a.X == b.X {
			return a.Y < b.Y
		}
		return a.X < b.X
	})
	for _, lar := range lars {
		if !hasRoom(output, lar.room) {
			if data.Intersects(lar.line, line) || line == lar.line {
				output = append(output, lar.room)
			}
		}
	}
	sort.SliceStable(lars, func(i, j int) bool {
		a, _ := data.LineToPt(lars[i].line)
		b, _ := data.LineToPt(lars[j].line)
		if a.Y == b.Y {
			return a.X < b.X
		}
		return a.Y < b.Y
	})
	for _, lar := range lars {
		if !hasRoom(output, lar.room) {
			if data.Intersects(lar.line, line) || line == lar.line {
				output = append(output, lar.room)
			}
		}
	}
	return output
}

func SetAdjacentRooms(rooms []*data.Room) {
	for _, room := range rooms {
		top, right, btm, left := data.RectToLines(room.Rect)
		for _, side := range []data.Line{top, right, btm, left} {
			for _, cnx := range Sweep(side, rooms) {
				if cnx.Rect == room.Rect {
					continue
				}
				if !hasRoom(room.Adj, cnx) {
					room.Adj = append(room.Adj, cnx)
				}
			}
		}
	}
}

// Where is B relative to A? Returns one of N, E, S, or W.
func RoomOrientation(a, b *data.Room) data.CardinalDirection {
	a0, _, a1, _ := data.RectToPts(*a.Rect)
	b0, _, b1, _ := data.RectToPts(*b.Rect)
	if a1.Y == b0.Y {
		return data.S
	} else if a1.X == b0.X {
		return data.W
	} else if a0.Y == b1.Y {
		return data.N
	} else {
		return data.E
	}
}

func AddDoor(a, b *data.Room) {
	door := data.Door{Orientation: RoomOrientation(a, b), Position: 0.5}
	a.Doors = append(a.Doors, &door)
	b.Doors = append(b.Doors, &door)
}

func AddDoorsBetweenRooms(rooms []*data.Room, roomEdges map[string][]string) {
	for _, room := range rooms {
		added := make(map[string][]string, 0)
		for _, neighbor := range room.Adj {
			if data.HasString(roomEdges[room.Rect.Label], neighbor.Rect.Label) &&
				!data.HasString(added[room.Rect.Label], neighbor.Rect.Label) &&
				!data.HasString(added[neighbor.Rect.Label], room.Rect.Label) {
				AddDoor(room, neighbor)
				added[room.Rect.Label] = append(added[room.Rect.Label], neighbor.Rect.Label)
			}
		}
	}
}
