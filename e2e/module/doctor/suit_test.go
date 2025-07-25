//go:build e2e
// +build e2e

package doctor

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestShiftModule(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Тестирование модуля докторов")
}

var (
	poolWrapper poolWrapper
)

var _ = BeforeSuite(func(ctx SpecContext) {
	commonfixture.SetupDatabase(ctx, 6)
	poolWrapper = commonfixture.SetupConnection(ctx, 6)
})

var _ = AfterSuite(func() {
	if poolWrapper != nil {
		poolWrapper.Close()
	}
})
