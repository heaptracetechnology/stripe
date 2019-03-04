package CustomerOperation

import (
	"bytes"
	"encoding/json"
	"github.com/heaptracetechnology/microservice-stripe/result"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

type Customer struct {
	Email       string `json:"email"`
	Description string `json:"description"`
}

var _ = Describe("Create Customer operations", func() {

	customer := Customer{}
	reqbody := new(bytes.Buffer)
	json.NewEncoder(reqbody).Encode(customer)

	req, err := http.NewRequest("POST", "/createcustomer", reqbody)
	if err != nil {
	}
	recorder := httptest.NewRecorder()

	Describe("create customer", func() {
		Context("CreateCustomer", func() {
			It("Should result http.StatusCreated", func() {
				CreateCustomer(recorder, req)
				Expect(result.GetResultCreated()).To(Equal(http.StatusCreated))
			})
		})
	})
})
