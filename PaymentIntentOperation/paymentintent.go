package PaymentIntentOperation

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/heaptracetechnology/microservice-stripe/result"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
	"net/http"
	"os"
)

type StripeObject struct {
	Paymentintent     string
	PaymentMethodType string
}

type Message struct {
	Success string      `json:"success"`
	Message interface{} `json:"Message"`
}

//Create PaymentIntent
func CreatePaymentIntent(w http.ResponseWriter, r *http.Request) {

	stripe.Key = os.Getenv("SECRET_KEY")

	decoder := json.NewDecoder(r.Body)
	var param *stripe.PaymentIntentParams
	err := decoder.Decode(&param)
	if err != nil {
		result.WriteErrorResponse(w, err)
		return
	}

	pi, err := paymentintent.New(param)
	if err != nil {
		result.WriteErrorResponse(w, err)
		return
	}
	bytes, err := json.Marshal(pi)
	if err != nil {
		result.WriteErrorResponse(w, err)
		return
	}
	result.WriteJsonResponse(w, bytes, http.StatusCreated)
}

//Retrieve PaymentIntent
func RetrievePaymentIntent(w http.ResponseWriter, r *http.Request) {

	stripe.Key = os.Getenv("SECRET_KEY")

	vars := mux.Vars(r)
	var id = vars["paymentintentid"]

	pi, err := paymentintent.Get(id, nil)
	if err != nil {
		result.WriteErrorResponse(w, err)
		return
	}
	bytes, err := json.Marshal(pi)
	if err != nil {
		result.WriteErrorResponse(w, err)
		return
	}
	result.WriteJsonResponse(w, bytes, http.StatusOK)
}

//Update PaymentIntent
func UpdatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	stripe.Key = os.Getenv("SECRET_KEY")

	vars := mux.Vars(r)
	var id = vars["paymentintentid"]

	decoderpi := json.NewDecoder(r.Body)
	var param *stripe.PaymentIntentParams
	err := decoderpi.Decode(&param)
	if err != nil {
		result.WriteErrorResponse(w, err)
		return
	}

	param.PaymentMethodTypes = []*string{stripe.String("card")}

	o, err := paymentintent.Update(id, param)
	if err != nil {
		result.WriteErrorResponse(w, err)
		return
	}
	bytes, err := json.Marshal(o)
	if err != nil {
		result.WriteErrorResponse(w, err)
		return
	}
	result.WriteJsonResponse(w, bytes, http.StatusOK)
}

//Capture PaymentIntent
func CapturePaymentIntent(w http.ResponseWriter, r *http.Request) {
	stripe.Key = os.Getenv("SECRET_KEY")
	vars := mux.Vars(r)
	var id = vars["paymentintentid"]

	decoderpi := json.NewDecoder(r.Body)
	var param *stripe.PaymentIntentCaptureParams
	err := decoderpi.Decode(&param)
	if err != nil {
		result.WriteErrorResponse(w, err)
		return
	}

	intent, err := paymentintent.Capture(id, param)
	if err != nil {
		result.WriteErrorResponse(w, err)
		return
	}
	bytes, err := json.Marshal(intent)
	if err != nil {
		result.WriteErrorResponse(w, err)
		return
	}
	result.WriteJsonResponse(w, bytes, http.StatusOK)
}

//Cancel PaymentIntent
func CancelPaymentIntent(w http.ResponseWriter, r *http.Request) {
	stripe.Key = os.Getenv("SECRET_KEY")
	vars := mux.Vars(r)
	var id = vars["paymentintentid"]

	intent, err := paymentintent.Cancel(id, nil)
	if err != nil {
		result.WriteErrorResponse(w, err)
		return
	}
	bytes, err := json.Marshal(intent)
	if err != nil {
		result.WriteErrorResponse(w, err)
		return
	}
	result.WriteJsonResponse(w, bytes, http.StatusOK)
}

//List All PaymentIntent
func ListAllPaymentIntent(w http.ResponseWriter, r *http.Request) {
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
		result.WriteErrorResponse(w, err)
		return
	}
	result.WriteJsonResponse(w, bytes, http.StatusOK)

}
