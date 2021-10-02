package models

// Delivery orchestrates the delivery of pizzas to houses
type Delivery struct {
	itinerary string
}

// NewDelivery creates a new Delivery
func NewDelivery(itinerary string) Delivery {
	return Delivery{
		itinerary: itinerary,
	}
}

// Valid validates the itinerary
func (d Delivery) Valid() bool {
	if d.itinerary == "" {
		return false
	}

	// TODO: Verify the input only contains allowed characters

	return true
}
