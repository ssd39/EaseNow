package app

import "math/big"

type RegisterStartPayload struct {
	PhoneNumber string `json:"phone_number"`
	CoutnryCode uint   `json:"country_code"`
}

type RegisterResponse struct {
	SessionId string `json:"session_id"`
	Success   bool   `json:"success"`
}

type RegisterValidatePayload struct {
	SessionId string `json:"session_id"`
	Otp       uint   `json:"otp"`
}

type RegisterValidateResponse struct {
	CreditLimit *big.Int `json:"credit_limit"`
	Success     bool     `json:"success"`
}

type RegisterFinalisePayload struct {
	SessionId       string `json:"session_id"`
	EmailId         string `json:"email_id"`
	PrivateData     string `json:"private_data"`
	WalletMessage   string `json:"wallet_message"`
	WalletSignature string `json:"wallet_sig"`
}

type RegisterFinaliseResponse struct {
	Success bool   `json:"success"`
	TxHash  string `json:"tx_hash"`
	//AuthToken string `json:"auth_token"`
}

type LoginPayload struct {
	WalletAddress string `json:"wallet_address"`
	WalletAuth    string `json:"wallet_auth"`
}

type LoginResponse struct {
	Success   bool   `json:"success"`
	AuthToken string `json:"auth_token"`
}

type UserTempModel struct {
	PhoneNumber    string
	CoutnryCode    uint
	Otp            uint
	IsOtpValidated bool
	CreditLimit    *big.Int
}

type ErrResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
