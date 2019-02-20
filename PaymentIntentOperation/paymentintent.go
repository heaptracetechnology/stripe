package PaymentIntentOperation

import (
	"encoding/json"
	"github.com/gorilla/mux"
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
		WriteErrorResponse(w, err)
		return
	}

	param.PaymentMethodTypes = []*string{stripe.String("card")}

	pi, err := paymentintent.New(param)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}
	bytes, err := json.Marshal(pi)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}
	WriteJsonResponse(w, bytes, http.StatusCreated)
}

//Retrieve PaymentIntent
func RetrievePaymentIntent(w http.ResponseWriter, r *http.Request) {

	stripe.Key = os.Getenv("SECRET_KEY")

	vars := mux.Vars(r)
	var id = vars["paymentintent_id"]

	pi, err := paymentintent.Get(id, nil)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}
	bytes, err := json.Marshal(pi)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}
	WriteJsonResponse(w, bytes, http.StatusOK)
}

//Update PaymentIntent
func UpdatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	stripe.Key = os.Getenv("SECRET_KEY")

	vars := mux.Vars(r)
	var id = vars["paymentintent_id"]

	decoderpi := json.NewDecoder(r.Body)
	var param *stripe.PaymentIntentParams
	err := decoderpi.Decode(&param)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}

	o, err := paymentintent.Update(id, param)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}
	bytes, err := json.Marshal(o)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}
	WriteJsonResponse(w, bytes, http.StatusOK)
}

//Confirm PaymentIntent
func ConfirmPaymentIntent(w http.ResponseWriter, r *http.Request) {

	stripe.Key = os.Getenv("SECRET_KEY")

	params := &stripe.PaymentIntentConfirmParams{
		Source: stripe.String("src_1E5PV5JytX7n0OoXVhMXhy4N"),
	}
	intent, err := paymentintent.Confirm("pi_1E5Q7FJytX7n0OoXaCfqt9KD", params)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}
	bytes, err := json.Marshal(intent)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}
	WriteJsonResponse(w, bytes, http.StatusOK)
}

//Capture PaymentIntent
func CapturePaymentIntent(w http.ResponseWriter, r *http.Request) {
	stripe.Key = os.Getenv("SECRET_KEY")
	vars := mux.Vars(r)
	var id = vars["paymentintent_id"]

	decoderpi := json.NewDecoder(r.Body)
	var param *stripe.PaymentIntentCaptureParams
	err := decoderpi.Decode(&param)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}

	intent, err := paymentintent.Capture(id, param)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}
	bytes, err := json.Marshal(intent)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}
	WriteJsonResponse(w, bytes, http.StatusOK)
}

//Cancel PaymentIntent
func CancelPaymentIntent(w http.ResponseWriter, r *http.Request) {
	stripe.Key = os.Getenv("SECRET_KEY")
	vars := mux.Vars(r)
	var id = vars["paymentintent_id"]

	intent, err := paymentintent.Cancel(id, nil)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}
	bytes, err := json.Marshal(intent)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}
	WriteJsonResponse(w, bytes, http.StatusOK)
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
		WriteErrorResponse(w, err)
		return
	}
	WriteJsonResponse(w, bytes, http.StatusOK)

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
