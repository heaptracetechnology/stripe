package PaymentIntentOperation

import (
	"bytes"
	"encoding/json"
	//"fmt"
	"github.com/gorilla/mux"
	"github.com/stripe/stripe-go"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type PaymentIntent1 struct {
	Amount              int64
	Currency            string
	Payment_Method_Type string
}

func TestListAllPaymentIntent(t *testing.T) {

	req, err := http.NewRequest("GET", "/listallpaymentintent", nil)
	if err != nil {
		t.Fatal(err)
	}
	os.Setenv("SECRET_KEY", "sk_test_gENQu8ecxwwMUsWlgsQeqbgI")
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ListAllPaymentIntent)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var listpaymentintent []stripe.PaymentIntent
	data := recorder.Body.String()
	dataByte := []byte(data)
	err = json.Unmarshal(dataByte, &listpaymentintent)

	if err != nil {
		t.Errorf("handler returned unexpected body: got %v want %v",
			recorder.Body.String(), listpaymentintent)
	}

}

func TestCreatePaymentIntent(t *testing.T) {
	paymentintent := PaymentIntent1{Amount: 5600, Currency: "usd", Payment_Method_Type: "card"}
	reqbody := new(bytes.Buffer)
	json.NewEncoder(reqbody).Encode(paymentintent)

	req, err := http.NewRequest("POST", "/createpaymentintent", reqbody)
	if err != nil {
		t.Fatal(err)
	}
	os.Setenv("SECRET_KEY", "sk_test_gENQu8ecxwwMUsWlgsQeqbgI")
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreatePaymentIntent)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestRetrievePaymentIntent(t *testing.T) {

	req, err := http.NewRequest("GET", "/retrievepaymentIntent/pi_1E5WMKJytX7n0OoXQ0o9SaLy", nil)
	if err != nil {
		t.Fatal(err)
	}

	vars := map[string]string{
		"paymentintent_id": "pi_1E5WMKJytX7n0OoXQ0o9SaLy",
	}

	req = mux.SetURLVars(req, vars)
	os.Setenv("SECRET_KEY", "sk_test_gENQu8ecxwwMUsWlgsQeqbgI")
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(RetrievePaymentIntent)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestUpdatePaymentIntent(t *testing.T) {

	paymentintent := PaymentIntent1{Amount: 6000, Currency: "usd"}
	reqbody := new(bytes.Buffer)
	json.NewEncoder(reqbody).Encode(paymentintent)

	req, err := http.NewRequest("PUT", "/updatepaymentintent/pi_1E6EC3JytX7n0OoXi0RHZfoe", reqbody)
	if err != nil {
		t.Fatal(err)
	}
	vars := map[string]string{
		"paymentintent_id": "pi_1E6EC3JytX7n0OoXi0RHZfoe",
	}

	req = mux.SetURLVars(req, vars)
	os.Setenv("SECRET_KEY", "sk_test_gENQu8ecxwwMUsWlgsQeqbgI")
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdatePaymentIntent)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestCancelPaymentIntent(t *testing.T) {

	req, err := http.NewRequest("POST", "/cancelpaymentintent/pi_1E6E79JytX7n0OoXpVFtXPzR", nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := map[string]string{
		"paymentintent_id": "pi_1E6E79JytX7n0OoXpVFtXPzR",
	}

	req = mux.SetURLVars(req, vars)
	os.Setenv("SECRET_KEY", "sk_test_gENQu8ecxwwMUsWlgsQeqbgI")
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CancelPaymentIntent)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
