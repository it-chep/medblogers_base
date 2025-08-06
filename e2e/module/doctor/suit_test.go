//go:build e2e
// +build e2e

package doctor

import (
	"medblogers_base/internal/pkg/postgres"
	"testing"

	commonfixture "medblogers_base/e2e/fixture"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestShiftModule(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Тестирование модуля докторов")
}

var conn postgres.PoolWrapper

var _ = BeforeSuite(func(ctx SpecContext) {
	commonfixture.SetupDatabase(ctx, 6)
	conn = commonfixture.SetupPoolConnections(ctx, 6)
})

var _ = AfterSuite(func() {
	// todo
})
