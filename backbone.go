package main

import (
	data "housify/data_structures"
)

func selectKey(g data.Graph) data.Pt {
	for k, _ := range g {
		return k
	}
	panic("Can't select key in empty graph.")
}

func connectToCenter(g data.Graph, rects []data.Rect) {
	for _, r := range rects {
		ne, se, sw, nw := data.RectToPts(r)
		center := r.Center()
		for _, corner := range []data.Pt{ne, se, sw, nw} {
			if _, ok := g[corner]; ok {
				data.Add(g, corner, center)
			}
		}
	}
}

func Backbone(bounds data.Rect, house *data.RTree) data.Graph {
	living, ok := house.Find("Living")
	if !ok {
		panic("House generated without a living room.")
	}
	leafs := house.Leafs()
	rects := data.RTreesToRects(leafs)
	var lines []data.Line
	for _, ln := range data.ResegmentLines(data.RectsToLines(rects)) {
		if !data.InPerimeter(bounds, ln) && !data.InPerimeter(living.Value, ln) {
			lines = append(lines, ln)
		}
	}
	g := data.LinesToGraph(lines)
	connectToCenter(g, rects)
	k := selectKey(g)

	centers := data.RectsToCenters(rects)
	astars := AStars(k, centers, g)
	return data.PathsToGraph(astars)
}
