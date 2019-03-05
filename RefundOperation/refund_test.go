package RefundOperation

import (
	"bytes"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"os"
)

type Refund struct {
	ChargeId string `json:"charge"`
}

var _ = Describe("Create Refund operations", func() {

	refund := Refund{ChargeId: "ch_1EAIJQJytX7n0OoXhSP6RTc9"}
	reqbody := new(bytes.Buffer)
	json.NewEncoder(reqbody).Encode(refund)

	os.Setenv("SECRET_KEY", "sk_test_gENQu8ecxwwMUsWlgsQeqbgI")
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
