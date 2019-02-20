package PaymentIntentOperation

import (
	"encoding/json"
	"fmt"
	//assert "github.com/stretchr/testify/require"
	"github.com/stripe/stripe-go"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

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
		fmt.Println("err ::: ", err)
		t.Errorf("handler returned unexpected body: got %v want %v",
			recorder.Body.String(), listpaymentintent)
	}

}
