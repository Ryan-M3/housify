package dataStructures

type Room struct {
	Rect  *Rect
	Doors []*Door
	Adj   []*Room
}
