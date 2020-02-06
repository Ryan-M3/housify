package data_structures

import (
	"testing"
)

func Test_Sum(t *testing.T) {
	s := Sum([]float64{1, 3, 5, 6})
	if s != 15.0 {
		t.Errorf("Got %.1f, but expected 15.0.", s)
	}
}
