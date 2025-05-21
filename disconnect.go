package main

import "github.com/urfave/cli/v2"

var disconnectAction = &cli.Command{
	Name: "disconnect",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "channel_id",
			Required: true,
			Usage:    "a unique id so you can later interact with this specific websocket",
		},
	},
	Action: func(c *cli.Context) error {
		return kill(c)
	},
}

func kill(c *cli.Context) error {
	return writeToSocket(c.String("channel_id"), "disconnect")
}
