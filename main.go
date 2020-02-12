package main

import (
	"fmt"
	data "housify/dataStructures"
	"math/rand"
	"time"
)

const (
	scaleAmt = 10.0
)

// for debugging purposes
func TestTree() (data.Rect, *data.FTree) {
	bounds := data.Rect{X0: 0, Y0: 0, X1: 60, Y1: 40, Label: "root"}
	a := data.FTree{Value: 60, Label: "Living", Cnx: nil}
	b := data.FTree{Value: 60, Label: "Kitchen", Cnx: nil}
	c := data.FTree{Value: 40, Label: "Master Bed", Cnx: nil}
	d := data.FTree{Value: 30, Label: "Bed", Cnx: nil}
	e := data.FTree{Value: 20, Label: "Bath", Cnx: nil}
	f := data.FTree{Value: 20, Label: "Laundry", Cnx: nil}
	g := data.FTree{Value: 10, Label: "Closet", Cnx: nil}
	areas := data.FTree{
		Value: 60 * 40,
		Label: "Root",
		Cnx:   []*data.FTree{&a, &b, &c, &d, &e, &f, &g},
	}
	return bounds, &areas
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	bounds, areas := GenHouse("data/room_edges.csv", "data/room_sizes.csv")
	squarified := Squarify(bounds, areas)
	squarified.Quantize(0) // prevents floating point errors
	backbone := Backbone(bounds, squarified)
	rooms := data.RectsToRooms(data.RTreesToRects(squarified.Leafs()))

	roomEdges, err := loadRoomEdges("data/room_edges.csv")
	if err != nil {
		panic(err)
	}
	SetAdjacentRooms(rooms)
	AddDoorsBetweenRooms(rooms, roomEdges)
	InsertHallway(RoomsAdjToBackbone(rooms, backbone), 6)

	Draw(bounds, rooms)
	fmt.Println("done!")
}
