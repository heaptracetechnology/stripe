package CustomerOperation

import (
	"encoding/json"
	"github.com/heaptracetechnology/microservice-stripe/result"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"net/http"
	"os"
)

//Create Cusytomer
func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	stripe.Key = os.Getenv("SECRET_KEY")

	decoder := json.NewDecoder(r.Body)
	var param *stripe.CustomerParams
	err := decoder.Decode(&param)
	if err != nil {
		result.WriteErrorResponse(w, err)
	}

	param.SetSource("tok_amex")

	cus, err := customer.New(param)
	if err != nil {
		result.WriteErrorResponse(w, err)
	}
	bytes, _ := json.Marshal(cus)
	result.WriteJsonResponse(w, bytes, http.StatusCreated)
}
