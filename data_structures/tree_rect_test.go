package data_structures

import (
	"testing"
)

func GetTree() RTree {
	a := RTree{Rect{X0: 0, Y0: 0, X1: 0, Y1: 0, Label: "A"}, nil}
	b := RTree{Rect{X0: 0, Y0: 0, X1: 0, Y1: 0, Label: "B"}, nil}
	c := RTree{
		Rect{X0: 0, Y0: 0, X1: 0, Y1: 0, Label: "C"},
		[]*RTree{
			&RTree{Rect{X0: 0, Y0: 0, X1: 0, Y1: 0, Label: "D"}, nil},
			&RTree{Rect{X0: 0, Y0: 0, X1: 0, Y1: 0, Label: "E"}, nil},
		},
	}
	return RTree{
		Rect{X0: 0, Y0: 0, X1: 0, Y1: 0, Label: "Root"},
		[]*RTree{&a, &b, &c},
	}
}

func Test_List(t *testing.T) {
	root := GetTree()
	nodes := root.List()
	if len(nodes) != 6 {
		t.Errorf("expected 6 items, got %d", len(nodes))
	}
}

func Test_Leafs(t *testing.T) {
	tree := GetTree()
	N := len(tree.Leafs())
	if N != 4 {
		t.Errorf("Expected len(tree.List()) == 4, got %d", N)
	}
}

func Test_Find(t *testing.T) {
	root := GetTree()
	found, ok := root.Find("E")
	if !ok || found.Value.Label != "E" {
		t.Errorf("Expected node labelled \"E\", but got %s instead.", found.Value.Label)
	}
}

func Test_CnxValues(t *testing.T) {
	tree := GetTree()
	N := len(tree.CnxValues())
	if N != 3 {
		t.Errorf("Expected %d, but got %d.", 3, N)
	}
}
