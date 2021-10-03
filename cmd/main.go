package cmd

import (
	"log"
	"os"
	"pizza-delivery/internal/pkg/delivery"
	"pizza-delivery/internal/pkg/reader"
	"sort"

	"github.com/urfave/cli/v2"
)

// Run launches the pizza delivery program
func Run() {
	app := &cli.App{
		Name:    "Pizza Delivery",
		Usage:   "Coordinate delivery staff to deliver pizzas",
		Version: "0.0.1",
		Flags:   globalFlags,
		Action: func(c *cli.Context) error {
			delivererCount := c.Int("deliverer-count")
			fileInput := c.String("file-input")

			input, err := reader.Read(fileInput)
			if err != nil {
				return err
			}

			housesVisited, err := delivery.Deliver(input, delivererCount)
			if err != nil {
				return err
			}

			log.Printf("number of houses receiving at least one pizza: %d\n", housesVisited)
			return nil
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
