package paymentintentoperation

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/heaptracetechnology/microservice-stripe/result"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

//PaymentIntentParams Struct
type PaymentIntentParams struct {
	Amount                    *int64            `json:"amount,omitempty"`
	ApplicationFeeAmount      *int64            `json:"applicationFeeAmount,omitempty"`
	CaptureMethod             *string           `json:"captureMethod,omitempty"`
	Confirm                   *bool             `json:"confirm,omitempty"`
	ConfirmationMethod        *string           `json:"confirmationMethod,omitempty"`
	Currency                  *string           `json:"currency,omitempty"`
	Customer                  *string           `json:"customer,omitempty"`
	Description               *string           `json:"description,omitempty"`
	OnBehalfOf                *string           `json:"onBehalfOf,omitempty"`
	PaymentMethod             *string           `json:"paymentMethod,omitempty"`
	PaymentMethodOptions      map[string]string `json:"paymentMethodOptions,omitempty"`
	PaymentMethodTypes        []*string         `json:"paymentMethodTypes,omitempty"`
	ReceiptEmail              *string           `json:"receiptEmail,omitempty"`
	SavePaymentMethod         *bool             `json:"savePaymentMethod,omitempty"`
	SetupFutureUsage          *string           `json:"setupFutureUsage,omitempty"`
	Shipping                  map[string]string `json:"shipping,omitempty"`
	Source                    *string           `json:"source,omitempty"`
	StatementDescriptor       *string           `json:"statementDescriptor,omitempty"`
	StatementDescriptorSuffix *string           `json:"statementDescriptorSuffix,omitempty"`
	TransferData              map[string]string `json:"transferData,omitempty"`
	TransferGroup             *string           `json:"transferGroup,omitempty"`
	OffSession                *bool             `json:"offSession,omitempty"`
}

//PaymentIntentCaptureParams struct
type PaymentIntentCaptureParams struct {
	AmountToCapture      *int64 `form:"amountToCapture,omitempty"`
	ApplicationFeeAmount *int64 `form:"applicationFeeAmount,omitempty"`
}

//PaymentIntentCancelParams struct
type PaymentIntentCancelParams struct {
	CancellationReason *string `form:"cancellationReason,omitempty"`
}

//CreatePaymentIntent Stripe
func CreatePaymentIntent(responseWriter http.ResponseWriter, request *http.Request) {

	stripe.Key = os.Getenv("SECRET_KEY")

	decoder := json.NewDecoder(request.Body)
	var param PaymentIntentParams
	err := decoder.Decode(&param)

	properties := stripe.PaymentIntentParams{
		Amount:               param.Amount,
		ApplicationFeeAmount: param.ApplicationFeeAmount,
		CaptureMethod:        param.CaptureMethod,
		Confirm:              param.Confirm,
		Currency:             param.Currency,
		Customer:             param.Customer,
		Description:          param.Description,
		OnBehalfOf:           param.OnBehalfOf,
		PaymentMethodTypes:   param.PaymentMethodTypes,
		ReceiptEmail:         param.ReceiptEmail,
		Source:               param.Source,
		StatementDescriptor:  param.StatementDescriptor,
		TransferGroup:        param.TransferGroup,
	}

	properties.PaymentMethodTypes = []*string{stripe.String("card")}

	pi, err := paymentintent.New(&properties)

	bytes, err := json.Marshal(pi)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
	}
	result.WriteJsonResponse(responseWriter, bytes, http.StatusCreated)
}

//RetrievePaymentIntent Stripe
func RetrievePaymentIntent(responseWriter http.ResponseWriter, request *http.Request) {

	stripe.Key = os.Getenv("SECRET_KEY")

	vars := mux.Vars(request)
	var id = vars["paymentIntentID"]

	pi, err := paymentintent.Get(id, nil)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
	}

	bytes, _ := json.Marshal(pi)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//UpdatePaymentIntent stripe
