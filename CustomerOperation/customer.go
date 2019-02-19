package CustomerOperation

import (
	"encoding/json"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"log"
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
		msgbytes, _ := json.Marshal(err)
		writeJsonResponse(w, msgbytes)
		return
	}

	param.SetSource("tok_amex")

	cus, err := customer.New(param)
	if err != nil {
		log.Fatalln(err)
	}
	bytes, _ := json.Marshal(cus)
	writeJsonResponse(w, bytes)
}

func writeJsonResponse(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bytes)
}
