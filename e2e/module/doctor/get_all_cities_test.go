//go:build e2e
// +build e2e

package doctor

import (
	"context"
	"medblogers_base/e2e/helper"
	"medblogers_base/e2e/module/doctor/app"
	"medblogers_base/e2e/module/doctor/fixture"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Пользователь запрашивает список доступных городов для регистрации", Label("rega", "get_all_cities"), func() {
	var module *app.TestableModule

	BeforeEach(func(ctx context.Context) {
		module = fixture.SetupModule(ctx, conn)
		helper.AddCity(ctx, conn, "Москва")
		helper.AddCity(ctx, conn, "Санкт-петербург")
		helper.AddCity(ctx, conn, "Екатеринбург")
	})

	It("Успешное получение городов", func(ctx context.Context) {
		cities, err := module.Module.Actions.AllCities.Do(ctx)
		Expect(err).NotTo(HaveOccurred())
		Expect(len(cities)).To(BeNumerically(">", 0))
	})
})
