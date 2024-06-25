package app

import (
	"bufio"
	"os"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	EaseNow "github.com/ssd39/easenow/core/easeNow"
	"github.com/ssd39/easenow/core/helper"
)

func Init(seedPath string) (*helper.Seed, error) {
	privateKeyBytes, err := helper.UnSealKey(seedPath)
	if err != nil {
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			return nil, err
		}
		privateKeyBytes := crypto.FromECDSA(privateKey)
		walletAddress, err := helper.GetWalletAddress(privateKey)
		if err != nil {
			return nil, err
		}
		err = helper.SealKeyToFile(privateKeyBytes, seedPath)
		if err != nil {
			return nil, err
		}
		Logger.Info("Wallet address: " + walletAddress)
		Logger.Info(("Please hit enter after topping up balance!"))
		input := bufio.NewScanner(os.Stdin)
		input.Scan()

		seed := &helper.Seed{PrivateKey: privateKey, WalletAddress: walletAddress}
		client, err := helper.GetEthClient()
		if err != nil {
			return nil, err
		}
		auth, err := helper.GetAuth(seed)
		if err != nil {
			return nil, err
		}
		instance, err := EaseNow.NewEaseNow(EaseNow.ContractAddress, client)
		if err != nil {
			return nil, err
		}
		tx, err := instance.Init(auth, [32]byte{})
		if err != nil {
			return nil, err
		}
		Logger.Info("Core Initiated!", "tx", tx.Hash().Hex())
		return seed, nil
	}

	privateKey, err := crypto.HexToECDSA(hexutil.Encode(privateKeyBytes)[2:])
	if err != nil {
		return nil, err
	}
	walletAddress, err := helper.GetWalletAddress(privateKey)
	if err != nil {
		return nil, err
	}
	Logger.Info("Wallet address: " + walletAddress)
	return &helper.Seed{PrivateKey: privateKey, WalletAddress: walletAddress}, nil
}
