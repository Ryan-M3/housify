package dataStructures

import (
	"fmt"
	"math"
	"testing"
	"testing/quick"
)

func CompareGraphs(g, h Graph) {
	for k, vs := range g {
		if _, ok := h[k]; ok {
			fmt.Printf("\ng and h: %v\n", k)
		} else {
			fmt.Printf("\ng only: %v\n", k)
		}
		for i, _ := range vs {
			if i < len(h[k]) {
				fmt.Printf("\tg: %v    h: %v\n", vs[i], h[k][i])
			} else {
				fmt.Printf("\tg: %v\n", vs[i])
			}
		}
		if len(h[k]) > len(vs) {
			for i := len(vs); i < len(h[k]); i++ {
				fmt.Printf("\t\t    h: %v\n", h[k][i])
			}
		}
	}
}

func Test_GraphToLines(t *testing.T) {
	fn := func(nums []float64) bool {
		for i := 0; i < len(nums); i++ {
			nums[i] = math.Mod(math.Round(nums[i]), 9)
		}
		g := make(Graph, 0)
		// Lines of zero length get added to the graph as self-connections,
		// which are then filtered out, which means they will disappear when
		// we convert back and forth from lines to points.
		var zeroLenPts []Pt
		for i := 0; i < len(nums)-3; i += 4 {
			a := Pt{nums[i+0], nums[i+1]}
			b := Pt{nums[i+2], nums[i+3]}
			if a == b {
				zeroLenPts = append(zeroLenPts, a)
			}
			Add(g, a, b)
		}
		h := LinesToGraph(GraphToLines(g))
		for k, vs := range g {
			for _, v := range vs {
				if !HasPt(h[k], v) && !HasPt(zeroLenPts, v) {
					lns := GraphToLines(g)
					fmt.Println(lns, "\n-----------------------------")
					CompareGraphs(g, h)
					t.Errorf("%v lacks value, %v", h[k], v)
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fn, nil); err != nil {
		t.Error(err)
	}
}
