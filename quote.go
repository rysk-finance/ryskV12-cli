package main

import "github.com/urfave/cli/v2"

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
			Name:     "reply_id",
			Required: true,
			Usage:    "the notification id to respond to",
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
		&cli.StringFlag{
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
	reply_id := c.String("reply_id")
	pk := c.String("private_key")

	payload := JsonRPCRequest{
		JsonRPC: "2.0",
		ID:      reply_id,
		Method:  "quote",
	}

	q := Quote{
		AssetAddress: c.String("asset"),
		ChainID:      c.Int("chain_id"),
		Expiry:       c.Int64("expiry"),
		IsPut:        c.Bool("is_put"),
		IsTakerBuy:   c.Bool("is_taker_buy"),
		Maker:        c.String("maker"),
		Nonce:        c.String("nonce"),
		Price:        c.String("price"),
		Quantity:     c.String("quantity"),
		Strike:       c.String("strike"),
		ValidUntil:   c.Int64("valid_until"),
	}

	msgHash, _, _ := CreateQuoteMessage(q)
	sig, _ := Sign(msgHash, pk)
	q.Signature = sig
	payload.Params = q

	return writeToFifo(channelID, payload)
}
