package dataStructures

import (
	"testing"
)

func GetTree() RTree {
	// Visual reference:
	//  2         + - - - +
	//            |   E   |
	//  1 + - + - + - - - +
	//    | A | B |   D   |
	//  0 + - + - + - - - +
	//    0   1   2   3   4
	a := RTree{Rect{X0: 0, Y0: 0, X1: 1, Y1: 1, Label: "A"}, nil}
	b := RTree{Rect{X0: 1, Y0: 0, X1: 2, Y1: 1, Label: "B"}, nil}
	c := RTree{
		Rect{X0: 2, Y0: 0, X1: 4, Y1: 2, Label: "C"},
		[]*RTree{
			&RTree{Rect{X0: 2, Y0: 0, X1: 4, Y1: 1, Label: "D"}, nil},
			&RTree{Rect{X0: 2, Y0: 1, X1: 4, Y1: 2, Label: "E"}, nil},
		},
	}
	return RTree{
		Rect{X0: 0, Y0: 0, X1: 4, Y1: 2, Label: "Root"},
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

func Test_FindNearestPt(t *testing.T) {
	tree := GetTree()
	insideA := Pt{X: 0.5, Y: 0.5}
	insideB := Pt{X: 1.5, Y: 0.5}
	insideD := Pt{X: 2.5, Y: 0.5}
	insideE := Pt{X: 3.5, Y: 1.5}
	aTwixtB := Pt{X: 1.0, Y: 0.5}
	bTwixtD := Pt{X: 2.0, Y: 0.5}
	bdeCrnr := Pt{X: 2.0, Y: 1.0}
	dTwixtE := Pt{X: 3.0, Y: 1.0}

	// There was probably a less boiler-platey way of writing these tests, but
	// vim macros are both a blessing and a curse in that they make this kind
	// of boilerplate easy.
	found := len(tree.FindNearestPt(insideA))
	if found != 1 {
		t.Errorf(
			"Expected 1 value for pt (%.2f, %.2f), but got %d",
			insideA.X, insideA.Y, found,
		)
	}
	found = len(tree.FindNearestPt(insideB))
	if found != 1 {
		t.Errorf(
			"Expected 1 value for pt (%.2f, %.2f), but got %d",
			insideB.X, insideB.Y, found,
		)
	}
	found = len(tree.FindNearestPt(insideD))
	if found != 1 {
		t.Errorf(
			"Expected 1 value for pt (%.2f, %.2f), but got %d",
			insideD.X, insideD.Y, found,
		)
	}
	found = len(tree.FindNearestPt(insideE))
	if found != 1 {
		t.Errorf(
			"Expected 1 value for pt (%.2f, %.2f), but got %d",
			insideE.X, insideE.Y, found,
		)
	}
	found = len(tree.FindNearestPt(aTwixtB))
	if found != 2 {
		t.Errorf(
			"Expected 2 values for pt (%.2f, %.2f), but got %d",
			aTwixtB.X, aTwixtB.Y, found,
		)
	}
	found = len(tree.FindNearestPt(bTwixtD))
	if found != 2 {
		t.Errorf(
			"Expected 2 values for pt (%.2f, %.2f), but got %d",
			bTwixtD.X, bTwixtD.Y, found,
		)
	}
	found = len(tree.FindNearestPt(dTwixtE))
	if found != 2 {
		t.Errorf(
			"Expected 2 values for pt (%.2f, %.2f), but got %d",
			dTwixtE.X, dTwixtE.Y, found,
		)
	}
	found = len(tree.FindNearestPt(bdeCrnr))
	if found != 3 {
		t.Errorf(
			"Expected 3 values for pt (%.2f, %.2f), but got %d",
			bdeCrnr.X, bdeCrnr.Y, found,
		)
	}
}
