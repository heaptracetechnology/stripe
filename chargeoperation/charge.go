package chargeoperation

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/heaptracetechnology/microservice-stripe/result"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

//ChargeParams struct
type ChargeParams struct {
	Amount               *int64            `json:"amount,omitempty"`
	ApplicationFeeAmount *int64            `json:"applicationFeeAmount,omitempty"`
	Capture              *bool             `json:"capture,omitempty"`
	Currency             *string           `json:"currency,omitempty"`
	Customer             *string           `json:"customer,omitempty"`
	Description          *string           `json:"description,omitempty"`
	OnBehalfOf           *string           `json:"onBehalfOf,omitempty"`
	ReceiptEmail         *string           `json:"receiptEmail,omitempty"`
	Shipping             map[string]string `json:"shipping,omitempty"`
	Source               map[string]string `json:"*"`
	TransferData         map[string]string `json:"transferData,omitempty"`
	TransferGroup        *string           `json:"transferGroup,omitempty"`
}

//CaptureParams struct
type CaptureParams struct {
	Amount               *int64   `json:"amount,omitempty"`
	ApplicationFeeAmount *int64   `json:"applicationFeeAmount,omitempty"`
	ExchangeRate         *float64 `json:"exchangeRate,omitempty"`
	ReceiptEmail         *string  `json:"receiptEmail,omitempty"`
	StatementDescriptor  *string  `json:"statementDescriptor,omitempty"`
	ApplicationFee       *int64   `json:"applicationFee,omitempty"`
}

//HealthCheck Stripe
func HealthCheck(responseWriter http.ResponseWriter, request *http.Request) {

	bytes, _ := json.Marshal("OK")
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//CreateCharge Stripe
func CreateCharge(responseWriter http.ResponseWriter, request *http.Request) {

	stripe.Key = os.Getenv("SECRET_KEY")

	decoder := json.NewDecoder(request.Body)
	var param ChargeParams
	err := decoder.Decode(&param)

	properties := stripe.ChargeParams{
		Amount:               param.Amount,
		ApplicationFeeAmount: param.ApplicationFeeAmount,
		Currency:             param.Currency,
		Description:          param.Description,
		Capture:              param.Capture,
		Customer:             param.Customer,
		OnBehalfOf:           param.OnBehalfOf,
		ReceiptEmail:         param.ReceiptEmail,
		Shipping:             &stripe.ShippingDetailsParams{},
		Source:               &stripe.SourceParams{},
		TransferData:         &stripe.ChargeTransferDataParams{},
		TransferGroup:        param.TransferGroup,
	}

	properties.SetSource("tok_visa")

	newCharge, err := charge.New(&properties)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
	}

	bytes, _ := json.Marshal(newCharge)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusCreated)
}

//CaptureCharge Stripe
func CaptureCharge(responseWriter http.ResponseWriter, request *http.Request) {

	stripe.Key = os.Getenv("SECRET_KEY")

	vars := mux.Vars(request)
	var chargeID = vars["charge"]

	decoder := json.NewDecoder(request.Body)
	var param CaptureParams
	err := decoder.Decode(&param)

	properties := stripe.CaptureParams{
		Amount:               param.Amount,
		ApplicationFeeAmount: param.ApplicationFeeAmount,
		ExchangeRate:         param.ExchangeRate,
		ReceiptEmail:         param.ReceiptEmail,
		StatementDescriptor:  param.StatementDescriptor,
		ApplicationFee:       param.ApplicationFee,
	}

	captureCharge, err := charge.Capture(chargeID, &properties)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
	}
	bytes, _ := json.Marshal(captureCharge)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}
