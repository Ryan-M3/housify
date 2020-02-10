package main

import (
	data "housify/data_structures"
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
			//!data.InPerimeter(living.Value, ln) &&
			//!data.EndsInPerimeter(bounds, ln) {
			lines = append(lines, ln)
		}
	}
	// Find the side of the living room closest to the middle of the house by sorting it.
	top, right, btm, left := data.RectToLines(&living.Value)
	sides := []data.Line{top, right, btm, left}
	sort.SliceStable(sides, func(i, j int) bool {
		a, b := data.LineToPt(sides[i])
		c, d := data.LineToPt(sides[j])
		mid1 := data.Pt{(a.X + b.X) / 2, (a.Y + b.Y) / 2}
		mid2 := data.Pt{(c.X + d.X) / 2, (c.Y + d.Y) / 2}
		return data.Distance(mid1, bounds.Center()) < data.Distance(mid2, bounds.Center())
	})
	paths := AStars(
		sides[0],
		data.LinesToGraph(lines),
		generateTargets(rects),
	)
	return data.PathsToGraph(paths)
	//return data.LinesToGraph(lines)
}
