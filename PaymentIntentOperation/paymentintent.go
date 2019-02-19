package PaymentIntentOperation

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
	"log"
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
		msgbytes, _ := json.Marshal(err)
		writeJsonResponse(w, msgbytes)
		return
	}

	param.PaymentMethodTypes = []*string{stripe.String("card")}

	pi, err := paymentintent.New(param)
	if err != nil {
		msgbytes, _ := json.Marshal(err)
		writeJsonResponse(w, msgbytes)
		return
	}
	bytes, _ := json.Marshal(pi)
	writeJsonResponse(w, bytes)
}

//Retrieve PaymentIntent
func RetrievePaymentIntent(w http.ResponseWriter, r *http.Request) {

	stripe.Key = os.Getenv("SECRET_KEY")

	vars := mux.Vars(r)
	var id = vars["paymentintent_id"]

	pi, err := paymentintent.Get(id, nil)
	if err != nil {
		log.Fatalln(err)
	}
	bytes, _ := json.Marshal(pi)
	writeJsonResponse(w, bytes)
}

//Update PaymentIntent
func UpdatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	stripe.Key = os.Getenv("SECRET_KEY")

	vars := mux.Vars(r)
	var id = vars["paymentintent_id"]

	decoderpi := json.NewDecoder(r.Body)
	var param *stripe.PaymentIntentParams
	errr := decoderpi.Decode(&param)
	if errr != nil {
		msgbytes, _ := json.Marshal(errr)
		writeJsonResponse(w, msgbytes)
		return
	}

	o, err := paymentintent.Update(id, param)
	if err != nil {
		msgbytes, _ := json.Marshal(err)
		writeJsonResponse(w, msgbytes)
		return
	}
	bytes, _ := json.Marshal(o)
	writeJsonResponse(w, bytes)
}

//Confirm PaymentIntent
func ConfirmPaymentIntent(w http.ResponseWriter, r *http.Request) {

	stripe.Key = os.Getenv("SECRET_KEY")

	params := &stripe.PaymentIntentConfirmParams{
		Source: stripe.String("src_1E5PV5JytX7n0OoXVhMXhy4N"),
	}
	intent, err := paymentintent.Confirm("pi_1E5Q7FJytX7n0OoXaCfqt9KD", params)
	if err != nil {
		msgbytes, _ := json.Marshal(err)
		writeJsonResponse(w, msgbytes)
		return
	}
	bytes, _ := json.Marshal(intent)
	writeJsonResponse(w, bytes)
}

//Capture PaymentIntent
func CapturePaymentIntent(w http.ResponseWriter, r *http.Request) {
	stripe.Key = os.Getenv("SECRET_KEY")
	vars := mux.Vars(r)
	var id = vars["paymentintent_id"]

	decoderpi := json.NewDecoder(r.Body)
	var param *stripe.PaymentIntentCaptureParams
	errr := decoderpi.Decode(&param)
	if errr != nil {
		msgbytes, _ := json.Marshal(errr)
		writeJsonResponse(w, msgbytes)
		return
	}

	intent, err := paymentintent.Capture(id, param)
	if err != nil {
		msgbytes, _ := json.Marshal(err)
		writeJsonResponse(w, msgbytes)
		return
	}
	bytes, _ := json.Marshal(intent)
	writeJsonResponse(w, bytes)
}

//Cancel PaymentIntent
func CancelPaymentIntent(w http.ResponseWriter, r *http.Request) {
	stripe.Key = os.Getenv("SECRET_KEY")
	vars := mux.Vars(r)
	var id = vars["paymentintent_id"]

	intent, err := paymentintent.Cancel(id, nil)
	if err != nil {
		log.Fatalln(err)
	}
	bytes, _ := json.Marshal(intent)
	writeJsonResponse(w, bytes)
}

//List All PaymentIntent
func ListAllPaymentIntent(w http.ResponseWriter, r *http.Request) {
	stripe.Key = os.Getenv("SECRET_KEY")

	params := &stripe.PaymentIntentListParams{}
	params.Filters.AddFilter("limit", "", "3")
	i := paymentintent.List(params)
	for i.Next() {
		p := i.PaymentIntent()
		fmt.Println(p)
		bytes, _ := json.Marshal(p)
		writeJsonResponse(w, bytes)
	}
}

func writeJsonResponse(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bytes)
}
