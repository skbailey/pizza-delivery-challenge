package main

import (
	"flag"
	"log"
	"pizza-delivery/internal/pkg/delivery"
)

func main() {
	var delivererCount int
	flag.IntVar(&delivererCount, "deliverer-count", 1, "number of deliverers available")
	flag.Parse()

	input := "^v^v^v^v^v"
	housesVisited, err := delivery.Deliver(input, delivererCount)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("number of houses receiving at least one pizza: %d\n", housesVisited)
}
