package dataStructures

import (
	"math"
	"testing"
	"testing/quick"
)

const (
	bigNum = 10000000000
)

func CheckRect(t *testing.T, r *Rect, x, y, z, w float64) {
	t.Helper()
	if !AboutEqual(r.X0, x) {
		t.Errorf("Rect.X0 failure; expected %.1f, but got %.1f.", x, r.X0)
	}
	if !AboutEqual(r.Y0, y) {
		t.Errorf("Rect.Y0 failure; expected %.1f, but got %.1f.", y, r.Y0)
	}
	if !AboutEqual(r.X1, z) {
		t.Errorf("Rect.Y1 failure; expected %.1f, but got %.1f.", z, r.X1)
	}
	if !AboutEqual(r.Y1, w) {
		t.Errorf("Rect.Y1 failure; expected %.1f, but got %.1f.", w, r.Y1)
	}
}

func AboutEqual(x float64, y float64) bool {
	precision := math.Pow(2, 32)
	return math.Round(precision*x) == math.Round(precision*y)
}

func Test_Width(t *testing.T) {
	fn := func(x, y, w, h float64) bool {
		x = math.Mod(math.Abs(math.Round(x)), bigNum)
		y = math.Mod(math.Abs(math.Round(y)), bigNum)
		w = math.Mod(math.Abs(math.Round(w)), bigNum)
		h = math.Mod(math.Abs(math.Round(h)), bigNum)
		r := Rect{x, y, x + w, y + h, ""}
		return AboutEqual(r.Width(), w)
	}

	if err := quick.Check(fn, nil); err != nil {
		t.Error(err)
	}
}

func Test_Height(t *testing.T) {
	fn := func(x, y, w, h float64) bool {
		x = math.Mod(math.Abs(math.Round(x)), bigNum)
		y = math.Mod(math.Abs(math.Round(y)), bigNum)
		w = math.Mod(math.Abs(math.Round(w)), bigNum)
		h = math.Mod(math.Abs(math.Round(h)), bigNum)
		r := Rect{x, y, x + w, y + h, ""}
		return AboutEqual(r.Height(), h)
	}

	if err := quick.Check(fn, nil); err != nil {
		t.Error(err)
	}
}

func Test_Area(t *testing.T) {
	fn := func(x, y, w, h float64) bool {
		x = math.Mod(math.Abs(math.Round(x)), bigNum)
		y = math.Mod(math.Abs(math.Round(y)), bigNum)
		w = math.Mod(math.Abs(math.Round(w)), bigNum)
		h = math.Mod(math.Abs(math.Round(h)), bigNum)
		r := Rect{x, y, x + w, y + h, ""}
		return AboutEqual(r.Area(), w*h)
	}

	if err := quick.Check(fn, nil); err != nil {
		t.Error(err)
	}
}

func Test_MoveTo(t *testing.T) {
	fn := func(x, y, w, h float64) bool {
		x = math.Mod(math.Abs(math.Round(x)), bigNum)
		y = math.Mod(math.Abs(math.Round(y)), bigNum)
		w = math.Mod(math.Abs(math.Round(w)), bigNum)
		h = math.Mod(math.Abs(math.Round(h)), bigNum)
		r := Rect{x, y, x + w, y + h, ""}
		r.MoveTo(x, y)
		return AboutEqual(x, r.X0) &&
			AboutEqual(y, r.Y0) &&
			AboutEqual(r.Width(), w) &&
			AboutEqual(r.Height(), h)
	}
	if err := quick.Check(fn, nil); err != nil {
		t.Error(err)
	}
}

func Test_SetWidthR(t *testing.T) {
	fn := func(x, y, w, h, newWidth float64) bool {
		x = math.Mod(math.Abs(math.Round(x)), bigNum)
		y = math.Mod(math.Abs(math.Round(y)), bigNum)
		w = math.Mod(math.Abs(math.Round(w)), bigNum)
		h = math.Mod(math.Abs(math.Round(h)), bigNum)
		newWidth = math.Mod(math.Abs(math.Round(newWidth)), bigNum)
		r := Rect{x, y, x + w, y + h, ""}
		r.SetWidthR(newWidth)
		return AboutEqual(r.Width(), newWidth) && AboutEqual(r.X0, x) && AboutEqual(r.Y0, y)
	}
	if err := quick.Check(fn, nil); err != nil {
		t.Error(err)
	}
}

func Test_SetWidthL(t *testing.T) {
	fn := func(x, y, w, h, newWidth float64) bool {
		x = math.Mod(math.Abs(math.Round(x)), bigNum)
		y = math.Mod(math.Abs(math.Round(y)), bigNum)
		w = math.Mod(math.Abs(math.Round(w)), bigNum)
		h = math.Mod(math.Abs(math.Round(h)), bigNum)
		newWidth = math.Mod(math.Abs(math.Round(newWidth)), bigNum)
		r := Rect{x, y, x + w, y + h, ""}
		r.SetWidthL(newWidth)
		return AboutEqual(r.Width(), newWidth) &&
			AboutEqual(r.X1, x+w) &&
			AboutEqual(r.Y1, y+h)
	}
	if err := quick.Check(fn, nil); err != nil {
		t.Error(err)
	}
}

