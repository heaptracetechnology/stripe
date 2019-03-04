package RefundOperation

import (
	"bytes"
	"encoding/json"
	"github.com/heaptracetechnology/microservice-stripe/result"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

type Refund struct {
	ChargeId string `json:"charge"`
}

var _ = Describe("Create Refund operations", func() {

	refund := Refund{}
	reqbody := new(bytes.Buffer)
	json.NewEncoder(reqbody).Encode(refund)

	req, err := http.NewRequest("POST", "/createrefund", reqbody)
	if err != nil {
	}
	recorder := httptest.NewRecorder()

	Describe("create refund", func() {
		Context("CreateRefund", func() {
			It("Should result http.StatusCreated", func() {
				CreateRefund(recorder, req)
				Expect(result.GetResultCreated()).To(Equal(http.StatusCreated))
			})
		})
	})
})
