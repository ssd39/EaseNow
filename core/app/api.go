package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ssd39/easenow/core/helper"
)

var seed *helper.Seed

var tempUserMap = make(map[string]UserTempModel)

func StartApi(seed_ *helper.Seed) error {
	seed = seed_
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/register-start", registerStart).Methods("POST")
	r.HandleFunc("/register-finalise", registerFinalise).Methods("POST")

	// Bind to a port and pass our router in
	Logger.Info("Started api on 8000")
	return http.ListenAndServe(":8000", r)
}

func registerStart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var payload RegisterStartPayload
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil || payload.CoutnryCode == 0 || payload.PhoneNumber == "" {
		w.WriteHeader(http.StatusBadRequest)
		errRes := ErrResponse{Success: false, Message: "Please fill all details!"}
		jsonRes, _ := json.Marshal(errRes)
		w.Write(jsonRes)
		return
	}
	sessionId := helper.RandStringRunes(12)
	tempUserMap[sessionId] = UserTempModel{PhoneNumber: payload.PhoneNumber, CoutnryCode: payload.CoutnryCode, CreditLimit: helper.RandRange(1, 3)}
	w.WriteHeader(http.StatusOK)
	res := RegisterResponse{SessionId: sessionId, Success: true}
	jsonRes, _ := json.Marshal(res)
	w.Write(jsonRes)
}

func registerFinalise(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Gorilla!\n"))
}
