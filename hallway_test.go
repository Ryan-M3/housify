package main

import (
	data "housify/dataStructures"
	"testing"
)

func Test_HasEscape(t *testing.T) {
	door1 := data.Door{Orientation: data.E, Position: 0.5}
	door2 := data.Door{Orientation: data.N, Position: 0.5}
	//door3 := data.Door{data.S, 0.5}
	r0 := data.Room{
		Rect:  &data.Rect{0, 0, 1, 1, "a"},
		Doors: []*data.Door{&door1},
		Adj: []*data.Room{
			&data.Room{
				Rect:  &data.Rect{1, 0, 2, 1, "b"},
				Doors: []*data.Door{&door1, &door2},
				Adj:   nil,
			},
		},
	}
	if !hasEscape(&r0) {
		t.Errorf("Room has an escape, but HasEscape returns false.")
	}
	r0.Doors = append(r0.Doors, &door2)
	if hasEscape(&r0) {
		t.Errorf("Room doesn't have an escape, but HasEscape returns true.")
	}
}
