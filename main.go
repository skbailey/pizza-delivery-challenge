package main

import (
	"flag"
	"fmt"
	"log"
	"pizza-delivery/internal/models"
	"strings"
)

func main() {
	var delivererCount int
	flag.IntVar(&delivererCount, "deliverer-count", 1, "number of deliverers available")
	flag.Parse()

	// Initialize deliverers
	current := models.NewLocation(0, 0)
	deliverers := make([]*models.GPS, delivererCount)
	for idx := range deliverers {
		deliverers[idx] = models.NewGPS(current)
	}

	tracker := models.NewTracker()
	tracker.AddLocation(current)

	// Validate input
	// input := []string{"^", ">", "v", "<"}
	input := strings.Split("^v^v^v^v^v", "")

	// Deliver pizza
	for idx, instruction := range input {
		delivererIdx := idx % len(deliverers)
		gps := deliverers[delivererIdx]
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
