package PaymentIntentOperation

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

type PaymentIntent struct {
	Amount             int64  `json:"amount"`
	Currency           string `json:"currency"`
	Paymentmethodtypes string `json:"paymentmethodtypes"`
}

var _ = Describe("Create PaymentIntent operations", func() {

	paymentintent := PaymentIntent{}
	reqbody := new(bytes.Buffer)
	json.NewEncoder(reqbody).Encode(paymentintent)
	os.Setenv("SECRET_KEY", "sk_test_gENQu8ecxwwMUsWlgsQeqbgI")
	req, err := http.NewRequest("POST", "/CreatePaymentIntent", reqbody)
	if err != nil {
	}
	recorder := httptest.NewRecorder()

	Describe("create payment intent", func() {
		Context("CreatePaymentIntent", func() {
			It("Should result http.StatusCreated", func() {
				CreatePaymentIntent(recorder, req)
				Expect(GetResultCreated()).To(Equal(http.StatusCreated))
			})
		})
	})
})

var _ = Describe("Retrive PaymentIntent operations", func() {

	req, err := http.NewRequest("GET", "/retrievepaymentIntent/paymentintent_id", nil)
	if err != nil {
	}
	vars := map[string]string{
		//"paymentintent_id": "pi_1E8p2MJytX7n0OoX1uYVO0Fz",
		"paymentintent_id": "",
	}

	req = mux.SetURLVars(req, vars)
	recorder := httptest.NewRecorder()

	Describe("retrive payment intent", func() {
		Context("RetrivePaymentIntent", func() {
			It("Should result http.StatusOK", func() {
				RetrievePaymentIntent(recorder, req)
				Expect(GetResultOK()).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("Update PaymentIntent operations", func() {

	paymentintent := PaymentIntent{}
	reqbody := new(bytes.Buffer)
	json.NewEncoder(reqbody).Encode(paymentintent)
	req, err := http.NewRequest("PUT", "/updatepaymentintent/paymentintent_id", reqbody)
	if err != nil {
	}
	vars := map[string]string{
		//"paymentintent_id": "pi_1E8p2MJytX7n0OoX1uYVO0Fz",
		"paymentintent_id": "",
	}

	req = mux.SetURLVars(req, vars)
	recorder := httptest.NewRecorder()

	Describe("update payment intent", func() {
		Context("UpdatePaymentIntent", func() {
			It("Should result http.StatusOK", func() {
				UpdatePaymentIntent(recorder, req)
				Expect(GetResultOK()).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("Capture PaymentIntent operations", func() {

	paymentintent := PaymentIntent{}
	reqbody := new(bytes.Buffer)
	json.NewEncoder(reqbody).Encode(paymentintent)
	req, err := http.NewRequest("PUT", "/capturecharge/paymentintent_id", reqbody)
	if err != nil {
	}
	vars := map[string]string{
		//"paymentintent_id": "pi_1E8p2MJytX7n0OoX1uYVO0Fz",
		"paymentintent_id": "",
	}

	req = mux.SetURLVars(req, vars)
	recorder := httptest.NewRecorder()

	Describe("capture payment intent", func() {
		Context("CapturePaymentIntent", func() {
			It("Should result http.StatusOK", func() {
				CapturePaymentIntent(recorder, req)
				Expect(GetResultOK()).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("Cancel PaymentIntent operations", func() {

	paymentintent := PaymentIntent{}
	reqbody := new(bytes.Buffer)
	json.NewEncoder(reqbody).Encode(paymentintent)
	req, err := http.NewRequest("POST", "/cancelpaymentintent/paymentintent_id", nil)
	if err != nil {
	}
	vars := map[string]string{
		//"paymentintent_id": "pi_1E8p2MJytX7n0OoX1uYVO0Fz",
		"paymentintent_id": "",
	}

	req = mux.SetURLVars(req, vars)
	recorder := httptest.NewRecorder()

	Describe("cancel payment intent", func() {
		Context("CancelPaymentIntent", func() {
			It("Should result http.StatusOK", func() {
				CancelPaymentIntent(recorder, req)
				Expect(GetResultOK()).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("List All PaymentIntent operations", func() {

	req, err := http.NewRequest("GET", "/listallpaymentintent", nil)
	if err != nil {
	}

	recorder := httptest.NewRecorder()

	Describe("list all payment intent", func() {
		Context("ListAllPaymentIntent", func() {
			It("Should result http.StatusOK", func() {
				ListAllPaymentIntent(recorder, req)
				Expect(GetResultOK()).To(Equal(http.StatusOK))
			})
		})
	})
})
