package data_structures

// Returns the top, right, bottom, and left lines of a rectangle.
func RectToLines(r *Rect) (Line, Line, Line, Line) {
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
	a := path[0]
	for _, b := range path[1:] {
		if !HasPt(g[a], b) {
			Add(g, a, b)
		}
		a = b
	}
}

func PathsToGraph(paths [][]Pt) Graph {
	g := make(Graph, 0)
	for _, path := range paths {
		addPath(g, path)
	}
	return g
}

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
