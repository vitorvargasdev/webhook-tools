package cmd

import "github.com/urfave/cli/v2"

func getFlags(flags *Flags) *[]cli.Flag {
	return &[]cli.Flag{
		&cli.IntFlag{
			Name:        "port",
			Aliases:     []string{"p"},
			Value:       8080,
			Usage:       "Port to listen",
			Destination: &flags.port,
		},
		&cli.IntFlag{
			Name:        "httpCode",
			Aliases:     []string{"c"},
			Value:       200,
			Usage:       "The HTTP status code that will be returned",
			Destination: &flags.httpCode,
		},
		&cli.StringFlag{
			Name:        "formatOutput",
			Aliases:     []string{"f"},
			Value:       "md",
			Usage:       "format output",
			Destination: &flags.formatOutput,
		},
		&cli.StringFlag{
			Name:        "output",
			Aliases:     []string{"o"},
			Value:       "logs",
			Usage:       "output directory",
			Destination: &flags.output,
		},
		&cli.BoolFlag{
			Name:        "web",
			Aliases:     []string{"w"},
			Value:       false,
			Usage:       "expose to worldwide",
			Destination: &flags.web,
		},
	}
}
