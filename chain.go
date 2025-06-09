package main

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	crypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	CHAIN_ID_ETH           int = 1
	CHAIN_ID_BASE          int = 8453
	CHAIN_ID_HYPE	       int = 999
	CHAIN_ID_BASE_SEPOLIA  int = 84532
	CHAIN_ID_HYPE_TESTNET  int = 998
	CHAIN_ID_MONAD_TESTNET int = 10143
)

type Addresses struct {
	MarginPool  common.Address
	MMarket     common.Address
	StrikeAsset common.Address
}

type Account struct {
	Public  common.Address
	Private *ecdsa.PrivateKey
}

var ADDRESSES = map[int]Addresses{
	CHAIN_ID_BASE_SEPOLIA: {
		MarginPool:  common.HexToAddress("0xf0ac2edd8eb37b4ac75a2ed02eeae3a5c6bb8e8e"),
		MMarket:     common.HexToAddress("0x6d23217cb228f3779f6506b44b1abaa3edf6fc58"),
		StrikeAsset: common.HexToAddress("0x98d56648c9b7f3cb49531f4135115b5000ab1733"),
	},
	CHAIN_ID_MONAD_TESTNET: {
		MarginPool:  common.HexToAddress("0x968d1f9dcda00df7be0bbc0acf3761853d06eb33"),
		MMarket:     common.HexToAddress("0x0d71dbdb11c35c72fa4a33beec00d5eb1fbcf1c8"),
		StrikeAsset: common.HexToAddress("0xf817257fed379853cde0fa4f97ab987181b1e5ea"),
	},
	CHAIN_ID_HYPE_TESTNET: {
		MarginPool: common.HexToAddress("0xB586aBB1b69611284A855BAECE6707bf353C035b"),
		MMarket:     common.HexToAddress("0x4bB4dDf8F4588780fe0cD31593bd8f5719eBfefd"),
		StrikeAsset: common.HexToAddress("0x69E37eCb273a23B8864A5f39520Bf97c6768b65C"),
	},
	CHAIN_ID_BASE: {
		MarginPool:  common.HexToAddress("0x6EBec8b078464B5d59eCc1c23F7F37e65b75e61f"),
		MMarket:     common.HexToAddress("0x3D9CB5D2Fa4600bF8d75fB59Fe01Db765dCced15"),
		StrikeAsset: common.HexToAddress("0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913"),
	},
	CHAIN_ID_HYPE: {
		MarginPool:  common.HexToAddress("0x24a44f1dc25540c62c1196FfC297dFC951C91aB4"),
		MMarket:     common.HexToAddress("0x691a5fc3a81a144e36c6C4fBCa1fC82843c80d0d"),
		StrikeAsset: common.HexToAddress("0xB8CE59FC3717ada4C02eaDF9682A9e934F625ebb"),
	},
}

func newAccountFromPrivateKey(pk string) (account Account, err error) {
	account = Account{}
	account.Private, err = crypto.HexToECDSA(strings.TrimPrefix(pk, "0x")) // hex w/o 0x
	if err != nil {
		return account, err
	}

	account.Public = crypto.PubkeyToAddress(account.Private.PublicKey)
	return account, nil

}

func (a *Account) newTransactionOpts(ctx context.Context, chainID int, c ethclient.Client) (*bind.TransactOpts, error) {
	nonce, err := c.PendingNonceAt(ctx, a.Public)
	if err != nil {
		return nil, err
	}
	gasprice, err := c.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}
	opts, err := bind.NewKeyedTransactorWithChainID(a.Private, new(big.Int).SetInt64(int64(chainID)))
	if err != nil {
		return nil, err
	}
	opts.Nonce = big.NewInt(int64(nonce))
	opts.GasPrice = gasprice
	return opts, nil
}

func (a *Account) approve(ctx context.Context, chainID int, client ethclient.Client, amount *big.Int) (txHash string, err error) {
	opts, err := a.newTransactionOpts(ctx, chainID, client)
	if err != nil {
		return "", nil
	}

	erc20, err := NewIERC20(ADDRESSES[chainID].StrikeAsset, &client)
	if err != nil {
		return "", err
	}

	tx, err := erc20.Approve(opts, ADDRESSES[chainID].MMarket, amount)
	if err != nil {
		return "", err
	}

	rx, err := bind.WaitMined(ctx, &client, tx)
	if err != nil {
		return "", err
	}

	return rx.TxHash.Hex(), nil
}
