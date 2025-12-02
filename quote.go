package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var quoteAction = &cli.Command{
	Name:  "quote",
	Usage: "Send a quote",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "channel_id",
			Required: true,
			Usage:    "the socket id to send messages into",
		},
		&cli.StringFlag{
			Name:     "rfq_id",
			Required: true,
			Usage:    "the rfq id to respond to",
		},
		&cli.StringFlag{
			Name:     "asset",
			Required: true,
			Usage:    "asset address",
		},
		&cli.IntFlag{
			Name:     "chain_id",
			Required: true,
		},
		&cli.Int64Flag{
			Name:     "expiry",
			Required: true,
		},
		&cli.BoolFlag{
			Name: "is_put",
		},
		&cli.BoolFlag{
			Name: "is_taker_buy",
		},
		&cli.StringFlag{
			Name:     "maker",
			Required: true,
		},
		&cli.Uint64Flag{
			Name:     "nonce",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "price",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "quantity",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "strike",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "valid_until",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "usd",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "collateral",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "private_key",
			Required: true,
			Usage:    "private key to sign messages with",
		},
	},
	Action: func(c *cli.Context) error {
		return quote(c)
	},
}

func quote(c *cli.Context) error {
	channelID := c.String("channel_id")
	rfq_id := c.String("rfq_id")
	pk := c.String("private_key")

	payload := JsonRPCRequest{
		JsonRPC: "2.0",
		ID:      rfq_id,
		Method:  "quote",
	}

	q := Quote{
		AssetAddress: c.String("asset"),
		ChainID:      c.Int("chain_id"),
		Expiry:       c.Int64("expiry"),
		IsPut:        c.Bool("is_put"),
		IsTakerBuy:   c.Bool("is_taker_buy"),
		Maker:        c.String("maker"),
		Nonce:        fmt.Sprintf("%d", c.Uint64("nonce")),
		Price:        c.String("price"),
		Quantity:     c.String("quantity"),
		Strike:       c.String("strike"),
		ValidUntil:   c.Int64("valid_until"),
		USD:          c.String("usd"),
		CollateralAsset: c.String("collateral"),
	}

	msgHash, _, err := CreateQuoteMessage(q)
	if err != nil {
		return err
	}
	sig, err := Sign(msgHash, pk)
	if err != nil {
		return err
	}
	q.Signature = sig
	payload.Params = q

	return writeToSocket(channelID, payload)
}
