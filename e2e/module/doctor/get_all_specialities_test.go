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

var _ = Describe("Пользователь запрашивает список доступных специальностей для регистрации", Label("rega", "get_all_specialities_test"), func() {
	var module *app.TestableModule

	BeforeEach(func(ctx context.Context) {
		module = fixture.SetupModule(ctx, conn)
		helper.AddSpeciality(ctx, conn, "Хирург")
		helper.AddSpeciality(ctx, conn, "Акушер-гинеколог")
		helper.AddSpeciality(ctx, conn, "Терапевт")
	})

	It("Успешное получение специальностей", func(ctx context.Context) {
		specialities, err := module.Module.Actions.AllSpecialities.Do(ctx)
		Expect(err).NotTo(HaveOccurred())
		Expect(len(specialities)).To(BeNumerically(">", 0))
	})
})
