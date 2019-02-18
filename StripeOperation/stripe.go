package StripeOperation

import (
	"encoding/json"
	"fmt"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/balance"
	"github.com/stripe/stripe-go/charge"
	"github.com/stripe/stripe-go/paymentintent"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type ChargeObject struct {
	AmountCh      int64  `json:"amount"`
	CurrencyCh    string `json:"currency"`
	DescriptionCh string `json:"description"`
}

//CreateCharge
func CreateCharge(w http.ResponseWriter, r *http.Request) {
	stripe.Key = os.Getenv("SECRET_KEY")

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var chargeobj ChargeObject
	err = json.Unmarshal(body, &chargeobj)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	chargeParams := &stripe.ChargeParams{
		Amount:      stripe.Int64(chargeobj.AmountCh),
		Currency:    stripe.String(string(stripe.CurrencyUSD)),
		Description: stripe.String(chargeobj.DescriptionCh),
	}

	chargeParams.SetSource("tok_visa")
	ch, err := charge.New(chargeParams)
	if err != nil {
		log.Fatalln(err)
	}
	bytes, _ := json.Marshal(ch)
	writeJsonResponse(w, bytes)
}

//GetBalance
func GetBalance(w http.ResponseWriter, r *http.Request) {

	stripe.Key = os.Getenv("SECRET_KEY")

	b, err := balance.Get(nil)
	if err != nil {
		log.Fatalln(err)
		fmt.Println(err)
	}
	bytes, _ := json.Marshal(b)
	writeJsonResponse(w, bytes)
}

//Create PaymentIntent
func CreatePaymentIntent(w http.ResponseWriter, r *http.Request) {

	stripe.Key = os.Getenv("SECRET_KEY")

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var chargeobj ChargeObject
	err = json.Unmarshal(body, &chargeobj)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(chargeobj.AmountCh),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		PaymentMethodTypes: []*string{
			stripe.String("card"),
		},
	}
	pi, err := paymentintent.New(params)
	if err != nil {
		log.Fatalln(err)
	}
	bytes, _ := json.Marshal(pi)
	writeJsonResponse(w, bytes)
}

func writeJsonResponse(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bytes)
}
