package dataStructures

type RTree struct {
	Value Rect
	Cnx   []*RTree
}

func (t *RTree) List() []*RTree {
	listed := []*RTree{t}
	for _, node := range t.Cnx {
		listed = append(listed, node.List()...)
	}
	return listed
}

func (t *RTree) Leafs() []*RTree {
	if len(t.Cnx) == 0 {
		return []*RTree{t}
	}
	listed := []*RTree{}
	for _, node := range t.Cnx {
		listed = append(listed, node.Leafs()...)
	}
	return listed
}

func (t *RTree) Find(label string) (*RTree, bool) {
	if t.Value.Label == label {
		return t, true
	}
	for _, node := range t.Cnx {
		if found, ok := node.Find(label); ok {
			return found, true
		}
	}
	return nil, false
}

func (t *RTree) CnxValues() []Rect {
	var output []Rect
	for _, node := range t.Cnx {
		output = append(output, node.Value)
	}
	return output
}

func (t *RTree) FindNearestPt(pt Pt) []*Rect {
	var output []*Rect
	for _, leaf := range t.Leafs() {
		if leaf.Value.Contains(pt) {
			output = append(output, &leaf.Value)
		}
	}
	return output
}

// Increase or decrease all corners. Doe NOT scale in place, it scales from the
// origin. Use this if you are iterating over rooms and want to make the entire
// picture bigger.
func (t *RTree) Scale(amt float64) {
	t.Value.Scale(amt)
}

func (t *RTree) Quantize(decimalPlaces float64) {
	for _, branch := range t.List() {
		branch.Value.X0 = RoundTo(branch.Value.X0, decimalPlaces)
		branch.Value.X1 = RoundTo(branch.Value.X1, decimalPlaces)
		branch.Value.Y0 = RoundTo(branch.Value.Y0, decimalPlaces)
		branch.Value.Y1 = RoundTo(branch.Value.Y1, decimalPlaces)
	}
}

func (t *RTree) Adjacent(room *Rect, house *RTree) []*Rect {
	var adj []*Rect
	a, b, c, d := RectToLines(room)
	lines := []Line{a, b, c, d}
	for _, candidate := range RTreesToRects(house.Leafs()) {
		if candidate != *room {
			inCommon := 0
			x, y, z, w := RectToLines(&candidate)
			for _, candidatePt := range []Line{x, y, z, w} {
				for _, ln := range lines {
					if Intersects(ln, candidatePt) {
						inCommon++
					}
				}
			}
			if inCommon > 1 {
				adj = append(adj, &candidate)
			}
		}
	}
	return adj
}
