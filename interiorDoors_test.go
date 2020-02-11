package main

import (
	data "housify/dataStructures"
	"testing"
)

func Test_AddDoorsBetweenRooms(t *testing.T) {
	bounds, areas := TestTree()
	squarified := Squarify(bounds, areas)
	squarified.Quantize(0)
	rooms := data.RectsToRooms(data.RTreesToRects(squarified.Leafs()))
	roomEdges, err := loadRoomEdges("data/room_edges.csv")
	if err != nil {
		t.Error(err)
	}
	SetAdjacentRooms(rooms)
	AddDoorsBetweenRooms(rooms, roomEdges)
}

func Test_SetAdjRooms(t *testing.T) {
	bounds, areas := TestTree()
	squarified := Squarify(bounds, areas)
	squarified.Quantize(0)
	rooms := data.RectsToRooms(data.RTreesToRects(squarified.Leafs()))
	SetAdjacentRooms(rooms)
	for _, room := range rooms {
		msg := "len(%s.Adj) should be %d, but got %d"
		switch room.Rect.Label {
		case "Living":
			if len(room.Adj) != 2 {
				t.Errorf(msg, room.Rect.Label, 2, len(room.Adj))
			}
		case "Kitchen":
			if len(room.Adj) != 3 {
				t.Errorf(msg, room.Rect.Label, 3, len(room.Adj))
			}
		case "Master Bed":
			if len(room.Adj) != 5 {
				t.Errorf(msg, room.Rect.Label, 5, len(room.Adj))
			}
		case "Bed":
			if len(room.Adj) != 3 {
				t.Errorf(msg, room.Rect.Label, 3, len(room.Adj))
			}
		case "Bath":
			if len(room.Adj) != 3 {
				t.Errorf(msg, room.Rect.Label, 3, len(room.Adj))
			}
		case "Laundry":
			if len(room.Adj) != 4 {
				t.Errorf(msg, room.Rect.Label, 4, len(room.Adj))
			}
		case "Closet":
			if len(room.Adj) != 2 {
				t.Errorf(msg, room.Rect.Label, 2, len(room.Adj))
			}
		}
	}
}
