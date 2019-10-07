package customeroperation

import (
	"testing"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
)

func TestCard(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("../test-report/cireport.txt")
	RunSpecsWithDefaultAndCustomReporters(t, "Customer Operations", []Reporter{junitReporter})
}
