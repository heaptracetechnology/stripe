package customeroperation

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/heaptracetechnology/microservice-stripe/result"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
)

//CreateCustomer stripe
func CreateCustomer(responseWriter http.ResponseWriter, request *http.Request) {

	stripe.Key = os.Getenv("SECRET_KEY")

	decoder := json.NewDecoder(request.Body)
	var param *stripe.CustomerParams
	err := decoder.Decode(&param)
	if err != nil {
		result.WriteErrorResponse(nil, err)
	}
	sourceErr := param.SetSource("tok_amex")
	if sourceErr != nil {
		result.WriteErrorResponse(nil, sourceErr)
	}

	customer, err := customer.New(param)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}
	bytes, _ := json.Marshal(customer)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusCreated)

}
