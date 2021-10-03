package main

import (
	"flag"
	"log"
	"pizza-delivery/internal/pkg/delivery"
	"pizza-delivery/internal/pkg/reader"
)

func main() {
	var delivererCount int
	flag.IntVar(&delivererCount, "deliverer-count", 1, "number of deliverers available")

	var fileInput string
	flag.StringVar(&fileInput, "file-input", "", "path to the input file")
	flag.Parse()

	input, err := reader.Read(fileInput)
	if err != nil {
		log.Fatal(err)
	}

	housesVisited, err := delivery.Deliver(input, delivererCount)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("number of houses receiving at least one pizza: %d\n", housesVisited)
}
