package BalanceOperation

import (
	"encoding/json"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/balance"
	"net/http"
	"os"
)

//GetBalance
func GetBalance(w http.ResponseWriter, r *http.Request) {

	stripe.Key = os.Getenv("SECRET_KEY")

	b, err := balance.Get(nil)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}
	bytes, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	WriteJsonResponse(w, bytes, http.StatusOK)
}

func GetResultOK() int {
	return http.StatusOK
}

func WriteErrorResponse(w http.ResponseWriter, err error) {
	msgbytes, _ := json.Marshal(err)
	WriteJsonResponse(w, msgbytes, http.StatusBadRequest)
}

func WriteJsonResponse(w http.ResponseWriter, bytes []byte, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bytes)
}
