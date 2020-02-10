package dataStructures

import (
	"math"
)

type Rect struct {
	X0, Y0, X1, Y1 float64
	Label          string
}

/////////////
// Getters //
/////////////

func (r *Rect) Width() float64 {
	return math.Abs(r.X1 - r.X0)
}

func (r *Rect) Height() float64 {
	return math.Abs(r.Y1 - r.Y0)
}

func (r *Rect) Area() float64 {
	return r.Width() * r.Height()
}

func (r *Rect) Center() Pt {
	return Pt{(r.X0 + r.X1) / 2, (r.Y0 + r.Y1) / 2}
}

func (r *Rect) Contains(pt Pt) bool {
	return r.X0 <= pt.X && pt.X <= r.X1 &&
		r.Y0 <= pt.Y && pt.Y <= r.Y1
}

/////////////
// Setters //
/////////////

func (r *Rect) Scale(amt float64) {
	r.X0 *= amt
	r.Y0 *= amt
	r.X1 *= amt
	r.Y1 *= amt
}

// Set the width to w by moving the right edge of the rectangle.
func (r *Rect) SetWidthR(w float64) {
	r.X1 = r.X0 + w
}

// Set the width to w by moving the left edge of the rectangle.
func (r *Rect) SetWidthL(w float64) {
	r.X0 = r.X1 - w
}

// Set the height to h by moving the top edge down.
func (r *Rect) SetHeightTop(h float64) {
	r.Y1 = r.Y0 + h
}

// Set the height to h by moving the bottom edge up.
func (r *Rect) SetHeightBtm(h float64) {
	r.Y0 = r.Y1 - h
}

// Move the rectangle so that the lower left hand corner has the given x and
// y coordinates.
func (r *Rect) MoveTo(x, y float64) {
	r.X1 = x + r.Width()
	r.Y1 = y + r.Height()
	r.X0 = x
	r.Y0 = y
}

func (r *Rect) AboveLine(ln Line) bool {
	//          y1
	//  + - - - +
	//  |       |
	//  + - - - +
	// y0
	//  +---e---+  e.y
	return r.Y0 >= ln.y1 && r.Y1 > ln.y1
}

func (r *Rect) BelowLine(ln Line) bool {
	return !r.AboveLine(ln)
}

func (r *Rect) RightOfLine(ln Line) bool {
	//    r.x0    r.x1
	//  +  + - - - +
	//  |  |       |
	//  +  + - - - +
	// e.x
	return r.X0 >= ln.x1 && r.X1 > ln.x1
}

func (r *Rect) LeftOfLine(ln Line) bool {
	return !r.RightOfLine(ln)
}

//////////////
// Stackers //
//////////////

// Stack a list of rooms one on top of the other.
func colStack(rs []*Rect) {
	x, y := rs[0].X0, rs[0].Y1 // the first room isn't moved
	for _, r := range rs[1:] {
		r.MoveTo(x, y)
		x = r.X0
		y = r.Y1
	}
}

// Stack a list of rooms one to the right of the last, like books on a
// bookshelf.
func rowStack(rs []*Rect) {
	x, y := rs[0].X1, rs[0].Y0
	for _, r := range rs[1:] {
		r.MoveTo(x, y)
		x = r.X1
		y = r.Y0
	}
}

// Sum a list of floats.
func sum(row []float64) float64 {
	s := 0.0
	for _, n := range row {
		s += n
	}
	return s
}

// Create rectangles and stack them into a column, resizing the rectangles if
// needed to fit them into that column perfectly.
func ColStackInto(areas *FTree, shelf *Rect) []*Rect {
	areas.scaleAreas(shelf.Height())
	var output []*Rect
	for _, branch := range areas.Cnx {
		r := *shelf
		r.SetWidthR(shelf.Width())
		r.SetHeightTop(branch.Value)
		r.Label = branch.Label
		output = append(output, &r)
	}
	output[0].MoveTo(shelf.X0, shelf.Y0)
	colStack(output)
	return output
}

// Create rooms and stack them like books into a bookshelf, resizing the room
// as needed to fit them into that shelf perfectly.
func RowStackInto(areas *FTree, shelf *Rect) []*Rect {
	areas.scaleAreas(shelf.Width())
	var output []*Rect
	for _, branch := range areas.Cnx {
		r := *shelf
		r.SetHeightTop(shelf.Height())
		r.SetWidthR(branch.Value)
		r.Label = branch.Label
		output = append(output, &r)
	}
	output[0].MoveTo(shelf.X0, shelf.Y0)
	rowStack(output)
	return output
}

// Create rooms with the given areas stacked side-by-side in a row or column
// entirely inside the row with no gaps and return those rooms.
func FitInto(bound *Rect, areas *FTree) []*Rect {
	if bound.Width() > bound.Height() {
		return RowStackInto(areas, bound)
	} else {
		return ColStackInto(areas, bound)
	}
}

///////////////////////////
// Bisectors / Splitters //
///////////////////////////

// Given a room like this:
//    +----------+
//    |          |
//    +----------+
//
// Split create two rooms by cutting top to bottom (in the example below,
// the cut is 33% into the room:
//    +---+------+
//    |   |      |
//    +---+------+
func SplitDownTheMiddle(r *Rect, percentFromLeft float64) (Rect, Rect) {
	left := *r
	right := *r
	left.SetWidthR(percentFromLeft * r.Width())
	right.SetWidthL((1 - percentFromLeft) * r.Width())
	return left, right
}

// Given a rectangle like this:
//    +-----+
//    |     |
//    |     |
//    |     |
//    |     |
//    +-----+
//
// Split create two rectangles by cutting top to bottom (in the example below,
// the cut is 33% into the rectangle:
//    +-----+
//    |     |
//    +-----+
//    |     |
//    |     |
//    +-----+
func SplitLeftToRight(r *Rect, percentFromBtm float64) (Rect, Rect) {
	btm := Rect{X0: r.X0, Y0: r.Y0, X1: r.X1, Y1: r.Y1, Label: r.Label}
	top := Rect{X0: r.X0, Y0: r.Y0, X1: r.X1, Y1: r.Y1, Label: r.Label}
	btm.SetHeightTop(percentFromBtm * r.Height())
	top.SetHeightBtm((1 - percentFromBtm) * r.Height())
	return btm, top
}
