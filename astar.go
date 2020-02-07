package main

import (
	"fmt"
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
func AStar(a, b data.Pt, g data.Graph, visited map[data.Pt]bool) ([]data.Pt, bool) {
	if a == b {
		return []data.Pt{b}, true
	}
	visited[a] = true
	if pts, ok := g[a]; ok {
		sort.SliceStable(pts, func(i, j int) bool {
			return heuristic(a, pts[i]) < heuristic(a, pts[j])
		})
		for _, pt := range pts {
			if _, ok := visited[pt]; ok {
				continue
			}
			if path, ok := AStar(pt, b, g, visited); ok {
				return append([]data.Pt{a}, path...), ok
			}
		}
	}
	return nil, false
}

func AStars(src data.Pt, tgts []data.Pt, g data.Graph) [][]data.Pt {
	paths := make([][]data.Pt, len(tgts))
	for i, tgt := range tgts {
		if path, ok := AStar(src, tgt, g, make(map[data.Pt]bool, 0)); ok {
			paths[i] = path
		} else {
			fmt.Printf("Error, no path between Src: %v and Tgt: %v\n", src, tgt)
			//return AStars(selectKey(g), tgts, g)
			//panic("invalid graph generated / no path found between points")
		}
	}
	return paths
}
