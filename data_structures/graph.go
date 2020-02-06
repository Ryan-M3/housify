package data_structures

import ()

type Graph = map[Pt][]Pt

func Add(g Graph, k, v Pt) {
	if _, ok := g[k]; ok {
		g[k] = append(g[k], v)
	} else {
		g[k] = []Pt{v}
	}
}
