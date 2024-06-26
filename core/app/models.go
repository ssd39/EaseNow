package app

type RegisterStartPayload struct {
	PhoneNumber string `json:"phone_number"`
	CoutnryCode uint   `json:"country_code"`
}

type RegisterResponse struct {
	SessionId string `json:"session_id"`
	Success   bool   `json:"success"`
}

type RegisterFinalisePayload struct {
	EmailId       string `json:"email_id"`
	FacialData    string `json:"facial_data"`
	PrivateData   string `json:"private_data"`
	IsNewWallet   bool   `json:"is_new_wallet"`
	WalletAddress string `json:"wallet_address"`
	WalletAuth    string `json:"wallet_auth"`
}

type RegisterFinaliseResponse struct {
	Success   bool   `json:"success"`
	TxHash    string `json:"tx_hash"`
	AuthToken string `json:"auth_token"`
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
	PhoneNumber string `json:"phone_number"`
	CoutnryCode uint   `json:"country_code"`
	CreditLimit int    `json:"credit_limit"`
}

type ErrResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
