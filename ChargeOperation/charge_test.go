package ChargeOperation

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type Charge struct {
	Amount      int64
	Currency    string
	Description string
}

func TestCreateCustomer(t *testing.T) {

	charge := Charge{Amount: 5600, Currency: "usd", Description: "5600 charge for testing@testmail.com"}
	reqbody := new(bytes.Buffer)
	json.NewEncoder(reqbody).Encode(charge)

	req, err := http.NewRequest("POST", "/createcharge", reqbody)
	if err != nil {
		t.Fatal(err)
	}
	os.Setenv("SECRET_KEY", "sk_test_gENQu8ecxwwMUsWlgsQeqbgI")
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCharge)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
