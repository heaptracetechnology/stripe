package refundoperation

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type Refund struct {
	ChargeId string `json:"charge"`
}

var secretKey = os.Getenv("STRIPE_SECRET_KEY")

var _ = Describe("Create Refund operations", func() {

	refund := Refund{ChargeId: "ch_1EAIJQJytX7n0OoXhSP6RTc9"}
	reqbody := new(bytes.Buffer)
	json.NewEncoder(reqbody).Encode(refund)

	os.Setenv("SECRET_KEY", secretKey)
	req, err := http.NewRequest("POST", "/createrefund", reqbody)
	if err != nil {
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateRefund)
	handler.ServeHTTP(recorder, req)

	Describe("create refund", func() {
		Context("CreateRefund", func() {
			It("Should result http.StatusBadRequest already refunded", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})
