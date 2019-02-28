package BalanceOperation

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetBalance(t *testing.T) {

	req, err := http.NewRequest("GET", "/getbalance", nil)
	if err != nil {
		t.Fatal(err)
	}
	os.Setenv("SECRET_KEY", "sk_test_gENQu8ecxwwMUsWlgsQeqbgI")
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetBalance)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
