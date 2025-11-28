package main

import (
	"fmt"
	"time"

	"github.com/urfave/cli/v2"
)

var transferAction = &cli.Command{
	Name:  "transfer",
	Usage: "request a transfer",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "channel_id",
			Required: true,
			Usage:    "the socket id to send messages into",
		},
		&cli.Int64Flag{
			Name:     "chain_id",
			Required: true,
			Usage:    "chain_id",
		},
		&cli.StringFlag{
			Name:     "asset",
			Required: true,
			Usage:    "asset address",
		},
		&cli.StringFlag{
			Name:     "amount",
			Required: true,
			Usage:    "amount to deposit",
		},
		&cli.BoolFlag{
			Name:  "is_deposit",
			Usage: "whether you want to deposit or withdraw",
		},
		&cli.Uint64Flag{
			Name:     "nonce",
			Required: true,
			Usage:    "nonce to sign the message with",
		},
		&cli.StringFlag{
			Name:     "private_key",
			Required: true,
			Usage:    "private key to sign messages with",
		},
	},
	Action: func(c *cli.Context) error {
		return transfer(c)
	},
}

func transfer(c *cli.Context) error {
	channelID := c.String("channel_id")
	method := "withdraw"
	if c.Bool("is_deposit") {
		method = "deposit"
	}

	payload := JsonRPCRequest{
		JsonRPC: "2.0",
		ID:      fmt.Sprintf("%d", time.Now().Unix()),
		Method:  method,
	}

	t := Transfer{
		User:      c.String("user"),
		Asset:     c.String("asset"),
		ChainID:   int(c.Int64("chain_id")),
		Amount:    c.String("amount"),
		IsDeposit: c.Bool("is_deposit"),
		Nonce:     fmt.Sprintf("%d", c.Uint64("nonce")),
	}

	msgHash, _, err := CreateTransferMessage(t)
	if err != nil {
		return err
	}
	sig, err := Sign(msgHash, c.String("private_key"))
	if err != nil {
		return err
	}
	t.Signature = sig
	payload.Params = t
	return writeToSocket(channelID, payload)
}
