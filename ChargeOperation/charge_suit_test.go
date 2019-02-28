package ChargeOperation

import (
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
	"testing"
)

func TestCharge(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("../test-report/cireport.txt")
	RunSpecsWithDefaultAndCustomReporters(t, "Charge Operations", []Reporter{junitReporter})
}
