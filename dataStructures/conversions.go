package dataStructures

import (
	"math"
)

// Returns the top, right, bottom, and left lines of a rectangle.
func RectToLines(r *Rect) (Line, Line, Line, Line) {
	// x0,y1   x1,y1
	//   + - - - +
	//   |       |
	//   |       |
	//   + - - - +
	// x0,y0    x1,y0
	return Line{r.X0, r.Y1, r.X1, r.Y1}, // top
		Line{r.X1, r.Y0, r.X1, r.Y1}, // right
		Line{r.X0, r.Y0, r.X1, r.Y0}, // bottom
		Line{r.X0, r.Y0, r.X0, r.Y1} // left
}

func RectsToLines(rects []Rect) []Line {
	lines := make([]Line, len(rects)*4)
	for i, r := range rects {
		top, right, btm, left := RectToLines(&r)
		lines[i*4+0] = top
		lines[i*4+1] = right
		lines[i*4+2] = btm
		lines[i*4+3] = left
	}
	return lines
}

func LinesToGraph(lines []Line) Graph {
	g := make(Graph, 0)
	for _, ln := range lines {
		left, right := LineToPt(ln)
		// no self-connections
		if left == right {
			continue
		}
		if !HasPt(g[left], right) {
			Add(g, left, right)
		}
		if !HasPt(g[right], left) {
			Add(g, right, left)
		}
	}
	return g
}

func addPath(g Graph, path []Pt) {
	if len(path) < 2 {
		return
	}
	a := path[0]
	for _, b := range path[1:] {
		if !HasPt(g[a], b) {
			Add(g, a, b)
		}
		a = b
	}
}

//func PathsToGraph(paths [][]Pt) Graph {
//	g := make(Graph, 0)
//	for _, path := range paths {
//		addPath(g, path)
//	}
//	return g
//}

func RTreesToRects(rtrees []*RTree) []Rect {
	rs := make([]Rect, len(rtrees))
	for i, rtree := range rtrees {
		rs[i] = rtree.Value
	}
	return rs
}

func LineToPt(ln Line) (Pt, Pt) {
	return Pt{ln.x0, ln.y0}, Pt{ln.x1, ln.y1}
}

func RectsToCenters(rects []Rect) []Pt {
	centers := make([]Pt, len(rects))
	for i, r := range rects {
		centers[i] = r.Center()
	}
	return centers
}

func GraphToLines(g Graph) []Line {
	var edges []Line
	added := make(map[Line]bool, 0)
	for k, vs := range g {
		for _, v := range vs {
			minx := math.Min(k.X, v.X)
			maxx := math.Max(k.X, v.X)
			miny := math.Min(k.Y, v.Y)
			maxy := math.Max(k.Y, v.Y)
			e := Line{minx, miny, maxx, maxy}
			// a line from point A to point B is consider the same line as the
			// line from point B to point A; below we check for duplicates
			// before adding it to the output
			if _, ok := added[e]; !ok {
				edges = append(edges, e)
				added[e] = true
				reversed := Line{e.x1, e.y1, e.x0, e.y0}
				added[reversed] = true
			}
		}
	}
	return edges
}

func RectToPts(r Rect) (Pt, Pt, Pt, Pt) {
	return Pt{r.X1, r.Y1}, // NE
		Pt{r.X1, r.Y0}, // SE
		Pt{r.X0, r.Y0}, // SW
		Pt{r.X0, r.Y1} // NW
}

func PtsToLine(a, b Pt) Line {
	return Line{a.X, a.Y, b.X, b.Y}
}

func PathsToGraph(paths [][]Line) Graph {
	g := make(Graph, 0)
	for _, path := range paths {
		for _, line := range path {
			ptA, ptB := LineToPt(line)
			Add(g, ptA, ptB)
		}
	}
	return g
}

func RectsToRooms(rects []Rect) []*Room {
	output := make([]*Room, len(rects))
	for i, _ := range rects {
		output[i] = &Room{&rects[i], nil, nil}
	}
	return output
}
