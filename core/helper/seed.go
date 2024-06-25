package helper

import "crypto/ecdsa"

type Seed struct {
	PrivateKey    *ecdsa.PrivateKey
	WalletAddress string
}
