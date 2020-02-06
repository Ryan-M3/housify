package data_structures

import (
	"sort"
)

type FTree struct {
	Value float64
	Label string
	Cnx   []*FTree
}

func (t *FTree) List() []*FTree {
	listed := []*FTree{t}
	for _, node := range t.Cnx {
		listed = append(listed, node.List()...)
	}
	return listed
}

func (t *FTree) Sort() {
	sort.SliceStable(t.Cnx, func(i, j int) bool {
		return t.Cnx[i].Value > t.Cnx[j].Value
	})
}

func (t *FTree) Leafs() []*FTree {
	if len(t.Cnx) == 0 {
		return []*FTree{t}
	}
	listed := []*FTree{}
	for _, node := range t.Cnx {
		listed = append(listed, node.Leafs()...)
	}
	return listed
}

func (t *FTree) Find(value interface{}) (*FTree, bool) {
	if t.Value == value {
		return t, true
	}
	for _, node := range t.Cnx {
		if found, ok := node.Find(value); ok {
			return found, true
		}
	}
	return nil, false
}

func (t *FTree) CnxValues() []float64 {
	t.Sort()
	values := make([]float64, len(t.Cnx))
	for i, node := range t.Cnx {
		values[i] = node.Value
	}
	return values
}

func (t *FTree) Add(value float64, label string) {
	node := FTree{Value: value, Label: label}
	t.Cnx = append(t.Cnx, &node)
}

// Resize a list of areas so that their squares would fit into a width. This is
// slightly different from just scaling the areas based no the proporition of
// area they take up because area is two dimensional and width is one
// dimensional.
func (areas *FTree) scaleAreas(width float64) {
	for i, _ := range areas.Cnx {
		areas.Cnx[i].Value /= width
	}
	total := sum(areas.CnxValues())
	for i, _ := range areas.Cnx {
		areas.Cnx[i].Value *= width / total
	}
}
