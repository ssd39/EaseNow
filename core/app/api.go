package app

import (
	"encoding/json"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/mux"
	"github.com/spruceid/siwe-go"
	EaseNow "github.com/ssd39/easenow/core/easeNow"
	"github.com/ssd39/easenow/core/helper"
)

var seed *helper.Seed

var tempUserMap = make(map[string]UserTempModel)

func StartApi(seed_ *helper.Seed) error {
	seed = seed_
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/register-start", registerStart).Methods("POST")
	r.HandleFunc("/register-validate", registerValidateOtp).Methods("POST")
	r.HandleFunc("/register-finalise", registerFinalise).Methods("POST")

	// Bind to a port and pass our router in
	Logger.Info("Started api on 4000")
	return http.ListenAndServe(":4000", r)
}

func registerStart(w http.ResponseWriter, r *http.Request) {
	var payload RegisterStartPayload
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil || payload.CoutnryCode == 0 || payload.PhoneNumber == "" {
		WriteBadReq(&w, "Please fill all details!")
		return
	}
	sessionId := helper.RandStringRunes(14)
	// here using hardcoded otp for mvp
	tempUserMap[sessionId] = UserTempModel{PhoneNumber: payload.PhoneNumber, CoutnryCode: payload.CoutnryCode, Otp: 111111}
	WriteJsonResponse(&w, http.StatusOK, &RegisterResponse{SessionId: sessionId, Success: true})
}

func registerValidateOtp(w http.ResponseWriter, r *http.Request) {
	var payload RegisterValidatePayload
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil || payload.SessionId == "" {
		WriteBadReq(&w, "Bad request payload!")
		return
	}
	userTempData, ok := tempUserMap[payload.SessionId]
	if !ok {
		WriteBadReq(&w, "Session not found!")
		return
	}
	if userTempData.Otp != payload.Otp {
		WriteBadReq(&w, "Wrong otp!")
		return
	}
	userTempData.IsOtpValidated = true
	// for mvp genrating random credit score in production use multiple thirdparty api to validate the score
	userTempData.CreditLimit = helper.EtherToWei(big.NewFloat(float64(helper.RandRange(1, 4)) - ((float64(helper.RandRange(1, 9)) / 10) * float64(helper.RandRange(0, 2)))))
	tempUserMap[payload.SessionId] = userTempData
	res := RegisterValidateResponse{
		CreditLimit: userTempData.CreditLimit,
		Success:     true,
	}
	WriteJsonResponse(&w, http.StatusOK, &res)
}

func registerFinalise(w http.ResponseWriter, r *http.Request) {
	var payload RegisterFinalisePayload
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil || payload.SessionId == "" {
		WriteBadReq(&w, "Please fill all details!")
		return
	}
	// here use the facial data and private data to confirm the person and if things gets confirmed store them on-chain
	userTempData, ok := tempUserMap[payload.SessionId]
	if !ok {
		WriteBadReq(&w, "Session not found!")
		return
	}

	var message *siwe.Message
	message, err := siwe.ParseMessage(payload.WalletMessage)
	if err != nil {
		WriteBadReq(&w, "Error while parsing siwe message!")
		return
	}
	publicKey, err := message.VerifyEIP191(payload.WalletSignature)
	if err != nil {
		WriteBadReq(&w, "Authentication failed!")
		return
	}

	if message.GetNonce() != payload.SessionId {
		WriteBadReq(&w, "Wrong session id provided!")
		return
	}

	walletAddress := helper.GetWalletAddressFromPubKey(publicKey)

	client, err := helper.GetEthClient()
	if err != nil {
		Logger.Error("Failed to get eth client", "error", err)
		WriteBadReq(&w, "Internal server error!")
		return
	}

	auth, err := helper.GetAuth(seed)
	if err != nil {
		Logger.Error("Failed to get internal wallet auth", "error", err)
		WriteBadReq(&w, "Internal server error!")
		return
	}

	instance, err := EaseNow.NewEaseNow(EaseNow.ContractAddress, client)
	if err != nil {
		Logger.Error("Error while getting cotract instance!", "error", err)
		WriteBadReq(&w, "Internal server error!")
		return
	}
	tx, err := instance.RegisterUser(auth, common.HexToAddress(walletAddress), userTempData.CreditLimit, [32]byte([]byte(payload.PrivateData)[:32]))
	if err != nil {
		Logger.Error("Failed to register the user!", "error", err)
		WriteBadReq(&w, "Register user TX failed!")
		return
	}

	res := RegisterFinaliseResponse{TxHash: tx.Hash().Hex(), Success: true}
	WriteJsonResponse(&w, http.StatusOK, &res)
}

func WriteBadReq(w *http.ResponseWriter, message string) {
	WriteJsonResponse(w, http.StatusBadRequest, &ErrResponse{Success: false, Message: message})
}

func WriteJsonResponse(w *http.ResponseWriter, statusCode int, v any) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(statusCode)
	jsonRes, _ := json.Marshal(v)
	(*w).Write(jsonRes)
}
