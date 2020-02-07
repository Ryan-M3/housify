package main

import (
	data "housify/data_structures"
	"testing"
)

func Test_AStar(t *testing.T) {
	g := data.Graph{
		data.Pt{X: 0, Y: 0}: []data.Pt{data.Pt{X: 0, Y: 1}, data.Pt{X: 1, Y: 0}},
		data.Pt{X: 0, Y: 1}: []data.Pt{data.Pt{X: 1, Y: 1}, data.Pt{X: -1, Y: 1}},
		data.Pt{X: 1, Y: 0}: []data.Pt{data.Pt{X: 4, Y: 4}, data.Pt{X: 2, Y: 0}},
	}
	want := []data.Pt{
		data.Pt{X: 0, Y: 0},
		data.Pt{X: 1, Y: 0},
		data.Pt{X: 2, Y: 0},
	}
	m := make(map[data.Pt]bool, 0)
	if got := AStar([]data.Pt{data.Pt{X: 0, Y: 0}}, data.pt{data.Pt{X: 2, Y: 0}}, g, m); ok {
		if len(got) != len(want) {
			t.Errorf("Not enough points returned; expected %v, but got %v", want, got)
		}
		for i := 0; i < len(got); i++ {
			if got[i] != want[i] {
				t.Errorf("Mismatch at index %d in want %v, and got %v", i, want, got)
			}
		}
	} else {
		t.Errorf("AStar could not find path.")
	}
}

func Test_AStars(t *testing.T) {
	g := data.Graph{
		data.Pt{X: 0, Y: 0}: []data.Pt{data.Pt{X: 0, Y: 1}, data.Pt{X: 1, Y: 0}},
		data.Pt{X: 0, Y: 1}: []data.Pt{data.Pt{X: 1, Y: 1}, data.Pt{X: -1, Y: 1}},
		data.Pt{X: 1, Y: 0}: []data.Pt{data.Pt{X: 4, Y: 4}, data.Pt{X: 2, Y: 0}},
	}
	src := data.Pt{X: 0, Y: 0}
	tgt1 := data.Pt{X: 2, Y: 0}
	tgt2 := data.Pt{X: -1, Y: 1}
	tgts := []data.Pt{tgt1, tgt2}
	path1 := []data.Pt{
		data.Pt{X: 0, Y: 0},
		data.Pt{X: 1, Y: 0},
		data.Pt{X: 2, Y: 0},
	}
	path2 := []data.Pt{
		data.Pt{X: 0, Y: 0},
		data.Pt{X: 0, Y: 1},
		data.Pt{X: -1, Y: 1},
	}
	want := [][]data.Pt{path1, path2}
	got := AStars(src, tgts, g)
	if len(got) != len(want) {
		t.Errorf("Not enough points returned; expected %v, but got %v", want, got)
	}
	for i := 0; i < len(got); i++ {
		for j := 0; j < len(got[i]); j++ {
			if got[i][j] != want[i][j] {
				t.Errorf("Mismatch at index %d, %d in want %v, and got %v", i, j, want, got)
			}
		}
	}
}
