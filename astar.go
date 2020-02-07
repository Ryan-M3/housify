package main

import (
	data "housify/data_structures"
	"math"
	"sort"
)

// Since lines are all at a 90 degree angle from eachother, Manhattan Distance
// is used as the heuristic function.
func heuristic(a, b data.Pt) float64 {
	return math.Abs(a.X-b.X) + math.Abs(a.Y-b.Y)
}

// AStar search is really just DFS where you prioritize closer nodes over
// farther nodes.
func AStar(g data.Graph, src data.Pt, tgts []data.Pt, visited map[data.Pt]bool) []data.Pt {
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
		toTgt := AStar(g, neighbor, tgts, visited)
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

func HasAnyPtInAnySublist(sublists [][]data.Pt, pts []data.Pt) bool {
	for _, pt := range pts {
		if HasPtInSublist(sublists, pt) {
			return true
		}
	}
	return false
}

func AStars(g data.Graph, tgtOpts [][]data.Pt) [][]data.Pt {
	src := data.SelectKey(g)
	output := make([][]data.Pt, len(tgtOpts))
	for i, tgts := range tgtOpts {
		if !HasAnyPtInAnySublist(output, tgts) {
			output[i] = AStar(g, src, tgts, make(map[data.Pt]bool, 0))
		}
	}
	return output
}
