package RefundOperation

import (
	"encoding/json"
	"github.com/heaptracetechnology/microservice-stripe/result"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/refund"
	"net/http"
	"os"
)

//Create Refund
func CreateRefund(w http.ResponseWriter, r *http.Request) {
	stripe.Key = os.Getenv("SECRET_KEY")

	decoder := json.NewDecoder(r.Body)
	var param *stripe.RefundParams
	err := decoder.Decode(&param)
	if err != nil {
		result.WriteErrorResponse(w, err)
	}

	refunded, err := refund.New(param)
	if err != nil {
		result.WriteErrorResponse(w, err)
	}
	bytes, _ := json.Marshal(refunded)
	result.WriteJsonResponse(w, bytes, http.StatusCreated)
}
