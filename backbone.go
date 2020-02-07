package main

import (
	data "housify/data_structures"
)

// Generate a nested list of points which connect to the primary leaf nodes.
func generateTargets(rs []data.Rect) [][]data.Pt {
	var pts [][]data.Pt
	for _, r := range rs {
		a, b, c, d := data.RectToPts(r)
		tgt := []data.Pt{a, b, c, d}
		pts = append(pts, tgt)
	}
	return pts
}

func Backbone(bounds data.Rect, house *data.RTree) data.Graph {
	leafs := house.Leafs()
	rects := data.RTreesToRects(leafs)
	var lines []data.Line
	for _, ln := range data.ResegmentLines(data.RectsToLines(rects)) {
		if !data.InPerimeter(bounds, ln) &&
			!data.EndsInPerimeter(bounds, ln) {
			lines = append(lines, ln)
		}
	}
	paths := AStars(
		data.LinesToGraph(lines),
		generateTargets(rects),
	)
	return data.PathsToGraph(paths)
}
