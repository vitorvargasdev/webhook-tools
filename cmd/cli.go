package cmd

import (
	"github.com/urfave/cli/v2"
)

type CLI struct {
  App *cli.App
}

func (c CLI) Prepare() CLI {
  app := cli.NewApp()
	app.Name = "Webhook Tools"
	app.Usage = "Utilities to help with webhook development"
  c.App = app

  return c 
}

