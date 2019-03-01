package CustomerOperation

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type Customer1 struct {
	Email       string
	Description string
}

func TestCreateCustomer(t *testing.T) {

	customer := Customer1{Email: "testing3@testmail.com", Description: "test case new customer"}
	reqbody := new(bytes.Buffer)
	json.NewEncoder(reqbody).Encode(customer)

	req, err := http.NewRequest("POST", "/createcustomer", reqbody)
	if err != nil {
		t.Fatal(err)
	}
	os.Setenv("SECRET_KEY", "sk_test_gENQu8ecxwwMUsWlgsQeqbgI")
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCustomer)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
