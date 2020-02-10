package dataStructures

import (
	"fmt"
)

type Graph = map[Pt][]Pt

func Add(g Graph, k, v Pt) {
	if _, ok := g[k]; ok {
		g[k] = append(g[k], v)
	} else {
		g[k] = []Pt{v}
	}
}

func SelectKey(g Graph) Pt {
	for k, _ := range g {
		return k
	}
	panic("Can't select key in empty graph.")
}

func PrintGraph(g Graph) {
	fmt.Println("Graph:")
	for k, vs := range g {
		fmt.Printf("  Key: (%.2f, %.2f)\n", k.X, k.Y)
		for _, v := range vs {
			fmt.Printf("    Val: (%.2f, %.2f)\n", v.X, v.Y)
		}
		fmt.Println()
	}
}
