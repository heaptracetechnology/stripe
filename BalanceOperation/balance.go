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
		msgbytes, _ := json.Marshal(err)
		writeJsonResponse(w, msgbytes)
		return
	}
	bytes, _ := json.Marshal(b)
	writeJsonResponse(w, bytes)
}

func writeJsonResponse(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bytes)
}
