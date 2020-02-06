package data_structures

import (
	"testing"
	"testing/quick"
)

func Test_HasPt(t *testing.T) {
	fn := func(x Pt, xs []Pt) bool {
		return !HasPt(xs, x) && HasPt(append(xs, x), x)
	}
	if err := quick.Check(fn, nil); err != nil {
		t.Error(err)
	}
}
