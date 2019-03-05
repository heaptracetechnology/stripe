package CustomerOperation

import (
	"bytes"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"os"
)

type Customer struct {
	Email       string `json:"email"`
	Description string `json:"description"`
}

var _ = Describe("Create Customer operations", func() {

	customer := Customer{Email: "testcust@demo.com", Description: "Test Demo Customer"}
	reqbody := new(bytes.Buffer)
	json.NewEncoder(reqbody).Encode(customer)

	os.Setenv("SECRET_KEY", "sk_test_gENQu8ecxwwMUsWlgsQeqbgI")

	req, err := http.NewRequest("POST", "/createcustomer", reqbody)
	if err != nil {
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCustomer)
	handler.ServeHTTP(recorder, req)

	Describe("create customer", func() {
		Context("CreateCustomer", func() {
			It("Should result http.StatusCreated", func() {
				//CreateCustomer(recorder, req)
				Expect(recorder.Code).To(Equal(http.StatusCreated))
			})
		})
	})
})
