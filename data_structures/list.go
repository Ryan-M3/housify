package data_structures

// Sum a list of floats.
func Sum(xs []float64) float64 {
	acc := 0.0
	for _, x := range xs {
		acc += x
	}
	return acc
}
