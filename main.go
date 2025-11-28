package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// updated at build time
var Version = ""

func main() {
	app := &cli.App{
		Name: "ryskV12",
		Commands: []*cli.Command{
			approveAction,
			balancesAction,
			connectAction,
			disconnectAction,
			positionsAction,
			quoteAction,
			transferAction,
			versionAction,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
