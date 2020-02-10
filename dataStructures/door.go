package dataStructures

type CardinalDirection int

const (
	N CardinalDirection = iota
	E CardinalDirection = iota
	S CardinalDirection = iota
	W CardinalDirection = iota
)

type Door struct {
	Orientation CardinalDirection
	Position    float64
}
