package main

import (
	"fmt"
	data "housify/data_structures"
)

func connectToCenter(g data.Graph, rects []data.Rect) {
	for _, r := range rects {
		ne, se, sw, nw := data.RectToPts(r)
		center := r.Center()
		for _, corner := range []data.Pt{ne, se, sw, nw} {
			data.Add(g, corner, center)
		}
	}
}

func GenerateTargets(rs []data.Rect) [][]data.Pt {
	var pts [][]data.Pt
	for _, r := range rs {
		a, b, c, d := data.RectToPts(r)
		tgt := []data.Pt{a, b, c, d}
		pts = append(pts, tgt)
	}
	return pts
}

func HouseGraph(bounds data.Rect, house *data.RTree) data.Graph {
	house.Quantize(0)
	leafs := house.Leafs()
	rects := data.RTreesToRects(leafs)
	var lines []data.Line
	for _, ln := range data.ResegmentLines(data.RectsToLines(rects)) {
		if !data.InPerimeter(bounds, ln) &&
			!data.EndsInPerimeter(bounds, ln) {
			lines = append(lines, ln)
		}
	}
	g := data.LinesToGraph(lines)
	tgts := GenerateTargets(rects)
	paths := DFSs(g, tgts)
	fmt.Println()
	return data.PathsToGraph(paths)
	//data.PrintGraph(g)
	//k := data.SelectKey(g)
	//astars := AStars(k, centers, g)
	//living, _ := house.Find("Living")
	//astars := DFSs(g, centers)
	//return data.PathsToGraph(astars)
	//return g
}

//func Backbone(bounds data.Rect, house *data.RTree) data.Graph {
//	_, ok := house.Find("Living")
//	if !ok {
//		panic("House generated without a living room.")
//	}
//	leafs := house.Leafs()
//	rects := data.RTreesToRects(leafs)
//	var lines []data.Line
//	for _, ln := range data.ResegmentLines(data.RectsToLines(rects)) {
//		if !data.InPerimeter(bounds, ln) && !data.EndsInPerimeter(bounds, ln) { // && !data.InPerimeter(living.Value, ln, 8) {
//			lines = append(lines, ln)
//		}
//	}
//	g := data.LinesToGraph(lines)
//
//	var rects2 []data.Rect // rects minus the living rooms rect
//	for _, r := range rects {
//		if r.Label != "Living" {
//			rects2 = append(rects2, r)
//		}
//	}
//	rects = rects2
//
//	connectToCenter(g, rects)
//	centers := data.RectsToCenters(rects)
//	//astars := AStars(k, centers, g)
//	astars := DFSs(g, centers)
//	return data.PathsToGraph(astars)
//}
