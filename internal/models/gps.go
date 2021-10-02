package models

import "errors"

// GPS directs a delivery driver to a location
type GPS struct {
	position Location
}

// NewGPS creates a new GPS
func NewGPS(position Location) *GPS {
	return &GPS{
		position: position,
	}
}

// Move moves to the next location
func (g *GPS) Move(direction string) (Location, error) {
	var next = g.position

	switch direction {
	case "^":
		// Go north
		next.Y++
	case "v":
		// Go south
		next.Y--
	case ">":
		// Go east
		next.X++
	case "<":
		// Go west
		next.X--
	default:
		return Location{}, errors.New("unexpected instruction received")
	}

	g.position = next
	return next, nil
}
