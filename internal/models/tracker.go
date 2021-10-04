package models

// Tracker keeps track of unique locations visited
type Tracker struct {
	locationCounter map[Location]struct{}
}

// NewTracker creates a new Tracker
func NewTracker() Tracker {
	return Tracker{
		locationCounter: make(map[Location]struct{}),
	}
}

// AddLocation add a location to the tracker
func (t Tracker) AddLocation(loc Location) {
	t.locationCounter[loc] = struct{}{}
}

// CountUniqueLocations returns a count of unique locations visited
func (t Tracker) CountUniqueLocations() int {
	return len(t.locationCounter)
}
