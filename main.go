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

	for _, instruction := range input {
		var next = current

		switch instruction {
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
			log.Fatalln("unexpected instruction received")
		}

		tracker.AddLocation(next)
		current = next
	}

	fmt.Printf("Unique locations: %#v\n", tracker.GetUniqueLocations())
}
