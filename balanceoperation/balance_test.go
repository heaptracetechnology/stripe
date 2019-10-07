package balanceoperation

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var secretKey = os.Getenv("STRIPE_SECRET_KEY")

var _ = Describe("Get Balance operations", func() {

	os.Setenv("SECRET_KEY", secretKey)
	req, err := http.NewRequest("GET", "/getBalance", nil)
	if err != nil {
		log.Fatal(err)
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
