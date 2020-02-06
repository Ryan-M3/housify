package data_structures

import (
	"math"
	"sort"
)

type Line struct {
	x0, y0, x1, y1 float64
}

func (ln *Line) Length() float64 {
	return math.Sqrt(math.Pow(math.Abs(ln.x0-ln.x1), 2) + math.Pow(math.Abs(ln.y0-ln.y1), 2))
}

func Horz(line Line) bool {
	return line.y0 == line.y1
}

func Vert(line Line) bool {
	return line.x0 == line.x1
}

func Collinear(a, b Line) bool {
	return (Horz(a) && Horz(b) && a.y0 == b.y0) ||
		(Vert(a) && Vert(b) && a.x0 == b.x0)
}

func InPerimeter(hull Rect, line Line) bool {
	if Horz(line) {
		if (line.y0 == hull.Y0) || (line.y0 == hull.Y1) {
			return hull.X0 <= line.x0 && line.x1 <= hull.X1
		}
	} else {
		if (line.x0 == hull.X0) || (line.x0 == hull.X1) {
			return hull.Y0 <= line.y0 && line.y1 <= hull.Y1
		}
	}
	return false
}

// Do two horizontal, collinear lines intersect?
// Given two lines, a:b, and c:d, two lines
// intesect if:
//		a < c < b
//
// Visually:
//		+---+-+-----+
// 		a   c b     d
//
// They're tangent if:
//		a < b && b == c
//
// and they're non-intersecting if:
//		a < b < c
func IntersectsHorz(left Line, right Line) bool {
	if left.y0 != right.y0 {
		return false
	} else if right.x0 < left.x0 {
		return IntersectsHorz(right, left)
	} else {
		return Horz(left) && Horz(right) && right.x0 < left.x1
	}
}

// Do two vertical, collinear lines intersect?
// Given two lines, a:b, and c:d, two lines
// intesect if:
//		a < c < b
//
// Visually:
//		+---+-+-----+
// 		a   c b     d
//
// They're tangent if:
//		a < b && b == c
//
// and they're non-intersecting if:
//		a < b < c
func IntersectsVert(btm Line, top Line) bool {
	if top.x0 != btm.x0 {
		return false
	} else if top.y0 < btm.y0 {
		return IntersectsVert(top, btm)
	} else {
		return Vert(btm) && Vert(top) && top.y0 < btm.y1
	}
}

// Given two lines, a:b, and c:d, two lines
// intesect if:
//		a < c < b
//
// Visually:
//		+---+-+-----+
// 		a   c b     d
//
// They're tangent if:
//		a < b && b == c
//
// and they're non-intersecting if:
//		a < b < c
func Intersects(left, right Line) bool {
	return IntersectsHorz(left, right) || IntersectsVert(left, right)
}

// Determine if an line is inside another line.
func Inside(bounding Line, line Line) bool {
	return (Horz(bounding) && Horz(line) && bounding.x0 <= line.x0 && line.x1 <= bounding.x1) ||
		(Vert(bounding) && Vert(line) && bounding.y0 <= line.y0 && line.y1 <= bounding.y1)
}

// Like Inside, but returns false if either line shares the same left or right edge.
func StrictlyInside(bounding Line, line Line) bool {
	return (Horz(bounding) && Horz(line) && bounding.x0 < line.x0 && line.x1 < bounding.x1) ||
		(Vert(bounding) && Vert(line) && bounding.y0 < line.y0 && line.y1 < bounding.y1)
}

func ResegmentLine(left Line, right Line) []Line {
	// If segments are horizontal, then the ys are all the same, and if they're
	// vertical then the xs are all the same. This allows us to combine
	// operations of resegmenting vertical and horizontal lines into a single
	// permuation of x and y coordinates into three lines.
	if StrictlyInside(left, right) {
		// Visual diagram for reference:
		// +-------------------+
		// |    +--------+     |
		// a    b       d      c
		// LX0  RXO     RX1    LX1
		return []Line{
			Line{left.x0, left.y0, right.x0, right.y0},
			Line{left.x0, left.y0, left.x1, left.y1},
			Line{right.x1, right.y1, left.x1, left.y1},
		}
	} else if Inside(left, right) {
		if Horz(left) && Horz(right) {
			xs := []float64{left.x0, left.x1, right.x0, right.x1}
			sort.SliceStable(xs, func(i, j int) bool {
				return xs[i] < xs[j]
			})
			y := left.y0
			return []Line{
				Line{xs[0], y, xs[1], y},
				Line{xs[2], y, xs[3], y},
			}
		} else if Vert(left) && Vert(right) {
			ys := []float64{left.y0, left.y1, right.y0, right.y1}
			sort.SliceStable(ys, func(i, j int) bool {
				return ys[i] < ys[j]
			})
			x := left.x0
			return []Line{
				Line{x, ys[0], x, ys[1]},
				Line{x, ys[2], x, ys[3]},
			}
		} else {
			panic("incorrect input in ResegmentLine()")
		}
	} else {
		// Visual diagram for reference:
		// +-------+
		// |    +--|-----+
		// a    b  c     d
		// LX0  RX0 LX1  RX1
		return []Line{
			Line{left.x0, left.y0, right.x0, right.y0},
			Line{right.x0, right.y0, left.x1, left.y1},
			Line{left.x1, left.y1, right.x1, right.y1},
		}
	}
}

func resegmentLinesHorz(lines []Line) []Line {
	sort.SliceStable(lines, func(i, j int) bool {
		return lines[i].y0 < lines[j].y0
	})
	sort.SliceStable(lines, func(i, j int) bool {
		return lines[i].x0 < lines[j].x0
	})
	output := []Line{}
	for i := 1; i < len(lines); i++ {
		if Intersects(lines[i-1], lines[i]) {
			newLines := ResegmentLine(lines[i-1], lines[i])
			output = append(output, newLines...)
		} else {
			output = append(output, lines[i-1])
			if i+1 == len(lines) {
				output = append(output, lines[i])
			}
		}
	}
	return output
}

func resegmentLinesVert(lines []Line) []Line {
	sort.SliceStable(lines, func(i, j int) bool {
		return lines[i].x0 < lines[j].x0
	})
	sort.SliceStable(lines, func(i, j int) bool {
		return lines[i].y0 < lines[j].y0
	})
	output := []Line{}
	for i := 1; i < len(lines); i++ {
		if Intersects(lines[i-1], lines[i]) {
			newLines := ResegmentLine(lines[i-1], lines[i])
			output = append(output, newLines...)
		} else {
			output = append(output, lines[i-1])
			if i+1 == len(lines) {
				output = append(output, lines[i])
			}
		}
	}
	return output
}

// Warning: lines of negative length break this.
func ResegmentLines(lines []Line) []Line {
	var hs, vs []Line
	for _, ln := range lines {
		if Horz(ln) {
			hs = append(hs, ln)
		} else {
			vs = append(vs, ln)
		}
	}
	hs = resegmentLinesHorz(hs)
	vs = resegmentLinesVert(vs)
	return append(hs, vs...)
}

func RectToPts(r Rect) (Pt, Pt, Pt, Pt) {
	return Pt{r.X1, r.Y1}, // NE
		Pt{r.X1, r.Y0}, // SE
		Pt{r.X0, r.Y0}, // SW
		Pt{r.X0, r.Y1} // NW
}

func FilterLines(lines []Line) []Line {
	added := make(map[Line]bool, 0)
	var output []Line
	for _, ln := range lines {
		if _, ok := added[ln]; !ok && ln.Length() > 0 {
			output = append(output, ln)
			added[ln] = true
		}
	}
	return output
}
