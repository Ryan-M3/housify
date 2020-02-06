package data_structures

import (
	"fmt"
	"math"
	"testing"
	"testing/quick"
)

const (
	maxNum = 1000000000 // 1 billion is absurdly large for our purposes
)

func Test_CollinearHorz(t *testing.T) {
	fn := func(x, y, w1, w2, spacing float64) bool {
		left := Line{x, y, x + w1, y}
		right := Line{left.x1 + spacing, y, left.x1 + spacing + w2, y}
		return Horz(left) && Horz(right) && Collinear(left, right)
	}

	if err := quick.Check(fn, nil); err != nil {
		t.Error(err)
	}
}

func Test_CollinearVert(t *testing.T) {
	fn := func(x, y, h1, h2, spacing float64) bool {
		left := Line{x, y, x, y + h1}
		right := Line{x, left.y1 + spacing, x, left.y1 + spacing + h2}
		return Vert(left) && Vert(right) && Collinear(left, right)
	}

	if err := quick.Check(fn, nil); err != nil {
		t.Error(err)
	}
}

func Test_Horz(t *testing.T) {
	fn := func(x, y, w1, w2, spacing float64) bool {
		x = math.Abs(math.Mod(math.Round(x), maxNum))
		y = math.Abs(math.Mod(math.Round(y), maxNum))
		w1 = math.Abs(math.Mod(math.Round(w1), maxNum))
		w2 = math.Abs(math.Mod(math.Round(w2), maxNum))
		spacing = math.Abs(math.Mod(math.Round(spacing), maxNum))
		left := Line{x, y, x + w1, y}
		right := Line{left.x1 + spacing, y, left.x1 + spacing + w2, y}
		top := Line{x, y, x, y + w1}
		btm := Line{x, top.y1 + spacing, x, top.y1 + spacing + w2}
		return Horz(left) && Horz(right) && !Horz(top) && !Horz(btm)
	}

	if err := quick.Check(fn, nil); err != nil {
		t.Error(err)
	}
}

func Test_Vert(t *testing.T) {
	fn := func(x, y, w1, w2, spacing float64) bool {
		x = math.Abs(math.Mod(math.Round(x), maxNum))
		y = math.Abs(math.Mod(math.Round(y), maxNum))
		w1 = math.Abs(math.Mod(math.Round(w1), maxNum))
		w2 = math.Abs(math.Mod(math.Round(w2), maxNum))
		spacing = math.Abs(math.Mod(math.Round(spacing), maxNum))
		left := Line{x, y, x + w1, y}
		right := Line{left.x1 + spacing, y, left.x1 + spacing + w2, y}
		top := Line{x, y, x, y + w1}
		btm := Line{x, top.y1 + spacing, x, top.y1 + spacing + w2}
		return !Vert(left) && !Vert(right) &&
			Vert(top) && Vert(btm)
	}

	if err := quick.Check(fn, nil); err != nil {
		t.Error(err)
	}
}

func Test_Exterior(t *testing.T) {
	fn := func(x, y, w, h float64) bool {
		x = math.Abs(math.Mod(math.Round(x), maxNum))
		y = math.Abs(math.Mod(math.Round(y), maxNum))
		w = math.Abs(math.Mod(math.Round(w), maxNum))
		h = math.Max(2, math.Abs(math.Mod(math.Round(h), maxNum)))
		r := Rect{X0: x, Y0: y, X1: x + w, Y1: y + h, Label: ""}
		ln1 := Line{
			x0: x + 0.1*w,
			y0: y,
			x1: x + 0.9*w,
			y1: y,
		}
		ln2 := Line{
			x0: x + 0.1*w,
			y0: y + 1,
			x1: x + 0.9*w,
			y1: y + 1,
		}
		return InPerimeter(r, ln1) && !InPerimeter(r, ln2)
	}

	if err := quick.Check(fn, nil); err != nil {
		t.Error(err)
	}
}

func Test_LinesVert(t *testing.T) {
	checkLine := func(ln Line, a, b, x, y float64) bool {
		return ln.x0 == a && ln.y0 == b && ln.x1 == x && ln.y1 == y
	}
	r := &Rect{0, 0, 1, 1, ""}
	_, right, _, left := RectToLines(r)
	if !checkLine(left, 0, 0, 0, 1) {
		t.Errorf("Got %v, expected 0, 0, 0, 1 rectangle.", left)
	}
	if !checkLine(right, 1, 0, 1, 1) {
		t.Errorf("Got %v, expected 1, 0, 1, 1 rectangle.", right)
	}
}

func Test_LinesHorz(t *testing.T) {
	checkLine := func(ln Line, a, b, x, y float64) bool {
		return ln.x0 == a && ln.y0 == b && ln.x1 == x && ln.y1 == y
	}
	r := &Rect{X0: 0, Y0: 0, X1: 1, Y1: 1, Label: ""}
	top, _, btm, _ := RectToLines(r)
	if !checkLine(btm, 0, 0, 1, 0) {
		t.Errorf("Got %v, expected 0, 0, 0, 1 rectangle.", btm)
	}
	if !checkLine(top, 0, 1, 1, 1) {
		t.Errorf("Got %v, expected 1, 0, 1, 1 rectangle.", top)
	}
}

