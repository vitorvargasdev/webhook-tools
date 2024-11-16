package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

type Flags struct {
  port int
  httpCode int
  formatOutput string
  output string
  web bool
}

type CLI struct {
	App *cli.App
  Flags Flags
}

func (c CLI) Prepare() CLI {
	app := cli.NewApp()
	app.Name = "Webhook Tools"
	app.Usage = "Utilities to help with webhook development"

	app.Flags = *getFlags(&c.Flags)

  app.Action = func(cCtx *cli.Context) error {
    fmt.Println(c.Flags)

    return nil
  }

	c.App = app

	return c
}
