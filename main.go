package main

import (
	"fmt"
	"log"
)

// Location identifies houses along the pizza delivery route
type Location struct {
	X int
	Y int
}

func main() {
	current := Location{0, 0}
	locations := []Location{current}
	input := []string{"^", ">", "v", "<"}

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

		locations = append(locations, next)
		current = next
	}

	fmt.Printf("Here are the visited locations: %#v\n", locations)
}
