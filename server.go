package main

import (
	"fmt"
	"github.com/user/stripe/route"
	"log"
	"net/http"
)

func main() {
	fmt.Println("in server-go")
	router := route.NewRouter()
	log.Fatal(http.ListenAndServe(":3000", router))
}
