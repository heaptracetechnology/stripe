package balanceoperation

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/heaptracetechnology/microservice-stripe/result"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/balance"
)

//GetBalance stripe
func GetBalance(responseWriter http.ResponseWriter, request *http.Request) {

	stripe.Key = os.Getenv("SECRET_KEY")
	bal, err := balance.Get(nil)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
	}
	bytes, _ := json.Marshal(bal)
	responseWriter.WriteHeader(http.StatusOK)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}
