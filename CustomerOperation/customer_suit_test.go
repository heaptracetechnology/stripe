package CustomerOperation

import (
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
	"testing"
)

func TestCard(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("../test-report/cireport.txt")
	RunSpecsWithDefaultAndCustomReporters(t, "Customer Operations", []Reporter{junitReporter})
}
