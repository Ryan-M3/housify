package data_structures

import (
	"testing"
)

func Test_Add(t *testing.T) {
	a := Pt{0, 0}
	b := Pt{1, 2}
	c := Pt{42, 123}
	g := make(Graph, 0)
	Add(g, a, b)
	if !HasPt(g[a], b) {
		t.Errorf("Failed to add point to graph with Add function.")
	}
	if len(g[a]) != 1 {
		t.Errorf("Incorrect length of graph: got %d, but expected 1", len(g[a]))
	}
	Add(g, a, c)
	if !HasPt(g[a], c) {
		t.Errorf("Failed to add point to graph with Add function.")
	}
	if len(g[a]) != 2 {
		t.Errorf("Incorrect length of graph: got %d, but expected 2", len(g[a]))
	}
	Add(g, c, a)
	if !HasPt(g[c], a) {
		t.Errorf("Failed to add point to graph with Add function.")
	}
	if len(g[c]) != 1 {
		t.Errorf("Incorrect length of graph: got %d, but expected 1", len(g[c]))
	}
}

func Test_LinesToGraph(t *testing.T) {
	// For visual reference:
	// 3      +
	//        |
	// 2      +---------+----+
	//        |         |    |
	// 1 +----+         |    |
	//        |         |    |
	// 0      +         +    +
	//   0    1    2    3    4    5
	got := LinesToGraph([]Line{
		{0, 1, 1, 1},
		{1, 1, 1, 0},
		{1, 1, 1, 2},
		{1, 2, 1, 3},
		{1, 2, 3, 2},
		{3, 2, 3, 0},
		{3, 2, 4, 2},
		{4, 2, 4, 0},
	})
	want := Graph{
		Pt{0, 1}: []Pt{Pt{1, 1}},
		Pt{1, 1}: []Pt{Pt{1, 0}, Pt{1, 2}, Pt{0, 1}},
		Pt{1, 2}: []Pt{Pt{1, 3}, Pt{3, 2}, Pt{1, 1}},
		Pt{3, 2}: []Pt{Pt{3, 0}, Pt{4, 2}, Pt{1, 2}},
		Pt{4, 2}: []Pt{Pt{4, 0}, Pt{3, 2}},
	}
	for k, vs := range want {
		for _, v := range vs {
			if !HasPt(got[k], v) {
				t.Errorf("%v missing expected value %v", got[k], v)
			}
		}
		if len(got[k]) != len(vs) {
			t.Errorf("Mismatch in number of entries between got (%v) and want (%v)", got[k], vs)
		}
	}
}

func Test_PathsToGraph(t *testing.T) {
	a := Pt{X: 0, Y: 0}
	b := Pt{X: 1, Y: 0}
	c := Pt{X: 2, Y: 0}
	d := Pt{X: -1, Y: 1}
	path1 := []Pt{a, b, c}
	path2 := []Pt{a, b, d}
	paths := [][]Pt{path1, path2}
	want := Graph{a: []Pt{b}, b: []Pt{c, d}}
	got := PathsToGraph(paths)
	if len(got) != len(want) {
		t.Errorf("Expected %v, but got %v", want, got)
	}
	if got[a][0] != want[a][0] || got[b][0] != want[b][0] || got[b][1] != want[b][1] {
		t.Errorf("Expected %v, but got %v", want, got)
	}
}
