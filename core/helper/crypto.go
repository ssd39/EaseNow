package helper

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"os"

	"github.com/edgelesssys/ego/ecrypto"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spruceid/siwe-go"
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

func Eip191Hash(m *siwe.Message) common.Hash {
	// Ref: https://stackoverflow.com/questions/49085737/geth-ecrecover-invalid-signature-recovery-id
	data := []byte(m.String())
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256Hash([]byte(msg))
}
