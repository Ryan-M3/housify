package data_structures

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

func (t *RTree) FindNearestPt(pt Pt) []Rect {
	var output []Rect
	for _, leaf := range t.Leafs() {
		if leaf.Value.Contains(pt) {
			output = append(output, leaf.Value)
		}
	}
	return output
}
