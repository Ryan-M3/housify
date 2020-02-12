package main

import (
	data "housify/dataStructures"
	"sort"
)

// Generate a nested list of points which connect to the primary leaf nodes.
func generateTargets(rs []data.Rect) [][]data.Line {
	var lines [][]data.Line
	for _, r := range rs {
		a, b, c, d := data.RectToLines(&r)
		tgt := []data.Line{a, b, c, d}
		lines = append(lines, tgt)
	}
	return lines
}

// The interior edges that a hallway could be inserted into.
func Backbone(bounds data.Rect, house *data.RTree) data.Graph {
	leafs := house.Leafs()
	rects := data.RTreesToRects(leafs)
	living, _ := house.Find("Living")
	var lines []data.Line
	for _, ln := range data.ResegmentLines(data.RectsToLines(rects)) {
		if !data.InPerimeter(bounds, ln) {
			lines = append(lines, ln)
		}
	}
	// Find the side of the living room closest to the middle of the house by
	// sorting it.
	top, right, btm, left := data.RectToLines(&living.Value)
	sides := []data.Line{top, right, btm, left}
	sort.SliceStable(sides, func(i, j int) bool {
		return data.Distance(sides[i].MidPt(), bounds.Center()) < data.Distance(sides[j].MidPt(), bounds.Center())
	})
	paths := AStars(
		sides[0],
		data.LinesToGraph(lines),
		generateTargets(rects),
	)
	return data.PathsToGraph(paths)
}
