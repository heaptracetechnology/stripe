package CardOperation

import (
	"bytes"
	"encoding/json"
	"github.com/heaptracetechnology/microservice-stripe/result"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

type Card struct {
	Source   interface{} `json:"source"`
	Customer string      `json:"customer"`
	Metadata interface{} `json:"metadata"`
}

var _ = Describe("Create Card operations", func() {
	os.Setenv("SECRET_KEY", "sk_test_Ma5sDiK8tRrvg1eHVew4SYEN")
	var param stripe.CardParams
	reqbody := new(bytes.Buffer)
	json.NewEncoder(reqbody).Encode(param)

	req, err := http.NewRequest("POST", "/createcard", reqbody)
	if err != nil {
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCard)
	handler.ServeHTTP(recorder, req)
	
	Describe("create card", func() {
		Context("CreateCard", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})