// Also tests Intersects() for horizontal intputs.
func Test_IntersectsHorz(t *testing.T) {
	ln1 := Line{0, 0, 10, 0}
	ln2 := Line{11, 0, 110, 0}
	ln3 := Line{5, 0, 12, 0}

	if IntersectsHorz(ln1, ln2) {
		t.Errorf("%v and %v don't intersect, but IntersectHorz says they do", ln1, ln2)
	}
	if !IntersectsHorz(ln1, ln3) {
		t.Errorf("%v and %v intersect, but IntersectsHorz says they don't", ln1, ln3)
	}
	if !IntersectsHorz(ln2, ln3) {
		t.Errorf("%v and %v intersect, but IntersectsHorz says they don't", ln2, ln3)
	}

	if Intersects(ln1, ln2) {
		t.Errorf("%v and %v don't intersect, but Intersect says they do", ln1, ln2)
	}
	if !Intersects(ln1, ln3) {
		t.Errorf("%v and %v intersect, but Intersects says they don't", ln1, ln3)
	}
	if !Intersects(ln2, ln3) {
		t.Errorf("%v and %v intersect, but Intersects says they don't", ln2, ln3)
	}
}

// Also tests Intersects() for vertical inputs.
func Test_IntersectsVert(t *testing.T) {
	ln1 := Line{0, 0, 0, 10}
	ln2 := Line{0, 11, 0, 110}
	ln3 := Line{0, 5, 0, 12}

	if IntersectsVert(ln1, ln2) {
		t.Errorf("%v and %v don't intersect, but IntersectVert says they do", ln1, ln2)
	}
	if !IntersectsVert(ln1, ln3) {
		t.Errorf("%v and %v intersect, but IntersectsVert says they don't", ln1, ln3)
	}
	if !IntersectsVert(ln2, ln3) {
		t.Errorf("%v and %v intersect, but IntersectsVert says they don't", ln2, ln3)
	}

	if Intersects(ln1, ln2) {
		t.Errorf("%v and %v don't intersect, but Intersect says they do", ln1, ln2)
	}
	if !Intersects(ln1, ln3) {
		t.Errorf("%v and %v intersect, but Intersects says they don't", ln1, ln3)
	}
	if !Intersects(ln2, ln3) {
		t.Errorf("%v and %v intersect, but Intersects says they don't", ln2, ln3)
	}
}

func VerifyLine(e Line, a, b, x, y float64) bool {
	return e.x0 == a && e.y0 == b && e.x1 == x && e.y1 == y
}

func Test_ResegmentSingle(t *testing.T) {
	left := Line{x0: 0, y0: 0, x1: 10, y1: 0}
	right := Line{x0: 5, y0: 0, x1: 15, y1: 0}
	newLines := ResegmentLine(left, right)
	beg, mid, end := newLines[0], newLines[1], newLines[2]
	if !VerifyLine(beg, 0, 0, 5, 0) {
		t.Errorf("Expected 0, 0, 5, 0 rect, got %v", beg)
	}
	if !VerifyLine(mid, 5, 0, 10, 0) {
		t.Errorf("Expected 5, 0, 10, 0 rect, got %v", beg)
	}
	if !VerifyLine(end, 10, 0, 15, 0) {
		t.Errorf("Expected 10, 0, 15, 0 rect, got %v", beg)
	}
}

func Test_Resegment(t *testing.T) {
	// +---+---+---+---+
	// 0   1   2   3   4
	// +-+---+
	// +---+
	// +-+-+-+
	input := []Line{
		// first rect
		{0, 0, 0, 1},
		{0, 0.5, 0, 0.85}, // overlaps
		{1, 0, 1, 1},
		{0, 0, 1, 0},
		{0, 0, 1.5, 0}, // overlaps
		{0, 1, 1, 1},
		// second rect
		{1, 1, 1, 2},
		{2, 1, 2, 2},
		{1, 1, 2, 1},
		{1, 2, 2, 2},
		// third rect
		{2, 2, 2, 3},
		{3, 2, 3, 3},
		{2, 2, 3, 2},
		{2, 3, 3, 3},
		// fourth rect
		{0, 0, 0, 2}, // overlaps
		{3, 0, 3, 2},
		{0, 0, 3, 0},
		{0, 2, 3, 2},
	}
	output := FilterLines(ResegmentLines(input))
	if len(output) != len(input)+3 {
		t.Errorf("Expected len(output) == %d, got %d", len(input)+3, len(output))
		for _, ln := range output {
			fmt.Println(ln)
		}
	}
}

func Test_Inside(t *testing.T) {
	outer := Line{x0: 0, y0: 0, x1: 10, y1: 0}
	inner := Line{x0: 1, y0: 0, x1: 9, y1: 0}
	if !Inside(outer, inner) {
		t.Errorf("Expected true in Inside(%v, %v), but got false.", outer, inner)
	}
	farAway := Line{x0: 100, y0: 0, x1: 900, y1: 0}
	if Inside(outer, farAway) {
		t.Errorf("Expected false in LineBoundBy(%v, %v), but got false.", outer, farAway)
	}
}

func AbsI(x int64) int64 {
	if x < -x {
		return -x
	}
	return x
}
