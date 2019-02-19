package SourceOperation

import (
	"encoding/json"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/source"
	"log"
	"net/http"
	"os"
)

//Create Source
func CreateSource(w http.ResponseWriter, r *http.Request) {

	stripe.Key = os.Getenv("SECRET_KEY")

	decoder := json.NewDecoder(r.Body)
	var param *stripe.SourceObjectParams
	err := decoder.Decode(&param)
	if err != nil {
		msgbytes, _ := json.Marshal(err)
		writeJsonResponse(w, msgbytes)
		return
	}

	params := &stripe.SourceObjectParams{
		Type:     stripe.String("ach_credit_transfer"),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Owner: &stripe.SourceOwnerParams{
			Email: stripe.String("jenny.rosen@example.com"),
		},
	}
	s, err := source.New(params)
	if err != nil {
		log.Fatalln(err)
	}
	bytes, _ := json.Marshal(s)
	writeJsonResponse(w, bytes)

}

func writeJsonResponse(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bytes)
}
