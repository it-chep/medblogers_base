//go:build e2e
// +build e2e

package doctor

import (
	"context"
	"medblogers_base/e2e/module/doctor/app"
	"medblogers_base/e2e/module/doctor/fixture"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Пользователь запрашивает список доступных городов для регистрации", Label("rega", "get_all_cities"), func() {
	var module *app.TestableModule

	BeforeEach(func(ctx context.Context) {
		module = fixture.SetupModule(ctx, conn)
	})

	It("Успешное получение городов", func(ctx context.Context) {
		cities, err := module.Module.Actions.AllCities.Do(ctx)
		Expect(err).NotTo(HaveOccurred())
		Expect(len(cities) > 0)
	})
})
