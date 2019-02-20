package SourceOperation

import (
	"encoding/json"
	"github.com/stripe/stripe-go"
	//"github.com/stripe/stripe-go/paymentsource"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stripe/stripe-go/source"
	"net/http"
	"os"
)

//Create Source
func CreateSource(w http.ResponseWriter, r *http.Request) {

	stripe.Key = os.Getenv("SECRET_KEY")

	vars := mux.Vars(r)
	var tranftype = vars["transfertype"]

	decoder := json.NewDecoder(r.Body)
	var param *stripe.SourceObjectParams
	err := decoder.Decode(&param)
	fmt.Println("param : ", param)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}

	param.Type = stripe.String(tranftype)

	s, errr := source.New(param)
	if errr != nil {
		WriteErrorResponse(w, err)
		return
	}
	fmt.Println(s)
	bytes, _ := json.Marshal(s)
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
