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
		msgbytes, _ := json.Marshal(err)
		writeJsonResponse(w, msgbytes)
		return
	}
	// if err != nil {
	// 	msgbytes, _ := json.Marshal(err)
	// 	writeJsonResponse(w, msgbytes)
	// 	return
	// }

	param.Type = stripe.String(tranftype)

	s, errr := source.New(param)
	if errr != nil {
		msgbytes, _ := json.Marshal(err)
		writeJsonResponse(w, msgbytes)
		return
	}
	fmt.Println(s)
	bytes, _ := json.Marshal(s)
	writeJsonResponse(w, bytes)

}

// //Attach Source
// func AttachSource(w http.ResponseWriter, r *http.Request) {
// 	stripe.Key = "sk_test_gENQu8ecxwwMUsWlgsQeqbgI"

// 	params := &stripe.CustomerSourceParams{
// 		Customer: stripe.String("cus_EYvkRauOZ0rXXO"),
// 		Source: &stripe.SourceParams{
// 			Token: stripe.String("src_1E5mUKJytX7n0OoXQoNS58Q9"),
// 		},
// 	}
// 	s, err := paymentsource.New(params)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	bytes, _ := json.Marshal(s)
// 	writeJsonResponse(w, bytes)
// }

func writeJsonResponse(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bytes)
}
