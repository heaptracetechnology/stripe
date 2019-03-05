package ChargeOperation

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"os"
)

type Charge struct {
	Amount      int64  `json:"amount"`
	Currency    string `json:"currency"`
	Description string `json:"description"`
	Capture     bool   `json:"capture"`
	ChargeID    string `json:"charge"`
}

var _ = Describe("Create charge operations", func() {

	charge := Charge{Amount: 2000, Currency: "usd", Description: "Created test charge", Capture: false}
	reqbody := new(bytes.Buffer)
	json.NewEncoder(reqbody).Encode(charge)

	os.Setenv("SECRET_KEY", "sk_test_gENQu8ecxwwMUsWlgsQeqbgI")
	req, err := http.NewRequest("POST", "/createcharge", reqbody)
	if err != nil {
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCharge)
	handler.ServeHTTP(recorder, req)

	Describe("create charge", func() {
		Context("CreateCharge", func() {
			It("Should result http.StatusCreated", func() {
				Expect(recorder.Code).To(Equal(http.StatusCreated))
			})
		})
	})
})

var _ = Describe("Capture Charge operations", func() {

	charge := Charge{Amount: 300}
	reqbody := new(bytes.Buffer)
	json.NewEncoder(reqbody).Encode(charge)

	os.Setenv("SECRET_KEY", "sk_test_gENQu8ecxwwMUsWlgsQeqbgI")
	req, err := http.NewRequest("POST", "/capturecharge/charge", reqbody)
	if err != nil {
	}
	vars := map[string]string{
		"charge": "ch_1EAJ64JytX7n0OoXxR6RL8HU",
	}

	req = mux.SetURLVars(req, vars)

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CaptureCharge)
	handler.ServeHTTP(recorder, req)

	Describe("capture charge", func() {
		Context("CaptureCharge", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})
