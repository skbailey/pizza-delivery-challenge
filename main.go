package main

import (
	"fmt"
	"log"
	"pizza-delivery/internal/models"
)

func main() {
	// Validate input
	input := []string{"^", ">", "v", "<"}

	// Handle initial delivery
	current := models.NewLocation(0, 0)
	tracker := models.NewTracker()
	tracker.AddLocation(current)

	// Deliver pizza
	gps := models.NewGPS(current)
	for _, instruction := range input {
		next, err := gps.Move(instruction)
		if err != nil {
			// Since we validate the input we should never get here. But in the exceptional
			// case that we do enter this conditional, terminate the program
			log.Fatal(err)
		}

		tracker.AddLocation(next)
	}

	fmt.Printf("Unique locations: %#v\n", tracker.GetUniqueLocations())
}
