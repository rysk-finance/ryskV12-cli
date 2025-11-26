package main

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	crypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/goccy/go-json"
)

var ZeroAddress = common.HexToAddress("0x0")

var EIP712_TYPES = &apitypes.Types{
	"EIP712Domain": {
		{
			Name: "name",
			Type: "string",
		},
		{
			Name: "version",
			Type: "string",
		},
		{
			Name: "chainId",
			Type: "uint256",
		},
		{
			Name: "verifyingContract",
			Type: "address",
		},
	},
	"Quote": {
		{
			Name: "assetAddress",
			Type: "address",
		},
		{
			Name: "chainId",
			Type: "uint256",
		},
		{
			Name: "isPut",
			Type: "bool",
		},
		{
			Name: "strike",
			Type: "uint256",
		},
		{
			Name: "expiry",
			Type: "uint64",
		},
		{
			Name: "maker",
			Type: "address",
		},
		{
			Name: "nonce",
			Type: "uint64",
		},
		{
			Name: "price",
			Type: "uint256",
		},
		{
			Name: "quantity",
			Type: "uint256",
		},
		{
			Name: "isTakerBuy",
			Type: "bool",
		},
		{
			Name: "validUntil",
			Type: "uint64",
		},
		{
			Name: "usd",
			Type: "address",
		},
	},
	"Transfer": {
		{
			Name: "asset",
			Type: "address",
		},
		{
			Name: "chainId",
			Type: "uint256",
		},
		{
			Name: "amount",
			Type: "uint256",
		},
		{
			Name: "isDeposit",
			Type: "bool",
		},
		{
			Name: "nonce",
			Type: "uint64",
		},
	},
}

// EncodeTypedData - Encoding the typed data
func EncodeTypedData(typedData *apitypes.TypedData) (common.Hash, error) {
	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return common.BytesToHash([]byte{}), err
	}
	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return common.BytesToHash([]byte{}), err
	}
	rawData := fmt.Appendf(nil, "\x19\x01%s%s", string(domainSeparator), string(typedDataHash))
	hash := common.BytesToHash(crypto.Keccak256(rawData))
	return hash, err
}

// Signs msg with EIP712 signing scheme
func Sign(message []byte, privateKey string) (string, error) {
	privateKeyEcdsa, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", err
	}
	sigBytes, err := signTypedData(message, privateKeyEcdsa)
	if err != nil {
		return "", err
	}
	signature := fmt.Sprintf("0x%s", common.Bytes2Hex(sigBytes))
	return signature, nil
}

func signTypedData(message []byte, privateKey *ecdsa.PrivateKey) (sig []byte, err error) {
	sig, err = crypto.Sign(message, privateKey)
	if err != nil {
		return sig, err
	}
	sig[64] += 27
	return
}

func createEIP712Domain(chainId int64) *apitypes.TypedDataDomain {
	return &apitypes.TypedDataDomain{
		Name:              "rysk",
		Version:           "0.0.0",
		ChainId:           math.NewHexOrDecimal256(chainId),
		VerifyingContract: ZeroAddress.String(),
	}
}

func createEIP712TypedData(chainId int64, msgType string, msg map[string]interface{}) *apitypes.TypedData {
	return &apitypes.TypedData{
		Types:       *EIP712_TYPES,
		PrimaryType: msgType,
		Domain:      *createEIP712Domain(chainId),
		Message:     msg,
	}
}

func CreateQuoteMessage(q Quote) (messageHash []byte, typedData *apitypes.TypedData, err error) {
	msg, _ := json.Marshal(q)
	var imessage map[string]interface{}
	json.Unmarshal(msg, &imessage)
	// remove extra fields
	delete(imessage, "signature")
	typedData = createEIP712TypedData(int64(q.ChainID), "Quote", imessage)
	hash, err := EncodeTypedData(typedData)
	if err != nil {
		return nil, typedData, err
	}
	messageHash = hash.Bytes()
	return messageHash, typedData, nil
}

func CreateTransferMessage(t Transfer) (messageHash []byte, typedData *apitypes.TypedData, err error) {
	msg, _ := json.Marshal(t)
	var imessage map[string]interface{}
	json.Unmarshal(msg, &imessage)
	// remove extra fields
	delete(imessage, "signature")
	typedData = createEIP712TypedData(int64(t.ChainID), "Transfer", imessage)
	hash, err := EncodeTypedData(typedData)
	if err != nil {
		return nil, typedData, err
	}
	messageHash = hash.Bytes()
	return messageHash, typedData, nil
}
