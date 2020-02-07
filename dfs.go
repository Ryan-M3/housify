package main

import (
	data "housify/data_structures"
	"sort"
)

func DFS(g data.Graph, src data.Pt, tgts []data.Pt, visited map[data.Pt]bool) []data.Pt {
	if data.HasPt(tgts, src) {
		return []data.Pt{src}
	}
	visited[src] = true
	sort.SliceStable(g[src], func(i, j int) bool {
		return heuristic(g[src][i], src) < heuristic(g[src][j], src)
	})
	for _, neighbor := range g[src] {
		if visited[neighbor] {
			continue
		}
		toTgt := DFS(g, neighbor, tgts, visited)
		if len(toTgt) > 0 {
			return append([]data.Pt{src}, toTgt...)
		}
	}
	return []data.Pt{}
}

func HasPtInSublist(sublists [][]data.Pt, pt data.Pt) bool {
	for _, sublist := range sublists {
		if data.HasPt(sublist, pt) {
			return true
		}
	}
	return false
}

func HasAnyInSublist(sublists [][]data.Pt, pts []data.Pt) bool {
	for _, pt := range pts {
		if HasPtInSublist(sublists, pt) {
			return true
		}
	}
	return false
}

func DFSs(g data.Graph, tgtOpts [][]data.Pt) [][]data.Pt {
	src := data.SelectKey(g)
	output := make([][]data.Pt, len(tgtOpts))
	for i, tgts := range tgtOpts {
		if !HasAnyInSublist(output, tgts) {
			output[i] = DFS(g, src, tgts, make(map[data.Pt]bool, 0))
		}
	}
	return output
}
