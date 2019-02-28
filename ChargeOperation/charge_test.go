package ChargeOperation

import (
	"bytes"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

type Charge struct {
	Amount      int64  `json:"amount"`
	Currency    string `json:"currency"`
	Description string `json:"description"`
}

var _ = Describe("Create Card operations", func() {

	charge := Charge{}
	reqbody := new(bytes.Buffer)
	json.NewEncoder(reqbody).Encode(charge)

	req, err := http.NewRequest("POST", "/createcharge", reqbody)
	if err != nil {
	}
	recorder := httptest.NewRecorder()

	Describe("create charge", func() {
		Context("CreateCharge", func() {
			It("Should result http.StatusCreated", func() {
				CreateCharge(recorder, req)
				Expect(GetResultCreated()).To(Equal(http.StatusCreated))
			})
		})
	})
})

var _ = Describe("Capture Charge operations", func() {

	charge := Charge{}
	reqbody := new(bytes.Buffer)
	json.NewEncoder(reqbody).Encode(charge)

	req, err := http.NewRequest("POST", "/capturecharge", reqbody)
	if err != nil {
	}
	recorder := httptest.NewRecorder()

	Describe("capture charge", func() {
		Context("CaptureCharge", func() {
			It("Should result http.StatusOK", func() {
				CaptureCharge(recorder, req)
				Expect(GetResultOK()).To(Equal(http.StatusOK))
			})
		})
	})
})
