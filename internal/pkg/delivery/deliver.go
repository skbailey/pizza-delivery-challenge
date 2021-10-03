package delivery

import (
	"pizza-delivery/internal/models"
	"pizza-delivery/internal/pkg/parser"
	"pizza-delivery/internal/pkg/validation"
)

// Deliver coordinates the delivery of pizzas
func Deliver(input string, delivererCount int) (int, error) {
	err := validation.ValidateDeliveryParams(input, delivererCount)
	if err != nil {
		return 0, err
	}

	itinerary := parser.Parse(input)
	if err != nil {
		return 0, err
	}

	// Initialize deliverers and tracker
	deliverers, tracker := initialize(delivererCount)

	// Deliver pizza
	for idx, direction := range itinerary {
		// Select deliverer
		delivererIdx := idx % len(deliverers)
		gps := deliverers[delivererIdx]

		// Move and deliver
		next, err := gps.Move(direction)
		if err != nil {
			// Since we validate the input we should never get here. But in the exceptional
			// case that we do enter this conditional, handle the error
			return 0, err
		}

		// Track location
		tracker.AddLocation(next)
	}

	return tracker.CountUniqueLocations(), nil
}

func initialize(delivererCount int) ([]*models.GPS, models.Tracker) {
	current := models.NewLocation(0, 0)
	deliverers := make([]*models.GPS, delivererCount)
	for idx := range deliverers {
		deliverers[idx] = models.NewGPS(current)
	}

	tracker := models.NewTracker()
	tracker.AddLocation(current)

	return deliverers, tracker
}
