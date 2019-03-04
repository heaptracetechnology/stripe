package BalanceOperation

import (
	"github.com/heaptracetechnology/microservice-stripe/result"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Get Balance operations", func() {

	req, err := http.NewRequest("GET", "/getbalance", nil)
	if err != nil {
	}
	recorder := httptest.NewRecorder()

	Describe("get balance", func() {
		Context("GetBalance", func() {
			It("Should result http.StatusCreated", func() {
				GetBalance(recorder, req)
				Expect(result.GetResultOK()).To(Equal(http.StatusOK))
			})
		})
	})
})
