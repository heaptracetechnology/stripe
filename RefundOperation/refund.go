package RefundOperation

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/heaptracetechnology/microservice-stripe/result"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/refund"
)

//Create Refund
func CreateRefund(w http.ResponseWriter, r *http.Request) {
	stripe.Key = os.Getenv("SECRET_KEY")

	decoder := json.NewDecoder(r.Body)
	var param *stripe.RefundParams
	err := decoder.Decode(&param)
	if err != nil {
		result.WriteErrorResponse(nil, err)
	}

	refunded, err := refund.New(param)
	if err != nil {
		result.WriteErrorResponse(nil, err)
	}
	bytes, _ := json.Marshal(refunded)
	result.WriteJsonResponse(w, bytes, http.StatusCreated)
}
