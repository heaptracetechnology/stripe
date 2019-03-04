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

	card := Card{}
	reqbody := new(bytes.Buffer)
	json.NewEncoder(reqbody).Encode(card)

	req, err := http.NewRequest("POST", "/createcard", reqbody)
	if err != nil {
	}
	recorder := httptest.NewRecorder()

	Describe("create card", func() {
		Context("CreateCard", func() {
			It("Should result http.StatusCreated", func() {
				CreateCard(recorder, req)
				Expect(result.GetResultCreated()).To(Equal(http.StatusCreated))
			})
		})
	})
})
