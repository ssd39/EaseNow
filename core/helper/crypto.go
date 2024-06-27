package helper

import (
	"crypto/ecdsa"
	"errors"
	"os"

	"github.com/edgelesssys/ego/ecrypto"
	"github.com/ethereum/go-ethereum/crypto"
)

func SealKeyToFile(seed []byte, filePath string) error {
	out, err := ecrypto.SealWithUniqueKey(seed, nil)
	if err != nil {
		return err
	}
	err = os.WriteFile(filePath, out, 0600)
	if err != nil {
		return err
	}
	return nil
}

func UnSealKey(filePath string) ([]byte, error) {
	encSeed, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	seed, err := ecrypto.Unseal(encSeed, nil)
	if err != nil {
		return nil, err
	}
	return seed, nil
}

func GetWalletAddress(privateKey *ecdsa.PrivateKey) (string, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("error casting public key to ECDSA")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return address, nil
}

func GetWalletAddressFromPubKey(publicKeyECDSA *ecdsa.PublicKey) string {
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return address
}
