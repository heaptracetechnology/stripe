package CardOperation

import (
	"encoding/json"
	"github.com/heaptracetechnology/microservice-stripe/result"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/card"
	"net/http"
	"os"
)

func CreateCard(w http.ResponseWriter, r *http.Request) {
	stripe.Key = os.Getenv("SECRET_KEY")

	decoder := json.NewDecoder(r.Body)
	var param *stripe.CardParams
	err := decoder.Decode(&param)
	if err != nil {
		result.WriteErrorResponse(w, err)
	}

	card, err := card.New(param)
	if err != nil {
		result.WriteErrorResponse(w, err)
	}
	bytes, _ := json.Marshal(card)
	result.WriteJsonResponse(w, bytes, http.StatusCreated)
}
