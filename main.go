package main

import (
	"fmt"
	"github.com/fogleman/gg"
	data "housify/data_structures"
	"math/rand"
	"time"
)

const (
	scaleAmt = 10.0
)

func Draw(parent *data.RTree) {
	parent.Scale(scaleAmt)
	w := int(parent.Value.Width())
	h := int(parent.Value.Height())
	dc := gg.NewContext(w, h)
	dc.SetRGB(1, 1, 1)
	dc.Fill()
	dc.Clear()
	if err := dc.LoadFontFace("./data/ubuntu.ttf", 12); err != nil {
		panic(err)
	}
	dc.SetLineWidth(1)
	dc.SetRGBA(0, 0, 0, 1)
	for _, r := range parent.Leafs() {
		if r.Value.Label == "Living" {
			r.Scale(scaleAmt)
			pt := r.Value.Center()
			dc.DrawStringAnchored(r.Value.Label, pt.X, pt.Y, 0.5, 0.5)
			continue
		}
		r.Scale(scaleAmt)
		// The coordinate system used winds up flipping the image across the
		// vertical axis. This is because the library used considers the NW
		// corner to be the origin, but I always think of points as being on a
		// Cartesian plane, so the SW corner is the origin. This shouldn't
		// really matter, however.
		dc.DrawRectangle(
			r.Value.X0,
			r.Value.Y0,
			r.Value.Width(),
			r.Value.Height(),
		)
		dc.Stroke()
		pt := r.Value.Center()
		dc.DrawStringAnchored(r.Value.Label, pt.X, pt.Y, 0.5, 0.5)
	}
	dc.SavePNG("out.png")
	fmt.Println("done!")
}

// for debugging purposes only
func drawWithBackbone(parent *data.RTree, backbone []data.Line) {
	parent.Scale(scaleAmt)
	w := int(parent.Value.Width())
	h := int(parent.Value.Height())
	dc := gg.NewContext(w, h)
	dc.SetRGB(1, 1, 1)
	dc.Fill()
	dc.Clear()
	if err := dc.LoadFontFace("./data/ubuntu.ttf", 12); err != nil {
		panic(err)
	}
	dc.SetLineWidth(1)
	dc.SetRGBA(0, 0, 0, 1)
	for _, r := range parent.Leafs() {
		r.Scale(scaleAmt)
		// The coordinate system used winds up flipping the image across the
		// vertical axis. This is because the library used considers the NW
		// corner to be the origin, but I always think of points as being on a
		// Cartesian plane, so the SW corner is the origin. This shouldn't
		// really matter, however.
		dc.DrawRectangle(
			r.Value.X0,
			r.Value.Y0,
			r.Value.Width(),
			r.Value.Height(),
		)
		dc.Stroke()
		pt := r.Value.Center()
		dc.DrawStringAnchored(r.Value.Label, pt.X, pt.Y, 0.5, 0.5)
	}
	dc.SetRGB255(255, 0, 0)
	dc.SetLineWidth(4.0)
	for _, ln := range backbone {
		a, b := data.LineToPt(ln)
		dc.DrawLine(a.X*scaleAmt, a.Y*scaleAmt, b.X*scaleAmt, b.Y*scaleAmt)
		dc.Stroke()
	}
	dc.SavePNG("out.png")
	fmt.Println("done!")
}

func TestTree() (data.Rect, *data.FTree) {
	bounds := data.Rect{X0: 0, Y0: 0, X1: 60, Y1: 40, Label: "root"}
	a := data.FTree{Value: 60, Label: "Living", Cnx: nil}
	b := data.FTree{Value: 60, Label: "Kitchen", Cnx: nil}
	c := data.FTree{Value: 40, Label: "Master Bed", Cnx: nil}
	d := data.FTree{Value: 30, Label: "Bed", Cnx: nil}
	e := data.FTree{Value: 20, Label: "Bath", Cnx: nil}
	f := data.FTree{Value: 20, Label: "Laundry", Cnx: nil}
	g := data.FTree{Value: 10, Label: "Closet", Cnx: nil}
	areas := data.FTree{
		Value: 60 * 40,
		Label: "Root",
		Cnx:   []*data.FTree{&a, &b, &c, &d, &e, &f, &g},
	}
	return bounds, &areas
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	bounds, areas := GenHouse("data/room_edges.csv", "data/room_sizes.csv")
	squarified := Squarify(bounds, areas)
	squarified.Quantize(0) // prevents floating point errors
	backbone := Backbone(bounds, squarified)
	InsertHallway(squarified, backbone, 4)
	Draw(squarified)
}
