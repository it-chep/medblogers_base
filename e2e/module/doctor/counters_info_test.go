package doctor

import (
	"context"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"medblogers_base/e2e/helper"
	"medblogers_base/e2e/module/doctor/app"
	"medblogers_base/e2e/module/doctor/fixture"
	"medblogers_base/e2e/shared_behavior"
	"medblogers_base/internal/modules/doctors/domain/doctor"
)

var _ = Describe("Пользователь запрашивает список количество докторов и подписчиков", Label("rega", "counters_info"), func() {
	var module *app.TestableModule

	BeforeEach(func(ctx context.Context) {
		module = fixture.SetupModule(ctx, conn)
		shared_behavior.ExpectGetAllSubscribersInfo(module.Http, "123", "подписчика", "11.05.2004")
		helper.AddNewDoctor(ctx, module.PoolWrapper, doctor.New(
			doctor.WithName("Инокентий"),
			doctor.WithSlug("inokenty"),
		))
	})

	It("Успешное получение количества докторов и подписчиков", func(ctx context.Context) {
		By("выполнение экшена")

		countersInfo, err := module.Module.Actions.CounterInfo.Do(ctx)

		Expect(err).NotTo(HaveOccurred())
		Expect(countersInfo.DoctorsCount).To(BeNumerically(">", 0))
		Expect(countersInfo.SubscribersCount).To(Equal("123"))
	})
})
