package models

// Location identifies houses along the pizza delivery route
type Location struct {
	X int
	Y int
}

// NewLocation creates a new Location
func NewLocation(x, y int) Location {
	return Location{
		X: x,
		Y: y,
	}
}
