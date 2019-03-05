package BalanceOperation

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"os"
)

var _ = Describe("Get Balance operations", func() {

	os.Setenv("SECRET_KEY", "sk_test_gENQu8ecxwwMUsWlgsQeqbgI")
	req, err := http.NewRequest("GET", "/getbalance", nil)
	if err != nil {
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetBalance)
	handler.ServeHTTP(recorder, req)

	Describe("get balance", func() {
		Context("GetBalance", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})
