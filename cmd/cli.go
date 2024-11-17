package cmd

import (
	"github.com/urfave/cli/v2"
)

type Flags struct {
	port         int
	httpCode     int
	formatOutput string
	output       string
	web          bool
}

type CLI struct {
	App   *cli.App
	Flags Flags
}

func (c CLI) Prepare() CLI {
	app := cli.NewApp()
	app.Name = "Webhook Tools"
	app.Usage = "Utilities to help with webhook development"
	flags := Flags{}

	app.Flags = *getFlags(&flags)
	app.Action = func(c *cli.Context) error {
    return StartServer(flags)
	}

	c.App = app

	return c
}
