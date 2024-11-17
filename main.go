package main

import (
	"log"
	"os"

	"github.com/vitorvargasdev/webhook-tools/cmd"
)

func main() {
	cli := cmd.CLI{}
	cli = cli.Prepare()

	if err := cli.App.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
