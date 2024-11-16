package main

import (
	"log"
	"os"

	"github.com/vitorvargasdev/webhook-tools/cmd"
)

func main() {
  cli := cmd.CLI{}
  if err := cli.Prepare().App.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
