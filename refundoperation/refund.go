package refundoperation

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/heaptracetechnology/microservice-stripe/result"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/refund"
)

//RefundParams Struct
type RefundParams struct {
	Amount               *int64  `json:"amount,omitempty"`
	Charge               *string `json:"charge,omitempty"`
	Reason               *string `json:"reason,omitempty"`
	RefundApplicationFee *bool   `json:"refund_application_fee,omitempty"`
	ReverseTransfer      *bool   `json:"reverse_transfer,omitempty"`
}

//CreateRefund stripe
func CreateRefund(responseWriter http.ResponseWriter, request *http.Request) {
	stripe.Key = os.Getenv("SECRET_KEY")

	decoder := json.NewDecoder(request.Body)
	var param RefundParams
	err := decoder.Decode(&param)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
	}

	properties := stripe.RefundParams{
		Amount:               param.Amount,
		Charge:               param.Charge,
		Reason:               param.Reason,
		RefundApplicationFee: param.RefundApplicationFee,
		ReverseTransfer:      param.ReverseTransfer,
	}
	refunded, err := refund.New(&properties)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
	}
	bytes, _ := json.Marshal(refunded)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusCreated)
}