func UpdatePaymentIntent(responseWriter http.ResponseWriter, request *http.Request) {
	stripe.Key = os.Getenv("SECRET_KEY")

	vars := mux.Vars(request)
	var id = vars["paymentIntentID"]

	decoderPI := json.NewDecoder(request.Body)
	var param PaymentIntentParams

	decodeErr := decoderPI.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
	}

	properties := stripe.PaymentIntentParams{
		Amount:               param.Amount,
		ApplicationFeeAmount: param.ApplicationFeeAmount,
		CaptureMethod:        param.CaptureMethod,
		Confirm:              param.Confirm,
		Currency:             param.Currency,
		Customer:             param.Customer,
		Description:          param.Description,
		OnBehalfOf:           param.OnBehalfOf,
		PaymentMethodTypes:   param.PaymentMethodTypes,
		ReceiptEmail:         param.ReceiptEmail,
		Source:               param.Source,
		StatementDescriptor:  param.StatementDescriptor,
		TransferGroup:        param.TransferGroup,
	}
	update, updateErr := paymentintent.Update(id, &properties)
	if updateErr != nil {
		result.WriteErrorResponse(responseWriter, updateErr)
	}

	bytes, _ := json.Marshal(update)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

// //ConfirmPaymentIntent stripe
// func ConfirmPaymentIntent(responseWriter http.ResponseWriter, request *http.Request) {

// 	stripe.Key = os.Getenv("SECRET_KEY")
// 	vars := mux.Vars(request)
// 	var id = vars["paymentIntentID"]

// 	decoderPI := json.NewDecoder(request.Body)
// 	var param *stripe.PaymentIntentConfirmParams
// 	decodeErr := decoderPI.Decode(&param)
// 	if decodeErr != nil {
// 		result.WriteErrorResponse(responseWriter, decodeErr)
// 		return
// 	}

// 	// var params1 *stripe.PaymentIntentConfirmParams

// 	// params1 = &stripe.PaymentIntentConfirmParams{
// 	// 	PaymentMethod: stripe.String("pm_card_visa"),
// 	// }
// 	// intent, err := paymentintent.Confirm("pi_1E558UJytX7n0OoXIoSZGbCf", params)

// 	// param.Params = stripe.Params("pm_card_visa")

// 	confirm, confirmErr := paymentintent.Confirm(id, param)
// 	if confirmErr != nil {
// 		result.WriteErrorResponse(responseWriter, confirmErr)
// 		return
// 	}

// 	bytes, _ := json.Marshal(confirm)
// 	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
// }

//CapturePaymentIntent stripe
func CapturePaymentIntent(responseWriter http.ResponseWriter, request *http.Request) {

	stripe.Key = os.Getenv("SECRET_KEY")
	vars := mux.Vars(request)
	var id = vars["paymentIntentID"]

	decoderPI := json.NewDecoder(request.Body)
	var param PaymentIntentCaptureParams
	decodeErr := decoderPI.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	properties := stripe.PaymentIntentCaptureParams{
		AmountToCapture:      param.AmountToCapture,
		ApplicationFeeAmount: param.ApplicationFeeAmount,
	}

	capture, captureErr := paymentintent.Capture(id, &properties)
	if captureErr != nil {
		result.WriteErrorResponse(responseWriter, captureErr)
		return
	}

	bytes, _ := json.Marshal(capture)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//CancelPaymentIntent stripe
func CancelPaymentIntent(responseWriter http.ResponseWriter, request *http.Request) {

	stripe.Key = os.Getenv("SECRET_KEY")
	vars := mux.Vars(request)
	var id = vars["paymentIntentID"]

	var param PaymentIntentCancelParams

	properties := stripe.PaymentIntentCancelParams{
		CancellationReason: param.CancellationReason,
	}

	intent, intentErr := paymentintent.Cancel(id, &properties)
	if intentErr != nil {
		result.WriteErrorResponse(responseWriter, intentErr)
		return
	}

	bytes, _ := json.Marshal(intent)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//ListAllPaymentIntent stripe
func ListAllPaymentIntent(responseWriter http.ResponseWriter, request *http.Request) {

	stripe.Key = os.Getenv("SECRET_KEY")

	params := stripe.PaymentIntentListParams{}
	res := paymentintent.List(&params)

	var listpaymentintent []stripe.PaymentIntent

	for res.Next() {
		paymentintent := stripe.PaymentIntent{}
		listpaymentintent = append(listpaymentintent, paymentintent)
	}

	bytes, err := json.Marshal(&listpaymentintent)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
	}
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}
