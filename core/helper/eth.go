package helper

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var ChainId int64 = 1337

func GetEthClient() (*ethclient.Client, error) {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetAuth(seed *Seed) (*bind.TransactOpts, error) {
	client, err := GetEthClient()
	if err != nil {
		return nil, err
	}
	nonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress(seed.WalletAddress))
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(seed.PrivateKey, big.NewInt(ChainId))
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	return auth, nil
}
