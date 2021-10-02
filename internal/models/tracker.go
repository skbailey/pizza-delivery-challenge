package models

// Tracker keeps track of unique locations visited
type Tracker struct {
	locationCounter map[Location]int
}

// NewTracker creates a new Tracker
func NewTracker() Tracker {
	return Tracker{
		locationCounter: map[Location]int{},
	}
}

// AddLocation add a location to the tracker
func (t Tracker) AddLocation(loc Location) {
	count, ok := t.locationCounter[loc]
	if ok {
		t.locationCounter[loc] = count + 1
	} else {
		t.locationCounter[loc] = 1
	}
}

// CountUniqueLocations returns a count of unique locations visited
func (t Tracker) CountUniqueLocations() int {
	locations := make([]Location, 0)

	for l := range t.locationCounter {
		locations = append(locations, l)
	}

	return len(locations)
}
