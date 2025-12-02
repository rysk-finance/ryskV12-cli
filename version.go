package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var versionAction = &cli.Command{
	Name: "version",
	Action: func(c *cli.Context) error {
		fmt.Println(Version)
		return nil
	},
}
