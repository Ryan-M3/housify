package main

import (
	data "housify/dataStructures"
	"math"
	"sort"
)

// Since lines are all at a 90 degree angle from eachother, Manhattan Distance
// is used as the heuristic function.
func heuristic(a, b data.Pt) float64 {
	return math.Abs(a.X-b.X) + math.Abs(a.Y-b.Y)
}

// we need to check that a path goes across a target room.
//func inBoth(xs, ys []*data.Rect) []*data.Rect {

// AStar search is really just DFS where you prioritize closer nodes over
// farther nodes.
func AStar(g data.Graph, src data.Line, tgts []data.Line, visited map[data.Line]bool) []data.Line {
	if data.HasLine(tgts, src) {
		return []data.Line{src}
	}
	visited[src] = true
	_, endpt := data.LineToPt(src)
	sort.SliceStable(g[endpt], func(i, j int) bool {
		return heuristic(g[endpt][i], endpt) < heuristic(g[endpt][j], endpt)
	})
	for _, neighbor := range g[endpt] {
		nextLine := data.PtsToLine(endpt, neighbor)
		if visited[nextLine] {
			continue
		}
		toTgt := AStar(g, nextLine, tgts, visited)
		if len(toTgt) > 0 {
			return append([]data.Line{src}, toTgt...)
		}
	}
	return []data.Line{}
}

func HasLineInSublists(sublists [][]data.Line, line data.Line) bool {
	for _, sublist := range sublists {
		if data.HasLine(sublist, line) {
			return true
		}
	}
	return false
}

func HasAnyLineInAnySublist(sublists [][]data.Line, lines []data.Line) bool {
	for _, ln := range lines {
		if HasLineInSublists(sublists, ln) {
			return true
		}
	}
	return false
}

func AStars(src data.Line, g data.Graph, tgtOpts [][]data.Line) [][]data.Line {
	if len(g) == 0 {
		return nil
	}
	output := make([][]data.Line, len(tgtOpts))
	for i, tgts := range tgtOpts {
		if !HasAnyLineInAnySublist(output, tgts) {
			output[i] = AStar(g, src, tgts, make(map[data.Line]bool, 0))
		}
	}
	return output
}
