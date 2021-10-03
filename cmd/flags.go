package cmd

import "github.com/urfave/cli/v2"

var globalFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    "file-input",
		Usage:   "Provide a file containing an itinerary of directions",
		EnvVars: []string{"FILE_INPUT"},
	},
	&cli.IntFlag{
		Name:    "deliverer-count",
		Value:   1,
		Usage:   "Assign the number of deliverers available (who deliver pizza)",
		EnvVars: []string{"DELIVERER_COUNT"},
	},
}
