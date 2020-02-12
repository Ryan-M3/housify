package main

import (
	"github.com/fogleman/gg"
	data "housify/dataStructures"
)

func Draw(bounds data.Rect, rooms []*data.Room) {
	bounds.Scale(scaleAmt)
	w := int(bounds.Width())
	h := int(bounds.Height())
	dc := gg.NewContext(w, h)
	dc.SetRGB(1, 1, 1)
	dc.Fill()
	dc.Clear()
	if err := dc.LoadFontFace("./data/ubuntu.ttf", 12); err != nil {
		panic(err)
	}
	dc.SetLineWidth(1)
	dc.SetRGBA(0, 0, 0, 1)
	for _, r := range rooms {
		r.Rect.Scale(scaleAmt)
		if r.Rect.Label == "Living" {
			pt := r.Rect.Center()
			dc.DrawStringAnchored(r.Rect.Label, pt.X, pt.Y, 0.5, 0.5)
			continue
		}
		// The coordinate system used winds up flipping the image across the
		// vertical axis. This is because the library used considers the NW
		// corner to be the origin, but I always think of points as being on a
		// Cartesian plane, so the SW corner is the origin. This shouldn't
		// really matter, however.
		dc.DrawRectangle(
			r.Rect.X0,
			r.Rect.Y0,
			r.Rect.Width(),
			r.Rect.Height(),
		)
		dc.Stroke()
		pt := r.Rect.Center()
		dc.DrawStringAnchored(r.Rect.Label, pt.X, pt.Y, 0.5, 0.5)
		DrawDoors(dc, r)
	}
	dc.SavePNG("out.png")
}

func DrawDoors(dc *gg.Context, room *data.Room) {
	var x0, y0, x1, y1, x2, y2, x3, y3 float64
	for _, door := range room.Doors {
		top, right, btm, left := data.RectToLines(room.Rect)
		dir := data.Pt{X: 40, Y: 40}
		var where data.Pt
		switch door.Orientation {
		case data.N:
			a, b := data.LineToPt(top)
			where = data.Lerp(door.Position, a, b)
			x0, y0 = where.X-dir.X/2, where.Y
			x1, y1 = x0, y0+dir.Y
			x2, y2 = x0+dir.X, y1
			x3, y3 = x0+dir.X, where.Y
		case data.E:
			a, b := data.LineToPt(right)
			where = data.Lerp(door.Position, a, b)
			x0, y0 = where.X, where.Y
			x3, y3 = x0, y0+dir.Y
			x2, y2 = x0+dir.X, y3
			x1, y1 = x0+dir.X, where.Y
		case data.S:
			a, b := data.LineToPt(btm)
			where = data.Lerp(door.Position, a, b)
			x0, y0 = where.X-dir.X/2, where.Y
			x1, y1 = x0, y0-dir.Y
			x2, y2 = x0+dir.X, y1
			x3, y3 = x0+dir.X, where.Y
		case data.W:
			a, b := data.LineToPt(left)
			where = data.Lerp(door.Position, a, b)
			dir.X *= -1
			x1, y1 = where.X+dir.X, where.Y
			x2, y2 = x1, y1+dir.Y
			x3, y3 = x1-dir.X, y2
			x0, y0 = x1-dir.X, where.Y
		}
		dc.MoveTo(x0, y0)
		dc.QuadraticTo(x1, y1, x2, y2)
		dc.Stroke()
		dc.DrawLine(x2, y2, x3, y3)
		dc.Stroke()
		dc.SetRGB255(255, 255, 255)
		dc.DrawLine(x0, y0, x3, y3)
		// You wouldn't think that calling this function multiple times would
		// make the previous stoke with 100% alpha any more white, but it does.
		dc.Stroke()
		dc.Stroke()
		dc.SetRGB255(0, 0, 0)
	}
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
}
