package BalanceOperation

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/heaptracetechnology/microservice-stripe/result"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/balance"
)

//GetBalance
func GetBalance(w http.ResponseWriter, r *http.Request) {

	stripe.Key = os.Getenv("SECRET_KEY")
	b, err := balance.Get(nil)
	if err != nil {
		result.WriteErrorResponse(w, err)
		return
	}
	bytes, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	result.WriteJsonResponse(w, bytes, http.StatusOK)
}
