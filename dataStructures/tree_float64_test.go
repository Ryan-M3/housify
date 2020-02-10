package dataStructures

import (
	"math/rand"
	"testing"
)

func Test_FTree_Add(t *testing.T) {
	for i := 0; i < 16; i++ {
		tree := FTree{0, "root", nil}
		for j := 0; j < 99; j++ {
			flt := rand.Float64()
			tree.Add(flt, "random")
			lastFlt := tree.Cnx[len(tree.Cnx)-1].Value
			if lastFlt != flt {
				t.Errorf("expected %f, but got %f in tree.Cnx[len(tree.Cnx)]", flt, lastFlt)
			}
		}
	}
}

func Test_FTree_CnxValues(t *testing.T) {
	root := FTree{0, "root", nil}
	for i := 0; i < 10; i++ {
		root.Add(rand.Float64(), "random")
		for j := 0; j < 10; j++ {
			root.Cnx[i].Add(rand.Float64(), "random")
		}
	}
	prev := 123456789.0
	for _, v := range root.CnxValues() {
		this := v
		// Remember, we're sorting in DESCENDING value.
		if prev < this {
			t.Errorf("FTree.Sort() failed to properly sort tree: %.2f < %.2f", prev, this)
		}
		prev = v
	}
}
