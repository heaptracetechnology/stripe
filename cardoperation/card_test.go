package cardoperation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	stripe "github.com/stripe/stripe-go"
)

var secretKey = os.Getenv("STRIPE_SECRET_KEY")

var _ = Describe("Create Card operations", func() {

	var param stripe.CardParams
	param.Customer = stripe.String("mockCustomerID")

	reqbody := new(bytes.Buffer)
	json.NewEncoder(reqbody).Encode(param)

	os.Setenv("SECRET_KEY", secretKey)

	req, err := http.NewRequest("POST", "/createCard", reqbody)
	if err != nil {
		fmt.Println("err test : ", err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCard)
	handler.ServeHTTP(recorder, req)

	Describe("create card", func() {
		Context("CreateCard", func() {
			It("Should result http.StatusCreated", func() {
				Expect(recorder.Code).To(Equal(http.StatusCreated))
			})
		})
	})

})