func Test_SetHeightTop(t *testing.T) {
	fn := func(x, y, w, h, newHeight float64) bool {
		x = math.Mod(math.Abs(math.Round(x)), bigNum)
		y = math.Mod(math.Abs(math.Round(y)), bigNum)
		w = math.Mod(math.Abs(math.Round(w)), bigNum)
		h = math.Mod(math.Abs(math.Round(h)), bigNum)
		newHeight = math.Mod(math.Abs(math.Round(newHeight)), bigNum)
		r := Rect{x, y, x + w, y + h, ""}
		r.SetHeightTop(newHeight)
		return AboutEqual(r.Height(), newHeight) &&
			AboutEqual(r.X0, x) &&
			AboutEqual(r.Y0, y)
	}
	if err := quick.Check(fn, nil); err != nil {
		t.Error(err)
	}
}

func Test_SetHeightBtm(t *testing.T) {
	fn := func(x, y, w, h, newHeight float64) bool {
		x = math.Mod(math.Abs(math.Round(x)), bigNum)
		y = math.Mod(math.Abs(math.Round(y)), bigNum)
		w = math.Mod(math.Abs(math.Round(w)), bigNum)
		h = math.Mod(math.Abs(math.Round(h)), bigNum)
		newHeight = math.Mod(math.Abs(math.Round(newHeight)), bigNum)
		r := Rect{x, y, x + w, y + h, ""}
		r.SetHeightBtm(newHeight)
		return AboutEqual(r.Height(), newHeight) &&
			AboutEqual(r.X1, x+w) &&
			AboutEqual(r.Y1, y+h)
	}
	if err := quick.Check(fn, nil); err != nil {
		t.Error(err)
	}
}

func Test_colStack(t *testing.T) {
	rs := []*Rect{
		&Rect{0, 0, 1, 2, ""},
		&Rect{0, 0, 2, 2, ""},
		&Rect{0, 0, 3, 2, ""},
		&Rect{0, 0, 1, 3, ""},
	}
	colStack(rs)
	CheckRect(t, rs[0], 0, 0, 1, 2)
	CheckRect(t, rs[1], 0, 2, 2, 4)
	CheckRect(t, rs[2], 0, 4, 3, 6)
	CheckRect(t, rs[3], 0, 6, 1, 9)
}

func Test_rowStack(t *testing.T) {
	rs := []*Rect{
		&Rect{-1, -1, 0, 1, ""},
		&Rect{0, 0, 2, 2, ""},
		&Rect{0, 0, 3, 2, ""},
		&Rect{0, 0, 1, 3, ""},
	}
	rowStack(rs)
	CheckRect(t, rs[0], -1, -1, 0, 1)
	CheckRect(t, rs[1], 0, -1, 2, 1)
	CheckRect(t, rs[2], 2, -1, 5, 1)
	CheckRect(t, rs[3], 5, -1, 6, 2)
}

func Test_rowStackInto(t *testing.T) {
	shelf := Rect{0, 0, 10, 2, ""}
	areas := []float64{4, 4, 4, 2, 2, 2, 2}
	ftree := FTree{0, "root", nil}
	for _, area := range areas {
		ftree.Add(area, "")
	}
	rs := RowStackInto(&ftree, &shelf)
	CheckRect(t, rs[0], 0, 0, 2, 2)
	CheckRect(t, rs[1], 2, 0, 4, 2)
	CheckRect(t, rs[2], 4, 0, 6, 2)
	CheckRect(t, rs[3], 6, 0, 7, 2)
	CheckRect(t, rs[4], 7, 0, 8, 2)
	CheckRect(t, rs[5], 8, 0, 9, 2)
	CheckRect(t, rs[6], 9, 0, 10, 2)
}

func Test_colStackInto(t *testing.T) {
	shelf := Rect{0, 0, 10, 2, ""}
	areas := []float64{4, 4, 4, 2, 2, 2, 2}
	ftree := FTree{0, "root", nil}
	for _, area := range areas {
		ftree.Add(area, "")
	}
	rs := ColStackInto(&ftree, &shelf)
	CheckRect(t, rs[0], 0.0, 0.0, 10.0, 4.0/10.0)
	CheckRect(t, rs[1], 0.0, 4.0/10.0, 10.0, 8.0/10.0)
	CheckRect(t, rs[2], 0.0, 8.0/10.0, 10.0, 12.0/10.0)
	CheckRect(t, rs[3], 0.0, 12.0/10.0, 10.0, 14.0/10.0)
	CheckRect(t, rs[4], 0.0, 14.0/10.0, 10.0, 16.0/10.0)
	CheckRect(t, rs[5], 0.0, 16.0/10.0, 10.0, 18.0/10.0)
	CheckRect(t, rs[6], 0.0, 18.0/10.0, 10.0, 20.0/10.0)
}

func Test(t *testing.T) {
	r := Rect{0, 0, 1, 1, ""}
	if !r.Contains(Pt{0.5, 0.5}) {
		t.Errorf("expected true in Contains(Pt{0.5, 0.5}), got false")
	}
	if r.Contains(Pt{2.5, 0.5}) {
		t.Errorf("expected false in Contains(Pt{2.5, 0.5}), got true")
	}
	if !r.Contains(Pt{0, 0.5}) {
		t.Errorf("expected true in Contains(Pt{0, 0.5}), got false")
	}
	if r.Contains(Pt{-1, -1}) {
		t.Errorf("expected false in Contains(Pt{-1, -1}), got true")
	}
}
